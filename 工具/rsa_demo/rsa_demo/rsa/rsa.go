package rsa

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"time"

	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"math/big"
)

const (
	MODE_PUBKEY_ENCRYPT = iota //公钥加密
	MODE_PUBKEY_DECRYPT        //公钥解密
	MODE_PRIKEY_ENCRYPT        //私钥加密
	MODE_PRIKEY_DECRYPT        //私钥解密
)

var (
	RSA = &rSASecurity{}
)

type rSASecurity struct {
	isFileInit bool            //true: 从文件读取密钥钥, false:字符串密钥
	ifCache    bool            //isFileInit＝＝true时有效。 true:只在初始化时读取密钥, false:每次都从文件读取
	pubStr     string          //isFileInit＝=true:公钥文件路径， isFileInit＝=false:公钥字符串
	priStr     string          //isFileInit＝=true:私钥文件路径， isFileInit＝=false:私钥字符串
	pubkey     *rsa.PublicKey  //公钥
	prikey     *rsa.PrivateKey //私钥
	pubModTime time.Time       //isFileInit＝＝true&&ifCache＝＝false时有效, 公钥文件最后的修改时间
	priModTime time.Time       //isFileInit＝＝true&&ifCache＝＝false时有效, 私钥文件最后的修改时间
}

func (this *rSASecurity) String(in string, mode int) (string, error) {
	var inByte []byte
	var err error
	if mode == MODE_PRIKEY_ENCRYPT || mode == MODE_PUBKEY_ENCRYPT {
		inByte = []byte(in)
	} else if mode == MODE_PRIKEY_DECRYPT || mode == MODE_PUBKEY_DECRYPT {
		inByte, err = base64.StdEncoding.DecodeString(in)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("mode not found")
	}
	inByte, err = this.Byte(inByte, mode)
	if err != nil {
		return "", err
	}
	if mode == MODE_PRIKEY_ENCRYPT || mode == MODE_PUBKEY_ENCRYPT {
		return base64.StdEncoding.EncodeToString(inByte), nil
	} else {
		return string(inByte), nil
	}
}

func (this *rSASecurity) Byte(in []byte, mode int) ([]byte, error) {
	out := bytes.NewBuffer(nil)
	err := this.IO(bytes.NewReader(in), out, mode)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(out)
}

func (this *rSASecurity) File(srcPath, distPath string, mode int) error {
	in, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer in.Close()
	isSuccess := false
	defer func() {
		if !isSuccess {
			os.Remove(distPath)
		}
	}()
	out, err := os.OpenFile(distPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0664)
	if err != nil {
		return err
	}
	defer out.Close()
	err = this.IO(in, out, mode)
	if err == nil {
		isSuccess = true
	}
	return err
}

func (this *rSASecurity) IO(in io.Reader, out io.Writer, mode int) error {
	switch mode {
	case MODE_PUBKEY_ENCRYPT:
		if key, err := this.getPubKey(); err != nil {
			return err
		} else {
			return pubKeyIO(key, in, out, true)
		}
	case MODE_PUBKEY_DECRYPT:
		if key, err := this.getPubKey(); err != nil {
			return err
		} else {
			return pubKeyIO(key, in, out, false)
		}
	case MODE_PRIKEY_ENCRYPT:
		if key, err := this.getPriKey(); err != nil {
			return err
		} else {
			return priKeyIO(key, in, out, true)
		}
	case MODE_PRIKEY_DECRYPT:
		if key, err := this.getPriKey(); err != nil {
			return err
		} else {
			return priKeyIO(key, in, out, false)
		}
	default:
		return errors.New("mode not found")
	}
}

func NewRSASecurity(pubStr, priStr string) (rsa *rSASecurity, pubKeyErr, priKeyErr error) {
	rsa = &rSASecurity{}
	pubKeyErr, priKeyErr = rsa.Init(pubStr, priStr)
	return
}

func NewRSASecurityByFile(pubFile, priFile string, ifCache bool) (rsa *rSASecurity, pubKeyErr, priKeyErr error) {
	rsa = &rSASecurity{}
	pubKeyErr, priKeyErr = rsa.InitByFile(pubFile, priFile, ifCache)
	return
}

func (this *rSASecurity) Init(pubStr, priStr string) (pubkeyErr, prikeyErr error) {
	this.isFileInit = false
	this.pubStr = pubStr
	this.priStr = priStr
	this.ifCache = true
	this.pubkey, pubkeyErr = getPubKey([]byte(this.pubStr))
	this.prikey, prikeyErr = getPriKey([]byte(this.priStr))
	return
}

