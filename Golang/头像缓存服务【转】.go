package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	_PATH    = "./imgs/"
	_URL     = "https://secure.gravatar.com/avatar/"
	_DEFAYLT = "./default"
	_HOME    = `头像缓存服务
调用例子:
一般调用调用: /avatar/63fe3f5adfe6fd8464babe98bd1e4e93?name=128.jpg
头像更新调用: /avatar/63fe3f5adfe6fd8464babe98bd1e4e93?name=128.jpg&tag=refresh`
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/avatar/", avatar)
	http.ListenAndServe(":80", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(_HOME))
}

func avatar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	md5 := strings.Split(r.URL.Path, "/")[2] //取出MD5

	if len(md5) != 32 { //如果长度不为32位返回默认头像
		http.ServeFile(w, r, _DEFAYLT)
		return
	}

	fileName := md5 //不附带GET参数的话使用MD5作为文件名
	if r.FormValue("name") != "" {
		fileName = r.FormValue("name") //附带s参数的话使用s参数作为文件名
	}

	if r.Form.Get("tag") == "refresh" { //更新头像
		os.RemoveAll(_PATH + md5 + "/") //删除储存图片的目录
	}

	_, err := os.Stat(_PATH + md5 + "/" + fileName)
	if err != nil {
		img := httpGet(_URL + md5 + "?" + r.Form.Encode()) //GET头像
		if img == nil {
			http.ServeFile(w, r, _DEFAYLT)
			return
		}
		os.MkdirAll(_PATH+md5, os.ModePerm)                  //文件不存在,创建文件目录
		file, err := os.Create(_PATH + md5 + "/" + fileName) //创建文件
		defer file.Close()

		if err != nil { //文件创建失败
			log.Println(err) //打印错误并返回默认头像
			http.ServeFile(w, r, _DEFAYLT)
			return
		}
		file.Write(img) //写入数据到创建好的文件里
	}
	http.ServeFile(w, r, _PATH+md5+"/"+fileName) //返回头像
}

func httpGet(url string) (body []byte) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return
}
