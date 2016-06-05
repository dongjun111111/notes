package Library

import (
	"strings"
	"path/filepath"
	"os"
	"os/exec"
	"fmt"
)
// 分离文件名与扩展名(包含.)
func SplitFilename(filename string) (baseName, ext string) {
	baseName = filename
	// 找到最后一个'.'
	ext = SubstringByte(filename, strings.LastIndex(filename, "."))
	baseName = strings.TrimRight(filename, ext)
	ext = strings.ToLower(ext)
	return;
}
// 转换文件的格式
// toExt包含.
func TransferExt(path string, toExt string) string {
	dir := filepath.Dir(path) + "/" // 文件路径
	name := filepath.Base(path) // 文件名 a.jpg
	// 获取文件名与路径
	baseName, _ := SplitFilename(name)
	return dir + baseName + toExt
}
func GetFilename(path string) string {
	return filepath.Base(path)
}
// file size
// length in bytes
func GetFilesize(path string) int64 {
	fileinfo, err := os.Stat(path)
	if err == nil {
		return fileinfo.Size()
	}
	return 0;
}
// 清空dir下所有的文件和文件夹
// RemoveAll会清空本文件夹, 所以还要创建之
func ClearDir(dir string) bool {
	err := os.RemoveAll(dir)
	if err != nil {
		return false
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return false
	}
	return true
}
// list dir's all file, return filenames
func ListDir(dir string) []string {
	f, err := os.Open(dir)
	if err != nil {
		return nil
	}
	names, _ := f.Readdirnames(0)
	return names
}



//当前应用的绝对路径
func GetAppRoot() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	return filepath.Dir(p)
}

//文件夹是否存在
func DirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

//检查文件夹是否存在，如果不存在，则创建一个新文件夹
func GetDir(path string) error {
	//文件夹是否存在
	if DirExists(path) {
		return nil
	} else {
		//创建文件夹
		if err := os.Mkdir(path, os.ModeDir); err != nil {
			return err
		}
		return nil
	}
}
//文件大小
func FileSize(size int) string {
	s := float32(size)
	if s > 1024*1024 {
		return fmt.Sprintf("%.1f M", s/(1024*1024))
	}
	if s > 1024 {
		return fmt.Sprintf("%.1f K", s/1024)
	}
	return fmt.Sprintf("%f B", s)
}
//是否文件
func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	if f.IsDir() {
		return false
	}
	return true
}
//是否目录
func IsDir(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return f.IsDir()
}
