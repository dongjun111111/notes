19
0
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18##不可变字符串
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

###接口
go语言中的接口并不是其他语言（C++,java等）中所提供的接口概念。go:非侵入式接口。
一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口：
定义一个File类：
<pre>
type File struct {
	//...
}
func (f *File) Read(buf []byte) (n int,err error)
func (f *File) Write(buf []byte) (n int ,err error)
func (f *File) Seek(off int64,whence int) (pos int64,err error)
func (f *File) Close() error
//上面定义了一个File类，并实现有Read(),Write(),Seek(),Close()等方法。实现一个基于File类的接口：
type IFile interface {
	Read(buf []byte) (n int,err error)
	Write(buf []byte) (n int,err error)
	Seek(off int64,whence int) (pos int64,err error)
	Close() error
}
type IReader interface {
	Read(buf []byte) (n int ,err error)
}
type IWriter interface {
	Write(buf []byte) (n int ,err error)
}
type ICloser interface {
	Close() error
}
</pre>
尽管File类并没有从这些接口继承，甚至可以不知道这些接口的存在，但是File类实现了这些接口，可以进行赋值:
var file1 IFile = new(File)
var file2 IReader = new(File)
var file3 IWriter = new(File)
var file4 ICloser = new(File)

####接口赋值
两种情况：

1. 将对象实例化赋值给接口；
第一种要求实现接口要求的所有方法：
<pre>
type Interger int 
func (a Integer) Less (b Integer) bool {
	return a <b
}
func (a *Integer) Add (b Integer){
	*a +=b
}
//对应的，我们定义接口LessAdder,如下：
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)  
}
</pre>
2. 将一个接口赋值给另一个接口。


###匿名函数
<pre>
package main 
import "fmt"
func main(){
	var j int = 5
	a := func()(func()){
		var i int = 10
		return func () {
			fmt.Printf("i,j :%d,%d \n",i,j)
		}
	}()
	a()
	j *=2
	a()
}
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
range关键字用于循环遍历数组，切片，管道或映射项目。数组和切片，它返回项目作为整数的索引。映射返回下一个键 - 值对的键。无论是范围返回一个或两个值。
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

<pre>
package main
import "fmt"
func main {
   /* create a slice */
   numbers := []int{0,1,2,3,4,5,6,7,8} 
   /* print the numbers */
   for i:= range numbers {
      fmt.Println("Slice item",i,"is",numbers[i])
   } 
   /* create a map*/
   coutryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}
   /* print map using keys*/
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   /* print map using key-value*/
   for country,capital := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",capital)
   }
}
output ==>
Slice item 0 is 0
Slice item 1 is 1
Slice item 2 is 2
Slice item 3 is 3
Slice item 4 is 4
Slice item 5 is 5
Slice item 6 is 6
Slice item 7 is 7
Slice item 8 is 8
Capital of France is Paris
Capital of Italy is Rome
Capital of Japan is Tokyo
Capital of France is Paris
Capital of Italy is Rome
Capital of Japan is Tokyo
</pre>
数组range用法
<pre>
package main 
import "fmt"
func main () {
	arr := [5]int{234,3,56,4,3}
	for d := range arr {
	fmt.Println(d)
}
}
</pre>

<pre>
package main
import "fmt"
func my(array [5]int){
	array[0] =100
	fmt.Println("array value:",array)
}
func main (){
	array := [5]int {1,23,4,5,6}
	my(array)
	fmt.Println("array:values:",array)
}
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


###Slices：切片
Slices是Go语言中的关键数据类型，它有比数组（arrays）更强的访问接口。
Go编程切片是一种抽象了Go编程数组。由于Go编程数组允许您定义的变量，可容纳同类的几个数据项类型，但它不提供任何内置的方法来动态地增加它的大小或得到一个子数组自身。切片覆盖这一限制。它提供了数组所需的多种效用函数，被广泛应用在Go编程。
####定义切片
要定义一个切片，你可以声明它作为一个数组时，不需要指定大小或使用make函数来创建。
<pre>
var numbers []int /* a slice of unspecified size */
/* numbers == []int{0,0,0,0,0}*/
numbers = make([]int,5,5) /* a slice of length 5 and capacity 5*/
</pre>

<pre>
package main 
import "fmt"
func main(){
	a := [10]int {1,2,3,4,5,6,7,8,9,10}
	s := a[2:8]  //取出2-8之间的数构成一个新数组，也就是一个切片
	s1 :=make([]int, 10 ,20)  // 取出10-20之间的数 ，结果为0,0,0...
	s2 := a[:3]  //1,2,3
	fmt.Println(s,s1,s2)
}
</pre>
####len() 和 cap() 函数
由于切片是一种抽象数组。它实际上使用数组作为底层structure.len()函数返回的元素呈现在cap()函数返回切片作为多少元素，它可以容纳的容量的切片。以下为例子来解释片的使用：
<pre>
package main
import "fmt"
func main {
   var numbers = make([]int,3,5)
   printSlice(numbers)
}
func printSlice(x []int){
   fmt.printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
output ==>
len=3 cap=5 slice=[0 0 0]
</pre>
切片：<br>
切片创建有两种形式：

* 基于数组：
<pre>
package main
import "fmt"
func main(){
	var myarr [10]int =[10]int {1,2,4,5,6,7,7,8,8,9}
	var myslice []int =myarr[:5]   //取数组前5个数据，相应的5:表示从第5位往后取
	fmt.Println("elements of myarr:")
	for 
	_,v := range myarr {//如果这里没有_,则v代表的是下标值，0-9，而并不是数组值，所有要两个值，第二个值是数组的值
		fmt.Print(v," ")
	}
	fmt.Println("\nElements of myslice :")
	for _,v :=range myslice {
		fmt.Print(v , " ")
	}
	fmt.Println()
}
</pre>

* 直接创建
<pre>
创建一个初始元素个数为5的数组切片，初始值为0：
myslice1 := make([]int,5)
创建提供初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间：
myslice2 := make([]int,5,10)
直接创建并初始化包含5个元素的数组切片：
myslice3 := []int {1,2,4,5,6}
</pre>
向一个数组切片中追加另一个数组切片，append； 注意要第二个切片后面要加上... 不然编译出错。因为按append语义，
从第二个参数起的所有参数都是待附加的元素，加上省略号相当于把myslice2包含的所有元素打散后传入。
<pre>
package main 
import "fmt"
func main(){
	myslice1 := []int {12,3,4,5}
	myslice2 := []int {7,8,9}
	myslice1 =append(myslice1,myslice2...)
	for _,i := range myslice1 {
		fmt.Print(i)
	}
}
</pre>
向数组切片追加元素：
<pre>
package main 
import "fmt"
func main(){
	myslice1 := []int {1234,54,6,7}
	myslice1 = append(myslice1,77,88)   //这里不是:=，而是=
	for _,i := range myslice1 {
		fmt.Println(i)
	}
}
</pre>
内容拷贝：
<pre>
slice1 := []int {1,2,4,5,6}
slice2 := []int {7,8,9]
copy(slice2,slice1) //只会复制slice1的前3个元素到slice2中
copy(slice1,slice2)  //只会复制slice2的3个元素到slice1的前3个位置
</pre>
####Nil 切片
如果一个切片，没有输入默认声明，它被初始化为为nil。其长度和容量都为零。
<pre>
package main
import "fmt"
func main {
   var numbers []int
   printSlice(numbers)
   if(numbers == nil){
      fmt.printf("slice is nil")
   }
}

func printSlice(x []int){
   fmt.printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
output ==>
len=0 cap=0 slice=[]
slice is nil
</pre>
####append() 和 copy() 函数
Slice允许增加使用切片的append()函数。使用copy()函数，源切片的内容复制到目标切片。下面是一个例子：
<pre>
package main

import "fmt"

func main {
   var numbers []int
   printSlice(numbers)
   
   /* append allows nil slice */
   numbers = append(numbers, 0)
   printSlice(numbers)
   
   /* add one element to slice*/
   numbers = append(numbers, 1)
   printSlice(numbers)
   
   /* add more than one element at a time*/
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)
   
   /* create a slice numbers1 with double the capacity of earlier slice*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)
   
   /* copy content of numbers to numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)   
}

func printSlice(x []int){
   fmt.printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

output ==>
len=0 cap=0 slice=[]
len=1 cap=2 slice=[0]
len=2 cap=2 slice=[0 1]
len=5 cap=8 slice=[0 1 2 3 4]
len=5 cap=16 slice=[0 1 2 3 4]
</pre>


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
map声明。创建。赋值。使用的实例：
<pre>
package main 
import "fmt"
type personinfo struct {
	id string
	name string
	address string
}
func main (){
	//map的声明 persondb是map的变量名，string是键的类型，personinfo是其中所存放的值得类型
	var persondb  map[string] personinfo
	//创建 键类型为string，值类型为personinfo的map
	persondb=make(map[string] personinfo)
	persondb["12345"] = personinfo{"12345","tom","room200"}
	persondb["1"] =personinfo{"1","jack","room333"}
    //map是一堆键值对的未排序集合。比如以身份证号作为唯一键来标识一个人的信息。
	//从这个map查找见为12345的信息
	person,ok := persondb["12345"]
	if ok {
		fmt.Println("Found person:",person.name,"with id 12345.")
	}else{
		fmt.Println("Not found")
	}

}
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
	// var st Student =Student {"tom",18}  st :=Student {"tom" 18}
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
###cap()函数
cap()函数返回的是数组切片分配的空间大小。
<pre>
package main
import "fmt"
func main(){
	myslice := make([]int,5,10)
	fmt.Println("len(myslice):",len(myslice))
	fmt.Println("cap(myslice):",cap(myslice))
}

Output==>
len(myslice):5
cap(myslice):10
</pre>
###unsafe
由于Go语言不能和C语言一样直接进行指针运算，所以需要引入unsafe包，通过它进行运算.unsafe.Pointer其实就是类似C的void *，在golang中是用于各种指针相互转换的桥梁。
在Go语言中，指针的本质是什么呢？是unsafe.Pointer和uintptr。

##import详解
我们在写Go代码的时候经常用到import这个命令用来导入包文件，而我们经常看到的方式参考如下：
<pre>
 import(
    "fmt"
)
</pre>
然后我们代码里面可以通过如下的方式调用
<pre>
 fmt.Println("hello world")
</pre>
上面这个fmt是Go语言的标准库，其实是去GOROOT环境变量指定目录下去加载该模块，当然Go的import还支持如下两种方式来加载自己写的模块：

* 相对路径
<pre>
import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import
</pre>
* 绝对路径
<pre>
import “shorturl/model”//加载gopath/src/shorturl/model模块
</pre>
上面展示了一些import常用的几种方式，但是还有一些特殊的import，让很多新手很费解，下面我们来一一讲解一下到底是怎么一回事

* 点操作<br>
我们有时候会看到如下的方式导入包
<pre>
 import(
    . "fmt"
)
</pre>
这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("helloworld")可以省略的写成Println("helloworld")

* 别名操作<br>
别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
<pre>
 import(
    f "fmt"
)
</pre>
别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("helloworld")

*  _ 操作<br>
这个操作经常是让很多人费解的一个操作符，请看下面这个import
<pre>
 import (
    "database/sql"
    _ "github.com/ziutek/mymysql/godrv"
)
</pre>
_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
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

<pre>
package main

import (
	"fmt"
)
type usb interface {   
	sname() string
 	connect() 
}
type PhoneConnecter struct {
	name string
	
}
func (pc PhoneConnecter) sname() string{
	return pc.name
}
func (pc PhoneConnecter)connect(){
	fmt.Println("Connect:",pc.name)
}
func main(){
	var a usb
	a =PhoneConnecter{"PhoneConnecter"}
	a.connect()
}

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
##go的变量类型变换
Go不支持隐式转换，必须手动指明。比如：
<pre>
var a int =2
var b float64=float64(a)
</pre>
##nil 错误
golang的nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。nil是预先说明的标识符，
也即通常意义上的关键字。在golang中，nil只能赋值给指针、channel、func、interface、map或slice类型的变量。
如果未遵循这个规则，则会引发panic。
<pre>
package main
import "fmt"
func main(){
    var a int
    var b float32
    var c bool====
    var d string
    var e []int
    var f map[string] int
    var g *int
    if nil == e{
        fmt.Print("e is nil \n")
    }
    if nil == f{
        fmt.Print("f is nil \n")
    }
    fmt.Print(a,b,c,d,e,f,g)
}
</pre>
##自动类型转换
<pre>
package main
import "fmt"
func main(){
	var b string
	b="Hello world"
	fmt.Print(b)
}
</pre>
上面的相当于：
<pre>
package main
import "fmt"
func main(){
	b := "Hello world"
	fmt.Print(b)
}
</pre>
go语言编译器自动会推断变量b的类型。

##Go中字符编码问题导致的len问题
Go中默认是UTF-8。
<pre>
package main
import "fmt"
func main () {
	a := "fffs"
	fmt.Println(a)
	fmt.Printf("%d\n",len(a))
}
Output==>
fffs
4
</pre>
###一个疑问： 
当一个字符串中既有英文又有中文时，会出现字符编码错误提示，待解决。


###指针
####什么是指针？
指针是一个变量，其值是另一个变量的地址，所述存储器位置，即，直接地址。就像变量或常量，必须声明指针之前，可以用它来存储任何变量的地址。指针变量声明的一般形式是：
<pre>
var var_name *var-type
</pre>
####如何使用指针？
有一些重要的操作，我们使用指针非常频繁。 （a）定义一个指针变量（b）分配一个变量的指针；（c）在指针变量的地址，可用地址来访问它的值。这可通过使用一元运算符 * ，返回位于其操作数所指定的地址的变量的值。
####在Go中的nil指针
Go语言编译一个 nil 值赋给一个没有被确切的地址分配的指针变量。这样做是在变量声明时，分配 nil 指针被称为nil指针。
<pre>
package main
import "fmt"
func main() {
   var  ptr *int
   fmt.Printf("The value of ptr is : %x\n", ptr  )
}
output ==>
The value of ptr is 0
</pre>
支持指针类型 *T，指针的指针 **T，以及包含包名前缀的  *<package>.T。

- 默认值，没有 NULL 常量。
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
	p := uintptr(unsafe.Pointer(&d)) /*  *struct -> Pointer -> uintptr  */
	p += unsafe.Offsetof(d.x) /* uintptr + offset */
	p2 := unsafe.Pointer(p) /* uintptr -> Pointer  */
	px := (*int)(p2) /* Pointer -> *int  */
	*px = 200   //d.x = 200
	fmt.Printf("%#v\n", d)
}

Output ==>
struct {s string; x int}{s:"abc",x:200}
</pre>
####注意：GC 把 uintptr 当成普通整数对象，它⽆无法阻⽌止 "关联" 对象被回收。
###Go语言指针数组
<pre>
package main
import "fmt"
const MAX int = 3
func main() {
   a := []int{10,100,200}
   var i int
   for i = 0; i < MAX; i++ {
      fmt.Printf("Value of a[%d] = %d\n", i, a[i] )
   }
}

output ==>
Value of a[0] = 10
Value of a[1] = 100
Value of a2] = 200
</pre>
##感悟
###Channel与锁谁轻量
Channel和锁谁轻量? 一句话告诉你: Channel本身用锁实现的. 因此在迫不得已时, 还是尽量减少使用Channel, 但Channel属于语言层支持, 适度使用, 可以改善代码可读写
###设计
踏入Golang, 就不要尝试设计模式
传统的OO在这里是非法的, 尝试模拟只是一种搞笑
把OO在Golang里换成复合+接口
对实现者来说, 把各种结构都复合起来, 对外暴露出一个或多个接口, 接口就好像使用者在实现模型上打出的很多洞
别怕全局函数, 包(Package)可以控制全局函数使用范围.
没必要什么都用interface对外封装, struct也是一种良好的封装方法
Golang无继承, 因此无需类派生图. 没有派生这种点对点的依赖, 因此不会在大量类关系到来时, 形成繁杂不可变化的树形结构
###容器
用了很长时间map, 才发现Golang把map内建为语言特性时, 已经去掉了外置型map的api特性. 一切的访问和获取都是按照语言特性来做的, 原子化
数组可以理解为底层对象, 你平时用的都是切片, 不是数组, 切片就是指针, 指向数组. 切片是轻量的, 即便值拷贝也是低损耗的
###内存
Golang在实际运行中, 你会发现内存可能会疯涨. 但跑上一段时间后, 就保持稳定. 这和Golang的内存分配, 垃圾回收有一定的关系
现代的编程语言的内存管理不会很粗暴的直接从OS那边分配很多内存. 而是按需的不断分配成块的内存.
对于非海量级应用, Golang本身的内存模型完全可以撑得下来. 无需像C++一样, 每个工程必做内存池和线程池
###错误
觉得Golang不停的处理err? 那是因为平时在其他语言根本没处理过错误, 要不然就是根部一次性try过所有的异常, 这是一种危险的行为
panic可以被捕获, 因此编写服务器时, 可以做到不挂
###危险的interface{}
这东西就跟C/C++里的void*一样的危险, nil被interface{}包裹后不会等于nil相等, 但print出来确实是nil
模板估计可以解决容器内带interface{}的问题. 但新东西引入, 估计又会让现在的哲学一些凌乱
###初学Tips
语言学习按照官网的教学走, 跑完基本就会了
下载一个LiteIDE, 配合Golang的Runtime,基本开环境就有了
Golang的类库设计方式和C#/C++都不同, 如果有Python经验的会感觉毫无违和感
有一万个理由造轮子都请住手, 类库里有你要的东西
写大工程请搜索: Golang项目目录结构组织
Golang语言本身本人没有发现bug, 即便有也早就被大神们捉住了. 唯一的一个感觉貌似bug的, 经常是结构体成员首字母小写, 但是json又无法序列化出来…
慎用cgo. 官方已经声明未来对cgo不提供完整兼容性. 任何一门语言在早期都需要对C做出支持, 但后期完善后的不兼容都是常态。
###golang的time.Format的坑
golang的time.Format设计的和其他语言都不一样, 其他语言总是使用一些格式化字符进行标示, 而golang呢, 查了网上一些坑例子 自己查了下golang的源码, 发现以下代码
// String returns the time formatted using the format string
//  "2006-01-02 15:04:05.999999999 -0700 MST"
func (t Time) String() string {
    return t.Format("2006-01-02 15:04:05.999999999 -0700 MST")
}
尝试将2006-01-02 15:04:05写入到自己的例子中
func nowTime() string {
    return time.Now().Format("2006-01-02 15:04:05")
}
结果返回正确. 询问了下, 据说这个日期是golang诞生的日子… 咋那么自恋呢…

###goruntime

- 尽可能实现无锁机制;
- CGO有限制的使用，它会将该回收的资源推迟到下一次才对其进行回收操作;

![]("image/image1.png")

![]("image/image2.png")

![]("image/image3.png")

![]("image/image4.png")

![]("image/image5.png")

![]("image/image6.png")

![]("image/image7.png")




###go的保留字/关键字(25)
break default func interface select case defer go map struct chan else goto 
package switch const fallthrough if range type continue for import return var <br>
var和const参考2.2Go语言基础里面的变量和常量申明<br>
package和import已经有过短暂的接触<br>
func 用于定义函数和方法<br>
return 用于从函数返回<br>
defer 用于类似析构函数<br>
go 用于并发<br>
select 用于选择不同类型的通讯<br>
interface 用于定义接口<br>
struct 用于定义抽象数据类型<br>
break、case、continue、for、fallthrough、else、if、switch、goto、default流程控制<br>
chan用于channel通讯<br>
type用于声明自定义类型<br>
map用于声明map类型数据<br>
range用于读取slice、map、channel数据<br>

###go数据类型
boolean numeric string derived(指针类型，数组类型，联盟类型，函数类型，切片类型，接口类型，地图类型，管道类型) <br>
####整型：
uint8(0-255)8位无符号整数 uint16(0-65535) uint32(0-4294967295) uint64(0-big) int8(-128 - 127)有符号8位整数 int16(-32768 - 32767) int32 int64
####浮点类型:
float32 float64 complex64 complex128
####其它数值类型:
byte（相当于uint8）
rune (相当于uint32)
uintptr (一个无符号整数来存储指针值的解释的比特位)
<pre>
package main
import "fmt"
func main() {
   var a, b, c = 3, 4, "foo"  
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)
}

Output ==>
3
4
foo
a is of type int
b is of type int
c is of type string
</pre>
%T输出该变量的数据类型<br>
<b>const 关键字</b>
<pre>
  const LENGTH int = 10
  const WIDTH int = 5  
</pre>
这是一个良好的编程习惯大写定义常量。
###Go语言其它运算符
还有其他一些重要的运算符，包括sizeof和?:在Go语言中也支持。
& 返回一个变量的地址 &a; 将得到变量的实际地址  <br>
* 指针的变量  *a; 将指向一个变量
<pre>
package main
import "fmt"
func main() {
   var a int = 4
   var b int32
   var c float32
   var ptr *int
   /* example of type operator */
   fmt.Printf("Line 1 - Type of variable a = %T\n", a );
   fmt.Printf("Line 2 - Type of variable b = %T\n", b );
   fmt.Printf("Line 3 - Type of variable c= %T\n", c );
   /* example of & and * operators */
   ptr = &a	/* 'ptr' now contains the address of 'a'*/
   fmt.Printf("value of a is  %d\n", a);
   fmt.Printf("*ptr is %d.\n", *ptr);
}

Ouptut ==>
Line 1 - Type of variable a = int
Line 2 - Type of variable b = int32
Line 3 - Type of variable c= float32
value of a is  4
*ptr is 4.
</pre>
###表达式Switch
<pre>
package main
import "fmt"
func main() {
   /* local variable definition */
   var grade string = "B"
   var marks int = 90
   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }
   switch {
      case grade == "A" :
         fmt.Printf("Excellent!\n" )     
      case grade == "B", grade == "C" :
         fmt.Printf("Well done\n" )      
      case grade == "D" :
         fmt.Printf("You passed\n" )      
      case grade == "F":
         fmt.Printf("Better try again\n" )
      default:
         fmt.Printf("Invalid grade\n" );
   }
   fmt.Printf("Your grade is  %s\n", grade );      
}

Output ==>
Well done
Excellent!
Your grade is  A
</pre>


###表达式select
以下规则适用于select语句：

- 可以有任意数量的范围内选择一个case语句。每一种情况下后跟的值进行比较，以及一个冒号。
- 对于case的类型必须是一个通信通道操作。
- 当通道运行下面发生的语句这种情况将执行。在case语句中break不是必需的。
- select语句可以有一个可选默认case，它必须出现在select的结束前。缺省情况下，可用于执行任务时没有的情况下是真实的。在默认情况下break不是必需的。
<pre>
package main
import "fmt"
func main() {
   var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    
}   

Output==>
no communication
</pre>
###go无限循环
<pre>
package main
import "fmt"
func main() {
   for true  {
       fmt.Printf("This loop will run forever.\n");
   }
}

</pre>
tip: 按Ctrl+ C键终止无限循环.
###Go语言continue语句
在Go编程语言中的continue语句有点像break语句。不是强制终止，只是继续循环下一个迭代发生，在两者之间跳过任何代码。
<pre>
package main
import "fmt"
func main() {
   /* local variable definition */
   var a int = 10
   /* do loop execution */
   for a < 20 {
      if a == 15 {
         /* skip the iteration */
         a = a + 1;
         continue;
      }
      fmt.Printf("value of a: %d\n", a);
      a++;     
   }  
}

output ==>
value of a: 10
value of a: 11
value of a: 12
value of a: 13
value of a: 14
value of a: 16
value of a: 17
value of a: 18
value of a: 19
</pre>
###Go语言goto语句
在Go编程语言中的goto语句提供无条件跳转从跳转到标记声明的功能。
注意：使用goto语句是高度劝阻的在任何编程语言，因为它使得难以跟踪程序的控制流程，使程序难以理解，难以修改。使用一个goto任何程序可以改写，以便它不需要goto。
<pre>
package main
import "fmt"
func main() {
   /* local variable definition */
   var a int = 10

   /* do loop execution */
   LOOP: for a < 20 {
      if a == 15 {
         /* skip the iteration */
         a = a + 1
         goto LOOP
      }
      fmt.Printf("value of a: %d\n", a)
      a++     
   }  
}

Output ==>
value of a: 10
value of a: 11
value of a: 12
value of a: 13
value of a: 14
value of a: 16
value of a: 17
value of a: 18
value of a: 19
</pre>

另一个例子：
<pre>
func myFunc() {
    i := 0
Here:   //这行的第一个词，以冒号结束作为标签
    println(i)
    i++
    goto Here   //跳转到Here去
}
</pre>
注意：标签名是大小写敏感的。

###嵌套for
下面的程序使用嵌套for循环从2至100找出的素数.
<pre>
package main
import "fmt"
func main() {
   /* local variable definition */
   var i, j int

   for i=2; i < 100; i++ {
      for j=2; j <= (i/j); j++ {
         if(i%j==0) {
            break; // if factor found, not prime
         }
      }
      if(j > (i/j)) {
         fmt.Printf("%d is prime\n", i);
      }
   }  
}

output ==>
2 is prime
3 is prime
5 is prime
7 is prime
11 is prime
13 is prime
17 is prime
19 is prime
23 is prime
29 is prime
31 is prime
37 is prime
41 is prime
43 is prime
47 is prime
53 is prime
59 is prime
61 is prime
67 is prime
71 is prime
73 is prime
79 is prime
83 is prime
89 is prime
97 is prime
</pre>
###从函数返回多个值
<pre>
package main
import "fmt"
func swap(x, y string) (string, string) {
   return y, x
}
func main() {
   a, b := swap("Mahesh", "Kumar")
   fmt.Println(a, b)
}
Output ==>
Kumar Mahesh
</pre>
###Go语言按值调用
<pre>
package main
import "fmt"
func main() {
   /* local variable definition */
   var a int = 100
   var b int = 200

   fmt.Printf("Before swap, value of a : %d\n", a )
   fmt.Printf("Before swap, value of b : %d\n", b )

   /* calling a function to swap the values */
   swap(a, b)

   fmt.Printf("After swap, value of a : %d\n", a )
   fmt.Printf("After swap, value of b : %d\n", b )
}
func swap(x, y int) int {
   var temp int

   temp = x /* save the value of x */
   x = y    /* put y into x */
   y = temp /* put temp into y */
   return temp;
}

Output ==>
Before swap, value of a :100
Before swap, value of b :200
After swap, value of a :100
After swap, value of b :200
</pre>
这表明，参数值没有被改变，虽然它们已经在函数内部改变。
###Go语言参考调用
通过传递函数参数拷贝参数的地址到形式参数的参考方法调用。在函数内部，地址是访问调用中使用的实际参数。这意味着，对参数的更改会影响传递的参数。
要通过引用传递的值，参数的指针被传递给函数就像任何其他的值。所以，相应的，需要声明函数的参数为指针类型如下面的函数swap()，它的交换两个整型变量的值指向它的参数。
<pre>
package main
import "fmt"
func main() {
   /* local variable definition */
   var a int = 100
   var b int= 200

   fmt.Printf("Before swap, value of a : %d\n", a )
   fmt.Printf("Before swap, value of b : %d\n", b )

   /* calling a function to swap the values.
   * &a indicates pointer to a ie. address of variable a and 
   * &b indicates pointer to b ie. address of variable b.
   */
   swap(&a, &b)

   fmt.Printf("After swap, value of a : %d\n", a )
   fmt.Printf("After swap, value of b : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* save the value at address x */
   *x = *y    /* put y into x */
   *y = temp    /* put temp into y */
}

output ==>
Before swap, value of a :100
Before swap, value of b :200
After swap, value of a :200
After swap, value of b :100
</pre>

###Go语言函数作为值
Go编程语言提供灵活性，以动态创建函数，并使用它们的值。在下面的例子中，我们已经与初始化函数定义的变量。此函数变量的目仅仅是为使用内置的Math.sqrt()函数。下面是一个例子：
<pre>
package main
import (
   "fmt"
   "math"
)
func main(){
   /* declare a function variable */
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }
   /* use the function */
   fmt.Println(getSquareRoot(9))
}
output ==>
3
</pre>
###Go语言函数闭合
Go编程语言支持匿名函数其可以作为函数闭包。当我们要定义一个函数内联不传递任何名称，它可以使用匿名函数。在我们的例子中，我们创建了一个函数getSequence()将返回另一个函数。该函数的目的是关闭了上层函数的变量i 形成一个闭合。下面是一个例子：
<pre>
package main
import "fmt"
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
	  return i  
   }
}
func main(){
   /* nextNumber is now a function with i as 0 */
   nextNumber := getSequence()  
   /* invoke nextNumber to increase i by 1 and return the same */
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   /* create a new sequence and see the result, i is 0 again*/
   nextNumber1 := getSequence()  
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}

output ==>
1
2
3
1
2
</pre>
###Go语言方法
Go编程语言支持特殊类型的函数调用的方法。在方法声明的语法中，“接收器”的存在是为了表示容器中的函数。该接收器可用于通过调用函数“.”运算符。下面是一个例子：
<pre>
package main
import (
   "fmt"
   "math"
)
/* define a circle */
type Circle strut {
   x,y,radius float64
}
/* define a method for circle */
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}
func main(){
   circle := Circle(x:0, y:0, radius:5)
   fmt.Printf("Circle area: %f", circle.area())
}

output ==>
Circle area: 78.539816
</pre>

###Go语言范围规则
在任何编程程序的作用域，其中一个定义的变量可以有它的存在，超出该变量的区域就不能访问。有三个地方变量可以在Go编程语言声明如下：

- 内部函数或这就是所谓的局部变量块
- 所有函数的外面的变量称为全局变量
- 在这被称为形式参数函数的参数的定义
- 让我们来解释一下什么是局部和全局变量和形式参数。

####局部变量
<pre>
package main
import "fmt"
func main() {
   /* local variable declaration */
   var a, b, c int 
   /* actual initialization */
   a = 10
   b = 20
   c = a + b
   fmt.Printf ("value of a = %d, b = %d and c = %d\n", a, b, c)
}
</pre>

####全局变量
全局变量函数的定义之外，通常在程序的顶部。全局变量的值在整个项目的生命周期，它们可以在里面任意的程序中定义的函数中访问。
全局变量可以被任何函数访问。也就是说，全局变量可以在整个程序中使用在它声明之后。下面是使用全局和局部变量的例子：
<pre>
package main
import "fmt"
/* global variable declaration */
var g int
func main() {
   /* local variable declaration */
   var a, b int
   /* actual initialization */
   a = 10
   b = 20
   g = a + b
   fmt.Printf("value of a = %d, b = %d and g = %d\n", a, b, g)
}
</pre>

####形式参数
<pre>
package main
import "fmt"
/* global variable declaration */
var a int = 20;
func main() {
   /* local variable declaration in main function */
   var a int = 10
   var b int = 20
   var c int = 0
   fmt.Printf("value of a in main() = %d\n",  a);
   c = sum( a, b);
   fmt.Printf("value of c in main() = %d\n",  c);
}
/* function to add two integers */
func sum(a, b int) int {
   fmt.Printf("value of a in sum() = %d\n",  a);
   fmt.Printf("value of b in sum() = %d\n",  b);
   return a + b;
}
output ==>
value of a in main() = 10
value of a in sum() = 10
value of b in sum() = 20
value of c in main() = 30
</pre>

####初始化局部和全局变量
当局部变量作为全局变量被初始化其对应值为0。指针被初始化为nil。
###Go语言数组
声明数组
<pre>
var balance [10] float32
</pre>
初始化数组
<pre>
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
</pre>
访问数组元素：
<pre>
package main
import "fmt"
func main() {
   var n [10]int /* n is an array of 10 integers */
   var i,j int
   /* initialize elements of array n to 0 */         
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* set element at location i to i + 100 */
   }
   /* output each array element's value */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}
output ==>
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109
</pre>

####访问二维数组元素
<pre>
package main
import "fmt"
func main() {
   /* an array with 5 rows and 2 columns*/
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   var i, j int
   /* output each array element's value */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
}
output ==>
a[0][0]: 0
a[0][1]: 0
a[1][0]: 1
a[1][1]: 2
a[2][0]: 2
a[2][1]: 4
a[3][0]: 3
a[3][1]: 6
a[4][0]: 4
a[4][1]: 8
</pre>

###Go语言传递数组到函数
如果想通过一个一维数组作为函数的参数，就必须声明函数形式参数在以下两种方式之一，以下两种声明方法产生类似的结果，因为每个告诉编译器，一个整数数组将会被接收。类似的方式，可以通过多维数组形式参数。
<pre>
package main
import "fmt"
func main() {
   /* an int array with 5 elements */
   var  balance = []int {1000, 2, 3, 17, 50}
   var avg float32
   /* pass array as an argument */
   avg = getAverage( balance, 5 ) ;
   /* output the returned value */
   fmt.Printf( "Average value is: %f ", avg );
}
func getAverage(arr []int, size int) float32 {
   var i,sum int
   var avg float32  
   for i = 0; i < size;i++ {
      sum += arr[i]
   }
   avg = float32(sum / size)
   return avg;
}
output ==>
Average value is: 214.400000
</pre>
###Go语言结构

####定义结构struct
定义一个结构，必须使用type和struct语句。该结构语句定义了一个新的数据类型，项目不止一个成员。类型语句是结构在我们的案例类型绑定的名称。

####访问结构成员
要访问结构的成员，我们使用成员访问运算符(.)。成员访问运算符是编码作为结构变量名，并且我们希望访问结构部件之间的周期。可使用struct关键字来定义结构类型的变量。
<pre>
package main
import "fmt"
type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books        /* Declare Book1 of type Book */
   var Book2 Books        /* Declare Book2 of type Book */
   /* book 1 specification */
   Book1.title = "Go Programming"
   Book1.author = "Mahesh Kumar"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407
   /* book 2 specification */
   Book2.title = "Telecom Billing"
   Book2.author = "Zara Ali"
   Book2.subject = "Telecom Billing Tutorial"
   Book2.book_id = 6495700
   /* print Book1 info */
   fmt.printf( "Book 1 title : %s\n", Book1.title)
   fmt.printf( "Book 1 author : %s\n", Book1.author)
   fmt.printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.printf( "Book 1 book_id : %d\n", Book1.book_id)
   /* print Book2 info */
   fmt.printf( "Book 2 title : %s\n", Book2.title)
   fmt.printf( "Book 2 author : %s\n", Book2.author)
   fmt.printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.printf( "Book 2 book_id : %d\n", Book2.book_id)
}

output ==>
Book 1 title : Go Programming
Book 1 author : Mahesh Kumar
Book 1 subject : Go Programming Tutorial
Book 1 book_id : 6495407
Book 2 title : Telecom Billing
Book 2 author : Zara Ali
Book 2 subject : Telecom Billing Tutorial
Book 2 book_id : 6495700
</pre>

####结构作为函数参数
<pre>
package main
import "fmt"
type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books        /* Declare Book1 of type Book */
   var Book2 Books        /* Declare Book2 of type Book */
   /* book 1 specification */
   Book1.title = "Go Programming"
   Book1.author = "Mahesh Kumar"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407
   /* book 2 specification */
   Book2.title = "Telecom Billing"
   Book2.author = "Zara Ali"
   Book2.subject = "Telecom Billing Tutorial"
   Book2.book_id = 6495700
   /* print Book1 info */
   printBook(Book1)
   /* print Book2 info */
   printBook(Book2)
}
func printBook( book Books )
{
   fmt.printf( "Book title : %s\n", book.title);
   fmt.printf( "Book author : %s\n", book.author);
   fmt.printf( "Book subject : %s\n", book.subject);
   fmt.printf( "Book book_id : %d\n", book.book_id);
}
output ==>
Book title : Go Programming
Book author : Mahesh Kumar
Book subject : Go Programming Tutorial
Book book_id : 6495407
Book title : Telecom Billing
Book author : Zara Ali
Book subject : Telecom Billing Tutorial
Book book_id : 6495700
</pre>

####指针结构
可以非常相似定义指针结构的方式，为您定义指向任何其他变量:
<pre>
var struct_pointer *Books
</pre>
使用结构指针重新写上面例子：
<pre>
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books        /* Declare Book1 of type Book */
   var Book2 Books        /* Declare Book2 of type Book */
 
   /* book 1 specification */
   Book1.title = "Go Programming"
   Book1.author = "Mahesh Kumar"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407

   /* book 2 specification */
   Book2.title = "Telecom Billing"
   Book2.author = "Zara Ali"
   Book2.subject = "Telecom Billing Tutorial"
   Book2.book_id = 6495700
 
   /* print Book1 info */
   printBook(&Book1)

   /* print Book2 info */
   printBook(&Book2)
}
func printBook( book *Books )
{
   fmt.printf( "Book title : %s\n", book.title);
   fmt.printf( "Book author : %s\n", book.author);
   fmt.printf( "Book subject : %s\n", book.subject);
   fmt.printf( "Book book_id : %d\n", book.book_id);
}
output ==>
Book title : Go Programming
Book author : Mahesh Kumar
Book subject : Go Programming Tutorial
Book book_id : 6495407
Book title : Telecom Billing
Book author : Zara Ali
Book subject : Telecom Billing Tutorial
Book book_id : 6495700
</pre>


###Go语言映射
Go编程提供另一个重要的数据类型是映射，唯一映射一个键到一个值。一个键要使用在以后检索值的对象。给定的键和值，可以在一个Map对象存储的值。值存储后，您可以使用它的键检索。

####定义映射
必须使用make函数来创建一个映射。
<pre>
/* declare a variable, by default map will be nil*/
var map_variable map[key_data_type]value_data_type

/* define the map as nil map can not be assigned any value*/
map_variable = make(map[key_data_type]value_data_type)
</pre>
例子:
<pre>
package main
import "fmt"
func main {
   var coutryCapitalMap map[string]string
   /* create a map*/
   coutryCapitalMap = make(map[string]string)
   
   /* insert key-value pairs in the map*/
   countryCapitalMap["France"] = "Paris"
   countryCapitalMap["Italy"] = "Rome"
   countryCapitalMap["Japan"] = "Tokyo"
   countryCapitalMap["India"] = "New Delhi"
   
   /* print map using keys*/
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   /* test if entry is present in the map or not*/
   captial, ok := countryCapitalMap["United States"]
   /* if ok is true, entry is present otherwise entry is absent*/
   if(ok){
      fmt.Println("Capital of United States is", capital)  
   }else {
      fmt.Println("Capital of United States is not present") 
   }
}

output ==>
Capital of India is New Delhi
Capital of France is Paris
Capital of Italy is Rome
Capital of Japan is Tokyo
Capital of United States is not present
</pre>

