package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func ListHref(html string) {
	//  (?m)开启多行文本模式，直到(?-m)。如果没出现就一直匹配到最后
	//  .*? 尽可能少的（贪婪模式关闭）匹配任意字符

	//var hrefRegexp = regexp.MustCompile("(?m)<a.*?[^<]>.*?</a>") //抓取所有链接地址
	//var imgRegexp = regexp.MustCompile("(?m)src=\"http://.*?(jpg|png|jpeg|gif|bmp|ico|tiff|swf|svg|eps|tga)") //抓取常见图片链接
	//var audioRegexp = regexp.MustCompile("(?m)\"http://.*?(mp3|wav|cda|wma|ra|rma|midi|ogg|ape|flac|aac|aiff|au)") //抓取常见音频链接
	//var hrefRegexp = regexp.MustCompile("(?m)\"http://.*?(.mp4|.avi|.rmvb|.wmv|.3gp|.rm|.mkv|.flv|.mpeg|.swf)") //抓取常见视频链接
	//var hrefRegexp = regexp.MustCompile(`<img(.*?)src="(.*?)"`) //图片2
	//var hrefRegexp = regexp.MustCompile(`<img.*?(?:>|\/>)`) //图片3
	var hrefRegexp = regexp.MustCompile(`<img(.*?)src="((http|https)://)(.*?)"`) //图片4
	match := hrefRegexp.FindAllString(html, -1)
	if match != nil {
		for i, v := range match {
			fmt.Println("[", i, "] --> ", v)
		}
	}
}

func main() {
	res, _ := http.Get("http://v.youku.com/v_show/id_XMTcxNjI0MTgxMg==.html?beta&f=28100028&from=y1.3-news-newgrid-123-10079.225627-225628.2-1")
	rrr, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	ListHref(string(rrr))
}
