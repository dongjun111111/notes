package main

/*
运行:go run t2.go path1 path2
path1 和path2是存放照片的路径，程序默认会按照默认的路径顺序来扫描文件，并且删除后面的目录中的相同的文件／照片
*/
import (
	"container/list"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileInfo struct {
	path string
	//  name string
	md5 string
}

var file_list map[string]FileInfo
var file_delete_list list.List

func getFileInfo(path string) bool {
	md5, err := calMd5(path)
	if err != nil {
		return false
	}
	var info FileInfo
	info.md5 = md5
	info.path = path
	addInfo(info)
	return true
}

func addInfo(info FileInfo) {
	_, ok := file_list[info.md5]
	if ok {
		file_delete_list.PushBack(info.path)
	} else {
		file_list[info.md5] = info
	}
}

func doInfo(path string, info os.FileInfo, err error) error {
	if info == nil {
		fmt.Println(path + " info is nil")
		return err
	}
	if info.IsDir() {
		// if return not nil ,walk func will break
		return nil
	}
	getFileInfo(path)
	return nil
}

func do_dir(path string) error {
	err := filepath.Walk(path, doInfo)
	return err
}

func calMd5(path string) (md5_str string, er error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return md5_str, err
	}
	h := md5.New()
	io.Copy(h, file)
	md5_bytes := h.Sum(nil)
	md5_str = fmt.Sprintf("%X", md5_bytes)
	return md5_str, err
}

func main() {
	flag.Parse()
	argc := flag.NArg()
	//  file_list = map[string]FileInfo
	file_list = make(map[string]FileInfo)
	fmt.Println(file_list)
	fmt.Println(argc)
	folder_list := list.New()
	for i := 0; i != argc; i++ {
		folder_list.PushBack(flag.Arg(i))
	}
	for folder := folder_list.Front(); folder != nil; folder = folder.Next() {
		root := folder.Value.(string)
		fmt.Println(root)
		err := do_dir(root)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("File to delete is :")
	for path := file_delete_list.Front(); path != nil; path = path.Next() {
		fmt.Println("deleteing " + path.Value.(string))
		os.Remove(path.Value.(string))
	}

	// for _, info := range file_list {
	//  fmt.Println(info.path)
	// }
}
