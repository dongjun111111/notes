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
当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutines转移到另一个系统线程上去，以使这些goroutines不阻塞