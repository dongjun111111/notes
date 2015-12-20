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




###go的保留字(25)
break default func interface select case defer go map struct chan else goto 
package switch const fallthrough if range type continue for import return var

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