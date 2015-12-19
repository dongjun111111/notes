##不可变字符串
<pre>
package main 
import "fmt"
func main() {
	cat := "cat is cat"
	var rr byte =cat[0]
	fmt.Println(string(rr))
	cat = "dog is not dog"
	fmt.Println(cat)
}

Output==>
c
dog is not dog
</pre>
字符串的不可变性，不允许修改字符串内容。很多初学者可能认为cat = "dog is not dog"不是改变字符串内容了吗？这种理解是错误的，它只是将变量cat指向了另一个内存地址，原来字符串并没改变，你改变的只是变量的地址。
在Go语言中，单引号表示一个Unicode字符。
<pre>
package main 
import "fmt"
func main() {
	var m [5]byte = [5]byte{}
	m[0] = 'c'
	fmt.Println(m)
}

Output==>
[99 0 0 0 0]
</pre>
这里使用长度为5的字节数组来存放，并且在第一个位置放入‘c’字符，99就对应了'c'字符的ascii码值。最后打印结果,效果如上。对于中文的编码需要3个字节。而该byte类型的数组，每个数组元素只有一个字节容量，所以放不下中文字符，那么如果我们非要放中文字符。解决方法是：
<pre>
package main 
import "fmt"
func main() {
	var m [5]rune=[5]rune{}
	m[0] = 'c'
	m[1] = '猫'
	fmt.Println(m)
}
Output==>
[99 29483 0 0 0]
</pre>
将byte数组换成rune类型的数组就行了。原因就是rune是有32位的长度，足够放下3个字节表示的中文字符了 。<br>
连接跨行行字符串时，"+" 必须在上⼀一⾏行末尾，否则导致编译错误。
<pre>
s := "hello" + 
"world"
</pre>
##字符串遍历
这里有两种情况：
<br>下面一种是纯英文格式，比较方便。
<pre>
package main 
import "fmt"
func main() {
	a := "my name is jason"
	for i:=0; i< len(a);i++{
		fmt.Printf("%c",a[i])
	}
}
Output==>
my name is jason
</pre>
%c表示格式化成字符.
<br>第二种是带有中文的。
<pre>
package main 
import "fmt"
func main() {
	a := "我是谁"
	b := []rune(a)
	for i:=0;i< len(b);i++{
		fmt.Printf("%c",b[i])
	}
}
</pre>
把字符串里面的byte数组转成rune数组就可以了.
##字符串拼接
直接用+就行了。
<pre>
package main
import "fmt"
func main() {
	a := "你"
	b := a + "好"
	fmt.Println(b)
}
Output==>
你好
</pre>
如果想提高性能，可以导入bytes包，如下：
<pre>
package main
import (
	"fmt"
	"bytes"
)
func main (){
	a := bytes.Buffer{}
	a.WriteString("你")
	a.WriteString("好")
	fmt.Println(a.String())
}
Output==>
你好
</pre>
##给定一个int型数组，找出其中的奇数
<pre>
package main
import "fmt"
func main (){
	arr := []int{1,3,-5,41,22,64}
	for _,num := range arr {
		if isOdd(num) {
			fmt.Printf("%d\n",num)
		}
	}
}
func isOdd(num int) bool {
	return num & 1 == 1
}
Output==>
1
3
-5
41
</pre>

###GO语言range的用法
range是go语言系统定义的一个函数。
函数的含义是在一个数组中遍历每一个值，返回该值的下标值和此处的实际值。
假如说a[0]=10，则遍历到a[0]的时候返回值为0，10两个值。
<pre>
package main
import (
	"fmt"
)
func main (){
	sum := 0.0
	var avg float64
	xs := []float64 {1,2,3,4,5,6}
	switch len(xs) {
		case 0:
			avg =0
		default:
		for _,v := range xs {
			sum += v
		}
		avg = sum /float64(len(xs))
	}
	fmt.Println(avg)
}

Output ==>
3.5
</pre>


###go二维数组
<pre>
package main
import "fmt"
func main (){
	var two [2][3]int
	for i:=0;i<2;i++{
		for j:= 0; j< 3; j++{
			two[i][j] =i + j
		}
	}
	fmt.Println("2d:",two)
}

Output ==>
2d: [[0 1 2] [1 2 3]]
</pre>


###Slices：切片
Slices是Go语言中的关键数据类型，它有比数组（arrays）更强的访问接口。
<pre>
package main
import "fmt"
func main() {
    //跟数组（arrays）不同，slices的类型跟所包含的元素类型一致（不是元素的数量）。使用内置的make命令，构建一个非零的长度的空slice对象。这里我们创建了一个包含了3个字符的字符串 。(初始化为零值zero-valued)
    s := make([]string, 3)
    fmt.Println("emp:", s)
    //我们可以像数组一样进行设置和读取操作。
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
    //获取到的长度就是当时设置的长度。
    fmt.Println("len:", len(s))
    //相对于这些基本的操作，slices支持一些更加复杂的功能。有一个就是内置的append，可以在现有的slice对象上添加一个或多个值。注意要对返回的append对象重新赋值，以获取最新的添加了元素的slice对象。
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)
    //Slices也可以被复制。这里我们将s复制到了c，长度一致。
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)
</pre>
<b>slices跟arrays是两种不同的数据类型，但是他们的fmt.Println打印方式很相似。</b>