func (this *rSASecurity) InitByFile(pubFile, priFile string, ifCache bool) (pubkeyErr, prikeyErr error) {
	this.isFileInit = true
	this.pubStr = pubFile
	this.priStr = priFile
	this.ifCache = ifCache
	_, pubkeyErr = this.getPubKey()
	_, prikeyErr = this.getPriKey()
	return
}

func (this *rSASecurity) getPubKey() (*rsa.PublicKey, error) {
	if this.isFileInit && !this.ifCache {
		f, err := os.Stat(this.pubStr)
		if err != nil {
			return nil, err
		}
		if f.ModTime().Equal(this.pubModTime) {
			if this.pubkey == nil {
				return nil, ErrPublicKey
			}
			return this.pubkey, nil
		} else {
			in, err := ioutil.ReadFile(this.pubStr)
			if err != nil {
				return nil, err
			}
			this.pubkey, err = getPubKey(in)
			if err == nil {
				this.pubModTime = f.ModTime()
			}
			return this.pubkey, err
		}
	} else {
		if this.pubkey == nil {
			return nil, ErrPublicKey
		}
		return this.pubkey, nil
	}
}

func (this *rSASecurity) getPriKey() (*rsa.PrivateKey, error) {
	if this.isFileInit && !this.ifCache {
		f, err := os.Stat(this.priStr)
		if err != nil {
			return nil, err
		}
		if f.ModTime().Equal(this.priModTime) {
			if this.prikey == nil {
				return nil, ErrPrivateKey
			}
			return this.prikey, nil
		} else {
			in, err := ioutil.ReadFile(this.priStr)
			if err != nil {
				return nil, err
			}
			this.prikey, err = getPriKey(in)
			if err == nil {
				this.priModTime = f.ModTime()
			}
			return this.prikey, err
		}
	} else {
		if this.prikey == nil {
			return nil, ErrPrivateKey
		}
		return this.prikey, nil
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	ErrDataToLarge     = errors.New("message too long for RSA public key size")
	ErrDataLen         = errors.New("data length error")
	ErrDataBroken      = errors.New("data broken, first byte is not zero")
	ErrKeyPairDismatch = errors.New("data is not encrypted by the private key")
	ErrDecryption      = errors.New("decryption error")
	ErrPublicKey       = errors.New("get public key error")
	ErrPrivateKey      = errors.New("get private key error")
)

/*公钥解密*/
func pubKeyDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if k != len(data) {
		return nil, ErrDataLen
	}
	m := new(big.Int).SetBytes(data)
	if m.Cmp(pub.N) > 0 {
		return nil, ErrDataToLarge
	}
	m.Exp(m, big.NewInt(int64(pub.E)), pub.N)
	d := leftPad(m.Bytes(), k)
	if d[0] != 0 {
		return nil, ErrDataBroken
	}
	if d[1] != 0 && d[1] != 1 {
		return nil, ErrKeyPairDismatch
	}
	var i = 2
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	i++
	if i == len(d) {
		return nil, nil
	}
	return d[i:], nil
}

/*私钥加密*/
func priKeyEncrypt(rand io.Reader, priv *rsa.PrivateKey, hashed []byte) ([]byte, error) {
	tLen := len(hashed)
	k := (priv.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, ErrDataLen
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], hashed)
	m := new(big.Int).SetBytes(em)
	c, err := decrypt(rand, priv, m)
	if err != nil {
		return nil, err
	}
	copyWithLeftPad(em, c.Bytes())
	return em, nil
}

