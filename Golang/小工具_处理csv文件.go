package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

var quick, first, unrepeat bool

func main() {
	path := flag.String("p", "", "-p 指定csv文件路径")
	logname := flag.String("l", "", "-l 指定结果输出路径,不指定则输出到标准输出")
	flag.BoolVar(&quick, "q", false, "-q 数据全部加载到内存中处理,默认是少量数据加载到内存操作,bool值")
	flag.BoolVar(&first, "f", false, "-f 使用首次遇到的一条数据,默认是使用最后遇到的一条数据,bool值")
	flag.BoolVar(&unrepeat, "u", false, "-u 默认是使用重复数据,bool值")
	flag.Parse()
	if *logname != "" {
		File, err := os.Create(*logname)
		if err != nil {
			log.Println("创建输出文件失败:", err.Error())
			return
		}
		log.SetOutput(File)
	}
	if *path == "" {
		flag.Usage()
		return
	}
	Read(*path)
}

func Read(path string) {
	File, err := os.Open(path)
	if err != nil {
		log.Println("读取csv文件失败,错误信息:", err.Error())
		return
	}
	defer File.Close()
	csvr := csv.NewReader(File)
	if quick {
		QMerge(csvr)
	} else {
		list := SMerge(csvr, File)
		FromLineRead(list, path)
	}
}

func QMerge(r *csv.Reader) {
	var m = make(map[string][]string)
	var list []string
	var repeat []string
	r.Read()
	var err error
	var key string
	for {
		list, err = r.Read()
		if err != nil {
			if err != io.EOF {
				log.Println("读取文件内容失败,错误信息:", err.Error())
			}
			break
		}
		if len(list) != 5 {
			log.Println("无效数据:", list)
			continue
		}
		key = strings.TrimSpace(list[1] + list[3])
		if key != "" {
			if _, ok := m[key]; ok {
				repeat = append(repeat, key)
				if !first {
					m[key] = list
				}
			} else {
				m[key] = list
			}
		}
	}
	for _, value := range repeat {
		if unrepeat {
			delete(m, value)
		} else {
			log.Println(m[value])
		}
	}

	if unrepeat {
		for _, value := range m {
			log.Println(value)
		}
	}
}

func SMerge(r *csv.Reader, seek io.Seeker) []int {
	var m = make(map[string]int)
	var list []string
	var repeat []string
	r.Read()
	var key string
	var err error
	var line int
	for {
		list, err = r.Read()
		if err != nil {
			if err != io.EOF {
				log.Println("读取文件内容失败,错误信息:", err.Error())
			}
			break
		}
		if len(list) != 5 {
			log.Println("无效数据:", list)
			line++
			continue
		}
		key = strings.TrimSpace(list[1] + list[3])
		if key != "" {
			if _, ok := m[key]; ok {
				repeat = append(repeat, key)
				if !first {
					m[key] = line
				}
			} else {
				m[key] = line
			}
		}
		line++
	}
	var lines = make([]int, 0, len(m))
	for _, v := range repeat {
		if unrepeat {
			delete(m, v)
		} else {
			lines = append(lines, m[v])
		}
	}
	if unrepeat {
		for _, v := range m {
			lines = append(lines, v)
		}
	}
	sort.Ints(lines)
	return lines
}

func FromLineRead(lines []int, path string) {
	File, err := os.Open(path)
	if err != nil {
		log.Println("读取csv文件失败:", err.Error())
		return
	}
	defer File.Close()
	r := csv.NewReader(File)
	r.Read()
	var list []string
	var line, index int
	for {
		list, err = r.Read()
		if err != nil {
			if err != io.EOF {
				log.Println("读取文件内容失败,错误信息:", err.Error())
			}
			break
		}
		if lines[index] == line {
			log.Println(list)
			index++
			if index >= len(lines) {
				break
			}
		}
		line++
	}
}
