package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const (
	//配置文件
	CONFFILE string = "config.txt"
)

func main() {
	for {
		//读取文件的信息
		bytes, err := ioutil.ReadFile(CONFFILE)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//按照换行符分割
		text := string(bytes)
		cmdarr := strings.Split(text, "\r\n")

		//是否新的开始
		isBegin := 1
		for _, val := range cmdarr {
			tmpval := strings.TrimSpace(val)

			//如果是新命令开始，那么是切换目录操作
			if tmpval != "" && isBegin == 1 {
				os.Chdir(tmpval)
			} else if tmpval != "" {
				//分割命令
				cmdarr := strings.Split(tmpval, " ")
				//命令名称
				command := cmdarr[0]
				//命令参数
				params := cmdarr[1:]
				//执行cmd命令
				execCommand(command, params)
			}

			//如果是空行，说明新的命令开始
			if tmpval == "" {
				isBegin = 1
				continue
			} else {
				isBegin = 0
			}
		}
	}
}

//执行命令函数
//commandName 命名名称，如cat，ls，git
//params 命令参数，如ls -l的-l，git log 的log
func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}