####delete() 函数
delete()函数是用于从映射中删除一个项目。映射和相应的键将被删除。下面是一个例子：
<pre>
package main
import "fmt"
func main {   
   /* create a map*/
   coutryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
   
   fmt.Println("Original map")   
   
   /* print map */
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   /* delete an entry */
   delete(countryCapitalMap,"France");
   fmt.Println("Entry for France is deleted")  
   
   fmt.Println("Updated map")   
   
   /* print map */
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
}

output==>
Original Map
Capital of France is Paris
Capital of Italy is Rome
Capital of Japan is Tokyo
Capital of India is New Delhi
Entry for France is deleted
Updated Map
Capital of India is New Delhi
Capital of Italy is Rome
Capital of Japan is Tokyo
</pre>

####Go语言递归
递归是以相似的方式重复项目的过程。同样适用于编程语言中，如果一个程序可以让你调用同一个函数被调用的函数，递归调用函数内使用如下。
<pre>
func recursion() {
   recursion() /* function calls itself */
}
func main() {
   recursion()
}
</pre>
Go编程语言支持递归，即要调用的函数本身。但是在使用递归时，程序员需要谨慎确定函数的退出条件，否则会造成无限循环。
递归函数是解决许多数学问题想计算一个数阶乘非常有用的，产生斐波系列等。

####数字阶乘
<pre>
package main
import "fmt"
func factorial(i int) {
   if(i <= 1) {
      return 1
   }
   return i * factorial(i - 1)
}
func main {  
    var i int = 15
    fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
}
output ==>
Factorial of 15 is 2004310016
</pre>

####斐波那契系列
<pre>
package main

import "fmt"

func fibonaci(i int) {
   if(i == 0) {
      return 0
   }
   if(i == 1) {
      return 1
   }
   return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
    var i int
    for i = 0; i < 10; i++ {
       fmt.Printf("%d\t%n", fibonaci(i))
    }    
}
outpt ==>
0	1	1	2	3	5	8	13	21	34
</pre>
###Go语言错误处理
使用返回值和错误消息。
<pre>
if err != nil {
   fmt.Println(err)
}
</pre>
例子：
<pre>
package main

import "errors"
import "fmt"
import "math"

func Sqrt(value float64)(float64, error) {
   if(value < 0){
      return 0, errors.New("Math: negative number passed to Sqrt")
   }
   return math.Sqrt(value)
}

func main() {
   result, err:= Sqrt(-1)

   if err != nil {
      fmt.Println(err)
   }else {
      fmt.Println(result)
   }
   
   result, err = Sqrt(9)

   if err != nil {
      fmt.Println(err)
   }else {
      fmt.Println(result)
   }
}
output ==>
Math: negative number passed to Sqrt
3
</pre>

###defer
Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照<b>逆序</b>执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。如下代码所示，我们一般写打开一个资源是这样操作的：
<pre>
func ReadWrite() bool {
    file.Open("file")
// 做一些工作
    if failureX {
        file.Close()
        return false
    }

    if failureY {
        file.Close()
        return false
    }

    file.Close()
    return true
}
</pre>
我们看到上面有很多重复的代码，Go的defer有效解决了这个问题。使用它后，不但代码量减少了很多，而且程序变得更优雅。在defer后指定的函数会在函数退出前调用。
<pre>
func ReadWrite() bool {
    file.Open("file")
    defer file.Close()
    if failureX {
        return false
    }
    if failureY {
        return false
    }
    return true
}
</pre>
如果有很多调用defer，那么defer是采用后进先出模式，所以如下代码会输出4 3 2 1 0
<pre>
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
</pre>

###面向对象
前面两章我们介绍了函数和struct，那你是否想过函数当作struct的字段一样来处理呢？今天我们就讲解一下函数的另一种形态，带有接收者的函数，我们称为<b>method</b>
现在假设有这么一个场景，你定义了一个struct叫做长方形，你现在想要计算他的面积，那么按照我们一般的思路应该会用下面的方式来实现:
<pre>
package main
import "fmt"
type Rectangle struct {
    width, height float64
}
func area(r Rectangle) float64 {
    return r.width*r.height
}
func main() {
    r1 := Rectangle{12, 2}
    r2 := Rectangle{9, 4}
    fmt.Println("Area of r1 is: ", area(r1))
    fmt.Println("Area of r2 is: ", area(r2))
}
</pre>
下面我们用最开始的例子用method来实现：
<pre>
package main
import (
    "fmt"
    "math"
)
type Rectangle struct {
    width, height float64
}
type Circle struct {
    radius float64
}

func (r Rectangle) area() float64 {
    return r.width*r.height
}

func (c Circle) area() float64 {
    return c.radius * c.radius * math.Pi
}


func main() {
    r1 := Rectangle{12, 2}
    r2 := Rectangle{9, 4}
    c1 := Circle{10}
    c2 := Circle{25}

    fmt.Println("Area of r1 is: ", r1.area())
    fmt.Println("Area of r2 is: ", r2.area())
    fmt.Println("Area of c1 is: ", c1.area())
    fmt.Println("Area of c2 is: ", c2.area())
}
</pre>

####method继承
前面一章我们学习了字段的继承，那么你也会发现Go的一个神奇之处，method也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。让我们来看下面这个例子:
<pre>
package main
import "fmt"

type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human //匿名字段
    school string
}

type Employee struct {
    Human //匿名字段
    company string
}

//在human上面定义了一个method
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
    mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
    sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

    mark.SayHi()
    sam.SayHi()
}
</pre>

####method重写
上面的例子中，如果Employee想要实现自己的SayHi,怎么办？简单，和匿名字段冲突一样的道理，我们可以在Employee上面定义一个method，重写了匿名字段的方法。请看下面的例子:
<pre>
package main
import "fmt"

type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human //匿名字段
    school string
}

type Employee struct {
    Human //匿名字段
    company string
}

//Human定义method
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi() {
    fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
        e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
    mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
    sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

    mark.SayHi()
    sam.SayHi()
}
</pre>

###并发
有人把Go比作21世纪的C语言，第一是因为Go语言设计简单，第二，21世纪最重要的就是并行程序设计，而Go从语言层面就支持了并行。

####1.goroutine
goroutine是Go并行设计的核心。goroutine说到底其实就是线程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。
goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。<br>
go hello(a, b, c)<br>
通过关键字go就启动了一个goroutine。我们来看一个例子
<pre>
package main

import (
    "fmt"
    "runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()
        fmt.Println(s)
    }
}

func main() {
    go say("world") //开一个新的Goroutines执行
    say("hello") //当前Goroutines执行
}

// 以上程序执行后将输出：
// hello
// world
// hello
// world
// hello
// world
// hello
// world
// hello
</pre>
我们可以看到go关键字很方便的就实现了并发编程。 上面的多个goroutine运行在同一个进程里面，共享内存数据，不过设计上我们要遵循：不要通过共享来通信，而要通过通信来共享。

####2.channels(channel)
goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通信呢，Go提供了一个很好的通信机制channel。channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：
<pre>
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
</pre>
channel通过操作符<-来接收和发送数据
<pre>
ch <- v    // 发送v到channel ch.
v := <-ch  // 从ch中接收数据，并赋值给v
</pre>
我们把这些应用到我们的例子中来：
<pre>
package main

import "fmt"

func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total += v
    }
    c <- total  // send total to c
}

func main() {
    a := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c  // receive from c

    fmt.Println(x, y, x + y)
}
</pre>
默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。

####3.Buffered Channels
上面我们介绍了默认的非缓存类型的channel，不过Go也允许指定channel的缓冲大小，很简单，就是channel可以存储多少元素。ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel。在这个channel 中，前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。
<pre>
ch := make(chan type, value)
value == 0 ! 无缓冲（阻塞）
value > 0 ! 缓冲（非阻塞，直到value 个元素）
</pre>
我们看一下下面这个例子，你可以在自己本机测试一下，修改相应的value值
<pre>
package main

import "fmt"

func main() {
    c := make(chan int, 2)//修改2为1就报错，修改2为3可以正常运行
    c <- 1
    c <- 2
    fmt.Println(<-c)
    fmt.Println(<-c)
}
    //修改为1报如下的错误:
    //fatal error: all goroutines are asleep - deadlock!
</pre>

####4.Range和Close
上面这个例子中，我们需要读取两次c，这样不是很方便，Go考虑到了这一点，所以也可以通过range，像操作slice或者map一样操作缓存类型的channel，请看下面的例子:
<pre>
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 1, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x + y
    }
    close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}
</pre>
for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。
记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
另外记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的

####5.Select
我们上面介绍的都是只有一个channel的情况，那么如果存在多个channel的时候，我们该如何操作呢，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。

select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。
<pre>
package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 1, 1
    for {
        select {
        case c <- x:
            x, y = y, x + y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}
</pre>
在select里面还有default语法，select其实就是类似switch的功能，default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。
<pre>
select {
case i := <-c:
    // use i
default:
    // 当c阻塞的时候执行这里
}
</pre>

####6.超时
有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？我们可以利用select来设置超时，通过如下的方式实现：
<pre>
func main() {
    c := make(chan int)
    o := make(chan bool)
    go func() {
        for {
            select {
                case v := <- c:
                    println(v)
                case <- time.After(5 * time.Second):
                    println("timeout")
                    o <- true
                    break
            }
        }
    }()
    <- o
}
</pre>

####7.runtime goroutine
runtime包中有几个处理goroutine的函数：

- Goexit
退出当前执行的goroutine，但是defer函数还会继续调用

- Gosched
让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

- NumCPU
返回 CPU 核数量

- NumGoroutine
返回正在执行和排队的任务总数

- GOMAXPROCS
用来设置可以并行计算的CPU核数的最大值，并返回之前的值。

###指针
声明指针：
var ip *int // pointer to an integer
var fp *float32 //pointer to a float
指针的作用很多，其实说白了就是直接操作内存，好处是：

- 效率更高，这个很容易理解，直接操作内存，效率必然更高
- 可以写复杂度更高的数据
- 结构，这个也好理解，程序员可以操作内存，当然可以写出灵活、复杂的数据结构
- 编写出简洁、紧凑、高效的程序
<pre>
package main
import (
	"fmt"
)
func main () {
	type person struct {
	Age int
	phone int
	name string
	}
	var s=person{2333,1222222,""}
	var p *person
	p =&s
	fmt.Printf("%p, %v\n",p,p.phone)
}
output ==>
0xc082004640,1222222
</pre>

来自PHP的格式值，这里可能用到：

- %% - 返回一个百分号 %
- %b - 二进制数
- %c - ASCII 值对应的字符,字符型
- %d - 包含正负号的十进制数（负数、0、正数）
- %e - 使用小写的科学计数法（例如 1.2e+2）
- %E - 使用大写的科学计数法（例如 1.2E+2）
- %u - 不包含正负号的十进制数（大于等于 0）
- %f - 浮点数（本地设置）
- %F - 浮点数（非本地设置）
- %g - 较短的 %e 和 %f
- %G - 较短的 %E 和 %f
- %o - 八进制数
- %s - 字符串
- %x - 十六进制数（小写字母）
- %X - 十六进制数（大写字母）


###函数
不支持 嵌套(nested)、重载(overload) 和默认参数（default parameter）
<pre>
package main
func add (x,y int)(z int){
	z = x+ y
	return
}
func main () {
	println(add(1 ,3))
}
output ==>
4
</pre>

###defer
- 简化资源的回收
<pre>
mu.Lock()
defer mu.Unlock()
</pre>
从简化资源的释放角度看, defer 类似一个语法糖, 好像不是必须的.

- panic异常的捕获
在Go语言中, panic用于抛出异常, recover用于捕获异常. recover只能在defer语句中使用, 直接调用recover是无效的.
<pre>
package main5
import "fmt"
func main () {
	f()
	fmt.Println("returned normally from f.")
}
func f() {
	defer func () {
	if r := recover(); r != nil {
		fmt.Println("Recovered in f",r)
	}
	}()
	fmt.Println("Calling g")
	g()
	fmt.Println("Returned normally form g.")
}
func g() {
	panic("ERROR");
}
output ==>
Calling g
Recovered in f ERROR
returned normally form f.
</pre>
因此, 如果要捕获Go语言中函数的异常, 就离不开defer语句了.

- 修改返回值
<pre>
package main
func test (a,b int) (sum int) {
	defer func () {
		sum  += 2
	}()
	sum =a + b
	return sum
}

func main () {
      print(test(2,3));
}
output==>
7
</pre>
 这个特性应该只是 defer 的副作用, 具体在什么场景使用就要由开发者自己决定了.

- 安全的回收资源
<pre>
func TestFailed(t *testing.T) {  
    var wg sync.WaitGroup  
    for i := 0; i < 2; i++ {  
        wg.Add(1)  
        go func(id int) {  
            // defer wg.Done()  
            t.Fatalf("TestFailed: id = %v\n", id)  
            wg.Done()  
        }(i)  
    }  
    wg.Wait()  
} 
</pre>


###Go语言并发编程总结
Golang :不要通过共享内存来通信，而应该通过通信来共享内存。

####通过golang中的 goroutine 与sync.Mutex进行 并发同步
<pre>
import( 
    "fmt"
    "sync"
    "runtime"
)
var count int =0;
func counter(lock * sync.Mutex){
      lock.Lock()
      count++
      fmt.Println(count)
      lock.Unlock()
}
func main(){
   lock:=&sync.Mutex{}
   for i:=0;i<10;i++{
      //传递指针是为了防止 函数内的锁和 调用锁不一致
      go counter(lock)  
     }
   for{
      lock.Lock()
      c:=count
      lock.Unlock()
      ///把时间片给别的goroutine  未来某个时刻运行该routine
      runtime.Gosched()
      if c>=10{
        fmt.Println("goroutine end")
        break
        }
   }    
}
</pre>

####goroutine之间通过 channel进行通信,channel是和类型相关的 可以理解为  是一种类型安全的管道。
<pre>
package main  
import "fmt"
func Count(ch chan int) {
    ch <- 1  
    fmt.Println("Counting")
}
func main() {
    chs := make([]chan int, 10)
for i := 0; i < 10; i++ {
        chs[i] = make(chan int)
  go Count(chs[i])
  fmt.Println("Count",i)
    }
for i, ch := range chs {
  <-ch
  fmt.Println("Counting",i)
    }  
} 
</pre>

####Go语言中的select是语言级内置  非堵塞
select {
case <-chan1: // 如果chan1成功读到数据，则进行该case处理语句  
case chan2 <- 1: // 如果成功向chan2写入数据，则进行该case处理语句  
default: // 如果上面都没有成功，则进入default处理流程  
}
可以看出，select不像switch，后面并不带判断条件，而是直接去查看case语句。每个
case语句都必须是一个面向channel的操作。比如上面的例子中，第一个case试图从chan1读取
一个数据并直接忽略读到的数据，而第二个case则是试图向chan2中写入一个整型数1，如果这
两者都没有成功，则到达default语句。 

####channel 的带缓冲读取写入
之前我们示范创建的都是不带缓冲的channel，这种做法对于传递单个数据的场景可以接受，
但对于需要持续传输大量数据的场景就有些不合适了。接下来我们介绍如何给channel带上缓冲，
从而达到消息队列的效果。
要创建一个带缓冲的channel，其实也非常容易：
c := make(chan int, 1024)
在调用make()时将缓冲区大小作为第二个参数传入即可，比如上面这个例子就创建了一个大小
为1024的int类型channel，即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被
填完之前都不会阻塞。
从带缓冲的channel中读取数据可以使用与常规非缓冲channel完全一致的方法，但我们也可
以使用range关键来实现更为简便的循环读取：
<pre>
for i := range c {
    fmt.Println("Received:", i)
} 
</pre>

####用goroutine模拟生产消费者
<pre>
package main
import "fmt"
import "time"
func Producer (queue chan<- int){
        for i:= 0; i < 10; i++ {
                queue <- i  
                }
}
func Consumer( queue <-chan int){
        for i :=0; i < 10; i++{
                v := <- queue
                fmt.Println("receive:", v)
        }
}
func main(){
        queue := make(chan int, 1)
        go Producer(queue)
        go Consumer(queue)
        time.Sleep(1e9) //让Producer与Consumer完成
}
</pre>

####通过make 创建通道 
 make(c1 chan int)   创建的是 同步channel ...读写完全对应
make(c1 chan int ,10) 闯进带缓冲的通道 上来可以写10次

####随机向通道中写入0或者1 
<pre>
package main
import "fmt"
import "time"
func main(){
       ch := make(chan int, 1)
 for {
   ///不停向channel中写入 0 或者1
  select {
   case ch <- 0:
   case ch <- 1:
  }
    //从通道中取出数据
    i := <-ch
    fmt.Println("Value received:",i)
    time.Sleep(1e8)
    }
}
</pre>

无缓冲的信道是一批数据一个一个的「流进流出」<br>
缓冲信道则是一个一个存储，然后一起流出去


####带缓冲的channel 
之前创建的都是不带缓冲的channel，这种做法对于传递单个数据的场景可以接受，
但对于需要持续传输大量数据的场景就有些不合适了。接下来我们介绍如何给channel带上缓冲，
从而达到消息队列的效果。
要创建一个带缓冲的channel，其实也非常容易：
c := make(chan int, 1024)
在调用make()时将缓冲区大小作为第二个参数传入即可，比如上面这个例子就创建了一个大小
为1024的int类型channel，即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被
填完之前都不会阻塞。
从带缓冲的channel中读取数据可以使用与常规非缓冲channel完全一致的方法，但我们也可
以使用range关键来实现更为简便的循环读取：
<pre>
for i := range c {
    fmt.Println("Received:", i)
}
</pre>
实例：
<pre>
package main
import "fmt"
import "time"
func A(c chan int){
 for i:=0;i<10;i++{
        c<- i
    }
}
func B(c chan int){
 for val:=range c {
      fmt.Println("Value:",val)  
    }
}
func main(){
    chs:=make(chan int,10)
    //只要有通道操作一定要放到goroutine中否则 会堵塞当前的主线程 并且导致程序退出
    //对于同步通道 或者带缓冲的通道 一定要封装成函数 使用 goroutine 包装
    go A(chs)
    go B(chs)
    time.Sleep(1e9)
}

</pre>

####关于创建多个goroutine具体到go语言会创建多少个线程
<pre>
import "os"
func main() {
    for i:=0; i<20; i++ {
        go func() {
            for {
                b:=make([]byte, 10)
                os.Stdin.Read(b) // will block
            }
        }()
    }
    select{}
}
</pre>
上面代码会产生21个线程：
runtime scheduler(src/pkg/runtime/proc.c)会维护一个线程池，当某个goroutine被block后，scheduler会创建一个新线程给其他ready的goroutine
GOMAXPROCS控制的是未被阻塞的所有goroutine被multiplex到多少个线程上运行

####在channel中也是可以传递channel的,Go语言的channel和map  slice等一样都是原生类型

需要注意的是，在Go语言中channel本身也是一个原生类型，与map之类的类型地位一样，因
此channel本身在定义后也可以通过channel来传递。
我们可以使用这个特性来实现*nix上非常常见的管道（pipe）特性。管道也是使用非常广泛
的一种设计模式，比如在处理数据时，我们可以采用管道设计，这样可以比较容易以插件的方式
增加数据的处理流程。
下面我们利用channel可被传递的特性来实现我们的管道。 为了简化表达， 我们假设在管道中
传递的数据只是一个整型数，在实际的应用场景中这通常会是一个数据块。
首先限定基本的数据结构：
<pre>
type PipeData struct {
    value int
    handler func(int) int
    next chan int
}
</pre>
然后我们写一个常规的处理函数。我们只要定义一系列PipeData的数据结构并一起传递给
这个函数，就可以达到流式处理数据的目的：
<pre>
func handle(queue chan *PipeData) {
for data := range queue {
        data.next <- data.handler(data.value)
    }
}
</pre>

####只读只写单向channel代码例子,遵循权限最小化的原则
<pre>
package main
import "fmt"
import "time"
//接受一个参数 是只允许读取通道  除非直接强制转换 要么你只能从channel中读取数据
func sCh(ch <-chan int){
   for val:= range ch {
     fmt.Println(val)
   }
}
func main(){
    //创建一个带100缓冲的通道 可以直接写入 而不会导致 主线程堵塞
    dch:=make(chan int,100)
    for i:=0;i<100;i++{
      dch<- i  
    }
    //传递进去 只读通道
    go sCh(dch)
    time.Sleep(1e9)
}
</pre>

####channel的关闭,以及判断channel的关闭
关闭channel非常简单，直接使用Go语言内置的close()函数即可：
close(ch)
在介绍了如何关闭channel之后，我们就多了一个问题：如何判断一个channel是否已经被关
闭？我们可以在读取的时候使用多重返回值的方式：
x, ok := <-ch
这个用法与map中的按键获取value的过程比较类似，只需要看第二个bool返回值即可，如
果返回值是false则表示ch已经被关闭。

####Go的多核并行化编程    高性能并发编程 必须设置GOMAXPROCS 为最大核数目 这个值由runtime.NumCPU()获取

在执行一些昂贵的计算任务时， 我们希望能够尽量利用现代服务器普遍具备的多核特性来尽
量将任务并行化，从而达到降低总计算时间的目的。此时我们需要了解CPU核心的数量，并针对
性地分解计算任务到多个goroutine中去并行运行。
下面我们来模拟一个完全可以并行的计算任务：计算N个整型数的总和。我们可以将所有整
型数分成M份，M即CPU的个数。让每个CPU开始计算分给它的那份计算任务，最后将每个CPU
的计算结果再做一次累加，这样就可以得到所有N个整型数的总和：
<pre>
type Vector []float64
// 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
for ; i < n; i++ {
         v[i] += u.Op(v[i])
     }
     c <- 1       
// 发信号告诉任务管理者我已经计算完成了
}
const NCPU = 16     
// 假设总共有16核   
func (v Vector) DoAll(u Vector) {   
    c := make(chan int, NCPU)  // 用于接收每个CPU的任务完成信号   
for i := 0; i < NCPU; i++ {   
go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
    } 
// 等待所有CPU的任务完成
for i := 0; i < NCPU; i++ {   
<-c    // 获取到一个数据，表示一个CPU计算完成了
    }
// 到这里表示所有计算已经结束
}
</pre>
这两个函数看起来设计非常合理。DoAll()会根据CPU核心的数目对任务进行分割，然后开
辟多个goroutine来并行执行这些计算任务。
是否可以将总的计算时间降到接近原来的1/N呢？答案是不一定。如果掐秒表（正常点的话，
应该用7.8节中介绍的Benchmark方法） ，会发现总的执行时间没有明显缩短。再去观察CPU运行
状态， 你会发现尽管我们有16个CPU核心， 但在计算过程中其实只有一个CPU核心处于繁忙状态，
这是会让很多Go语言初学者迷惑的问题。
官方的答案是，这是当前版本的Go编译器还不能很智能地去发现和利用多核的优势。虽然
我们确实创建了多个goroutine，并且从运行状态看这些goroutine也都在并行运行，但实际上所有
这些goroutine都运行在同一个CPU核心上， 在一个goroutine得到时间片执行的时候， 其他goroutine
都会处于等待状态。从这一点可以看出，虽然goroutine简化了我们写并行代码的过程，但实际上
整体运行效率并不真正高于单线程程序。
在Go语言升级到默认支持多CPU的某个版本之前，我们可以先通过设置环境变量
GOMAXPROCS的值来控制使用多少个CPU核心。具体操作方法是通过直接设置环境变量
GOMAXPROCS的值，或者在代码中启动goroutine之前先调用以下这个语句以设置使用16个CPU
核心：
runtime.GOMAXPROCS(16)
到底应该设置多少个CPU核心呢，其实runtime包中还提供了另外一个函数NumCPU()来获
取核心数。可以看到，Go语言其实已经感知到所有的环境信息，下一版本中完全可以利用这些
信息将goroutine调度到所有CPU核心上，从而最大化地利用服务器的多核计算能力。抛弃
GOMAXPROCS只是个时间问题

####主动出让时间片给其他 goroutine 在未来的某一时刻再来执行当前goroutine
我们可以在每个goroutine中控制何时主动出让时间片给其他goroutine，这可以使用runtime
包中的Gosched()函数实现。
实际上，如果要比较精细地控制goroutine的行为，就必须比较深入地了解Go语言开发包中
runtime包所提供的具体功能。

####Go中的同步
倡导用通信来共享数据，而不是通过共享数据来进行通信，但考虑
到即使成功地用channel来作为通信手段，还是避免不了多个goroutine之间共享数据的问题，Go
语言的设计者虽然对channel有极高的期望，但也提供了妥善的资源锁方案。

####Go中的同步锁
倡导用通信来共享数据，而不是通过共享数据来进行通信，但考虑
到即使成功地用channel来作为通信手段，还是避免不了多个goroutine之间共享数据的问题，Go
语言的设计者虽然对channel有极高的期望，但也提供了妥善的资源锁方案。
对于这两种锁类型， 任何一个Lock()或RLock()均需要保证对应有Unlock()或RUnlock()
调用与之对应，否则可能导致等待该锁的所有goroutine处于饥饿状态，甚至可能导致死锁。锁的
典型使用模式如下：
<pre>
var l sync.Mutex  
func foo() {
l.Lock()  
//延迟调用 在函数退出 并且局部资源被释放的时候 调用
defer l.Unlock()  
//...
}  
</pre>
这里我们再一次见证了Go语言defer关键字带来的优雅

####全局唯一操作 sync.Once.Do();sync.atomic原子操作子包
对于从全局的角度只需要运行一次的代码，比如全局初始化操作，Go语言提供了一个Once
类型来保证全局的唯一性操作，具体代码如下：
<pre>
var a string
var once sync.Once  
func setup() {
a = "hello, world"
}  
func doprint() {
once.Do(setup)
print(a)  
}  
func twoprint() {
go doprint()
go doprint()  
}
</pre>
如果这段代码没有引入Once， setup()将会被每一个goroutine先调用一次， 这至少对于这个
例子是多余的。在现实中，我们也经常会遇到这样的情况。Go语言标准库为我们引入了Once类
型以解决这个问题。once的Do()方法可以保证在全局范围内只调用指定的函数一次（这里指
setup()函数） ，而且所有其他goroutine在调用到此语句时，将会先被阻塞，直至全局唯一的
once.Do()调用结束后才继续。
这个机制比较轻巧地解决了使用其他语言时开发者不得不自行设计和实现这种Once效果的
难题，也是Go语言为并发性编程做了尽量多考虑的一种体现。
如果没有once.Do()，我们很可能只能添加一个全局的bool变量，在函数setup()的最后
一行将该bool变量设置为true。在对setup()的所有调用之前，需要先判断该bool变量是否已
经被设置为true，如果该值仍然是false，则调用一次setup()，否则应跳过该语句。实现代码
<pre>
var done bool = false
func setup() {
a = "hello, world" 
done = true
}     
func doprint() { 
if !done {
        setup()
    }   
print(a)  
}  
</pre>
这段代码初看起来比较合理， 但是细看还是会有问题， 因为setup()并不是一个原子性操作，
这种写法可能导致setup()函数被多次调用，从而无法达到全局只执行一次的目标。这个问题的
复杂性也更加体现了Once类型的价值。
为了更好地控制并行中的原子性操作，sync包中还包含一个atomic子包，它提供了对于一
些基础数据类型的原子操作函数，比如下面这个函数：
func CompareAndSwapUint64(val *uint64, old, new uint64) (swapped bool)
就提供了比较和交换两个uint64类型数据的操作。这让开发者无需再为这样的操作专门添加
Lock操作。

###go的if else坑点
<pre>
package main
import "fmt"
func main () {
	var a int =1
	var b int =4
	if(a == b){
	fmt.Println("true");	
}else{
	fmt.Println("false");
}
}
</pre>
else必须跟在中括号后面，成一行。

####printf与println的区别
 printf可以解析变量，一般用%d %s %t %p等代替要输出的变量；println不可以，原样输出%d %T，所以一般直接把变量放在括号里面输出.

###函数返回多个值
<pre>
package main
import "fmt"
func swap (x,y string)(string , string){   //格式注意
	return y,x
}
func main() {
	a ,b := swap("fff","dddd")
	fmt.Println(a,b)
}
</pre>

####获取变量在内存中的存储位置
<pre>
package main
import "fmt"
func main (){
	var a int = 10
	fmt.Printf("%x \n",&a)
}
</pre>

####访问指针存储地址与值的方法与区别
<pre>
package main
import "fmt"
func main() {
	var a int = 30
	var ip *int
	ip = &a
	fmt.Printf("%x\n",ip)     //ip存储地址
	fmt.Printf("%d\n",*ip)    //*ip值

}
</pre>
空指针与nil
<pre>
package main
import "fmt"
func main () {
	var ptr *int
	if(ptr == nil){  //如果ptr是空指针
		fmt.Println("ptr is a null pointer")
	}
}
</pre>
指针数组
<pre>
package main
import "fmt"
const MAX int = 3
func main () {
	a := []int {10,20,30}
	var i int
	var ptr [MAX]*int
	for i =0 ;i< MAX ;i ++ {
		ptr[i] = &a[i] //整数地址赋值给指针数组
	}
	
	for i=0;i< MAX; i++ {
		fmt.Printf("a[%d] = %d \n",i,*ptr[i])
	}
}
output ==>
a[0]= 10
a[1]= 20
a[2]= 30
</pre>

###Go hello world
<pre>
package main 
import (
	"io"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter,r *http.Request){
		io.WriteString(w,"Hello ,world")
}
func main(){
	http.HandleFunc("/hello",helloHandler)
	err :=http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServe:",err.Error())
	}
}
</pre>
1. go run t.go
2. in browser,input http://localhost:8080/hello,then "hello world" is showing up.Amazing!

###MD5与sha1加密
字符串加密
<pre>
package main 
import (
	"fmt"
	"crypto/sha1"
	"crypto/md5"
)
func main(){
	TestString :="Hi,Jason"

	Md5Inst:=md5.New()
	Md5Inst.Write([]byte(TestString))
	Result:=Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n",Result)

	Sha1Inst :=sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result=Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n",Result)

}
</pre>
文件加密：
<pre>
package main 
import (
	"io"
	"fmt"
	"os"
	"crypto/md5"
	"crypto/sha1"
)
func main () {
	TestFile :="123.txt"
	infile,inerr :=os.Open(TestFile)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h,infile)
		fmt.Printf("%x %s\n",md5h.Sum([]byte("")))
		sha1h := sha1.New()
		io.Copy(sha1h,infile)
		fmt.Printf("%x %s\n",sha1h.Sum([]byte("")))
	}else{
		fmt.Println(inerr)
		os.Exit(1)
	}
}
</pre>


###代码格式化
在go的命令行工具中，用
go fmt 文件名
可以实现对代码的格式化，美观易读。

###远程import支持
go 语言可以调用远程的包。例如：
<pre>
package main
import (
	"fmt"
    "github.com/myteam/exp/crc32"
)
</pre>
然后再执行go build 或者 go install之前，只要先执行<br>
go get github.com/myteam/exp/crc32就行了。

###反射reflect（建议少用）
####简单类型的反射操作
通过使用Type和Value,我们可以对一个类型进行各项灵活的操作。
<pre>
package main 
import (
	"fmt"
	"reflect"
)
func main (){
	var x float64 =34.5
	fmt.Println("type:" ,reflect.TypeOf(x))
}

output==>
float64
</pre>
<pre>
package main 
import (
	"fmt"
	"reflect"
)
func main (){
	var x float64= 3.4
	v :=reflect.ValueOf(x) 
	fmt.Println("type:" , v.Type())
}
output==>
float64
</pre>

<pre>
package main 
import (
	"fmt"
	"reflect"
)
func main (){
	var x float64 =34.5
	fmt.Println("value:",reflect.ValueOf(x))
}
output ==>
34.5
</pre>
Kind():
<pre>
package main 
import (
	"fmt"
	"reflect"
)
func main (){
	var x float64= 3.4
	v :=reflect.ValueOf(x)   
	fmt.Println("kind is float64 :" , v.Kind() == reflect.Float64)
}
output==>
kind of float64 : true
</pre>
Canset():()这里要说明为什么要少用reflect,在欲改变变量值的时候更要慎用：
<pre>
package main 
import (
	"fmt"
	"reflect"
)
func main(){
	var x float64=4.5
	p := reflect.ValueOf(&x)  //得到x的地址
	fmt.Println("settability of v :",p.CanSet())
}
output==>
settability of v : false
</pre>
这说明：在上面的情况下不能通过反射对x重新赋值，这就是reflect慎用的原因，可以通用
下面的方法实现：
<pre>
package main 
import (
	"fmt"
	"reflect"
)
func main(){
	var x float64=4.5
	p := reflect.ValueOf(&x)  //得到x的地址
	fmt.Println("settability of v :",p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v :" ,v.CanSet())
}	
output ==>
settability of v : false
settability of v : true  //这时候可以重新赋值
</pre>
<pre>
package main 
import (
	"fmt"
	"reflect"
)
func main(){
	var x float64=4.5
	p := reflect.ValueOf(&x)  //得到x的地址
	fmt.Println("settability of v :",p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v :" ,v.CanSet())
	v.SetFloat(7.6)
	fmt.Println(v.Interface())
	fmt.Println(x)
}	
output ==>
settability of v : false
settability of v : true
7.6
7.6
</pre>
这时候，v已经重新赋值。

####对结构的反射操作
获取一个结构中的所有成员的值：
<pre>
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	S string
}

func main() {
	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d : %s %s =%v \n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

output ==>
0: A int= 203
1: B string = mh203
</pre>


###A.1.2 完整包列表
<pre>
bufio 实现缓冲的I/O
bytes 提供了对字节切片操作的函数
crypto 收集了常见的加密常数
errors 实现了操作错误的函数
Expvar 为公共变量提供了一个标准的接口，如服务器中的运算计数器
flag 实现了命令行标记解析
fmt 实现了格式化输入输出
hash 提供了哈希函数接口
html 实现了一个HTML5兼容的分词器和解析器
image 实现了一个基本的二维图像库
io 提供了对I/O原语的基本接口
log 它是一个简单的记录包，提供最基本的日志功能
math 提供了一些基本的常量和数学函数
mine 实现了部分的MIME规范
net 提供了一个对UNIX网络套接字的可移植接口，包括TCP/IP、UDP域名解析和
UNIX域套接字
os 为操作系统功能实现了一个平台无关的接口
path 实现了对斜线分割的文件名路径的操作
reflect 实现了运行时反射，允许一个程序以任意类型操作对象
regexp 实现了一个简单的正则表达式库
runtime 包含与Go运行时系统交互的操作，如控制goroutine的函数
sort 提供对集合排序的基础函数集
strconv 实现了在基本数据类型和字符串之间的转换
strings 实现了操作字符串的简单函数
sync 提供了基本的同步机制，如互斥锁
syscall 包含一个低级的操作系统原语的接口
testing 提供对自动测试Go包的支持
time 提供测量和显示时间的功能
unicode Unicode编码相关的基础函数
archive tar 实现对tar压缩文档的访问
zip 提供对ZIP压缩文档的读和写支持
bzip2 实现了bzip2解压缩
flate 实现了RFC 1951中所定义的DEFLATE压缩数据格式
gzip 实现了RFC 1951中所定义的gzip格式压缩文件的读和写
lzw 实现了Lempel-Ziv-Welch编码格式的压缩的数据格式，参见T. A. Welch, A
Technique for High-Performance Data Compression, Computer, 17(6) (June 1984), pp
compress
zlib 实现了RFC 1950中所定义的zlib格式压缩数据的读和写
heap 提供了实现heap.Interface接口的任何类型的堆操作
list 实现了一个双链表
container
ring 实现了对循环链表的操作
aes 实现了AES加密（以前的Rijndael），详见美国联邦信息处理标准（197号文）
cipher 实现了标准的密码块模式，该模式可包装进低级的块加密实现中
des 实现了数据加密标准（Data Encryption Standard，DES）和三重数据加密算法（Triple
Data Encryption Algorithm，TDEA），详见美国联邦信息处理标准（46-3号文）
dsa 实现了FIPS 186-3所定义的数据签名算法（Digital Signature Algorithm）
ecdsa 实现了FIPS 186-3所定义的椭圆曲线数据签名算法（Elliptic Curve Digital Signature
Algorithm）
elliptic 实现了素数域上几个标准的椭圆曲线
hmac 实现了键控哈希消息身份验证码（Keyed-Hash Message Authentication Code，
HMAC），详见美国联邦信息处理标准（198号文）
md5 实现了RFC 1321中所定义的MD5哈希算法
rand 实现了一个加密安全的伪随机数生成器
rc4 实现了RC4加密，其定义见Bruce Schneier的应用密码学（Applied Cryptography）
rsa 实现了PKCS#1中所定义的RSA加密
sha1 实现了RFC 3174中所定义的SHA1哈希算法
sha256 实现了FIPS 180-2中所定义的SHA224和SHA256哈希算法
sha512 实现了FIPS 180-2中所定义的SHA384和SHA512哈希算法
subtle 实现了一些有用的加密函数，但需要仔细考虑以便正确应用它们
tls 部分实现了RFC 4346所定义的TLS 1.1协议
x509 可解析X.509编码的键值和证书
crypto
x509/pkix 包含用于对X.509证书、CRL和OCSP的ASN.1解析和序列化的共享的、低级的结构
database sql 围绕SQL提供了一个通用的接口
sql/driver 定义了数据库驱动所需实现的接口，同sql包的使用方式
dwarf 提供了对从可执行文件加载的DWARF调试信息的访问，这个包对于实现Go语言
的调试器非常有价值
elf 实现了对ELF对象文件的访问。ELF是一种常见的二进制可执行文件和共享库的
文件格式。Linux采用了ELF格式
gosym 访问Go语言二进制程序中的调试信息。对于可视化调试很有价值
macho 实现了对 http://developer.apple.com/mac/library/documentation/DeveloperTools/Conceptual/
MachORuntime/Reference/reference.html 所定义的Mach-O对象文件的访问
debug
pe 实现了对PE（Microsoft Windows Portable Executable）文件的访问
ascii85 实现了ascii85数据编码，用于btoa工具和Adobe’s PostScript以及PDF文档格式
asn1 实现了解析DER编码的ASN.1数据结构，其定义见ITU-T Rec X.690
base32 实现了RFC 4648中所定义的base32编码
base64 实现了RFC 4648中所定义的base64编码
binary 实现了在无符号整数值和字节串之间的转化，以及对固定尺寸值的读和写
csv 可读和写由逗号分割的数值（csv）文件
gob 管理gob流——在编码器（发送者）和解码器（接收者）之间进行二进制值交换
hex 实现了十六进制的编码和解码
json 实现了定义于RFC 4627中的JSON对象的编码和解码
pem 实现了PEM（Privacy Enhanced Mail）数据编码
encoding
xml 实现了一个简单的可理解XML名字空间的XML 1.0解析器
ast 声明了用于展示Go包中的语法树类型
build 提供了构建Go包的工具
doc 从一个Go AST（抽象语法树）中提取源代码文档
parser 实现了一个Go源文件解析器
printer 实现了对AST（抽象语法树）的打印
scanner 实现了一个Go源代码文本的扫描器
go
token 定义了代表Go编程语言中词法标记以及基本操作标记（printing、predicates）的常
量
adler32 实现了Adler-32校验和
crc32 实现了32位的循环冗余校验或CRC-32校验和
crc64 实现了64位的循环冗余校验或CRC-64校验和
hash
fnv 实现了Glenn Fowler、Landon Curt Noll和Phong Vo所创建的FNV-1和FNV-1a未加
密哈希函数
html template 它自动构建HTML输出，并可防止代码注入
color 实现了一个基本的颜色库
draw 提供一些做图函数
gif 实现了一个GIF图像解码器
jpeg 实现了一个JPEG图像解码器和编码器
image
png 实现了一个PNG图像解码器和编码器
index suffixarray 通过构建内存索引实现的高速字符串匹配查找算法
io ioutil 实现了一些实用的I/O函数
log syslog 提供了对系统日志服务的简单接口
big 实现了多精度的算术运算（大数）
cmplx 为复数提供了基本的常量和数学函数
Math
rand 实现了伪随机数生成器
mime multipart 实现了在RFC 2046中定义的MIME多个部分的解析
http 提供了HTTP客户端和服务器的实现
mail 实现了对邮件消息的解析
rpc 提供了对一个来自网络或其他I/O连接的对象可导出的方法的访问
smtp 实现了定义于RFC 5321中的简单邮件传输协议（Simple Mail Transfer Protocol)
textproto 实现了在HTTP、NNTP和SMTP中基于文本的通用的请求/响应协议
url 解析URL并实现查询转义
http/cgi 实现了定义于RFC 3875中的CGI（通用网关接口）
http/fcgi 实现了FastCGI协议
http/httptest 提供了一些HTTP测试应用
http/httputil 提供了一些HTTP应用函数，这些是对net/http包中的东西的补充，只不过相对
不太常用
http/pprof 通过其HTTP服务器运行时提供性能测试数据，该数据的格式正是pprof可视化工
具需要的
net
rpc/jsonrpc 为rpc包实现了一个JSON-RPC ClientCodec和ServerCodec
os exec 可运行外部命令
user 通过名称和id进行用户账户检查
path filepath 实现了以与目标操作系统定义文件路径相兼容的方式处理文件名路径
regexp syntax 将正则表达式解析为语法树
runtime debug 包含当程序在运行时调试其自身的功能
pprof 以pprof可视化工具需要的格式写运行时性能测试数据
sync atomic 提供了低级的用于实现同步算法的原子级的内存机制
iotest 提供一系列测试目的的类型，实现了Reader和Writer标准接口
quick 实现了用于黑箱测试的实用函数
testing
script 帮助测试使用通道的代码
scanner 为UTF-8文本提供了一个扫描器和分词器
tabwriter 实现了一个写筛选器（tabwriter.Writer），它可将一个输入的tab分割的列
翻译为适当对齐的文本
template 数据驱动的模板引擎，用于生成类似HTML的文本输出格式
template/parse 为template构建解析树
unicode/utf16 实现了UTF-16序列的的编码和解码
text
unicode/utf8 实现了支持以UTF-8编码的文本的函数和常数
</pre>

####goroutine与channel实现并发
并行计算，两个goroutine进行并行的累加计算，都完成后打印
<pre>
package main

import "fmt"

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum //将计算结果发送到channel
}

