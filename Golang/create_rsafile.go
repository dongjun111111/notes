package main


import (
"crypto/rand"
"crypto/rsa"
"crypto/x509"
"encoding/pem"
"errors"


"fmt"
"io/ioutil"
"log"
"os"
"time"
)


func main() {
var bits int
bits = 2048
if err := GenRsaKey(bits); err != nil {
log.Fatal("密钥文件生成失败！")
}
log.Println("密钥文件生成成功！")


initData := "abcdefghijklmnopq"
init := []byte(initData)


data, err := RsaEncrypt(init)
if err != nil {
panic(err)
}
pre := time.Now()
origData, err := RsaDecrypt(data)
if err != nil {
panic(err)
}
now := time.Now()
fmt.Println(now.Sub(pre))
fmt.Println(string(origData))
for {


}
}


var decrypted string
var privateKey, publicKey []byte


func init() {
var err error
// flag.StringVar(&decrypted, "d", "", "加密过的数据")
// flag.Parse()
publicKey, err = ioutil.ReadFile("public.pem")
if err != nil {
os.Exit(-1)
}
privateKey, err = ioutil.ReadFile("private.pem")
if err != nil {
os.Exit(-1)
}
}


func GenRsaKey(bits int) error {
// 生成私钥文件
privateKey, err := rsa.GenerateKey(rand.Reader, bits)
if err != nil {
return err
}
derStream := x509.MarshalPKCS1PrivateKey(privateKey)
block := &pem.Block{
Type:  "私钥",
Bytes: derStream,
}
file, err := os.Create("private.pem")
if err != nil {
return err
}
err = pem.Encode(file, block)
if err != nil {
return err
}
// 生成公钥文件
publicKey := &privateKey.PublicKey
derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
if err != nil {
return err
}
block = &pem.Block{
Type:  "公钥",
Bytes: derPkix,
}
file, err = os.Create("public.pem")
if err != nil {
return err
}
err = pem.Encode(file, block)
if err != nil {
return err
}
return nil
}


// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
block, _ := pem.Decode(publicKey)
if block == nil {
return nil, errors.New("public key error")
}
pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
if err != nil {
return nil, err
}
pub := pubInterface.(*rsa.PublicKey)
return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}


// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
block, _ := pem.Decode(privateKey)
if block == nil {
return nil, errors.New("private key error!")
}
priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
if err != nil {
return nil, err
}
return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
