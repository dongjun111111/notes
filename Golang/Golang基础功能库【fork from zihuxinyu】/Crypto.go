package Library

import (
	"encoding/hex"
	"crypto/md5"
	"io"
	"crypto/rand"
	"encoding/base64"
	math_rand "math/rand"
	"time"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"net/url"
	"bytes"
)

// Guid
func NewGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return MD5(base64.URLEncoding.EncodeToString(b))
}
// 后面加个str生成之, 更有保障, 确保唯一
func NewGuidWith(str string) string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return MD5(base64.URLEncoding.EncodeToString([]byte(string(b) + str)))
}


var key = MD5byte("&^&*^&*$##%#@%$#@%$@$#@$%@")
var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
//md5加密
func MD5(s string) string {
	//是否使用高强度密码
	if b, _ := beego.AppConfig.Bool("StrongPassword"); b {
		return MD5Ex(s)
	} else {
		return hex.EncodeToString(MD5byte(s))
	}
}
func MD5byte(s string) []byte {
	h := md5.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
//加盐强密码
func MD5Ex(s string) string {
	h := md5.New()
	h.Write(key)
	h.Write([]byte(s))
	h.Write(iv)
	//fmt.Println(hex.EncodeToString(h.Sum(nil)), fmt.Sprintf("%x", h.Sum(nil)), MD5(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
//sha1加密
func SHA1(s string) string {
	return hex.EncodeToString(SHA1Byte(s))
}
func SHA1Byte(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
//Base64编码
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
//Base64解码
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
//AES编码
func AesEncode(src []byte) ([]byte, error) {
	var s []byte
	c, err := aes.NewCipher(key)
	if err == nil {
		cfb := cipher.NewCFBEncrypter(c, iv)
		s = make([]byte, len(src))
		cfb.XORKeyStream(s, src)
	}
	return s, err
}
//AES解码
func AesDecode(src []byte) ([]byte, error) {
	var s []byte
	c, err := aes.NewCipher(key)
	if err == nil {
		cfb := cipher.NewCFBDecrypter(c, iv)
		s = make([]byte, len(src))
		cfb.XORKeyStream(s, src)
	}
	return s, err
}
//utf-8转gbk
func Utf8ToGBK(str string) string {
	//字符集转换
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(str)
}
//gbk转utf-8
func GBKToUtf8(str string) string {
	//字符集转换
	enc := mahonia.NewDecoder("gbk")
	return enc.ConvertString(str)
}
//url编码
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}



// 随机密码
// num 几位
func RandomPwd(num int) string {
	chars := make([]byte, 62)
	j := 0
	for i := 48; i <= 57; i++ {
		chars[j] = byte(i)
		j++
	}
	for i := 65; i <= 90; i++ {
		chars[j] = byte(i)
		j++
	}
	for i := 97; i <= 122; i++ {
		chars[j] = byte(i)
		j++
	}
	j--;
	str := ""
	math_rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		x := math_rand.Intn(j)
		str += string(chars[x])
	}
	return str
}
//生成随机字符串
func RandomString(num int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < num; {
		if string(RandomInt(65, 90)) != temp {
			temp = string(RandomInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}
//生成随机数字
func RandomInt(min int, max int) int {
	math_rand.Seed(time.Now().UTC().UnixNano())
	return min + math_rand.Intn(max-min)
}
