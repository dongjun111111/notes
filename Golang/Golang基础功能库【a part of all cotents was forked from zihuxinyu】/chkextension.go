// 检测文件后缀名
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ChkExtension("example.zip", "zip"))
}

// 检测文件后缀名
func ChkExtension(path string, e string) bool {
	return strings.LastIndex(path, "."+e)+len("."+e) == len(path)
}
