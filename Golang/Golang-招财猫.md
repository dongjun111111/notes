#喵
##基础知识
###常量 const iota 
const可以放到func外面，其他变量的声明不可以放到外面。
<pre>
package main

import "fmt"
import "os"

const z string = "这是string"

//常量组合声明时，iota每次引用会逐步自增，初始值为0，步进值为1
const (
	a uint8  = iota
	b uint8  = iota
	c uint16 = iota
)
//即使iota不是在常量组内第一个开始引用，也会按组内常量数量递增
const (
	a1        = 4
	a2 string = "d"
	a3 bool   = true
	a4 int    = iota
)
//枚举的常量都为同一类型时，可以使用简单序列格式(组内复用表达式).
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
###数组 Array
<pre>
package main

import "fmt"

func main() {
	var a [3]int = [3]int{3, 4, 5}
	var b [2]int = [2]int{} //[0 0]
	//使用...自动计算数组的长度
	var c = [...]int{5, 6, 7, 8, 9} //[5 6 7 8 9]
	d := [6]int{}                   //[0 0 0 0 0 0]
	//多维数组.多维数组只能自动计算最外围数组长度
	e := [...][3]int{{1, 2, 5}, {4, 6, 4}} //[[1 2 5] [4 6 4]]
	//初始化指定索引的数组元素，未指定初始化的元素保持默认零值
	var f = [...]string{2: "first", 4: "second"}
	fmt.Println(a[2])
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//通过下标访问多维数组元素
	fmt.Println(e, "++++++++", e[1][2]) //[[1 2 5] [4 6 4]] ++++++++ 4
	fmt.Println("f数组所有：", f, "f数组元素：", f[2])
}
output==>
5
[0 0]
[5 6 7 8 9]
[0 0 0 0 0 0]
[[1 2 5] [4 6 4]] ++++++++ 4
f数组所有： [  first  second] f数组元素： first
</pre>