###Maps：键值对
Maps是Go语言中的关联数据类型（在其它语言中有时会被称之为哈希表[hashes]或字典[dicts]）
<pre>
package main
import "fmt"
func main (){
	m := make(map[string]int)
	m["k1"]=7
	m["k2"]=13
	fmt.Println(m)
}
Output==>
map[k1:7 k2:13]
</pre>
###引用类型
引用类型包括slice、map、channel,他们有复杂的内部结构，除了申请内存外，还需要初始化相关属性。<br>
内置函数new计算类型大小，为其分配零值内存。返回指针。而make会被编译器翻译成具体的创建函数，而其分配内存和初始化成员结构，返回对象而不是指针。
##枚举
<pre>
package main
import "fmt"

const (
	sunday =iota
	monday
	tuesdy
	wednesday
	thurday
	friday
	saturday
)
func main (){
	fmt.Println(friday)
}

Output ==>
5
</pre>
Go语言中通过关键字const来定义枚举，上面的例子中，定义了一个关于星期的枚举，当打印Friday时候输出5。打印Sunday输出0。其实，在Go语言中，枚举似乎就是常量一种特殊形式，只不过在上述代码中出现了关键字<b>iota</b>，这个是一个非常有用的东西，可以帮你省写很多东西，在上面他会初始化为0，然后每一行就会增加1，因此可以认为是一个自增量。于是我们就不必这样写了：Sunday=1   Monday=2……一个iota帮你解决一切烦恼，而且在后续中还能对iota进行操作：例如可以Monday = iota*2于是Monday就等于2了。上面说到一行定义一个iota会自增赋值给常量，那么可以一行定义多个吗？答案是可以，但是必须得明确指定值，不然会报错：
<pre>
package main
import "fmt"

const (
	sunday =iota
	monday =iota*2
	tuesday
	wednesday
	thurday
	friday,satuday=15,16
)
func main (){
	fmt.Println(satuday)
}
Output ==>
16
</pre>
发现上面枚举的值都是整数，当然其它类型的也可以，只要相应的赋值就行了，如Sunday = "sun"。

##结构体
Go语言之结构体定义：
结构体，对于学过C语言的应该很熟悉，对于C这样的语言，没有类的概念，结构体在很大程度上是作为封装的主要方式，那么在Go语言中。结构体又是如何的呢？
<pre>
package main
import "fmt"
type Student struct {
	Name string
	age int
}
func main () {
	var st Student =Student {"tom",18}  //st :=Student {"tom" 18}
	st.Name ="cat"
	fmt.Printf("%s %d",st.Name ,st.age)
}
Output ==>
cat 18
</pre>
发现和c语言差不多么，如果仔细看你会发现结构体中的Name首字母N是大写的，而age的首字母a是小写的。这可不是随便的哦。虽然这里我是随便的。<b>在Go语言中如果结构的Field首字母大写，那么它是public的，可以在package外访问。而age首字母是小写的，那么它只能在本package中被访问。</b>是否和java中类的字段用private关键字或者public定义类似呢？
上述代码中我们声明并初始化st变量是一起进行的，当然也可以分开：
<pre>
if err != nil {
	return err
}

if err !=nil{
 panic(err)
}
</pre>

###基础知识
- %d 十进制整形
- %ld 十进制长整形
- %s 字符串
- %c 字符型
- %f 浮点型

###方法
一般我们把类的成员函数叫做Methods（方法）。
<pre>
package main
import "fmt"
type mystring string
func (str mystring) prefix(preStr string)(newStr mystring){
	newStr =mystring(preStr) + str
	return
}

func main (){
	var before mystring ="go"
	after := before.prefix("let`s")
	fmt.Printf("%s\n",before)
	fmt.Printf("%s\n",after)
}