func main() {
	values := []int{1, 3, 4, 5, 6, 7, 7, 8}
	resultChan := make(chan int, 2)  //定义2个goroutine
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <- resultChan, <-resultChan  //接收结果
	fmt.Println("result:", sum1 + sum2)
}
output ==>
36
</pre>
上面的是用2个切片完成，接下来分成3个切片完成求和的目的。
<pre>
package main

import "fmt"

func sum(value []int, resultChan chan int) {
	sum := 0
	for _, value := range value {
		sum += value
	}
	resultChan <- sum
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	resultChan := make(chan int, 3)
	go sum(arr[:len(arr)/3], resultChan)
	go sum(arr[len(arr)/3:len(arr)/3*2],resultChan)
	go sum(arr[len(arr)/3*2:],resultChan)
	sum1, sum2,sum3 := <-resultChan, <-resultChan,<-resultChan
	fmt.Println("result:",sum1+sum2+sum3)
}
output ==>
36
</pre>
带缓冲的channel：
<pre>
package main 
import "fmt"
func main() {
	c := make(chan int ,3)
	c <- 15
	c <- 34
	c <- 65
	close(c)
	fmt.Printf("%d\n",<-c)
	fmt.Printf("%d\n",<-c)
	fmt.Printf("%d\n",<-c)
}
output==>
15
34
65
</pre>
上面的虽然已经close了，但是我们依旧可以从中读出关闭前写入的3个值，下面的情况，则会出现错误提示：
<pre>
package main 
import "fmt"
func main() {
	c := make(chan int ,3)
	c <- 15
	c <- 34
	c <- 65
	close(c)
	c <- 1
	fmt.Printf("%d\n",<-c)
	fmt.Printf("%d\n",<-c)
	fmt.Printf("%d\n",<-c)
}
output==>
panic:send on closed channel
</pre>
第四次读取时，则会返回该channel类型的零值。向这类channel写入操作也会触发panic。<br>
close还可以协同多个Goroutines。比如下面这个例子，我们创建了100个Worker Goroutine，这些Goroutine在被创建出来后都阻塞在"<-start"上，直到我们在main goroutine中给出开工的信号："close(start)"，这些goroutines才开始真正的并发运行起来。
<pre>
package main 
import "fmt"
func worker(start chan bool,index int){
	<-start
	fmt.Println("This is worker:",index)
}
func main(){
	start := make(chan bool)
	for i:=1;i<100;i++{
		go worker(start,i)
	}
	close(start)
	select{}
}
</pre>
这里又引出一个话题：select{}的用法。
惯用方法：for/select
我们在使用select时很少只是对其进行一次evaluation，我们常常将其与for {}结合在一起使用，并选择适当时机从for{}中退出。
<pre>
for {
select {
case x := <- somechan:
// … 使用x进行一些操作
case y, ok := <- someOtherchan:
// … 使用y进行一些操作，
// 检查ok值判断someOtherchan是否已经关闭
case outputChan <- z:
// … z值被成功发送到Channel上时
default:
// … 上面case均无法通信时，执行此分支
}
}
</pre>

###go中有''
<pre>
package main 
import "fmt"
func main(){
	s := "hello"
	c := []byte(s)  //将字符串s转换为[]byte类型
	c[0] = 'c'  //这里如果是""，则是错误
	s2 := string(c) //再转换成string类型
	fmt.Printf("%s \n",s2)
}
</pre>
数组平均数
<pre>
package main

import (
	"fmt"
)
func getAvg(arr []int,size int) float32 {
	var sum int
	var avg float32
	for _,v :=range arr {
		sum += v
	}
	
	avg =float32(sum / size)
	return avg
}
func main(){
	arr :=[]int{4,5,7,7,6,4,54,13}
	var avg float32
	avg =getAvg(arr,len(arr))
	fmt.Println(avg)
}
</pre>
指针
<pre>
package main 
import "fmt"
func main(){
	var ip *int
	var a int = 20 
	ip =&a
	fmt.Printf("address of a : %x\n",&a)
	fmt.Printf("address of ip:%x\n",ip)
	fmt.Printf("ip is : %d \n",*ip)
}
</pre>

nil指针
<pre>
package main   
import "fmt"
func main(){  //会产生一个nil指针，即值为0
	var ptr *int
	fmt.Printf("ptr : %x \n",ptr)
}
</pre>
指针数组（将普通数组的值指针存到指定的指针数组，再由指针数组通过地址取出对应的值）
<pre>
package main
import "fmt"
const MAX int = 3
func main(){
	a :=[]int{10,100,200}
	var i int
	var ptr [MAX]*int    //指针数组的声明，数组元素的个数一定要写上,即[num]
	for i =0;i<MAX;i++ {
		ptr[i]=&a[i]
	}
	for i =0 ;i < MAX;i++ {
		fmt.Printf("value of a[%d] :%d\n",i, *ptr[i])
	}
}
</pre>
给指针数组赋值时出现的一个错误，关于range.
<pre>
package main

import (
	"fmt"
)
func main(){
	arr :=[4]int{3,5,67,8}
	var ptr [4]*int    
	/*这里错误，输出的都是8,最后一个值,用range循环赋值的注意点
	for k,v:=range arr {
		ptr[k] = &v
	}
	*/
	for i:=0;i<len(arr);i++{  //这一种则可以正确输出
		ptr[i] = &arr[i]
	}
	for i :=0;i<len(ptr);i++{
		fmt.Printf("%d\n",*ptr[i])
	}
}
</pre>
指针的指针
<pre>
package main 
import "fmt"
func main(){
	a := 4
	var ptr *int
	var pptr **int
	ptr= &a
	pptr = &ptr
	fmt.Printf("%d\n",a)
	fmt.Printf("%d\n",*ptr)
	fmt.Printf("%d\n",**pptr)
}
</pre>
函数指针
<pre>
package main
import "fmt"
func main(){
	var a int = 3
	var b int =6
	fmt.Printf("%d\n",a)	
	fmt.Printf("%d\n",b)
	swap(&a,&b)
	fmt.Printf("%d\n",a)
	fmt.Printf("%d\n",b)
} 
//Go语言允许您将指针传递给函数。要做到这一点，只需声明函数参数为指针类型
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x =*y
	*y = temp
}
</pre>
结构的例子
<pre>
package main

import (
	"fmt"
)
type book struct{
	title string
	author string
	subject string
	id int
}
func main(){
	var book1 book //声明一个book的结构
	book1.title= "bob"
	book1.author = "jason"
	book1.subject ="cs"
	book1.id=45
	fmt.Println(book1.title,book1.author,book1.subject,book1.id)
}
</pre>
结构作为函数参数，指针形式
<pre>
package main

import (
	"fmt"
)
type person struct {
	name string
	age int
}
func show(person1 *person){ //声明一个变量作为指向person结构的指针
	fmt.Println(person1.name)
	fmt.Println(person1.age)
}
func main(){
	var person1 person
	person1.name= "jason"
	person1.age =54
	show(&person1)  //&对一个内存地址进行访问
}
</pre>
字典map
<pre>
package main //map
import (
	"fmt"
)
func main(){
	var mapp map[string]string  //声明（必须）
	mapp=make(map[string]string) //创建（必须）,可以理解为对象的实例化
	mapp["name"]="jason"    //key-value
	mapp["age"] ="45"
	mapp["tool"]="knife"
	
	for _,v:=range mapp{
		fmt.Println(v,mapp[v])
	}
}
</pre>
###array|slice|struct|map|pointer|interface声明创建区别

+ array:
var arr [5]int  //声明<br>
arr :=[]int {...} //快捷创建

+ slice：
var sl make([]int)   //声明<br>
sl := []int{...}  //快速创建

+ struct:
<pre>
type str strcut{  //声明<br>
   ...
}
</pre>
str.name = ""  //创建元素
或者下面一种
<pre>
type rectangle struct {
	width,height int
}
r1 :=rectangle{12,2}
</pre>
+ map:
type mapp map[string]string  //声明<br>
mapp = make(map[string]string)//可以理解为实例化<br>
mapp["key"] ="value" //创建

+ 指针:
var ptr *int //声明一个变量指向int型变量的指针，存放内存地址<br>
 &ptr  //通过指针访问变量的在内存中的存放地址，显示变量值

+interface
 type inter interface {  //声明
	area() int 
}
带有接收者的函数，更加灵活
method
<pre>
package main
import (
	"fmt"
	"math"
)
type rectangle struct {
	width,height float64
}
type circle struct {
	radius float64
}
func (r rectangle) area() float64 {
	return r.width*r.height
}
func (c circle) area()float64{
	return c.radius*c.radius*math.Pi
}
func main(){
	r1 :=rectangle{12,2}
	r2 :=rectangle{6,4}
	c1 :=circle{2}
	c2 :=circle{5}
	fmt.Println(r1.area())
	fmt.Println(r2.area())
	fmt.Println(c1.area())
	fmt.Println(c2.area())
}
</pre>
结构的继承与带有接收者的函数的综合使用，
使函数优雅
<pre>
package main
import "fmt"
type human struct {
	name string
	age int
	phone string
}
type student struct {
	human
	school string
}
type employee struct {
	human
	company string
}
func (h *human) sayhi(){ //带有接收者的函数,完整的可以是(h *human) sayhi() int{}
	
	fmt.Printf("%s , %s\n",h.name,h.phone)
}
func main(){
	mark :=student{human{"mark",23,"3445343"},"MIT"}
	sam :=employee{human{"sam",45,"26564"},"Google"}
	mark.sayhi()
	sam.sayhi()
}
</pre>
并发(原则：不要通过共享来通信，而要通过通信来共享。)
<pre>
package main

