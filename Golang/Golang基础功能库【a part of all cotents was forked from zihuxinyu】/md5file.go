// 对文件进行md5编码
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("md5test.log")
	if err != nil {
		panic(err)
	}

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return
	}
	fmt.Printf("%x\n", h.Sum(nil))
	// output: 43c6359298645ded23f3c2ee44acf564

	// 经过io.Copy操作后，file的偏移量(seek)被指向了最后面
	// 如果还需要使用则需要修改file色偏移量(seek)
	// 该行代码输出为空，因为file的seed已经位于最后了
	io.Copy(os.Stdin, file)
	// output:

	fmt.Print("\n")

	file.Seek(0, 0)

	// 该行输出文件的内容，因为file的偏移量(seek)被设置为0了
	io.Copy(os.Stdin, file)
	// output: md5test.log

	fmt.Print("\n")
}
