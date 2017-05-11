package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// ---------------
// Encode
// 用gob进行数据编码
//
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// -----------------
// Decode
// 用gob进行数据解码
//
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

type User struct {
	Name string
	Age  int
}

type Out struct {
	Age  int
	Name string
}

func New() *User {
	return &User{Name: "viney", Age: 32}
}

func main() {
	// 实例化User
	u := New()

	// 对User编码
	b, err := Encode(u)
	if err != nil {
		fmt.Println("encode fail: " + err.Error())
	}

	// 对User解码
	var out Out
	if err := Decode(b, &out); err != nil {
		fmt.Println("decode fail: " + err.Error())
	}

	fmt.Println(out)
}