import (
	"fmt"
	"runtime"
)
func say(s string){
	for i:=0;i<2;i++{
		runtime.Gosched() //让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}
func main(){
	go say("world") //开一个新的Gotuntines,后执行
	say("hello") //当前Goruntines
}
结果是:
hello
world
hello
</pre>
channel之声明及简单应用
<pre>
package main

import (
	"fmt"
)
func sum(a []int,c chan int){
	total :=0
	for _,v :=range a {
		total += v
	}
	c <- total //send total to c
}
func main(){
	a :=[]int{7,3,5,6,8}
	c :=make(chan int)
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
	x ,y := <-c,<-c
	fmt.Println(x,y,x+y)
}

</pre>
golang写一个http服务器
<pre>
package main
import (
	"log"
	"net/http"
	"fmt"
	"strings"
)
func sayhello(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v :=range r.Form {
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"<html><body><font color=red>Hello jason</font></body></html>") //输出到浏览器
}
func main(){
	http.HandleFunc("/",sayhello)
	err :=http.ListenAndServe(":9090",nil)
	if err !=nil {
		log.Fatal("ListenAndServe:",err)
	}
}
</pre>
<pre>
package main 

import (
	"time"
	"fmt"
)
func main(){
	a := 1
	go func (){
		a =2
	}()
a =3
fmt.Println("a is" ,a)
time.Sleep(2 * time.Second)
}
output==>
a is 3
</pre>
###其实golang中有'',它用于byte类型的变量的处理。比如:
<pre>
//将string类型转换成byte类型后再转换成string输出
package main 
func main(){
	ss :="hello world"
	bss :=[]byte(ss)
	bss[0]='H'    //这里必须单引号，单引号，单引号	
	fmt.Println(s,d,string(bss))
}
output ==>
Hello world
</pre>
类似的有rune,而rune更适合有中文的字符串。如下：
<pre>
package main
import "fmt"
func main(){
	str := "惦念"
	strrune :=[]rune(str)
	strtune[0] = '想'
	fmt.Println(string(strrune))
}
output==>
想念
</pre>
关于函数返回值：
<pre>
package main

import (
	"fmt"
)
func test()(int, string){ //这里int与string中间有','
	return 4,"星星"
}
func main(){
	str := "hello world"
	strrune := []rune(str)
	strrune[0] = 'H'
	_,xing := test()
	fmt.Println(string(strrune),xing)
}
</pre>
<pre>
package main

import (
	"fmt"
)
func main(){
	str := "大家好，我是谁"
	for _,v :=range str {
		fmt.Printf("%c",v) //%c单个字符输出
	}
}
output==>
大家好，我是谁	
</pre>
注意：range在循环时候操作的不是原数据，而是新复制的对象。
而for则不会。
<pre>
package main
import (
	"fmt"
)
func main(){
	var s string
	s = "leeeoo jason"
	a :=[3]int{0,2,5}
	println(s)
	/*for i,v :=range a {
		if i ==0 {
			a[1],a[2] =999,999
			fmt.Println(v)
		}
		fmt.Println(v)
	}   //输出0，2，5
	*/
	for i:=0;i<len(a);i++{
		if i == 1 {
			a[i] = 10
		}
		fmt.Println(a[i])
	}	//输出0,10,5
}
</pre>
将函数作为函数的返回值输出
<pre>
package main
import (
	"fmt"
)
func test() func(){
	return func(){
	x :=100
		fmt.Printf("x (%p) = %d \n",&x,x)
	}
}
func main(){
	f :=test()
	f()
}
output ==>
x (0xc0820022e0) = 100 
</pre>
数组小注意点：
<pre>
a:= [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."
</pre>
位拷贝拷贝的是地址，而值拷贝则拷贝的是内容。数组是值拷贝，值拷⻉贝⾏行为会造成性能问题，通常会建议使⽤用 slice，或数组指针。

####map字典
<pre>
package main
func main(){
	m := map[int]struct {
		name string
		age int
	}{
		1 : {"use1",10},
		5 : {"fgf",60},
	}
	println(m[5].age)
}
output==>
60
</pre>
####interface
小实例：
<pre>
package main 

import (
	"fmt"
)
type user struct {
	id int
	name string
}
func main (){
	u :=user{35,"jason"}
	var i interface{} = u
	fmt.Printf("%v\n",i.(user))
}
output ==>
{35 jason}
</pre>
%T --- >输出的是变量的类型，如int,string等等
<pre>
package main

import (
	"fmt"
)
func print (v interface{}){
	fmt.Printf("%T：%v\n",v,v)
}
func main(){
	print(1)
	print("Hello wolrd")
}
output ==>
int :1
string : hello wolrd
</pre>
如何通过interface接口修改数值：
接口转型其返回临时对象，只有使用指针才能修改其状态。
<pre>
package main
import (
	"fmt"
)
type user struct {
	name string
	age int
}
func main(){
	u :=user{"jack",12}
	var pi,vi interface {} = u,&u 
	//pi.(user).name = "jason1"     //Errot : cannot assign to pi.(user).name  | 只有指针才能修改
	vi.(*user).name = "jason2"
	fmt.Println(pi.(user))
	fmt.Println(vi.(*user))
}

output==>
{jack 12}     //不是指针形式修改不了
&{jason2 12}
</pre>
####channel
<pre>
package main
import "fmt"
func main() {
data := make(chan int) // 数据交换队列
exit := make(chan bool) // 退出通知
go func() {
for d := range data { // 从队列迭代接收数据，直到 close 。
fmt.Println(d)
}
fmt.Println("recv over.")
exit <- true // 发出退出通知。
}()
data <- 1 // 发送数据。
data <- 2
data <- 3
close(data) // 关闭队列。
fmt.Println("send over.")
<-exit // 等待退出通知。
}

----------
output ==>
1
2
3
send over.
recv over.
</pre>
简单http服务器
<pre>
package main
import (
	"fmt"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"yes")
}
func main(){
	http.HandleFunc("/test",handler)
	http.ListenAndServe(":8080",nil)
}
</pre>
####goto跳转
<pre>
package main

import (
	"fmt"
)
func main(){
	i:=0
	Loop:      //定义标签
	fmt.Printf("%d\n",i)	
	if i<10 {
		i++
		goto Loop //跳转回标签
	}
}
//实现的一个循环
</pre>
####强制转换成rune切片
<pre>
package main
import "fmt"
func main(){
	s := "是功夫f"
	fmt.Println(len([]rune(s)))  //强制转换
}
output==>
4   //不转换结果是10
</pre>
####goroutine 并发编程
<pre>
package main

import (
	"sync"
	"math"
)
func sum(id int){
	var x int64
	for i :=0;i<math.MaxUint32;i++{
		 x +=int64(i)
	}
	println(id,x)
}
func main(){
	wg :=new(sync.WaitGroup)
	wg.Add(2)
	for i :=0;i<2;i++{
		go func(id int){
			defer wg.Done()
			sum(id)
		}(i)
	}
	wg.Wait()
}
</pre>
####延迟调用defer
<pre>
package main
func main(){
	x ,y := 10,20
	defer func(i int){
		println("defer:",i,y)
	}(y)    //这里（y）相当于调用了该匿名函数并传入了一个值：y
	x += 10
	y += 100
	println("x=",x,"y=",y)
}
output ==>
x=10 y=120
defer 20 120
</pre>
###数组Array
- 数组是值类型，赋值和传参会复制整个数组，⽽而不是指针。

- 数组⻓长度必须是常量，且是类型的组成部分。[2]int 和 [3]int 是不同类型。

-  ⽀支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
-  数组是值类型，所以性能较切片和指针差，建议能够使用其他的则用它们代替数组。
####slice实例  (注意slice中范围数字的选取)
<pre>
package main

import (
	"fmt"
)
func main(){
	data :=[...]int{0,1,2,3,4,5,6,7,8,9}
	s :=data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
}
output ==>
[102,103]
</pre>
<pre>
package main
func main(){
	s :=[]int{1,34,5,6}  //slice
	d :=[4]int{4,5,6} //array
	f :=make([]int,3,4)  //make声明slice，注意这里是[]int,而不是[]int{}!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	f[2] = 3
	println(s[0])
	println(d[1])
	println(f[2])
}
output==>
1
5
3
</pre>
<pre>
package main
func main(){
	s :=[]int{1,2,3,4,5,6,7}
	s1 :=s[2:5]
	s2 :=append(s1,77)
	println(s1[0])  //3,4,5
	println(s2[0])	//3,4,5,77
}
</pre>
####map字典
<pre>
package main    //map字典
func main(){
	m :=make(map[string]int,1000)  //1000是提前申请的内存容量
	m =map[string]int{
		"a":1,
	}
	if v ,ok :=m["a"];ok{
		println(v)
	}
	m["b"] = 2
	println(m["b"])
	delete(m,"b")  //删除的语法！！！不是delete(m["b"])
	println(m["b"])
}
output==>
1
2
0
</pre>
<pre>
package main
func main(){
	type user struct {name string}
	m :=map[int]user{
		1 : {"sir"},   //这里逗号不可少
	}
	println(m[1].name)
}
output ==>
sir
</pre>
<pre>
package main    
type user struct {
	id int
	name string
}
func main(){     //换一种形式
	m :=map[user]int{
		user{1,"tom"}:100,
	}
	println(m[user{1,"tom"}])
}
output ==>
100
</pre>
####struct
<pre>
package main
func main(){
	type user struct {
		name string
	}
	type manage struct {
		user
		title string
	}
	m :=manage{
		user : user{"jason"},
		title : "administrator",
	}
	println(m.user.name)
	println(m.title)
}
output==>
jason
administrator
</pre>
<pre>
package main

import (
	"fmt"
	"encoding/json"
)
type person struct{
	username string
	age int
	friend []string   //切片
	addr string
}
func testjson(){
	p1 :=&person{
		"jason",
		23,
		[]string{"lisi","wangyu"},  //这里有[]string
		"hz",
	}
	p,err :=json.Marshal(p1)
	if err !=nil{
		fmt.Print(err.Error())
	}
	fmt.Print(p)
}
func main(){
	testjson()
}

</pre>
<pre>
package main

import (
	"fmt"
)
type user struct{
	id int
	name string
}
func (self user)test(){
	fmt.Println(self)
}
func main(){
	u :=user{2,"jason"}
	mvalue :=u.test
	u.id,u.name = 6 , "jack"
	u.test()
	mvalue()
}
</pre>
<pre>
package main

import (
	"fmt"
)
func Print(v interface{}){     //空接口
	fmt.Printf("%T:%v\n",v,v) //这里必须是%T,打印变量类型
}
func main(){
	Print(1)
}
output==>
int:1
</pre>
<pre>
package main

import (
	"fmt"
)
type Tester struct{
	s interface {     //声明一个接口
		String() string
	}
}
type User struct {
	id int
	name string
}
func (self *User) String() string {
	return fmt.Sprintf("user %d,%s",self.id,self.name)
}
func main(){
	t := Tester{&User{2,"jason"}}
	fmt.Println(t.s.String())   //调用接口
}
output==>
user 2,jason
</pre>
接口转型返回临时对象，只有使用指针才能修改其内容。
<pre>
package main

import (
	"fmt"
)
type user struct {
	id int
	name string
}
func main(){
	u := user{1,"tom"}
	var vi interface{} = &u         //这里,如果只是u，则报错
	vi.(*user).name = "jack"		//这里
	fmt.Println(vi.(*user))         //这里
}
output==>
&{1 jack}		
</pre>
<pre>
package main

import (
	"runtime"
	"sync"
)
func main(){
	wg :=new(sync.WaitGroup)
	wg.Add(1)  //添加或者减少等待goroutine的数量
	go func(){
		defer wg.Done()  //相当于Add(-1)
		defer println("A.defer")
		func() {
			defer println("B.defer")
			runtime.Goexit()
			println("8")
		}()
		println("A")
	}()
	wg.Wait() //执行阻塞，直到所有的WaitGroup数量变成0
}
/*WaitGroup的特点是Wait()可以用来阻塞直到队列中的所有任务
都完成时才解除阻塞，而不需要sleep一个固定的时间来等待．但是
其缺点是无法指定固定的goroutine数目．但是其缺点是无法指定
固定的goroutine数目．可能通过使用channel解决此问题。*/
output==>
B.defer
A.defer
</pre>
一个goroutine和主程序通信的例子。
<pre>
package main
func main(){
	channel :=make(chan string)   //d定义
	go func(){
		channel <- "hello"   //给channel赋值
	}()
	msg := <- channel    //将channel值赋给msg
	println(msg)
}
output ==>
hello
</pre>
给结构体起别名
<pre>
package main
type data struct{
	name string
	age int
}
type ddata data
func main(){
	data :=ddata{"nihao",34}
	println(data.name)
}
output==>
nihao 
</pre>
数组是值类型，所以不能通过传递参数在函数内部进行修改：
<pre>
package main
func modify(arr[5]int){
	arr[1] = 5
	println("arr[1] = ",arr[1])
}
func main(){
	td :=[5]int{1,2,3,4,5}
	modify(td)
	for _,val := range td{
		println(val)
	}
}
output ==>
arr[1] = 5
1
2
3
4
5
</pre>
var ch chan int<br>
var ch1 chan<- int  //ch1只能写<br>
var ch2 <-chan int  //ch2只能读<br>
channel是类型相关的，也就是一个channel只能传递一种类型。例如，上面的ch只能传递int。<br>
在go语言中，有4种引用类型：slice，map，channel，interface。<br>
channel消息传递
<pre>
package main

import (
	"time"
)
func producer(queue chan<- int){    //只能写
	for i:=0;i<10;i++{
		queue <- i
	}
}
func consumer(queue <- chan int){	//只能读
	for i:=0;i<10;i++{
		v :=<- queue
		println("rceive:",v)
	}
}
func main(){
	queue :=make(chan int , 1)
	go producer(queue)
	go consumer(queue)
	time.Sleep(1e9)   //等待所有并发完成
}
output==>
rceive: 0
rceive: 1
rceive: 2
rceive: 3
rceive: 4
rceive: 5
rceive: 6
rceive: 7
rceive: 8
rceive: 9
</pre>
多个channel
<pre>
package main

import (
	"time"
)
/*一个goroutine中处理多个channel的情况。我们不可能阻塞在
两个channel，这时就该select场了。与C语言中的select可以监
控多个fd一样，go语言中select可以等待多个channel。
*/
func main(){
	c1 :=make(chan string)
	c2 :=make(chan string)
	go func(){
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func(){
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	for i :=0;i<2;i++{
		select {
			case msg1 := <- c1:
			println("receive:",msg1)
			case msg2 := <- c2 :
			println("receive:",msg2)
		}
	}
}

output==>
receive: one
receive: two
</pre>
<pre>
package main
func main(){
	data :=make(chan int)     
	exit :=make(chan bool)  
	go func(){    
		for d:=range data{  //从队列接收数据，直到close
			println(d)
		}
		println("recv over")
		exit<-true	//发出退出通知
	}()
	data <- 1		//发送数据
	data <- 2
	data <- 3
	close(data)		//关闭队列
	println("send over")
	<-exit
}
output==>
1
2
3
send over
recv over
</pre>
<pre>
package main
func main(){
	data := make(chan int ,3)  // 缓冲区可以存储 3 个元素
	exit := make(chan bool)
	data <- 1					// 在缓冲区未满前，不会阻塞。(这里只能存入缓冲区允许的数量)
	data <- 2					//这里如果传入数据大于3个，则报错
	data <- 3
	go func(){
		for d :=range data{
			println(d)
		}
		exit <- true
	}()
	data <- 4
	data <- 5
	close(data)
	<- exit 
}
output ==>
1
2
3
4
5
</pre>
<pre>
package main
func main(){
	d1 :=make(chan int)
	d2 :=make(chan int,3)
	d2 <- 1
	d2 <- 2
	d2 <- 3
	//d2 <- 4 缓冲区数量超出允许的数值，报错
	println(len(d1),cap(d1))
	println(len(d2),cap(d2))
}
output ==>
0 0 
3 3 
</pre>
时间格式化
<pre>
package main
import (
	"time"
)
func main(){
	println(time.Now().Format("2006-01-02 15:04:05"))  //time.Now().Format("2006-01-02 15:04:05")格式化时间的标准格式,只能是这个时间，据说是go诞生之日, 记忆方法:6-1-2-3-4-5
}
output ==>
2016-02-21 15:53:09   //当前时间
</pre>
<pre>
package main
/*将通道作为变量形式进行赋值与输出,区别在于要有 "<-" */
import (
	"math/rand"
)
func test()chan int{
	c :=make(chan int)
	go func() {
		c <- rand.Int()     //c<- 写入值
	}()
	return c
}
func main(){
	t:= test()
	println(<-t)    //<-t 读取值
/*
等价于上面的
    t:=<-test()
	println(t)
*/
}
</pre>
初始化函数
<pre>
package main
/*
• 每个源⽂文件都可以定义一个或多个初始化函数。
• 编译器不保证多个初始化函数执行次序。
• 初始化函数在单一线程被调⽤用，仅执行一次。
• 初始化函数在包所有全局变量初始化后执⾏行。
• 在所有初始化函数结束后才执行 main.main。
• 无法调用初始化函数。
*/
import (
	"fmt"
	"time"
)
var now =time.Now()
func init(){
	fmt.Printf("%v",now)
}
func main(){
	
}
output ==>
2016-02-21 16:56:40.4286884 +0800 CST
</pre>
go语言的int转换成string有3种方法：

- 1、int32位，strconv.Itoa
- 2、大于32位，strconv.FormatInt()
- 3、万恶的fmt.Sprintf...好吧，这个我在php里是经常用来做格式化
<pre>
Golang简单写文件操作的四种方法
 /***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
 if checkFileIsExist(filename) {  //如果文件存在
  f, err1 = os.OpenFile(filename, os.O_APPEND, 0666)  //打开文件
  fmt.Println("文件存在");
 }else {
  f, err1 = os.Create(filename)  //创建文件
  fmt.Println("文件不存在");
 }
 check(err1)
 n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
 check(err1)
 fmt.Printf("写入 %d 个字节n", n);

/*****************************  第二种方式: 使用 ioutil.WriteFile 写入文件 ***********************************************/
 var d1 = []byte(wireteString);
 err2 := ioutil.WriteFile("./output2.txt", d1, 0666)  //写入文件(字节数组)
 check(err2)

/*****************************  第三种方式:  使用 File(Write,WriteString) 写入文件 ***********************************************/
 f, err3 := os.Create("./output3.txt")  //创建文件
 check(err3)
 defer f.Close()
 n2, err3 := f.Write(d1)  //写入文件(字节数组)
 check(err3)
 fmt.Printf("写入 %d 个字节n", n2)
 n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
 fmt.Printf("写入 %d 个字节n", n3)
 f.Sync()




 /***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 ***********************************************/
 w := bufio.NewWriter(f)  //创建新的 Writer 对象
 n4, err3 := w.WriteString("bufferedn")
 fmt.Printf("写入 %d 个字节n", n4)
 w.Flush()
 f.Close()
</pre>
sha1与md5加密
<pre>
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"crypto/md5"
)
func md5t(data string) string{
	t :=md5.New()
	io.WriteString(t,data) //使用 io.WriteString 写入文件|字符串
	return fmt.Sprintf("%x",t.Sum(nil)) //Sprintf(格式化)将int转成string
}
func sha1t(data string) string{
	t :=sha1.New()
	io.WriteString(t,data)
	return fmt.Sprintf("%x",t.Sum(nil))
}
func main(){
	var data string = "123"
	fmt.Printf("Md5:%s\n",md5t(data))
	fmt.Printf("Sha1:%s\n",sha1t(data))
}
output ==>
Md5:202cb962ac59075b964b07152d234b70
Sha1:40bd001563085fc35165329ea1ff5c5ecbdbbeef
</pre>
golang 中的json注意点
<br><b>
在使用json时，golang有一个大坑，非常非常大的坑，初学者很容易栽，那就是你定义的struct，如果某个字段需要被encoding到json的数据中，那这个字段必须是可导出的，也就是说，必须以大写字母开头！类似的情况也经常发生在”html/template”中，初学者务必小心。
</b>
###Golang读取文件
对于小文件，一次性读取：
<pre>
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
func readall(filepath string)([]byte,error){
	f,err :=os.Open(filepath)
	if err != nil{
		return nil,err
	}
	return ioutil.ReadAll(f)
}
func main(){
	d,_ :=readall("test.txt")
	fmt.Printf("%s",d)     
}
output ==>
test jason
</pre>
对于一般性文件，则是分块读取，可在速度和内存占用之间取得很好的平衡：
<pre>
package main

import (
	"io"
	"bufio"
	"os"
)
func processblock(line []byte){
	os.Stdout.Write(line)
}
func readblock(filepath string,bufsize int,hookfn func([]byte)) error{
	f ,err :=os.Open(filepath)
	if err != nil{
		return err
	}
	defer f.Close()
	buf :=make([]byte,bufsize) //规定一次读取多少字节
	bfrd :=bufio.NewReader(f)
	for {
		n ,err :=bfrd.Read(buf)
		hookfn(buf[:n]) // n是成功读取字节数
		if err != nil{ //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF { 
				return nil
			}
			return err
		}
	}
	return nil
}
func main(){
	readblock("test.txt",1000,processblock)
}
output ==>
test jason
</pre>
对于大文件来说，逐行读取很方便，性能可能慢一些，但是仅占用极少的内存空间：
<pre>
package main

import (
	"io"
	"bufio"
	"os"
)
func processline(line []byte){
	os.Stdout.Write(line)
}
func readline(filepath string,hookfn func([]byte)) error{
	f,err :=os.Open(filepath)
	if err != nil{
		return err
	}
	defer f.Close()
	bfrd :=bufio.NewReader(f)
	for{
		line,err :=bfrd.ReadBytes('\n')
		hookfn(line)
		if err != nil{
			if err == io.EOF{
				return nil
			}
			return err
		}
	}
	return nil
}
func main(){
	readline("test.txt",processline)
}
output ==>
test jason
</pre>
func Count(s, sep []byte) int 计算子字节切片 sep 在字节切片 s 中出现的非重叠实例的数量。
<pre>
package main

import (
	"bytes"
	"fmt"
)
func main(){
	s :=[]byte("banana")
	sep1 :=[]byte("an")
	fmt.Println(bytes.Count(s,sep1))
}
output ==>
2
</pre>
字节切片比较函数

- func Compare(a, b[]byte) int : 返回整数：1, 0 ,-1 
- func Equal(a, b []byte) bool: 返回true or false 
- func EqualFold(a, b []byte) bool : 忽略大小写：返回 true or false
<pre>
package main
import (
    "bytes"
    "fmt"
)
func main() {
    a := []byte("abc")
    b := []byte("ABC")
    s := []byte("GOLANG")
    t := []byte("golang")

    fmt.Println(bytes.Compare(a, b))
    fmt.Println(bytes.Equal(a, b))
    fmt.Println(bytes.EqualFold(s, t))
}
输出为：
1
false
true
</pre>
<pre>
package main
/*func Contains(b, subslice [] byte) bool 
检查字节切片 b ，是否包含子字节切片.检查是否存在某个字节 */
import (
	"bytes"
	"fmt"
)
func main(){
	b :=[]byte("golang")
	subslice1 :=[]byte("go")
	subslice2 :=[]byte("Go")
	fmt.Println(bytes.Contains(b,subslice1))
	fmt.Println(bytes.Contains(b,subslice2))
}
output==>
true 
false
</pre>
func Join(s [][]byte, sep []byte) []byte<br>
用字节切片 sep 吧 s中的每个字节切片连接成一个，并且返回
<pre>
package main
import (
    "bytes"
    "fmt"
)
func main() {
    // 字节切片 的每个元素，依旧是字节切片。
    s := [][]byte{
        []byte("你好"),
        []byte("世界"),  //这里的逗号，必不可少
    }
    sep := []byte(",")
    fmt.Println(string(bytes.Join(s, sep)))

    var a = []int{1,
        2,
        3,
        5,  //这里的逗号，也必不可少
    }
    fmt.Println(a)

    var b = []int{1, 2, 3, 4, 5}  //这里最后一个元素不需要逗号
    fmt.Println(b)
}
output==>
你好,世界
[1 2 3 5]
[1 2 3 4 5]
</pre>
<pre>
package main
type vertex struct {
	x int
	y int
}
func main(){
	p:=vertex{1,2}
	q :=&p.x
	println(*q)
}
output==>
1
</pre>
<pre>
package main
	myvar := 1 //error只能在函数内部使用简短的变量声明
	func main() {  
}
</pre>
<pre>
package main
	func main() {  
	    one := 0
	    one := 1 //error无法使用精简的赋值语句对变量重新赋值
	}
//下面的正确
	package main
	func main() {  
	    one := 0
	    one, two := 1,2
	    one,two = two,one
	}
</pre>
####除非特别指定，否则无法使用 nil 对变量赋值
nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型，所以会报错。
<pre>
package main
func main() {  
    var x = nil //error
    _ = x
}
//下面的正确
package main
func main() {  
    var x interface{} = nil
    _ = x
}
</pre>
####Map是定长的
创建 Map 的时候可以指定 Map 的长度，但是在运行时是无法使用 cap() 功能重新指定 Map 的大小，Map 是定长的。
<pre>
package main
func main() {  
    m := make(map[string]int,99)
    cap(m) //error
}
</pre>
####Go语言中，传递的数组不是内存地址，而是原数组的拷贝
所以是无法通过传递数组的方法去修改原地址的数据的。
如果需要修改原数组的数据，需要使用数组指针（array pointer）。
<pre>
package main
import "fmt"
func main() {  
    x := [3]int{1,2,3}
    func(arr *[3]int) {
        (*arr)[0] = 7
        fmt.Println(arr) //prints &[7 2 3]
    }(&x)
    fmt.Println(x) //prints [7 2 3]
}
</pre>
####试图访问不存在的 Map 键值
并不能在所有情况下都能通过判断 map 的记录值是不是 nil 判断记录是否存在。在 Go 语言中，对于“零值”是 nil 的数据类型可以这样判断，但是其他的数据类型不可以。简而言之，这种做法并不可靠（例如布尔变量的“零值”是 false）。最可靠的做法是检查 map 记录的第二返回值。
<pre>
//错误代码
package main

import "fmt"

func main() {  
    x := map[string]string{"one":"a","two":"","three":"c"}

    if v := x["two"]; v == "" { //incorrect
        fmt.Println("no entry")
    }
}
//修正代码
package main

import "fmt"

func main() {  
    x := map[string]string{"one":"a","two":"","three":"c"}

    if _,ok := x["two"]; !ok {
        fmt.Println("no entry")
    }
}
</pre>
####String不可变
对于 String 中单个字符的操作会导致编译失败。String 是带有一些附加属性的只读的字节片（Byte Slices）。所以如果想要对 String 操作的话，应当使用字节片操作，而不是将它转换为 String 类型。
<pre>
//错误代码
package main

import "fmt"

func main() {  
    x := "text"
    x[0] = 'T'

    fmt.Println(x)
}
//修改代码
package main

import "fmt"

func main() {  
    x := "text"
    xbytes := []byte(x)
    xbytes[0] = 'T'

    fmt.Println(string(xbytes)) //prints Text
}
</pre>
####String 与下标
和其他语言不同，String 的下表返回值是 Byte 类型的值，而不是字符类型。
<pre>
package main

import "fmt"

func main() {  
    x := "text"
    fmt.Println(x[0]) //print 116
    fmt.Printf("%T",x[0]) //prints uint8
}
//如果需要在 UTF8 类型的 String 中取出指定字符，那么需要用到 unicode/utf8 与实验性的 utf8string 包。utf8string 包包含 AT() 方法，可以取出字符，也可以将 String 转换为 Rune SLice。
</pre>
####String并不一定是UTF8格式
String 类型不一定是 UTF8 格式，String 中也可以包含自定义的文字/字节。只有需要将字符串显示出来的时候才需要用 UTF8 格式，其他情况下可以随便用转义来表示任意字符。<br>
可以使用 unicode/utf8 包中体重的 ValidString() 方法判断是否是 UTF8 类型的文本。
<pre>
package main

import (  
    "fmt"
    "unicode/utf8"
)

func main() {  
    data1 := "ABC"
    fmt.Println(utf8.ValidString(data1)) //prints: true

    data2 := "A\xfeC"
    fmt.Println(utf8.ValidString(data2)) //prints: false
}
</pre>
####log.Fatal 与 log.Panic 在后台悄悄做了一些事情
日志库提供了不同级别的日志记录，如果使用 Fatal 和 Panic 级别的日志，那么记录完这条日志后，应用程序便会退出而不会继续执行。
<pre>
package main

import "log"

func main() {  
    log.Fatalln("Fatal Level: log entry") //app exits here
    log.Println("Normal Level: log entry")
}
</pre>
go http
<pre>
package main

import (
	"io"
	"log"
	"net/http"
)
func main(){
	mux := http.NewServeMux()
	mux.Handle("/",&myHandler{})
	err :=http.ListenAndServe(":8089",mux)
	if err != nil {
		log.Fatal(err)
	}
}
type myHandler struct{}
func (*myHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	io.WriteString(w, r.URL.String())
}
/*package main

import (
	"io"
	"log"
	"net/http"
)
func main(){
	http.HandleFunc("/",sayhello)
	err :=http.ListenAndServe(":8088",nil)
	if err !=nil{
		log.Fatal(err)
	}
}
func sayhello(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"hello world,this is version 1.")
}
*/
</pre>
os包
<pre>
package main

import (
	"os"
	"fmt"
)
func main(){
	var goos string = os.Getenv("GOOS") //操作系统类型
	fmt.Printf("the operating system is :%s \n",goos)
	path := os.Getenv("PATH") //环境变量
	fmt.Printf("Path is %s\n",path)
}
</pre>
<pre>
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "reflect"
    "time"
)

func main() {
    dir, _ := os.Getwd()
    fmt.Println("dir:", dir)
    err := os.Chdir("d:/project/test2")
    dir, _ = os.Getwd()
    fmt.Println("dir:", dir)

    //参数不区分大小写
    //不存在环境变量就返回空字符串 len(path) = 0
    path := os.Getenv("gopath")
    fmt.Println(path)

    //返回有效group id
    egid := os.Getegid()
    fmt.Println("egid:", egid)

    //返回有效UID
    euid := os.Geteuid()
    fmt.Println("euid:", euid)

    gid := os.Getgid()
    fmt.Println("gid:", gid)

    uid := os.Getuid()
    fmt.Println("uid:", uid)

    //err:getgroups: not supported by windows
    g, err := os.Getgroups()
    fmt.Println(g, "error", err)

    pagesize := os.Getpagesize()
    fmt.Println("pagesize:", pagesize)

    ppid := os.Getppid()
    fmt.Println("ppid", ppid)

    //filemode, err := os.Stat("main.go")
    //不存在文件返回GetFileAttributesEx test2: The system cannot find the file specified.
    filemode, err := os.Stat("main.go")
    if err == nil {
        fmt.Println("Filename:", filemode.Name())
        fmt.Println("Filesize:", filemode.Size())
        fmt.Println("Filemode:", filemode.Mode())
        fmt.Println("Modtime:", filemode.ModTime())
        fmt.Println("IS_DIR", filemode.IsDir())
        fmt.Println("SYS", filemode.Sys())
    } else {
        fmt.Println("os.Stat error", err)
    }

    //Chmod is not supported under windows.
    //在windows变化是这样子的 -rw-rw-rw- => -r--r--r--
    err = os.Chmod("main.go", 7777)
    fmt.Println("chmod:", err)
    filemode, err = os.Stat("main.go")
    fmt.Println("Filemode:", filemode.Mode())

    //access time modification time
    err = os.Chtimes("main.go", time.Now(), time.Now())
    fmt.Println("Chtime error:", err)

    //获取全部的环境变量
    data := os.Environ()
    for _, val := range data {
        fmt.Println(val)
    }
    fmt.Println("---------end---environ----------------------")

    mapping := func(s string) string {
        m := map[string]string{"xx": "sssssssssssss",
            "yy": "ttttttttttttttt"}
        return m[s]
    }
    datas := "hello $xx blog address $yy"
    //这个函数感觉还蛮有用处
    expandStr := os.Expand(datas, mapping)
    fmt.Println(expandStr)
    datas = "GOBIN PATH $gopaTh" //不区分大小写
    fmt.Println(os.ExpandEnv(datas))

    hostname, err := os.Hostname()
    fmt.Println("hostname:", hostname)

    _, err = os.Open("WWWW.XX")
    if err != nil {
        fmt.Println(os.IsNotExist(err))
        fmt.Println(err)
    }

    f, err := os.Open("WWWW.XX")
    if err != nil && !os.IsExist(err) {
        fmt.Println(f, "not exist")
    }

    //windows 下两个都是true
    fmt.Println(os.IsPathSeparator('/'))
    fmt.Println(os.IsPathSeparator('\\'))
    fmt.Println(os.IsPathSeparator('.'))

    //判断返回的error 是否是因为权限的问题
    //func IsPermission(err error) bool

    // not supported by windows
    err = os.Link("main.go", "newmain.go")
    if err != nil {
        fmt.Println(err)
    }

    var pathSep string
    if os.IsPathSeparator('\\') {
        pathSep = "\\"
    } else {
        pathSep = "/"
    }
    fmt.Println("PathSeparator:", pathSep)
    //MkdirAll 创建的是所有下级目录，如果没有就创建他
    //Mkdir 创建目录，如果是多级目录遇到还未创建的就会报错
    err = os.Mkdir(dir+pathSep+"md"+pathSep+"md"+pathSep+"md"+pathSep+"md"+pathSep+"md", os.ModePerm)
    if err != nil {
        fmt.Println(os.IsExist(err), err)
    }

    err = os.RemoveAll(dir + "md\\md\\md\\md\\md")
    fmt.Println("removall", err)

    //rename 实际上通过movefile来实现的
    err = os.Rename("main.go", "main1.go")

    f1, _ := os.Stat("main.go")
    f2, _ := os.Stat("main1.go")
    if os.SameFile(f1, f2) {
        fmt.Println("the sanme")
    } else {
        fmt.Println("not same")
    }

    //os.Setenv 这个函数是设置环境变量的很简单
    evn := os.Getenv("WD_PATH")
    fmt.Println("WD_PATH:", evn)
    err = os.Setenv("WD_PATH", "D:/project")
    if err != nil {
        fmt.Println(err)
    }

    tmp, _ := ioutil.TempDir(dir, "tmp")
    fmt.Println(tmp)
    tmp = os.TempDir()
    fmt.Println(tmp)

    cf, err := os.Create("golang.go")
    defer cf.Close()
    fmt.Println(err)
    fmt.Println(reflect.ValueOf(f).Type())

    of, err := os.OpenFile("golang.goss", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
    defer of.Close()
    fmt.Println("os.OpenFile:", err)

    oof, err := os.Open("golang.goss")
    defer oof.Close()
    fmt.Println("os.Open file:", oof.Fd())
    fmt.Println("os.Open err:", err)
    oof.Close()

    r, w, err := os.Pipe()
    w.Write([]byte("1111"))
    var buf = make([]byte, 4)
    r.Read(buf)
    fmt.Println(buf)
    w.Write([]byte("2222"))
    r.Read(buf) // 如果没有调用w.Write(),r.Read()就会阻塞
    fmt.Println("ssss--", buf)

    b := make([]byte, 100)
    ff, _ := os.Open("main.go")
    n, _ := ff.Read(b)
    fmt.Println(n)
    fmt.Println(string(b[:n]))

    //第二个参数，是指，从第几位开始读取
    n, _ = ff.ReadAt(b, 20)
    fmt.Println(n)
    fmt.Println(string(b[:n]))

    //获取文件夹下文件的列表
    dirs, err := os.Open("md")
    if err != nil {
        fmt.Println(err)
    }
    defer dirs.Close()
    //参数小于或等去0，表示读取所有的文件
    //另外一个只读取文件名的函数
    //fs, err := dirs.Readdirname(0)
    fs, err := dirs.Readdir(-1)
    if err == nil {
        for _, file := range fs {
            fmt.Println(file.Name())
        }
    } else {
        fmt.Println("Readdir:", err)
    }

    //func (f *File) WriteString(s string) (ret int, err error)
    //写入字符串函数原型，哪个个函数比较快呢？？

    //p, _ := os.FindProcess(628)
    //fmt.Println(p)
    //p.Kill()
    attr := &os.ProcAttr{
        Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
    }
    //参数也可以这么写 `c:\windows\system32\notepad.EXE`  用的是反单引号
    p, err := os.StartProcess("c:\\windows\\system32\\notepad.EXE", []string{"c:\\windows\\system32\\notepad.EXE", "d:/1.txt"}, attr)
    p.Release()
    time.Sleep(1000000000)
    p.Signal(os.Kill)
    os.Exit(10)
}
</pre>
net/http包
<pre>
package main
//http客户端 
import (
	"io/ioutil"
	"net/http"
	"fmt"
)
func main(){
	response,err :=http.Get("http://baidu.com")
	if err !=nil {
		println(err)
	}
	defer response.Body.Close() //使用完成后关闭
    body,_ := ioutil.ReadAll(response.Body) 
    fmt.Println(string(body)) 
}
</pre>
###并发与并行
单核只能叫并发(伪并行)，多核才是并行。<br>
单核多线程并发有时间片的概念，你新开的 goroutine 里执行的代码可能在一个时间片里就执行完了。这里<br>
c <- true<br>
执行后，在同一个时间片里仍然继续执行了接下来的fmt.Println("go end").
之后调度器才切换到另一个goroutine，所以每次都打印了 go end。<br>
如果调度器让一个时间片缩短到只进行一个原子操作，那样才接近真并行，但太影响效率了，这中间只能有一个取舍。
###阻塞等待所有goroutines都完成
<b>官方推荐方案</b>
<pre>
package main

import (
	"time"
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup //定义一个同步等待的组
func main(){
	maxProcs :=runtime.NumCPU() //获取CPU个数
	runtime.GOMAXPROCS(maxProcs)//限制同时运行的goroutines数量
	for i :=0;i <10;i++{
		wg.Add(1) //为同步等待组增加一个成员
		go Printer(i) //并发一个goroutine
	}
	wg.Wait() //阻塞等待所有组内成员都执行完毕退栈
	fmt.Println("We Done")
}
//定义一个函数用于并发
func Printer(a int)(){
	time.Sleep(2000 * time.Millisecond)
	fmt.Printf("i am %d\n",a)
	defer wg.Done()
}
output==>
i am 0
i am 5
i am 1
i am 9
i am 8
i am 2
i am 7
i am 3
i am 4
i am 6
We Done
</pre>
另一种是利用channel的阻塞机制
<pre>
package main
//利用channel的阻塞机制
import (
	"time"
	"fmt"
	"runtime"
)
var num =14  //定义共并发多少数量
var cnum chan int
func main(){
	maxprocs :=runtime.NumCPU()
	runtime.GOMAXPROCS(maxprocs)//限制同时运行的goroutines数量
	cnum =make(chan int,num)//make一个chan,缓存为num
	for i :=0;i<num;i++{
		go Printer(i)
	}
	//利用channel的阻塞，一直从信道取数据，直到取得跟并发数一样的个数的数据，
	//则视为所有goroutines完成
	for i :=0;i<num;i++{
		<- cnum
	}
	fmt.Println("Done!!!")
}

func Printer(a int){
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("I am %d\n",a)
	cnum <- 1 //goroutine结束时传送一个标示给信道。
}
output==>
I am 0
I am 1
I am 7
I am 13
I am 12
I am 11
I am 10
I am 9
I am 8
I am 4
I am 6
I am 5
I am 3
I am 2
Done!!!
</pre>
将上面的稍微改装一下,输出结果将不是无序的，我猜想改装之后不是并发执行，而是单核并行，没有利用到多核带来的高性能。
<pre>
package main
//利用channel的阻塞机制
import(
	"fmt"
	"runtime"
)
var num =14  //定义共并发多少数量
var cnum chan int
func main(){
	maxprocs :=runtime.NumCPU()
	runtime.GOMAXPROCS(maxprocs)//限制同时运行的goroutines数量
	cnum =make(chan int,num)//make一个chan,缓存为num
	//利用channel的阻塞，一直从信道取数据，直到取得跟并发数一样的个数的数据，
	//则视为所有goroutines完成
	for i :=0;i<num;i++{
		go Printer(i)
		<- cnum
	}
	fmt.Println("Done!!!")
}

func Printer(a int){
	//time.Sleep(200 * time.Millisecond)
	fmt.Printf("I am %d\n",a)
	cnum <- 1 //goroutine结束时传送一个标示给信道。
}
output==>
I am 0
I am 1
I am 2
I am 3
I am 4
I am 5
I am 6
I am 7
I am 8
I am 9
I am 10
I am 11
I am 12
I am 13
Done!!!
</pre>
默认的，信道的存消息和取消息都是阻塞的 (叫做无缓冲的信道，不过缓冲这个概念稍后了解，先说阻塞的问题)。<br>
<b>阻塞</b>:也就是说, 无缓冲的信道在取消息和存消息的时候都会挂起当前的goroutine，除非另一端已经准备好。
<b>无缓冲的信道</b>:其实，无缓冲的信道永远不会存储数据，只负责数据的流通，为什么这么讲呢？<br>
从无缓冲信道取数据，必须要有数据流进来才可以，否则当前线阻塞<br>
数据流入无缓冲信道, 如果没有其他goroutine来拿走这个数据，那么当前线阻塞<br>
所以，你可以测试下，无论如何，我们测试到的无缓冲信道的大小都是0 (len(channel))<br>
如果信道正有数据在流动，我们还要加入数据，或者信道干涩，我们一直向无数据流入的空信道取数据呢？ 就会引起死锁。你会看到下面的报错内容：<br>
fatal error: all goroutines are asleep - deadlock!<br>
何谓死锁? 操作系统有讲过的，所有的线程或进程都在等待资源的释放。如上的程序中, 只有一个goroutine, 所以当你向里面加数据或者存数据的话，都会锁死信道， 并且阻塞当前 goroutine, 也就是所有的goroutine(其实就main线一个)都在等待信道的开放(没人拿走数据信道是不会开放的)，也就是死锁咯。<br>
一个死锁的案例
<pre>
package main
import (
	"fmt"
)
func main(){
	ch :=make(chan int)
	ch <- 1 //1流入信道，堵塞当前线, 没人取走数据信道不会打开,造成死锁
	fmt.Println("this line code won`t run")
}
</pre>
但是，是否果真 所有不成对向信道存取数据的情况都是死锁?如下是个反例:
<pre>
package main
func main() {
    c := make(chan int)

    go func() {
       c <- 1
    }()
}
成功执行
</pre>
解释：程序正常退出了，很简单，并不是我们那个总结不起作用了，还是因为一个让人很囧的原因，main又没等待其它goroutine，自己先跑完了， 所以没有数据流入c信道，一共执行了一个goroutine, 并且没有发生阻塞，所以没有死锁错误。
<b>那么死锁的解决办法呢？</b><br>

- 把没取走的数据取走，没放入的数据放入， 因为无缓冲信道不能承载数据，那么就赶紧拿走
<pre>
c, quit := make(chan int), make(chan int)

  go func() {
    c <- 1
    quit <- 0
  }()

  <- c // 取走c的数据！
  <-quit
 </pre>

- 另一个解决办法是缓冲信道, 即设置c有一个数据的缓冲大小
<pre>
c := make(chan int, 1)
</pre>
这样的话，c可以缓存一个数据。也就是说，放入一个数据，c并不会挂起当前线, 再放一个才会挂起当前线直到第一个数据被其他goroutine取走, 也就是只阻塞在容量一定的时候，不达容量不阻塞。这十分类似Python中的队列Queue.
####缓冲信道
缓存信道用英文来讲更为达意: buffered channel.<br>
缓冲这个词意思是，缓冲信道不仅可以流通数据，还可以缓存数据。它是有容量的，存入一个数据的话 , 可以先放在信道里，不必阻塞当前线而等待该数据取走。而前面说的无缓冲信道只能流通数据，而不能存储。<br>
当缓冲信道达到满的状态的时候，就会表现出阻塞了，因为这时再也不能承载更多的数据了，「你们必须把 数据拿走，才可以流入数据。<br>
在声明一个信道的时候，我们给make以第二个参数来指明它的容量(默认为0，即无缓冲):
<pre>
var ch chan int = make(chan int, 2) // 写入2个元素都不会阻塞当前goroutine, 存储个数达到2的时候会阻塞
</pre>
例子：
<pre>
package main
func main(){
	ch :=make(chan int,3)
	//向缓存信道流入3个元素，程序正常执行
	ch <- 1
	ch <- 3
	ch <- 4
}
</pre>
其实，缓冲信道是先进先出的，我们可以把缓冲信道看作为一个线程安全的队列：
<pre>
package main
func main(){
	var ch chan int = make(chan int,2)
	ch <- 3
	ch <- 2
	println(<- ch)  //缓冲信道先进先出
	println(<- ch)
	
}
output==>
3
2
</pre>
####信道数据读取和信道关闭
信道数据的读取在使用range时候需要注意。一般情况下，range不等到信道关闭是不会结束读取。也就是如果 缓冲信道干涸了，那么range就会阻塞当前goroutine, 所以死锁。
<pre>
package main
funcm main(){
	ch :=make(chan int ,3)
	ch <- 1
	ch <- 4
	ch <- 3
	for v :=range ch {
		println(v)
	}
}
报错，range不等到信道关闭不会结束读取，导致出现死锁
</pre>
那么怎么改进呢？方法之一是：显式地关闭信道。<br>
被关闭的信道会禁止数据流入, 是只读的。我们仍然可以从关闭的信道中取出数据，但是不能再写入数据了。
<pre>
package main
func main(){
	ch :=make(chan int,3)
	ch <- 1
	ch <- 4
	ch <- 3
	close(ch)
	for v :=range ch {
		println(v)
	}
}
output==>
1
4
3
</pre>
方法之二：读到信道为空的时候就结束读取：
<pre>
package main
func main(){
	var ch chan int =make(chan int ,3)
	ch <- 1
	ch <- 4
	ch <- 3
	for i := 0;i<3;i++{
		println(<-ch)
		if(len(ch) <= 0){
			break
		}
	}
}
output==>
1
4
3
</pre>
小结：

- 无缓冲的信道是一批数据一个一个的"流进流出".
- 缓冲信道则是一个一个存储，然后一起流出去.



####使用channel的close发送广播
<pre>
package main

import (
	"fmt"
	"math/rand"
	"time"
)
func waiter(i int,block,done chan struct{} ){
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Println(i,"waiting...")
	<- block
	fmt.Println(i,"done!")
	done <- struct{}{}
}
func main(){
	block,done :=make(chan struct{}),make(chan struct{})
	for i:=0;i<4;i++{
		go waiter(i,block,done)
	}
	time.Sleep(5 * time.Second)
	close(block)
	for i:=0;i<4;i++{
		<- done
	}
}
output==>
3 waiting...
2 waiting...
1 waiting...
0 waiting...
3 done!
2 done!
1 done!
0 done!
</pre>

###goroutine背后的系统知识

1. 操作系统与运行库
2. 并发与并行 (Concurrency and Parallelism)
3. 线程的调度
4. 并发编程框架
5. goroutine

1. 操作系统与运行库

对于普通的电脑用户来说，能理解应用程序是运行在操作系统之上就足够了，可对于开发者，我们还需要了解我们写的程序是如何在操作系统之上运行起来的，操作系统如何为应用程序提供服务，这样我们才能分清楚哪些服务是操作系统提供的，而哪些服务是由我们所使用的语言的运行库提供的。

除了内存管理、文件管理、进程管理、外设管理等等内部模块以外，操作系统还提供了许多外部接口供应用程序使用，这些接口就是所谓的“系统调用”。从DOS时代开始，系统调用就是通过软中断的形式来提供，也就是著名的INT 21，程序把需要调用的功能编号放入AH寄存器，把参数放入其他指定的寄存器，然后调用INT 21，中断返回后，程序从指定的寄存器(通常是AL)里取得返回值。这样的做法一直到奔腾2也就是P6出来之前都没有变，譬如windows通过INT 2E提供系统调用，Linux则是INT 80，只不过后来的寄存器比以前大一些，而且可能再多一层跳转表查询。后来，Intel和AMD分别提供了效率更高的SYSENTER/SYSEXIT和SYSCALL/SYSRET指令来代替之前的中断方式，略过了耗时的特权级别检查以及寄存器压栈出栈的操作，直接完成从RING 3代码段到RING 0代码段的转换。

系统调用都提供什么功能呢？用操作系统的名字加上对应的中断编号到谷歌上一查就可以得到完整的列表 (Windows, Linux)，这个列表就是操作系统和应用程序之间沟通的协议，如果需要超出此协议的功能，我们就只能在自己的代码里去实现，譬如，对于内存管理，操作系统只提供进程级别的内存段的管理，譬如Windows的virtualmemory系列，或是Linux的brk，操作系统不会去在乎应用程序如何为新建对象分配内存，或是如何做垃圾回收，这些都需要应用程序自己去实现。如果超出此协议的功能无法自己实现，那我们就说该操作系统不支持该功能，举个例子，Linux在2.6之前是不支持多线程的，无论如何在程序里模拟，我们都无法做出多个可以同时运行的并符合POSIX 1003.1c语义标准的调度单元。

可是，我们写程序并不需要去调用中断或是SYSCALL指令，这是因为操作系统提供了一层封装，在Windows上，它是NTDLL.DLL，也就是常说的Native API，我们不但不需要去直接调用INT 2E或SYSCALL，准确的说，我们不能直接去调用INT 2E或SYSCALL，因为Windows并没有公开其调用规范，直接使用INT 2E或SYSCALL无法保证未来的兼容性。在Linux上则没有这个问题，系统调用的列表都是公开的，而且Linus非常看重兼容性，不会去做任何更改，glibc里甚至专门提供了syscall(2)来方便用户直接用编号调用，不过，为了解决glibc和内核之间不同版本兼容性带来的麻烦，以及为了提高某些调用的效率(譬如__NR_ gettimeofday)，Linux上还是对部分系统调用做了一层封装，就是VDSO (早期叫linux-gate.so)。

可是，我们写程序也很少直接调用NTDLL或者VDSO，而是通过更上一层的封装，这一层处理了参数准备和返回值格式转换、以及出错处理和错误代码转换，这就是我们所使用语言的运行库，对于C语言，Linux上是glibc，Windows上是kernel32(或调用msvcrt)，对于其他语言，譬如Java，则是JRE，这些“其他语言”的运行库通常最终还是调用glibc或kernel32。

“运行库”这个词其实不止包括用于和编译后的目标执行程序进行链接的库文件，也包括了脚本语言或字节码解释型语言的运行环境，譬如Python，C#的CLR，Java的JRE。

对系统调用的封装只是运行库的很小一部分功能，运行库通常还提供了诸如字符串处理、数学计算、常用数据结构容器等等不需要操作系统支持的功能，同时，运行库也会对操作系统支持的功能提供更易用更高级的封装，譬如带缓存和格式的IO、线程池。

所以，在我们说“某某语言新增了某某功能”的时候，通常是这么几种可能：
1. 支持新的语义或语法，从而便于我们描述和解决问题。譬如Java的泛型、Annotation、lambda表达式。
2. 提供了新的工具或类库，减少了我们开发的代码量。譬如Python 2.7的argparse
3. 对系统调用有了更良好更全面的封装，使我们可以做到以前在这个语言环境里做不到或很难做到的事情。譬如Java NIO

但任何一门语言，包括其运行库和运行环境，都不可能创造出操作系统不支持的功能，Go语言也是这样，不管它的特性描述看起来多么炫丽，那必然都是其他语言也可以做到的，只不过Go提供了更方便更清晰的语义和支持，提高了开发的效率。

2. 并发与并行 (Concurrency and Parallelism)

并发是指程序的逻辑结构。非并发的程序就是一根竹竿捅到底，只有一个逻辑控制流，也就是顺序执行的(Sequential)程序，在任何时刻，程序只会处在这个逻辑控制流的某个位置。而如果某个程序有多个独立的逻辑控制流，也就是可以同时处理(deal)多件事情，我们就说这个程序是并发的。这里的“同时”，并不一定要是真正在时钟的某一时刻(那是运行状态而不是逻辑结构)，而是指：如果把这些逻辑控制流画成时序流程图，它们在时间线上是可以重叠的。

并行是指程序的运行状态。如果一个程序在某一时刻被多个CPU流水线同时进行处理，那么我们就说这个程序是以并行的形式在运行。（严格意义上讲，我们不能说某程序是“并行”的，因为“并行”不是描述程序本身，而是描述程序的运行状态，但这篇小文里就不那么咬文嚼字，以下说到“并行”的时候，就是指代“以并行的形式运行”）显然，并行一定是需要硬件支持的。

而且不难理解：

1. 并发是并行的必要条件，如果一个程序本身就不是并发的，也就是只有一个逻辑控制流，那么我们不可能让其被并行处理。

2. 并发不是并行的充分条件，一个并发的程序，如果只被一个CPU流水线进行处理(通过分时)，那么它就不是并行的。

3. 并发只是更符合现实问题本质的表达方式，并发的最初目的是简化代码逻辑，而不是使程序运行的更快；

这几段略微抽象，我们可以用一个最简单的例子来把这些概念实例化：用C语言写一个最简单的HelloWorld，它就是非并发的，如果我们建立多个线程，每个线程里打印一个HelloWorld，那么这个程序就是并发的，如果这个程序运行在老式的单核CPU上，那么这个并发程序还不是并行的，如果我们用多核多CPU且支持多任务的操作系统来运行它，那么这个并发程序就是并行的。

还有一个略微复杂的例子，更能说明并发不一定可以并行，而且并发不是为了效率，就是Go语言例子里计算素数的sieve.go。我们从小到大针对每一个因子启动一个代码片段，如果当前验证的数能被当前因子除尽，则该数不是素数，如果不能，则把该数发送给下一个因子的代码片段，直到最后一个因子也无法除尽，则该数为素数，我们再启动一个它的代码片段，用于验证更大的数字。这是符合我们计算素数的逻辑的，而且每个因子的代码处理片段都是相同的，所以程序非常的简洁，但它无法被并行，因为每个片段都依赖于前一个片段的处理结果和输出。

并发可以通过以下方式做到：

1. 显式地定义并触发多个代码片段，也就是逻辑控制流，由应用程序或操作系统对它们进行调度。它们可以是独立无关的，也可以是相互依赖需要交互的，譬如上面提到的素数计算，其实它也是个经典的生产者和消费者的问题：两个逻辑控制流A和B，A产生输出，当有了输出后，B取得A的输出进行处理。线程只是实现并发的其中一个手段，除此之外，运行库或是应用程序本身也有多种手段来实现并发，这是下节的主要内容。

2. 隐式地放置多个代码片段，在系统事件发生时触发执行相应的代码片段，也就是事件驱动的方式，譬如某个端口或管道接收到了数据(多路IO的情况下)，再譬如进程接收到了某个信号(signal)。

并行可以在四个层面上做到：

1. 多台机器。自然我们就有了多个CPU流水线，譬如Hadoop集群里的MapReduce任务。

2. 多CPU。不管是真的多颗CPU还是多核还是超线程，总之我们有了多个CPU流水线。

3. 单CPU核里的ILP(Instruction-level parallelism)，指令级并行。通过复杂的制造工艺和对指令的解析以及分支预测和乱序执行，现在的CPU可以在单个时钟周期内执行多条指令，从而，即使是非并发的程序，也可能是以并行的形式执行。

4. 单指令多数据(Single instruction, multiple data. SIMD)，为了多媒体数据的处理，现在的CPU的指令集支持单条指令对多条数据进行操作。

其中，1牵涉到分布式处理，包括数据的分布和任务的同步等等，而且是基于网络的。3和4通常是编译器和CPU的开发人员需要考虑的。这里我们说的并行主要针对第2种：单台机器内的多核CPU并行。

关于并发与并行的问题，Go语言的作者Rob Pike专门就此写过一个幻灯片：http://talks.golang.org/2012/waza.slide

在CMU那本著名的《Computer Systems: A Programmer’s Perspective》里的这张图也非常直观清晰：


3. 线程的调度

上一节主要说的是并发和并行的概念，而线程是最直观的并发的实现，这一节我们主要说操作系统如何让多个线程并发的执行，当然在多CPU的时候，也就是并行的执行。我们不讨论进程，进程的意义是“隔离的执行环境”，而不是“单独的执行序列”。

我们首先需要理解IA-32 CPU的指令控制方式，这样才能理解如何在多个指令序列(也就是逻辑控制流)之间进行切换。CPU通过CS:EIP寄存器的值确定下一条指令的位置，但是CPU并不允许直接使用MOV指令来更改EIP的值，必须通过JMP系列指令、CALL/RET指令、或INT中断指令来实现代码的跳转；在指令序列间切换的时候，除了更改EIP之外，我们还要保证代码可能会使用到的各个寄存器的值，尤其是栈指针SS:ESP，以及EFLAGS标志位等，都能够恢复到目标指令序列上次执行到这个位置时候的状态。

线程是操作系统对外提供的服务，应用程序可以通过系统调用让操作系统启动线程，并负责随后的线程调度和切换。我们先考虑单颗单核CPU，操作系统内核与应用程序其实是也是在共享同一个CPU，当EIP在应用程序代码段的时候，内核并没有控制权，内核并不是一个进程或线程，内核只是以实模式运行的，代码段权限为RING 0的内存中的程序，只有当产生中断或是应用程序呼叫系统调用的时候，控制权才转移到内核，在内核里，所有代码都在同一个地址空间，为了给不同的线程提供服务，内核会为每一个线程建立一个内核堆栈，这是线程切换的关键。通常，内核会在时钟中断里或系统调用返回前(考虑到性能，通常是在不频繁发生的系统调用返回前)，对整个系统的线程进行调度，计算当前线程的剩余时间片，如果需要切换，就在“可运行”的线程队列里计算优先级，选出目标线程后，则保存当前线程的运行环境，并恢复目标线程的运行环境，其中最重要的，就是切换堆栈指针ESP，然后再把EIP指向目标线程上次被移出CPU时的指令。Linux内核在实现线程切换时，耍了个花枪，它并不是直接JMP，而是先把ESP切换为目标线程的内核栈，把目标线程的代码地址压栈，然后JMP到__switch_to()，相当于伪造了一个CALL __switch_to()指令，然后，在__switch_to()的最后使用RET指令返回，这样就把栈里的目标线程的代码地址放入了EIP，接下来CPU就开始执行目标线程的代码了，其实也就是上次停在switch_to这个宏展开的地方。

这里需要补充几点：(1) 虽然IA-32提供了TSS (Task State Segment)，试图简化操作系统进行线程调度的流程，但由于其效率低下，而且并不是通用标准，不利于移植，所以主流操作系统都没有去利用TSS。更严格的说，其实还是用了TSS，因为只有通过TSS才能把堆栈切换到内核堆栈指针SS0:ESP0，但除此之外的TSS的功能就完全没有被使用了。(2) 线程从用户态进入内核的时候，相关的寄存器以及用户态代码段的EIP已经保存了一次，所以，在上面所说的内核态线程切换时，需要保存和恢复的内容并不多。(3) 以上描述的都是抢占式(preemptively)的调度方式，内核以及其中的硬件驱动也会在等待外部资源可用的时候主动调用schedule()，用户态的代码也可以通过sched_yield()系统调用主动发起调度，让出CPU。

现在我们一台普通的PC或服务里通常都有多颗CPU (physical package)，每颗CPU又有多个核 (processor core)，每个核又可以支持超线程 (two logical processors for each core)，也就是逻辑处理器。每个逻辑处理器都有自己的一套完整的寄存器，其中包括了CS:EIP和SS:ESP，从而，以操作系统和应用的角度来看，每个逻辑处理器都是一个单独的流水线。在多处理器的情况下，线程切换的原理和流程其实和单处理器时是基本一致的，内核代码只有一份，当某个CPU上发生时钟中断或是系统调用时，该CPU的CS:EIP和控制权又回到了内核，内核根据调度策略的结果进行线程切换。但在这个时候，如果我们的程序用线程实现了并发，那么操作系统可以使我们的程序在多个CPU上实现并行。

这里也需要补充两点：(1) 多核的场景里，各个核之间并不是完全对等的，譬如在同一个核上的两个超线程是共享L1/L2缓存的；在有NUMA支持的场景里，每个核访问内存不同区域的延迟是不一样的；所以，多核场景里的线程调度又引入了“调度域”(scheduling domains)的概念，但这不影响我们理解线程切换机制。(2) 多核的场景下，中断发给哪个CPU？软中断(包括除以0，缺页异常，INT指令)自然是在触发该中断的CPU上产生，而硬中断则又分两种情况，一种是每个CPU自己产生的中断，譬如时钟，这是每个CPU处理自己的，还有一种是外部中断，譬如IO，可以通过APIC来指定其送给哪个CPU；因为调度程序只能控制当前的CPU，所以，如果IO中断没有进行均匀的分配的话，那么和IO相关的线程就只能在某些CPU上运行，导致CPU负载不均，进而影响整个系统的效率。

4. 并发编程框架

以上大概介绍了一个用多线程来实现并发的程序是如何被操作系统调度以及并行执行(在有多个逻辑处理器时)，同时大家也可以看到，代码片段或者说逻辑控制流的调度和切换其实并不神秘，理论上，我们也可以不依赖操作系统和其提供的线程，在自己程序的代码段里定义多个片段，然后在我们自己程序里对其进行调度和切换。

为了描述方便，我们接下来把“代码片段”称为“任务”。

和内核的实现类似，只是我们不需要考虑中断和系统调用，那么，我们的程序本质上就是一个循环，这个循环本身就是调度程序schedule()，我们需要维护一个任务的列表，根据我们定义的策略，先进先出或是有优先级等等，每次从列表里挑选出一个任务，然后恢复各个寄存器的值，并且JMP到该任务上次被暂停的地方，所有这些需要保存的信息都可以作为该任务的属性，存放在任务列表里。

看起来很简单啊，可是我们还需要解决几个问题：

(1) 我们运行在用户态，是没有中断或系统调用这样的机制来打断代码执行的，那么，一旦我们的schedule()代码把控制权交给了任务的代码，我们下次的调度在什么时候发生？答案是，不会发生，只有靠任务主动调用schedule()，我们才有机会进行调度，所以，这里的任务不能像线程一样依赖内核调度从而毫无顾忌的执行，我们的任务里一定要显式的调用schedule()，这就是所谓的协作式(cooperative)调度。(虽然我们可以通过注册信号处理函数来模拟内核里的时钟中断并取得控制权，可问题在于，信号处理函数是由内核调用的，在其结束的时候，内核重新获得控制权，随后返回用户态并继续沿着信号发生时被中断的代码路径执行，从而我们无法在信号处理函数内进行任务切换)

(2) 堆栈。和内核调度线程的原理一样，我们也需要为每个任务单独分配堆栈，并且把其堆栈信息保存在任务属性里，在任务切换时也保存或恢复当前的SS:ESP。任务堆栈的空间可以是在当前线程的堆栈上分配，也可以是在堆上分配，但通常是在堆上分配比较好：几乎没有大小或任务总数的限制、堆栈大小可以动态扩展(gcc有split stack，但太复杂了)、便于把任务切换到其他线程。

到这里，我们大概知道了如何构造一个并发的编程框架，可如何让任务可以并行的在多个逻辑处理器上执行呢？只有内核才有调度CPU的权限，所以，我们还是必须通过系统调用创建线程，才可以实现并行。在多线程处理多任务的时候，我们还需要考虑几个问题：

(1) 如果某个任务发起了一个系统调用，譬如长时间等待IO，那当前线程就被内核放入了等待调度的队列，岂不是让其他任务都没有机会执行？

在单线程的情况下，我们只有一个解决办法，就是使用非阻塞的IO系统调用，并让出CPU，然后在schedule()里统一进行轮询，有数据时切换回该fd对应的任务；效率略低的做法是不进行统一轮询，让各个任务在轮到自己执行时再次用非阻塞方式进行IO，直到有数据可用。

如果我们采用多线程来构造我们整个的程序，那么我们可以封装系统调用的接口，当某个任务进入系统调用时，我们就把当前线程留给它(暂时)独享，并开启新的线程来处理其他任务。

(2) 任务同步。譬如我们上节提到的生产者和消费者的例子，如何让消费者在数据还没有被生产出来的时候进入等待，并且在数据可用时触发消费者继续执行呢？

在单线程的情况下，我们可以定义一个结构，其中有变量用于存放交互数据本身，以及数据的当前可用状态，以及负责读写此数据的两个任务的编号。然后我们的并发编程框架再提供read和write方法供任务调用，在read方法里，我们循环检查数据是否可用，如果数据还不可用，我们就调用schedule()让出CPU进入等待；在write方法里，我们往结构里写入数据，更改数据可用状态，然后返回；在schedule()里，我们检查数据可用状态，如果可用，则激活需要读取此数据的任务，该任务继续循环检测数据是否可用，发现可用，读取，更改状态为不可用，返回。代码的简单逻辑如下：
<pre>
struct chan {
    bool ready,
    int data
};

int read (struct chan *c) {
    while (1) {
        if (c->ready) {
            c->ready = false;
            return c->data;
        } else {
            schedule();
        }
    }
}

void write (struct chan *c, int i) {
    while (1) {
        if (c->ready) {
            schedule(); 
        } else {
            c->data = i;
            c->ready = true;
            schedule(); // optional
            return;
        }
    }
}
</pre>
很显然，如果是多线程的话，我们需要通过线程库或系统调用提供的同步机制来保护对这个结构体内数据的访问。

以上就是最简化的一个并发框架的设计考虑，在我们实际开发工作中遇到的并发框架可能由于语言和运行库的不同而有所不同，在功能和易用性上也可能各有取舍，但底层的原理都是殊途同归。

譬如，glic里的getcontext/setcontext/swapcontext系列库函数可以方便的用来保存和恢复任务执行状态；Windows提供了Fiber系列的SDK API；这二者都不是系统调用，getcontext和setcontext的man page虽然是在section 2，但那只是SVR4时的历史遗留问题，其实现代码是在glibc而不是kernel；CreateFiber是在kernel32里提供的，NTDLL里并没有对应的NtCreateFiber。

在其他语言里，我们所谓的“任务”更多时候被称为“协程”，也就是Coroutine。譬如C++里最常用的是Boost.Coroutine；Java因为有一层字节码解释，比较麻烦，但也有支持协程的JVM补丁，或是动态修改字节码以支持协程的项目；PHP和Python的generator和yield其实已经是协程的支持，在此之上可以封装出更通用的协程接口和调度；另外还有原生支持协程的Erlang等，笔者不懂，就不说了，具体可参见Wikipedia的页面：http://en.wikipedia.org/wiki/Coroutine

由于保存和恢复任务执行状态需要访问CPU寄存器，所以相关的运行库也都会列出所支持的CPU列表。

从操作系统层面提供协程以及其并行调度的，好像只有OS X和iOS的Grand Central Dispatch，其大部分功能也是在运行库里实现的。

5. goroutine

Go语言通过goroutine提供了目前为止所有(我所了解的)语言里对于并发编程的最清晰最直接的支持，Go语言的文档里对其特性也描述的非常全面甚至超过了，在这里，基于我们上面的系统知识介绍，列举一下goroutine的特性，算是小结：

(1) goroutine是Go语言运行库的功能，不是操作系统提供的功能，goroutine不是用线程实现的。具体可参见Go语言源码里的pkg/runtime/proc.c

(2) goroutine就是一段代码，一个函数入口，以及在堆上为其分配的一个堆栈。所以它非常廉价，我们可以很轻松的创建上万个goroutine，但它们并不是被操作系统所调度执行

(3) 除了被系统调用阻塞的线程外，Go运行库最多会启动$GOMAXPROCS个线程来运行goroutine

(4) goroutine是协作式调度的，如果goroutine会执行很长时间，而且不是通过等待读取或写入channel的数据来同步的话，就需要主动调用Gosched()来让出CPU

(5) 和所有其他并发框架里的协程一样，goroutine里所谓“无锁”的优点只在单线程下有效，如果$GOMAXPROCS > 1并且协程间需要通信，Go运行库会负责加锁保护数据，这也是为什么sieve.go这样的例子在多CPU多线程时反而更慢的原因

(6) Web等服务端程序要处理的请求从本质上来讲是并行处理的问题，每个请求基本独立，互不依赖，几乎没有数据交互，这不是一个并发编程的模型，而并发编程框架只是解决了其语义表述的复杂性，并不是从根本上提高处理的效率，也许是并发连接和并发编程的英文都是concurrent吧，很容易产生“并发编程框架和coroutine可以高效处理大量并发连接”的误解。

(7) Go语言运行库封装了异步IO，所以可以写出貌似并发数很多的服务端，可即使我们通过调整$GOMAXPROCS来充分利用多核CPU并行处理，其效率也不如我们利用IO事件驱动设计的、按照事务类型划分好合适比例的线程池。在响应时间上，协作式调度是硬伤。

(8) goroutine最大的价值是其实现了并发协程和实际并行执行的线程的映射以及动态扩展，随着其运行库的不断发展和完善，其性能一定会越来越好，尤其是在CPU核数越来越多的未来，终有一天我们会为了代码的简洁和可维护性而放弃那一点点性能的差别。

####runtime.Gosched的作用分析
runtime.Gosched()用于让出CPU时间片，就像是跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A停止执行，B继续执行。
<pre>
package main

import (
	"runtime"
)
func say(s string){
	for i:=0;i<2;i++{
		runtime.Gosched()
		println(s)
		
	}
}
func main(){
	go say("world")
	say("hello")
}
output==>
hello
world
hello
解析：
注意结果：
1、先输出了hello,后输出了world.
2、hello输出了2个，world输出了1个（因为第2个hello输出完，主线程就退出了，第2个world没机会了）
把代码中的runtime.Gosched()注释掉，执行结果是：
hello
hello
因为say("hello")这句占用了时间，等它执行完，线程也结束了，say("world")就没有机会了。
这里同时可以看出，go中的goroutins并不是同时在运行。事实上，如果没有在代码中通过
runtime.GOMAXPROCS(n) 其中n是整数，
指定使用多核的话，goroutins都是在一个线程里的，它们之间通过不停的让出时间片轮流运行，达到类似同时运行的效果。
</pre>
当然，要牢记一句话：<br>
当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutines转移到另一个系统线程上去，以使这些goroutines不阻塞.<br>
并发关乎结构，并行关乎执行;并发是指同时处理很多事情,在程序设计阶段；并行是指同时能完成很多事情，在执行阶段。
并发提供了一种方式让我们能够设计一种方案将问题(非必须的)并行的解决；同时执行(通常是相关的)计算任务的编程技术。

####select
注意到 select 的代码形式和 switch 非常相似， 不过 select 的 case 里的操作语句只能是【IO 操作】.<br>
此示例里面 select 会一直等待等到某个 case 语句完成， 也就是等到成功从 ch1 或者 ch2 中读到数据。 则 select 语句结束。<br>
【使用 select 实现 timeout 机制】
<pre>
package main
import (
	"time"
)
func main(){
	timeout := make(chan bool,1)
	var ch  chan int=make(chan int)
	go func(){
		ch <- 5 
		time.Sleep(1e9)
		timeout <- true
	}()
	select {
		case <- ch :
		println("收到数据")
		case <-timeout:
		println("timeout!")
	}
}
</pre>

###Golang中的单引号
golang里单引号只能有一个字符，如果输出会返回这个字符的ascii码 ，如果想输出为字符需要用string()函数转换一下。


###Golang之cond锁定期唤醒锁
cond的主要作用就是获取锁之后，wait()方法会等待一个通知，来进行下一步锁释放等操作，以此控制锁合适释放，释放频率。
<pre>
package main
import (
	"time"
	"fmt"
	"sync"
)
var locker = new(sync.Mutex)
var cond =sync.NewCond(locker)
func test(x int){
	cond.L.Lock()             //获取锁
	cond.Wait()               //等待通知，暂时阻塞
	fmt.Println(x)
 	time.Sleep(time.Second)
  	cond.L.Unlock()          //释放锁
}
func main(){
	for i :=0;i<40;i++{
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	fmt.Println("broadcast")
	cond.Signal()            //下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Signal()            //1秒后下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 2)
	cond.Broadcast()         //2秒之后下发广播给所有等待的goroutine
	time.Sleep(time.Second * 10)
}
output ==>
start all
broadcast
3
0
20
1
2
6
4
5
8
7
9
10
</pre>
带有404页面的Go-web
<pre>
package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

//设置路由器
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello gerryyang, version 2!\n")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9095", mux)
}

</pre>
####Golang之iinterface
<pre>
package main
func main(){
	var s interface{} =[]string{"left","right"}
	for _,x :=range s.([]string){
		println(x)
	}
}
output==>
left
right
</pre>

###Golang字符串
<pre>
package main
import (
	"strings"
	"fmt"
)
func main(){
	fmt.Println("是否包含：",strings.Contains("test","f")) //false
	fmt.Println("出现字符的次数：",strings.Count("teast","t")) //1
	fmt.Println("判断字符串首部是不是某个字符：",strings.HasPrefix("test","te")) //true
	fmt.Println("判断字符串尾部是不是某个字符：",strings.HasSuffix("test,","rf")) //false
	fmt.Println("查询字符串位置:",strings.Index("teast","e")) //1
	fmt.Println(strings.LastIndex("go gopher", "go")) // 3
	fmt.Println("字符串数组连接：",strings.Join([]string{"a","b"},"-")) //a-b
	fmt.Println("重复一个字符串：",strings.Repeat("a",3)) //aaa
	fmt.Println("字符串替换，若指定起始位置为0，则全部替换：",strings.Replace("foo","o","3",-1))//f33
	fmt.Println("字符串替换 指定起始位置1:",strings.Replace("foo","o","22",1)) //f22o
	fmt.Println("字符串切割：",strings.Split("a-b-c-d-e","-")) //[a b c d e]
	fmt.Println("小写转换：",strings.ToLower("TEST")) //test
	fmt.Println("大写转换：",strings.ToUpper("up")) //UP
	fmt.Println("长度：",len("helo")) //4
	fmt.Println("标取字符串中的字符：","hello"[1]) //101
	fmt.Println(strings.ContainsRune("我是中国", '我')) // true  注意第二个参数，用的是字符（单引号）
	fmt.Println(strings.EqualFold("Go", "go")) // true 大小写忽略
	fmt.Println(strings.IndexRune("我是中国人", '中')) // 在存在返回 6
	fmt.Println("Fields are: %q", strings.Fields(" foo bar baz ")) //["foo" "bar" "baz"] 返回一个列表

}
</pre>
读文件
<pre>
package main

import (
	"os"
)
func main(){
	userfile :="yes.txt"
	fin,err :=os.Open(userfile)
	defer fin.Close()
	if err != nil {
		println(userfile,err)
		return
	}
	buf :=make([]byte,1024)
	for {
		n ,_:=fin.Read(buf)
		if 0 == n {break}
		os.Stdout.Write(buf[:n])
	}
}
</pre>
重写文件
<pre>
package main

import (
	"os"
)
func main(){
	userfile :="yes.txt"
	fout,err :=os.Create(userfile)
	defer fout.Close()
	if err != nil {
		println(userfile,err)
		return 
	}
	for i := 0;i<10;i++{
		fout.WriteString("Just a test!\r\n")
	}
}
</pre>

####Golang原子操作
原子操作即是进行过程中不能被中断的操作。针对某个值的原子操作在被进行的过程中，CPU绝不会再去进行其他的针对该值的操作。为了实现这样的严谨性，原子操作仅会由一个独立的CPU指令代表和完成。

- GO语言提供的原子操作都是非入侵式的，由标准库sync/atomic中的众多函数代表 
- 类型包括int32,int64,uint32,uint64,uintptr,unsafe.Pointer，共六个。
- 这些函数提供的原子操作共有五种：增或减，比较并交换，载入，存储和交换
####增或减Add
<pre>
package main
import (
    "fmt"
    "sync/atomic"
)

func main(){
    var i32 int32
    fmt.Println("=====old i32 value=====")
    fmt.Println(i32)
    //第一个参数值必须是一个指针类型的值,因为该函数需要获得被操作值在内存中的存放位置,以便施加特殊的CPU指令
    //结束时会返回原子操作后的新值
    newI32 := atomic.AddInt32(&i32,3)
    fmt.Println("=====new i32 value=====")
    fmt.Println(i32)
    fmt.Println(newI32)

    var i64 int64
    fmt.Println("=====old i64 value=====")
    fmt.Println(i64)
    newI64 := atomic.AddInt64(&i64,-3)
    fmt.Println("=====new i64 value=====")
    fmt.Println(i64)
    fmt.Println(newI64)

}
</pre>
####比较并交换CAS
Compare And Swap 简称CAS，在sync/atomic包种，这类原子操作由名称以‘CompareAndSwap’为前缀的若干个函数代表。
<pre>
package main

import (
    "fmt"
    "sync/atomic"
)

var value int32

func main()  {
    fmt.Println("======old value=======")
    fmt.Println(value)
    fmt.Println("======CAS value=======")
    addValue(3)
    fmt.Println(value)


}

//不断地尝试原子地更新value的值,直到操作成功为止
func addValue(delta int32){
    //在被操作值被频繁变更的情况下,CAS操作并不那么容易成功
    //so 不得不利用for循环以进行多次尝试
    for {
        v := value
        if atomic.CompareAndSwapInt32(&value, v, (v + delta)){
            //在函数的结果值为true时,退出循环
            break
        }
        //操作失败的缘由总会是value的旧值已不与v的值相等了.
        //CAS操作虽然不会让某个Goroutine阻塞在某条语句上,但是仍可能会使流产的执行暂时停一下,不过时间大都极其短暂.
    }
}
</pre>
####载入Load
上面的比较并交换案例总 v:= value为变量v赋值，但… 要注意，在进行读取value的操作的过程中,其他对此值的读写操作是可以被同时进行的,那么这个读操作很可能会读取到一个只被修改了一半的数据.
<pre>
package main

import (
    "fmt"
    "sync/atomic"
)

var value int32

func main()  {
    fmt.Println("======old value=======")
    fmt.Println(value)
    fmt.Println("======CAS value=======")
    addValue(3)
    fmt.Println(value)


}

//不断地尝试原子地更新value的值,直到操作成功为止
func addValue(delta int32){
    //在被操作值被频繁变更的情况下,CAS操作并不那么容易成功
    //so 不得不利用for循环以进行多次尝试
    for {
        //v := value
        //在进行读取value的操作的过程中,其他对此值的读写操作是可以被同时进行的,那么这个读操作很可能会读取到一个只被修改了一半的数据.
        //因此我们要使用载入
        v := atomic.LoadInt32(&value)
        if atomic.CompareAndSwapInt32(&value, v, (v + delta)){
            //在函数的结果值为true时,退出循环
            break
        }
        //操作失败的缘由总会是value的旧值已不与v的值相等了.
        //CAS操作虽然不会让某个Goroutine阻塞在某条语句上,但是仍可能会使流产的执行暂时停一下,不过时间大都极其短暂.
    }
}
</pre>
####存储Store
与读取操作相对应的是写入操作。 而sync/atomic包也提供了与原子的载入函数相对应的原子的值存储函数。 以Store为前缀
在原子地存储某个值的过程中，任何CPU都不会进行针对同一个值的读或写操作。原子的值存储操作总会成功，因为它并不会关心被操作值的旧值是什么和CAS操作有着明显的区别
<pre>
    fmt.Println("======Store value=======")
    atomic.StoreInt32(&value, 10)
    fmt.Println(value)
</pre>

####bufio.Reader
创建支持缓存读取的具有缺省长度缓冲区的Reader对象，Reader对象会从底层的io.Reader接口读取尽量多的数据进行缓存。
<pre>
package main

import (
	"fmt"
	"bufio"
	"bytes"
)
func main(){
	rb :=bytes.NewReader([]byte("a string to be read"))
	r :=bufio.NewReader(rb)
	var buf [128]byte
	n ,_:=r.Read(buf[:])
	fmt.Println(string(buf[:n]))
}
output==>
a string to be read
</pre>
####archive/tar读取tar文件
<pre>
package main

import (
    "os"
    "io"
    "archive/tar"
	"fmt"
)
func handleError(err error) {
	fmt.Println(err)
}
func main() {
    fr, err := os.Open("demo.tar")  // 打开tar包文件，返回*io.Reader
    handleError(err)    // handleError为错误处理函数，下同
    defer fr.Close()

    // 实例化新的tar.Reader
    tr := tar.NewReader(fr)

    for {
        hdr, err := tr.Next()   // 获取下一个文件，第一个文件也用此方法获取
        if err == io.EOF {
            break   // 已读到文件尾
        }
        handleError(err)

        // 通过创建文件获得*io.Writer
        fw, _ := os.Create("demo/" + hdr.Name)
        handleError(err)

        // 拷贝数据
        _, err = io.Copy(fw, tr)
        handleError(err)
    }
}
output ==>
<nil>
</pre>
tar创建新文件
<pre>
package main

import (
    "os"
    "io"
    "archive/tar"
)
func handleError(err error){
	println(err)
}
func main() {
    fw, err := os.Create("do.tar")    // 创建tar包文件，返回*io.Writer
    handleError(err)    // handleError为错误处理函数，下同
    defer fw.Close()

    // 实例化新的tar.Writer
    tw := tar.NewWriter(fw)
    defer tw.Close()

    // 获取要打包的文件的内容
    fr, err := os.Open("do.txt")
    handleError(err)
    defer fr.Close()

    // 获取文件信息
    fi, err := fr.Stat()
    handleError(err)

    // 创建tar.Header
    hdr := new(tar.Header)
    hdr.Name = fi.Name()
    hdr.Size = fi.Size()
    hdr.Mode = int64(fi.Mode())
    hdr.ModTime = fi.ModTime()

    // 写入数据头
    err = tw.WriteHeader(hdr)
    handleError(err)

    // 写入文件数据
    _, err = io.Copy(tw, fr)
    handleError(err)
}
//已经创建do.tar文件
</pre>
####bufio.NewReaderSize
创建的支持缓存读取的具有指定长度缓冲区的Reader对象，Reader对象会从底层的io.Reader接口读取尽量多的数据进行缓存。
<pre>
package main

import (
	"bufio"
	"bytes"
)
func main(){
	rb :=bytes.NewBuffer([]byte("this is a new string"))
	r :=bufio.NewReaderSize(rb,8132)
	var buf [128]byte
	n ,_:=r.Read(buf[:])
	println(string(buf[:n]))
}
output==>
this is a new string
</pre>
####bufio.NewWriter
创建支持缓存写的具有缺省长度缓冲区的Writer对象，Writer对象会将缓存的数据批量写入底层的io.Writer接口
<pre>
package main

import (
	"fmt"
	"bufio"
	"bytes"
)
func main(){
	wb :=bytes.NewBuffer(nil)
	w :=bufio.NewWriter(wb)
	w.Write([]byte("hello."))
	w.Write([]byte("world!"))
	fmt.Printf("%d:%s\n",len(wb.Bytes()),string(wb.Bytes()))
	w.Flush()
	fmt.Printf("%d:%s\n",len(wb.Bytes()),string(wb.Bytes()))
}
output==>
0:
12:hello.world!
</pre>
####bufio.NewReaderSize
创建支持缓存写的具有指定长度缓冲区的Writer对象，Writer对象会将缓存的数据批量写入底层的io.Writer接口
<pre>
package main
import (
    "bytes"
    "bufio"
    "fmt"
)

func main() {
    wb := bytes.NewBuffer(nil)
    w := bufio.NewWriterSize(wb, 8192)
    w.Write([]byte("hello,"))
    w.Write([]byte("world!"))
    fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
    w.Flush()
    fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
}
output==>
0: 12:hello,world! 
</pre>
###peek
读取指定字节数的数据，这些被读取的数据不会从缓冲区中清除。在下次读取之后，本次返回的字节切片会失效。如果Peek返回的字节数不足n字节，则会同时返回一个错误说明原因。如果n比缓冲区要大，则错误为ErrBufferFull。
<pre>
package main

import (
	"fmt"
	"bufio"
	"bytes"
)
func main(){
	rb :=bytes.NewBuffer([]byte("12345678"))
	r :=bufio.NewReader(rb)
	b1 ,_:=r.Peek(4)
	fmt.Println(string(b1))
	b2,_:=r.Peek(8)
	fmt.Println(string(b2))
}
output==>
1234
12345678
</pre>
读取数据存放到p中，返回已读取的字节数。
<pre>
package main
import (
	"fmt"
	"bufio"
	"bytes"
)
func main(){
	rb :=bytes.NewBuffer([]byte("1234567890"))
	r :=bufio.NewReader(rb)
	var buf [128]byte
	n,err :=r.Read(buf[:]) //n 读取的字节数,err 错误
	fmt.Printf("%d,%v\n",n,err)
	fmt.Println(string(buf[:n]))
}
output==>
10,<nil>
1234567890
</pre>
返回第一个字符
<pre>
package main

import (
	"fmt"
	"bufio"
	"bytes"
)
func main(){
	rb :=bytes.NewBuffer([]byte("987654321"))
	r :=bufio.NewReader(rb)
	b,err :=r.ReadByte()
	fmt.Printf("%c,%v\n",b,err)
}
output==>
9,<nil>
</pre>
读取到指定字符
<pre>
package main

import (
	"fmt"
	"bufio"
	"bytes"
)
func main(){
	rb :=bytes.NewReader([]byte("123$456"))
	r :=bufio.NewReader(rb)
	b,err :=r.ReadBytes('$')
	fmt.Printf("%s,%v\n",string(b),err)
}
output==>
123$
</pre>
###container
常用的容器工具包，目前有heap、list、ring三种数据结构.
####heap
任何实现了heap.Interface接口的对象都可以使用heap包提供的方法对堆进行操作(堆是一个完全二叉树)。通过对heap.Interface中的Less方法的不同实现，来实现最大堆和最小堆。通常堆的数据结构为一个一维数组。
####list
list包实现了双向链表的功能。<br>
遍历一个list的代码实例（其中l为*list对象）：
<pre>
    for e := l.Front(); e != nil; e = e.Next() {
        // do something with e.Value
    }
</pre>
####ring
ring包实现了环形双向链表的功能。
###AES加密
<pre>
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "os"
)

func main() {
    // 消息明文
    src := []byte("hello, world")
    // 密钥，长度可以为16、24、32字节
    key := "1234567890abcdef"
    // 初始向量，长度必须为16个字节(128bit)
    var iv = []byte("abcdef1234567890")[:aes.BlockSize]
    // 得到块，用于加密和解密
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
        os.Exit(1)
    }
    fmt.Printf("NewClipher(%d bytes)\n", len(key))

    // 加密，使用CFB模式(密文反馈模式)，其他模式参见crypto/cipher
    encrypter := cipher.NewCFBEncrypter(block, iv)

    encrypted := make([]byte, len(src))
    encrypter.XORKeyStream(encrypted, src)
    fmt.Printf("Encrypting %s : %v -> %v\n", src, []byte(src), encrypted)

    // 解密
    decrypter := cipher.NewCFBDecrypter(block, iv)

    decrypted := make([]byte, len(src))
    decrypter.XORKeyStream(decrypted, encrypted)
    fmt.Printf("Decrypting %v -> %v : %s\n", encrypted, decrypted, decrypted)
}
output==>
Encrypting hello, world : [104 101 108 108 111 44 32 119 111 114 108 100] -> [235 32 43 140 87 212 167 232 74 65 110 69]
Decrypting [235 32 43 140 87 212 167 232 74 65 110 69] -> [104 101 108 108 111 44 32 119 111 114 108 100] : hello, world
</pre>
###Errors
<pre>
package main
import (
	"errors"
	"fmt"
)
func main(){
	fmt.Println(errors.New("Err"))
}
output==>
Err
</pre>
###EscapeString
EscapeString用于将特殊字符转移为html实体,如把<转义成&lt;<br>
它只会转义下列五种字符: < > & ' " <br>
需要注意的是 UnescapeString(EscapeString(s)) == s 返回结果一定是true,但是反过来就不一定是true了.
<pre>
package main
import (
	"html"
	"fmt"
)
func main(){
	var s string = "<script>alert('xss');</script>"
	fmt.Println(html.EscapeString(s))
}
output==>
&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;
</pre>
UnescapeString
<pre>
package main

