package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// 获取文件修改时间
func FileMTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// 获取文件大小
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// 删除文件
func Unlink(file string) error {
	return os.Remove(file)
}

// 重命名文件
func Rename(file string, to string) error {
	return os.Rename(file, to)
}

// 往文件写入字符串
func FilePutContent(file string, content string) (int, error) {
	fs, e := os.Create(file)
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	return fs.WriteString(content)
}

// 从文件中获取字符串
func FileGetContent(file string) (string, error) {
	if !IsFile(file) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

// it returns false when it's a directory or does not exist.
func IsFile(file string) bool {
	f, e := os.Stat(file)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 创建文件
func CreateFile(dir string, name string) (string, error) {
	src := dir + name + "/"
	if IsExist(src) {
		return src, nil
	}

	if err := os.MkdirAll(src, 0777); err != nil {
		if os.IsPermission(err) {
			fmt.Println("你不够权限创建文件")
		}
		return "", err
	}

	return src, nil
}

type FileRepos []Repository

type Repository struct {
	Name     string
	FileTime int64
}

func (r FileRepos) Len() int {
	return len(r)
}

func (r FileRepos) Less(i, j int) bool {
	return r[i].FileTime < r[j].FileTime
}

func (r FileRepos) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// 获取所有文件
//如果文件达到最上限，按时间删除
func delFile(files []os.FileInfo, count int, fileDir string) {
	if len(files) <= count {
		return
	}

	result := new(FileRepos)

	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			*result = append(*result, Repository{Name: file.Name(), FileTime: file.ModTime().Unix()})
		}
	}

	sort.Sort(result)
	deleteNum := len(files) - count
	for k, v := range *result {
		if k+1 > deleteNum {
			break
		}
		Unlink(fileDir + v.Name)
	}
	return
}

func main() {
	size, _ := FileSize("2016-09-09startfromiteye.log")
	fmt.Println("文件大小:", size)
	//生成文件
	dir := "E:/golang/myorm/src/"
	file, err := CreateFile(dir, "20130829")

	if err != nil {
		return
	}

	//写文件
	content := "teststttttt"
	l, e := FilePutContent(file+"1.txt", content)

	if e != nil && l <= 0 {
		return
	}

	//读文件
	// str, _ := FileGetContent(file + "1.txt")
	// fmt.Println("str", str)
	// size, _ := FileSize(file + "1.txt")
	// fmt.Println("size", size)
	// ftime, _ := FileMTime(file + "1.txt")
	// fmt.Println("ftime", ftime)

	// 获取所有文件
	//如果文件达到最上限，按时间删除
	files, _ := ioutil.ReadDir(file)
	delFile(files, 1, file)
	fmt.Println("count", len(files))
}
