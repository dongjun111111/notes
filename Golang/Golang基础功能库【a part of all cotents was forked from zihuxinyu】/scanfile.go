// 递归读取指定目录下的文件
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("./",
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				fmt.Println("dir:", path)
				return nil
			}
			fmt.Println("file:", path)
			return nil
		})
}