import (
	"html"
	"fmt"
)
func main(){
	var s string = "&lt;script&gt;alert(&#39;xss&#39;);&lt;/script&gt"
	fmt.Println(html.UnescapeString(s))
}
output==>
<script>alert('xss');</script>
</pre>
####image包
<pre>
package main  
import (  
    "fmt"  
    "image"  
    "image/color"  
    "image/jpeg"  
    "log"  
    "os"  
)  
  
const (  
    dx = 500  
    dy = 200  
)  
  
func main() {  
  
    file, err := os.Create("test.jpeg")  
    if err != nil {  
        log.Fatal(err)  
    }  
    defer file.Close()  
    alpha := image.NewAlpha(image.Rect(0, 0, dx, dy))  
    for x := 0; x < dx; x++ {  
        for y := 0; y < dy; y++ {  
            alpha.Set(x, y, color.Alpha{uint8(x % 256)})   //设定alpha图片的透明度  
        }  
    }  
  
    fmt.Println(alpha.At(400, 100))    //144 在指定位置的像素  
    fmt.Println(alpha.Bounds())        //(0,0)-(500,200)，图片边界  
    fmt.Println(alpha.Opaque())        //false，是否图片完全透明  
    fmt.Println(alpha.PixOffset(1, 1)) //501，指定点相对于第一个点的距离  
    fmt.Println(alpha.Stride)          //500，两个垂直像素之间的距离  
    jpeg.Encode(file, alpha, nil)      //将image信息写入文件中  
}  
</pre>
####suffixarray
功能说明： 根据正则表达式查找所有的索引，并返回匹配结果在数据中的位置（结果已排序）
<pre>
package main
import (
	"fmt"
	"index/suffixarray"
)
func main(){
	data := []byte("YeS")
	index :=suffixarray.New(data)
	fmt.Println(index.Bytes())
}
output==>
[89 101 83]
</pre>
####index.Lookup
功能说明：根据输入byte数组查找索引，并返回匹配结果在数据中的位置（结果未排序）
<pre>
package main

import (
	"fmt"
	"index/suffixarray"
)
func main(){
	data :=[]byte("helloyork")
	index:=suffixarray.New(data)
	str :=[]byte("y")
	res :=index.Lookup(str,1)
	fmt.Println(res)
}
output==>
[5]
</pre>
返回Index类型,功能说明： 创建一个索引.
<pre>
package main

import (
	"fmt"
	"index/suffixarray"
)
func main(){
	data :=[]byte("UpUp")
	index :=suffixarray.New(data)
	fmt.Println(index)
}
ouptut==>
&{[85 112 85 112] [2 0 3 1]}
</pre>
####copy
文件复制操作。<br>
功能说明： 向dst拷贝src的全部数据；读取src中数据直到EOF，故不会返回io.EOF.
可能的异常： io.ErrShortWrite:写入数据不等于读取数据.
<pre>
package main
import (
	"io"
	"os"
)
func main(){
	yes,_ :=os.Open("yes.txt")
	df,_ :=os.Create("copy.txt")
	wr,err :=io.Copy(yes,df) //关键的一步,wr指的是复制的内存大小(拷贝字节数)
	if err == nil{
		println("Copy success, total",wr,"bytes")
	}
}
</pre>
####io.EOF
io.EOF：当n>src字节数时，拷贝src全部数据并返回src字节数<br>
io.ErrShortWrite:写入数据不等于读取数据<br>
<pre>
package main 

import (
	"io"
	"os"
)
func main(){
	sf,_ :=os.Open("yes.txt")
	df,_ :=os.Create("df.txt")
	written,err :=io.Copy(df,sf) 
	if err == nil{
		println("success!,total",written,"bytes")
	}
	if err ==io.EOF{
		println("copy all total",written,"bytes")
	}
}
output==>
success!,total 140 bytes
</pre>
####limitreader
获得一个只能从r读到n比特数据的Reader
<pre>
package main

import (
    "io"
    "fmt"
    "os"
    "reflect"
)

func main() {
    reader, _ := os.Open("yes.txt")
    limitReader := io.LimitReader(reader, 20)
    fmt.Println(reflect.TypeOf(limitReader))
}
output==>
*io.LimitedReader
</pre>
####MultiReader
获得一个可以对参数中的多个Reader进行连续读取的reader,一次性读取多个文件。
<pre>
import (
    "io"
    "fmt"
    "os"
    "reflect"
)

func main() {
    reader1, _ := os.Open("src1.txt")
    reader2, _ := os.Open("src2.txt")
    multiReader := io.MultiReader(reader1, reader2)
    fmt.Println(reflect.TypeOf(multiReader))
}
</pre>
####同样的就有将一个数据同时写入多个文件中。MultiWriter
<pre>
package main

import (
    "io"
    "fmt"
    "os"
    "reflect"
)

func main() {
    writer1, _ := os.Create("dst1.txt")
    writer2, _ := os.Create("dst2.txt")
    multiWriter := io.MultiWriter(writer1, writer2)
    fmt.Println(reflect.TypeOf(multiWriter))
}
</pre>
####Log包
基本常量
<pre>
package main

import (
	"log"
)
func main(){
	//Llongfile完整文件名:行号
	//Lshortfile 文件名LstdFlags日期(Y-m-d)与时间(i:h:s)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println()
}
output==>
2016/03/02 22:05:17 lite.go:10:
</pre>
log.Fatal:
打印日志并且退出；相当于调用了Print()和os.Exit(1)
<pre>
package main

import (
	"log"
)
func main(){
	age := 4
	log.Fatal("Hi age = ",age)
	
}
output==>
2016/03/02 22:07:24 Hi age = 4
进程退出
</pre>
类似的是log.Fatalf:<br>
按格式输出日志，并退出。相当于调用Printf()并调用os.Exit(1)
<pre>
package main
import(
    "log"
)
func main(){
    name := "golang"
    log.Fatalf("%8d,%8s", 23, name)
}
output==>
2016/03/02 22:11:20       23,  golang
进程退出
</pre>
类似的还有log.Fatalln
<pre>
package main

import (
	"log"
)
func main(){
	log.Fatalln("bytes")
}
output==>
2016/03/02 22:15:08 bytes
进程退出
</pre>
####log.New自定义日志输出
这个方法用来自定义logger，指定输出目标、格式等。
<pre>
package main
import (
	"log"
	"os"
)
func main(){
	file,err :=os.OpenFile("copy.txt",os.O_WRONLY,0666) //copy.txt必须已经存在，指定输出位置在copy.txt
	if err != nil{
		panic(err)
	}
	defer file.Close()
	l :=log.New(file,"logger",log.Ldate)
	l.Println("log to file copy.txt")
}
</pre>
####log.Panic
这个方法相当于调用Print()及panic()。<br>
类似的还有log.Panicf和log.Panicln.
<pre>
package main
import(
    "log"
    "fmt"
)
func main(){

    defer func(){
        if err := recover(); err !=nil{
            fmt.Println(err)    //output : "call panic and stop"
            handleException()
        }
    }()
    log.Panic("call panic and stop")
    log.Println("this will not be called.")
}
func handleException(){
    log.Println("recovering...")
}
</pre>
####log.Print
输出日志到标准logger
<pre>
package main

import (
	"log"
)
func main(){
	log.Print("string")
}
output==>
2016/03/02 22:39:04 string
</pre>
####rand
float32
<pre>
package main
import (
    "fmt"
    "math/rand"
)

func main() {
    n := 10
    i := 0
    for i < n {
        x := rand.Float32()
        fmt.Println(x)
        i += 1
    }
}
output==>
0.6046603
0.9405091
0.6645601
0.4377142
0.4246375
0.68682307
0.06563702
0.15651925
0.09696952
0.30091187
</pre>
float64
<pre>
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    n := 10
    i := 0
    for i < n {
        x := rand.Float64()
        fmt.Println(x)
        i += 1
    }
}
output==>
0.6046602879796196
0.9405090880450124
0.6645600532184904
0.4377141871869802
0.4246374970712657
0.6868230728671094
0.06563701921747622
0.15651925473279124
0.09696951891448456
0.30091186058528707
</pre>
####Int
返回一个非负伪随机整数
<pre>
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    n := 10
    i := 0
    for i < n {
        x := rand.Int()
        fmt.Println(x)
        i += 1
    }
}
output==>
5577006791947779410
8674665223082153551
6129484611666145821
4037200794235010051
3916589616287113937
6334824724549167320
605394647632969758
1443635317331776148
894385949183117216
2775422040480279449
</pre>
####New
该函数主要返回了一个新的Rand实例,并以6作为随机值产生器.
<pre>
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    r := rand.New(rand.NewSource(6))
    fmt.Println(r.Int())
}
output==>
3305628230121721621
</pre>
####rand.seed
以提供的参数作为种子值,来初始化随机产生器.如果Seed没有被调用,那么随机数调用前默认调用Seed(1).<br>
小贴士,一般用传入time.Now().UnixNano()给Seed函数来实现不同运行次数看到不同结果的目的.
<pre>
package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
   rand.Seed(time.Now().UnixNano())
   fmt.Println(rand.Int())
}
output==>
4905570427655584144
</pre>
###Net包
net 包是处理网络I/O的包.
####CanonicalHeaderKey
<pre>
package main
//功能说明： 这个函数返回一个规范好的标准的http header字符串。
// 标准化的http header 字符串格式为： 第一个字母和跟着“-”字符
//后面的第一个字母大写，剩下的字符全部小写。 举个例子： 对于一个
//http头（header） "accept-encoding" 来说， 规范好的标准
//化格式就是 "Accept-Encoding"。
import (
	"net/http"
	"fmt"
)
func main(){
	fmt.Println(http.CanonicalHeaderKey("accept-encoding"))
	fmt.Println(http.CanonicalHeaderKey("accept-Language"))
}
output==>
Accept-Encoding
Accept-Language
</pre>
####Error
<pre>
package main

import (
	_ "io"
    "log"
    "net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    // 正常情况的输出,暂时注释掉了
    //io.WriteString(w, "hello, world!\n")

    // 输出用户自定义错误
    http.Error(w, "this is a error", 404)
}

func main() {
    // 指定当用户访问 http://www.xxx.com:mmmm/hello 的时候(注意，请不要在hello后面加上/变成hello/)
    // 调用HelloServer这个函数来处理
    http.HandleFunc("/hello", HelloServer)

    // 侦听本地的8888端口
    // 客户可以通过浏览器来访问
    // 可以输入 http://localhost:8888/hello来访问
    // 其中localhost会被转换为127.0.0.1,这是本地ip地址
    // 这里需要注意的问题有
    // 1. 360误报，可以删除360装qq管家
    // 2. 某些浏览器不识别这个端口，需要手动配置一下,或者你可以使用80端口，8080端口
    //    或者换成chrome浏览器尝试一下
    // 3. 本地防火墙阻止
    // 如果顺利的话，你的浏览器会输出 this is a error

    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
浏览器地址栏中输入localhost:8888/hello
404
</pre>
自定义多个页面链接入口
<pre>
package main

import (
	"log"
	"io"
	"net/http"
)
func helloServer(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"hello world")
}
func IndexPhpServer(w http.ResponseWriter ,req *http.Request){
	io.WriteString(w,"<html><body><font color=red>Yes,这是红的</font></body></html>")
}
func jiang(w http.ResponseWriter,req *http.Request){
	io.WriteString(w,"这是jiang的网站")
}
func main(){
	http.HandleFunc("/hello",helloServer)
	http.HandleFunc("/index.php",IndexPhpServer)
	http.HandleFunc("/jiang",jiang)
	err :=http.ListenAndServe(":8008",nil)
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}
</pre>
####客户端发送JSON数据
<pre>
package main

import (
    "fmt"
    "net/http"
    "strings"
)

func main() {
    json := `{"content":"hello,world"}`
    b := strings.NewReader(json)

    http.Post("http://localhost:8888/hello", "image/jpeg", b)
    fmt.Println("post ok")
}
output==>
post ok
</pre>
####http.NotFound直接输出404
<pre>
package main

import (
	"log"
	"net/http"
)
func helloServer(w http.ResponseWriter , req *http.Request){
	http.NotFound(w,req)
}
func main(){
	http.HandleFunc("/hello",helloServer)
	err :=http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatal("listenAndServe:",err)
	}
}
</pre>
####http.ParseHTTPVersion
用来解析一个HTTP版本字符串
<pre>
package main

import (
	"fmt"
	"net/http"
)
func main(){
	major,minor,ok :=http.ParseHTTPVersion("HTTP/1.1")
	fmt.Println(major,minor,ok) //major主版本号，minor从版本号
}
output==>
1 1 true
</pre>
####http.ProxyFromEnvironment
1. *url.URL 代理的URL,如果没有使用代理或者代理全局变量没有定义则返回nil
2. error 错误，如果没有使用代理或者代理全局变量没有定义则返回nil
3. 功能说明： ProxyFromEnvironment返回给定request的代理url. 一般该URL由用户的环境变量$HTTP_PROXY和$NO_PROXY (或$http_proxy和$no_proxy)指定。 如果用户的全局代理环境无效则返回一个错误。 如果用户没有使用代理或者全局环境变量没有定义则会返回一个nil的URL和一个nil的错误。
<pre>
package main

import (
    "io"
    "log"
    "net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "<html>\n<body>\n")
    io.WriteString(w, "hello, world!<br/>\n")

    myurl, _ := http.ProxyFromEnvironment(req)
    if myurl != nil {
        io.WriteString(w, myurl.Scheme+"<br/>\n")
        io.WriteString(w, myurl.Opaque+"<br/>\n")
        io.WriteString(w, myurl.Host+"<br/>\n")
        io.WriteString(w, myurl.Path+"<br/>\n")
        io.WriteString(w, myurl.RawQuery+"<br/>\n")
        io.WriteString(w, myurl.Fragment+"<br/>\n")
    } else {
        io.WriteString(w, "url is null <br/>\n")
    }
    io.WriteString(w, "</body>\n</html>\n")

}

