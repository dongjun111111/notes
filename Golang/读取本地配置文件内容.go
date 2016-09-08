package main

//读取配置文件内容
import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	//配置文件
	CONFFILE string = "config.txt"
)

func main() {

	//读取文件的信息
	bytes, err := ioutil.ReadFile(CONFFILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//按照换行符分割
	text := string(bytes)
	cmdarr := strings.Split(text, "\r\n")

	for _, val := range cmdarr {
		fmt.Println(val)
	}
}
