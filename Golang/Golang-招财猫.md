#喵
##基础知识
const iota ;const可以放到func外面，其他变量的声明不可以放到外面。
<pre>
package main

import "fmt"
import "os"

const z string = "这是string"
const (
	a uint8  = iota
	b uint8  = iota
	c uint16 = iota
)

const (
	a1        = 4
	a2 string = "d"
	a3 bool   = true
	a4 int    = iota
)
const (
	x = iota
	x1
	x2
)

//定制iota序列初始值与步进值
const (
	z1 = (iota + 2) * 3 //初始值 6 ，步进值 3
	z2
	z3
)

func main() {
	var i int
	i = 4
	j := "hello"
	t := "你好"
	const x int = 4
	fmt.Println(i, j, t, "ok")
	fmt.Println(x, z, "c iota:", c, "a4 iota:", a4) //c等于2 a4等于3
	fmt.Println("x2 iota:", x2)                     //x2等于2
	fmt.Println("z1 value :", z1, "z2 value:", z2, "z3 value:", z3)
	os.Exit(0)
}
output==>
4 hello 你好 ok
4 这是string c iota: 2 a4 iota: 3
x2 iota: 2
z1 value : 6 z2 value: 9 z3 value: 12
</pre>