func main() {

    http.HandleFunc("/hello", HelloServer)

    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
</pre>
####http.Redirect
页面跳转
<pre>
package main

import (
	"log"
	"io"
	"net/http"
)
func helloServer(w http.ResponseWriter, req *http.Request){
	http.Redirect(w,req,"world",http.StatusFound)
}
func worldServer(w http.ResponseWriter,req *http.Request){
	io.WriteString(w,"world server")
}
func main(){
	http.HandleFunc("/hello",helloServer)
	http.HandleFunc("/world",worldServer)
	err :=http.ListenAndServe(":9999",nil)
	if err != nil{
		log.Fatal("listenAndServe:",err)
	}
}
在地址栏中输入localhost:9999/hello,结果是直接跳到world页面中。
</pre>
####http.Serve
tcp监听端口，在端口上启用http服务。
<pre>
package main
import (
	"log"
	"net"
	"io"
	"net/http"
)
func helloServer(w http.ResponseWriter,req *http.Request){
	io.WriteString(w,"hellowolrd server")
}
func main(){
	http.HandleFunc("/hi",helloServer)
	// 首先，创建用tcp协议监听8888端口
	l,e :=net.Listen("tcp",":8888")
	if e != nil{
		log.Fatal("Listen:",e)
	}
	// 然后在监听的这个端口上启用http服务进行http服务
	err:=http.Serve(l,nil)
	if err != nil {
		log.Fatal("serve:",err)
	}
}
</pre>
####http.ServeFile
将本地文件输出到浏览器,通过tcp监听端口，通过http服务。
<pre>
package main

import (
	"log"
	"net"
	"net/http"
)
func helloserver(w http.ResponseWriter,req *http.Request){
	http.ServeFile(w,req,"yes.txt")
}
func main(){
	http.HandleFunc("/hello",helloserver)
	l,e :=net.Listen("tcp",":8009")
	if e != nil {
		log.Fatal("listen:",e)
	}
	err :=http.Serve(l,nil)
	if err != nil {
		log.Fatal("serve:",err)
	}
}
output=>
Just a test!
Just a test!
Just a test!
Just a test!
Just a test!
Just a test!
Just a test!
Just a test!
Just a test!
Just a test!
</pre>
####cookie
<pre>
package main

import (
    "log"
    "net"
    "net/http"
    "time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    expire := time.Now().AddDate(0, 0, 1)
    mycookie := http.Cookie{"test", "testcookie", "/", "www.sanguohelp.com", expire, expire.Format(time.UnixDate),
        86400, true, true, "test=testcookie", []string{"test=tcookie"}}

    http.SetCookie(w, &mycookie)

}

func main() {

    http.HandleFunc("/hello", HelloServer)

    l, e := net.Listen("tcp", ":8888")
    if e != nil {
        log.Fatal("Listen: ", e)
    }

    err := http.Serve(l, nil)
    if err != nil {
        log.Fatal("Serve: ", err)
    }
}
</pre>
####http.StatusText
string对应状态码的字符串
<pre>
package main

import (
	"net/http"
	"fmt"
)
func main(){
	fmt.Println(http.StatusText(404))
	fmt.Println(http.StatusText(202))
}
output==>
Not Found
Accepted
</pre>
###Os包
os.Getwd
<pre>
package main

import (
	"fmt"
	"os"
)
func main(){
	pwd,err :=os.Getwd()
	if err != nil{
		fmt.Printf("Error:%s\n",err)
		return
	}
	fmt.Println("The current directory is:",pwd)
}
output==>
The current directory is: C:\mygo\src\lite
</pre>
####path.Base
这个函数主要是用来返回最后一个元素的路径,如果路径为空返回.如果路径由斜线组成,返回/
<pre>
package main

import (
	"path"
	"fmt"
)
func main(){
	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base(""))  // .    
    fmt.Println(path.Base("////"))  // /
}
output==>
b
</pre>
####path.Dir
这个函数主要是返回path中最后一个元素的路径
<pre>
package main

import (
    "fmt"
    "path"
)

func main() {
    fmt.Println(path.Dir("/a/b/c")) // /a/b
    fmt.Println(path.Dir("")) // .
}
</pre>
####path.Ext
<pre>
package main

import (
    "fmt"
    "path"
)

func main() {
    fmt.Println(path.Ext("/a/b/c/bar.css")) // .css
    fmt.Println(path.Ext("/a/b/c/bar")) // ""
}
</pre>
####path.IsAbs
这个函数主要是判断路径是不是绝对路径，如果是绝对路径返回true
<pre>
package main

import (
    "fmt"
    "path"
)

func main() {
    fmt.Println(path.IsAbs("/home/zzz/go.pdf")) // true
    fmt.Println(path.IsAbs("home/zzz/go.pdf"))  // false
}
</pre>
####path.join
<pre>
package main

import (
    "fmt"
    "path"
)

func main() {
    fmt.Println(path.Join("a", "b", "c")) // a/b/c
    fmt.Println(path.Join("a", "", "c"))  // a/c
    fmt.Println(path.Join("a", "../bb/../c", "c")) // c/c
}
</pre>
####path.Split
这个函数主要是分离路径中的文件目录和文件
<pre>
package main

import (
    "fmt"
    "path"
)

func main() {
    fmt.Println(path.Split("static/myfile.css")) // static/ myfile.css
    fmt.Println(path.Split("static"))   // "" static
}
</pre>
###reflect
reflect.Append返回新的切片.<br>
追加一个切片x值到切片，并返回所创建的Slice。在Go中，每一个x值必须是分配给切片的元素类型。
<pre>
package main

import (
	"fmt"
	"reflect"
)
func main(){
	var a []int
	var value reflect.Value = reflect.ValueOf(&a)
	//判断指针是否指向内存地址
	if !value.CanSet(){
		value =value.Elem() //使指针指向内存地址
	}
	value = reflect.Append(value,reflect.ValueOf(1))
	value = reflect.Append(value, reflect.ValueOf(2))
	value = reflect.Append(value, reflect.ValueOf(3), reflect.ValueOf(4)) //支持可变参数
	fmt.Println(value.Kind(),value.Slice(0, value.Len()).Interface())
}
output==>
slice [1 2 3 4]
</pre>
####reflect.Copy
Copy 复制src的内容复制到dst，直到dst已被填补满，或src已经耗尽。它返回复制的元素的数量。每个 dst 和 src 的 Kind（样）都必须切片(Slice)“或”数组(Array)，dst和src必须具有相同的元素类型。
<pre>
package main
import (
    "fmt"
    "reflect"
)

type A struct {
  A0 []int
  A1 []int
}

func main(){
    var a A
    a.A0 = append(a.A0, []int{1,2,3,4,5,6,7}...)
    a.A1 = append(a.A1, 9, 8, 7, 6)
    var n = reflect.Copy(reflect.ValueOf(a.A0), reflect.ValueOf(a.A1))
    fmt.Println(n, a)
    //>>4 {[9 8 7 6 5 6 7] [9 8 7 6]}}
}
4 {[9 8 7 6 5 6 7] [9 8 7 6]}
</pre>
####reflect.TypeOf(x).Kind()
reflect.Kind 有以下常量成员
reflect.Invalid       // 无效
reflect.Bool          // 布尔
reflect.Int           // 整数（有符号）
reflect.Int8          // 整数8位（有符号）
reflect.Int16         // 整数16位（有符号）
reflect.Int32         // 整数32位（有符号）
reflect.Int64         // 整数64位（有符号）
reflect.Uint          // 整数（无符号）
reflect.Uint8         // 整数8（无符号）
reflect.Uint16        // 整数16（无符号）
reflect.Uint32        // 整数（无符号）
reflect.Uint64        // 整数（无符号）
reflect.Uintptr       // 整数（指针,无符号）
reflect.Float32       // 浮点数32位
reflect.Float64       // 浮点数64位
reflect.Complex64     // 复数64位
reflect.Complex128    // 复数128位
reflect.Array         // 数组
reflect.Chan          // 信道
reflect.Func          // 函数
reflect.Interface     // 接口
reflect.Map           // 地图
reflect.Ptr           // 指针
reflect.Slice         // 切片
reflect.String        // 字符
reflect.Struct        // 结构
reflect.UnsafePointer // 安全指针
<pre>
 package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var a string
        var kind reflect.Kind = reflect.TypeOf(a).Kind()
        fmt.Println(kind, kind == reflect.String, kind == reflect.Int)
        //>>string true false
    }
output==>
string true false
</pre>
###sort包
####sort.Float64s()
Float64s 以升序排列 float64 切片
<pre>
package main 

import (
	"fmt"
	"sort"
)
func main(){
	a :=[]float64{5,6,7,8,9,6,7}
	sort.Float64s(a)
	fmt.Println(a)  //注意这里的fmt.Println若换成println则输出结果不一样
}
output==>
[5 6 6 7 7 8 9]
</pre>
####sort.Float64sAreSorted
sort.Float64sAreSorted判断float64切片是否已经排序过，当然，它已经把float64切片重新排序了.
<pre>
package main

import (
	"sort"
	"fmt"
)
func main(){
	a :=[]float64{2,4,6,7,4,5,6,7,5,7,9}
	fmt.Println(sort.Float64sAreSorted(a))
	sort.Float64s(a)
	fmt.Println(sort.Float64sAreSorted(a))
	fmt.Println(a)
}
output==>
false
true
[2 4 4 5 5 6 6 7 7 7 9]
</pre>
####sort.Ints()
以升序排列int切片
<pre>
package main

import (
	"fmt"
	"sort"
)
func main(){
	a :=[]int{34,5,6,4,3,4,5,6,564,3,43,3,5,566,56,43}
	sort.Ints(a)
	fmt.Println(a)
}
ouptut==>
[3 3 3 4 4 5 5 5 6 6 34 43 43 56 564 566]
</pre>
####sort.IntsAreSorted
判断int切片是否已经排序过
<pre>
package main

import (
	"sort"
	"fmt"
)
func main(){
	a :=[]int{3,4,5,6,5,7,8}
	sort.Ints(a)
	fmt.Println(sort.IntsAreSorted(a))
	fmt.Println(a)
}
output==>
true
[3 4 5 5 6 7 8]
</pre>
####sort.IsSorted
判断数据是否已经按升序排列
<pre>
package main

import (
	"sort"
	"fmt"
)
func main(){
	d :=[]int{23,5,6,45,3,2}
	sort.Sort(sort.IntSlice(d))
	fmt.Println(sort.IsSorted(sort.IntSlice(d)))
	fmt.Println(d)
}
output==>
true
[2 3 5 6 23 45]
</pre>
####sort.Reverse
返回逆序(倒序)的数据
<pre>
package main

import (
	"fmt"
	"sort"
)
func main(){
	s :=[]int{3,5,64,5,6,3,54,4,23}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
output==>
[64 54 23 6 5 5 4 3 3]
</pre>
####sort.Search
二分法找某个数
<pre>
package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
oupput==>
Pick an integer from 0 to 100.
Is your number <= 50? 
Is your number <= 75? 
Is your number <= 88? 
Is your number <= 94? 
Is your number <= 97? 
Is your number <= 99? 
Your number is 100.
</pre>
####StringSlice
有以下方法：Less . Search . Swap。
<pre>
package main

import (
	"fmt"
	"sort"
)
func main(){
	p :=sort.StringSlice{"php","golang","java","python","c"}
	fmt.Println(p.Len())  //5  返回数组/切片的长度
	fmt.Println(p.Less(0,1)) //false 返回 p[i] < p[j] 是否为真
	fmt.Println(p.Search("java"))  //2
	p.Swap(0,1)  //交换键名为0与1的键值
	fmt.Println(p)
	p.Sort()
	fmt.Println(p)  //排序 a-z
}
output==>
5
false
2
[golang php java python c]
[c golang java php python]
</pre>
####Strings
以升序排列string切片
<pre>
package main

import (
	"fmt"
	"sort"
)
func main(){
	a :=[]string{"php","golang","java","python","c"}
	sort.Strings(a)
	fmt.Println(a)
}
output==>
[c golang java php python]
</pre>
####StringsAreSorted
判断是否已经按升序排列
<pre>
package main

import (
	"sort"
	"fmt"
)
func main(){
	a :=[]string{"php","golang","java"}
	fmt.Println(sort.StringsAreSorted(a)) //false
	sort.Strings(a)
	fmt.Println(sort.StringsAreSorted(a)) //true
}
output==>
false
true
</pre>
###Strconv
####strconv.AppendBool
将布尔值 b 转换为字符串 "true" 或 "false" 然后将结果追加到 dst 的尾部，返回追加后的 []byte.
<pre>
package main
import (
    "fmt"
    "strconv"
)

func main() {
   	list := strconv.AppendBool(make([]byte, 3), false)
    fmt.Println(list)//[0 0 0 102 97 108 115 101]
    newlist := strconv.AppendBool(list, true)
    fmt.Println(newlist)//[0 0 0 102 97 108 115 101 116 114 117 101]
}
output==>
[0 0 0 102 97 108 115 101]
[0 0 0 102 97 108 115 101 116 114 117 101]
</pre>
####Atoi
Atoi是函数ParseInt(s, 10, 0)的简写。把字符串格式的数字如“12345”转化为数字12345.
<pre>
package main
import (
    "fmt"
    "strconv"
)
func main() {
    fmt.Println(strconv.Atoi("12345"))
    fmt.Println(strconv.Atoi("abcde"))
}
output==>
12345 <nil>
0 strconv.ParseInt: parsing "abcde": invalid syntax
</pre>
####FormatBool
将true或者false转换成字符串
<pre>
package main

import (
    "fmt"
    "strconv"
)

func main() {
    b := strconv.FormatBool(true)
    fmt.Println(b)
    b = strconv.FormatBool(false)
    fmt.Println(b)
}
output==>
true
false
</pre>
####Cond
*cond:创建Cond。
<br>Wait():添加waiter，使用时要注意先调用c.L.Lock()。
<br>Broadcast():唤醒所有waiter，包括上一代和新一代。
<pre>
package main

import (
    "fmt"
    "time"
    "sync"
)

func waiter(cond *sync.Cond,id int) {
    cond.L.Lock()
    cond.Wait()
    cond.L.Unlock()
    fmt.Printf("Waiter:%d wake up!\n",id)
}

func main() {
    locker := new(sync.Mutex)
    cond := sync.NewCond(locker)  //使用Mutex作为Locker

    for i := 0; i < 3; i++ {        //生成waiter
        go waiter(cond,i)
    }
    time.Sleep(time.Second * 1)     //等待waiter到位

    cond.L.Lock()
    cond.Signal()                   //唤醒一个waiter
    cond.L.Unlock()

    for i := 3; i < 5; i++ {        //生成新一代waiter
        go waiter(cond,i)
    }
    time.Sleep(time.Second * 1)

    cond.L.Lock()
    cond.Signal()                   //唤醒的将是上一代（id<3）的waiter之一
    cond.L.Unlock()

    cond.L.Lock()
    cond.Broadcast()                //唤醒所以waiter
    cond.L.Unlock()
    time.Sleep(time.Second * 1)
}
output==>
Waiter:0 wake up!
Waiter:1 wake up!
Waiter:2 wake up!
Waiter:3 wake up!
Waiter:4 wake up!
</pre>
####Mutex
state:成员state用表明当前锁是处于被占用（state==1）还是空闲（state==0）。使用atomic.CompareAndSwapInt32()进行修改。<br>
sema:当前锁占用失败，系统监听成员sema。当sema==1，表明锁被释放。当前争用的的过程将被系统唤醒尝试去获取锁。当sema==0，当前争用过程会被系统暂停。<br>
Lock()和Unlock()：获取和释放当前锁。
<pre>
package main


import (
    "fmt"
    "runtime"
    "sync"
)

func click(total *int,ch chan int) {
    for i := 0; i < 1000; i++ {
        *total += 1
    }
    ch <- 1
}

func clickWithMutex(total *int,m *sync.Mutex, ch chan int) {
    for i := 0; i < 1000; i++ {
        m.Lock()
        *total += 1
        m.Unlock()
    }
    ch <- 1
}


func main() {

    runtime.GOMAXPROCS(2)       //使用多个处理器，不然都是顺序执行。

    m := new(sync.Mutex)
    count1 := 0;
    count2 := 0;

    ch := make(chan int, 200)       //保证输出时count完了

    for i := 0; i < 100; i++ {
        go click(&count1, ch)
    }
    for i := 0; i < 100; i++ {
        go clickWithMutex(&count2, m, ch)
    }

    for i := 0; i < 200; i++ {
        <-ch
    }

    fmt.Printf("count1:%d\ncount2:%d\n", count1,count2)
}
output==>
count1:100000
count2:100000
</pre>
####Once
<pre>
package main

import (
	"fmt"
	"sync"
)
func main(){
	once :=new(sync.Once)
	ch :=make(chan int,3)
	for i:=0;i<3;i++{
		go func(x int){
			once.Do(func(){
				fmt.Printf("once %d\n",x)
			})
			fmt.Printf("%d\n",x)
			ch <-1
		}(i)
	}
	for i :=0;i<3;i++{
		<- ch 
	}
}
output==>
once 0
0
1
2
</pre>
####RWMutex
成员

w ： 用于写锁，拒绝其他写操作。
writerSem ： 写锁信号量，用于阻塞或唤醒写锁争夺锁的行为。
readerSem ： 读锁信号量，写锁来时唤醒要进来的读锁。
readerCount ： 读锁的持有数。
readerWait ： 写锁需要等待的读操作数，最后一个读离开唤醒争用写锁的过程。

- Rlock()  :获取读锁，当之前以有一个写锁存在，阻塞。
- RUlock() :放弃读锁，当存在有尝试进入的写锁且当前是最后一个读锁时，唤醒写锁争用过程。
- Lock() :获取写锁，当之前存在读锁，阻塞。
-  Unlock() :放弃写锁，当存在有读锁争用过程被阻塞，唤醒所以读锁。
<pre>
package main


import (
    "fmt"
    "runtime"
    "sync"
)

func clickWithMutex(total *int,m *sync.RWMutex, ch chan int) {
    for i := 0; i < 1000; i++ {
        m.Lock()
        *total += 1
        m.Unlock()

        if i==500 {
            m.RLock()
            fmt.Println(*total)
            m.RUnlock()
        }
    }
    ch <- 1
}


func main() {

    runtime.GOMAXPROCS(2)       //使用多个处理器，不然都是顺序执行。

    m := new(sync.RWMutex)
    count := 0;

    ch := make(chan int, 10)        //保证输出时count完了

    for i := 0; i < 10; i++ {
        go clickWithMutex(&count, m, ch)
    }

    for i := 0; i < 10; i++ {
        <-ch
    }

    fmt.Printf("count:%d\n", count)
}
output==>
501
1109
4048
4944
5856
5886
6447
6753
6826
8630
count:10000
</pre>
####WaitGroup
- Add(1):添加delta个元素进入WaitGroup。个人觉得这里的接口设计有问题，如果Add(2)，那就要掉用两次Done()。而且会有Add(2)的需求吗？
- Done():对WaitGroup的counter减1。
- Wait():阻塞，直到WaitGroup中的所以过程完成。
<pre>
package main

import (
	"fmt"
	"sync"
)
func wgProcess(wg *sync.WaitGroup,id int){
	fmt.Printf("process %d is going!\n",id)
	wg.Done()
}
func main(){
	wg :=new(sync.WaitGroup)
	for i:=0;i<3;i++{
		wg.Add(1)
		go wgProcess(wg,i)
	}
	wg.Wait()
}
output==>
process 2 is going!
</pre>
####syscall
获取磁盘空间
<pre>
package main
import (
	"log"
	"unsafe"
	"syscall"
)
//获取磁盘空间
func main() {
     h := syscall.MustLoadDLL("kernel32.dll")
    c := h.MustFindProc("GetDiskFreeSpaceExW")
    lpFreeBytesAvailable := int64(0)
    lpTotalNumberOfBytes := int64(0)
    lpTotalNumberOfFreeBytes := int64(0)
    r2, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("F:"))),
        uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
        uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
        uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
    if r2 != 0 {
        log.Println(r2, err, lpFreeBytesAvailable/1024/1024)
    }
}
output==>
2016/03/05 14:22:53 1 The operation completed successfully. 33265
</pre>
####scanner
初始化一个词法解析器。
<pre>
package main

import (
	"text/scanner"
	"fmt"
	"strings"
)
func main(){
	src:=strings.NewReader("int hello = 3;hello+23;print hello;")
	fmt.Println(src)
	var s scanner.Scanner
	s.Init(src)
	
	tok :=s.Scan()
	fmt.Println(s.TokenText())
	for tok !=scanner.EOF {
		tok = s.Scan()
		fmt.Println(s.TokenText())
	}
}
output==>
&{int hello = 3;hello+23;print hello; 0 -1}
int
hello
=
3
;
hello
+
23
;
print
hello
;
</pre>
###Time
####Add
返回时间（t + d）
<pre>
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println("now:", now)
    fmt.Println("after 3 hours", now.Add(3*time.Hour))
}
output==>
now: 2016-03-05 14:34:07.2542474 +0800 +0800
after 3 hours 2016-03-05 17:34:07.2542474 +0800 +0800
</pre>
####AddDate
生成增加后的年月日之后的时间。
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	now :=time.Now()
	fmt.Println(now)
	fmt.Println(now.AddDate(2,3,4))
}
output==>
2016-03-05 14:36:26.3602038 +0800 +0800
2018-06-09 14:36:26.3602038 +0800 +0800
</pre>
####After
等待指定时间段之后将当前时间发送给返回的chan中。等价于NewTimer(d).C
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	result :=make(chan int)
	go func(ch chan int){
		time.Sleep(3 * time.Second)
		ch <- 4
	}(result)
	select {
		case <- time.After(2 * time.Second):
		fmt.Println("time out")
		case <- result:
		fmt.Println(result)
	}
}
output==>
time out
</pre>
####Format
返回根据layout指定的格式格式化之后的字符串，layout定义了标准时间的显示格式。预定义的layout有ANSIC，UnixDate，RFC3339等。
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	t :=time.Now()
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.UnixDate))
}
output==>
Sat Mar  5 14:44:27 2016
Sat Mar  5 14:44:27 +0800 2016
</pre>
####Month
返回一年中的某个月
<pre>
package main

import (
    "fmt"
    "time"
)

func main() {
    _, M, _ := time.Now().Date()
    fmt.Println(M)
}
output==>
March
</pre>
####ticker
新建一个Ticker，包含了time channel，每隔指定d间隔的时间发送时间给这个channel。d必须大于0，否则函数会崩溃.每隔一段时间操作一次指定动作。
<pre>
package main

import (
	"fmt"
	"time"
)
func tick(ch <-chan time.Time){
	for t :=range ch {
		fmt.Println(t)
	}
}
func main(){
	ticker :=time.NewTicker(time.Second)
	go tick(ticker.C)
	time.Sleep(5 * time.Second)
}
output==>
2016-03-05 14:52:16.2305333 +0800 +0800
2016-03-05 14:52:17.2325906 +0800 +0800
2016-03-05 14:52:18.2306477 +0800 +0800
2016-03-05 14:52:19.2307049 +0800 +0800
2016-03-05 14:52:20.2307621 +0800 +0800
</pre>
####timer
新建一个Timer，在时间d之后将当前时间发送给C.
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	timer :=time.NewTimer(2 * time.Second)
	t := <- timer.C
	fmt.Println(t)
}
output==>
2016-03-05 14:55:39.2671463 +0800 +0800
</pre>
####utf16.Decode()
将utf-16序列 解码成Unicode字符序列并返回
<pre>
package main

import (
 "fmt"
"unicode/utf16"
)
func main() {
u := []uint16{72, 101, 108, 108, 111, 32, 19990, 30028}
s := utf16.Decode(u)
fmt.Printf("%c", s)
//[H e l l o   世 界]
}
output==>
[H e l l o   世 界]
</pre>
####utf16.Encode
将s编码成 UTF-16 序列并返回.
<pre>
package main

import (
	"fmt"
	"unicode/utf16"
)
func main(){
	s :=[]rune("hello世界")
	u :=utf16.Encode(s)
	fmt.Printf("%v",u)
}
output==>
[104 101 108 108 111 19990 30028]
</pre>
####strings.HasPrefix
判断某个字符串的前缀是否是某个字符。
<pre>
package main
import (
	"strings"
	"fmt"
)
func main(){
	var str string ="this is an example of a string\n"
	fmt.Printf("%s",str)
	fmt.Printf("%t\n",strings.HasPrefix(str,"th"))
}
output==>
this is an example of a string
true
</pre>
####strings.HasSuffix
<pre>
package main
import (
	"strings"
	"fmt"
)
func main(){
	var str string ="this is an example of a string\n"
	fmt.Printf("%s",str)
	fmt.Printf("%t\n",strings.HasSuffix(str,"nn"))
}
output==>
this is an example of a string
false
</pre>
####Count
输出字符串中有多少个指定字符
<pre>
package main

import (
	"strings"
	"fmt"
)
func main(){
	var str string = "this is it going"
	fmt.Println(strings.Count(str,"i"))
}
output==>
4
</pre>
####Repeat
重复字符串组成新的字符串。
<pre>
package main
import (
	"fmt"
	"strings"
)
func main(){
	var or string = "hi "
	var new string
	new =strings.Repeat(or,4)
	fmt.Printf("%s",new)
}
output==>
hi hi hi hi
</pre>
####ToLower与ToUpper
大小写转换
<pre>
package main

import (
	"strings"
)
func main(){
	var or string = "hey,how are you"
	var lower string
	var upper string
	lower = strings.ToLower(or)
	upper= strings.ToUpper(or)
	println(lower)
	println(upper)
}
output==>
hey,how are you
HEY,HOW ARE YOU
</pre>
####strings.Fields
<pre>
package main

import (
	"fmt"
	"strings"
)
func main(){
	str :="the quick brown fox jumps over the lazy dog"
	//会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割
	//成若干小块，组成一个slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。
	//相当于PHP中的implode(" ",)
	sl :=strings.Fields(str)
	fmt.Printf("分割形成的切片是:%v\n",sl)
	for _,val :=range sl{
		fmt.Printf("%s -",val)
	}
}
output==>
分割形成的切片是:[the quick brown fox jumps over the lazy dog]
the -quick -brown -fox -jumps -over -the -lazy -dog -
</pre>
####strings.Split与strings.Join
类似与PHP中的explode与implode。
<pre>
package main
import (
	"fmt"
	"strings"
)
func main(){
	str2 := "GO1|The ABC of Go|25"
	sl :=strings.Split(str2,"|")//字符串分割成切片
	for _,val :=range sl {
		fmt.Println(val)
	}
	str3 :=strings.Join(sl,",")  //将切片组合成字符串
	fmt.Println(str3)
}
output==>
GO1
The ABC of Go
25
GO1,The ABC of Go,25
</pre>
当在进行大量的计算时，提升性能最直接有效的一种方式就是避免重复计算。通过在内存中缓存和重复利用相同计算的结果，称之为内存缓存。最明显的例子就是生成斐波那契数列的程序

####Copy与Append
<pre>
package main
import "fmt"

func main() {
    sl_from := []int{1, 2, 3}
    sl_to := make([]int, 10)

    n := copy(sl_to, sl_from)   //复制
    fmt.Println(sl_to)
    fmt.Printf("Copied %d elements\n", n) // n == 3

    sl3 := []int{1, 2, 3}
    sl3 = append(sl3, 4, 5, 6)    //追加
    fmt.Println(sl3)
}
output==>
[1 2 3 0 0 0 0 0 0 0]
Copied 3 elements
[1 2 3 4 5 6]
</pre>
###channel消息传递
<pre>
package main
//生产者－消费者模型的消息传递
import (
	"time"
	"fmt"
)
func producer (queue chan<- int){  //定义一个只写的channel
	for i:=0;i<10;i++{
		queue <- i
	}
}
func consumer(queue <-chan int){  // 定义一个只读的channel
	for i :=0;i<10;i++{
		v :=<-queue
		fmt.Println("receive:",v)
	}
}
func main(){
	queue :=make(chan int,1)
	go producer(queue)
	go consumer(queue)
	time.Sleep(1e9)  //等待producer与consumer完成
}
output==>
receive: 0
receive: 1
receive: 2
receive: 3
receive: 4
receive: 5
receive: 6
receive: 7
receive: 8
receive: 9
</pre>
###select
go中用select可以等待多个channel
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	c1 :=make(chan string)
	c2 :=make(chan string)
	go func (){
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func(){
		time.Sleep(time.Second * 2)
		c2 <- "c2"
	}()
	for i:=0;i<2;i++{
		select {
			case msg1 :=<-c1:
			fmt.Println("receive:",msg1)
			case msg2 :=<- c2:
			fmt.Println("receive:",msg2)
		}
	}
}
output==>
receive: one
receive: c2
</pre>
###sync.WaitGroup
程序中需要并发，需要创建多个goroutine，并且一定要等到这些并发全部完成后才继续接下来的程序执行。WaitGroup的特点是Wait()可以用来阻塞直到队列中的所有任务都完成时才解除阻塞，而不需要sleep一个固定的时间来等待，但是其缺点时候无法指定固定的goroutine数目。
<pre>
package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func function(num int){
	fmt.Println(num)
	wg.Done()
}
func main(){
	for i:=0;i<20;i++{
		wg.Add(1)
		go function(i)
	}
	wg.Wait()
}
output==>
19
0
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
</pre>
###无缓冲channel
<pre>
package main
func afunc(ch chan int){
	println("yes")
	<- ch 
}
func main(){
	ch :=make(chan int)
	go afunc(ch)
	ch <- 1
}
output==>
"yes"
</pre>
代码分析：
首先创建一个无缓冲channel ch,然后执行go afunc()此时执行<-ch,则afunc便会阻塞，不再继续往下执行，直到主进程中ch <- 1向channel ch 中注入数据才会解除afunc该协程的阻塞。
<br>无缓冲总结：
对于无缓存的channel,放入channel和从channel中向外面取数据这两个操作不能放在同一个协程中，防止死锁的发生；同时应该先利用go 开一个协程对channel进行操作，此时阻塞该go 协程，然后再在主协程中进行channel的相反操作（与go 协程对channel进行相反的操作），实现go 协程解锁．即必须go协程在前，解锁协程在后．
###select超时处理案例
<pre>
package main
import (
	"time"
	"fmt"
)
func main(){
	c :=make(chan int)
	o :=make(chan bool)
	go func(){
		select {
			case  i := <- c :
			fmt.Println(i)
			//设置超时时间为３ｓ，如果channel　3s钟没有响应，一直阻塞，则报告超时，进行超时处理
			case <- time.After(time.Duration(3) * time.Second) :
			fmt.Println("3秒内没有数据,超时处理")
			o <- true
			break
		}
	}()
	<- o
	
}
output==>
3秒内没有数据,超时处理
</pre>
###channel死锁的一种情况
<pre>
package main
func afunc(ch chan int){
	println("ch")
	<- ch
}
func main(){
	ch :=make( chan int)
	for i:=0;i<10;i++{
		
		go afunc(ch)
		ch <- 1
		ch <- 1
		ch <- 1
		ch <- 1
	}
}
output==>
ch
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	C:/mygo/src/lite/main.go:12 +0xce
运行错误
</pre>
上面这段运行和之前那一段基本上原理是一样的，但是运行后却会发生死锁。为什么呢？其实总结起来就一句话，"放得太快，取得太慢了"。
按理说，我们应该在我们主routine中创建子goroutine并每次向channel中放入数据，而子goroutine负责从channel中取出数据。但是我们的这段代码在创建了子goroutine后，每个routine会向channel中放入5个数据。这样，每向channel中放入6个数据才会执行一次取出操作，这样一来就可能会有某一时刻，channel已经满了，但是所有的routine都在执行放入操作(因为它们当前执行放入操作的概率是执行取出操作的6倍)，这样一来，所有的routine都阻塞了，从而导致死锁。
###channel避免空读取
<pre>
package main

import (
	"fmt"
)
func main(){
	ch :=make(chan int,5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	//如果这里没有close,循环到第四个以后就会出错,因为对取channel行为一直存在
	for {
		data,ok :=<-ch
		if !ok{
			break
		}
		fmt.Println(data)
	}
}
output==>
1
2
3
4
</pre>
####通过控制channel的个数来控制输出结果的个数。
<pre>
package main

import (
	"fmt"
)
func afunc(ch chan int){
	fmt.Println("finish")
	<- ch
}
func main(){
	ch :=make(chan int)
	//通过控制channelcount数值大小，来控制输出的值的个数
	//这里如果channelcount为2 ，则输出两个finished
	channelcount := 1
	for i:=0;i<channelcount;i++{
		go afunc(ch)
	}
	for i:=0;i<channelcount;i++{ 
		ch <- 1
	}
	
}
output==>
finish
</pre>
同步时main取得协程的变量。
<pre>
package main
import (
	"fmt"
)
var a string
var c =make(chan int,10)
func f(){
	fmt.Println("那我跟着跑")
	a = "hello world"
	c <- 0
}
func main(){
	fmt.Println("我是main函数，我先跑")
	go f()
	<- c
	fmt.Println("在这里我取到了a的值",a)
}
output==>
我是main函数，我先跑
那我跟着跑
在这里我取到了a的值 hello world
</pre>
####golang大小端的判断
<pre>
package main
import (
    "fmt"
    "unsafe"
)
 
const N int = int(unsafe.Sizeof(1)) //取一个int型占用字节数
 
func main() {
    x := 0x1234 // 4*4一共占用16位，2字节
    p := unsafe.Pointer(&x) //取地址
    fmt.Printf("sizeof N %v\n", N)
    p2 := (*[N]byte)(p) // 32位，4字节 类型转化
    fmt.Printf("%v %v %v %v \n", p2[0], p2[1], p2[2], p2[3]) 
    if p2[0] == 0 {
        fmt.Println("本机器：大端") // 
    } else {
        fmt.Println("本机器：小端") // 52（=3*16+4） 18（=1*16+2） 0 0
    }
}
output==>
sizeof N 8
52 18 0 0 
本机器：小端
</pre>
###Channel实现类似并行
<pre>
package main
import (
	"time"
	"fmt"
)
var quit chan int
func foo(id int){
	fmt.Println(id)
	time.Sleep(time.Second *3)  //所有输出完之后会停顿3秒
	quit <- 0
}
func main(){
	count := 10
	quit =make(chan int,count)  //缓冲
	for i :=0;i<count;i++{
		go foo(i)
	}
	for i:=0;i<count;i++{
		<- quit
	}
}
output==>
0
1
2
3
4
5
6
7
8
9
</pre>
####下面的情况还是单核运行
<pre>
package  main

import (
	"runtime"
	"fmt"
)
var quit chan int = make(chan int)
func loop(){
	for i:=0;i<10;i++{
//显式地让出CPU时间给其他goroutine，结果是两个一样的同时输出
		//runtime.Gosched() 
		fmt.Printf("%d\n",i)
	}
	quit <- 0
}
func main(){
	core :=runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	go loop()
	go loop()
	println("计算机的可运行核心个数是：",core)
	for i:=0;i<2;i++{
		<- quit
	}
}
output==>
计算机的可运行核心个数是： 4
0
1
2
3
4
5
6
7
8
9
0
1
2
3
4
5
6
7
8
9
</pre>
####g关于goroutine的一个问题
下面的代码并没有任何输出，why？
<pre>
package main
import (
	"fmt"
)
var ch chan int
func say(s string){
	for i:=0;i<5;i++{
		fmt.Println(s)
	}
	ch <- 1
}
func main(){
	go say("world")

	
}
output ==>

</pre>
解析：这里Go仍然在使用单核，for死循环占据了单核CPU所有的资源，而main线和say两个goroutine都在一个线程里面， 所以say没有机会执行。解决方案还是两个：<br>
允许Go使用多核(runtime.GOMAXPROCS)；
手动显式调动(runtime.Gosched)。
###总结
关于runtime包几个函数:

1. Gosched 让出cpu
2. NumCPU 返回当前系统的CPU核数量
3. GOMAXPROCS 设置最大的可同时使用的CPU核数

Goexit 退出当前goroutine(但是defer语句会照常执行)
我们从例子中可以看到，默认的, 所有goroutine会在一个原生线程里跑，也就是只使用了一个CPU核。<br>
在同一个原生线程里，如果当前goroutine不发生阻塞，它是不会让出CPU时间给其他同线程的goroutines的，这是Go运行时对goroutine的调度，我们也可以使用runtime包来手工调度。
####真正的并行小案例
<pre>
package main

import (
	"runtime"
	"fmt"
)
var quit chan int = make(chan int)
func loop(id int){
	for i:=0;i<20;i++{
		fmt.Printf("%d",id)
	}
	quit <- 0
}
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i:=0;i<3;i++{
		go loop(i)
	}
	for i:=0;i<3;i++{
		<- quit
	}
}
output==>
000211111111111112000000211111112222222222222222200000000000
</pre>
执行它我们会发现以下现象:

有时会发生抢占式输出(说明Go开了不止一个原生线程，达到了真正的并行)
有时会顺序输出, 打印完0再打印1, 再打印2(说明Go开一个原生线程，单线程上的goroutine不阻塞不松开CPU)
那么，我们还会观察到一个现象，无论是抢占地输出还是顺序的输出，都会有那么两个数字表现出这样的现象:

一个数字的所有输出都会在另一个数字的所有输出之前
原因是， 3个goroutine分配到至多4个线程上，就会至少两个goroutine分配到同一个线程里，单线程里的goroutine 不阻塞不放开CPU, 也就发生了顺序输出。
<<<<<<< HEAD

####json
<pre>
package main  
import (  
  "fmt"
  "encoding/json"
)
type Person struct {  
  FirstName string `json:"first_name"` //FirstName <=> firest_name 
  LastName string `json:"last_name"` 
  MiddleName string `json:"middle_name,omitempty"` 
} 
func main() {  
  json_string := ` { "first_name": "John", "last_name": "Smith" }`
  person := new(Person)
  json.Unmarshal([]byte(json_string), person) //将json数据转为Person Struct 
  fmt.Println(person) 
  new_json, _ := json.Marshal(person) //将Person Sturct 转为json格式   
  fmt.Printf("%s\n", new_json) 
} 
output==>
&{John Smith }
{"first_name":"John","last_name":"Smith"}
</pre>
各种类型转成json格式
<pre>
package main 
import ( 
    "encoding/json" 
    "fmt" 
) 

//tag中的第一个参数是用来指定别名 
//比如Name 指定别名为 username `json:"username"` 
//如果不想指定别名但是想指定其他参数用逗号来分隔 
//omitempty 指定到一个field时 如果在赋值时对该属性赋值 或者 对该属性赋值为 zero value 
//那么将Person序列化成json时会忽略该字段 
//- 指定到一个field时 
//无论有没有值将Person序列化成json时都会忽略该字段 
//string 指定到一个field时 
//比如Person中的Count为int类型 如果没有任何指定在序列化 
//到json之后也是int 比如这个样子 "Count":0 
//但是如果指定了string之后序列化之后也是string类型的 
//那么就是这个样子"Count":"0" 
type Person struct { 
    Name        string `json:"username"` 
    Age         int 
    Gender      bool `json:",omitempty"` 
    Profile     string 
    OmitContent string `json:"-"` 
    Count       int    `json:",string"` 
} 

