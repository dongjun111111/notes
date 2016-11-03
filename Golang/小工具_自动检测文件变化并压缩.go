package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	paths := []string{`需要压缩的文件目录地址`, `需要压缩的文件目录地址`}
	ExampleNewWatcher(paths)
}

var eventTime = make(map[string]int64)

func ExampleNewWatcher(paths []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					iscompress := true
					if !checkIfWatchExt(event.Name) {
						continue
					}
					mt := getFileModTime(event.Name)
					if t := eventTime[event.Name]; mt == t {
						//fmt.Println("[SKIP] # %s #", event.String())
						iscompress = false
					}
					//fmt.Println("file:", event.Name, iscompress, event.Op, fsnotify.Write, event.Op&fsnotify.Write)
					if iscompress {
						fmt.Println("modified file:", event.Name)
						icmd := exec.Command("uglifyjs", event.Name, "-m", "-o", event.Name)
						err1 := icmd.Run()

						if err1 != nil {
							fmt.Println(err1.Error())
						} else {
							eventTime[event.Name] = time.Now().Unix()
						}
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()
	for _, path := range paths {
		fmt.Println("[TRAC] Directory( %s )", path)
		err = watcher.Add(path)
		if err != nil {
			fmt.Println("[ERRO] Fail to watch directory[ %s ]", err.Error())
			//os.Exit(2)
		}
	}
	<-done
}

var watchExts = []string{".js"}

// checkIfWatchExt returns true if the name HasSuffix <watch_ext>.
func checkIfWatchExt(name string) bool {
	for _, s := range watchExts {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

// getFileModTime retuens unix timestamp of `os.File.ModTime` by given path.
func getFileModTime(path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	fmt.Println(path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("[ERRO] Fail to open file[ %s ]", err.Error())
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		fmt.Println("[ERRO] Fail to get file information[ %s ]", err.Error())
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
