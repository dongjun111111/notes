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
###Slice
<pre>
package main

import "fmt"

func main() {
	//slice
	var a []int
	fmt.Println(a, "len:", len(a), "cap:", cap(a)) //	[] len: 0 cap: 0
	var b []int = []int{5, 78, 5, 56, 45, 3}
	fmt.Println(b)         //[5 78 5 56 45 3]
	for _, pp := range b { //5,78,5,56,45,3,
		fmt.Printf("%d,", pp)
		if len(b) == 6 {
			fmt.Println()
		}
	}
	//使用内置函数make初始化slice，第一参数是slice类型，第二参数是长度，第三参数是容量(省略时与长度相同)
	var c = make([]int, 3, 10)
	fmt.Println(c, "len:", len(c), "cap:", cap(c))
	var d = new([]int)                               //var d= new([]int ,2,3)写法错误
	fmt.Println(d, "len:", len(*d), "cap:", cap(*d)) //这里用地址的形式访问
	e := []int{4, 5, 7, 8, 54}
	fmt.Println("e:", e, "len:", len(e), "cap:", cap(e))
	e1 := e[0:2] //值是e[0] e[1]，不包括e[2].  [4 5]
	fmt.Println("e1:", e1, "len:", len(e1), "cap:", cap(e1))
	e2 := e[:3] //值是e[0] e[1] e[2] ,不包括e[3]
	fmt.Println(e2)
	e3 := e[:] //相当于复制了一个e切片
	fmt.Println("e3:", e3, "len:", len(e3), "cap:", cap(e3))
	//向slice中增加/修改元素
	f := []string{} //空的slice
	f = append(f, "Jason")
	f = append(f, "Miao")
	fmt.Println(f, "len:", len(f), "cap:", cap(f))
	/*删除slice中指定的元素,因为slice引用指向底层数组，数组的长度不变元素是不能删除的，
	所以删除的原理就是排除待删除元素后用其他元素重新构造一个数组*/
	index := 2 //删除第三个元素
	var ee []int
	ee = append(e[:index], e[index+1:]...)
	fmt.Println(ee) //[4 5 8 54]
	//向slice中间插入元素 注意：保存后部剩余元素，必须新建一个临时切片
	rear := append([]int{}, ee[index:]...)
	ee = append(ee[0:index], 100)
	ee = append(ee, rear...)
	fmt.Println("after insert:", ee)

}
output==>
[] len: 0 cap: 0
[5 78 5 56 45 3]
5,
78,
5,
56,
45,
3,
[0 0 0] len: 3 cap: 10
&[] len: 0 cap: 0
e: [4 5 7 8 54] len: 5 cap: 5
e1: [4 5] len: 2 cap: 5
[4 5 7]
e3: [4 5 7 8 54] len: 5 cap: 5
[Jason Miao] len: 2 cap: 2
[4 5 8 54]
after insert: [4 5 100 8 54]
</pre>
###字典/映射 Map
<pre>
package main

import "fmt"

func main() {
	/*map是引用类型，使用内置函数 make进行初始化，
	未初始化的map零值为 nil长度为0，并且不能赋值元素
	var m map[int]int
	fmt.Println(len(m))   -->false can`t enter value to a nil map
	*/
	//使用内置函数make初始化map
	var m map[int]int = make(map[int]int)
	m[0] = 4
	m[3] = 2
	m[4] = 1
	m[7] = 2
	fmt.Println(m)
	fmt.Println("is nil:", nil == m) //false
	//直接赋值初始化map
	n1 := map[string]int{
		"jason": 3,
		"miao":  2, //最后的逗号一定要加上
	}
	type S struct {
		age  int
		name int
	}
	n2 := map[string]S{
		"a": S{3, 5},
		"b": {23, 7}, //最后的逗号一定要加上;类型名称可忽略
	}
	fmt.Println(n1, n2)
	//map的使用:修改、删除元素
	fmt.Println(n2["b"]) //{23 7}
	n2["b"] = S{77, 77}  //修改
	fmt.Println(n2["b"])
	delete(n2, "b")      //删除
	fmt.Println(n2["b"]) //空的map是{0,0}   {0 0}
}
output==>
map[4:1 7:2 0:4 3:2]
is nil: false
map[jason:3 miao:2] map[b:{23 7} a:{3 5}]
{23 7}
{77 77}
{0 0}
</pre>