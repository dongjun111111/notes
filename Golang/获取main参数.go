package main

//获取main函数入参的两种方式
import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// os.Args方式
	args := os.Args
	if args == nil || len(args) < 2 {
		fmt.Println("Hello 世界!")
	} else {
		fmt.Println("Hello ", args[1]) // 第二个参数，第一个参数为命令名
	}

	// flag.Args方式
	flag.Parse()
	var ch []string = flag.Args()
	if ch != nil && len(ch) > 0 {
		fmt.Println("Hello ", ch[0]) // 第一个参数开始
	}
}
