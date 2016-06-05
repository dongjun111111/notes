// 解压zip文件
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 打开一个zip格式文件
	r, err := zip.OpenReader("file.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// 迭代压缩文件中的文件，打印出文件中的内容
	for _, f := range r.File {
		fmt.Printf("文件名 %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64))
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}

}