func main() { 

    var p *Person = &Person{ 
        Name:        "brainwu", 
        Age:         21, 
        Gender:      true, 
        Profile:     "I am ghj1976", 
        OmitContent: "OmitConent", 
    } 
    if bs, err := json.Marshal(&p); err != nil { 
        panic(err) 
    } else { 
        //result --> {"username":"brainwu","Age":21,"Gender":true,"Profile":"I am ghj1976","Count":"0"} 
        fmt.Println(string(bs)) 
    } 

    var p2 *Person = &Person{ 
        Name:        "brainwu", 
        Age:         21, 
        Profile:     "I am ghj1976", 
        OmitContent: "OmitConent", 
    } 
    if bs, err := json.Marshal(&p2); err != nil { 
        panic(err) 
    } else { 
        //result --> {"username":"brainwu","Age":21,"Profile":"I am ghj1976","Count":"0"} 
        fmt.Println(string(bs)) 
    } 

    // slice 序列化为json 
    var aStr []string = []string{"Go", "Java", "Python", "Android"} 
    if bs, err := json.Marshal(aStr); err != nil { 
        panic(err) 
    } else { 
        //result --> ["Go","Java","Python","Android"] 
        fmt.Println(string(bs)) 
    } 

    //map 序列化为json 
    var m map[string]string = make(map[string]string) 
    m["Go"] = "No.1" 
    m["Java"] = "No.2" 
    m["C"] = "No.3" 
    if bs, err := json.Marshal(m); err != nil { 
        panic(err) 
    } else { 
        //result --> {"C":"No.3","Go":"No.1","Java":"No.2"} 
        fmt.Println(string(bs)) 
    } 
}


</pre>
###Golang之RPC接触
RPC（Remote Procedure Call Protocol）——远程过程调用协议，它是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。<br>
golang服务器端代码：这里暴露了一个RPC接口，一个HTTP接口
<pre>
package main 

import ( 
    "fmt" 
    "io" 
    "net" 
    "net/http" 
    "net/rpc" 
) 

type Watcher int 

func (w *Watcher) GetInfo(arg int, result *int) error { 
    *result = 1 
    return nil 
} 

func main() { 

    http.HandleFunc("/ghj1976", Ghj1976Test) 

    watcher := new(Watcher) 
    rpc.Register(watcher) 
    rpc.HandleHTTP() 

    l, err := net.Listen("tcp", ":1234") 
    if err != nil { 
        fmt.Println("监听失败，端口可能已经被占用") 
    } 
    fmt.Println("正在监听1234端口") 
    http.Serve(l, nil) 
} 

func Ghj1976Test(w http.ResponseWriter, r *http.Request) { 
    io.WriteString(w, "<html><body>ghj1976-123</body></html>") 
}
</pre>
客户端代码：
<pre>
package main 

import ( 
    "fmt" 
    "net/rpc" 
) 

func main() { 
    client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234") 
    if err != nil { 
        fmt.Println("链接rpc服务器失败:", err) 
    } 
    var reply int 
    err = client.Call("Watcher.GetInfo", 1, &reply) 
    if err != nil { 
        fmt.Println("调用远程服务失败", err) 
    } 
    fmt.Println("远程服务返回结果：", reply) 
}
=======
####一个问题
<pre>
package main
import (
	"fmt"
)
var complete chan int = make(chan int)

func loop() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", i)
    }
	println("time") //只输出一次，why?
    complete <- 0 
}


func main() {
    go loop()
    <- complete // 直到线程跑完, 取到消息. main在此阻塞住
}
output==>
0 time1 2 3 4 5 6 7 8 9     //这里time只输出一次，why?
</pre>
####普通生成器模型
<pre>
package main
//普通生成器模型
import "fmt"
import "math/rand"
func rand_generator_2() chan int {
   out := make(chan int)
   go func() {
       for {
       out <- rand.Int()
       }
  }()
  return out
} 
func main() {
    rand_service_handler :=rand_generator_2()
    fmt.Printf("%d\n",<-rand_service_handler)
}
</pre>
####golang的版本信息
<pre>
package main

import (
	"runtime"
	"fmt"
)
func main(){
	fmt.Println(runtime.Version())
}
output==>
go1.4rc2
>>>>>>> 99a36c9796d03006b29c940be236a036f19dbaf9
</pre>
####Go实现AES加解密
AES简介
密码学中的高级加密标准（Advanced Encryption Standard，AES），又称Rijndael加密法，这个标准用来替代原先的DES。AES加密数据块分组长度必须为128bit，密钥长度可以是128bit、192bit、256bit中的任意一个。<br>
AES也是对称加密算法。
核心代码如下：
<pre>
func AesEncrypt(origData, key []byte) ([]byte, error) {
     block, err := aes.NewCipher(key)
     if err != nil {
          return nil, err
     }
     blockSize := block.BlockSize()
     origData = PKCS5Padding(origData, blockSize)
     // origData = ZeroPadding(origData, block.BlockSize())
     blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
     crypted := make([]byte, len(origData))
     // 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
     // crypted := origData
     blockMode.CryptBlocks(crypted, origData)
     return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
     block, err := aes.NewCipher(key)
     if err != nil {
          return nil, err
     }
     blockSize := block.BlockSize()
     blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
     origData := make([]byte, len(crypted))
     // origData := crypted
     blockMode.CryptBlocks(origData, crypted)
     origData = PKCS5UnPadding(origData)
     // origData = ZeroUnPadding(origData)
     return origData, nil
}
</pre>
完整的是：
<pre>
package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

func main() {
	testAes()
}

func testAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9fd&fwfl")
	result, err := AesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
</pre>
##Go By Example http://gobyexample.everyx.in/
####变参函数
这个函数使用任意数目的 int 作为参数。
如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)。
<pre>
package main

import (
	"fmt"
)
func sum(nums ...int){
	fmt.Print(nums," ")
	total := 0
	for _,num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main(){
	sum(1,2)
	sum(1,2,3,4)
	nums :=[]int{4,5,6,7,8,9}
	sum(nums...)
}
output==>
[1 2] 3
[1 2 3 4] 10
[4 5 6 7 8 9] 39
</pre>
####闭包
我们调用 intSeq 函数，将返回值（也是一个函数）赋给nextInt。这个函数的值包含了自己的值 i，这样在每次调用 nextInt 是都会更新 i 的值。通过多次调用 nextInt 来看看闭包的效果。
<pre>
package main

import (
	"fmt"
)
func intSeq() func() int{
	i := 0
	return func() int{
		i += 1
		return i
	}
}
func main(){
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	newInts := intSeq()
	fmt.Println(newInts())
}
output==>
1
2
3
1
</pre>
####递归
face 函数在到达 face(0) 前一直调用自身。
<pre>
package main

import (
	"fmt"
)
func fact(n int) int{
	if n == 0{
		return 1
	}
	return n*fact(n-1)
}
func main(){
	fmt.Println(fact(7))
}
output==>
5040
</pre>
####指针
<pre>
package main

import (
	"fmt"
)
func zeroval(ival int){
	ival = 0
}
func zeroptr(iptr *int){
	*iptr = 0
}
func main(){
	i := 1
	fmt.Println("initial:",i)
	zeroval(i)
	fmt.Println("zeroval:",i)
	zeroptr(&i)
	fmt.Println("zeroptr:",i)
	//通过 &i 语法来取得 i 的内存地址，例如一个变量i 的指针
	fmt.Println("pointer:",&i) 
	
}
output==>
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc0820062e0
</pre>
####结构体
<pre>
package main

import (
	"fmt"
)
type person struct {
	name string
	age int
}
func main(){
	fmt.Println(person{"jason",34})
	s := person{"bob",12}
	fmt.Println(s.name)
	s.age = 66
	fmt.Println(s.age)
	sp :=&s.age
	fmt.Println(sp)
	
}
output==>
{jason 34}
bob
66
0xc08200a650
</pre>
####方法
Go 自动处理方法调用时的值和指针之间的转化。你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据
<pre>
package main

import (
	"fmt"
)
type rect struct {
	width,height int
}
func (r *rect)area() int {  //这一种可以将函数作为变量的方法
	return r.width * r.height
}
func perim(r rect) int {     //这一种只是普通的函数
	return 2*r.width + 2*r.height
}
func main(){
	r :=rect{width:4,height:5}
	fmt.Println("area:",r.area())  //调用方法之一
	fmt.Println("perim:",perim(r))  //调用方法之二，对比上面的
}
output==>
area: 20
perim: 18
</pre>
####接口
<pre>
package main

import (
	"fmt"
	"math"
)
type geometry interface{  //声明接口
	area() float64
	perim() float64
}
type rect struct {
	width,height float64
}
type circle struct {
	radius float64
}
func (r rect) area() float64{//实现接口中的area()方法
	return r.width * r.height
}
func (r rect) perim() float64{//实现接口中的perim()方法
	return 2*r.width + 2*r.height
}
func (c circle) area() float64{//实现接口中的area()方法
	return math.Pi * c.radius*c.radius
}
func (c circle) perim() float64 {//实现接口中的perim()方法
	return 2*math.Pi*c.radius
}
func measure(g geometry){
	fmt.Println(g)  //打印本身
	fmt.Println(g.area()) //打印面积
	fmt.Println(g.perim()) //打印周长
}
func main(){
	r :=rect{width:3,height:4}
	c :=circle{radius:5}
	measure(r)
	measure(c)
}
output==>
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
</pre>
####协程
<pre>
package main

import (

	"fmt"
)
func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}
func main(){
	f("direct")
	go f("goroutine")
	go func(msg string){//该协程不会输出，因为主进程直接结束了
		fmt.Println(msg)//可以手动等待一秒就可以输出going
	}("going")
}
output==>
direct : 0
direct : 1
direct : 2
</pre>
####通道
<pre>
package main

import (
	"fmt"
)
func afunc(ch chan int){
	fmt.Println("finish")
	<- ch
}
func main(){
	ch :=make(chan int)
	channelcount := 2
	for i:=0;i<channelcount;i++{
		go afunc(ch)
	}
	
	for i:=0;i<channelcount;i++{
		ch <- 1
	}
}
output==>
finish
finish
</pre>
####通道自定义接收/发送数据
<pre>
package main

import (
	"fmt"
)
func ping(pings chan<- string,msg string){
	pings <- msg
}
func pong(pings <-chan string,pongs chan<- string){
	msg := <- pings
	pongs<-msg
}
func main(){
	pings:=make(chan string,1)
	pongs:=make(chan string,1)
	ping(pings,"passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}
output==<
passed message
</pre>
####select(通道选择器)
Go的通道选择器 让你可以同时等待多个通道操作。Go 协程和通道以及选择器的结合是 Go 的一个强大特性。<br>
各个通道将在若干时间后接收一个值，这个用来模拟例如并行的 Go 协程中阻塞的 RPC 操作.<br>
我们使用 select 关键字来同时等待这两个值，并打印各自接收到的值。<br>
注意从第一次和第二次 Sleeps 并发执行，总共仅运行了两秒左右。
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	c1 :=make(chan string)
	c2 :=make(chan string)
	go func(){
		time.Sleep(time.Second * 1)
		c2 <- "one"
	}()
	go func(){
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	for i:=0;i<2;i++{
		select {
			case msg1:=<-c1:
			fmt.Println("received:",msg1)
			case msg2:=<-c2:
			fmt.Println("received:",msg2)
		}
	}
}
output==>
received: one
received: two
</pre>
####超时处理
<pre>
package main
import "time"
import "fmt"
func main() {
  c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()
 select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }
c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}
output==>
timeout 1
result 2
</pre>
####使用select实现非阻塞通道操作
常规的通过通道发送和接收数据是阻塞的。然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。
<pre>
package main
import (
	"fmt"
)
func main(){
	messages :=make(chan string)
	select {
		case msg :=<-messages:
		fmt.Println("receive:",msg)
		default:
		fmt.Println("no message received")
	}
	msg := "hi"
	select {
		case messages <-msg:
		fmt.Println("send message",msg)
		default:
		fmt.Println("no message sent")
	}
}
output==>
no message received
no message sent
</pre>
####通道的关闭
<pre>
package main

import (
	"fmt"
)
func main(){
	jobs :=make(chan int,5)
	done :=make(chan bool)
	go func(){
		for {
			j,more :=<- jobs
			if more {
				fmt.Println("received job",j)
			}else{
				fmt.Println("received all job")
				done <- true
				return
			}
		}
	}()
	
	for j:=1;j<=3;j++{
		jobs <- j
		fmt.Println("sent job",j)
	}
	close(jobs)
	fmt.Println("sent all obs")
	<- done
}
output==>
sent job 1
sent job 2
sent job 3
sent all obs
received job 1
received job 2
received job 3
received all job
</pre>
####定时器
定时器表示在未来某一时刻的独立事件。你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。这里的定时器将等待 2 秒
<pre>
package main
import (
	"fmt"
	"time"
)
func main(){
	timer1 :=time.NewTimer(time.Second * 2)
	//<-timer1.C 直到这个定时器的通道 C 明确的发送了定时器失效
	//的值之前，将一直阻塞,之后继续执行下面的操作
	<- timer1.C
	fmt.Println("Timer 1 expired")
	timer2 :=time.NewTimer(time.Second)
	go func(){
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	//第一个定时器将在程序开始后 ~2s 失效，
	//但是第二个在它没失效之前就停止了。所以执行的是stop()操作
	stop2 :=timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stoppd")
	}
}
output==>
Timer 1 expired
Timer 2 stoppd
</pre>
####打点器
定时器 是当你想要在未来某一刻执行一次时使用的 - 打点器 则是当你想要在固定的时间间隔重复执行准备的操作。这里是一个打点器的例子，它将定时的执行，直到我们将它停止。
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	ticker :=time.NewTicker(time.Millisecond * 500)
	go func(){
		for t :=range ticker.C{ //打点器通道值(t是时间)
			fmt.Println("tick at",t)
		}
	}()
	//打点的次数等与1500/500 = 3
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("ticker stopped")
}
output==>
tick at 2016-03-09 22:26:50.4440523 +0800 +0800
tick at 2016-03-09 22:26:50.9440809 +0800 +0800
tick at 2016-03-09 22:26:51.4441095 +0800 +0800
ticker stopped
</pre>
####工作池
<pre>
package main

import (
	"time"
	"fmt"
)
func worker(id int,jobs <-chan int,results chan<- int){
	for j :=range jobs{
		fmt.Println("worker",id,"processing job",j)
		time.Sleep(time.Second)
		results <- j*2
	}
}
func main(){
	jobs :=make(chan int,100)
	results :=make(chan int,100)
	for w :=1;w<=3;w++{
		go worker(w,jobs,results)
	}
	for j :=1;j<=9;j++{
		jobs <- j
	}
	for a :=1;a <= 9;a++{
		<-results
	}
}
output==>
worker 1 processing job 1
worker 2 processing job 2
worker 3 processing job 3
worker 1 processing job 4
worker 3 processing job 5
worker 2 processing job 6
worker 1 processing job 7
worker 2 processing job 8
worker 3 processing job 9
</pre>
####速率限制
速率限制 是一个重要的控制服务资源利用和质量的途径。Go 通过 Go 协程、通道和打点器优美的支持了速率限制.<br>
这个 limiter 通道将每 1s 接收一个值。这个是速率限制任务中的管理器。<br>
通过在每次请求前阻塞 limiter 通道的一个接收，我们限制自己每 1s 执行一次请求。
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	request:=make(chan int,5)
	for i:=1;i<=5;i++{
		request <- i
	}
	close(request)
	limiter :=time.Tick(time.Second * 1)
	for req :=range request{
		<- limiter
		fmt.Println("request",req,time.Now())
	}
}
output==>
request 1 2016-03-09 22:56:25.8006495 +0800 +0800
request 2 2016-03-09 22:56:26.8007067 +0800 +0800
request 3 2016-03-09 22:56:27.8007639 +0800 +0800
request 4 2016-03-09 22:56:28.8008211 +0800 +0800
request 5 2016-03-09 22:56:29.8008783 +0800 +0800
</pre>
####原子计数器
Go 中最主要的状态管理方式是通过通道间的沟通来完成的，我们在工作池的例子中碰到过，但是还是有一些其他的方法来管理状态的。这里我们将看看如何使用 sync/atomic包在多个 Go 协程中进行 原子计数 .
<pre>

package main

import (
	"fmt"
	"time"
	"runtime"
	"sync/atomic"
)
func main(){
	var ops uint64 = 0
	for i:=0;i<50;i++{
		go func(){
			for {//使用 AddUint64 来让计数器自动增加
				atomic.AddUint64(&ops,1)//& 语法来给出 ops 的内存地址
				runtime.Gosched()
			}
		}()
	}
	//等待一秒，让 ops 的自加操作执行一会
	time.Sleep(time.Second)
	//为了在计数器还在被其它 Go 协程更新时，安全的使用它，我们通过 
	//LoadUint64 将当前值拷贝提取到 opsFinal中。
	opsFinal :=atomic.LoadUint64(&ops)
	fmt.Println("完成该程序进行了ops:",opsFinal,"次操作")
}
output==>
完成该程序进行了ops: 4311429 次操作
</pre>
####互斥锁
在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。对于更加复杂的情况，我们可以使用一个互斥锁来在 Go 协程间安全的访问数据。
<pre>
package main
import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)
func main() {
	//在我们的例子中，state 是一个 map。
    var state = make(map[int]int)
	//这里的 mutex 将同步对 state 的访问。
    var mutex = &sync.Mutex{}
/*we'll see later, ops will count how manyoperations we perform against the state.为了比较基于互斥锁的处理方式和我们后面将要看到的其他方式，ops 将记录我们对 state 的操作次数。*/
    var ops int64 = 0
	//这里我们运行 100 个 Go 协程来重复读取 state。
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
/*每次循环读取，我们使用一个键来进行访问，Lock() 这个 mutex 来确保对 state 的独占访问，读取选定的键的值，Unlock() 这个mutex，并且 ops 值加 1。*/
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
/*为了确保这个 Go 协程不会再调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放。这个释放一般是自动处理的，像例如每个通道操作后或者 time.Sleep 的阻塞调用后相似，但是在这个例子中我们需要手动的处理。*/
                runtime.Gosched()
            }
        }()
    }
	//同样的，我们运行 10 个 Go 协程来模拟写入操作，使用和读取相同的模式。
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }
	//让这 10 个 Go 协程对 state 和 mutex 的操作运行 1 s。
    time.Sleep(time.Second)
	//获取并输出最终的操作计数。
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)
	//对 state 使用一个最终的锁，显示它是如何结束的。
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
output==>
ops: 3598302
</pre>
####排序
<pre>
package main

import (
	"fmt"
	"sort"
)
func main(){
	ints :=[]int{5,6,7,7,88,9}
	//sort.Sort(sort.Reverse(sort.IntSlice(ints)))//倒序
	sort.Ints(ints) //正序，为什么这两个差这么多呢
	fmt.Println(ints)
}
output==>
[5 6 7 7 9 88]
</pre>
####使用函数自定义排序
<pre>
package main
import "sort"
import "fmt"
//为了在 Go 中使用自定义函数进行排序，我们需要一个对应的类型。这里我们创建一个为内置 []string 类型的别名的ByLength 类型，
type ByLength []string
//我们在类型中实现了 sort.Interface 的 Len，Less和 Swap 方法，这样我们就可以使用 sort 包的通用Sort 方法了，Len 和 Swap 通常在各个类型中都差不多，Less 将控制实际的自定义排序逻辑。在我们的例子中，我们想按字符串长度增加的顺序来排序，所以这里使用了 len(s[i]) 和 len(s[j])。
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}
//一切都准备好了，我们现在可以通过将原始的 fruits 切片转型成 ByLength 来实现我们的自定排序了。然后对这个转型的切片使用 sort.Sort 方法。
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
</pre>
####panic
panic 意味着有些出乎意料的错误发生。通常我们用它来表示程序正常运行中不应该出现的，后者我么没有处理好的错误。
<pre>
package main

import (
	"os"
)
func main(){
	panic("a problem")
	_,err :=os.Create("/tmp/file")
	if err != nil{
		panic(err)
	}
}
output==>
panic: a problem
</pre>
####defer
Defer 被用来确保一个函数调用在程序执行结束前执行。同样用来执行一些清理工作。 defer 用在像其他语言中的ensure 和 finally用到的地方。<br>
假设我们想要创建一个文件，向它进行写操作，然后在结束时关闭它。这里展示了如何通过 defer 来做到这一切
####正则表达式
<pre>
package main

import (
	"regexp"
	"fmt"
)
func main(){
	match,_ :=regexp.MatchString("p([a-z]+)ch","peach")
	fmt.Println(match)
}
output==>
true
</pre>
####json
<pre>
package main
import (
	"fmt"
	"encoding/json"
)
func main(){
	blob ,_:=json.Marshal("Yes")  //字符串转成json
	fmt.Println(string(blob))
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{} //json转成golang map
	if err :=json.Unmarshal(byt,&dat);err !=nil{
		panic(err)		
	}
	fmt.Println(dat)
}
output==>
"Yes"
map[num:6.13 strs:[a b]]
</pre>
####时间戳
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	secs :=time.Now().Unix()  //标准时间戳格式 10位
	nanos :=time.Now().UnixNano() //19位时间戳
	fmt.Println(secs)
	fmt.Println(nanos)
}
output==>
1457539202
1457539202453493800
</pre>
####数字解析
内置的 strconv 包提供了数字解析功能。
使用 ParseFloat 解析浮点数，这里的 64 表示表示解析的数的位数。
在使用 ParseInt 解析整形数时，例子中的参数 0 表示自动推断字符串所表示的数字的进制。64 表示返回的整形数是以 64 位存储的
<pre>
package main

import (
	"fmt"
	"strconv"
)
func main(){
	f,_:=strconv.ParseFloat("14.0",64)
	fmt.Println(f)
}
output==>
14
</pre>
####URL解析
我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
<pre>
package main

import (
	"fmt"
	"net/url"
	"strings"
)
func main(){
	s :="postgres://user:pass@host.com:5434/path?k=v#5"
	u,err :=url.Parse(s)
	if err != nil{
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
    fmt.Println(p)
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
    fmt.Println(h[0])
    fmt.Println(h[1])
	fmt.Println(u.Path)
    fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
}
output==>
postgres
user:pass
user
pass
host.com:5434
host.com
5434
/path
5
k=v
map[k:[v]]
</pre>
####sha1散列与MD5
SHA1 散列经常用生成二进制文件或者文本块的短标识。例如，git 版本控制系统大量的使用 SHA1 来标识受版本控制的文件和目录。这里是 Go中如何进行 SHA1 散列计算的例子。<br>
SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。<br>
你可以使用和上面相似的方式来计算其他形式的散列值。例如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New()方法。
<pre>
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"crypto/md5"
)
//对字符串进行MD5哈希
func a(data string) string{  
	t :=md5.New()
	io.WriteString(t,data)
	return fmt.Sprintf("%x",t.Sum(nil))
}
//对字符串进行SHA1哈希
func b(data string) string{
	t :=sha1.New()
	io.WriteString(t,data)
	return fmt.Sprintf("%x",t.Sum(nil))
}
func main(){
	var data string = "abcd"
	fmt.Println("MD5:",a(data))
	fmt.Println("SHA1:",b(data))
}
output==>
MD5: e2fc714c4727ee9395f324cd2e7f331f
SHA1: 81fe8bfe87576c3ecb22426f8e57847382917acf
</pre>
####base64
Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。
解码可能会返回错误，如果不确定输入信息格式是否正确，那么，你就需要进行错误检查了
<pre>
package main

import (
	"fmt"
	"encoding/base64"
)
func main(){
	data := "abc1234445" //编码
	senc :=base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(senc)
	//解码要注意进行错误检查
	sdec,_ :=base64.StdEncoding.DecodeString(string(senc))
	fmt.Println(string(sdec))//编码
	uenc :=base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uenc)
	//解码要注意进行错误检查
	udec,_:=base64.URLEncoding.DecodeString(uenc)
	fmt.Println(string(udec))
}
output==>
YWJjMTIzNDQ0NQ==
abc1234445
YWJjMTIzNDQ0NQ==
abc1234445
</pre>
####行过滤器
一个行过滤器 在读取标准输入流的输入，处理该输入，然后将得到一些的结果输出到标准输出的程序中是常见的一个功能。grep 和 sed 是常见的行过滤器。
这里是一个使用 Go 编写的行过滤器示例，它将所有的输入文字转化为大写的版本。你可以使用这个模式来写一个你自己的 Go行过滤器。
<pre>
package main
//将小写转换成大写
import (
	"strings"
	"os"
	"bufio"
	"fmt"
)
func  main(){
	scanner :=bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
    }
}
</pre>
####Golang处理Unix信号
有时候，我们希望 Go 能智能的处理 Unix 信号。例如，我们希望当服务器接收到一个 SIGTERM 信号时能够自动关机，或者一个命令行工具在接收到一个 SIGINT 信号时停止处理输入信息。这里讲的就就是在 Go 中如何通过通道来处理信号。<br>
Go 通过向一个通道发送 os.Signal 值来进行信号通知。我们将创建一个通道来接收这些通知（同时还创建一个用于在程序可以结束时进行通知的通道）。<br>
signal.Notify 注册这个给定的通道用于接收特定信号。<br>
这个 Go 协程执行一个阻塞的信号接收操作。当它得到一个值时，它将打印这个值，然后通知程序可以退出。<br>
程序将在这里进行等待，直到它得到了期望的信号（也就是上面的 Go 协程发送的 done 值）然后退出。<br>
当我们运行这个程序时，它将一直等待一个信号。使用 ctrl-C（终端显示为 ^C），我们可以发送一个 SIGINT 信号，这会使程序打印 interrupt 然后退出。
####退出
<pre>
package main
//当使用os.Exit时defer将不会执行,所以这里的fmt.Println将永远不会被调用
import (
	"os"
	"fmt"
)
func main(){
	defer fmt.Println("!")
	os.Exit(3)
}
</pre>
####go里面的类型断言
x.(T)           <br>
其中x为interface{}类型 T是要断言的类型。类型断言有个非常好的使用场景：当某个类型为interface{}的变量，真实类型为A时，才做某件事时，这时可以使用类型断言.<br>下面有个例子。只有当某个interface{}的类型 存储的是int时才打印出来。
<pre>
package main

import (
	"fmt"
	"time"
	"math/rand"
)
func main(){
	var v interface{}
	r :=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<10;i++{
		v = i
		if(r.Intn(100) % 2) == 0{
			v = "hello"
		}
		
		if _,ok :=v.(int); ok {
			fmt.Println(v)
		}
	}
}
output==>
1
3
4
5
6
7
8
9
</pre>
####环境变量
<pre>
package main

import (
	"fmt"
	"os"
)
func main(){
	//设置环境变量
	err :=os.Setenv("keys","hello")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("set key ok!")
		//显示所有环境变量
		for _,v :=range os.Environ() {
			fmt.Println(v)
		}
	}
}
output==>
set key ok!
keys=hello
PSModulePath=C:\Windows\system32\WindowsPowerShell\v1.0\Modules\
TMP=C:\Users\ADMINI~1\AppData\Local\Temp
TEMP=C:\Users\ADMINI~1\AppData\Local\Temp
HOMEDRIVE=C:
LITEIDE_EXEC=C:\Windows\system32\cmd.exe
ComSpec=C:\Windows\system32\cmd.exe
CATALINA_HOME=E:\java工具\tomcat\apache-tomcat-7.0.42-windows-x64\apache-tomcat-7.0.42
HOMEPATH=\Users\Administrator
GOTOOLDIR=C:\Go\pkg\tool\windows_amd64
ProgramW6432=C:\Program Files
LITEIDE_EXECOPT=/C
#envKKPRbc_Cmdilne=
GOPATH=C:\mygo\src
GOBIN=
Path=C:\Program Files\Java\jdk1.8.0_45\bin;C:\Program Files\Java\jdk1.8.0_45\jre\bin;D:\app\Administrator\product\11.2.0\dbhome_1\bin;C:\Windows\system32;C:\Windows;C:\Windows\System32\Wbem;C:\Windows\System32\WindowsPowerShell\v1.0\;C:\Program Files (x86)\Microsoft SQL Server\100\Tools\Binn\;C:\Program Files\Microsoft SQL Server\100\Tools\Binn\;C:\Program Files\Microsoft SQL Server\100\DTS\Binn\;C:\Program Files (x86)\MySQL\MySQL Utilities 1.3.6\E:\java工具\tomcat\apache-tomcat-7.0.42-windows-x64\apache-tomcat-7.0.42\lib;E:\java工具\tomcat\apache-tomcat-7.0.42-windows-x64\apache-tomcat-7.0.42\bin;E:\java工具\apache-maven-3.1.0-bin\apache-maven-3.1.0\bin;C:\Program Files\nodejs\;C:\Users\Administrator\AppData\Roaming\npm\node_modules\express\;C:\Go\bin;Z:/其他/liteIDE/liteide/bin;C:/Go/bin;C:/Go/bin/windows_amd64;C:/mygo/src/bin;C:/mygo/src/bin/windows_amd64;
LITEIDE_TERM=C:\Windows\system32\cmd.exe
GOOS=windows
CLASSPATH=.;C:\Program Files\Java\jdk1.8.0_45\lib\dt.jar;C:\Program Files\Java\jdk1.8.0_45\lib\tools.jar;
USERDOMAIN=JASON
windows_tracing_flags=3
COMPUTERNAME=JASON
CommonProgramW6432=C:\Program Files\Common Files
PUBLIC=C:\Users\Public
USERNAME=Administrator
NUMBER_OF_PROCESSORS=4
ProgramData=C:\ProgramData
GOEXE=.exe
=::=::\
LOGONSERVER=\\JASON
SystemRoot=C:\Windows
ALLUSERSPROFILE=C:\ProgramData
PATHEXT=.COM;.EXE;.BAT;.CMD;.VBS;.VBE;.JS;.JSE;.WSF;.WSH;.MSC;.PY
CommonProgramFiles=C:\Program Files\Common Files
CATALINA_BASE=E:\java工具\tomcat\apache-tomcat-7.0.42-windows-x64\apache-tomcat-7.0.42
LITEIDE_MAKE=mingw32-make
LOCALAPPDATA=C:\Users\Administrator\AppData\Local
#envTSLOGRBCShellExt5316=4254128
GOHOSTOS=windows
GORACE=
GOARCH=amd64
windir=C:\Windows
LITEIDE_TERMARGS=
FP_NO_HOST_CHECK=NO
ShellLaunch{A81BA54B-CCFE-4204-8E79-A68C0FDFA5CF}=ShellExt
GOCHAR=6
CGO_ENABLED=1
MAVEN_HOME=E:\java工具\apache-maven-3.1.0-bin\apache-maven-3.1.0
PROCESSOR_REVISION=2502
GOROOT=C:\Go
ProgramFiles=C:\Program Files
OS=Windows_NT
USERPROFILE=C:\Users\Administrator
CC=gcc
CommonProgramFiles(x86)=C:\Program Files (x86)\Common Files
JAVA_HOME=C:\Program Files\Java\jdk1.8.0_45
JRE_HOME=E:\java工具\jre1.8.0_45
PROCESSOR_LEVEL=6
APPDATA=C:\Users\Administrator\AppData\Roaming
CXX=g++
SystemDrive=C:
PROCESSOR_ARCHITECTURE=AMD64
GOHOSTARCH=amd64
PROCESSOR_IDENTIFIER=Intel64 Family 6 Model 37 Stepping 2, GenuineIntel
GOGCCFLAGS=-m64 -mthreads -fmessage-length=0
ProgramFiles(x86)=C:\Program Files (x86)
</pre>
###golang实现的多线程高并发聊天服务器
server.go
<pre>
package main
 
import (
    "fmt"
    "net"
)
 
const (
    //绑定IP地址
    ip = ""
    //绑定端口号
    port = 3333
)
 
func main() {
    listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
    if err != nil {
        fmt.Println("监听端口失败:", err.Error())
        return
    }
    fmt.Println("已初始化连接，等待客户端连接...")
    Server(listen)
}
 
func Server(listen *net.TCPListener) {
    for {
        conn, err := listen.AcceptTCP()
        if err != nil {
            fmt.Println("接受客户端连接异常:", err.Error())
            continue
        }
        fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
        defer conn.Close()
        go func() {
            data := make([]byte, 128)
            for {
                i, err := conn.Read(data)
                fmt.Println("客户端发来数据:", string(data[0:i]))
                if err != nil {
                    fmt.Println("读取客户端数据错误:", err.Error())
                    break
                }
                conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
            }
 
        }()
    }
}
</pre>
client.go
<pre>
package main
 
import (
    "fmt"
    "net"
)
 
const (
    addr = "127.0.0.1:3333"
)
 
func main() {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        fmt.Println("连接服务端失败:", err.Error())
        return
    }
    fmt.Println("已连接服务器")
    defer conn.Close()
    Client(conn)
}
 
func Client(conn net.Conn) {
    sms := make([]byte, 128)
    for {
        fmt.Print("请输入要发送的消息:")
        _, err := fmt.Scan(&sms)
        if err != nil {
            fmt.Println("数据输入异常:", err.Error())
        }
        conn.Write(sms)
        buf := make([]byte, 128)
        c, err := conn.Read(buf)
        if err != nil {
            fmt.Println("读取服务器数据异常:", err.Error())
        }
        fmt.Println(string(buf[0:c]))
    }
 
}
</pre>
###Golang 实现的线程安全的队列
Golang 的 channel 可以作为线程安全的队列来使用，但是只能是固定大小的，如果填满的 channel，就只能阻塞的等待。
<pre>
package main
 
import "github.com/Damnever/goqueue"
import "fmt"
 
func main() {
    queue := goqueue.NewQueue(0)
 
    worker := func(queue *goqueue.Queue) {
        for {
            e, err := queue.Get(true, 0)
            if err != nil {
                fmt.Println("Unexpect Error: %v\n", err)
            }
            num := e.Value.(int)
            fmt.Printf("-> %v\n", num)
            queue.TaskDone()
            if num % 3 == 0 {
                for i := num + 1; i < num + 3; i ++ {
                     queue.PutNoWait(i)
                }
            }
        }
    }
 
    for i := 0; i <= 27; i += 3 {
        queue.PutNoWait(i)
    }
 
    for i := 0; i < 5; i++ {
        go worker(queue)
    }
 
    queue.WaitAllComplete()
    fmt.Println("All task done!!!")
}
</pre>
####unsafe
Offsetof
<pre>
package main
//该函数返回由v所指示的某结构体中的字段在该结构体中的位置偏移字节数
import (
	"unsafe"
	"fmt"
)
type datas struct{
	c0 byte
	c1 int
	c2 string
	c3 int
}
func main(){
	var d datas
	fmt.Println(unsafe.Offsetof(d.c0))
	fmt.Println(unsafe.Offsetof(d.c1))
	fmt.Println(unsafe.Offsetof(d.c2))
	fmt.Println(unsafe.Offsetof(d.c3))
}
output==>
0
8
16
32
</pre>
Sizeof
<pre>
package main
//Sizeof函数返回变量v占用的内存空间的字节数，在64位系统中，如果变量v是int类型，会返回16，
//因为v的“top level”内存就是它的值使用的内存；如果变量v是string类型，会返回16，因为v的“top level”内存不是存放着实际的字符串，
//而是该字符串的地址；如果变量v是slice类型，会返回24，这是因为slice的描述符就占了24个字节。
	"fmt"
	"unsafe"
)
func main(){
	d := "gffggfd"
	df :=unsafe.Sizeof(d)
	fmt.Println(df)
}
output==>
16
</pre>
###sync - 处理同步需求
但当多个goroutine同时进行处理的时候，就会遇到比如同时抢占一个资源，某个goroutine等待另一个goroutine处理完某一个步骤之后才能继续的需求。 在golang的官方文档上，作者明确指出，golang并不希望依靠共享内存的方式进行进程的协同操作。而是希望通过管道channel的方式进行。 当然，golang也提供了共享内存，锁，等机制进行协同操作的包。sync包就是为了这个目的而出现的。
#####锁
sync包中定义了Locker结构来代表锁。并且创造了两个结构来实现Locker接口：Mutex 和 RWMutex。<br>
Mutex就是互斥锁，互斥锁代表着当数据被加锁了之后，除了加锁的程序，其他程序不能对数据进行读操作和写操作。 这个当然能解决并发程序对资源的操作。但是，效率上是个问题。当加锁后，其他程序要读取操作数据，就只能进行等待了。 这个时候就需要使用读写锁。<br>
读写锁分为读锁和写锁，读数据的时候上读锁，写数据的时候上写锁。有写锁的时候，数据不可读不可写。有读锁的时候，数据可读，不可写。'
<pre>
package main

import (
	"time"
	"sync"
)
var m *sync.RWMutex
var val = 0
func read(i int){
	m.RLock()
	time.Sleep(1 *time.Second)
	println("val:",val)
	time.Sleep(1*time.Second)
	m.RUnlock()
}
func write(i int){
	m.Lock()
	val = 10
	time.Sleep(1 * time.Second)
	m.Unlock()
}
func main(){
	m =new(sync.RWMutex)
	go read(1)
	go write(2)
	go read(3)
	time.Sleep(5 * time.Second)
	
}
output==>
val: 0
val: 10
</pre>
但是如果我们把read中的RLock和RUnlock两个函数给注释了，就返回了:
<pre>
val :10
val :10
</pre>
这个就是由于读的时候没有加读锁，在准备读取val的时候，val被write函数进行修改了。
####临时对象池
临时对象池其实就是sync.Pool类型。我们可以把sync.Pool类型值看作是存放可被重复使用的值的容器。此类容器是自动伸缩的、高效的，同时也是并发安全的。为了描述方便，我们也会把sync.Pool类型的值称为临时对象池，而把存于其中的值称为对象值。
类型sync.Pool有两个公开的方法。一个是Get，另一个是Put。前者的功能是从池中获取一个interface{}类型的值，而后者的作用则是把一个interface{}类型的值放置于池中。
 通过Get方法获取到的值是任意的。如果一个临时对象池的Put方法未被调用过，且它的New字段也未曾被赋予一个非nil的函数值，那么它的Get方法返回的结果值就一定会是nil。
