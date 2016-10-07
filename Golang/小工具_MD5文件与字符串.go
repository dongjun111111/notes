package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func md5File(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	h := md5.New()
	file.Seek(0, 0)
	_, err = io.Copy(h, file)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	fmt.Println(md5File("testioutilfile.txt"))
	h := md5.New()
	io.WriteString(h, "MD5 string!")
	fmt.Printf("%x\n", h.Sum(nil))
}
