package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urls := "http://fengyun5.com/Sibao/600/1.html"
	res, _ := http.Get(urls)
	file, _ := os.Create("downloadFiles.html")
	io.Copy(file, res.Body)
	fmt.Println("下载完成！")
}