Ouptut ==>
go
let`s go
</pre>
面的程序中，第4行我们定义了一种新类型mystring，其实就是string的别名。当然，你可以定义你想要的类型，比如上篇中的结构体。

<pre>
package main
import "fmt"
type Person struct {
	name string
	age int8
}
func (p Person)getName()string {
	return p.name
}
func main (){
	p:=Person{"slick",21}
	fmt.Printf("%s \n",p.getName())
	p.name="gogo"
	fmt.Printf("%s\n",p.getName())
}

Output ==>
slick 
gogo
</pre>

###GO中的接口
最基本的接口形式：
<pre>
type show interface {
	draw()
	count()
}
</pre>
和定义一个结构体类似，只不过将struct换成了interface，然后声明了两个函数：draw()和count()。就这么简单，一个接口就定义好了，那么如何实现接口呢？
<pre>
package main
import (
    "fmt"
)
 
type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
 
type Xi []int
type Xs []string
 
func (p Xi) Len() int { return len(p) }
func (p Xi) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i int, j int) { p[i], p[j] = p[j], p[i] }
 
func (p Xs) Len() int { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i int, j int) { p[i], p[j] = p[j], p[i] }
 
 
func Sort(x Sorter) {
    for i := 0; i < x.Len() - 1; i++ {
        for j := i + 1; j < x.Len(); j++ {
            if x.Less(i, j) {
                x.Swap(i, j)
            }
        }
    }
}
func main() {
    ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
    strings := Xs{"nut", "ape", "elephant", "zoo", "go"}
    Sort(ints)
    fmt.Printf("%v\n", ints)
    Sort(strings)
    fmt.Printf("%v\n", strings)
}

Output ==>
[3 8 9 10 14 17 44 67 73 89]
[ape element go nut zoo]
</pre>

##函数
<pre>
package main
import "fmt"
func say(str string,args... interface {}) (int,error){
	_,err := fmt.Printf(str,args...)
	return len(args),err
}
func main(){
	count := 1
	closure := func (msg string) {
		say("%d %s\n",count,msg)
		count++
	}
	closure("Say one")
	closure("Say again")
}

Output ==>
Say one
Say again
</pre>
 在上述的代码中，我们一共声明并定义了两个函数，一个是say，另一个则是一个匿名函数，而且这里通过匿名函数，生成了一个函数闭包。在Go语言中

使用func关键字声明一个函数。因此，如果你要声明一个函数，马上要想到func，不管是不是匿名函数，唯一的区别就是匿名函数后面没有函数名称了，直接

func（参数列表）（返回值）。从上面我们也看到了，Go语言函数的返回类型在函数名的后面，和它声明变量的类型一样，这也与大部分语言不同的。而且函数的返回值可以是一个，也可以多个。比如上面的say函数，我们就返回了两个，一个整数类型，一个error。其中整数类型的是可变参数的长度，error类型则是从fmt包中Printf函数返回的值中的其中一个，而且我们看到接受fmt.Printf()函数返回值的第一个变量我们使用了"_"符号，这个代表我们不关心第一个返回值，将它忽略。接下来再来看say函数的第二个参数，是一个...interface{}类型，三个点是Go语言的一种类型Slices，类似数组，但是有所不同，这个将在后续文章中继续介绍，既然是一个类似数组的类型，当然也可以想到可变参数可以接收任意多个，但是必须是相同类型的，而这里使用一个空接口类型作为Slices的元素类型，使得可以接收任意类型参数的元素，之后可以通过缺省参数推断出每一个元素真实的类型。
###指针
支持指针类型 *T，指针的指针 **T，以及包含包名前缀的  *<package>.T。

- 默认值 nil，没有 NULL 常量。
- 操作符 "&" 取变量地址，"*" 透过指针访问目标对象。
- 不⽀支持指针运算，不⽀支持 "->" 运算符，直接用 "." 访问目标成员。
<pre>
package main
import "fmt"
func main (){
	type data struct {a int}
	var d= data{1234}
	var p *data
	p =&d
	fmt.Printf("%p,%v\n",p,p.a)
}
Output ==>
0x08020022f0,1234
</pre>
<br>格式控制符“%p”中的p是pointer（指针）的缩写.
<br>最简单的方法是用"%v"标志，它可以以适当的格式输出任意的类型（包括数组和结构）。下面是解释%v的乱入程序.<br>
//start
<pre>
package main
import "fmt"
 type T struct {
         a int
         b string
}
func main(){
        t := T{77, "Sunset Strip"}
        a := []int{1, 2, 3, 4}
        fmt.Printf("%v %v %v\n", u64, t, a)
}
Output ==>
 18446744073709551615 {77 Sunset Strip} [1 2 3 4]
</pre>
如果是使用"Print"或"Println"函数的话，甚至不需要格式化字符串。这些函数会针对数据类型 自动作转换。"Print"函数默认将每个参数以"%v"格式输出，"Println"函数则是在"Print"函数 的输出基础上增加一个换行。一下两种输出方式和前面的输出结果是一致的。
<pre>
 fmt.Print(u64, " ", t, " ", a, "\n")
 fmt.Println(u64, t, a)
</pre>

//end
#####不能对指针做加减法等运算。
下面的这个是错误的：
<pre>
x := 1234
p := &x
p++  //Error :不能对指针进行运算
</pre>
那么，怎么对指针进行运算操作呢？将 Pointer 转换成 uintptr，可变相实现指针运算。
<pre>
package main
import "fmt"
import "unsafe"
func main() {
	d := struct {
	s string
	x int
	}{"abc", 100}
	p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(d.x) // uintptr + offset
	p2 := unsafe.Pointer(p) // uintptr -> Pointer
	px := (*int)(p2) // Pointer -> *int
	*px = 200   //d.x = 200
	fmt.Printf("%#v\n", d)
}

Output ==>
struct {s string; x int}{s:"abc",x:200}
</pre>
####注意：GC 把 uintptr 当成普通整数对象，它⽆无法阻⽌止 "关联" 对象被回收。