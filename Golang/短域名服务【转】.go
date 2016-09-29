// 短网址 路由跳转
package main

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	list map[string]string
)

func main() {
	list = make(map[string]string)
	addUrl("djason", "localhost:2333")
	addUrl("163", "163.com")
	addUrl("baidu", "baidu.com")
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}
func addUrl(key string, url string) bool {
	if strings.LastIndex(url, "http") != 0 {
		url = "http://" + url
	}
	if list["/"+key] != "" { //Key已存在
		return false
	}
	list["/"+key] = url
	return true
}
func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.URL.Path
	if url == "/" || url == "/favicon.ico" {
		if url == "/favicon.ico" { //处理图标请求
			http.NotFound(w, r)
			return
		}
		if len(r.Form["key"]) > 0 && len(r.Form["url"]) > 0 {
			if list["/"+r.Form["key"][0]] != "" {
				fmt.Fprintln(w, "该key已经被占用,请重新选择")
				return
			}
			if addUrl(r.Form["key"][0], r.Form["url"][0]) { //将新的 Url 插入到 Map 中
				fmt.Fprintln(w, "添加成功")
				fmt.Fprintln(w, "原Url："+r.Form["url"][0])
				fmt.Fprintln(w, "新Url：http://"+r.Host+"/"+r.Form["key"][0])
			} else {
				http.Redirect(w, r, "http://"+r.Host, http.StatusFound) //添加失败，重定向到首页
			}
		} else {
			fmt.Fprintln(w, `没有添加必要参数`)
		}
	} else {
		jump := list[url]
		if jump != "" {
			http.Redirect(w, r, jump, http.StatusFound) //302重定向
		} else {
			fmt.Fprintln(w, "没有找到链接")
		}
	}
}