/*公钥加密或解密Reader*/
func pubKeyIO(pub *rsa.PublicKey, in io.Reader, out io.Writer, isEncrytp bool) error {
	k := (pub.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	buf := make([]byte, k)
	var b []byte
	var err error
	size := 0
	for {
		size, err = in.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		if isEncrytp {
			b, err = rsa.EncryptPKCS1v15(rand.Reader, pub, b)
		} else {
			b, err = pubKeyDecrypt(pub, b)
		}
		if err != nil {
			return err
		}
		if _, err = out.Write(b); err != nil {
			return err
		}
	}
	return nil
}

/*私钥加密或解密Reader*/
func priKeyIO(pri *rsa.PrivateKey, r io.Reader, w io.Writer, isEncrytp bool) error {
	k := (pri.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	buf := make([]byte, k)
	var err error
	var b []byte
	size := 0
	for {
		size, err = r.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		if isEncrytp {
			b, err = priKeyEncrypt(rand.Reader, pri, b)
		} else {
			b, err = rsa.DecryptPKCS1v15(rand.Reader, pri, b)
		}

		if err != nil {
			return err
		}
		if _, err = w.Write(b); err != nil {
			return err
		}
	}
	return nil
}

/*公钥加密或解密byte*/
func pubKeyByte(pub *rsa.PublicKey, in []byte, isEncrytp bool) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	if len(in) <= k {
		if isEncrytp {
			return rsa.EncryptPKCS1v15(rand.Reader, pub, in)
		} else {
			return pubKeyDecrypt(pub, in)
		}
	} else {
		iv := make([]byte, k)
		out := bytes.NewBuffer(iv)
		if err := pubKeyIO(pub, bytes.NewReader(in), out, isEncrytp); err != nil {
			return nil, err
		}
		return ioutil.ReadAll(out)
	}
}

/*私钥加密或解密byte*/
func priKeyByte(pri *rsa.PrivateKey, in []byte, isEncrytp bool) ([]byte, error) {
	k := (pri.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	if len(in) <= k {
		if isEncrytp {
			return priKeyEncrypt(rand.Reader, pri, in)
		} else {
			return rsa.DecryptPKCS1v15(rand.Reader, pri, in)
		}
	} else {
		iv := make([]byte, k)
		out := bytes.NewBuffer(iv)
		if err := priKeyIO(pri, bytes.NewReader(in), out, isEncrytp); err != nil {
			return nil, err
		}
		return ioutil.ReadAll(out)
	}
}

var pemStart = []byte("-----BEGIN ")

/*读取公钥*/
//公钥可以没有如 -----BEGIN PUBLIC KEY-----的前缀后缀
func getPubKey(in []byte) (*rsa.PublicKey, error) {
	var pubKeyBytes []byte
	if bytes.HasPrefix(in, pemStart) {
		block, _ := pem.Decode(in)
		if block == nil {
			return nil, ErrPublicKey
		}
		pubKeyBytes = block.Bytes
	} else {
		var err error
		pubKeyBytes, err = base64.StdEncoding.DecodeString(string(in))
		if err != nil {
			return nil, ErrPublicKey
		}
	}

	pub, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, err
	} else {
		return pub.(*rsa.PublicKey), err
	}

}

/*读取私钥*/
func getPriKey(in []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(in)
	if block == nil {
		return nil, ErrPrivateKey
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return pri, nil
	}
	pri2, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	} else {
		return pri2.(*rsa.PrivateKey), nil
	}
}

/*从crypto/rsa复制 */
var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)

/*从crypto/rsa复制 */
func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

/*从crypto/rsa复制 */
func decrypt(random io.Reader, priv *rsa.PrivateKey, c *big.Int) (m *big.Int, err error) {
	if c.Cmp(priv.N) > 0 {
		err = ErrDecryption
		return
	}
	var ir *big.Int
	if random != nil {
		var r *big.Int

		for {
			r, err = rand.Int(random, priv.N)
			if err != nil {
				return
			}
			if r.Cmp(bigZero) == 0 {
				r = bigOne
			}
			var ok bool
			ir, ok = modInverse(r, priv.N)
			if ok {
				break
			}
		}
		bigE := big.NewInt(int64(priv.E))
		rpowe := new(big.Int).Exp(r, bigE, priv.N)
		cCopy := new(big.Int).Set(c)
		cCopy.Mul(cCopy, rpowe)
		cCopy.Mod(cCopy, priv.N)
		c = cCopy
	}

	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}
	if ir != nil {
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}

	return
}

/*从crypto/rsa复制 */
func copyWithLeftPad(dest, src []byte) {
	numPaddingBytes := len(dest) - len(src)
	for i := 0; i < numPaddingBytes; i++ {
		dest[i] = 0
	}
	copy(dest[numPaddingBytes:], src)
}

/*从crypto/rsa复制 */
func nonZeroRandomBytes(s []byte, rand io.Reader) (err error) {
	_, err = io.ReadFull(rand, s)
	if err != nil {
		return
	}
	for i := 0; i < len(s); i++ {
		for s[i] == 0 {
			_, err = io.ReadFull(rand, s[i:i+1])
			if err != nil {
				return
			}
			s[i] ^= 0x42
		}
	}
	return
}

/*从crypto/rsa复制 */
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

/*从crypto/rsa复制 */
func modInverse(a, n *big.Int) (ia *big.Int, ok bool) {
	g := new(big.Int)
	x := new(big.Int)
	y := new(big.Int)
	g.GCD(x, y, a, n)
	if g.Cmp(bigOne) != 0 {
		return
	}
	if x.Cmp(bigOne) < 0 {
		x.Add(x, n)
	}
	return x, true
}
