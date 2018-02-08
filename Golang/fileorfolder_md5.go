package main

import (
	"crypto/md5"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

var directory, file string
var merge bool
var limit int

func init() {
	// directory = flag.String("d", "", "The directory contains all the files that need to calculate the md5 value")
	// file = flag.String("f", "", "The file that need to caclulate the md5 value")
	// merge = flag.Bool("merge", false, "Merging all md5 values to one (Folder type only)")
	// limit = flag.Int("max", 0, "limit the max files to caclulate.")
}
func Md5SumFile(file string) (value [md5.Size]byte, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	value = md5.Sum(data)
	return
}

type result struct {
	path   string
	md5Sum [md5.Size]byte
	err    error
}

func Md5SumFolder(folder string, limit int) (map[string][md5.Size]byte, error) {
	returnValue := make(map[string][md5.Size]byte)
	var limitChannel chan (struct{})
	if limit != 0 {
		limitChannel = make(chan struct{}, limit)
	}

	done := make(chan struct{})
	defer close(done)

	c := make(chan result)
	errc := make(chan error, 1)
	var wg sync.WaitGroup
	go func() {
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			if limit != 0 {
				//如果已经满了则阻塞在这里
				limitChannel <- struct{}{}
			}

			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path: path, md5Sum: md5.Sum(data), err: err}:
				case <-done:
				}
				if limit != 0 {
					//读出数据，这样就有新的文件可以处理
					<-limitChannel
				}

				wg.Done()
			}()
			select {
			case <-done:
				return errors.New("Canceled")
			default:
				return nil
			}
		})
		errc <- err
		go func() {
			wg.Wait()
			close(c)
		}()
	}()
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		returnValue[r.path] = r.md5Sum
	}
	if err := <-errc; err != nil {
		return nil, err
	}
	return returnValue, nil
}

func main() {
	file = ""
	directory = "./"
	merge = false // true 会将文件夹下所有文件的MD5显示出来，不然只会把文件夹的一个MD5显示出来
	limit = 100
	timeStart := time.Now()
	flag.Parse()
	if directory == "" && file == "" {
		return
	}
	if file != "" {
		md5Value, err := Md5SumFile(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%x %s\n", md5Value, file)
		return
	}
	if directory != "" {
		result, err := Md5SumFolder(directory, limit)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var paths []string
		for path := range result {
			paths = append(paths, path)
		}
		sort.Strings(paths)
		if merge {
			var md5value string
			for _, path := range paths {
				md5value += fmt.Sprintf("%x", result[path])
			}
			fmt.Printf("%x %s\n", md5.Sum([]byte(md5value)), directory)
		} else {
			for _, path := range paths {
				fmt.Printf("%x %s\n", result[path], path)
			}
		}

	}
	fmt.Println(time.Since(timeStart).String())
}
