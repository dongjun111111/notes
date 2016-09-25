package main

//删除代码中注释内容 (适配常见注释类型)
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var sem = make(chan int)

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil {
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

//读取文件内容
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

//删除文件中注释内容
func RewriteFileContent(filename string) {
	sem <- 1
	contents, _ := os.OpenFile(filename, os.O_RDONLY, 0777)
	defer contents.Close()
	buff := bufio.NewReader(contents)
	var basestr string
	for {
		str, err := buff.ReadString('\n')
		if err != nil && err.Error() == "EOF" {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if strings.Contains(str, "//") {
			str = strings.TrimSpace(str)
		}
		filenameres := strings.Split(filename, ".")
		if len(filenameres) <= 1 || len(filenameres) > 2 {
			panic("文件类型错误，请检查源文件是否符合条件")
			return
		}
		switch filenameres[1] {
		case "go":
			//适用于go
			if !strings.HasPrefix(str, "//") || strings.HasPrefix(str, "// @") {
				basestr += str
			}
		case "html", "htm":
			//适用于html
			if !strings.HasPrefix(str, "<!--") || !strings.HasSuffix(str, "-->") {
				basestr += str
			}
		case "sh":
			//适用于shell
			if !strings.HasPrefix(str, "#") {
				basestr += str
			}
		case "java":
			//适用于java
			if !strings.HasPrefix(str, "/**") && !strings.HasSuffix(str, "*/") {
				basestr += str
			}
		case "cpp", "c":
			//适用于c、c++
			if !strings.HasPrefix(str, "/*") && !strings.HasSuffix(str, "*/") {
				basestr += str
			}
		}
	}
	var filecontent = []byte(basestr)
	err2 := ioutil.WriteFile(filename, filecontent, 0666)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(filename + " - 删除注释内容成功!")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	files, _ := WalkDir("D:\\gopath\\src\\test\\ngrok", ".go")
	for _, v := range files {
		go RewriteFileContent(v)
		<-sem
	}
}
