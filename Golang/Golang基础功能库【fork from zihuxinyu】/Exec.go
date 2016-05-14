//
// Created by weibaohui on 14-8-28.
//
//文件作用:存放系统执行相关类
package Library

import (
	"bytes"
	"os/exec"
	"fmt"
)

//示例	system("ps -ef|grep ./server|grep -v grep|awk '{printf $2}'|xargs kill -1")

func ExecSystem(s string) {
	cmd := exec.Command("/bin/sh", "-c", s) //调用Command函数
	var out bytes.Buffer //缓冲字节

	cmd.Stdout = &out //标准输出
	err := cmd.Run() //运行指令 ，做判断
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out.String()) //输出执行结果
}
