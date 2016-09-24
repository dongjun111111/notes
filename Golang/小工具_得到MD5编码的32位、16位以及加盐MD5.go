package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type MD5Client struct{}

var MD5 = MD5Client{}

// 得到32位的MD5 string
func (this *MD5Client) Encrypt(plantext []byte) string {
	result := md5.Sum(plantext)
	return hex.EncodeToString(result[:])
}

//据说这种方式效率较上面的高
func (this *MD5Client) EasyEncrypt(plantext []byte) string {
	m := md5.Sum([]byte(plantext))
	return hex.EncodeToString(m[:])
}

// 得到 16位 MD5 string
func (this *MD5Client) Get16MD5String(data string) string {
	return data[8:24]
}

//加盐 ，得到32位的string
func (this *MD5Client) EncryptWithSalt(plantext []byte, salt []byte) string {
	hash := md5.New()
	hash.Write(plantext)
	hash.Write(salt)
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	md5_1 := MD5.Encrypt([]byte(`Jason`))
	fmt.Println(md5_1)
	fmt.Println(MD5.Get16MD5String(md5_1))
}
