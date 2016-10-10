package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	upload_path string = "./cache-basedir/"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

//上传
func uploadHandle(w http.ResponseWriter, r *http.Request) {
	//从请求当中判断方法
	if r.Method == "GET" {
		io.WriteString(w, "<html><head><title>我的第一个页面</title><style>body { text-align:center;}</style></head><body><form action='' method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file'  /><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
	} else {
		//获取文件内容 要这样获取
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//创建文件
		fW, err := os.Create(upload_path + head.Filename)
		if err != nil {
			fmt.Println("文件创建失败")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}
		io.WriteString(w, head.Filename+" 保存成功")
		//http.Redirect(w, r, "/hello", http.StatusFound)
	}
}
func main() {
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/uploadfile", uploadHandle)
	http.HandleFunc("/", NotFoundHandler) //定义404页
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败")
		return
	}
}
