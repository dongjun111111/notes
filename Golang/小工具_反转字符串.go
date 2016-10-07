package main

import (
	"fmt"
)

func reverseString(s string) string {
	//将字符串string转换成rune类型，而后才能进行对调操作
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverseString("Hello世界"))
}