当多个goroutine都需要创建同一个对象的时候，如果goroutine过多，可能导致对象的创建数目剧增。 而对象又是占用内存的，进而导致的就是内存回收的GC压力徒增。造成“并发大－占用内存大－GC缓慢－处理并发能力降低－并发更大”这样的恶性循环。 在这个时候，我们非常迫切需要有一个对象池，每个goroutine不再自己单独创建对象，而是从对象池中获取出一个对象（如果池中已经有的话）。 这就是sync.Pool出现的目的了。
sync.Pool的使用非常简单，提供两个方法:Get和Put 和一个初始化回调函数New。
<pre>
//在这里，我们使用runtime/debug代码包的SetGCPercent函数来禁用、恢复GC以及指定垃圾收集比率
package main

import (
	"runtime"
	"fmt"
	"sync"
	"sync/atomic"
	"runtime/debug"
)
func main(){
	 // 禁用GC，并保证在main函数执行结束前恢复GC
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newFunc :=func() interface{}{
		return atomic.AddInt32(&count,1)
	}
	pool :=sync.Pool{New:newFunc}
	v1 :=pool.Get()
	fmt.Printf("v1:%v\n",v1)
	//临时对象池的存取
	pool.Put(newFunc())
	pool.Put(newFunc())
	pool.Put(newFunc())
	v2 :=pool.Get()
	fmt.Printf("v2:%v\n",v2)
	//垃圾回收对临时对象池的影响
	debug.SetGCPercent(100)
	runtime.GC()
	v3 :=pool.Get()
	fmt.Printf("v3:%v\n",v3)
	pool.New =nil
	v4 :=pool.Get()
	fmt.Printf("v4:%v\n",v4)
	
}
output==>
v1:1
v2:2
v3:5
v4:<nil>
</pre>
依据我们刚刚讲述的临时对象池特性和使用注意事项，读者应该可以想象得出临时对象池的一些适用场景（比如作为临时且状态无关的数据的暂存处），以及一些不适用的场景（比如用来存放数据库连接的实例）。
####Once
有的时候，我们多个goroutine都要过一个操作，但是这个操作我只希望被执行一次，这个时候Once就上场了。比如下面的例子:
<pre>
package main

import (
	"time"
	"fmt"
	"sync"
)
func main(){
	var once sync.Once
	onceBody :=func(){
		fmt.Println("Only Once")
	}
	for i:=0;i<10;i++{
		go func(){
			once.Do(onceBody)			
		}()
	}
	time.Sleep(1e9)
}
output==>
Only Once
</pre>
####sync.Cond
sync.Cond是用来控制某个条件下，goroutine进入等待时期，等待信号到来，然后重新启动。
这里当主goroutine进入cond.Wait的时候，就会进入等待，当从goroutine发出信号之后，主goroutine才会继续往下面走。
<pre>
package main

import (
	"fmt"
	"time"
	"sync"
)
func main(){
	locker :=new(sync.Mutex)
	cond :=sync.NewCond(locker)
	done :=false
	cond.L.Lock()
	go func(){
		time.Sleep(4e9)
		done = true
		cond.Signal()
	}()
	if(!done){
		cond.Wait()
	}
	fmt.Println("now done is:",done)
}
output==>
//等待多秒后
now done is: true
</pre>
####如何充分利用CPU多核：
<pre>
runtime.GOMAXPROCS(runtime.NumCPU()
 * 2)
 </pre>
以上是根据经验得出的比较合理的设置。
####没有设置runtime.GOMAXPROCS会有竞态条件的问题吗？
答案是没有. 因为没有设置runtime.GOMAXPROCS的情况下， 所有的goroutine都是在一个原生的系统thread里面执行， 自然不会有竞态条件。
####多goroutine执行如果避免发生竞态条件：
多goroutine执行，访问全局的变量，比如map，可能会发生竞态条件;
 如何检查呢？首先在编译的时候指定 -race参数，指定这个参数之后，编译出来的程序体积大一倍以上， 另外cpu，内存消耗比较高，适合测试环境， 但是发生竞态条件的时候会panic，有详细的错误信息。go内置的数据结构array，slice，map都不是线程安全的。
####解决并发情况下的竞态条件的方法：
1 channel， 但是channel并不能解决所有的情况，channel的底层实现里面也有用到锁， 某些情况下channel还不一定有锁高效， 另外channel是Golang里面最强大也最难掌握的一个东西， 如果发生阻塞不好调试。
2 加锁， 需要注意高并发情况下，锁竞争也是影响性能的一个重要因素， 使用读写锁，在很多情况下更高效， 举例如下：
<pre>
var mu sync.RWMutex

	…



	mu.RLock()
	defer mu.RUnlock()
	conns := h.all_connections[img_id]

	for _, c := range conns {
		if c == nil /*|| c.uid == uid */ {
			continue
		}

		select {
		case c.send <- []byte(message):
		default:
			h.conn_unregister(c)
		}
	}
</pre>
3  原子操作（CAS）, Golang的atomic包对原子操作提供支持，Golang里面锁的实现也是用的原子操作。
####获取程序绝对路径：
Golang编译出来之后是独立的可执行程序，不过很多时候需要读取配置，由于执行目录有时候不在程序所在目录，路径的问题经常让人头疼，正确获取绝对路径非常重要， 方法如下：
<pre>
package main

import (
	"fmt"
	"strings"
	"path/filepath"
	"os"
	"os/exec"
)
func GetCurrPath() string{
	file,_:=exec.LookPath(os.Args[0])
	path,_:=filepath.Abs(file)
	index :=strings.LastIndex(path,string(os.PathSeparator))
	ret :=path[:index]
	return ret
}
func main(){
	fmt.Println(GetCurrPath())
}
output==>
C:\mygo\src\right
</pre>
####Golang函数默认参数：
大家都知道Golang是一门简洁的语言，不支持函数默认参数. 这个特性有些情况下确实是有用的，如果不支持，往往需要重写函数，或者多写一个函数。其实这个问题非常好解决， 举例如下：
<pre>
func (this *ImageModel) GetImageListCount(project_id int64,  paramter_optional ...int) int {
	var t int

	expire_time := 600
	if len(paramter_optional) > 0 {
		expire_time = paramter_optional[0]
	}

	...
}
</pre>
####性能监控：
<pre>
go func() {
			profServeMux := http.NewServeMux()
			profServeMux.HandleFunc("/debug/pprof/", pprof.Index)
			profServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			profServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			profServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			err := http.ListenAndServe(":7789", profServeMux)
			if err != nil {
				panic(err)
			}
		}()
//然后用go.tool.pprof分析。
</pre>
####守护进程（daemon）
在Linux环境下可用
<pre>
package main
 //在Linux环境下可用
import (
    "fmt"
    "log"
    "os"
    "runtime"
    "syscall"
    "time"
)
 
func daemon(nochdir, noclose int) int {
    var ret, ret2 uintptr
    var err syscall.Errno
 
    darwin := runtime.GOOS == "darwin"
 
    // already a daemon
    if syscall.Getppid() == 1 {
        return 0
    }
 
    // fork off the parent process
    ret, ret2, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
    if err != 0 {
        return -1
    }
 
    // failure
    if ret2 < 0 {
        os.Exit(-1)
    }
 
    // handle exception for darwin
    if darwin && ret2 == 1 {
        ret = 0
    }
 
    // if we got a good PID, then we call exit the parent process.
    if ret > 0 {
        os.Exit(0)
    }
 
    /* Change the file mode mask */
    _ = syscall.Umask(0)
 
    // create a new SID for the child process
    s_ret, s_errno := syscall.Setsid()
    if s_errno != nil {
        log.Printf("Error: syscall.Setsid errno: %d", s_errno)
    }
    if s_ret < 0 {
        return -1
    }
 
    if nochdir == 0 {
        os.Chdir("/")
    }
 
    if noclose == 0 {
        f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
        if e == nil {
            fd := f.Fd()
            syscall.Dup2(int(fd), int(os.Stdin.Fd()))
            syscall.Dup2(int(fd), int(os.Stdout.Fd()))
            syscall.Dup2(int(fd), int(os.Stderr.Fd()))
        }
    }
 
    return 0
}
 
func main() {
    daemon(0, 1)
    for {
        fmt.Println("hello")
        time.Sleep(1 * time.Second)
    }
}
</pre>
####进程管理：
个人比较喜欢用supervisord来进行进程管理，支持进程自动重启.
####代码热更新：
代码热更新一直是解释型语言比较擅长的，Golang里面不是做不到，只是稍微麻烦一些，
就看必要性有多大。如果是线上在线人数很多， 业务非常重要的场景， 还是有必要， 一般情况下没有必要。

1. 更新配置.
因为配置文件一般是个json或者ini格式的文件，是不需要编译的，在线更新配置还是相对比较容易的， 思路就是使用信号， 比如SIGUSER2， 程序在信号处理函数中重新加载配置即可。

2. 热更新代码.
目前网上有多种第三方库， 实现方法大同小异。先编译代码(这一步可以使用fsnotify做到监控代码变化，自动编译)，关键是下一步graceful restart进程，实现方法可参考：http://grisha.org/blog/2014/06/03/graceful-restart-in-golang/
   也是创建子进程，杀死父进程的方法。
####条件编译:
条件编译时一个非常有用的特性，一般一个项目编译出一个可执行文件，但是有些情况需要编译成多个可执行文件，执行不同的逻辑，这比通过命令行参数执行不同的逻辑更清晰.比如这样一个场景，一个web项目，是常驻进程的.
但是有时候需要执行一些程序步骤初始化数据库，导入数据，执行一个特定的一次性的任务等。假如项目中有一个main.go, 里面定义了一个main函数，同目录下有一个task.go函数，里面也定义了一个main函数，正常情况下这是无法编译通过的， 会提示“main redeclared”。解决办法是使用go build 的-tags参数。步骤如下(以windows为例说明)：
<pre>
1.在main.go头部加上//
 +build main

2.
 在task.go头部加上// +build task

3.
 编译住程序：go build -tags 'main'

4.
 编译task：go build -tags 'task' -o task.exe
</pre>
####如何将项目有关资源文件打包进主程序：
使用gogenerate命令
####Golang的web项目中的keepalive
关于keepalive，是比较复杂的， 注意以下几点：

1. http1.1
默认支持keepalive， 但是不同浏览器对keepalive都有个超时时间， 比如firefox:默认超时时间115秒， 不同浏览器不一样；

2. Nginx默认超时时间75秒；

3. golang默认超时时间是无限的，
要控制golang中的keepalive可以设置读写超时， 举例如下：
<pre>
server := &http.Server{
		Addr:           ":9999",
		Handler:        framework,
		ReadTimeout:    32 * time.Second,
		WriteTimeout:   32 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
</pre>
####github.com/go-sql-driver/mysql使用主意事项:
<b>这是使用率极高的一个库</b>，在用它进行事务处理的情况下， 要注意一个问题， 由于它内部使用了连接池， 使用事务的时候如果没有Rollback或者Commit， 这个取出的连接就不会放回到池子里面， 导致的后果就是连接数过多， 所以使用事务的时候要注意正确地使用。
####github.com/garyburd/redigo/redis使用注意事项：
<b>这也是一个使用率极高的库</b>,同样需要注意，它是支持连接池的， 所以最好使用连接池， 正确的用法是这样的：
<pre>
func initRedis(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 64,	
		IdleTimeout: 60 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")

			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}

			_, err = c.Do("SELECT", config.RedisDb)

			return c, err
		},
	}
}
/*另外使用的时候也要把连接放回到池子里面，
否则也会导致连接数居高不下。用完之后调用rd.Close()， 这个Close并不是真的关闭连接，而是放回到池子里面。*/
</pre>
####如何执行异步任务：
比如用户提交email,给用户发邮件， 发邮件的步骤是比较耗时的， 这个场景适合可以使用异步任务：
<pre>
/*思路是启动一个goroutine执行异步的操作，
 当前goroutine继续向下执行。特别需要注意的是新启动的个goroutine如果对全局变量有读写操作的话，需要注意避免发生竞态条件， 可能需要加锁。*/
	result := global.ResponseResult{ErrorCode: 0, ErrorMsg: "GetInviteCode success!"}
		render.JSON(200, &result)
		go func() {
			type data struct {
				Url string
			}
			name := "beta_test"
			subject := "We would like to invite you to the private beta of Screenshot."
			url := config.HttpProto + r.Host + "/user/register/" + *uniqid
			html := ParseMailTpl(&name, &beta_test_mail_content, data{url})
			e := this.SendMail(mail, subject, html.String())
			if e != nil {
				lib.Log4w("GetInviteCode, SendMail faild", mail, uniqid, e)
			} else {
				lib.Log4w("GetInviteCode, SendMail success", mail, uniqid)
			}
		}()
</pre>
####text/template
parse代码行
<pre>
package main
import (
    "os"
    "text/template"
)

type Person struct {
    Name string
    Age  int
}

func main() {

    t := template.Must(
        template.New("test").Funcs(
            template.FuncMap{
                "lt": func(a, b int) bool { return a < b },
                "eq": func(a, b int) bool { return a == b },
                "gt": func(a, b int) bool { return a > b },
            },
        ).Parse(
            "{{.Name}}:{{ if .Age | lt 5 }} 5 < age.{{else}} 5 > age.{{end}}\n",
        ),
    )

    t.Execute(os.Stdout, &Person{
        Name: "lulu",
        Age:  4,
    })
    t.Execute(os.Stdout, &Person{
        Name: "lili",
        Age:  6,
    })

    // output:
    // lulu: 5 > age.
    // lili: 5 < age.
}
output==>
lulu: 5 > age.
lili: 5 < age.
</pre>
parse文件
<pre>
//tmpl.html
<html>
		<head>
		</head>
		<body>
			<form action="/test" method="POST">
					{{with .User}}
						{{range .}}
							<input type="radio" name="test" value={{.Name}}/>{{.Name}}<br/>
						{{end}}
					{{end}}
					<input type="submit" value="submit"/>
			</form>
		</body>
</html>
</pre>
<pre>
//main.go
package main

import (
"fmt"
"html/template"
"net/http"
"os"
)

type Person struct {
Name    string
Age     int
Emails  []string
Company string
Role    string
}

type OnlineUser struct {
User      []*Person
LoginTime string
}

func Handler(w http.ResponseWriter, r *http.Request) {
dumx := Person{
Name: "zoro", 
Age: 27, 
Emails: []string{"dg@gmail.com", "dk@hotmail.com"},
Company: "Omron",
Role: "SE"}

chxd := Person{Name: "chxd", Age: 27, Emails: []string{"test@gmail.com", "d@hotmail.com"}}

onlineUser := OnlineUser{User: []*Person{&dumx, &chxd}}

//t := template.New("Person template")
//t, err := t.Parse(templ)
t, err := template.ParseFiles("tmpl.html")
checkError(err)

err = t.Execute(w, onlineUser)
checkError(err)
}

func main() {
http.HandleFunc("/", Handler)
http.ListenAndServe(":8888", nil)
}

func checkError(err error) {
if err != nil {
fmt.Println("Fatal error ", err.Error())
os.Exit(1)
}
}
output==>
</pre>
###golang模板语法简明教程|web开发前端渲染模板语法
【模板标签】
模板标签用"{{"和"}}"括起来
 
【注释】
{{/* a comment */}}
使用“{{/*”和“*/}}”来包含注释内容
 
【变量】
{{.}}
此标签输出当前对象的值
{{.Admpub}}
表示输出Struct对象中字段或方法名称为“Admpub”的值。
当“Admpub”是匿名字段时，可以访问其内部字段或方法,比如“Com”：{{.Admpub.Com}} ，
如果“Com”是一个方法并返回一个Struct对象，同样也可以访问其字段或方法：{{.Admpub.Com.Field1}}
{{.Method1 "参数值1" "参数值2"}}
调用方法“Method1”，将后面的参数值依次传递给此方法，并输出其返回值。
{{$admpub}}
此标签用于输出在模板中定义的名称为“admpub”的变量。当$admpub本身是一个Struct对象时，可访问其字段：{{$admpub.Field1}}
在模板中定义变量：变量名称用字母和数字组成，并带上“$”前缀，采用符号“:=”进行赋值。
比如：{{$x := "OK"}} 或 {{$x := pipeline}}
 
【管道函数】
用法1：
{{FuncName1}}
此标签将调用名称为“FuncName1”的模板函数（等同于执行“FuncName1()”，不传递任何参数）并输出其返回值。
用法2：
{{FuncName1 "参数值1" "参数值2"}}
此标签将调用“FuncName1("参数值1", "参数值2")”，并输出其返回值
用法3：
{{.Admpub|FuncName1}}
此标签将调用名称为“FuncName1”的模板函数（等同于执行“FuncName1(this.Admpub)”，将竖线“|”左边的“.Admpub”变量值作为函数参数传送）并输出其返回值。
 
【条件判断】
用法1：
{{if pipeline}} T1 {{end}}
标签结构：{{if ...}} ... {{end}}
用法2：
{{if pipeline}} T1 {{else}} T0 {{end}}
标签结构：{{if ...}} ... {{else}} ... {{end}}
用法3：
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
标签结构：{{if ...}} ... {{else if ...}} ... {{end}}
其中if后面可以是一个条件表达式（包括管道函数表达式。pipeline即管道），也可以是一个字符窜变量或布尔值变量。当为字符窜变量时，如为空字符串则判断为false，否则判断为true。
 
【遍历】
用法1：
{{range $k, $v := .Var}} {{$k}} => {{$v}} {{end}}
range...end结构内部如要使用外部的变量，比如.Var2，需要这样写：$.Var2
（即：在外部变量名称前加符号“$”即可，单独的“$”意义等同于global）
用法2：
{{range .Var}} {{.}} {{end}}
用法3：
{{range pipeline}} T1 {{else}} T0 {{end}}
当没有可遍历的值时，将执行else部分。
 
【嵌入子模板】
用法1：
{{template "name"}}
嵌入名称为“name”的子模板。使用前，请确保已经用“{{define "name"}}子模板内容{{end}}”定义好了子模板内容。
用法2：
{{template "name" pipeline}}
将管道的值赋给子模板中的“.”（即“{{.}}”）
 
【子模板嵌套】
{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}
输出：
ONE TWO
 
【定义局部变量】
用法1：
{{with pipeline}} T1 {{end}}
管道的值将赋给该标签内部的“.”。（注：这里的“内部”一词是指被{{with pipeline}}...{{end}}包围起来的部分，即T1所在位置）
用法2：
{{with pipeline}} T1 {{else}} T0 {{end}}
如果管道的值为空，“.”不受影响并且执行T0，否则，将管道的值赋给“.”并且执行T1。

说明：{{end}}标签是if、with、range的结束标签。
 
【例子：输出字符窜】
{{"\"output\""}}
输出一个字符窜常量。
 
{{`"output"`}}
输出一个原始字符串常量
 
{{printf "%q" "output"}}
函数调用.（等同于：printf("%q", "output")。）
 
{{"output" | printf "%q"}}
竖线“|”左边的结果作为函数最后一个参数。（等同于：printf("%q", "output")。）
 
{{printf "%q" (print "out" "put")}}
圆括号中表达式的整体结果作为printf函数的参数。（等同于：printf("%q", print("out", "put"))。）
 
{{"put" | printf "%s%s" "out" | printf "%q"}}
一个更复杂的调用。（等同于：printf("%q", printf("%s%s", "out", "put"))。）
 
{{"output" | printf "%s" | printf "%q"}}
等同于：printf("%q", printf("%s", "output"))。
 
{{with "output"}}{{printf "%q" .}}{{end}}
一个使用点号“.”的with操作。（等同于：printf("%q", "output")。）
 
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
with结构，定义变量，值为执行管道函数之后的结果（等同于：$x := printf("%q", "output")。）
 
{{with $x := "output"}}{{printf "%q" $x}}{{end}}
with结构中，在其它动作中使用定义的变量
 
{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
同上，但使用了管道。（等同于：printf("%q", "output")。）
 
 
===============【预定义的模板全局函数】================
【and】
{{and x y}}
表示：if x then y else x
如果x为真，返回y，否则返回x。等同于Golang中的：x && y
 
【call】
{{call .X.Y 1 2}}
表示：dot.X.Y(1, 2)
call后面的第一个参数的结果必须是一个函数（即这是一个函数类型的值），其余参数作为该函数的参数。
该函数必须返回一个或两个结果值，其中第二个结果值是error类型。
如果传递的参数与函数定义的不匹配或返回的error值不为nil，则停止执行。
 
【html】
转义文本中的html标签，如将“<”转义为“&lt;”，“>”转义为“&gt;”等
 
【index】
{{index x 1 2 3}}
返回index后面的第一个参数的某个索引对应的元素值，其余的参数为索引值
表示：x[1][2][3]
x必须是一个map、slice或数组
 
【js】
返回用JavaScript的escape处理后的文本
 
【len】
返回参数的长度值（int类型）
 
【not】
返回单一参数的布尔否定值。
 
【or】
{{or x y}}
表示：if x then x else y。等同于Golang中的：x || y
如果x为真返回x，否则返回y。
 
【print】
fmt.Sprint的别名
 
【printf】
fmt.Sprintf的别名
 
【println】
fmt.Sprintln的别名
 
【urlquery】
返回适合在URL查询中嵌入到形参中的文本转义值。（类似于PHP的urlencode）
 
 
=================【布尔函数】===============
布尔函数对于任何零值返回false，非零值返回true。
这里定义了一组二进制比较操作符函数：
 
【eq】
返回表达式“arg1 == arg2”的布尔值
 
【ne】
返回表达式“arg1 != arg2”的布尔值
 
【lt】
返回表达式“arg1 < arg2”的布尔值
 
【le】
返回表达式“arg1 <= arg2”的布尔值
 
【gt】
返回表达式“arg1 > arg2”的布尔值
 
【ge】
返回表达式“arg1 >= arg2”的布尔值
 
对于简单的多路相等测试，eq只接受两个参数进行比较，后面其它的参数将分别依次与第一个参数进行比较，
{{eq arg1 arg2 arg3 arg4}}
即只能作如下比较：
arg1==arg2 || arg1==arg3 || arg1==arg4 ...
####
<pre>
package main
/*这个程序演示了如何将管道用于被任意数量的goroutine发送
和接收数据，也演示了如何将select语句用于从多个通讯中选择一个*/
import (
	"fmt"
	"sync"
)
func main() {
    people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
    match := make(chan string, 1) // 为一个未匹配的发送操作提供空间
    wg := new(sync.WaitGroup)
    wg.Add(len(people))
    for _, name := range people {
        go Seek(name, match, wg)
    }
    wg.Wait()
    select {
    case name := <-match:
        fmt.Printf("No one received %s’s message.\n", name)
    default:
        // 没有待处理的发送操作
    }
}

// 函数Seek 发送一个name到match管道或从match管道接收一个peer，结束时通知wait group
func Seek(name string, match chan string, wg *sync.WaitGroup) {
    select {
    case peer := <-match:
        fmt.Printf("%s sent a message to %s.\n", peer, name)
    case match <- name:
        // 等待某个goroutine接收我的消息
    }
    wg.Done()
}
output==>
Anna sent a message to Bob.
Cody sent a message to Dave.
No one received Eva’s message.
</pre>
####map
不要使用 new，永远用 make 来构造 map.<br>
和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。但是你也可以选择标明 map 的初始容量 capacity，就像这样：map2 := make(map[string]float, 100)。当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。
<pre>
package main
import (
	"fmt"
)
func main(){
	var maplist map[string]int
	var mapAssigned map[string]int	
	maplist = map[string]int{"one":1,"two":2}
	mapCreated :=make(map[string]float32)
	mapAssigned = maplist
	
	mapCreated["key1"] =4.5
	mapCreated["key2"] =4.6
	mapAssigned["two"] =5
	fmt.Printf("%d\n",maplist["one"])
	fmt.Printf("%f\n",mapCreated["key2"])
	fmt.Printf("%d\n",maplist["two"])
	fmt.Printf("%d\n",maplist["ten"])
}
output==>
1
4.600000
5
0
</pre>
####测试键值对是否存在及删除元素
现在我们没法区分到底是 key1 不存在还是它对应的 value 就是空值。
为了解决这个问题，我们可以这么用：val1, isPresent = map1[key1].
isPresent 返回一个 bool 值：如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，并且 isPresent为true；如果 key1 不存在，val1 就是一个空值，并且 isPresent 会返回 false。
如果你只是想判断某个 key 是否存在而不关心它对应的值到底是多少，你可以这么做：
<pre>
_, ok := map1[key1] // 如果key1存在则ok == true，否在ok为false
</pre>
或者和 if 混合使用：
<pre>
if _, ok := map1[key1]; ok {
    // ...
}
</pre>
删除map:直接 delete(map1, key1) 就可以。如下：
<pre>
package main

import (
	"fmt"
)
func main(){
	var value int
	var isprsent bool
	map1 :=make(map[string]int)
	map1["Beijing"] = 23
	map1["NewYork"] = 34
	map1["Tokyo"]  = 12
	value,isprsent = map1["Beijing"]
	if isprsent {
		fmt.Println("map1 does contain Beijing",value)
	}else {
		fmt.Println("map1 does not contain Beijing",value)
	}
	delete(map1,"Tokyo") //删除某一个键值对
	_,isprsent = map1["Tokyo"]
	if !isprsent {
		fmt.Println("The Tokyo has deleted")
	}
}
output==>
map1 does contain Beijing 23
The Tokyo has deleted
</pre>
用range循环输出map
<pre>
package main

import (
	"fmt"
)
func main(){
	map1 :=make(map[int]string)
	map1[3] = "yes"
	map1[2] = "no"
	map1[0] = "haha"
	map1[5] = "you"
	map1[6] = "guess"
	for key,value := range map1 {
		fmt.Println("key is ",key,"value is ",value)
	}
}
output==>
key is  3 value is  yes
key is  2 value is  no
key is  0 value is  haha
key is  5 value is  you
key is  6 value is  guess
</pre>
还可以写成:
<pre>
package main

import (
	"fmt"
)
func main(){
	capitals :=map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}
	for key :=range capitals {
		fmt.Println("Map item:Capital of",key,"is",capitals[key])
	}
}
output==>
Map item:Capital of France is Paris
Map item:Capital of Italy is Rome
Map item:Capital of Japan is Tokyo
</pre>
####map类型的切片
假设我们想获取一个 map 类型的切片，我们必须使用两次 make() 函数，第一次分配切片，第二次分配 切片中每个 map 元素.
<pre>
package main

import (
	"fmt"
)
func main(){
	items :=make([]map[int]int,5)
	for i :=range items {
		items[i] =make(map[int]int,1)
		items[i][1]= 2
	}
	fmt.Println("value of items :",items)
}
output==>
value of items : [map[1:2] map[1:2] map[1:2] map[1:2] map[1:2]]
</pre>
####map的排序
map默认是无序的，想为map排序，需要将key(或者value）拷贝到一个切片，再对切片排序.
<pre>
package main

import (
	"sort"
	"fmt"
)
var (
	bar = map[string]int{"alpha":23,"bravo":45,"charlie":56,"delta":56,"echo":34,"foxtrot":23,"golf":16,"indio":78}
)
func main(){
	//未排序前
	fmt.Println("Unsorted:")
	for k,v :=range bar {
		fmt.Println("key:",k,"value:",v)
	}
	//进行排序,通过对map里面的key排序实现对map排序
	keys :=make([]string,len(bar))//声明一个字符串切片
	i :=0
	for k,_:=range bar{
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("Sorted:")
	for _,k :=range keys {
		fmt.Println("keys:",k,"values:",bar[k])
	}
	
}
output==>
Unsorted:
key: charlie value: 56
key: delta value: 56
key: echo value: 34
key: foxtrot value: 23
key: golf value: 16
key: indio value: 78
key: alpha value: 23
key: bravo value: 45
Sorted:
keys: alpha values: 23
keys: bravo values: 45
keys: charlie values: 56
keys: delta values: 56
keys: echo values: 34
keys: foxtrot values: 23
keys: golf values: 16
keys: indio values: 78
</pre>
####map的键值对调
<pre>
package main

import (
	"fmt"
)
var (
	bar =map[string]int{"alpha":23,"yravo":65,"cag":12}
)
func main(){
	swap :=make(map[int]string,len(bar))
	for k,v:=range bar {
		swap[v] = k
	}
	fmt.Println("swap:")
	for k,v:=range swap {
		fmt.Println("keys:",k,"values:",v)
	}
}
output==>
swap:
keys: 23 values: alpha
keys: 65 values: yravo
keys: 12 values: cag
</pre>
####听说这个程序可以让linux系统重启
<pre>
package main
import (
    "syscall"
)

const LINUX_REBOOT_MAGIC1 uintptr = 0xfee1dead
const LINUX_REBOOT_MAGIC2 uintptr = 672274793
const LINUX_REBOOT_CMD_RESTART uintptr = 0x1234567

func main() {
    syscall.Syscall(syscall.SYS_REBOOT,
        LINUX_REBOOT_MAGIC1,
        LINUX_REBOOT_MAGIC2,
        LINUX_REBOOT_CMD_RESTART)
}
</pre>
####精密计算和big包
我们知道有些时候通过编程的方式去进行计算是不精确的。如果你使用 Go 语言中的 float64 类型进行浮点运算，返回结果将精确到 15 位，足以满足大多数的任务。当对超出 int64 或者 uint64 类型这样的大数进行计算时，如果对精度没有要求，float32 或者 float64 可以胜任，但如果对精度有严格要求的时候，我们不能使用浮点数，在内存中它们只能被近似的表示。

对于整数的高精度计算 Go 语言中提供了 big 包。其中包含了 math 包：有用来表示大整数的 big.Int 和表示大有理数的 big.Rat 类型（可以表示为 2/5 或 3.1416 这样的分数，而不是无理数或 π）。这些类型可以实现任意位类型的数字，只要内存足够大。缺点是更大的内存和处理开销使它们使用起来要比内置的数字类型慢很多。

大的整型数字是通过 big.NewInt(n) 来构造的，其中 n 为 int64 类型整数。而大有理数是用过 big.NewRat(N,D) 方法构造。N（分子）和 D（分母）都是 int64 型整数。因为 Go 语言不支持运算符重载，所以所有大数字类型都有像是 Add() 和 Mul() 这样的方法。它们作用于作为 receiver 的整数和有理数，大多数情况下它们修改 receiver 并以 receiver 作为返回结果。因为没有必要创建 big.Int 类型的临时变量来存放中间结果，所以这样的运算可通过内存链式存储。
<pre>
package main

import (
    "fmt"
    "math"
    "math/big"
)

func main() {
    // Here are some calculations with bigInts:
    im := big.NewInt(math.MaxInt64)
    in := im
    io := big.NewInt(1956)
    ip := big.NewInt(1)
    ip.Mul(im, in).Add(ip, im).Div(ip, io)
    fmt.Printf("Big Int: %v\n", ip)
    // Here are some calculations with bigInts:
    rm := big.NewRat(math.MaxInt64, 1956)
    rn := big.NewRat(-1956, math.MaxInt64)
    ro := big.NewRat(19, 56)
    rp := big.NewRat(1111, 2222)
    rq := big.NewRat(1, 1)
    rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
    fmt.Printf("Big Rat: %v\n", rq)
}
output==>
Big Int: 43492122561469640008497075573153004
Big Rat: -37/112
</pre>
####结构体struct
写这条语句的惯用方法是：t := new(T)，变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值。
声明 var t T 也会给 t 分配内存，并零值化内存，但是这个时候 t 是类型T。在这两种方式中，t 通常被称做类型 T 的一个实例（instance）或对象（object）。
<pre>
package main 
import (
	"fmt"
)
type st struct {    //第一种struct
	i1 int 
	f1 float32
	str string
}

func main(){
	st2 := struct { //第二种struct
		name string
		phone int
	}{
		"jack",1234567890,  //注意这里需要在末尾加上,
	}
	ms :=new(st) //结构体的声明，用new关键字
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "jason"
	fmt.Println("The int is :",ms.i1)
	fmt.Println("The float32 is :",ms.f1)
	fmt.Println("The string is :",ms.str)
	fmt.Println(st2.name)
}
output==>
The int is : 10
The float32 is : 15.5
The string is : jason
jack
</pre>
对结构的元素进行修改
<pre>
package main 

import (
	"fmt"
	"strings"
)
type person struct {
	firstname string
	lastname string
}
func up(p *person){ //如果这里不用指针的形式则不能成功地修改原数据
	p.firstname = strings.ToUpper(p.firstname)
	p.lastname = strings.ToLower(p.lastname)
}
func main(){
	var person1 person
	person1.firstname = "jason"
	person1.lastname = "Bob"
	up(&person1)   //这里
	fmt.Println("The name is the person is",person1.firstname,person1.lastname)
}
output==>
The name is the person is JASON bob
</pre>
给结构体取别名
<pre>
package main

import (
	"fmt"
)
type number struct {
	f float32
}
type aliasnum number  //给结构体number取别名
func main(){
	a :=number{5.8}
	b := aliasnum{4.5}
	var c = number(b)
	fmt.Println(a,b,c)
}
output==>
{5.8} {4.5} {4.5}
</pre>
####匿名字段和内嵌结构体
结构体可以包含一个或多个 匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体。
可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。Go 语言中的继承是通过内嵌或组合来实现的，所以可以说，在 Go 语言中，相比较于继承，组合更受青睐。
<pre>
package main

import (
	"fmt"
)
type inner struct {
	in1 int
	in2 int
}
type outer struct {
	b int
	c float32
	int //匿名字段
	inner //内嵌结构体  (结构体内嵌结构体)
}
func main(){
	outernew := new(outer)
	outernew.b = 6
	outernew.c = 7.6
	outernew.int = 40
	outernew.in1 = 34
	outernew.in2 = 29
	fmt.Println(outernew.b)
	fmt.Println(outernew.c)
	fmt.Println(outernew.int)
	fmt.Println(outernew.in1)
	fmt.Println(outernew.in2)
}
output==>
6
7.6
40
34
29
</pre>
####方法
以一个结构体来示例
<pre>
package main

import (
	"fmt"
)
type twoint struct {
	a int 
	b int
}
func (tn *twoint) AddThem() int{
//这种方式很奇妙，将AddThem作为结构体twoint的方法
	return tn.a + tn.b
}
func main(){
	two :=new(twoint)
	two.a = 23
	two.b = 12
	fmt.Println(two.a)
	fmt.Println(two.b)
	fmt.Println(two.AddThem())
}
output==>
23
12
35
</pre>
以一个非结构体示例
<pre>
package main

import (
	"fmt"
)
type IntVector []int
func (v IntVector) sum()(s int){
	for _,x := range v {
		s += x
	}
	return s
}
func main(){
	fmt.Println(IntVector{2,2,3}.sum())
}
output==>
7
</pre>
####类型的 String()方法和格式化描述符
当定义了一个有很多方法的类型时，十之八九你会使用 String() 方法来定制类型的字符串形式的输出，换句话说：一种可阅读性和打印性的输出。如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：<font color=red><b>等同于使用格式化描述符 %v 产生的输出。还有 fmt.Print() 和 fmt.Println()也会自动使用 String() 方法</b></font>。
<pre>
package main

import (
	"strconv"
	"fmt"
)
type two struct {
	a int
	b int
}
func (tn *two) String() string {
	return "(" + strconv.Itoa(tn.a) +"|" + strconv.Itoa(tn.b)+")"
}
func main(){
	two :=new(two)
	two.a = 2
	two.b = 1
	fmt.Println(two)
	
}
output==>
(2|1)
</pre>
####垃圾回收和SetFinalizer
Go 开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存，在 Go 运行时中有一个独立的进程，即垃圾收集器（GC），会处理这些事情，它搜索不再使用的变量然后释放它们的内存。可以通过 runtime 包访问 GC 进程。
通过调用 runtime.GC() 函数可以显式的触发 GC，但这只在某些罕见的场景下才有用，比如当内存资源不足时调用 runtime.GC()，它会此函数执行的点上立即释放一大片内存，此时程序可能会有短时的性能下降（因为 GC 进程在执行）。
<pre>
package main

import (
	"fmt"
	"runtime"
)
/*如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，
比如写到日志文件中，可以通过如下方式调用函数来实现：
runtime.SetFinalizer(obj, func(obj *typeObj))
*/
func main(){
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc/1024)
}
output==>
34
</pre>
####类型断言：如何检测和转换接口变量的类型
一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的 动态 类型，即运行时在变量中存储的值的实际类型。在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型。通常我们可以使用 类型断言 来测试在某个时刻 varI 是否包含类型 T 的值：
<pre>
	v := varI.(T)    //varI 必须是一个接口变量，否则编译器会报错
</pre>
类型断言可能是无效的，虽然编译器会尽力检查转换是否有效，但是它不可能预见所有的可能性。如果转换在程序运行时失败会导致错误发生。更安全的方式是使用以下形式来进行类型断言：
<pre>
if v, ok := varI.(T); ok {  // checked type assertion
    Process(v)
    return
}
// varI is not of type T
</pre>
<pre>
package main

import (
    "fmt"
    "math"
)

type Square struct {
    side float32
}

type Circle struct {
    radius float32
}

type Shaper interface {
    Area() float32
}

func main() {
    var areaIntf Shaper
    sq1 := new(Square)
    sq1.side = 5

    areaIntf = sq1
    // Is Square the type of areaIntf?
    if t, ok := areaIntf.(*Square); ok {
        fmt.Printf("The type of areaIntf is: %T\n", t)
    }
    if u, ok := areaIntf.(*Circle); ok {
        fmt.Printf("The type of areaIntf is: %T\n", u)
    } else {
        fmt.Println("areaIntf does not contain a variable of type Circle")
    }
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
    return ci.radius * ci.radius * math.Pi
}
output==>
The type of areaIntf is: *main.Square
areaIntf does not contain a variable of type Circle
</pre>
#### 测试一个值是否实现了某个接口
使用接口使代码更具有普适性。
<pre>
type Stringer interface {
    String() string
}

if sv, ok := v.(Stringer); ok {
    fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
}
</pre>
#### 读取用户的输入
我们如何读取用户的键盘（控制台）输入呢？从键盘和标准输入 os.Stdin 读取输入，最简单的办法是使用 fmt 包提供的 Scan 和 Sscan 开头的函数。
<pre>
package main

import (
	"fmt"
)
var (
	firstname,lastname string
)
func main(){
	fmt.Println("Please enter your fullname :")
//Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行
	fmt.Scanln(&firstname,&lastname)
	fmt.Println("hi",firstname,lastname)
}
input==>
jason D
output==>
hi jason D
</pre>
也可以使用 bufio 包提供的缓冲读取（buffered reader）来读取数据:
<pre>
package main

import (
	"fmt"
	"os"
	"bufio"
)
var inputreader *bufio.Reader
var input string
var err error
func main(){
	inputreader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	input,err =inputreader.ReadString('\n')
	//以回车为节点
	if err == nil{
		fmt.Println(input)
	}
}
input==>
hello world
output==>
Please enter some input:
hello world
hello world
</pre>