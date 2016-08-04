#喵
##基础知识
<font color=red>在Golang中不能用string(int) ==>string，只能用strconv.Itoa(int) ==>string,同理,将string转成int只能用strconv.Atoi；string强制转换只能用于将切片转成string</font>。

字符串

因为Golan中的字符串是不可变的，所以不能像其他语言那样很容易就修改字符串的内容。但是还是有至少下面两种方式来实现字符串内容的修改。

第一种：转成[]byte类型
<pre>
package main

func main() {
/*
Go中字符串是不可变的,所以var s string = "hello" s[0] = 'c' println(s)报错
*/
	var s string = "hello"
	c := []byte(s)  //将字符串 s 转换成 []byte 类型
	c[0] = 'c'
	s = string(c)  //再转回 string 类型
	println(s)
}
</pre>
第二种：切片操作
<pre>
package main

import "fmt"

func main() {
	s := "hello"
	s = "c" + s[1:] //切片操作
	fmt.Println(s)
}
output==>
cello
</pre>
数组  --值类型
<pre>
package main

import "fmt"
//数组
func main() {
	var a [10]int
	a = [10]int{2, 3, 4, 5}
	b := [5]string{"f", "d", "e"}
	c := [...]int{45, 56, 67, 78, 8, 89, 8900, 8} //不定长度
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(a)
}
output==>
[45 56 67 78 8 89 8900 8]
[f d e  ]
[2 3 4 5 0 0 0 0 0 0]
</pre>
切片 -- 引用类型

在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice
<pre>
package main

import "fmt"

//数组
func main() {
	var ar = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	var a, b []byte
	a = ar[:5]
	b = ar[4:7]
	fmt.Println(string(a))
	fmt.Println(b)
}
output==>
abcde
[101 102 103]
</pre>
map 
<pre>
package main

import "fmt"

//数组
func main() {
	var numbers map[string]int = make(map[string]int)
	numbers["w"] = 3
	numbers["s"] = 2
	fmt.Println(numbers["w"])
}
ouput==>
3
</pre>
map 无序的，可能每次打印的map不是相同顺序的；通过 delete 删除 map 元素：
<pre>
package main 

import "fmt"

func main(){
	var a map[int]int = make(map[int]int)
	a[4] = 3
	a[2] = 1
	a[5] = 4
	a[6] = 5
	delete(a,6)   //删除map元素
	fmt.Println(a)
}
两次output==>
[ `go run main.go` | done: 2.5041432s ]
	map[5:4 4:3 2:1]
[ `go run main.go` | done: 2.2391281s ]
	map[4:3 2:1 5:4]        /*说明map是无序的*/
</pre>
<pre>
package main

import "fmt"

func main() {
	rating := map[string]float32{"c": 3, "d": 4, "f": 7, "h": 8}
	//map有两个返回值，分别是value与key，不存在key则是false
	resu, ok := rating["c"]
	if ok {
		fmt.Println(resu) //输出 value 值
	} else {
		fmt.Println("something error")
	}

}
output==>
3
</pre>
<pre>
package main

import "fmt"

//证明map是引用类型
func main() {
	m := make(map[string]string)
	m["hello"] = "Bonjour"
	m1 := m
	m1["hello"] = "salut"
	fmt.Println("m[\"hello\"]", m["hello"])
	fmt.Println("m1[\"hello\"]", m1["hello"])
}
output==>
m["hello"] salut
m1["hello"] salut
</pre>
需要注意的是：map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制。

传值与传指针

先看一下下面两个例子，传值：
<pre>
package main

import "fmt"

func add1(a int) int {
	a = a + 1
	return a
}
func main() {
	x := 4
	fmt.Println("x =", x)

	x1 := add1(x)

	fmt.Println("x+1 =", x1)
	fmt.Println("x =", x)
}
output==>
x = 4
x+1 = 5
x = 4			
</pre>
看到了吗？虽然我们调用了add1函数，并且在add1中执行a = a+1操作，但是上面例子中x变量的值没有发生变化

理由很简单：因为当我们调用add1的时候，add1接收的参数其实是x的copy，而不是x本身。

那你也许会问了，如果真的需要传这个x本身,该怎么办呢？

传指针:
<pre>
package main

import "fmt"

func add1(a *int) int {
	*a = *a + 1
	return *a
}
func main() {
	x := 4
	fmt.Println("x =", x)

	x1 := add1(&x)

	fmt.Println("x+1 =", x1)
	fmt.Println("x =", x)
}
output==>
x = 4
x+1 = 5
x = 5
</pre>
这样，我们就达到了修改x的目的。那么到底传指针有什么好处呢？

- 传指针使得多个函数能操作同一个对象。
- 传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
- Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）

defer 

1.在defer后指定的函数会在函数退出前调用
<pre>
package main
//2.后进先出
import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
output==>
4
3
2
1
0
</pre>
函数作为值、类型
<pre>
package main

import "fmt"

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are :", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are :", even)
}
output==>
slice =  [1 2 3 4 5 6 7]
Odd elements of slice are : [1 3 5 7]
Even elements of slice are : [2 4 6]
</pre>
init函数
<pre>
package main

import "fmt"

//init函数是golang程序自动调用的
func init() {
	fmt.Println("Golang")
}

func main() {

}
output==>
Golang
</pre>
struct
<pre>
package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}
func main() {
	var tom person
	tom.name, tom.age = "Tom", 18
	bob := person{"bob", 45}
	paul := person{"paul", 25}
	tb_older, tb_diff := Older(tom, bob)
	tp_older, tp_diff := Older(tom, paul)
	fmt.Println(tom.name, bob.name, tb_older.name, tb_diff)
	fmt.Println(tom.name, paul.name, tp_older.name, tp_diff)
}
output==>
Tom bob bob 27
Tom paul paul 7
</pre>
匿名结构
<pre>
package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human  //匿名字段，默认Student就包含了Human的所有字段
	school string
}

func main() {
	mark := Student{Human{"mark", 34, 60}, "THU"}
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His school is ", mark.school)
}
output==>
His name is  mark
His age is  34
His school is  THU
</pre>
匿名的其他类型 []string int等
<pre>
package main

import "fmt"

type Skills []string
type Human struct {
	name   string
	age    int
	weight int
}
type Student struct {
	Human  //匿名字段,自定义的struct
	Skills //匿名字段，自定义的类型string slice
	int    // 内置类型作为匿名字段
	school string
}

func main() {
	jack := Student{Human: Human{"jack", 34, 80}, school: "CMU"}
	fmt.Println(jack.name)
	fmt.Println(jack.age)
	fmt.Println(jack.school)
	jack.Skills = []string{"there is a long long words"}
	jack.Skills = append(jack.Skills, ",so how it is going on next")
	fmt.Println(jack.Skills)
	jack.int = 77
	fmt.Println(jack.int)
}
output==>
jack
34
CMU
[there is a long long words ,so how it is going on next]
77
</pre>
<pre>
package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Skills []string
type Student struct {
	Human
	Skills
	int
	num string
}

func main() {
	jason := Student{Human: Human{"Jason", 24}, Skills: []string{"Less is more"}, int: 77, num: "you guess"}
	fmt.Println(jason.name)
	fmt.Println(jason.int)
	fmt.Println(jason.num)
}
output==>
Jason
77
you guess
</pre>
这里有一个问题：如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，那么该怎么办呢？

Go里面很简单的解决了这个问题，最外层的优先访问，也就是当你通过student.phone访问的时候，是访问student里面的字段，而不是human里面的字段。

这样就允许我们去重载通过匿名字段继承的一些字段，当然如果我们想访问重载后对应匿名类型里面的字段，可以通过匿名字段名来访问。请看下面的例子
<pre>
package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Employee struct {
	Human
	age int
}

func main() {
	bob := Employee{Human{"bob", 23}, 12}

	fmt.Println("bob`s age is :", bob.age)
	fmt.Println("bob`s human age is :", bob.Human.age)
}
output==>
bob`s age is : 12
bob`s human age is : 23
</pre>
还有一个例子也是说重载的：
<pre>
package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	school string
}

func (h *Human) say() {
	fmt.Println(h.name, h.age)
}

func (h *Human) run() {
	fmt.Println(h.name + "is running")
}

func (s *Student) say() {
	fmt.Println("a student whois name is " + s.name + " ,saying something")
}

func (s *Student) study() {
	fmt.Println("stupid study")
}

func main() {
	jason := Student{Human{"jason", 12}, "MIT"}
	jason.say()       //a student whois name is jason ,saying something
	jason.Human.say() //jason 12
}
output==>
a student whois name is jason ,saying something
jason 12
</pre>
面向对象

函数的另一种形态，带有接收者的函数，我们称为method.
<pre>
package main

import "fmt"

type rectangle struct {
	width, height float64
}

func area(r rectangle) float64 {
	return r.width * r.height
}

func main() {
	rec := rectangle{34, 6}
	fmt.Println(area(rec))
}
output==>
204		
</pre>
将上面的改写成下面的形式(将area作为struct rectangle的一种属性或者说是方法):
<pre>
package main

import "fmt"

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func main() {
	ar := rectangle{4, 6}
	fmt.Println(ar.area())
}
output==>
24
</pre>
method继承

如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。让我们来看下面这个例子.
<pre>
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

//在human上面定义了一个method
func (h *Human) sayhi() {
	fmt.Println(h.name, h.phone)
}

func main() {
	mark := Student{Human{"mark", 34, "4543445454"}, "MIT"}
	jason := Employee{Human{"Jason", 23, "543434223"}, "ZCMLC"}
	mark.sayhi()
	jason.sayhi()
}
output==>
mark 4543445454
Jason 543434223
</pre>
method重写

上面的例子中，如果Employee想要实现自己的SayHi,怎么办？简单，和匿名字段冲突一样的道理，我们可以在Employee上面定义一个method，重写了匿名字段的方法。请看下面的例子.
<pre>
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) sayhi() {         //Human定义的sayhi方法
	fmt.Println(h.name, h.phone)
}

func (e *Employee) sayhi() {      //Employee定义的sayhi方法
	fmt.Println(e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"mark", 23, "34545"}, "CMU"}
	jason := Employee{Human{"jason", 24, "55753"}, "Golang Inc"}
	mark.sayhi()
	jason.sayhi()
}
output==>
mark 34545
jason Golang Inc 55753
</pre>
interface

Golang里面设计最精妙的应该算interface，它让面向对象，内容组织实现非常的方便，当你看完这一章，你就会被interface的巧妙设计所折服。

简单的说，interface是一组method签名的组合，我们通过interface来定义对象的一组行为。

我们前面一章最后一个例子中Student和Employee都能SayHi，虽然他们的内部实现不一样，但是那不重要，重要的是他们都能say hi

让我们来继续做更多的扩展，Student和Employee实现另一个方法Sing，然后Student实现方法BorrowMoney而Employee实现SpendSalary。

这样Student实现了三个方法：SayHi、Sing、BorrowMoney；而Employee实现了SayHi、Sing、SpendSalary。

上面这些方法的组合称为interface(被对象Student和Employee实现)。例如Student和Employee都实现了interface：SayHi和Sing，也就是这两个对象是该interface类型。而Employee没有实现这个interface：SayHi、Sing和BorrowMoney，因为Employee没有实现BorrowMoney这个方法。

interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。详细的语法参考下面这个例子。
<pre>
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) Sayhi() {
	fmt.Println(h.name, h.phone)
}

func (h Human) Sing(lyric string) {
	fmt.Println("Lalala...", lyric)
}

func (e Employee) Sayhi() {
	fmt.Println(e.name, e.company, e.phone)
}

type Men interface {
	Sayhi()
	Sing(lyric string)
}

func main() {
	mike := Student{Human{"mike", 23, "2455656"}, "MIT", 6}
	jason := Employee{Human{"jason", 24, "1247890"}, "Google Inc", 200}
	//定义Men的变量i
	var i Men
	i = mike
	fmt.Println("This is mike,a student:")
	i.Sayhi()
	i.Sing("I love U")

	i = jason
	fmt.Println("This is jason,an employee:")
	i.Sayhi()
	i.Sing("Born to be wild")

}
output==>
This is mike,a student:
mike 2455656
Lalala... I love U
This is jason,an employee:
jason Google Inc 1247890
Lalala... Born to be wild
</pre>
interface 与指针
<pre>
package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	school string
}

func (h *Human) say() {
	fmt.Println(h.name, h.age)
}

func (h *Human) run() {
	fmt.Println(h.name + " is running")
}

func (s *Student) say() {
	fmt.Println("a student whois name is " + s.name + " ,saying something")
}

func (s *Student) Study() {
	fmt.Println("stupid study")
}

type M interface {
	say()
	run()
}

func main() {
	jason := Student{Human{"jason", 6}, "MIT"}
	var m M
	m = &jason      //注意这里的 &
	m.run()
}
output==>
jason is running
</pre>
interface函数参数

interface的变量可以持有任意实现该interface类型的对象，这给我们编写函数(包括method)提供了一些额外的思考，我们是不是可以通过定义interface参数，让函数接受各种类型的参数。

举个例子：fmt.Println是我们常用的一个函数，但是你是否注意到它可以接受任意类型的数据。打开fmt的源码文件，你会看到这样一个定义:
<pre>
type Stringer interface {
	String() string
}
</pre>
也就是说，任何实现了String方法的类型都能作为参数被fmt.Println()调用，比如：
<pre>
package main

import "fmt"
import "strconv"

type Human struct {
	name  string
	age   int
	phone string
}

//通过这个方法Human 实现了 fmt.Stringer
func (h Human) String() string {
	return ".." + h.name + strconv.Itoa(h.age) + " years " + h.phone
}

func main() {
	bob := Human{"bob", 34, "57357878"}
	fmt.Println(bob.String())
}
output==>
..bob34 years 57357878
</pre>
并发
<pre>
package main

import "fmt"
import "runtime"

//runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func main() {
	go say("world")  //开启一个新的Goroutines执行
	say("hello")  //当前的Goroutines执行
}
output==>
hello
world
hello
world
hello
world
hello
world
hello
</pre>
超时
<pre>
package main

import "fmt"
import "time"

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(2 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
output==>
timeout
</pre>
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
###自增长 ID 生成器
<pre>
package main

import "fmt"

//自增长 ID 生成器
type AutoInc struct {
	start, step int
	queue       chan int
	running     bool
}

func New(start, step int) (ai *AutoInc) {
	ai = &AutoInc{
		start:   start,
		step:    step,
		running: true,
		queue:   make(chan int, 4),
	}
	go ai.process()
	return
}

func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
	}
}

func (ai *AutoInc) Id() int {
	return <-ai.queue
}

func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}

func main() {
	ai := New(1, 7)
	defer ai.Close()
	for i := 0; i < 10; i = i + 2 {
		if id := ai.Id(); id != i {
			fmt.Println(id)
		}
	}
}
output==>
1
8
15
22
29
</pre>
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
###结构体Struct
<pre>
package main

import "fmt"

func main() {
	type S struct {
		a int
		b string
	}
	//结构体初始化通过结构体字段的值作为列表来新分配一个结构体
	var s S = S{4, "jason"}
	fmt.Println(s)   //{4 jason}
	fmt.Println(s.a) //4
	//结构体是值类型，传递时会复制值，其默认零值不是nil
	var a S
	var b = S{}
	fmt.Println(a)      //{0 }
	fmt.Println(b)      //{0 }
	fmt.Println(a == b) //true
	type People struct {
		name  string
		age   int
		phone int
	}
	var jason People = People{"jason", 12, 123434545}
	fmt.Println("jason`s phone :", jason.phone)
	fmt.Println("jason`s name :", jason.name)
	//匿名结构体
	//匿名结构体声明时省略了type关键字，并且没有名称
	var x struct{}
	var y struct{ x int }
	fmt.Println(x, y) //{} {0}
	y.x = 3
	fmt.Println(y.x) //3
	
}
output==>
{4 jason}
4
{0 }
{0 }
true
jason`s phone : 123434545
jason`s name : jason
{} {0}
3
</pre>
结构体的tag属性，类似注释内容。 下面利用reflect包将tag内容输出出来。
<pre>
package main

import "fmt"
import "reflect"

type User struct {
	Name     string "This is an user name"       //这后面的是tag
	Password string "This is an user password"	 //这后面的是tag
}

func main() {
	user := &User{"Jason", "password"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag)  //将tag输出出来
	}
}
output==>
This is an user name
This is an user password
</pre>
###指针pointer
<pre>
package main

import "fmt"

func main() {
	var i int = 1
	pi := &i
	fmt.Println(pi) //0xc082006288
	a := []int{4, 5, 6}
	pa := &a
	fmt.Println(pa) //&[4 5 6]
	//使用*读取/修改指针指向的值
	i1 := new(int)
	*i1 = 3
	fmt.Println(i1, *i1) //0xc082048098 3
}
output==>
0xc082048038
&[4 5 6]
0xc082048098 3
</pre>
###通道 Channel
channel用于两个goroutine之间传递指定类型的值来同步运行和通讯。操作符<-用于指定channel的方向，发送或接收。如果未指定方向，则为双向channel。

channel是引用类型，使用make函数来初始化。未初始化的channel零值是nil，且不能用于发送和接收值。
<pre>
package main

import "fmt"

func main() {
	/*
		关闭channel，只能用于双向或只发送类型的channel
		只能由 发送方调用close函数来关闭channel
		接收方取出已关闭的channel中发送的值后，后续再从channel中取值时会以非阻塞的方式立即返回channel传递类型的零值。
	*/
	ch := make(chan string, 1)
	ch <- "hello"
	close(ch)
	s, ok := <-ch
	if ok {
		fmt.Println("receive value from sender:", s)
	} else {
		fmt.Println("get zero value from closed channel")
	}
	/*//向已关闭的通道发送值会产生运行时恐慌panic
	ch <- "hi"
	fmt.Println(<-ch)
	// 再次关闭已经关闭的通道也会产生运行时恐慌panic
	close(ch)
	*/

	//使用for range语句依次读取发送到channel的值，直到channel关闭。
	var chh = make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			chh <- i
		}
		close(chh)
	}()
	for x := range chh {
		fmt.Printf("%d-", x)
	}
}
output==>
receive value from sender: hello
0-1-2-3-4-
</pre>
在实际的项目中，我们的程序一般都是很多个goroutine同时工作，知道所有goroutine是否都完成不是一件容易的事情。以前的经验是通过轮询的方式，但是在golang中这种方式比较浪费性能。
<pre>
package main

import (
	"fmt"
)

var (
	flag bool
	str  string
)

var ch chan string = make(chan string)
/*
不要用无限轮询的方式来检查goroutine是否完成,而是要通过使用channel，
让foo()和main()实现通信，让foo()执行完毕后通过channel发送一个消息给main()，
告诉它自己的事儿完成了，然后main()收到消息后继续执行其他操作
*/
func foo() {
	flag = true
	str = "setup complete"
	ch <- "I am complete"
}
func main() {
	go foo()
	<-ch
	for !flag {
	}
	fmt.Println(str)
}

</pre>
使用channel精确控制goroutine的数量
<pre>
package main

import "fmt"

var ch chan int = make(chan int)

func afunction(ch chan int, t int) {
	fmt.Println(t+1, ":finished")
	<-ch
}

func main() {
	for i := 0; i < 7; i++ {
		go afunction(ch, i)
		ch <- 1
	}
}
output==>
1 :finished
2 :finished
3 :finished
4 :finished
5 :finished
6 :finished
7 :finished
</pre>
###switch case goto break continue
<pre>
package main

import "fmt"

func main() {
	x := 2
	//分支选择 switch
	switch x {
	case 0:
		fmt.Println("x=0")
	case 1:
		fmt.Println("x=1")
	case 2:
		fmt.Println("x=2")
	default:
		fmt.Println("default value")
	}
	switch {
	case x == 1:
		fmt.Println("1")
	case x == 2:
		fmt.Println("bingo 2")
	default:
		fmt.Println("default value")
	}
	//循环
	sl := []int{2, 4, 5, 6, 7}
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%d ", sl[i])
	}
	fmt.Println()
	for k, v := range sl {
		fmt.Println("key:", k, "value:", v) //下标,值
	}
	//循环的继续、中断、跳转
	for k, v := range sl {
		if v == 2 {
			fmt.Println(k)
			continue
		} else if v == 5 {
			break
		} else {
			goto JASON
		}
	JASON:
		fmt.Println("goto action done")
	}

}
output==>
x=2
bingo 2
2 4 5 6 7 
key: 0 value: 2
key: 1 value: 4
key: 2 value: 5
key: 3 value: 6
key: 4 value: 7
0
goto action done
</pre>
###有缓冲 无缓冲channel
无缓冲
<pre>
package main

import "fmt"

func writeRoutine(test_chan chan int, value int) {
	test_chan <- value
}
func readRoutine(test_chan chan int) {
	<-test_chan
	return
}
func main() {
	c := make(chan int)
	x := 100
	go writeRoutine(c, x)
	readRoutine(c)
	fmt.Println(x)
}
outout==>
100
</pre>
有缓冲
<pre>
package main 
import "fmt"
var c = make(chan int, 1)

func f() {
     c <- 'c'
 
     fmt.Println("在goroutine内")
 }
 
 func main() {
     go f()
 
     c <- 'c'  //要是没有这个，则"在main中"没有打印
     <-c
     <-c
 
     fmt.Println("在main中")  
 }
output==>
在goroutine内
在main中
</pre>
###方法
_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
<pre>
package main

import "fmt"

type A struct {
	x, y int
}

/*// 定义结构体的方法，'_'表示方法内忽略使用结构体、字段及其他方法
func (_ A) echo_A() {
	fmt.Println("(_ A)")
} */
// 同上
func (A) echoA(s string) {
	fmt.Println("(_A)", s)
}

/*
func (_ *A) echo_PA() {
	fmt.Println("(_ *A)")
}
*/
// 同上
func (*A) echo_PA(s string) {
	fmt.Println("(*A)", s)
}

//定义结构体的方法，方法内可以引用结构，字段及其他方法
func (a A) setX(x int) {
	a.x = x
}

// 定义结构体指针的方法，方法内可以引用结构体、结构体指针、字段及其他方法
func (a *A) setY(y int) {
	a.y = y
}
func main() {
	var a A
	a.setX(3)
	a.setY(9)
	fmt.Println(a.x, a.y)
	a.echoA("a")       //(_A) a
	a.echo_PA("jason") //(*A) jason
}
output==>
0 9
(_A) a
(*A) jason
</pre>
###并发 Concurrency
使用关键字go调用一个函数/方法，启动一个新的协程goroutine
<pre>
package main

import "time"

//主协程goroutine输出0，其他由go启动的几个子协程分别输出1～5
func say(i int) {
	println("goroutine:", i)
}
func main() {
	for i := 1; i <= 5; i++ {
		go say(i)
	}
	say(0)
	time.Sleep(5 * time.Second)
}
output==>
goroutine: 0
goroutine: 1
goroutine: 2
goroutine: 3
goroutine: 4
goroutine: 5
</pre>
goroutine 在相同的地址空间中运行，因此访问共享内存必须进行同步
<pre>
package main

import "time"
import "sync"

var mu sync.Mutex
var i int

func add() {
	/*
		使用互斥锁防止多个协程goroutine同时修改共享变量
		只能限制同时访问此方法修改变量，在方法外修改则限制是无效的
	*/
	mu.Lock()
	defer mu.Unlock()
	i++
}
func main() {
	for range [100]byte{} {
		go add()
	}
	time.Sleep(1 * time.Second)
	println(i)
}
output==>
100
</pre>
使用通道channel进行同步
<pre>
package main

import "time"

var i int
var ch = make(chan byte, 1)

//将channel用作同步开关
func main() {
	for range [100]byte{} {
		go add()
	}
	time.Sleep(1 * time.Second)
	println(i)
}
func add() {
	ch <- 0
	i++
	<-ch
}
output==>
100
</pre>
使用channel在不同的goroutine之间通信
<pre>
package main

import "time"

var i int
var ch = make(chan int, 1)

func add() {
	x := <-ch
	x++
	ch <- x
}
func main() {
	for range [100]byte{} {
		go add()
	}
	ch <- i
	time.Sleep(1 * time.Second)
	i = <-ch
	println(i)
}
output==>
100
</pre>
###测试 Testing
Go中自带轻量级的测试框架testing和自带的go test命令来实现单元测试和基准测试
####单元测试 Unit
- 测试源文件名必须是_test.go结尾的，go test的时候才会执行到相应的代码
- 必须import testing包
- 所有的测试用例函数必须以Test开头
- 测试用例按照源码中编写的顺序依次执行
- 测试函数TestXxx()的参数是*testing.T，可以使用该类型来记录错误或者是测试状态
- 测试格式：func TestXxx (t *testing.T)，Xxx部分可以为任意的字母数字的组合，首字母不能是小写字母[a-z]，例如Testsum是错误的函数名。
- 函数中通过调用*testing.T的Error，Errorf，FailNow，Fatal，FatalIf方法标注测试不通过，调用Log方法用来记录测试的信息。

测试分两个文件，分别是:

- test.go
<pre>
package testgo
import "math"
func Sum(min, max int) (sum int) {
	if min < 0 || max < 0 || max > math.MaxInt32 || min > max {
		return 0
	}
	for ; min <= max; min++ {
		sum += min
	}
	return
}
</pre>

- test_test.go

<pre>
package testgo
import "testing"
func TestSum(t *testing.T) {
	s := Sum(1, 0)
	t.Log("Sum 1 to 0:", s)
	if 0 != s {
		t.Error("not equal.")
	}
	s = Sum(1, 10)
	t.Log("Sum 1 to 10:", s)
	if 55 != s {
		t.Error("not equal.")
	}
}
</pre>
在当前位置执行测试：
<pre>
go test -v

输出：
=== RUN   TestSum
--- PASS: TestSum (0.00s)
	test_test.go:7: Sum 1 to 0: 0
	test_test.go:12: Sum 1 to 10: 55
PASS
ok  	test	0.237s
</pre>
####基准测试 Benchmark
基准测试 Benchmark用来检测函数/方法的性能。

- 基准测试用例函数必须以Benchmark开头
- go test默认不会执行基准测试的函数，需要加上参数-test.bench，语法:-test.bench="test_name_regex"，例如go test -test.bench=".*"表示测试全部的基准测试函数
- 在基准测试用例中，在循环体内使用testing.B.N，使测试可以正常的运行

测试分两个文件，分别是:

- test.go
<pre>
//test.go
package testgo
import "math"
func Sum(min, max int) (sum int) {
	if min < 0 || max < 0 || max > math.MaxInt32 || min > max {
		return 0
	}

	for ; min <= max; min++ {
		sum += min
	}
	return
}
</pre>
- test_test.go
<pre>
package testgo
import "testing"
func BenchmarkSum(b *testing.B) {
    b.Logf("Sum 1 to %d: %d\n", b.N, Sum(1, b.N))
}
</pre>
在当前位置执行测试： (注意bench后面还有一个 . )
<pre>
go test -v -bench .   

	输出：
	PASS
	BenchmarkSum-2	2000000000	         0.69 ns/op
	--- BENCH: BenchmarkSum-2
		test_test.go:6: Sum 1 to 1: 1
		test_test.go:6: Sum 1 to 100: 5050
		test_test.go:6: Sum 1 to 10000: 50005000
		test_test.go:6: Sum 1 to 1000000: 500000500000
		test_test.go:6: Sum 1 to 100000000: 5000000050000000
		test_test.go:6: Sum 1 to 2000000000: 2000000001000000000
	ok  	test	1.697s
	
	testing: warning: no tests to run
</pre>

###简单web服务器

<pre>
package main
import "fmt"
import "net/http"
import "log"
func sayhi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Jason")
}
func jason(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am jason")
}
func main() {
	http.HandleFunc("/", sayhi)
	http.HandleFunc("/jason", jason)
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal("listenandserve:", err)
	}
}
</pre>
###模板 template
<pre>
func temp(w http.ResponseWriter, r *http.Request) {
	t := template.New("some template")      //创建一个模板
	t, _ = t.ParseFiles("welcome.tpl", nil) //解析模板文件
	user := Getuser()                       //获取用户信息
	t.Execute(w, user)                      //执行模板的merge操作
}
/*
Parse与ParseFiles：Parse可以直接测试一个字符串，而不需要额外的文件；ParseFiles只能解析文件;
os.Stdout与http.ResponseWriter:os.Stdout实现了io.Writer接口，而http.ResponseWriter没有.
实现了io.Writer接口就可以用于t.Execute(),比如 temp.Execute(os.Stdout,element) [temp:待渲染的模板或则字符串,element:写入模板的参数] 。没有实现io.Writer接口的则会报如下的错误：
.\test.go:15: cannot use http.ResponseWriter (type func(...interface {}) (int, error)) as type io.Writer in argument to t.Execute:

*/
</pre>
####模板中如何插入数据？
上面我们演示了如何解析并渲染模板，接下来让我们来更加详细的了解如何把数据渲染出来。一个模板都是应用在一个Go的对象之上，Go对象的字段如何插入到模板中呢？
#####字段操作
Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象，这和Java或者C++中的this类似，如果要访问当前对象的字段通过{{.FieldName}},但是需要注意一点：这个字段必须是导出的(字段首字母必须是大写的),否则在渲染的时候就会报错，请看下面的这个例子：
<pre>
package main

import "html/template"
import "os"

type Person struct {
	Username string // 字段首字母必须大写，否则失败
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.Username }}")
	p := Person{Username: "Jason"}
	t.Execute(os.Stdout, p)
}
output==>
hello Jason
</pre>
#####输出嵌套字段内容
上面我们例子展示了如何针对一个对象的字段输出，那么如果字段里面还有对象，如何来循环的输出这些内容呢？我们可以使用{{with …}}…{{end}}和{{range …}}{{end}}来进行数据的输出。

- {{range}} 这个和Go语法里面的range类似，循环操作数据
- {{with}}操作是指当前对象的值，类似上下文的概念

<pre>
package main
import "html/template"
import "os"
type Friend struct {
	Fname string
}
type Person struct {
	Username string
	Emails   []string
	Friends  []*Friend
}
func main() {
	f1 := Friend{Fname: "Jack"}
	f2 := Friend{Fname: "Jason"}
	t := template.New("fieldname template")
	/*{{with}}操作是指当前对象的值，类似上下文的概念*/
	t, _ = t.Parse(`
		hello {{.Username}}!
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}  
		{{range .}}
			my friend name is {{.Fname}}
		{{end}}
		{{end}}
	`)
	p := Person{Username: "Jason",
		Emails:  []string{"jason@qq.com", "jason@163.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
output==>
hello Jason!
			
	an email jason@qq.com

	an email jason@163.com

	my friend name is Jack

	my friend name is Jason
</pre>

#####条件处理
在Go模板里面如果需要进行条件判断，那么我们可以使用和Go语言的if-else语法类似的方式来处理，如果pipeline为空，那么if就认为是false，下面的例子展示了如何使用if-else语法：
<pre>
package main

import (
	"os"
	"text/template"
)

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空的pipeline if demo:{{if ``}}不会输出.{{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的pipeline if demo:{{if `anything`}} 我有内容，我会输出.{{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}
output==>
空的pipeline if demo:
不为空的pipeline if demo: 我有内容，我会输出.
if-else demo:  if部分
</pre>
#####pipelines
Unix用户已经很熟悉什么是pipe了，ls | grep "example"类似这样的语法你是不是经常使用，过滤当前目录下面的文件，显示含有"example"的数据，表达的意思就是前面的输出可以当做后面的输入，最后显示我们想要的数据，而Go语言模板最强大的一点就是支持pipe数据，<font color=red>在Go语言里面任何{{}}里面的都是pipelines数据</font>，例如我们上面输出的email里面如果还有一些可能引起XSS注入的，那么我们如何来进行转化呢？
<pre>
{{. | html}}
</pre>
在email输出的地方我们可以采用如上方式可以把输出全部转化html的实体，上面的这种方式和我们平常写Unix的方式是不是一模一样，操作起来相当的简便，调用其他的函数也是类似的方式。
#####模板变量
有时候，我们在模板使用过程中需要定义一些局部变量，我们可以在一些操作中申明局部变量，例如withrangeif过程中申明局部变量，这个变量的作用域是{{end}}之前，Go语言通过申明的局部变量格式如下所示：
<pre>
$variable := pipeline
</pre>
详细的例子看下面的：
<pre>
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
{{with $x := "output"}}{{printf "%q" $x}}{{end}}
</pre>
#####模板函数
模板在输出对象的字段值时，采用了fmt包把对象转化成了字符串。但是有时候我们的需求可能不是这样的，例如有时候我们为了防止垃圾邮件发送者通过采集网页的方式来发送给我们的邮箱信息，我们希望把@替换成at例如：jason at jason.info，如果要实现这样的功能，我们就需要自定义函数来做这个功能。

每一个模板函数都有一个唯一值的名字，然后与一个Go函数关联，通过如下的方式来关联
<pre>
type FuncMap map[string]interface{}
</pre>
例如，如果我们想要的email函数的模板函数名是emailDeal，它关联的Go函数名称是EmailDealWith,那么我们可以通过下面的方式来注册这个函数
<pre>
t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
</pre>
EmailDealWith这个函数的参数和返回值定义如下：
<pre>
func EmailDealWith(args …interface{}) string
</pre>
例子如下：
<pre>
package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	Username string
	Emails   []string
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	return (substrs[0] + " AT " + substrs[1])
}
func main() {
	t := template.New("filedname test")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`
		hello {{.Username}}!
	        {{range .Emails}}
	            an emails {{.|emailDeal}}
	        {{end}}
	`)
	p := Person{Username: "Jason", Emails: []string{"jason@qq.com", "jason@163.com", "jack@sina.com"}}
	t.Execute(os.Stdout, p)
}
output==>
hello Jason!
		        
    an emails jason AT qq.com

    an emails jason AT 163.com

    an emails jack AT sina.com
</pre>
类似的还有一个是重写gt eq lt等常见筛选条件：
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
}
output==>
lulu: 5 > age.
lili: 5 < age.
</pre>
上面演示了如何自定义函数，其实，在模板包内部已经有内置的实现函数，下面代码截取自模板包里面
<pre>
var builtins = FuncMap{
    "and":      and,
    "call":     call,
    "html":     HTMLEscaper,
    "index":    index,
    "js":       JSEscaper,
    "len":      length,
    "not":      not,
    "or":       or,
    "print":    fmt.Sprint,
    "printf":   fmt.Sprintf,
    "println":  fmt.Sprintln,
    "urlquery": URLQueryEscaper,
}
</pre>
#####Must操作
模板包里面有一个函数Must，它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写。接下来我们演示一个例子，用Must来判断模板是否正确：
<pre>
package main

import "fmt"
import "text/template"

func main() {
	tok := template.New("first")
	template.Must(tok.Parse("some static text /*and a comment*/"))
	fmt.Println("The first one parsed ok")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("the second one parsed ok")

	fmt.Println("the next one ought to fail")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse("some static text {{ .Name }"))
}
output==>
The first one parsed ok
the second one parsed ok
the next one ought to fail

panic: template: check parse error with Must:1: unexpected "}" in operand

goroutine 1 [running]:
panic(0x556d40, 0xc08202a3e0)
	D:/go/src/runtime/panic.go:481 +0x3f4
text/template.Must(0x0, 0x760000, 0xc08202a3e0, 0x0)
	D:/go/src/text/template/helper.go:23 +0x52
main.main()
	D:/gopath/src/test/test.go:16 +0x928
exit status 2

exit status 1
</pre>
#####嵌套模板
我们平常开发Web应用的时候，经常会遇到一些模板有些部分是固定不变的，然后可以抽取出来作为一个独立的部分，例如一个博客的头部和尾部是不变的，而唯一改变的是中间的内容部分。所以我们可以定义成header、content、footer三个部分。Go语言中通过如下的语法来申明
<pre>
{{define "子模板名称"}}内容{{end}}
</pre>
通过如下方式来调用：
<pre>
{{template "子模板名称"}}
</pre>
接下来我们演示如何使用嵌套模板，我们定义三个文件，header.tmpl、content.tmpl、footer.tmpl文件，里面的内容如下

header.tmpl
<pre>
{{define "header"}}
<html>
<head>
    <title>演示信息</title>
</head>
<body>
{{end}}
</pre>
content.tmpl
<pre>
//content.tmpl
{{define "content"}}
{{template "header"}}
<h1>演示嵌套</h1>
<ul>
    <li>嵌套使用define定义子模板</li>
    <li>调用使用template</li>
</ul>
{{template "footer"}}
{{end}}
</pre>
footer.tmpl
<pre>
//footer.tmpl
{{define "footer"}}
</body>
</html>
{{end}}
</pre>
演示如下：
<pre>
package main

import (
    "fmt"
    "os"
    "text/template"
)

func main() {
    s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
    s1.ExecuteTemplate(os.Stdout, "header", nil)
    fmt.Println()
    s1.ExecuteTemplate(os.Stdout, "content", nil)
    fmt.Println()
    s1.ExecuteTemplate(os.Stdout, "footer", nil)
    fmt.Println()
    s1.Execute(os.Stdout, nil)
}
</pre>
通过上面的例子我们可以看到通过template.ParseFiles把所有的嵌套模板全部解析到模板里面，其实每一个定义的{{define}}都是一个独立的模板，他们相互独立，是并行存在的关系，内部其实存储的是类似map的一种关系(key是模板的名称，value是模板的内容)，然后我们通过ExecuteTemplate来执行相应的子模板内容，我们可以看到header、footer都是相对独立的，都能输出内容，content 中因为嵌套了header和footer的内容，就会同时输出三个的内容。但是当我们执行s1.Execute，没有任何的输出，因为在默认的情况下没有默认的子模板，所以不会输出任何的东西。
#####将struct传入模板
<pre>
package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{"wool", 34}
	muban := "{{.Count}} items are made of {{.Material}}"
	tmpl, err := template.New("test").Parse(muban)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters) //将struct与模板合成，合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}
outout==>
34 items are made of wool
</pre>

##深入理解 net/http 
###第一版
<pre>
package main

import (
	"io"
	"net/http"
)

/*
Hander是啥呢，它是一个接口。这个接口很简单，只要某个struct
有ServeHTTP(http.ResponseWriter, *http.Request)这个方法，
那这个struct就自动实现了Hander接口
*/

func sayhi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello jason")
}

func main() {
	http.HandleFunc("/", sayhi)       //注册一个sayhello函数给“/”，当浏览器浏览“/”的时候，会调用sayhello函数
	http.ListenAndServe(":8089", nil) //开始监听和服务
}
</pre>
###第二版
认识http.ResponseWriter

当http.ListenAndServe(":8080", &a{})后，开始等待有访问请求
一旦有访问请求过来，http包帮我们处理了一系列动作后，最后他会去调用a的ServeHTTP这个方法，并把自己已经处理好的http.ResponseWriter, *http.Request传进去
而a的ServeHTTP这个方法，拿到*http.ResponseWriter后，并往里面写东西，客户端的网页就显示出来了
<pre>
package main

//重写ServeHTTP方法
import (
	"io"
	"net/http"
)

type a struct{}

func (*a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 1")
}
func main() {
	http.ListenAndServe(":8089", &a{})
}
</pre>
认识*http.Request
<pre>
package main

import (
	"io"
	"net/http"
)

type a struct{}

func (*a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String() //获取访问的路径
	io.WriteString(w, path)
}

func main() {
	http.ListenAndServe(":8089", &a{})
}
output==>
地址栏输入：http://localhost:8089/ffffffffffffffffffffff
/ffffffffffffffffffffff
</pre>
一个非常简单的网站
<pre>
package main

import (
	"io"
	"net/http"
)

type a struct{}

func (*a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	switch path {
	case "/":
		io.WriteString(w, "<h1>ROOT</h1><a href=\"abc\">abc</a> | <a href=\"hello\">hello</a>")
	case "/abc":
		io.WriteString(w, "<h1>ABC</h1><a href=\"/\">root</a>")
	case "/hello":
		io.WriteString(w, "<h1>HELLO</h1><a href=\"/\">root</a>")
	}
}

func main() {
	http.ListenAndServe(":8089", &a{})
}
/*
运行后，可以看出，一个case就是一个页面
如果一个网站有上百个页面，那是否要上百个case？
很不幸，是的
那管理起来岂不是要累死？
要累死，不过，还好有ServeMux
*/
</pre>
###第三版-用ServeMux拯救上面的问题
ServeMux大致作用是，他有一张map表，map里的key记录的是r.URL.String()，而value记录的是一个方法，这个方法ServeHTTP是一样的，这个方法有一个别名，叫HandlerFunc.ServeMux还有一个方法名字是Handle，他是用来注册HandlerFunc 的.
ServeMux还有另一个方法名字是ServeHTTP，这样ServeMux是实现Handler接口的，否者无法当http.ListenAndServe的第二个参数传输.
<pre>
package main

import (
	"io"
	"net/http"
)

type b struct{}

func (*b) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 2")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hi", &b{})
	http.ListenAndServe(":8089", mux)
}
/*
mux := http.NewServeMux():新建一个ServeMux。
mux.Handle("/", &b{}):注册路由，把"/"注册给b这个实现Handler接口的struct，注册到map表中。
http.ListenAndServe(":8080", mux)第二个参数是mux。
运行时，因为第二个参数是mux，所以http会调用mux的ServeHTTP方法。
ServeHTTP方法执行时，会检查map表（表里有一条数据，key是“/h”，value是&b{}的ServeHTTP方法）
如果用户访问/h的话，mux因为匹配上了，mux的ServeHTTP方法会去调用&b{}的 ServeHTTP方法，从而打印hello
如果用户访问/abc的话，mux因为没有匹配上，从而打印404 page not found

ServeMux就是个二传手！
*/
</pre>
ServeMux的HandleFunc方法
<pre>
package main

/*
发现了没有，b这个struct仅仅是为了装一个ServeHTTP而存在，所以能否跳过b呢，
ServeMux说：可以 mux.HandleFunc是用来注册func到map表中的
*/
import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hi")
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ROOT")
	})

	http.ListenAndServe(":8089", mux)
}
</pre>
###time 
time.Sleep
<pre>
fmt.Println("start sleeping...")
time.Sleep(time.Second)
fmt.Println("end sleep.")
//【结果】打印start sleeping后，等了正好1秒后，打印了end sleep
//会阻塞，Sleep时，什么事情都不会做
</pre>
time.After
<pre>
package main

import "fmt"
import "time"

func main() {
	fmt.Println("the 1")
	//返回一个time.C这个管道，1秒(time.Second)后会在此管道中放入一个时间点
	tc := time.After(time.Second)
	fmt.Println("the 2")
	fmt.Println("the 3")
	<-tc
	fmt.Println("the 4")
}
//【结果】立即打印123，等了1秒不到一点点的时间，打印了4，结束
//打印the 1后，获得了一个空管道，这个管道1秒后会有数据进来
//打印the 2，（这里可以做更多事情）
//打印the 3
//等待，直到可以取出管道的数据（取出数据的时间与获得tc管道的时间正好差1秒钟）
//打印the 4
</pre>
time.AfterFunc

time.AfterFunc(time.Duration,func());
和After差不多，意思是多少时间之后在goroutine line执行函数.
<pre>
package main

import "time"
import "fmt"

func main() {
	f := func() {
		fmt.Println("time out")
	}
	time.AfterFunc(time.Second, f)
	time.Sleep(2 * time.Second)
}

//【结果】运行了1秒后，打印出timeout，又过了1秒，程序退出
//将一个间隔和一个函数给AfterFunc后
//间隔时间过后，执行传入的函数
</pre>
time.Tick

每隔多少时间后
<pre>
package main

import "time"
import "fmt"

func main() {
	fmt.Println("the 1")
	tc := time.Tick(time.Second)

	for i := 1; i <= 5; i++ {
		<-tc
		fmt.Println("hello")
	}
}
/*
首先打印一个 the 1
然后每隔1秒，打印一个hello
*/
</pre>
Before & After方法、

判断一个时间点是否在另一个时间点的前面（后面），返回true或false
<pre>
t1:=time.Now()
time.Sleep(time.Second)
t2:=time.Now()
a:=t2.After(t1)     //t2的记录时间是否在t1记录时间的**后面**呢，是的话，a就是true
fmt.Println(a)       //true
b:=t2.Before(t1)     //t2的记录时间是否在t1记录时间的**前面**呢，是的话，b就是true
fmt.Println(b)       //false
</pre>
Sub方法

两个时间点相减，获得时间差（Duration）
<pre>
t1:=time.Now()
time.Sleep(time.Second)
t2:=time.Now()
d:=t2.Sub(t1)     //时间2减去时间1
fmt.Println(d)       //打印结果差不多为1.000123几秒，因为Sleep无法做到精确的睡1秒
后发生的时间  减去   先发生时间，是正数
</pre>
Add方法

拿一个时间点，add一个时长，获得另一个时间点
<pre>
t1:=time.Now()              //现在是12点整（假设）,那t1记录的就是12点整
t2:=t1.Add(time.Hour)          //那t1的时间点 **加上(Add)** 1个小时，是几点呢？
fmt.Println(t2)       //13点（呵呵）
</pre>
###Golang处理表单输入
<pre>
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数,对于POST则解析响应包的主体(request body)
	//注意:如果没有调用ParseForm方法,下面无法获取表单的数据
	/*fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	} */
	fmt.Fprintf(w, "Hello Jason!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		/*
			默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，
			你才能对这个表单数据进行操作。我们修改一下代码，
			在fmt.Println("username:", r.Form["username"])之前加一行r.ParseForm(),重新编译
		*/
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
//其中一种情况
method: GET
method: POST
username: [rrtrt]
password: [rtt]
</pre>
###验证表单的输入
验证项目包括

- 必填字段
- 是否数字
- 是否中文
- 英文
- 电子邮件地址
- 手机号码
- 下拉菜单
- 单选按钮
- 复选框
- 日期和时间
- 身份证号码

见 https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/04.2.md
###客户端上传文件
<pre>
package main

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"
)

func postFile(filename string, targetUrl string) error {
    bodyBuf := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuf)

    //关键的一步操作
    fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
    if err != nil {
        fmt.Println("error writing to buffer")
        return err
    }

    //打开文件句柄操作
    fh, err := os.Open(filename)
    if err != nil {
        fmt.Println("error opening file")
        return err
    }
    defer fh.Close()

    //iocopy
    _, err = io.Copy(fileWriter, fh)
    if err != nil {
        return err
    }

    contentType := bodyWriter.FormDataContentType()
    bodyWriter.Close()

    resp, err := http.Post(targetUrl, contentType, bodyBuf)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    resp_body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    fmt.Println(resp.Status)
    fmt.Println(string(resp_body))
    return nil
}
//sample usage
func main() {
    target_url := "http://localhost:9090/upload"
    filename := "./astaxie.pdf"
    postFile(filename, target_url)
}
</pre>
###数据库接口
在我们使用database/sql接口和第三方库的时候经常看到如下:
<pre>
  import (
      "database/sql"
      _ "github.com/mattn/go-sqlite3"
  )
</pre>
新手都会被这个 _ 所迷惑，其实这个就是Go设计的巧妙之处，我们在变量赋值的时候经常看到这个符号，它是用来忽略变量赋值的占位符，那么包引入用到这个符号也是相似的作用，这儿使用_的意思是引入后面的包名而不直接使用这个包中定义的函数，变量等资源。
###Golang Session Cookie 
- Cookie

Golang中通过net/http包中的SetCookie来设置：
<pre>
http.SetCookie(w ResponseWriter, cookie *Cookie)
</pre>
w表示需要写入的response，cookie是一个struct，让我们来看一下cookie对象是怎么样的
<pre>
type Cookie struct {
    Name       string
    Value      string
    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string

// MaxAge=0 means no 'Max-Age' attribute specified.
// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}
</pre>
我们来看一个例子，如何设置cookie
<pre>
expiration := time.Now()
expiration = expiration.AddDate(1, 0, 0)
cookie := http.Cookie{Name: "username", Value: "jason", Expires: expiration}
http.SetCookie(w, &cookie)
</pre>
Golang读取cookie
上面的例子演示了如何设置cookie数据，我们这里来演示一下如何读取cookie
<pre>
cookie, _ := r.Cookie("username")
fmt.Fprint(w, cookie)
</pre>
还有另外一种读取方式
<pre>
for _, cookie := range r.Cookies() {
    fmt.Fprint(w, cookie.Name)
}
</pre>
可以看到通过request获取cookie非常方便。

- Session 

session是在服务器端实现的一种用户和服务器之间认证的解决方案，目前Go标准包没有为session提供任何支持，这小节我们将会自己动手来实现go版本的session管理和创建。
<pre>
//created by astaxie
package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxlifetime int64)
}

var provides = make(map[string]Provider)

// Register makes a session provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, provide Provider) {
	if provide == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provide
}

type Manager struct {
	cookieName  string     //private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

//get Session
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

//Destroy sessionid
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime)*time.Second, func() { manager.GC() })
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
</pre>
###xml文件处理
servers.xml
<pre>
<?xml version="1.0" encoding="utf-8"?>
<servers version="1">
    <server>
        <serverName>Shanghai_VPN</serverName>
        <serverIP>127.0.0.1</serverIP>
    </server>
    <server>
        <serverName>Beijing_VPN</serverName>
        <serverIP>127.0.0.2</serverIP>
    </server>
</servers>
</pre>
处理代码是：
<pre>
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}
</pre>
###Json处理

- 解析到结构体
<pre>
package main
import "encoding/json"
import "fmt"
type Server struct {
	Servername string
	Serverip   string
}
type Serverslice struct {
	Servers []Server
}
func main() {
	var s Serverslice
	str := `{"Servers":[{"Servername":"shanghai","ServerIP":"127.0.0.1"},{"Servername":"beijing","Serverip":"127.0.0.3"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
output==>
{[{shanghai 127.0.0.1} {beijing 127.0.0.3}]}
</pre>

- 解析到interface

我们知道interface{}可以用来存储任意数据类型的对象，这种数据结构正好用于存储解析的未知结构的json数据的结果。JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。Go类型和JSON类型的对应关系如下：

- bool 代表 JSON booleans
- float64 代表 JSON numbers
- string 代表 JSON strings
- nil 代表 JSON null

对于未知结构的json，建议使用https://github.com/bitly/go-simplejson。

- 生成Json

我们开发很多应用的时候，最后都是要输出JSON数据串，那么如何来处理呢？JSON包里面通过Marshal函数来处理，函数定义如下：
<pre>
func Marshal(v interface{}) ([]byte, error)
</pre>
例子 ：
<pre>
package main

import "encoding/json"
import "fmt"

type Server struct {
	Servername string
	Serverip   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{Servername: "shanghai", Serverip: "1234.56.45.56"})
	s.Servers = append(s.Servers, Server{Servername: "beijing", Serverip: "55.87.67.8"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err :", err)
	}
	fmt.Println(string(b))
}
output==>
{"Servers":[{"Servername":"shanghai","Serverip":"1234.56.45.56"},{"Servername":"beijing","Serverip":"55.87.67.8"}]}
</pre>
###regexp 正则
使用正则来过滤或截取抓取到的百度搜索首页内容
<pre>
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}
output==>
百度一下，你就知道
输入法
手写
拼音
关闭
百度首页
设置
登录
糯米
新闻
hao123
地图
视频
贴吧
登录
设置
更多产品
网页
新闻
贴吧
知道
音乐
图片
视频
地图
文库
更多»
手机百度
快人一步
百度糯米
一元大餐
把百度设为主页
把百度设为主页
关于百度
About&nbsp;&nbsp;Baidu
&copy;2016&nbsp;Baidu&nbsp;
使用百度前必读
&nbsp;
意见反馈
&nbsp;京ICP证030173号&nbsp;
京公网安备11000002000001号
</pre>
<pre>
package main

import (
    "fmt"
    "regexp"
)

func main() {
    a := "I am learning Go language"

    re, _ := regexp.Compile("[a-z]{2,4}")

    //查找符合正则的第一个
    one := re.Find([]byte(a))
    fmt.Println("Find:", string(one))

    //查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
    all := re.FindAll([]byte(a), -1)
    fmt.Println("FindAll", all)

    //查找符合条件的index位置,开始位置和结束位置
    index := re.FindIndex([]byte(a))
    fmt.Println("FindIndex", index)

    //查找符合条件的所有的index位置，n同上
    allindex := re.FindAllIndex([]byte(a), -1)
    fmt.Println("FindAllIndex", allindex)

    re2, _ := regexp.Compile("am(.*)lang(.*)")

    //查找Submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，第三个是第二个()里面的
    //下面的输出第一个元素是"am learning Go language"
    //第二个元素是" learning Go "，注意包含空格的输出
    //第三个元素是"uage"
    submatch := re2.FindSubmatch([]byte(a))
    fmt.Println("FindSubmatch", submatch)
    for _, v := range submatch {
        fmt.Println(string(v))
    }

    //定义和上面的FindIndex一样
    submatchindex := re2.FindSubmatchIndex([]byte(a))
    fmt.Println(submatchindex)

    //FindAllSubmatch,查找所有符合条件的子匹配
    submatchall := re2.FindAllSubmatch([]byte(a), -1)
    fmt.Println(submatchall)

    //FindAllSubmatchIndex,查找所有字匹配的index
    submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
    fmt.Println(submatchallindex)
}
output==>
Find: am
FindAll [[97 109] [108 101 97 114] [110 105 110 103] [108 97 110 103] [117 97 103 101]]
FindIndex [2 4]
FindAllIndex [[2 4] [5 9] [9 13] [17 21] [21 25]]
FindSubmatch [[97 109 32 108 101 97 114 110 105 110 103 32 71 111 32 108 97 110 103 117 97 103 101] [32 108 101 97 114 110 105 110 103 32 71 111 32] [117 97 103 101]]
am learning Go language
 learning Go 
uage
[2 25 4 17 21 25]
[[[97 109 32 108 101 97 114 110 105 110 103 32 71 111 32 108 97 110 103 117 97 103 101] [32 108 101 97 114 110 105 110 103 32 71 111 32] [117 97 103 101]]]
[[2 25 4 17 21 25]]
</pre>
###sync.WaitGroup
<pre>
package main

import (
	"fmt"
	"sync"
)

/*
sync.WaitGroup
sync包中的WaitGroup实现了一个类似任务队列的结构，
你可以向队列中加入任务，任务完成后就把任务从队列中移除，
如果队列中的任务没有全部完成，队列就会触发阻塞以阻止程序继续运行，
具体用法参考如下代码：
*/
var waitgroup sync.WaitGroup

func Afunction(shownum int) {
	fmt.Println(shownum)
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}

func main() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go Afunction(i)
	}
	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
}
output==>
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
</pre>
###Beego相关
<pre>
									|model--->db.go-->|
地址栏访问->路由router->controller-->|          		  |----->Render html
									|---------------->|


1. router  形式

beego.Router("/base/addbincard", &controllers.BankInfoController{}, "get:BindCardGet;post:BindCardPost") //如果以get方式访问该方法调用BindCardGet方法，要是以post方式访问则调用BinCardPost方法

2. controller    形式

//Jason test page
func (this *AboutController) Jason() {
	this.Ctx.WriteString("Jason")
}

//Jason test page two
func (this *AboutController) Jason2() {
	//加载公用controller，比如对是否登录的判断等等基本的操作
	this.BaseController.Get()
	this.Data["header"] = 4
	this.TplName = "about/jason.html"
}

3. model 形式

func (this *JoinController) Join() {
	this.BaseController.Get()
	this.Data["header"] = 4
	links, _ := models.GetLinks()
	this.Data["links"] = links

	//招聘人员信息
	joinlist, err := models.GetJoinList() //models中db.go中的GetJoinList()方法，它返回数据集与错误信息
	if err == nil {
		this.Data["list"] = joinlist // joinlist是来自数据库的信息；list是传递到模板的东西
	}
	this.TplName = "about/join.html"
}


4. db

func GetJoinList() (v []ZcmNews, err error) {
	o := orm.NewOrm()
	sql := " SELECT *  from zcm_news a where a.cid = 58 and a.status = 1  ORDER BY create_time desc "
	_, err = o.Raw(sql).QueryRows(&v)
	if err != nil {
		return nil, err
	}
	return v, nil

}

5. 前端页面

<!DOCTYPE html>
<html class="screen_bg">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>加入我们_招财猫理财一家有诚意的互联网金融投资理财平台</title>
<meta name="keywords" content="招财猫招聘,招财猫理财招聘,招财猫招聘信息,招财猫求职"/>
<meta name="description" content="欢迎有志于投入互联网金融创业的有激情梦想的小伙伴加入我们，一起推动中国的金融服务创新，诚挚欢迎领域精英与我们一起，挥洒满腔热情、成就宏伟事业，我在招财猫理财等你！"/>

<link href="/static/css/base.css?y=20160418" rel="stylesheet"/>
<link href="/static/css/aboutus.css" rel="stylesheet" />
</head>

<body>
<!--header start-->
{{template "layout/header.tpl" .}}
<!--header end-->
<!--nav start -->
<div class="sub_nav">
    <ul class="w1000">
        <li><a href="/cooperate">共襄问鼎</a></li>
        <li><a href="/finance">融创未来</a></li>
        <li><a href="/about">公司简介</a></li>
        <li><a href="/branch">分支机构</a></li>
        <!-- <li><a href="">团队介绍</a></li> -->
        <li><a href="/partner">合作伙伴</a></li>
        <li><a href="/media">企业动态</a></li>
        <li><a href="/news">行业资讯</a></li>
        <li><a href="/contact">联系我们</a></li>
        <li><a href="/join" class="curr">加入我们</a></li>
    </ul>
</div>
<!--nav end -->
<!--main start-->
<div class="join_banner"></div>
<div class="w1000 join_us">

   {{range .list}}
    <dl>
        <dt><h1>{{.Title}}</h1><a>查看详情<i></i></a></dt>
        <dd class="join_cnt_hide">
            <span class="join_arrowtop"></span>
            <ul>
                <li>{{.Content}}</li>
                <li>
                    <a class="join_hide_box">收起</a>
                    <input type="button" onclick="location='mailto:hr@zcmlc.com'" value="投个简历" />
                </li>
            </ul>
        </dd>
    </dl>
   {{end}}
</div>
<!--main end-->
<!--footer start-->
{{template "layout/footer.tpl" .}}
<!--footer end-->
<script type="text/javascript" src="/static/js/about.js?y=20160316"></script>
</body>
</html>
</pre>
beego 中路由参数与表单参数

beego的路由映射支持灵活的结构，比如对于这种/blog/:catName可以表示的是某一个分类下的blog列表，那么这里的:catName就是路由参数；如果说我们要对这个分类下面的blog进行分页，想查看第10页的blog，那么我们的url可能变成了/blog/:catName?page=10这种格式，那么这里的page就是表单参数。表单参数既可以是GET类型的参数也可以是POST类型的参数，总之都叫做表单参数。

1. 获取路由参数的方法（可以使用下面的方法来获取路由参数）
<pre>
方法 	原型
GetInt 	func (c *Controller) GetInt(key string) (int64, error)
GetBool 	func (c *Controller) GetBool(key string) (bool, error)
GetFloat 	func (c *Controller) GetFloat(key string) (float64, error)
GetString 	func (c *Controller) GetString(key string) string
</pre>
2. 获取表单参数的方法

上面我们看过了获取路由参数的方法，这里我们再看一下获取表单参数的方法。在上面的获取路由参数的讲解最后，我们发现可以使用和上面相同的方法来获取表单参数。
<pre>
方法 	原型
GetInt 	func (c *Controller) GetInt(key string) (int64, error)         // value = GetInt(name)
GetBool 	func (c *Controller) GetBool(key string) (bool, error)     // value = GetBool(name)
GetFloat 	func (c *Controller) GetFloat(key string) (float64, error) // value = GetFloat(name)
GetString 	func (c *Controller) GetString(key string) string          // value = GetString(name)
GetStrings(key string) []string										   // value = GetStrings(name)
</pre>
验证很简单，使用这样的url：http://localhost:8080/blog/30/beego/true/98.45?page=10 和代码：
<pre>
page, _ := this.GetInt("page")       //page是name,得到的结果是value   (name=>value)
beego.Debug("Page", page)
</pre>
输出:
<pre>
2014/09/02 14:41:07 [D] Page 10
</pre>
将表单内容赋值到一个struct里

如果要把表单里的内容赋值到一个 struct 里，除了用上面的方法一个一个获取再赋值外，beego 提供了通过另外一个更便捷的方式，就是通过 struct 的字段名或 tag 与表单字段对应直接解析到 struct。

模板html里面：
<pre>
<form id="user" method="post">
    名字：<input name="username" type="text" />
    年龄：<input name="age" type="text" />
    邮箱：<input name="Email" type="text" />
    <input type="submit" value="提交" />
</form>
</pre>
constroller里面定义struct:
<pre>
type user struct {
    Id    int         `form:"-"`           // - 表示忽略（不使用的意思）
    Name  interface{} `form:"username"`
    Age   int         `form:"age"`
    Email string
}

//controller里面解析：
func (this *MainController) Post() {     
    u := user{}
    if err := this.ParseForm(&u); err != nil {
        //handle error
    }
}
</pre>
Beego中cookie操作 redirect跳转操作 WriteString Abort操作需要引入github.com/astaxie/beego/context包

context 对象是对 Input 和 Output 的封装，里面封装了几个方法：

- Redirect
- Abort
- WriteString
- GetCookie
- SetCookie


Input 对象 （很多方法，下面列出常用）

- Domain 请求的域名，例如 beego.me
- UserAgent 返回请求的 UserAgent，例如 Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36
- Param 在路由设置的时候可以设置参数，这个是用来获取那些参数的，例如 Param(":id"),返回12
- Query 该函数返回 Get 请求和 Post 请求中的所有数据，和 PHP 中 $_REQUEST 类似
- Cookie 返回请求中的 cookie 数据，例如 Cookie("username")，就可以获取请求头中携带的 cookie 信息中 username 对应的值
- Session session 是用户可以初始化的信息，默认采用了 beego 的 session 模块中的 Session 对象，用来获取存储在服务
器端中的数据。
- 。。。

Output 对象 （很多方法，下面列出常用）

- Cookie 设置输出的 cookie 信息，例如 Cookie("sessionID","beegoSessionID")
- Json 把 Data 格式化为 Json，然后调用 Body 输出数据
- Jsonp 把 Data 格式化为 Jsonp，然后调用 Body 输出数据
- Session 设置在服务器端保存的值，例如 Session("username","astaxie")，这样用户就可以在下次使用的时候读取
- 。。。


Session操作可以引入github.com/astaxie/beego/session包，使用方法见http://beego.me/docs/module/session.md

在beego框架中直接在controller中使用GetSession SetSession DelSession 就能完成基本的操作。同样的，cookie操作也是类似，见下例子:
<pre>
//设置cookie
	this.Ctx.SetCookie("guess", "guesscookie")
	//清除cookie
	this.Ctx.SetCookie("guess", "0", -1)
	//读取cookie
	guess := this.Ctx.GetCookie("guess")
</pre>
json Marshal
<pre>
package main
import (
    "encoding/json"
    "fmt"
    "os"
)
func main ( ) {
    type ColorGroup struct {
        ID     int
        Name   string
        Colors [ ] string
    }
    group := ColorGroup {
        ID :     1 ,
        Name :   "Reds" ,
        Colors : [ ] string { "Crimson" , "Red" , "Ruby" , "Maroon" } ,
    }
    b , err := json. Marshal ( group )
    if err != nil {
        fmt. Println ( "error:" , err )
    }
    os. Stdout . Write ( b )
}
output==>
	{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
</pre>
json Unmarshal
<pre>
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(` [
        { "Name" : "Platypus" , "Order" : "Monotremata" } ,
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" }
    ] `)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
output==>
	[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
</pre>
beego中存储日志信息
<pre>
//将项目中所有beego.Emergency/beego.Warning等信息保存在logs/log.log文件中
beego.BeeLogger.SetLogger("file", `{"filename": "logs/log.log"}`)
</pre>
###beego中设置404页面
在入口页面main.go中编辑：
<pre>
var FilterPermission = func(ctx *context.Context) {
	id := ctx.Input.Session("loginid")
	if id == nil && ctx.Request.RequestURI != "/" {
		ctx.Redirect(302, "/")
	}
}

var FilterWap = func(ctx *context.Context) {
	//检查cookie
	var value = ctx.GetCookie("FromWap")
	if value != "true" {
		agent := strings.ToLower(ctx.Request.Header.Get("User-Agent"))
		var agents = []string{"android", "iphone", "ipad", "windows phone", "ipod", "blackberry", "mobile"}
		var mobile = false
		for _, a := range agents {
			if strings.Contains(agent, a) {
				mobile = true
				break
			}
		}
		if mobile {
			//移动端处理
			var wap = ctx.Request.FormValue("fromwap")
			if wap != "1" {
				//跳转到wap
				ctx.Redirect(302, "http://???.com")
			} else {
				//如果是wap端并且是跳转过来的则设置Cookie
				ctx.SetCookie("FromWap", "true")
			}
		}
	}
}
func main() {
	beego.SetStaticPath("/static/*", "/static")
	beego.InsertFilter("/user/recharge/*", beego.BeforeRouter, FilterPermission) //过滤
	beego.InsertFilter("/", beego.BeforeRouter, FilterWap)                       //wap访问过滤
	beego.ErrorHandler("404", page_not_found)                                    //404跳转
	beego.ErrorController(&controllers.ErrorController{})                        //自定义错误处理
	beego.Run()
}
//404处理
func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
</pre>
####beego内置模板函数
目前beego内置的模板函数有如下：

- markdown

实现了把markdown文本转化为html信息，使用方法{{markdown .Content}}

- dateformat

实现了时间的格式化，返回字符串，使用方法{{dateformat .Time "2006-01-02T15:04:05Z07:00"}}

- date

实现了类似PHP的date函数，可以很方便的根据字符串返回时间，使用方法{{date .T "Y-m-d H:i:s"}}

- compare

实现了比较两个对象的比较，如果相同返回true，否者false，使用方法{{compare .A .B}}

- substr

实现了字符串的截取，支持中文截取的完美截取，使用方法{{substr .Str 0 30}}

- html2str

实现了把html转化为字符串，剔除一些script、css之类的元素，返回纯文本信息，使用方法{{html2str .Htmlinfo}}

- str2html

实现了把相应的字符串当作HTML来输出，不转义，使用方法{{str2html .Strhtml}}

- htmlquote

实现了基本的html字符转义，使用方法{{htmlquote .quote}}

- htmlunquote

实现了基本的反转移字符，使用方法{{htmlunquote .unquote}}


###reflect
reflect包有两个数据类型我们必须知道，一个是Type，一个是Value。

Type就是定义的类型的一个数据类型，Value是值的类。
<pre>
package main

import (
	"fmt"
	"reflect"
)

type Mystruct struct {
	name string
}

func (this *Mystruct) Getname() string {
	return this.name
}
func (this *Mystruct) Gettwo() string {
	return this.name + "hello"
}

func main() {
	s := "this is a string"
	fmt.Println("s`s type  is:", reflect.TypeOf(s))
	fmt.Println("s`s value is:", reflect.ValueOf(s))

	var x int = 3
	fmt.Println("x`s type is:", reflect.TypeOf(x))
	fmt.Println("x`s value is:", reflect.ValueOf(x))

	a := new(Mystruct)
	a.name = "Jason"
	typ := reflect.TypeOf(a)
	fmt.Println("typ`s type is:", reflect.TypeOf(typ))
	fmt.Println("typ`s value is:", reflect.ValueOf(typ))
	fmt.Println("typ`s NumMethod is:", typ.NumMethod())//属于typ的方法的总个数（这个太6了),分别是Getname与Gettwo
}
output==>
s`s type  is: string
s`s value is: this is a string
x`s type is: int
x`s value is: 3
typ`s type is: *reflect.rtype
typ`s value is: &{8 8 1536160849 0 8 8 54 0x5a0930 0x543f50 0x51af60 0x4f1708 <nil>}
typ`s NumMethod is: 2 
</pre>
reflect.ValueOf().FieldByName ,获取结构体内某一个属性的值
<pre>
package main

import (
	"fmt"
	"reflect"
)

type Mystruct struct {
	name string
	age  int
}

func (this *Mystruct) Getname() string {
	return this.name
}

func main() {
	var a Mystruct
	b := new(Mystruct)
	fmt.Println("a value:", reflect.ValueOf(a))
	fmt.Println("b value:", reflect.ValueOf(b))

	a.name = "jack"
	a.age = 4
	b.name = "jason"

	val := reflect.ValueOf(a).FieldByName("age")
	fmt.Println("a FieldByName:", val)
}
output==>
a value: { 0}
b value: &{ 0}
a FieldByName: 4
</pre>

<pre>
package main

import (
	"fmt"
	"reflect"
)

type Info struct {
	name string `abc:"type,attr,omitempty" nnn:"xxx"`
	//pass struct{} `test`
}

func main() {
	info := Info{"hello"}
	ref := reflect.ValueOf(info)
	fmt.Println(ref.Kind())
	fmt.Println(reflect.Interface)
	fmt.Println(ref.Type())
	typ := reflect.TypeOf(info)
	n := typ.NumField()
	for i := 0; i < n; i++ {
		f := typ.Field(i)
		fmt.Println(f.Tag)
		fmt.Println(f.Tag.Get("nnn"))
		fmt.Println(f.Name)
	}
}
output==>
struct
interface
main.Info
abc:"type,attr,omitempty" nnn:"xxx"
xxx
name
</pre>
#####beego中使用beego/orm获取一个数据集的例子(多个数据)
models.go
<pre>
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)
type JasonTest struct {
	Id    int    `orm:"column(id);pk"`
	Name  string `orm:"column(name);size(20)"`
	Phone string `orm:"column(phone);size(100)"`
}

/*视情况而定
func init() {
	orm.RegisterModel(new(JasonTest))
}*/

//获取数据集
func GetNamesBySort(condition string) (names []JasonTest, err error) {
	o := orm.NewOrm()
	sql := "select name from jason_test where id>?"
	_, err = o.Raw(sql, condition).QueryRows(&names)
	fmt.Println(sql, err)
	return names, err
}
</pre>
controllers.go
<pre>
import "`/models"
func (this *...) ...(){
	results ,_:= models.GetNamesBySort("1")
	fmt.Println("所查询到的数据集的个数为：",len(results))
	var stringword string
	for k, _ := range results {               //将结构体形式的数据集转成字符串类型
		stringword += results[k].Name + "|"
	}
} 
</pre>
#####beego中使用beego/orm更新数据
models.go
<pre>
import "..."
type JasonTest struct {
	Id    int    `orm:"column(id);pk"`
	Name  string `orm:"column(name);size(20)"`
	Phone string `orm:"column(phone);size(100)"`
}
func init() {
	orm.RegisterModel(new(Jack))   //1.注册模型是必须的(注册模型：如果使用 orm.QuerySeter 进行高级查询的话，这个是必须的。反之，如果只使用Raw查询和 map struct，是无需这一步的。)
}
//这个方法会默认更新一条数据的所有字段
func UpdateAllData(jt *JasonTest)(int64,error){
	o:= orm.NewOrm()	
	num ,err := o.Update(jt)   	  //2.这里不能是&jt(指针的指针)
	return num ,err
}
//这个方法只会更新指定字段(这里只会更新name字段的值)
func UpdateOneData(jt *JasonTest)(int64,error){
	o:= orm.NewOrm()	
	num ,err := o.Update(jt,"name")
	return num ,err
}
</pre>
controllers.go
<pre>
import "..."
var jasontest models.JasonTest
jasontest.Id = 5                 //3.这里的Id（表的主键绝对不能为空）
jasontest.Name = "all"
jasontest.Phone = "11111"
_, err := models.UpdateAllData(&jasontest)
if err == nil {
	words += "所有字段更新成功"
} else {
	words += "所有字段更新失败"
}
_, err2 := models.UpdateOneData(&jasontest)
if err2 == nil {
	words += "name更新成功"
} else {
	words += "name更新失败"
}
</pre>
#####beego中使用beego/orm插入数据
models.go
<pre>
import "..."
type JasonTest struct {
	Id    int    `orm:"column(id);pk"`
	Name  string `orm:"column(name);size(20)"`
	Phone string `orm:"column(phone);size(100)"`
}
func InsertIntoJason(a int, b, c string) (err error) {
	o := orm.NewOrm()
	sql := "insert into jason_test(id,name,phone) values(?,?,?)"
	_, err = o.Raw(sql, a, b, c).Exec()
	fmt.Println(sql, err)
	return err
}
</pre>
controllers.go
<pre>
err := models.InsertIntoJason(9, "ll", "2433445")
if err == nil {
	words +="添加成功"
} else {
	words +="添加失败"
}
</pre>
###Beego中使用事务
<pre>
o := NewOrm()
err := o.Begin()
// 事务处理过程
...
...
// 此过程中的所有使用 o Ormer 对象的查询都在事务处理范围内
if SomeError {
    err = o.Rollback()
} else {
    err = o.Commit()
}
</pre>
### Socket编程
多并发执行,当有新的客户端请求到达并同意接受Accept该请求的时候他会反馈当前的时间信息。值得注意的是，在代码中for循环里，当有错误发生时，直接continue而不是退出，是因为在服务器端跑代码的时候，当有错误发生的情况下最好是由服务端记录错误，然后当前连接的客户端直接报错而退出，从而不会影响到当前服务端运行的整个服务。
<pre>
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime)) // don't care about return value
	// we're finished with this client
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
</pre>
###gob/Gob/GOB
<pre>
package main

/*
1 P和Q是两个结构体，应该说是“相似”的两个结构体
2 Encode是将结构体传递过来，但是Decode的函数参数却是一个pointer！

gob包是Go提供的"私有"的编解码方式，文档中也说了它的效率会比json，
xml等更高（虽然我也没有验证）。因此在两个Go服务之间的相互通信建议
不要再使用json传递了，完全可以直接使用gob来进行数据传递。
*/
import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Jason"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println(q)
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)

}
output==>
{0xc08202ab18 0xc08202ab1c Jason}
"Jason": {3,4}
</pre>
###Beego中正则路由的使用（实现一个路由下多个数据的逐一展示）
router.go
<pre>
	beego.Router("/jack2/:id([0-9]+)", &controllers.JackController{}, "get:Jack2")   //这里将get:Jack2改为*:Jack2，则无论何种访问方式都可以，不然只能通过get方式访问
</pre>
controller.go
<pre>
func (this *JackController) Jack2() {
	id := this.Ctx.Input.Param(":id") //这里的 :id  要与router里面的 :id([0-9]+) 对应
	idd, _ := strconv.Atoi(id)
	reult, err := models.ShowEveryData(idd)
	fmt.Println(idd)
	if err == nil && reult != nil {
		fmt.Println(reult.Name)
		Age := strconv.Itoa(reult.Age)
		this.Ctx.WriteString(reult.Name + "-" + Age)
	} else {
		this.Ctx.WriteString("出错了")
	}
}
</pre>
model.go
<pre>
type Jack struct {
	Id   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
	Age  int    `orm:"column(age)"`
}
//实现根据id访问每条数据
func ShowEveryData(id int) (data *Jack, err error) {
	o := orm.NewOrm()
	sql := "select * from jack where id = ?"
	err = o.Raw(sql, id).QueryRow(&data)
	if err != nil {
		fmt.Println("results:", sql, err)
	}
	return data, err
}
</pre>
###RESTful
<pre>
package main

import (
    "fmt"
    "github.com/drone/routes"
    "net/http"
)

func getuser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are get user %s", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are modify user %s", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are delete user %s", uid)
}

func adduser(w http.ResponseWriter, r *http.Request) {
    uid := r.FormValue("uid")
    fmt.Fprint(w, "you are add user %s", uid)
}

func main() {
    mux := routes.New()
    mux.Get("/user/:uid", getuser)
    mux.Post("/user/", adduser)
    mux.Del("/user/:uid", deleteuser)
    mux.Put("/user/:uid", modifyuser)
    http.Handle("/", mux)
    http.ListenAndServe(":8088", nil)
}
</pre>
###加密与解密
base64
<pre>
package main

//base64加密与解密
import "encoding/base64"
import "fmt"

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	hello := "hello world"
	debyte := base64Encode([]byte(hello))
	fmt.Println(debyte)
	fmt.Println(string(debyte))

	enbyte, err := base64Decode(debyte)
	if err == nil {
		fmt.Println(string(enbyte))
	}
}
output==>
[97 71 86 115 98 71 56 103 100 50 57 121 98 71 81 61]
aGVsbG8gd29ybGQ=
hello world
</pre>
高级加解密 aes/AES des/DES
<pre>
//des与aes用法类似，这里讲的是aes
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	//需要去加密的字符串
	plaintext := []byte("I am jason")
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	//aes的加密字符串
	key_text := "jason78gv798akljzmknm.ahkjkljl;k" //长度固定
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}

	fmt.Println(len(key_text))

	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}
output==>
32
I am jason=>4cb0daaa99eb2483e0bb
4cb0daaa99eb2483e0bb=>I am jason
</pre>
本地化资源
<pre>
package main

import "fmt"

var locales map[string]map[string]string

func main() {
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	locales["en"] = en
	cn := make(map[string]string, 10)
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	locales["zh-CN"] = cn
	lang := "zh-CN"
	fmt.Println(msg(lang, "pea"))
	fmt.Println(msg(lang, "bean"))
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}
output==>
豌豆
毛豆
</pre>
上面示例演示了不同locale的文本翻译，实现了中文和英文对于同一个key显示不同语言的实现，上面实现了中文的文本消息，如果想切换到英文版本，只需要把lang设置为en即可。
###自定义用户认证（登录或注册）
自定义的认证一般都是和session结合验证的。
<pre>
//登陆处理
func (this *LoginController) Post() {
    this.TplNames = "login.tpl"
    this.Ctx.Request.ParseForm()
    username := this.Ctx.Request.Form.Get("username")
    password := this.Ctx.Request.Form.Get("password")
    md5Password := md5.New()
    io.WriteString(md5Password, password)
    buffer := bytes.NewBuffer(nil)
    fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
    newPass := buffer.String()

    now := time.Now().Format("2006-01-02 15:04:05")

    userInfo := models.GetUserInfo(username)
    if userInfo.Password == newPass {
        var users models.User
        users.Last_logintime = now
        models.UpdateUserInfo(users)

        //登录成功设置session
        sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
        sess.Set("uid", userInfo.Id)
        sess.Set("uname", userInfo.Username)

        this.Ctx.Redirect(302, "/")
    }   
}

//注册处理
func (this *RegController) Post() {
    this.TplNames = "reg.tpl"
    this.Ctx.Request.ParseForm()
    username := this.Ctx.Request.Form.Get("username")
    password := this.Ctx.Request.Form.Get("password")
    usererr := checkUsername(username)
    fmt.Println(usererr)
    if usererr == false {
        this.Data["UsernameErr"] = "Username error, Please to again"
        return
    }

    passerr := checkPassword(password)
    if passerr == false {
        this.Data["PasswordErr"] = "Password error, Please to again"
        return
    }

    md5Password := md5.New()
    io.WriteString(md5Password, password)
    buffer := bytes.NewBuffer(nil)
    fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
    newPass := buffer.String()

    now := time.Now().Format("2006-01-02 15:04:05")

    userInfo := models.GetUserInfo(username)

    if userInfo.Username == "" {
        var users models.User
        users.Username = username
        users.Password = newPass
        users.Created = now
        users.Last_logintime = now
        models.AddUser(users)

        //登录成功设置session
        sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
        sess.Set("uid", userInfo.Id)
        sess.Set("uname", userInfo.Username)
        this.Ctx.Redirect(302, "/")
    } else {
        this.Data["UsernameErr"] = "User already exists"
    }

}

func checkPassword(password string) (b bool) {
    if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
        return false
    }
    return true
}

func checkUsername(username string) (b bool) {
    if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
        return false
    }
    return true
}
</pre>
###sync.Mutex加锁
<pre>
package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex *sync.Mutex

func lock(i int) {
	fmt.Println(i, " lock start")
	mutex.Lock()
	fmt.Println(i, "lock")
	time.Sleep(5 * time.Second)
	mutex.Unlock()
	fmt.Println(i, "unlock")
}

func main() {
	mutex = new(sync.Mutex)
	go lock(1)
	time.Sleep(time.Second)
	lock(2)
	fmt.Println("exit")
}
output==>
1  lock start
1 lock
2  lock start
1 unlock
2 lock
2 unlock
exit
</pre>
sync.WaitGroup

例子1
<pre>
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("Jason", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
output==>
Jason 9
Jason 0
Jason 1
Jason 2
Jason 3
Jason 4
Jason 5
Jason 6
Jason 7
Jason 8
</pre>
例子2
<pre>
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://baidu.com", "http://sina.com", "http://163.com",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Println(url)
		}(url)		
	}
	wg.Wait()
	fmt.Println("over!")
}
output==>
http://163.com
http://baidu.com
http://sina.com
over!
</pre>
http.Get()
<pre>
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := new(http.Client)
	reg, err := http.NewRequest("GET", "http://sina.com", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	reg.Header.Set("HTTP", "2.0")
	reg.Header.Set("Accept", "*/*")
	reg.Header.Set("Accept-Language", "zh-cn")
	reg.Header.Set(`User-Agent`, `AppStore/2.0 iOS/7.1.2 model/iPod5,1 build/11D257 (4; dt:81)`)
	reg.Header.Set(`Host`, `sina.com`)
	reg.Header.Set(`Connection`, `keep-alive`)
	reg.Header.Set(`X-Apple-Store-Front`, `143465-19,21 t:native`)
	reg.Header.Set(`X-Dsid`, `932530590`)

	resp, err := client.Do(reg)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}
output==>
......(so many contents)
</pre>
sync.Once()
<pre>
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func GetDemo(addr string) {
	res, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(robots))
}

func ExampleWaitGroup() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://sina.com",
		"http://sohu.com",
		"http://baidu.com",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			GetDemo(url)
		}(url)
	}
	wg.Wait()
	fmt.Println("-----------Group wait over---------------")
}

func ExampleOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	fmt.Println("-----------------Once over --------------")
}

func main() {
	ExampleOnce()
	ExampleWaitGroup()
}
output==>
...(so many contents...)
</pre>
###atomic原子操作
原子操作即是进行过程中不能被中断的操作。针对某个值的原子操作在被进行的过程中，CPU绝不会再去进行其他的针对该值的操作。为了实现这样的严谨性，原子操作仅会由一个独立的CPU指令代表和完成。

Golang提供的原子操作都是非入侵式的，由标准库sync/atomic中的众多函数代表类型包括int32,int64,uint32,uint64,uintptr,unsafe.Pointer，共六个。这些函数提供的原子操作共有五种：增或减，比较并交换，载入，存储和交换。
<pre>
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i32 int32
	fmt.Println("old i32 value===")
	fmt.Println(i32)
	newi32 := atomic.AddInt32(&i32, 3)
	fmt.Println("new i32 value===")
	fmt.Println(i32)
	fmt.Println(newi32)
}
output==>
old i32 value===
0
new i32 value===
3
3
</pre>
高并发下atomic
<pre>
package main

import (
        "fmt"
        "sync/atomic"
        "time"
)

func main() {

        var cnt uint32 = 0

        // 启动10个goroutine
        for i := 0; i < 10; i++ {
                go func() {
                        // 每个goroutine都做20次自增运算
                        for i := 0; i < 20; i++ {
                                time.Sleep(time.Millisecond)
                                atomic.AddUint32(&cnt, 1)
                        }
                }()
        }

        // 等待2s, 等goroutine完成
        time.Sleep(time.Second * 2)
        // 取最终结果
        cntFinal := atomic.LoadUint32(&cnt)

        fmt.Println("cnt:", cntFinal)
}
output==>
cnt=200
</pre>
卖票的故事 ~ ~ ~
<pre>
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var total_tickets int32 = 10
var mutex = &sync.Mutex{}

func sell_tickets(i int) {

	for total_tickets > 0 {

		mutex.Lock()
		// 如果有票就卖
		if total_tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			// 卖一张票
			total_tickets--
			fmt.Println("id:", i, " ticket:", total_tickets)
		}
		mutex.Unlock()
	}
}

func main() {

	// 设置真正意义上的并发
	runtime.GOMAXPROCS(4)

	// 生成随机种子
	rand.Seed(time.Now().Unix())

	// 并发5个goroutine来卖票
	for i := 0; i < 5; i++ {
		go sell_tickets(i)
	}

	// 等待线程执行完
	var input string
	fmt.Scanln(&input)
	// 退出时打印还有多少票
	fmt.Println(total_tickets, "done")
}
output==>
id: 1  ticket: 9
id: 1  ticket: 8
id: 1  ticket: 7
id: 1  ticket: 6
id: 1  ticket: 5
id: 1  ticket: 4
id: 1  ticket: 3
id: 1  ticket: 2
id: 1  ticket: 1
id: 1  ticket: 0
</pre>
###strconv中的类型转换
<pre>
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//ParseBool 将字符串转换为布尔值
	fmt.Println(strconv.ParseBool("1"))     //true
	fmt.Println(strconv.ParseBool("t"))     //true
	fmt.Println(strconv.ParseBool("true"))  //true
	fmt.Println(strconv.ParseBool("0"))     //false
	fmt.Println(strconv.ParseBool("f"))     //false
	fmt.Println(strconv.ParseBool("false")) //false

	// FormatBool 将布尔值转换为字符串 "true" 或 "false"
	fmt.Println(strconv.FormatBool(0 < 1)) //true
	fmt.Println(strconv.FormatBool(0 > 1)) // false

	// ParseFloat 将字符串转换为浮点数
	s := "0.65645434345"
	fmt.Println("f`s type was:", reflect.TypeOf(s))
	f, _ := strconv.ParseFloat(s, 64)
	fmt.Println("f`s type is:", reflect.TypeOf(f))

	// ParseInt 将字符串转换为 int 类型
	fmt.Println(strconv.ParseInt("123", 10, 8))
	dd, _ := strconv.ParseInt("123", 10, 8)
	fmt.Println("dd`s type is:", reflect.TypeOf(dd)) //int64

	// Atoi 相当于 ParseInt(s, 10, 0)
	fmt.Println(strconv.Atoi("234546764"))
	at, _ := strconv.Atoi("234546764")
	fmt.Println("at`s type is:", reflect.TypeOf(at)) //int

	// FormatFloat 将浮点数 f 转换为字符串值
	fot := 100.12345678901234567890123456789
	fmt.Println(strconv.FormatFloat(fot, 'b', 5, 32))
	fotconv := strconv.FormatFloat(fot, 'b', 5, 32)
	fmt.Println("fotconv`s type is:", reflect.TypeOf(fotconv)) //string

	//FormatInt将 int 型整数 i 转换为字符串形式
	// Itoa 相当于 FormatInt(i, 10)
	fmt.Println(strconv.Itoa(2048))
	itoaa := strconv.Itoa(2048)
	fmt.Println("itoa`s type is:", reflect.TypeOf(itoaa)) //string

	//Quote将字符串转换成 "双引号" 包裹起来的字符串
	fmt.Println(strconv.Quote(`Jason`)) //"Jason"

	// AppendQuoteToASCII 将字符串 s 转换为“双引号”引起来的 ASCII 字符串
	stas := "hello,世界"
	stab := make([]byte, 0)
	stab = strconv.AppendQuoteToASCII(stab, stas)
	fmt.Printf("%s", stab) //"hello,\u4e16\u754c"

}
output==>
true <nil>
true <nil>
true <nil>
false <nil>
false <nil>
false <nil>
true
false
f`s type was: string
f`s type is: float64
123 <nil>
dd`s type is: int64
234546764 <nil>
at`s type is: int
13123382p-17
fotconv`s type is: string
2048
itoa`s type is: string
"Jason"
"hello,\u4e16\u754c"
</pre>
###将其他类型的变量转换成字符串类型
<pre>
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ToString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func main() {
	var s int = 55555
	fmt.Println(ToString(s))
	fmt.Println("s`s type is:", reflect.TypeOf(ToString(s)))
}
output==>
55555
s`s type is: string
</pre>
###for循环
<pre>
package main

import "fmt"

//for 循环
func main() {
	//第一种,单一条件循环
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	//第二种，经典的循环条件初始化/条件判断/循环后条件变化
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//第三种，无条件for循环是死循环，除非用break跳出或则return从函数返回
	for {
		fmt.Println("loop")
		fmt.Println("loop")
		fmt.Println("loop")
		return //break也可以
	}
}
output==>
1
2
3
7
8
9
loop
loop
loop
</pre>
os.Expand()
<pre>
package main

import (
	"fmt"
	"os"
)

//将字符串中的特定变量替换
func main() {
	mapping := func(s string) string {
		m := map[string]string{"X": "XXXXXXXX", "Y": "YYYYYY"}
		return m[s]
	}
	data := "hello $X and $Y"
	fmt.Printf("%s", os.Expand(data, mapping)) //输出hello widuu blog address www.widuu.com
}
output==>
hello XXXXXXXX and YYYYYY
</pre>
###Golang自定义类型
<pre>
package main

import (
	"fmt"
	"strings"
)

type Server struct {
	Name string
}

type Servers []Server

func ListenServers() Servers {
	return []Server{
		{Name: "jason"},
		{Name: "jack"},
		{Name: "kitty"},
	}
}

func (s Servers) Filter(name string) Servers {
	filtered := make(Servers, 0)
	for _, server := range s {
		if strings.Contains(server.Name, name) {
			filtered = append(filtered, server)
			fmt.Printf("%v\n", &filtered)
		}
	}
	return filtered
}

func main() {
	servers := ListenServers()
	servers = servers.Filter("")
	fmt.Printf("servers %+v\n", servers)
}
output==>
&[{jason}]
&[{jason} {jack}]
&[{jason} {jack} {kitty}]
servers [{Name:jason} {Name:jack} {Name:kitty}]
</pre>
简化上面的可以写成：
<pre>
package main

import "fmt"

type Server struct {
	name string
}

type Servers []Server

func returnServers() Servers {
	return []Server{
		{name: "jason"},
		{name: "jack"},
		{name: "jakiry"},
	}
}

func main() {
	fmt.Println(returnServers())
}
output==>
[{jason} {jack} {jakiry}]
</pre>
###Golang异步编程
<pre>
package main

import (
	"log"
	"math/rand"
	"time"
)

func UploadNetvalueFile(done chan bool) {
	//利用随机函数模拟不同文件的处理时间
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Intn(20)
	log.Println("UploadNetvalueFile: ", x)
	time.Sleep(time.Second * time.Duration(x))
	log.Println("UploadNetvalueFile OK")
	done <- true

}

func UplaodSaleFareFile(done chan bool) {
	//利用随机函数模拟不同文件的处理时间
	r := rand.New(rand.NewSource(time.Now().UnixNano() + 10))
	x := r.Intn(5)
	log.Println("UplaodSaleFareFile: ", x)
	time.Sleep(time.Second * time.Duration(x))
	log.Println("UplaodSaleFareFile OK ")
	done <- true

}

func main() {
	t1 := time.Now()
	tasknum := 2

	done := make(chan bool, tasknum)

	// 模拟生成文件的时间
	time.Sleep(time.Second * 1)

	log.Println("create netvalue file !")

	// 并发执行

	go UploadNetvalueFile(done)

	// 模拟生成文件的时间
	time.Sleep(time.Second * 1)

	log.Println("create salefare file !")

	// 并发执行

	go UplaodSaleFareFile(done)

	//等待所有并发完成
	func() {
		for i := 0; i < tasknum; i++ {
			log.Println(<-done)
		}

	}()

	t2 := time.Now()
	log.Printf("耗时:%d\n", t2.Sub(t1))
}
output==>
2016/05/31 20:10:38 create netvalue file !
2016/05/31 20:10:38 UploadNetvalueFile:  6
2016/05/31 20:10:39 create salefare file !
2016/05/31 20:10:39 UplaodSaleFareFile:  2
2016/05/31 20:10:41 UplaodSaleFareFile OK
2016/05/31 20:10:41 true
2016/05/31 20:10:44 UploadNetvalueFile OK
2016/05/31 20:10:44 true
2016/05/31 20:10:44 耗时:7029402100
</pre>
###Golang一个方法实现对数字或者字符串进行排序
<pre>
package main

import "fmt"

type Xi []int
type Xs []string

func (xi Xi) Len() int {
	return len(xi)
}

func (xi Xi) Less(i, j int) bool {
	return xi[j] < xi[i]
}
func (xi Xi) Swap(i, j int) {
	xi[i], xi[j] = xi[j], xi[i]
}

func (xs Xs) Len() int {
	return len(xs)
}

func (xs Xs) Less(i, j int) bool {
	return xs[j] < xs[i]
}

func (xs Xs) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

func main() {
	ints := Xi{67, 8, 9, 56, 8, 8, 990, 54, 5667, 788, 99, 674, 5, 7, 845}
	strings := Xs{"jack", "hrrgy", "abi", "frog"}
	Sort(ints)
	fmt.Println(ints)
	Sort(strings)
	fmt.Println(strings)
}
output==>
[5 7 8 8 8 9 54 56 67 99 674 788 845 990 5667]
[abi frog hrrgy jack]
</pre>
###Golang不定二维数组
<pre>
package main

import "fmt"

func main() {
	s := [...][2]string{{"ss", "dd"}, {"gg", "hh"}, {"kk", "jj"}, {"55", "88"}}
	b := s
	for _, v := range b {
		for _, t := range v {
			fmt.Println(t)
		}
	}
}
output==>
ss
dd
gg
hh
kk
jj
55
88
</pre>
###Litte Tools 小工具
<pre>
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

//对字符串进行SHA1哈希
func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串进行MD5编码处理
func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//隐藏手机号中间4位
func HidePhoneNumber(number string) string {
	return number[:3] + "****" + number[7:]
}

func main() {
	s := "hello"
	fmt.Println(Sha1(s))
	fmt.Println(Md5(s))
	phone := "15567874567"
	fmt.Println(HidePhoneNumber(phone))
}
output==>
aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
5d41402abc4b2a76b9719d911017c592
155****4567
</pre>
###万能类型 interface{} 与 .(string)的使用 |类型断言
<pre>
package main

import (
	"fmt"
	"strings"
)

//万能类型(interface{})很神奇
func getname(pa ...interface{}) {
	var pasl []string
	for _, p := range pa {
		pasl = append(pasl, p.(string))
	}
	aa := strings.Join(pasl, "_")
	fmt.Println(aa)
}

func main() {
	getname("hello", "jason", "I", "am", "world")
}
output==>
hello_jason_I_am_world
</pre>
下面是利用类型断言来求和的例子：
<pre>
package main

import "fmt"

func Add(f ...interface{}) int {
	var ff int
	for _, v := range f {
		ff += v.(int)
	}
	return ff
}

func main() {
	fmt.Println(Add(5, 6, 7, 78, 4, 8, 0))
}
output==>
108
</pre>
类型断言。其实Golang中的所有类型包括string.int.struct等都是默认继承了interface{}，所以interface{}也就是万能类型的来历。现在有一个问题：如何将一个interface{}形式的变量转换成string。那么这里就必须用到类型断言。见下：
<pre>
package main

func main() {
	var val interface{} = "hello"
	println(val.(string))    //这样子用滴···，那么输出的是string形式，如果val.(int)那么就是int类型
}
output==>
hello
</pre>
###Golang截断float string类型数据
<pre>
package main

import "fmt"
import "strings"
import "strconv"

var whoareyou = make(map[string]string)

func WhoAreYou(account string) string {
	key := string(account[:3])
	return key
}

//对字符串进行截取
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

//格式化float类型，截断，不四舍五入
func FormatFloat64(f float64, l int) float64 {
	str1 := fmt.Sprintf("%f", f)
	strl := strings.Split(str1, ".")
	strl[1] = Substr(strl[1], 0, l)
	strre := strl[0] + "." + strl[1]
	fre, _ := strconv.ParseFloat(strre, 64)
	return fre
}

func main() {
	fmt.Println(WhoAreYou("13588484126"))
	fmt.Println(FormatFloat64(45.04676, 2))
}
output==>
135
45.04
</pre>
进行四舍五入的：
<pre>
package main
 
import (
    "fmt"
)
 
func main() {
    var fs []float64 = []float64{1.1234456, 1.1234567, 1.1234678, 1.1}
    for _, f := range fs {
        s := fmt.Sprintf("%.5f", f)
        fmt.Println(f, "->", s)
    }
}
output==>
1.1234456 -> 1.12345
1.1234567 -> 1.12346
1.1234678 -> 1.12347
1.1 -> 1.10000
</pre>
###发送Get请求
<pre>
package main

import "fmt"
import "net/http"
import "io/ioutil"
import "errors"
import "encoding/base64"

const (
	base64Table = "173QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

//发送get请求
func SendGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.Status == "404 Not Found" {
		return nil, errors.New("服务器异常!")
	}
	if bodyByte, err := ioutil.ReadAll(resp.Body); err == nil {
		return base64Encode(bodyByte), err //得到经过base64编码处理过的返回的数据
	} else {
		return nil, err
	}
}

func main() {
	result, _ := SendGet("http://www.sina.com")
	fmt.Println(string(result)) 
}
output==>  //得到base64处理过的切片
Y3SRj2WKAL7SaTG2BAo13voGdi2FAq7osAbkAWWGCAZFDWZFVv1pWu2oWu2oVq1pXjGqVvGpX37sa32lYFG8AUHlBQ6cYTGxDAZ13u1Fa318BAL....
</pre>
//发送post请求(1)/发送的数据是[]byte形式
<pre>
package main

import "fmt"
import "net/http"
import "io/ioutil"
import "errors"
import "bytes"

func SendPost(url string, body []byte) ([]byte, error) {
	requestBody := body //此处没有加密
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	if resp.Status == "404 Not Found" {
		return nil, errors.New("服务器异常!")
	}
	if bodyByte, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Println(string(bodyByte), "--------------")
		responseBody := bodyByte
		return responseBody, err //此处没有加密
	} else {
		return nil, err
	}
}

func main() {
	res, _ := SendPost("http://baidu.com", []byte{'3', '4'})
	fmt.Println(res)

}
output==>
200 OK
<html>
<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
</html>
--------------
[60 104 116 109 108 62 10 60 109 101 116 97 32 104 116 116 112 45 101 113 117 105 118 61 34 114 101 102 114 101 115 104 34 32 99 111 110 116 101 110 116 61 34 48 59 117 114 108 61 104 116 116 112 58 47 47 119 119 119 46 98 97 105 100 117 46 99 111 109 47 34 62 10 60 47 104 116 109 108 62 10]
</pre>
发送put请求
<pre>
package main

import "fmt"
import "net/http"
import "io/ioutil"
import "bytes"

//发送put请求
func SendPut(url string, body []byte) ([]byte, error) {
	client := &http.Client{}
	requestBody := body
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if bodyByte, err := ioutil.ReadAll(resp.Body); err == nil {
		responseBody := bodyByte
		return responseBody, err
	} else {
		return nil, err
	}
}

func main() {
	//待测试
}
</pre>
发送post请求(2)/发送的数据是map形式，也就是常见的表单提交形式，Content-Type：application/x-www-form-urlencoded
<pre>
package main

import (
	"fmt"
	"net/http"
)

func main() {
	//这里添加post的body内容
	var data map[string][]string
	data = make(map[string][]string)
	data["json"] = []string{"jsonStr"}

	//把post表单发送给目标服务器
	res, err := http.PostForm("https://api.huageya.com:8443/server/sign", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("post send success")
	fmt.Println(res)
}
output==>
post send success
&{200 OK 200 HTTP/1.1 1 1 map[Server:[Apache-Coyote/1.1] Set-Cookie:[JSESSIONID=F64A6408B481272B95172638CA8184A5; Path=/server/; Secure; HttpOnly] Content-Type:[text/html;charset=utf-8] Content-Length:[51] Date:[Sun, 12 Jun 2016 08:00:21 GMT]] 0xc0820af080 51 [] false map[] 0xc08209e000 0xc0820c24d0}
</pre>
发送post请求(3)//使用tcp形式发送
<pre>
package main

import (
	"fmt"
	"net"
)

func main() {
	//因为post方法属于HTTP协议，HTTP协议以TCP为基础，所以先建立一个
	//TCP连接，通过这个TCP连接来发送我们的post请求
	conn, err := net.Dial("tcp", "api.huageya.com:8443")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	//构造post请求
	var post string
	post += "POST /postpage HTTP/1.1\r\n"
	post += "Content-Type: application/x-www-form-urlencoded\r\n"
	post += "Content-Length: 37\r\n"
	post += "Connection: keep-alive\r\n"
	post += "Accept-Language: zh-CN,zh;q=0.8,en;q=0.6\r\n"
	post += "\r\n"
	post += "key=json&value=jsonStr\r\n"
	res, err := conn.Write([]byte(post))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("post send success")
	fmt.Println(res)
}
output==>
post send success
186
</pre>
并行计算
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

//并行计算，两个goroutine进行并行的累加计算，都完成后打印
func main() {
	values := []int{12, 3, 4, 5, 6, 7, 7, 8}
	resultChan := make(chan int, 2) //定义2个goroutine
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan //接收结果
	fmt.Println("results:", sum1, sum2)
}
output==>
results: 28 24
</pre>
特定元素过滤
<pre>
package main

import "fmt"
import "strings"

//特定元素过滤,有则返回true
func filter(num string) bool {
	if strings.Contains(num, "123") || strings.Contains(num, "456") {
		return true
	}
	return false
}
func main() {
	s := "thddhhs23gffgbf"
	fmt.Println(filter(s))
}
output==>
false
</pre>
简单时间处理：
<pre>
package main

import "fmt"
import "time"

func GetToday() string {
	today := time.Now().Format("2006-01-02")
	return today
}

//获取当前时间
func GetNowFormart() string {
	now := time.Now().Format("2006-01-02 15:04:05")
	return now
}

//获取今天剩余时间的秒数
func GetTodayLastSecond() time.Duration {
	today := GetToday() + " 23:59:59"
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", today, time.Local)
	return end.Sub(time.Now())
}

//获取活动结束 ，剩余时间的秒数
func GetAtivityLastSecondByEndtime(timestr string) time.Duration {
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
	return end.Sub(time.Now())
}

func main() {
	fmt.Println(GetToday(), GetTodayLastSecond(), GetAtivityLastSecondByEndtime("2016-06-09"))
}
output==>
2016-06-03 3h28m25.0986392s -2562047h47m16.854775808s
</pre>
###Golang json处理实例
<pre>
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigStruct struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}

func main() {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`

	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("==========json str 转map==========")
		fmt.Println(dat)
		fmt.Println(dat["host"])
	}

	//json str 转struct
	var config ConfigStruct
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("==========json str 转struct==========")
		fmt.Println(config)
		fmt.Println(config.Host)
	}

	//json str 转struct(部份字段)
	var part Other
	if err := json.Unmarshal([]byte(jsonStr), &part); err == nil {
		fmt.Println("==========json str 转struct==========")
		fmt.Println(part)
		fmt.Println(part.SerTcpSocketPort)
	}

	//struct 到json str
	if b, err := json.Marshal(config); err == nil {
		fmt.Println("==========struct 到json str==========")
		fmt.Println(string(b))
	}

	//map 到json str
	fmt.Println("==========map 到json str==========")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)

	//array 到 json str
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("==========array 到 json str==========")
		fmt.Println(string(lang))
	}

	//json 到 []string
	var wo []string
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("==========json 到 []string==========")
		fmt.Println(wo)
	}
}
output==>
==========json str 转map==========
map[static_dir:E:/Project/goTest/src/ templates_dir:E:/Project/goTest/src/templates/ serTcpSocketHost::12340 serTcpSocketPort:12340 port:9090 analytics_file: static_file_version:1 host:http://localhost:9090 fruits:[apple peach]]
http://localhost:9090
==========json str 转struct==========
{http://localhost:9090 9090  1 E:/Project/goTest/src/ E:/Project/goTest/src/templates/ :12340 12340 [apple peach]}
http://localhost:9090
==========json str 转struct==========
{:12340 12340 [apple peach]}
12340
==========struct 到json str==========
{"host":"http://localhost:9090","port":9090,"analytics_file":"","static_file_version":1,"static_dir":"E:/Project/goTest/src/","templates_dir":"E:/Project/goTest/src/templates/","serTcpSocketHost":":12340","serTcpSocketPort":12340,"fruits":["apple","peach"]}
==========map 到json str==========
{"analytics_file":"","fruits":["apple","peach"],"host":"http://localhost:9090","port":9090,"serTcpSocketHost":":12340","serTcpSocketPort":12340,"static_dir":"E:/Project/goTest/src/","static_file_version":1,"templates_dir":"E:/Project/goTest/src/templates/"}
==========array 到 json str==========
["hello","apple","python","golang","base","peach","pear"]
==========json 到 []string==========
[hello apple python golang base peach pear]
</pre>
###Golang生成RSAF公私钥
<pre>
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

var outFilePath = flag.String("outpath", "./", "Generate rsa file save path")

func main() {
	flag.Parse()
	if err := genRsaKey(1, *outFilePath); err != nil {
		fmt.Println("密钥文件生成失败！")
	} else {
		fmt.Println("密钥文件生成成功！")
	}
}

func genRsaKey(bits int, filePath string) error {
	//检测生成证书
	if bits > 1024 {
		bits = 2048
	} else {
		bits = 1024
	}

	//查看目录是否存在
	_, err := os.Stat(filePath)
	if err != nil {
		os.Mkdir(filePath, 0777)
	}

	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create(fmt.Sprintf("%s/private.pem", filePath))
	if err != nil {
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create(fmt.Sprintf("%s/public.pem", filePath))
	if err != nil {
		return err
	}
	//fmt.Println("私钥：", base64.StdEncoding.EncodeToString(block.Bytes))
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
output==> //本地目录生成两个文件，分别是public.pem与private.pem
密钥文件生成成功！
</pre>
针对上面的RSA密钥的生成，下面是它的一个应用(RSA加解密)：
<pre>
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDXQrrwRLUhFsgWw4snjemCYryE9+RigE9v7q2/smQI7/4H9TZn
rce/jm80VMxqB6c4t60HbmvuO8WspAp2oN/tZ7jx9qSfCWJoRqI7WbgLovOvwW+h
Yo6orIFsc9lFG8bVRxUUCozijTmJiorKzPis/DUv/yPObav+d1bwKzsJ9QIDAQAB
AoGAP2sdgCP96R25HVvG54Rbw1oriFEwLAT5YlTDQ7Le3fM2uEl6GdmM+9aO1LAW
+TYAAim7BHF3wtxBRLefjYuf7NsLuLmujvXQoxdWm+lXuNQxxiiaTxtuj/3NhymR
NROww9JsAATHSFi6CtLbni3Y+sKYA+NQextcpgSrN4Urv5ECQQDdC8qHZ1q2Rj9F
F+/2g+50L8f0/HDnkMwkRZgj4YnfOn219qXv1tO58W7+da7eQJ+SQVuISi6mFGOZ
4FVQ28y/AkEA+Uy9SvkOK8Nrq926fxrXLXU5LQVAa7+bIpk2EWGpCJo+YtJYdKr1
a+J+nMGQsjeX8m2DK6Q9IvmK1L/Ka5dySwJATmNlEjmT0Ln+q/j+LxTAVlGvfnCb
dXNDAcXwWyEbbJ9of0QVuoUbloBJFVIUjlqqfApTdHSiMGFgpOwKNV+NLwJBAJVQ
FP/Wc1pazR4+yvhdxwr+7qO8RX1DYVMzmGKIr4jrePoPKdOWoS9glJymgld7XJJi
bPGyiLtt4mzSAha2ukkCQGV1X9i+aYN1YP45jGxiizVuRphIrfdAnbiKbYMEXm51
NCtHHhLdIDnVu6tuP981eg87nEm8QQTvwjevSbWUmYY=
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDXQrrwRLUhFsgWw4snjemCYryE
9+RigE9v7q2/smQI7/4H9TZnrce/jm80VMxqB6c4t60HbmvuO8WspAp2oN/tZ7jx
9qSfCWJoRqI7WbgLovOvwW+hYo6orIFsc9lFG8bVRxUUCozijTmJiorKzPis/DUv
/yPObav+d1bwKzsJ9QIDAQAB
-----END PUBLIC KEY-----
`)

func main() {
	var base64Str string

	// rsa加密
	// rsa加密成功后base64编码
	data, err := RsaEncrypt([]byte("Jason RSA"))
	if err == nil {
		base64Str = base64.StdEncoding.EncodeToString(data)
		fmt.Println("加密编码内容=>", base64Str)
	} else {
		fmt.Println("rsa加密结果", err)
	}

	// base64解码
	// 确码成功后rsa解密
	dataStr, bErr := base64.StdEncoding.DecodeString(base64Str)
	if bErr == nil {
		origData, rErr := RsaDecrypt(dataStr)
		if rErr == nil {
			fmt.Println("解码解密内容=>", string(origData))
		} else {
			fmt.Println("rsa解密结果：", rErr)
		}

	} else {
		fmt.Println("base64解码结果：", bErr)
	}

}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
output==>
加密编码内容=> yeOHQY5nXXl3wiNyjWBQO3s0zsT0mZezhgh+ycmve5+uV0/odAsw0bBe/4Innbf6DYxZzPsf8nHUow5MAZKLATjCyfWUGpGndbjfRzNWZ35LvMpZZKy/+B9SD3zisZwGb3JpLYKQt7R8oBHdeyKVGEc97UYNHew/0kmaMzUa4JE=
解码解密内容=> Jason RSA
</pre>
###Golang写日志保存到本地
<pre>
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	logFileName = flag.String("log", "ServerLog.log", "Log file name")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	//set logfile Stdout
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	} else {
		fmt.Println("Log  was wrote!")
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//write log
	log.Printf("Server abort! Cause:%v \n", "test log file\r\n")

}
output==>
Log  was wrote!
</pre>
###Golang struct结构体转buffer缓冲区
<pre>
package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
)

func writeBuf(w io.Writer, v reflect.Value) (n int, err error) {
	newBuf := bytes.NewBuffer(nil)
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.Struct:
			n, err := writeBuf(newBuf, v.Field(i))
			if err != nil {
				return n, err
			}
		case reflect.Bool:
			boolByte := []byte{0}
			if v.Field(i).Bool() {
				boolByte = []byte{1}
			}
			newBuf.Write(boolByte)
		case reflect.String:
			newBuf.WriteString(v.Field(i).String())
		case reflect.Slice:
			newBuf.Write(v.Field(i).Bytes())
		case reflect.Int:
			binary.Write(newBuf, binary.LittleEndian, int32(v.Field(i).Int()))
		case reflect.Uint:
			binary.Write(newBuf, binary.LittleEndian, uint32(v.Field(i).Uint()))
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
			binary.Write(newBuf, binary.LittleEndian, v.Field(i).Interface())
		}
	}
	return w.Write(newBuf.Bytes())
}

func WriteStructToBuffer(w io.Writer, data interface{}) error {
	v := reflect.Indirect(reflect.ValueOf(data))
	if v.Kind() == reflect.Struct {
		fmt.Println("test")
		_, err := writeBuf(w, v)
		return err
	}
	return errors.New("invalid type Not a struct")
}

func StringFixedLength(s string, length int) []byte {
	sLength := len(s)
	if sLength >= length {
		return []byte(s[:length])
	} else {
		b := make([]byte, length-sLength)
		return append([]byte(s), b...)
	}
	return nil
}

func main() {
	s := "1gtrhgrt2345456"
	fmt.Println(string(StringFixedLength(s, 4)))

}
output==>
1gtr
</pre>
###Golang的Echo服务器
<pre>
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var running bool
var (
	host  = "127.0.0.1"
	port  = "8989"
	laddr = host + ":" + port
)

func main() {
	//flag.Parse();
	running = true

	log.Printf("Starting... ")
	conn, err := net.Dial("tcp", laddr)
	if err != nil {
		log.Printf("与服务器%s握手失败;\n 错误：%s", host, err)
		os.Exit(1)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadBytes('\n')
	conn.Write(name[0 : len(name)-1])
	go client_receiver(conn)
	go client_sender(conn)
	for running {
		time.Sleep(1 * 1e9)
	}
}

//客户端发送器
func client_sender(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("ClientToServer>")
		input, _ := reader.ReadBytes('\n')
		if string(input) == "/quit\n" {
			conn.Write([]byte("/quit\n"))
			running = false
			break
		}
		conn.Write(input[0 : len(input)-1])
	}

}

//客户端接收器
func client_receiver(conn net.Conn) {
	//无限循环
	for running {
		fmt.Println("===%s", readabc(conn))
		fmt.Print("You>")
	}
}

//读书连接
func readabc(conn net.Conn) string {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Printf("警告!连接已断开: %s \n", err)
		os.Exit(1)
	}
	return string(buf)
}
output==>
2016/06/04 15:27:03 Starting...
ClientToServer>
ClientToServer>gg
ClientToServer>/quit
===%s HTTP/1.1 400 Bad Request
Content-Type: text/plain
Connection: close

400 Bad Request
You>
</pre>
###beego中post、put与delete请求处理要点（下面以最常见的post表单提交为例）
需要注意xsrf(cross-site request forgery)的安全设置

在发送post请求的cotroller里
<pre>
this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
</pre>
同时在模板中设置
<pre>
<form action="example" method="post">
	{{ .xsrfdata }}
	<input type="text" name="name" /> 
	<input type="text" name="..." />
</form>
</pre>
xsrf支持controller 级别的屏蔽

XSRF 之前是全局设置的一个参数,如果设置了那么所有的API请求都会进行验证,但是有些时候API 逻辑是不需要进行验证的,因此现在支持在controller 级别设置屏蔽:
<pre>
type AdminController struct{
    beego.Controller
}

func (a *AdminController) Prepare() {
    a.EnableXSRF = false
}
</pre>
###先说结论：Golang中所以类型都是值类型，slice/map/channel也都是传值的
<pre>

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 4}
	fmt.Println(a)
	mo(a)
	fmt.Println(a)
}

func mo(data []int) {
	data = nil
}
output==>
[1 3 4]
[1 3 4]           //函数mo修改了a,但是结果却是a没有改变，说明slice不是引用类型
</pre>
###Golang发送手机短信
<pre>
package main

//使用【容联 - 云通讯】发送手机短信
import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 云通讯短信请求
type ReqBody struct {
	To         string   `json:"to"`
	AppId      string   `json:"appId"`
	TemplateId string   `json:"templateId"`
	Datas      []string `json:"datas,omitempty"`
}

// 构建云通讯短信请求
func newSmsRequest(mobile, verifyCode string) *http.Request {
	now := time.Now().Format("20060102150405")
	sigParameter := calcSigParameter(now)
	url := fmt.Sprintf("%s/2013-12-26/Accounts/%s/SMS/TemplateSMS?sig=%s", baseUrl, accountSid, sigParameter)
	b, _ := json.Marshal(ReqBody{
		To:         mobile,
		AppId:      appId,
		TemplateId: "1",
		Datas:      []string{verifyCode, "10"},
	})
	request, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Content-Length", fmt.Sprintf("%d", len(b)))
	request.Header.Set("Authorization", calcAuthorization(now))
	return request
}
func calcSigParameter(now string) string {
	h := md5.New()
	io.WriteString(h, accountSid+authToken+now)
	sign := fmt.Sprintf("%x", h.Sum(nil))
	return strings.ToUpper(sign)
}
func calcAuthorization(now string) string {
	return base64.StdEncoding.EncodeToString([]byte(accountSid + ":" + now))
}

// 云通讯响应
type RespBody struct {
	StatusMsg   string      `json:"statusMsg"`
	StatusCode  string      `json:"statusCode"`
	TemplateSMS interface{} `json:"TemplateSMS"`
}

// 解析云通讯响应
func parseSmsResp(resp *http.Response) (RespBody, error) {
	var data RespBody
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func main() {
	// 发送短信
	request := newSmsRequest("13488888888", "123456")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		// 处理错误
	}
	defer resp.Body.Close()
	data, err := parseSmsResp(resp)
	if err != nil || data.StatusCode != "000000" {
		// 处理错误
	}
}
</pre>
###中国电信sms/发送短信接口
<pre>
package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type accessTokenResponse struct {
	Access_token string
	Expires_in   int
	Res_code     string
	Res_message  string
}

func GetAccessToken(state string) (*accessTokenResponse, error) {
	u := url.Values{}
	u.Add("grant_type", "client_credentials")
	u.Add("app_id", APP_ID)
	u.Add("app_secret", APP_SECRET)
	u.Add("state", state)
	u.Add("scope", "")

	client := http.Client{}
	req, err := http.NewRequest("POST", "https://oauth.api.189.cn/emp/oauth2/v3/access_token", strings.NewReader(u.Encode()))
	if err != nil {
		return nil, err
	}
	// 设置请求头，表示为表单提交
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	// 解析到结构体中
	atr := new(accessTokenResponse)
	err = json.Unmarshal(r, atr)
	if err != nil {
		return nil, err
	}
	return atr, nil
}

const (
	//中国电信提供
	APP_ID     string = ""
	APP_SECRET string = ""
)

type SmsResponse struct {
	Identifier string
	Create_at  string
}

type Sms struct {
	AccessToken string
	Code        float64
	Message     string
	Token       string
}

func NewSms(accesstoken string) (*Sms, error) {
	s := new(Sms)
	s.AccessToken = accesstoken
	err := s.getToken()
	if err != nil {
		return nil, err
	}
	return s, nil

}

func (s *Sms) CustomSms(phone, randcode, exp_time string) (*SmsResponse, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	querys := map[string]string{}
	querys["app_id"] = APP_ID
	querys["access_token"] = s.AccessToken
	querys["timestamp"] = timestamp
	querys["token"] = s.Token
	querys["phone"] = phone
	querys["randcode"] = randcode
	querys["exp_time"] = exp_time
	u := url.Values{}
	for k, v := range querys {
		u.Set(k, v)
	}
	q := createSignQuery(querys)
	// 生成签名
	u.Set("sign", createSign(q))
	// 发送请求
	response, err := http.PostForm("http://api.189.cn/v2/dm/randcode/sendSms", u)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resStruct := new(SmsResponse)
	err = json.Unmarshal(result, resStruct)
	return resStruct, err
}

type tokenResult struct {
	Res_code int
	Token    string
}

func (s *Sms) getToken() error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	u, err := url.Parse("http://api.189.cn/v2/dm/randcode/token")
	if err != nil {
		return err
	}
	q := u.Query()
	querys := map[string]string{}
	querys["app_id"] = APP_ID
	querys["access_token"] = s.AccessToken
	querys["timestamp"] = timestamp
	for k, v := range querys {
		q.Set(k, v)
	}
	sq := createSignQuery(querys)
	// 签名
	q.Set("sign", createSign(sq))
	u.RawQuery = q.Encode()
	// 发送get请求
	response, err := http.Get(u.String())
	if err != nil {
		return err
	}
	// 读取返回结果
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	// 把结果解析到结构体
	tokenresult := new(tokenResult)
	err = json.Unmarshal(result, tokenresult)
	if err != nil {
		return err
	}
	if tokenresult.Res_code == 0 {
		s.Token = tokenresult.Token
		return nil
	} else {
		return errors.New(strconv.Itoa(tokenresult.Res_code) + "  " + tokenresult.Token)
	}
}

func getMapKeys(themap map[string]string) []string {
	keys := []string{}
	for k := range themap {
		keys = append(keys, k)
	}
	return keys
}

func createSignQuery(params map[string]string) string {
	keys := getMapKeys(params)
	sort.Strings(keys)
	q := ""
	for _, v := range keys {
		q += "&" + v + "=" + params[v]
	}
	return strings.TrimLeft(q, "&")
}

func createSign(querys string) string {
	return base64Encode(sha1Encode(querys))
}

func sha1Encode(str string) []byte {
	h := hmac.New(sha1.New, []byte(APP_SECRET))
	h.Write([]byte(str))
	return h.Sum(nil)
}

func base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

func main() {
	ats, err := GetAccessToken("")
	if err != nil {

	}
	s, err := NewSms(ats.Access_token)
	if err != nil {

	}
	rand.Seed(time.Now().UnixNano())
	randstr := fmt.Sprintf("%06d", rand.Intn(999999))
	// 验证码有效时间5分钟
	s.CustomSms("13888888888", randstr, "5")
}
</pre>
###2016-6-7工作记录

- 在有if的地方如果有变量声明的时候，不要使用 := ，而是应该尽量在if判断的外面用 var 先声明一下；
- 模板函数语法：  {{if and (eq a 3) (gt b 4)}} ，对，golang的语法就是这么浪。

###一套高效的超时通知机制(ticker + channel)
<pre>
package main

//一套高效的超时通知机制
import (
	"sync"
	"time"
)

type TimingWheel struct {
	sync.Mutex

	interval time.Duration

	ticker *time.Ticker
	quit   chan struct{}

	maxTimeout time.Duration

	cs []chan struct{}

	pos int
}

func NewTimingWheel(interval time.Duration, buckets int) *TimingWheel {
	w := new(TimingWheel)

	w.interval = interval

	w.quit = make(chan struct{})
	w.pos = 0

	w.maxTimeout = time.Duration(interval * (time.Duration(buckets)))

	w.cs = make([]chan struct{}, buckets)

	for i := range w.cs {
		w.cs[i] = make(chan struct{})
	}

	w.ticker = time.NewTicker(interval)
	go w.run()

	return w
}

func (w *TimingWheel) Stop() {
	close(w.quit)
}

func (w *TimingWheel) After(timeout time.Duration) <-chan struct{} {
	if timeout >= w.maxTimeout {
		panic("timeout too much, over maxtimeout")
	}

	w.Lock()

	index := (w.pos + int(timeout/w.interval)) % len(w.cs)

	b := w.cs[index]

	w.Unlock()

	return b
}

func (w *TimingWheel) run() {
	for {
		select {
		case <-w.ticker.C:
			w.onTicker()
			println("10")
		case <-w.quit:
			w.ticker.Stop()
			return
		}
	}
}

func (w *TimingWheel) onTicker() {
	w.Lock()

	lastC := w.cs[w.pos]
	w.cs[w.pos] = make(chan struct{})

	w.pos = (w.pos + 1) % len(w.cs)

	w.Unlock()

	close(lastC)
}

func main() {
	//这里我们创建了一个timingwheel，精度是0.1s，最大的超时等待时间为10s
	w := NewTimingWheel(100*time.Millisecond, 10)
	for {
		select {
		//等待0.8s
		case <-w.After(800 * time.Millisecond):
			return
		}
	}
}
output==>
10
10
10
10
10
10
10
10
10
</pre>
###Golang shal加密
<pre>
package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	//sha1
	h := sha1.New()
	io.WriteString(h, "aaaaaa")
	fmt.Printf("%x\n", h.Sum(nil))

	//hmac ,use sha1
	key := []byte("123456")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("aaaaaa"))
	fmt.Printf("%x\n", mac.Sum(nil))
}
output==>
f7a9e24777ec23212c54d7a350bc5bea5477fdbb
049988f47afd5ba680e84472db1dd31a6e6051cb
</pre>
Golang实现的数据存储管理器
<pre>
package main

//golang实现数据存储管理器
import (
	"container/list"
	"log"
	"sync"
	"time"
)

type DataManager struct {
	Lock    sync.Mutex
	DL      *list.List
	Expires int
}
type Data struct {
	Key     string
	Expires int
	Value   interface{}
}

var DM *DataManager

func (DM *DataManager) NewData(key string, value interface{}) *Data {
	return &Data{
		Key:     key,
		Expires: DM.Expires,
		Value:   value,
	}
}
func (DM *DataManager) Listen() {
	time.AfterFunc(time.Second, func() { DM.Listen() })
	DM.Lock.Lock()
	defer DM.Lock.Unlock()
	for e := DM.DL.Front(); e != nil; e = e.Next() {
		if e.Value.(*Data).Expires == 0 {
			DM.DL.Remove(e)
		} else {
			e.Value.(*Data).Expires--
		}
	}
}
func (DM *DataManager) Get(key string) interface{} {
	DM.Lock.Lock()
	defer DM.Lock.Unlock()
	for e := DM.DL.Front(); e != nil; e = e.Next() {
		if key == e.Value.(*Data).Key {
			e.Value.(*Data).Expires = DM.Expires
			DM.DL.MoveToBack(e)
			return e.Value.(*Data).Value
		}
	}
	return nil
}
func (DM *DataManager) Set(key string, value interface{}) {
	DM.Lock.Lock()
	defer DM.Lock.Unlock()
	for e := DM.DL.Front(); e != nil; e = e.Next() {
		if key == e.Value.(*Data).Key {
			DM.DL.Remove(e)
			return
		}
	}
	DM.DL.PushBack(DM.NewData(key, value))
}
func (DM *DataManager) Del(key string) {
	DM.Lock.Lock()
	defer DM.Lock.Unlock()
	for e := DM.DL.Front(); e != nil; e = e.Next() {
		if key == e.Value.(*Data).Key {
			DM.DL.Remove(e)
			return
		}
	}
	return
}
func init() {
	DM = &DataManager{
		DL:      list.New(),
		Expires: 5,
	}
	DM.Listen()
}

func main() {
	DM.Set("name", "jason")
	DM.Set("id", "33")
	log.Println(DM.Get("name").(string))
	log.Println(DM.DL.Len())
	log.Println(DM.DL.Len())
	log.Println(DM.Get("id").(string))
	time.Sleep(time.Second)
}
output==>
2016/06/07 19:26:14 jason
2016/06/07 19:26:14 2
2016/06/07 19:26:14 2
2016/06/07 19:26:14 33
</pre>
###slice不定参数
<pre>
package main

import "fmt"

/*
如果我们传入的是slice...这种形式的参数，
go不会创建新的slice,性能不会受到影响
*/
func t(args ...int) {
	fmt.Println(args)
}

func main() {
	a := []int{3, 7, 8, 2, 3, 7, 8, 2}
	b := a[1:]
	t(a...)
	t(b...)
	fmt.Println("b:", b)
}
output==>
[3 7 8 2 3 7 8 2]
[7 8 2 3 7 8 2]
b: [7 8 2 3 7 8 2]
</pre>
###Golang实时监控文件夹变化
<pre>
package main

import (
	"flag"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/go-fsnotify/fsnotify"
)

var (
	sleeptime int
	path      string
	cmd       string
	args      []string
)

func init() {
	flag.IntVar(&sleeptime, "t", 30, "-t=30")
	flag.StringVar(&path, "p", "./", "-p=filepath or dirpath")
	flag.StringVar(&cmd, "c", "", "-c=command")
	str := flag.String("a", "", `-a="args1 args2"`)
	flag.Parse()
	args = strings.Split(*str, " ")
}

func main() {
	Watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Init monitor error: ", err.Error())
		return
	}
	if err := Watch.Add(path); err != nil {
		log.Println("Add monitor path error: ", path)
		return
	}
	var (
		cron bool = false
		lock      = new(sync.Mutex)
	)
	for {
		select {
		case event := <-Watch.Events:
			log.Printf("Monitor event %s", event.String())
			if !cron {
				cron = true
				go func() {
					T := time.After(time.Second * time.Duration(sleeptime))
					<-T
					if err := call(cmd, args...); err != nil {
						log.Println(err)
					}
					lock.Lock()
					cron = false
					lock.Unlock()
				}()
			}
		case err := <-Watch.Errors:
			log.Println(err)
			return
		}
	}
}

func call(programe string, args ...string) error {
	cmd := exec.Command(programe, args...)
	buf, err := cmd.Output()
	if err != nil {
		return err
	}
	log.Printf("\n%s\n", string(buf))
	return nil
}
output==>
2016/06/08 09:27:55 Monitor event ".\\新建文本文档.txt": CREATE
2016/06/08 09:28:07 Monitor event ".\\新建文本文档.txt": RENAME
2016/06/08 09:28:07 Monitor event ".\\test.txt": CREATE
2016/06/08 09:28:07 Monitor event ".\\test.txt": WRITE
2016/06/08 09:28:23 Monitor event ".\\test.txt": RENAME
2016/06/08 09:28:23 Monitor event ".\\tr.txt": CREATE
2016/06/08 09:28:23 Monitor event ".\\tr.txt": WRITE
2016/06/08 09:28:25 exec: "": executable file not found in %PATH%
</pre>
###Golang监控goroutine是否异常退出
在Golang中，我们可以很轻易产生数以万计的goroutine，不过这也带来了麻烦：在运行中某一个goroutine异常退出，怎么办？

在erlang中，有link原语，2个进程可以链接在一起，一个在异常退出的时候，向另一个进程呼喊崩溃的原因，然后由另一个进程处理这些信号，包括是否重启这个进程。在这方面，erlang的确做得很好，估计以后这个特性会在golang中得到实现。据此，有了下面的实现：
<pre>
package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

type message struct {
	normal bool
	state  map[string]interface{}
}

func worker(mess chan message) {
	defer func() {
		exit_message := message{state: make(map[string]interface{})}
		i := recover()
		if i != nil {
			exit_message.normal = false
		} else {
			exit_message.normal = true
		}
		mess <- exit_message
	}()
	now := time.Now()
	seed := now.UnixNano()
	rand.Seed(seed)
	num := rand.Int63()
	if num%2 != 0 {
		panic("not evening")
	} else {
		runtime.Goexit()
	}
}

func supervisor(mess chan message) {
	for i := 0; i < 40; i++ {
		m := <-mess
		switch m.normal {
		case true:
			log.Println("Goroutine exit normal, nothing serious :)")
		case false:
			log.Println("Goroutine exit abnormal, something went wrong!!异常退出啦")
		}

	}
}

func init() {
	fmt.Println("当前时间戳:", strconv.Itoa(int(time.Now().UnixNano())), "当前时间:", time.Now().Format("2006-01-02 15:04:05 -0700"))
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	mess := make(chan message, 50)
	for i := 0; i < 40; i++ {
		go worker(mess)
	}
	supervisor(mess)
	time.Sleep(time.Second)
	//当前的时间戳字符串
	fmt.Println("当前时间戳:", strconv.Itoa(int(time.Now().UnixNano())), "当前时间:", time.Now().Format("2006-01-02 15:04:05 -0700"))
}

output==>
当前时间戳: 1465350575441214100 当前时间: 2016-06-08 09:49:35 +0800
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit abnormal, something went wrong!!异常退出啦
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
2016/06/08 09:49:35 Goroutine exit normal, nothing serious :)
当前时间戳: 1465350576468272900 当前时间: 2016-06-08 09:49:36 +0800
</pre>
###Golang模拟双色球彩票号码
<pre>
package main

import (
	"fmt"
)

func Combinaion(arr []string, m int) []string {
	if m == 1 {
		return arr
	}
	result := make([]string, 0)
	if len(arr) == m {
		var str string
		for i := 0; i < len(arr); i++ {
			str = str + arr[i]
			if i != len(arr)-1 {
				str = str + ","
			}
		}
		result = append(result, str)
		return result
	}

	firstItem := arr[0]
	tempArr1 := Combinaion(append(arr[1:]), m-1)
	for i := 0; i < len(tempArr1); i++ {
		result = append(result, firstItem+","+tempArr1[i])
	}
	tempArr2 := Combinaion(append(arr[1:]), m)
	result = append(result, tempArr2...)
	return result
}

func main() {
	reds := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "22", "23", "24", "25", "26", "28", "29", "30", "31", "32", "33"}
	result := Combinaion(reds, 6)
	fmt.Println("length:\r", len(result))
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i] + "\r")
	}
}
output==>
//跑一下有惊喜
</pre>
###交叉编译工具gox
交叉编译也就是你可以在linux上编译出可以在windows上运行的程序，在32位系统编译出64位系统运行的程序。进入到程序目录中，直接运行gox。程序会一口气生成17个文件。横跨windows,linux,mac,freebsd,netbsd五大操作系统。以及3种了下的处理器(386、amd64、arm),nice.

https://github.com/mitchellh/gox
###Golang三个常用方法-字符串Md5/获取Guid/字符串截取
<pre>
package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

//md5方法
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//Guid方法
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//字串截取
func SubString(s string, startpos, length int) string {
	runes := []rune(s)
	l := startpos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[startpos:l])
}
func main() {
	println(GetGuid())
	str := "hellojason"
	println(SubString(str, 2, 5))
}
output==>
c074eb9467c702c1a93d2b5ac0bce59f
lloja
</pre>
###Golang实现代理访问
<pre>
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//指定代理ip
func getTransportFieldURL(proxy_addr *string) (transport *http.Transport) {
	url_i := url.URL{}
	url_proxy, _ := url_i.Parse(*proxy_addr)
	transport = &http.Transport{Proxy: http.ProxyURL(url_proxy)}
	return
}

//从环境变量$http_proxy或$HTTP_PROXY中获取HTTP代理地址
func getTransportFromEnvironment() (transport *http.Transport) {
	transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	return
}
func fetch(url, proxy_addr *string) (html string) {
	transport := getTransportFieldURL(proxy_addr)
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		html = string(robots)
	} else {
		html = ""
	}
	return
}
func main() {
	proxy_addr := "http://221.10.251.196:80/"
	url := "http://zituo.net"
	html := fetch(&url, &proxy_addr)
	fmt.Println(html)
}
</pre>
###Golang获取本机MAC唯一标识
<pre>
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	var macAddr string
	var hws, err = net.Interfaces()
	if err != nil {
		fmt.Println("[InitLockerInfo]", "获取服务器MAC地址失败", err.Error())
		return
	}
	for _, h := range hws {
		var addr = h.HardwareAddr.String()
		if addr != "" && !strings.HasPrefix(addr, "00:00:00:00:00:00") {
			macAddr = h.HardwareAddr.String()
			break
		}
	}
	if macAddr != "" {
		fmt.Println("本机MAC硬件地址:", macAddr)
	} else {
		fmt.Println("[InitLockerInfo]", "未找到MAC硬件地址")
	}
}
output==>
</pre>
###Golang显示本机IP
<pre>
package main

import (
	"net"
	"strings"
)

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

func main() {
	println(GetLocalIP())
}
</pre>
###Golang实现tcp转发代理
<pre>
package main  
  
import (  
    "flag"  
    "fmt"  
    "io"  
    "net"  
    "os"  
    "strings"  
    "sync"  
)  
  
var lock sync.Mutex  
var trueList []string  
var ip string  
var list string  
  
func main() {  
    flag.StringVar(&ip, "l", ":9897", "-l=0.0.0.0:9897 指定服务监听的端口")  
    flag.StringVar(&list, "d", "127.0.0.1:1789,127.0.0.1:1788", "-d=127.0.0.1:1789,127.0.0.1:1788 指定后端的IP和端口,多个用','隔开")  
    flag.Parse()  
    trueList = strings.Split(list, ",")  
    if len(trueList) <= 0 {  
        fmt.Println("后端IP和端口不能空,或者无效")  
        os.Exit(1)  
    }  
    server()  
}  
  
func server() {  
    lis, err := net.Listen("tcp", ip)  
    if err != nil {  
        fmt.Println(err)  
        return  
    }  
    defer lis.Close()  
    for {  
        conn, err := lis.Accept()  
        if err != nil {  
            fmt.Println("建立连接错误:%v\n", err)  
            continue  
        }  
        fmt.Println(conn.RemoteAddr(), conn.LocalAddr())  
        go handle(conn)  
    }  
}  
  
func handle(sconn net.Conn) {  
    defer sconn.Close()  
    ip, ok := getIP()  
    if !ok {  
        return  
    }  
    dconn, err := net.Dial("tcp", ip)  
    if err != nil {  
        fmt.Printf("连接%v失败:%v\n", ip, err)  
        return  
    }  
    ExitChan := make(chan bool, 1)  
    go func(sconn net.Conn, dconn net.Conn, Exit chan bool) {  
        _, err := io.Copy(dconn, sconn)  
        fmt.Printf("往%v发送数据失败:%v\n", ip, err)  
        ExitChan <- true  
    }(sconn, dconn, ExitChan)  
    go func(sconn net.Conn, dconn net.Conn, Exit chan bool) {  
        _, err := io.Copy(sconn, dconn)  
        fmt.Printf("从%v接收数据失败:%v\n", ip, err)  
        ExitChan <- true  
    }(sconn, dconn, ExitChan)  
    <-ExitChan  
    dconn.Close()  
}  
  
func getIP() (string, bool) {  
    lock.Lock()  
    defer lock.Unlock()  
  
    if len(trueList) < 1 {  
        return "", false  
    }  
    ip := trueList[0]  
    trueList = append(trueList[1:], ip)  
    return ip, true  
} 
</pre>
###Golang使用pprof监控性能
下面的代码是有问题的！在此仅作为使用pprof的试验器材：
<pre>
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func counter() {
	list := []int{1}
	c := 1
	for i := 0; i < 10000000; i++ {
		httpGet()
		c = i + 1 + 2 + 3 + 4 - 5
		list = append(list, c)
	}
	fmt.Println(c)
	fmt.Println(list[0])
}

func work(wg *sync.WaitGroup) {
	for {
		counter()
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func httpGet() int {
	queue := []string{"start..."}
	resp, err := http.Get("http://www.163.com")
	if err != nil {
		// handle error
	}

	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	queue = append(queue, string(body))
	return len(queue)
}

func main() {
	flag.Parse()

	//这里实现了远程获取pprof数据的接口
	go func() {
		log.Println(http.ListenAndServe("localhost:7777", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 100; i++ {
		go work(&wg)
	}

	wg.Wait()
	time.Sleep(3 * time.Second)
}
//
http://localhost:7777/debug/pprof/
查看打印的内容
</pre>
###Golang给目录按时间排序，使用正则、时间条件搜索文件
<pre>
package filepath  
  
import (  
    "io/ioutil"  
    "os"  
    "path/filepath"  
    "regexp"  
    "strings"  
    "time"  
)  
  
//path表示搜索的路径,FullDir表示是不是递归查询,MatchDir表示是否匹配目录.  
type FindFiles struct {  
    Path     string `json:path`  
    FullDir  bool   `json:fulldir`  
    MatchDir bool   `json:matchdir`  
}  
  
//date小于等于0的时候表示查找最近这段时间的文件  
func (self FindFiles) DateFindFile(date int64) ([]string, error) {  
    date = date * 24 * 60 * 60  
    var less bool  
    switch {  
    case date <= 0:  
        date = time.Now().Unix() + date  
        less = true  
    case date > 0:  
        date = time.Now().Unix() - date  
        less = false  
    }  
    return datewalk(date, less, self.FullDir, self.MatchDir, self.Path)  
}  
  
func (self FindFiles) RegFindFile(reg string) ([]string, error) {  
    if strings.Index(reg, "*") == 0 {  
        reg = "." + reg  
    } else {  
        reg = "^" + reg  
    }  
    reg += "$"  
    Reg, err := regexp.Compile(reg)  
    if err != nil {  
        return []string{}, nil  
    }  
    if self.FullDir {  
        return namewalk(Reg, self.MatchDir, self.Path)  
    }  
    var list []string  
    infos, err := readDir(self.Path)  
    if err != nil {  
        return list, nil  
    }  
    path := filepath.ToSlash(self.Path)  
    if !strings.HasSuffix(path, "/") {  
        path += "/"  
    }  
    for _, v := range infos {  
        if Reg.MatchString(v.Name()) {  
            if v.IsDir() && !self.MatchDir {  
                continue  
            }  
            list = append(list, path+v.Name())  
        }  
    }  
    return list, nil  
}  
  
func (self FindFiles) DateAndRegexp(date int64, reg string) ([]string, error) {  
    var l []string  
    list, err := self.RegFindFile(reg)  
    if err != nil {  
        return l, err  
    }  
    date = date * 24 * 60 * 60  
    var less bool = false  
    if date <= 0 {  
        date = time.Now().Unix() + date  
        less = true  
    } else {  
        date = time.Now().Unix() - date  
    }  
    for _, v := range list {  
        info, err := os.Stat(v)  
        if err != nil {  
            continue  
        }  
        if less {  
            if date > info.ModTime().Unix() {  
                continue  
            }  
        } else {  
            if date < info.ModTime().Unix() {  
                continue  
            }  
        }  
        l = append(l, v)  
    }  
    return l, nil  
}  
  
func datewalk(date int64, less bool, fulldir, matchdir bool, path string) ([]string, error) {  
    var list []string  
    if !strings.HasSuffix(path, "/") {  
        path += "/"  
    }  
    if !fulldir {  
        infos, err := readDir(path)  
        if err != nil {  
            return list, err  
        }  
        for _, info := range infos {  
            file, ok := dResolve(date, less, matchdir, path, info)  
            if ok {  
                file = path + file  
                list = append(list, file)  
            }  
        }  
        return list, nil  
    }  
    return list, filepath.Walk(path, func(root string, info os.FileInfo, err error) error {  
        if err != nil {  
            return err  
        }  
        _, ok := dResolve(date, less, matchdir, root, info)  
        if ok {  
            root = filepath.ToSlash(root)  
            list = append(list, root)  
        }  
        return nil  
    })  
}  
  
func dResolve(date int64, less, matchdir bool, root string, info os.FileInfo) (string, bool) {  
    if less {  
        if date > info.ModTime().Unix() {  
            return "", false  
        }  
    } else {  
        if date < info.ModTime().Unix() {  
            return "", false  
        }  
    }  
    root = filepath.ToSlash(root)  
    if info.IsDir() && !matchdir {  
        return "", false  
    }  
  
    return info.Name(), true  
}  
  
func namewalk(reg *regexp.Regexp, matchdir bool, path string) ([]string, error) {  
    var list []string  
    return list, filepath.Walk(path, func(root string, info os.FileInfo, err error) error {  
        if err != nil {  
            return err  
        }  
        if !reg.MatchString(info.Name()) {  
            return nil  
        }  
        root = filepath.ToSlash(root)  
        if info.IsDir() && !matchdir {  
            return nil  
        }  
        list = append(list, root)  
        return nil  
    })  
}  
  
func readDir(path string) ([]os.FileInfo, error) {  
    info, err := os.Stat(path)  
    if err != nil {  
        return nil, err  
    }  
    if info.IsDir() {  
        return ioutil.ReadDir(path)  
    }  
    return []os.FileInfo{info}, nil  
} 
</pre>
###一种web路由的实现
Router
<pre>
package main

import (
	"net/http"
	"sync"
)

type Router map[string]func(w http.ResponseWriter, r *http.Request)

var routerMap Router = make(Router)
var lock *sync.RWMutex = new(sync.RWMutex)

func main() {
	routerMap.Regist("/", ce)
	hand := routerMap.Handler("/")
	http.HandleFunc("/", hand)
	http.ListenAndServe(":1789", nil)
}

func ce(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func (self Router) Regist(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	lock.Lock()
	defer lock.Unlock()
	self[pattern] = f
}

func (self Router) Handler(pattern string) func(w http.ResponseWriter, r *http.Request) {
	lock.RLock()
	defer lock.RUnlock()
	return self[pattern]
}
</pre>
###Golang为错误分级
<pre>
package main

import (
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

type l struct {
	logs  *log.Logger
	level int
	io.Closer
}

func NewLog(HttpLogPath string, level int) *l {
	file, err := os.OpenFile(HttpLogPath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Println("Error", err)
		log.Println("Error", "日志输出到标准输出.")
		return nil
	}
	var Log *log.Logger = log.New(os.Stdout, now(), 0)
	if file != nil {
		Log = log.New(file, now(), 0)
		go flushLogFile(file)
	}
	file.Seek(0, 2)
	return &l{Log, level, file}
}

func now() string {
	return time.Now().Format("2006-01-02 15:04:05 ")
}

func flushLogFile(File *os.File) {
	for _ = range time.NewTicker(50 * time.Second).C {
		if File == nil {
			return
		}
		File.Sync()
	}
}

func (self *l) SetLogLevel(level int) {
	if level > 4 {
		return
	}
	self.level = level
}

func (self *l) Print(v ...interface{}) {
	self.logs.Print(v)
}

func (self *l) Printf(formate string, v ...interface{}) {
	self.logs.Printf(formate, v)
}

func (self *l) Println(v ...interface{}) {
	self.logs.Println(v)
}

func (self *l) PrintfI(formate string, v ...interface{}) {
	if self.level > 1 {
		return
	}
	self.logs.Printf("Info->"+formate, v...)
}

func (self *l) PrintfW(formate string, v ...interface{}) {
	if self.level > 2 {
		return
	}
	self.logs.Printf("Warn->"+formate, v...)
}

func (self *l) PrintfE(formate string, v ...interface{}) {
	if self.level > 3 {
		return
	}
	self.logs.Printf("Error->"+formate, v...)
}

func (self *l) PrintfF(formate string, v ...interface{}) {
	if self.level > 4 {
		return
	}
	self.logs.Fatalf("Fatal->"+formate, v...)
}

func (self *l) InfoPrintf(callers int, formate string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(callers + 1)
	if !ok {
		return
	}
	self.logs.Printf("File->%s Line->%d\n", file, line)
	self.logs.Printf(formate, v...)
}

func main() {

	ltest := NewLog("http.log", 0)

	ltest.InfoPrintf(0, "\r\n%s\r\n", "错误在这里↑")

	ltest.PrintfI("%s\r\n", "Info")

	ltest.PrintfW("%s\r\n", "Warn")

	ltest.PrintfE("%s\r\n", "Error")

	ltest.PrintfF("%s\r\n", "Fatal")

}
</pre>
###Golang将ip地址作为Int形式存储
<pre>
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type IntIP struct {
	IP    string
	Intip int
}

func main() {
	var x *IntIP = &IntIP{IP: "192.168.1.1"}
	fmt.Println(x)
	x.ToIntIp()
	fmt.Println(*x)
}

func (self *IntIP) String() string {
	return self.IP
}

func (self *IntIP) ToIntIp() (int, error) {
	Intip, err := ConvertToIntIP(self.IP)
	if err != nil {
		return 0, err
	}
	self.Intip = Intip
	return Intip, nil
}

func (self *IntIP) ToString() (string, error) {
	i4 := self.Intip & 255
	i3 := self.Intip >> 8 & 255
	i2 := self.Intip >> 16 & 255
	i1 := self.Intip >> 24 & 255
	if i1 > 255 || i2 > 255 || i3 > 255 || i4 > 255 {
		return "", errors.New("Isn't a IntIP Type.")
	}
	ipstring := fmt.Sprintf("%d.%d.%d.%d", i4, i3, i2, i1)
	self.IP = ipstring
	return ipstring, nil
}
func ConvertToIntIP(ip string) (int, error) {
	ips := strings.Split(ip, ".")
	E := errors.New("Not A IP.")
	if len(ips) != 4 {
		return 0, E
	}
	var intIP int
	for k, v := range ips {
		i, err := strconv.Atoi(v)
		if err != nil || i > 255 {
			return 0, E
		}
		intIP = intIP | i<<uint(8*(3-k))
	}
	return intIP, nil
}
output==>
192.168.1.1
{192.168.1.1 3232235777}
</pre>
###使用Golang实现游戏批量搭服的小程序
<pre>
package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

type Template struct {
	Name       string
	Md5        string
	ConfigList []string
}

var url string
var homedir string
var Temp_Path map[string]string = make(map[string]string)
var HeadList []string
var Replational []string

func main() {
	flag.StringVar(&url, "u", "http://127.0.0.1:1789", "Specify server address.")
	flag.StringVar(&homedir, "d", "/data/gamehome", "Specify home directory.")
	L := flag.Bool("l", false, "local generation config.")
	cfgTempRelational := flag.String("c", "CfgTempRelational.ini", "Specify CfgTempRelational config")
	relationalTable := flag.String("r", "RelationalTable.ini", "Specify RelationalTable config.")
	name := flag.String("n", "server.zip", "Specify server package name.")
	flag.Parse()
	if !strings.HasSuffix(homedir, "/") {
		homedir = homedir + "/"
	}
	if *L {
		if Unzip(*name, homedir) {
			LocalReplace(*cfgTempRelational, *relationalTable)
			return
		}
		return
	}
	fmt.Printf("%s Start request template info.From %s\n", GetNow(), url)
	req, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s Request info error\n", GetNow())
		return
	}
	var TemplateInfo Template
	func(t *Template) {
		defer req.Body.Close()
		err := gob.NewDecoder(req.Body).Decode(t)
		if err != nil {
			fmt.Printf("%s Parse template info:\n%s\n", GetNow(), err)
			os.Exit(1)
		}
	}(&TemplateInfo)
	fmt.Printf("%s Parse Template info is OK.\n", GetNow())
	fmt.Printf("%s Print Template Info:\n%s\n", GetNow(), TemplateInfo)
	if Download(TemplateInfo) {
		if !Unzip(TemplateInfo.Name, homedir) {
			fmt.Printf("%s Unzip file error.\n", GetNow())
			os.Exit(2)
		}
	}
	ConfigDownload(TemplateInfo.ConfigList)
}

func LocalReplace(CfgTempRelational, RelationalTable string) {
	fmt.Printf("start parse %s\n", CfgTempRelational)
	GetPathConfig(CfgTempRelational)
	fmt.Printf("start parse %s\n", RelationalTable)
	ParseServerConfig(RelationalTable)
	localip := GetLocalIP()
	fmt.Printf("get local ip %s\n", localip)
	for _, v := range localip {
		info := Matching(v)
		if len(info) > 0 {
			valueList := Split(info)
			fmt.Printf("Match values %s\n", valueList)
			Key_Map := Merge(valueList)
			for k, v := range Temp_Path {
				fmt.Printf("create config %s%s\n", homedir, v)
				File, err := os.Create(fmt.Sprintf("%s%s", homedir, v))
				if err != nil {
					fmt.Printf("%s create %s faild\n%s\n", GetNow(), v, err)
					os.Exit(1)
				}
				ExecuteReplace(File, k, Key_Map)
			}
			return
		}
	}
}

func ConfigDownload(TemplateList []string) {
	for _, v := range TemplateList {
		resp, err := http.Get(fmt.Sprintf("%s/config?key=%s", url, v))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		path := homedir + resp.Header.Get("path")
		fmt.Printf("%s Start create %s file\n", GetNow(), path)
		File, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		io.Copy(File, resp.Body)
	}
	fmt.Printf("%s Create config is Ok.\n", GetNow())
}

func Download(TemplateInfo Template) bool {
	_, err := os.Lstat(TemplateInfo.Name)
	if err == nil {
		if Md5(TemplateInfo.Name) == TemplateInfo.Md5 {
			fmt.Printf("%s local is exist %s\n", GetNow(), TemplateInfo.Name)
			return true
		}
	}
	fmt.Printf("%s Start download %s\n", GetNow(), TemplateInfo.Name)
	u := fmt.Sprintf("%s/template/%s", url, TemplateInfo.Name)
	req, err := http.Get(u)
	if err != nil {
		fmt.Printf("%s Open %s error:\n%s\n", GetNow(), u, err)
		return false
	}
	defer req.Body.Close()
	File, err := os.Create(TemplateInfo.Name)
	if err != nil {
		fmt.Printf("%s Create %s error:\n%s\n", GetNow(), TemplateInfo.Name, err)
		return false
	}
	io.Copy(File, req.Body)
	if Md5(TemplateInfo.Name) != TemplateInfo.Md5 {
		fmt.Printf("%s Check md5 faild!\n", GetNow())
		return false
	}
	return true
}

func Md5(path string) string {
	fmt.Printf("%s Check md5 %s\n", GetNow(), path)
	File, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s Check md5 error:\n%s\n", GetNow(), err)
		return ""
	}
	m := md5.New()
	io.Copy(m, File)
	return fmt.Sprintf("%X", string(m.Sum([]byte{})))
}

func Unzip(filename, dir string) bool {
	fmt.Printf("%s Unzip to %s\n", GetNow(), dir)
	File, err := zip.OpenReader(filename)
	if err != nil {
		fmt.Printf("%s Open zip faild:\n%s\n", GetNow(), err)
		return false
	}
	defer File.Close()
	for _, v := range File.File {
		v.Name = fmt.Sprintf("%s%s", dir, v.Name)
		info := v.FileInfo()
		if info.IsDir() {
			err := os.MkdirAll(v.Name, 0644)
			if err != nil {
				fmt.Printf("%s Create direcotry %s faild:\n%s\n", GetNow(), v.Name, err)
				return false
			}
			continue
		}
		srcFile, err := v.Open()
		if err != nil {
			fmt.Printf("%s Read from zip faild:\n%s\n", GetNow(), err)
			return false
		}
		defer srcFile.Close()
		newFile, err := os.Create(v.Name)
		if err != nil {
			fmt.Printf("%s Create file faild:\n%s\n", GetNow(), err)
			return false
		}
		io.Copy(newFile, srcFile)
		newFile.Close()
	}
	return true
}

func ExecuteReplace(w *os.File, temp_path string, funcs map[string]string) error {
	fmt.Println(funcs)
	T := template.New("")
	buf, err := ioutil.ReadFile(temp_path)
	if err != nil {
		return err
	}
	T, err = T.Parse(string(buf))
	if err != nil {
		return err
	}
	err = T.Execute(w, funcs)
	if err != nil {
		return err
	}
	return nil
}

func GetPathConfig(path string) {
	File, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	defer File.Close()
	M := make(map[string]string)
	Buf := bufio.NewReader(File)
	var linenum int = 1
	for {
		line, _, err := Buf.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			os.Exit(5)
		}
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		list := bytes.Split(line, []byte("="))
		if len(list) != 2 {
			fmt.Printf("check config %s ,line %d\n", path, linenum)
			os.Exit(6)
		}
		key := string(bytes.TrimSpace(list[0]))
		value := string(bytes.TrimSpace(list[1]))
		TestPath(key)
		M[key] = value
		Check(key)
		linenum++
	}
	if len(M) < 1 {
		fmt.Printf("config %s can't emptey!\n", path)
		os.Exit(7)
	}
	Temp_Path = M
}
func TestPath(path string) {
	info, err := os.Lstat(path)
	if err != nil {
		fmt.Printf("check %s .error_info :%s", path, err)
		os.Exit(9)
	}
	if info.IsDir() {
		fmt.Printf("check %s is directory", path)
		os.Exit(10)
	}
	return
}
func Merge(list []string) map[string]string {
	ExecuteReplaceValue := make(map[string]string)
	for k, v := range HeadList {
		ExecuteReplaceValue[v] = list[k]
	}
	return ExecuteReplaceValue
}

func ParseServerConfig(path string) {
	File, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer File.Close()
	buf := bufio.NewReader(File)
	var linenum int = 1
	for i := 0; i < 1001; i++ { //just init top 1000
		line, _, err := buf.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		if len(HeadList) == 0 {
			list := Split(string(line))
			if len(list) > 0 {
				HeadList = list
			}
			continue
		}
		list := Split(string(line))
		if len(list) != len(HeadList) {
			fmt.Printf("Line %d parse error.", linenum)
			continue
		}
		Replational = append(Replational, string(line))
	}
	if len(Replational) <= 0 {
		fmt.Println("read config error.")
		os.Exit(1)
	}
}

func Split(str string) []string {
	var l []string
	list := strings.Split(str, " ")
	for _, v := range list {
		if len(v) == 0 {
			continue
		}
		if strings.Contains(v, "    ") {
			list := strings.Split(v, "  ")
			for _, v := range list {
				if len(v) == 0 {
					continue
				}
				l = append(l, v)
			}
			continue
		}
		l = append(l, v)
	}
	return l
}
func GetLocalIP() []string {
	list, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	var l []string
	for _, v := range list {
		if strings.Contains(v.String(), ":") {
			continue
		}
		ip := strings.Split(v.String(), "/")
		if len(ip) != 2 {
			continue
		}
		l = append(l, ip[0])
	}
	return l
}
func Matching(srcip string) string {
	for _, v := range Replational {
		if strings.Contains(v, srcip) {
			return v
		}
	}
	return ""
}

func GetNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Check(path string) {
	File, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer File.Close()
	buf := bufio.NewReader(File)
	var num int = 1
	var errornum int = 0
	s := []byte("{{")
	e := []byte("}}")
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return
		}
		if bytes.Count(line, s) != bytes.Count(line, e) {
			fmt.Printf("Line%d: %s\n", num, string(line))
			errornum++
			continue
		}
		if bytes.Count(line, []byte("{{.}}")) != 0 {
			fmt.Printf("Line%d: %s\n", num, string(line))
			errornum++
			continue
		}

		for i := 0; i < bytes.Count(line, s); i++ {
			first := bytes.Index(line, s)
			last := bytes.Index(line, e)
			if first == -1 || last == -1 {
				continue
			}
			if bytes.Index(line[first:last], []byte("{{.")) != 0 {
				fmt.Printf("Error Line %d: %s\n", num, string(line))
				errornum++
				break
			}
			line = line[last:]
		}
	}
	if errornum != 0 {
		fmt.Printf("Error num %d From %s\n", errornum, path)
		return
	}
	return
}
</pre> 
###Golang实现异步生成log程序
<pre>
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Logger struct {
	console bool
	warn    bool
	info    bool
	tformat func() string
	file    chan string
}

func NewLog(level string, console bool, File *os.File, buf int) (*Logger, error) {
	log := &Logger{console: console, tformat: format}
	if File != nil {
		FileInfo, err := File.Stat()
		if err != nil {
			return nil, err
		}
		mode := strings.Split(FileInfo.Mode().String(), "-")
		if strings.Contains(mode[1], "w") {
			str_chan := make(chan string, buf)
			log.file = str_chan
			go func() {
				for {
					fmt.Fprintln(File, <-str_chan)
				}
			}()
			defer func() {
				for len(str_chan) > 0 {
					time.Sleep(1e9)
				}
			}()
		} else {
			return nil, errors.New("can't write.")
		}
	}
	switch level {
	case "Warn":
		log.warn = true
		return log, nil
	case "Info":
		log.warn = true
		log.info = true
		return log, nil
	}
	return nil, errors.New("level must be Warn or Info.")
}

func (self *Logger) Error(info interface{}) {
	if self.console {
		fmt.Println("Error", self.tformat(), info)
	}
	if self.file != nil {
		self.file <- fmt.Sprintf("Error %s %s", self.tformat(), info)

	}
}

func (self *Logger) Warn(info interface{}) {
	if self.warn && self.console {
		fmt.Println("Warn", self.tformat(), info)
	}
	if self.file != nil {
		self.file <- fmt.Sprintf("Warn %s %s", self.tformat(), info)
	}
}
func (self *Logger) Info(info interface{}) {
	if self.info && self.console {
		fmt.Println("Info", self.tformat(), info)
	}
	if self.file != nil {
		self.file <- fmt.Sprintf("Info %s %s", self.tformat(), info)
	}
}
func (self *Logger) Close() {
	for len(self.file) > 0 {
		time.Sleep(1e8)
	}
}
func format() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func main() {
	File, _ := os.Create("log")
	log, err := NewLog("Info", true, File, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer log.Close()
	for i := 0; i < 1000; i++ {
		log.Warn("Jason")
		log.Info("warning")
	}

}
</pre>
###Golang实现简单文件服务器
<pre>
package main

import (
	"fmt"

	"net/http"
)

func main() {

	h := http.FileServer(http.Dir("./"))

	http.ListenAndServe(":1789", ce(h))

}

func ce(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.URL.Path)

		h.ServeHTTP(w, r)

	})

}
</pre>
###Golang使用正则路由实现http服务器
<pre>
package main

import (
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", route)
	http.ListenAndServe(":8080", nil)
}

var num = regexp.MustCompile(`\d`)
var str = regexp.MustCompile(`\w`)

func route(w http.ResponseWriter, r *http.Request) {
	switch {
	case num.MatchString(r.URL.Path):
		digits(w, r)
	case str.MatchString(r.URL.Path):
		sstr(w, r)
	default:
		w.Write([]byte("位置匹配项"))
	}
}

func digits(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("receive digits"))
}

func sstr(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("receive string"))
}
output==>
//打开浏览器，分别访问
localhost:8080/
localhost:8080/8
localhost:8080/ff
</pre>
###Golang实现命令行输入关键字在文件夹/文件中搜索
<pre>
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var path *string = flag.String("p", "./", "搜索的路径")
var re_string *string = flag.String("r", "", "关键字")

func main() {
	flag.Parse()
	if *re_string == "" {
		fmt.Println("搜索的关键字不能为空")
		return
	}
	fmt.Println("搜索的路径：", *path, "搜索的关键字：", *re_string)
	re, _ := regexp.Compile(*re_string)
	filepath.Walk(*path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Println("文件名称/路径：", path)
		File, _ := os.Open(path)
		r := bufio.NewReader(File)
		for {
			b, _, e := r.ReadLine()
			if e != nil {
				break
			}
			if b, _ := regexp.Match(string([]byte{0}), b); b {
				break
			}
			if re.Match(b) {
				fmt.Println(string(b))
			}
		}
		File.Close()
		return nil
	})
}
output==>
$ go run test.go -r fmt
搜索的路径： ./ 搜索的关键字： fmt
文件名称/路径： test.exe
文件名称/路径： test.go
        "fmt"
                fmt.Println("搜索的关键字不能为空")
        fmt.Println("搜索的路径：", *path, "搜索的关键字：", *re_string)
                fmt.Println("文件名称/路径：", path)
                                fmt.Println(string(b))


</pre>
###Golanngs实现类似Proxy的小程序{可以用来访问goolge} 
<pre>
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", route)
	e := http.ListenAndServe(":80", nil)
	if e != nil {
		fmt.Println(e)
	}
}

func route(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(r.Method, "", r.Body)
	req.URL = r.URL
	req.URL.Host = "www.qq.com" //"www.google.com"
	req.URL.Scheme = "http"
	for _, v := range r.Cookies() {
		req.AddCookie(v)
	}
	//req.Header = r.Header 这里的Header就不要使用了,使用的话他会自动跳转到https,代理就出问题了.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Here:", err)
		return
	}
	for k, v := range resp.Header {
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}
	for _, cookie := range resp.Cookies() {
		w.Header().Add("Set-Cookie", cookie.Raw)
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
	r.Body.Close()
}
</pre>
###Golang为结构体/struct排序
根据修改时间为某一目录下的所有文件排序
<pre>
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type info struct {
	Name string
	Time time.Time
}
type newlist []*info

func main() {
	l, e := getFilelist("./")
	if e != nil {
		fmt.Println(e)
	}
	sort.Sort(newlist(l)) //调用标准库的sort.Sort必须要先实现Len(),Less(),Swap() 三个方法.
	for _, v := range l {
		fmt.Println("文件名：", v.Name, "修改时间：", v.Time.Unix())
	}
}

func getFilelist(path string) ([]*info, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return []*info{}, err
	}
	var list []*info
	for _, v := range l {
		list = append(list, &info{v.Name(), v.ModTime()})
	}
	return list, nil
}

func (I newlist) Len() int {
	return len(I)
}
func (I newlist) Less(i, j int) bool {
	return I[i].Time.Unix() < I[j].Time.Unix()
}
func (I newlist) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}
output==>
文件名： test.exe 修改时间： 1464094340
文件名： test.go 修改时间： 1465400102
文件名： te.go 修改时间： 1465400411
文件名： tes.go 修改时间： 1465400422
</pre>
###工作小计
nohup ./httpserver &

nohup这个命令可以把程序放后台运行，顺便通过1>和2>把标准输出和标准错误重定向到文件，这样程序崩溃时才会有记录可查，这两者和程序的日志最好是分开，混在一起没办法判断轻重缓急:

nohup ./server 1> server.out 2> server.err
###Golang文件共享
<pre>
package main  
  
import (  
    "flag"  
    "fmt"  
    "io/ioutil"  
    "net/http"  
    "path/filepath"  
    "sort"  
    "sync"  
    "text/template"  
    "time"  
)  
  
const L = `<html>  
<title>文件列表</title>  
<body>  
    {{$ip := .IP}}  
    {{$dir := .Dir}}  
    <table>  
    {{range $k,$v := .List}}<tr><td><a href="http://{{$ip}}/{{$dir}}/{{$v.Name}}">文件名：{{$v.Name}}</a></td><td>  修改时间：{{$v.Time}}</td></tr>  
{{end}}  
    </table>  
</body>  
</html>`  
  
type info struct {  
    Name string  
    Time time.Time  
}  
  
type newlist []*info  
  
type Dirinfo struct {  
    lock sync.Mutex  
    IP   string  
    Dir  string  
    List newlist  
}  
  
var x Dirinfo  
var name, dir string  
var path *string = flag.String("p", "/tmp", "共享的路径")  
var port *string = flag.String("l", ":1789", "监听的IP:端口")  
  
func main() {  
    flag.Parse()  
    name = filepath.Base(*path)  
    dir = filepath.Dir(*path)  
    fmt.Println("共享的目录：", *path)  
    http.Handle(fmt.Sprintf("/%s/", name), http.FileServer(http.Dir(dir)))  
    http.HandleFunc("/", router)  
    http.ListenAndServe(*port, nil)  
}  
  
func router(w http.ResponseWriter, r *http.Request) {  
    l, _ := getFilelist(*path)  
    x.lock.Lock()  
    x.Dir = name  
    x.List = l  
    x.IP = r.Host  
    x.lock.Unlock()  
    t := template.New("")  
    t.Parse(L)  
    t.Execute(w, x)  
}  
  
func getFilelist(path string) (newlist, error) {  
    l, err := ioutil.ReadDir(path)  
    if err != nil {  
        return []*info{}, err  
    }  
    var list []*info  
    for _, v := range l {  
        list = append(list, &info{v.Name(), v.ModTime()})  
    }  
    sort.Sort(newlist(list))  
    return list, nil  
}  
  
func (I newlist) Len() int {  
    return len(I)  
}  
func (I newlist) Less(i, j int) bool {  
    return I[i].Time.Unix() < I[j].Time.Unix()  
}  
func (I newlist) Swap(i, j int) {  
    I[i], I[j] = I[j], I[i]  
} 
</pre>
###Golang自定义结构体标签的重要应用
在Golang中首字母大小写,决定着这此变量是否能被外部调用。比如下面：
<pre>
package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	name string
	Age  int
}

func main() {

	var info T = T{"fyxichen", 24}
	fmt.Println("编码前：", info)
	b, _ := json.Marshal(info)
	fmt.Println("编码后：", string(b))

}
output==>
编码前： {fyxichen 24}           //看这里
编码后： {"Age":24}              //看这里
</pre>
在这里name的值并未被编码,原因接收首字母是小写,用json编码后不能被外部使用。这时候，自定义标签的用处凸显出来了，请看下面。
<pre>
package main

import (
	"encoding/json"
	"fmt"
)

type T1 struct {
	Name string
	Age  int
}
type T2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var info1 T1 = T1{"fyxichen", 24}
	var info2 T2 = T2{"fyxichen", 24}
	b, _ := json.Marshal(info1)
	fmt.Println("Struct1:", string(b))
	b, _ = json.Marshal(info2)
	fmt.Println("Struct2:", string(b))
}
output==>
Struct1: {"Name":"fyxichen","Age":24}	  //看这里
Struct2: {"name":"fyxichen","age":24}     //看这里
</pre>
###Golang用堆排序的方法将一千万个int随机数排序 与 用快速排序法排序性能/正确性比较
是如果用快速排序法对重复率很高的slice排序的时候,时间复杂度会激增,速度相当慢，而且并没有达到排序的目的。  
所以尝试了一下堆排序，效果不错。

二叉树的特性:  

最后一个非叶子节点 ： root = length/2(当length为奇数的时候root向下取整) 在GO语言中的索引位置：root - 1,  左右孩子节点:child_l = 2*root,索引位置:child_l-1,右孩子的节点: 2*root+1 索引位置。 
<pre>
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Num := 10000000
	var list []int
	for i := Num; i > 0; i-- {
		list = append(list, rand.Intn(10000))
	} //生成一千万个0---10000的随机数.
	length := len(list)
	for root := length/2 - 1; root >= 0; root-- {
		sort(list, root, length)
	} //第一次建立大顶堆
	for i := length - 1; i >= 1; i-- {
		list[0], list[i] = list[i], list[0]
		sort(list, 0, i)
	} //调整位置并建并从第一个root开始建堆.如果不明白为什么,大家多把图画几遍就应该明朗了
	fmt.Println(list)
}
func sort(list []int, root, length int) {
	for {
		child := 2*root + 1
		if child >= length {
			break
		}
		if child+1 < length && list[child] < list[child+1] {
			child++ //这里重点讲一下,就是调整堆的时候,以左右孩子为节点的堆可能也需要调整
		}
		if list[root] > list[child] {
			return
		}
		list[root], list[child] = list[child], list[root]
		root = child
	}
}
output=>
//占用内存313984k
...
</pre>
而用快速排序法非但没有实现排序目的，还占用了更多的内存，如下:
<pre>
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var z []int
	for i := 0; i < 10000000; i++ {
		z = append(z, rand.Intn(10000))
	}
	sort(z)
	fmt.Println(z)
}
func sort(list []int) {
	if len(list) <= 1 {
		return //退出条件
	}
	i, j := 0, len(list)-1
	index := 1     //表示第一次比较的索引位置.
	key := list[0] //第一次比较的参考值.这里选择第一个数
	if list[index] > key {
		list[i], list[j] = list[j], list[i]
		j-- //表示取大值跟末尾的数替换位置,使大于参考值的数在后面
	} else {
		list[i], list[index] = list[index], list[i]
		i++ //表示取小的值跟前面的替换位置,使小于参考值的数在前面
		index++
	}
	sort(list[:i])   //处理参考值前面的数组
	sort(list[i+1:]) //处理参考值后面的数组
}
output==>
//占用内存1022936k
...
</pre>
###Golang实现slice去重
<pre>
package main

import (
	"fmt"
)

func main() {
	a := []int{2, 1, 2, 5, 6, 3, 4, 5, 2, 3, 9}
	z := Rm_duplicate(&a)
	fmt.Println(z)
}

func Rm_duplicate(list *[]int) []int {
	var x []int = []int{}
	for _, i := range *list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}
output==>
[2 1 5 6 3 4 9]
</pre>
###Golang抓取网站磁力链接
<pre>
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Print("输入要查询的字符：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadBytes('\n')
	x := string(input[0 : len(input)-2])
	const url, page string = "http://www.btcherry.com/search?keyword=", "&p="
	var Find string
	FileResult, _ := os.OpenFile("re.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 777)
	defer func() {
		time.Sleep(1e9 * 2)
		FileResult.Sync()
		FileResult.Close()
	}()

	for i := 1; i < 101; i++ {
		Find = url + x + page + strconv.Itoa(i)
		h := strings.Repeat("#", i/2) + strings.Repeat(" ", 50-i/2)
		fmt.Printf("\r%02d%%[%s]", i, h)
		time.Sleep(1e6 * 5)
		go Resolve(Find, FileResult)
	}
}

func Resolve(Find string, FileResult io.Writer) {
	Re0, _ := regexp.Compile("<h5.*h5>")
	Re1, _ := regexp.Compile("<h5 class='h' name='rsrc'")
	Re2, _ := regexp.Compile("<span class='highlight'>")
	Re3, _ := regexp.Compile("</span")
	Re4, _ := regexp.Compile("</h5>")
	Re5, _ := regexp.Compile(">")
	Re6, _ := regexp.Compile("data-hash=")
	Resp, err := http.Get(Find)
	if err != nil {
		fmt.Println(err)
	}
	Buf, _ := ioutil.ReadAll(Resp.Body)
	buf := Re0.FindAll(Buf, 1000)
	for _, line := range buf {
		line = Re1.ReplaceAll(line, []byte(""))
		line = Re2.ReplaceAll(line, []byte(""))
		line = Re3.ReplaceAll(line, []byte(""))
		line = Re4.ReplaceAll(line, []byte(""))
		line = Re5.ReplaceAll(line, []byte(""))
		line = Re6.ReplaceAll(line, []byte("magnet:?xt=urn:btih:"))
		FileResult.Write(line)
		FileResult.Write([]byte("\n"))
	}
}
</pre>
###根据手机号判断手机运营商
<pre>
package main

import (
	"fmt"
)

var whoareyou = make(map[string]string)

func init() {

	var yidong []string = []string{"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "178", "182", "183", "184", "187", "188"}
	var liantong []string = []string{"130", "131", "132", "145", "155", "156", "176", "185", "186"}
	var dianxin []string = []string{"133", "153", "177", "180", "181", "189"}
	var qita []string = []string{"170"} //1700,1705,1709
	for i := 0; i < len(yidong); i++ {
		whoareyou[yidong[i]] = "y-移动" //移动
	}
	for i := 0; i < len(liantong); i++ {
		whoareyou[liantong[i]] = "l-联通" //联通
	}
	for i := 0; i < len(dianxin); i++ {
		whoareyou[dianxin[i]] = "d-电信" //电信
	}
	for i := 0; i < len(qita); i++ {
		whoareyou[qita[i]] = "q-其他" //其他
	}
}

func WhoAreYou(account string) string {
	key := string(account[:3])
	return whoareyou[key]
}

func main() {
	fmt.Println(WhoAreYou("18856464532"))
}
output==>
y-移动
</pre>
###Golan的defer 
defer的用法及意义

- defer 在声明时不会立即执行，而是在函数 return 后，再按照 FILO （先进后出）的原则依次执行每一个 defer，一般用于异常处理、释放资源、清理数据、记录日志等。这有点像面向对象语言的析构函数，优雅又简洁，是 Golang 的亮点之一。
- defer 还有一个重要的特性，就是即便函数抛出了异常，也会被执行的。 这样就不会因程序出现了错误，而导致资源不会释放了。
###生成Hmac
<pre>
package main

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
)

//生成Hmac
func GenerateHmac(str string) string {
	var hash = hmac.New(md5.New, []byte("tesstring"))
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))

}

func main() {
	fmt.Println(GenerateHmac("EEE"))
}
output==>
9412ddc1f9772447951ec7281ce49ec8
</pre>
###Golang时间
<pre>
package main

import (
	"fmt"
	"time"
)

func main() {
	currenttime := time.Now()     //当前时间
	begintime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2015-12-01 00:00:00", time.Local)  //格式化特定时间
	endtime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-12-31 23:59:59", time.Local)   //格式化特定时间
	fmt.Println(currenttime, begintime, endtime)
}
output==>
2016-06-12 14:07:03.2768511 +0800 CST 2015-12-01 00:00:00 +0800 CST 2016-12-31 23:59:59 +0800 CST
</pre>
###Golang对金额进行判断
<pre>
package main

import (
	"strconv"
)

//对字符串进行截取
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func OperationMoney(money int) string {
	if money >= 10000 {
		bb := strconv.Itoa(money)
		aa := Substr(strconv.Itoa(money), 0, len([]byte(bb))-4)
		return aa + "万"
	} else {
		return strconv.Itoa(money) + "元"
	}
}
func main() {
	println(OperationMoney(135634))
}
output==>
13万
</pre>
###C开发后端的技术要求
- 熟悉linux下c开发，熟悉网络、进程/线程间通信编程；
- 全面的软件知识结构（操作系统、软件工程、设计模式、数据结构、数据库系统、网络安全)。
###Beego下的api开发
注意：在使用beego的自动注解路由的时候，每个方法的上面的自动注解内容格式必须是：
<pre>
// @Title 显示标题  
// @Description get TradeRecord
// @Param	app		path 	string	true		"The app for"
// @Success 200
// @Failure 403
// @router /lastedtraderecord [get]                                                                 //这里的router 必须是小写的，写成Router的就是错误的
</pre>
然后beego就会自动在router/CommentsRouter_*.go文件中自动生成对应的restful接口代码。

在对接口返回数据进行映射解析的时候，在api端 struct设计的时候以方法（sql语句）返回数据形式为准，需要注意的是在将数据以json格式输出的时候，它的数据结构形式会影响到接收端的数据结构形式设计，因为接收端的数据结构形式会以api返回数据的形式统一。

基本结构如下：
<pre>
routers：
-------router.go
router.go -->

package routers
import (
	"github.com/astaxie/beego"
)
func init() {
	beego.GlobalControllerRouter["？？？/controllers:ActivityController"] = append(beego.GlobalControllerRouter["zcm/controllers:ActivityController"],
		beego.ControllerComments{
			"GetHomepage",       //结构体的方法名
			`/home/`,            //路径，或者说是路由
			[]string{"get"},     //访问形式及返回数据类型
			nil})
......
}

-------commentsRouter_controllers.go

package routers
import (
	"github.com/astaxie/beego"
	"zcm/controllers"
)

func init() {
	ns := beego.NewNamespace("/x1",

		beego.NSNamespace("/activity",   
			beego.NSInclude(
				&controllers.ActivityController{},    //继承自beegoController的结构体
			),
		),
......
	beego.AddNamespace(ns)
}
</pre>
###Golang实现并发get请求
goroutine实现并发
<pre>
package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
)

type RemoteResult struct {
	Url    string
	Result string
}

func RemoteGet(requestUrl string, resultChan chan RemoteResult) {
	request := httplib.NewBeegoRequest(requestUrl, "GET")
	request.SetTimeout(2*time.Second, 5*time.Second)
	//request.String()
	content, err := request.String()
	if err != nil {
		content = "" + err.Error()
	}
	resultChan <- RemoteResult{Url: requestUrl, Result: content}
}
func MultiGet(urls []string) []RemoteResult {
	fmt.Println(time.Now())
	resultChan := make(chan RemoteResult, len(urls))
	defer close(resultChan)
	var result []RemoteResult
	for _, url := range urls {
		go RemoteGet(url, resultChan)
	}
	for i := 0; i < len(urls); i++ {
		res := <-resultChan
		result = append(result, res)
	}
	fmt.Println(time.Now())
	return result
}

func main() {
	urls := []string{
		"http://baidu.com",
		"http://soso.com",
		"http://bing.com",
		"http://qq.com",
		"http://yahoo.com"}
	content := MultiGet(urls)
	fmt.Println(content)
}
output==>
	2016-06-13 09:31:44.9976743 +0800 CST
	2016-06-13 09:31:46.5737645 +0800 CST
	[{http://baidu.com <html>
	<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
	</html>
	} {http://qq.com <!DOCTYPE html>
	<html lang="zh-CN">
	<head>
	<meta content="text/html; charset=gb2312" http-equiv="Content-Type">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<title>��Ѷ��ҳ</title>
	<script type="text/javascript">
	if(window.location.toString().index
...
</pre>
###利用Golang的反射包，实现根据函数名自动调用函数
<pre>
package main

import "fmt"
import "reflect"
import "encoding/xml"

type st struct {
}

func (this *st) Echo() {
	fmt.Println("echo()")
}

func (this *st) Echo2() {
	fmt.Println("echo--------------------()")
}

var xmlstr string = `<root>  
    <func>Echo</func>  
    <func>Echo2</func>  
    </root>`

type st2 struct {
	E []string `xml:"func"`
}

func main() {
	s2 := st2{}
	xml.Unmarshal([]byte(xmlstr), &s2)

	s := &st{}
	v := reflect.ValueOf(s)

	v.MethodByName(s2.E[1]).Call(nil)
}
output==>
echo--------------------()
</pre>
###Golang遍历某个目录下的文件，并读取文件名到一个csv文件 
<pre>
package main

import (
	"container/list"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

var outputFileName string = "filesName.csv"

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}

func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func PrintFilesName(path string) {
	fullPath := GetFullPath(path)

	listStr := list.New()

	filepath.Walk(fullPath, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}

		name := fi.Name()
		if outputFileName != name {
			listStr.PushBack(name)
		}

		return nil
	})

	OutputFilesName(listStr)
}

func ConvertToSlice(listStr *list.List) []string {
	sli := []string{}
	for el := listStr.Front(); nil != el; el = el.Next() {
		sli = append(sli, el.Value.(string))
	}

	return sli
}

func OutputFilesName(listStr *list.List) {
	files := ConvertToSlice(listStr)
	//sort.StringSlice(files).Sort()// sort

	f, err := os.Create(outputFileName)
	//f, err := os.OpenFile(outputFileName, os.O_APPEND | os.O_CREATE, os.ModeAppend)
	CheckErr(err)
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(f)

	length := len(files)
	for i := 0; i < length; i++ {
		writer.Write([]string{files[i]})
	}

	writer.Flush()
}

func main() {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path, _ = os.Getwd()
	}
	PrintFilesName(path)

	fmt.Println("done!")
}
output==>
done!
</pre>
###Golang生成自定义长度的密码
<pre>
package main

import (
	"fmt"
	ran "math/rand"
	"strconv"
	"time"
)

//生成自定义长度的密码
func GetRandom(length int) string {
	r := ran.New(ran.NewSource(time.Now().UnixNano()))
	var result string
	for i := 0; i < length; i++ {

		if int(r.Intn(2))%2 == 0 {
			var choice int
			if int(r.Intn(2))%2 == 0 {
				choice = 65
			} else {
				choice = 97
			}
			result = result + string(choice+r.Intn(26))
		} else {
			result = result + strconv.Itoa(r.Intn(10))
		}
	}
	return result
}

func main() {
	fmt.Println(GetRandom(45))
}
output==>
csB9065oLYlS9Kk0GE0cJhbDwIS8247Nouo1n0541Pwwa
</pre>
###Golan将数字转换成繁体中文
<pre>
package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func AmountConvert(p_money float64, p_Round bool) string {
	var NumberUpper = []string{"壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖", "零"}
	var Unit = []string{"分", "角", "圆", "拾", "佰", "仟", "万", "拾", "佰", "仟", "亿", "拾", "佰", "仟"}
	var regex = [][]string{
		{"零拾", "零"}, {"零佰", "零"}, {"零仟", "零"}, {"零零零", "零"}, {"零零", "零"},
		{"零角零分", "整"}, {"零分", "整"}, {"零角", "零"}, {"零亿零万零元", "亿元"},
		{"亿零万零元", "亿元"}, {"零亿零万", "亿"}, {"零万零元", "万元"}, {"万零元", "万元"},
		{"零亿", "亿"}, {"零万", "万"}, {"拾零圆", "拾元"}, {"零圆", "元"}, {"零零", "零"}}
	Str, DigitUpper, UnitLen, Round := "", "", 0, 0

	if p_money == 0 {
		return "零"
	}
	if p_money < 0 {
		Str = "负"
		p_money = math.Abs(p_money)
	}
	if p_Round {
		Round = 1
	} else {
		Round = 2
	}

	Digit_byte := []byte(strconv.FormatFloat(p_money, 'f', Round+1, 64)) //注意币种四舍五入
	UnitLen = len(Digit_byte) - Round

	for _, v := range Digit_byte {
		if UnitLen >= 1 && v != 46 {
			s, _ := strconv.ParseInt(string(v), 10, 0)
			if s != 0 {
				DigitUpper = NumberUpper[s-1]

			} else {
				DigitUpper = "零"
			}
			Str = Str + DigitUpper + Unit[UnitLen-1]
			UnitLen = UnitLen - 1
		}
	}

	for i, _ := range regex {
		reg := regexp.MustCompile(regex[i][0])
		Str = reg.ReplaceAllString(Str, regex[i][1])
	}

	if string(Str[0:3]) == "元" {
		Str = string(Str[3:len(Str)])
	}

	if string(Str[0:3]) == "零" {
		Str = string(Str[3:len(Str)])
	}
	return Str
}

func main() {
	fmt.Println(AmountConvert(454565, false))
}
output==>
肆拾伍万肆仟伍佰陆拾伍圆整
</pre>
###Golang获取字符串长度
<pre>
package main

import (
	"fmt"
)

func Strlen(s string) int {
	rs := []rune(s)
	rl := len(rs)
	return rl
}
func main() {
	fmt.Println(Strlen("hello冬季"))
}
output==>
7
</pre>
###梦网语音
<pre>
type SmsResult struct {
	XMLname       xml.Name `xml:"returnsms"`
	Status        string   `xml:"returnstatus"`
	Message       string   `xml:"message"`
	Remainpoint   string   `xml:"remainpoint"`
	TaskID        string   `xml:"taskID"`
	SuccessCounts string   `xml:"successCounts"`
}

//向梦网语音平台发送验证码，获取语音验证码
func GetForSmMengWang(account string, uid int) (vc string, b bool) {
	vcode := utils.GetRand2Digit() + utils.GetRand2Digit()
	v := url.Values{}
	v.Set("userId", "??????")
	v.Set("password", "??????")
	v.Set("pszMobis", account)
	v.Set("pszMsg", vcode) //4位验证码
	v.Set("iMobiCount", "1")
	v.Set("pszSubPort", "4009917005") //回拨显示的号码，目前不支持使用，请勿输入，（可用不传），如果输入会下单不成功返回错误码 MW:1094（显示号码不合法）
	dtOrder := time.Now().Local().Format("20060102150405") + vcode
	v.Set("MsgId", dtOrder)
	v.Set("PtTmplId", "100102")
	v.Set("msgType", "1")

	cache.RecordMessages(account, "语音验证码:==>"+vcode+"")     //向zcmlc_log数据库中message加入短信日志记录
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req_url := "http://61.145.229.28:5001/voiceprepose"
	req, err := http.NewRequest("POST", req_url+"/MongateSendSubmit", body)
	if err != nil {
		cache.RecordNewLogs(uid, account, "语音验证码:==>"+vcode+"发送失败："+err.Error(), "", "SMS", "", "")
		return vcode, false
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	resp, err := client.Do(req)                                                      //发送
	if err != nil {
		cache.RecordNewLogs(uid, account, "语音验证码:==>"+vcode+"发送失败："+err.Error(), "", "SMS", "", "")
		return vcode, false
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cache.RecordNewLogs(uid, account, "语音验证码:==>"+vcode+"解析失败："+err.Error(), "", "SMS", "", "")
		return vcode, false
	}
	cache.RecordNewLogs(uid, account, "语音验证码:==>"+vcode+",申请发送语音验证码,返回的结果:==>"+string(data), "", "SMS", "", "")
	return vcode, true
}

//向梦网平台获取验证码结果
func GetStateForMengWang(account string, uid int) bool {
	v := url.Values{}
	v.Set("userId", "??????")
	v.Set("password", "??????")
	v.Set("iReqType", "2")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req_url := "http://61.145.229.28:5001/voiceprepose"
	req, err := http.NewRequest("POST", req_url+"/MongateGetDeliver", body)
	if err != nil {
		cache.RecordNewLogs(uid, account, "语音验证码状态查询:==>"+err.Error(), "", "SMS", "", "")
		return false
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	resp, err := client.Do(req)                                                      //发送
	if err != nil {
		cache.RecordNewLogs(uid, account, "语音验证码状态查询:==>"+err.Error(), "", "SMS", "", "")
		return false
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cache.RecordNewLogs(uid, account, "语音验证码状态查询解析失败:==>"+err.Error(), "", "SMS", "", "")
		return false
	}
	cache.RecordNewLogs(uid, account, "语音验证码是否接听状态:==>"+string(data), "", "SMS", "", "")
	//	beego.Emergency(string(data), err)
	return true
}
</pre>
###姓名与身份证号的正则匹配
<pre>
r, _ := regexp.Compile("^([\u4E00-\u9FA5]{2,5}(?:·[\u4E00-\u9FA5]{2,5})*)$")//应该是匹配两个或以上的汉字，\u4e00-\u9fa5是所有汉字的unicode编码范围
if !r.MatchString(cardName) { 
	errmsg = "你输入的姓名有误或暂不支持,请核对后再试~"
	return
}
r2, _ := regexp.Compile("^(\\d{15}$|^\\d{18}$|^\\d{17}(\\d|X|x))$")//身份证号匹配
if !r2.MatchString(idCard) {
	errmsg = "你输入的身份证号有误或暂不支持,请核对后再试~"
	return
}
</pre>
###Golang调用API列出window下所有运行的进程
<pre>
package main

import (
	"fmt"
	"strconv"
	"syscall"
	"unsafe"
)

type ulong int32
type ulong_ptr uintptr

type PROCESSENTRY32 struct {
	dwSize              ulong
	cntUsage            ulong
	th32ProcessID       ulong
	th32DefaultHeapID   ulong_ptr
	th32ModuleID        ulong
	cntThreads          ulong
	th32ParentProcessID ulong
	pcPriClassBase      ulong
	dwFlags             ulong
	szExeFile           [260]byte
}

func main() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
	pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
	if int(pHandle) == -1 {
		return
	}
	Process32Next := kernel32.NewProc("Process32Next")
	for {
		var proc PROCESSENTRY32
		proc.dwSize = ulong(unsafe.Sizeof(proc))
		if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
			fmt.Println("ProcessName : " + string(proc.szExeFile[0:20]))![](ProcessName : [System Process])
			fmt.Println("ProcessID : " + strconv.Itoa(int(proc.th32ProcessID)))
		} else {
			break
		}
	}
	CloseHandle := kernel32.NewProc("CloseHandle")
	_, _, _ = CloseHandle.Call(pHandle)
}
output==>
ProcessName : [System Process]
ProcessID : 0
ProcessName : System
...
</pre>
###Golang获取当前毫秒时间戳
<pre>
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000, 10)
	fmt.Println(timestamp)
}
output==>
1465894621538550
</pre>
###Shell脚本
<pre>
比较字符写法：
-eq 等于
-ne 不等于
-gt 大于
-lt 小于
-le 小于等于
-ge 大于等于
-z 空串
* = 两个字符相等
* != 两个字符不等
* -n 非空串
举例：
if [ $# -ge 2 ];then
echo "你提供的参数过多！"
exit 1
else
 if [ $# -ne 0 ];then
  if [ $1 = s ];then
  echo "one"
  else
  echo "two"
  fi
 else
 echo "3"
 fi
fi
</pre>
###Golang生成图片验证码
//图片验证码工具
<per>
package utils

import (
	"math/rand"
	"strconv"
	"time"
	"github.com/dchest/captcha"
)

// 生成4位的数字图片验证码
// return:验证码字符串,验证码图片
func GeneratePicCode() (string, *captcha.Image) {
	rand.Seed(time.Now().UnixNano())
	var code = rand.Intn(8999) + 1000
	var codestr = strconv.Itoa(code)
	var codeArray = make([]byte, 4)
	for i := 0; i < 4; i++ {
		codeArray[3-i] = byte(code % 10)
		code /= 10
	}
	return codestr, captcha.NewImage("code", codeArray, 188, 72)
}
</pre>
使用它
<pre>
//获取登录图形验证码
func (this *LoginController) PictureCode() {
	var code, img = utils.GeneratePicCode()
	this.SetSession("loginCode", code)
	this.SetSession("loginCodeDeadline", time.Now().Add(time.Minute*3).Unix())
	img.WriteTo(this.Ctx.ResponseWriter) //输出到屏幕
}
</pre>
###Golang常用加密工具包/3DES/BASE64/MD5
<pre>
//加密工具类，用了3des和base64,md5
package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

const key = `dg.d!ehtg_78nmb?#r5ew1q!`

//若解码出错，返回空字符串 手机号解密
func DesCode(acc string) (account string) {
	defer func() {
		err := recover()
		if err != nil {
			account = ""
		}
	}()
	accbyt := []byte(acc)
	accbyt, _ = DesBase64Decrypt(accbyt)
	account = string(accbyt)
	return
}

//若解码出错，返回空字符串 手机号加密
func EncCode(acc string) (account string) {
	defer func() {
		err := recover()
		if err != nil {
			account = ""
		}
	}()
	accbyt := []byte(acc)
	accbyt, _ = DesBase64Encrypt(accbyt)
	account = string(accbyt)
	return
}

//md5加密
func MD5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

//des3 + base64 encrypt
func DesBase64Encrypt(origData []byte) ([]byte, error) {
	result, err := TripleDesEncrypt(origData, []byte(key))
	if err != nil {
		return nil, err
	}
	return []byte(base64.StdEncoding.EncodeToString(result)), nil
}

func DesBase64Decrypt(crypted []byte) ([]byte, error) {
	result, _ := base64.StdEncoding.DecodeString(string(crypted))
	origData, err := TripleDesDecrypt(result, []byte(key))
	if err != nil {
		return nil, err
	}
	return origData, nil
}

// 3DES加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("WTF-------------", err)
			panic("参数异常")
		}
	}()
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
</pre>
###Golang打印当前行数，文件名与函数名
<pre>
package main

import (
	"fmt"
	"runtime"
)

func main() {
	funcName, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Println("FuncName : " + runtime.FuncForPC(funcName).Name())
		fmt.Printf("file : %s\nline : %d", file, line)
	}
}
output==>
FuncName : main.main
file : D:/gopath/src/test/test.go
line : 9
</pre>
###Golang统计某一目录下所有文件代码行数并且统计总数
<pre>
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	linesum int
	mutex   *sync.Mutex = new(sync.Mutex)
)

var (
	// the dir where souce file stored
	rootPath string = "/go/src"
	// exclude these sub dirs
	nodirs [5]string = [...]string{"/bitbucket.org", "/github.com", "/goplayer", "/uniqush", "/code.google.com"}
	// the suffix name you care
	suffixname string = ".go"
)

func main() {
	argsLen := len(os.Args)
	if argsLen == 2 {
		rootPath = os.Args[1]
	} else if argsLen == 3 {
		rootPath = os.Args[1]
		suffixname = os.Args[2]
	}
	// sync chan using for waiting
	done := make(chan bool)
	go codeLineSum(rootPath, done)
	<-done

	fmt.Println("total line:", linesum)
}

// compute souce file line number
func codeLineSum(root string, done chan bool) {
	var goes int              // children goroutines number
	godone := make(chan bool) // sync chan using for waiting all his children goroutines finished
	isDstDir := checkDir(root)
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("root: %s, panic:%#v\n", root, pan)
		}

		// waiting for his children done
		for i := 0; i < goes; i++ {
			<-godone
		}

		// this goroutine done, notify his parent
		done <- true
	}()
	if !isDstDir {
		return
	}

	rootfi, err := os.Lstat(root)
	checkerr(err)

	rootdir, err := os.Open(root)
	checkerr(err)
	defer rootdir.Close()

	if rootfi.IsDir() {
		fis, err := rootdir.Readdir(0)
		checkerr(err)
		for _, fi := range fis {
			if strings.HasPrefix(fi.Name(), ".") {
				continue
			}
			goes++
			if fi.IsDir() {
				go codeLineSum(root+"/"+fi.Name(), godone)
			} else {
				go readfile(root+"/"+fi.Name(), godone)
			}
		}
	} else {
		goes = 1 // if rootfi is a file, current goroutine has only one child
		go readfile(root, godone)
	}
}

func readfile(filename string, done chan bool) {
	var line int
	isDstFile := strings.HasSuffix(filename, suffixname)
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("filename: %s, panic:%#v\n", filename, pan)
		}
		if isDstFile {
			addLineNum(line)
			fmt.Printf("file %s complete, line = %d\n", filename, line)
		}
		// this goroutine done, notify his parent
		done <- true
	}()
	if !isDstFile {
		return
	}

	file, err := os.Open(filename)
	checkerr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {
			line++
		}
	}
}

// check whether this dir is the dest dir
func checkDir(dirpath string) bool {
	// 判断该文件夹是否在被排除的范围之内
	for _, dir := range nodirs {
		if rootPath+dir == dirpath {
			return false
		}
	}
	return true
}

func addLineNum(num int) {
	// 获取锁
	mutex.Lock()
	// defer语句在函数返回时调用, 确保锁被释放
	defer mutex.Unlock()
	linesum += num
}

// if error happened, throw a panic, and the panic will be recover in defer function
func checkerr(err error) {
	if err != nil {
		// 在发生错误时调用panic, 程序将立即停止正常执行, 开始沿调用栈往上抛, 直到遇到recover
		// 对于java程序员, 可以将panic类比为exception, 而recover则是try...catch
		panic(err.Error())
	}
}
output==>
file /go/src/math/unsafe.go complete, line = 21
file /go/src/bufio/scan_test.go complete, line = 544
file /go/src/mime/type_windows.go complete, line = 41
file /go/src/net/url/url_test.go complete, line = 1439
file /go/src/math/nextafter.go complete, line = 49
file /go/src/math/pow.go complete, line = 137
file /go/src/math/pow10.go complete, line = 40
file /go/src/builtin/builtin.go complete, line = 256
file /go/src/math/rand/zipf.go complete, line = 77
file /go/src/math/remainder.go complete, line = 85
file /go/src/math/signbit.go complete, line = 10
file /go/src/math/sin.go complete, line = 224
file /go/src/math/sincos.go complete, line = 69
file /go/src/math/sinh.go complete, line = 77
file /go/src/math/sqrt.go complete, line = 148
file /go/src/math/tan.go complete, line = 130
file /go/src/math/tanh.go complete, line = 97
file /go/src/mime/encodedword.go complete, line = 434
file /go/src/mime/encodedword_test.go complete, line = 208
file /go/src/mime/example_test.go complete, line = 98
file /go/src/mime/grammar.go complete, line = 32
file /go/src/mime/mediatype.go complete, line = 361
file /go/src/mime/mediatype_test.go complete, line = 311
file /go/src/mime/multipart/writer_test.go complete, line = 128
file /go/src/mime/quotedprintable/writer_test.go complete, line = 158
file /go/src/mime/type.go complete, line = 187
file /go/src/mime/type_dragonfly.go complete, line = 9
......
total line: 820210
</pre>
###Golang中http读取大文件必须读完 
<pre>
package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://mirrors.ustc.edu.cn/opensuse/distribution/12.3/iso/openSUSE-12.3-GNOME-Live-i686.iso", nil)
	req.Header.Set("Connection", "close")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Resp code:", resp.StatusCode)
	resp.Body.Close()
}
output==>
Resp code: 200
</pre>
###Golang channel
Golang Channel的基本操作语法如下：

- c := make(chan bool) //创建一个无缓冲的bool型Channel 
- c <- x        //向一个Channel发送一个值
- <- c          //从一个Channel中接收一个值
- x = <- c      //从Channel c接收一个值并将其存储到x中
- x, ok = <- c  //从Channel接收一个值，如果channel关闭了或没有数据，那么ok将被置为false

不带缓冲的Channel兼具通信和同步两种特性，颇受青睐。
<pre>
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin doing something")
	c := make(chan bool)
	go func() {
		fmt.Println("Doing something ...")
		c <- true //事情做完后向channel中写入值来作为通知
	}()
	<-c //等待读取channel中的值，读不到一直阻塞进程
	fmt.Println("Done!")
}
output==>
Begin doing something
Doing something ...
Done!
</pre>
多个goroutine
<pre>
package main

import (
	"fmt"
)

func worker(start chan bool, index int) {
	<-start
	fmt.Println("This  is  worker:", index)
}
func main() {
	start := make(chan bool)
	for i := 1; i <= 10; i++ {
		go worker(start, i)
		start <- true
	}
}
output==>
This  is  worker: 1
This  is  worker: 2
This  is  worker: 3
This  is  worker: 4
This  is  worker: 5
This  is  worker: 6
This  is  worker: 7
This  is  worker: 8
This  is  worker: 9
This  is  worker: 10
</pre>
###Golang发送邮件给多人
<pre>
func SendEmaliToUsers(user, password, content, title string) {
	host := "smtp.exmail.qq.com:25"
	to := strings.Split(user, ";") //收件人用;号隔开
	content_type := "Content-Type: text/plain" + "; charset=UTF-8"
	str := content // 邮件内容
	msg := []byte("To: cuinan\r\nFrom: " + user + ">\r\nSubject:" + title + "\r\n" + content_type + "\r\n" + str)
	err := smtp.SendMail(host, smtp.PlainAuth("", "???@???.com", password, "smtp.exmail.qq.com"), "???@???.com", to, []byte(msg))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
</pre>
###Golang中的ecb/ECB加密
<pre>
package main

import (
	"bytes"
	"crypto/des"
	"errors"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("程序开始....")
	key := []byte{0xD5, 0x92, 0x86, 0x02, 0x2A, 0x0B, 0x3E, 0x64}
	data := []byte("hello world")

	out, _ := MyEncrypt(data, key)
	log.Println("加密后:", out)
	out, _ = MyDecrypt(out, key)
	log.Println("解密后:", string(out))
}
func MyEncrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out, nil
}
func MyDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return out, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
output==>
2016/06/15 17:01:21 test.go:31: 程序开始....
2016/06/15 17:01:21 test.go:36: 加密后: [169 102 193 176 238 138 72 19 199 11 51 238 250 56 193 150]
2016/06/15 17:01:21 test.go:38: 解密后: hello world
</pre>
下面是一个ecb封装的工具方法
<pre>
package utils

import (
	"crypto/cipher"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
</pre>
###Golang解析sina登录页
<pre>
package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func loginPre1() map[string]interface{} {

	client := &http.Client{}

	reqest, err := http.NewRequest("GET", "http://login.sina.com.cn/sso/prelogin.php?entry=weibo&callback=sinaSSOController.preloginCallBack&su=Z3V5dWV0ZnRiJTQwMTYzLmNvbQ%3D%3D&rsakt=mod&checkpin=1&client=ssologin.js(v1.4.5)&_=", nil)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Encoding", "gzip, deflate")
	reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Host", "login.sina.com.cn")
	reqest.Header.Add("Referer", "http://weibo.com/")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	response, err := client.Do(reqest)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	if response.StatusCode == 200 {

		var body string

		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					panic(err)
				}

				if n == 0 {
					break
				}
				body += string(buf)
			}
		default:
			bodyByte, _ := ioutil.ReadAll(response.Body)
			body = string(bodyByte)
		}

		r := regexp.MustCompile(`sinaSSOController.preloginCallBack\((.*?)\)`)
		rs := r.FindStringSubmatch(body)

		//json decode
		header := make(map[string]interface{})
		err = json.Unmarshal([]byte(rs[1]), &header)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			os.Exit(0)
		}

		t := fmt.Sprintf("%f", header["servertime"])

		header["servertime"] = strings.Trim(t, ".000000")

		return header
	}

	return nil
}

func main() {
	res := loginPre1()
	for k, v := range res {
		fmt.Println(k, "==>", v)
	}
}
output==>
servertime ==> 1465991817
pcid ==> gz-28505aff1e897f38d4ba0cad5eca27acc317
is_openlock ==> 0
showpin ==> 1
retcode ==> 0
nonce ==> 6FIJCZ
pubkey ==> EB2A38568661887FA180BDDB5CABD5F21C7BFD59C090CB2D245A87AC253062882729293E5506350508E7F9AA3BB77F4333231490F915F6D63C55FE2F08A49B353F444AD3993CACC02DB784ABBB8E42A9B1BBFFFB38BE18D78E87A0E41B9B8F73A928EE0CCEE1F6739884B9777E4FE9E88A1BBE495927AC4A799B3181D6442443
rsakv ==> 1330428213
exectime ==> 357
</pre>
###Golang实现权重轮询调度算法/WRRS
<pre>
package main

import (
	"fmt"
	"time"
)

var slaveDns = map[int]map[string]interface{}{
	0: {"connectstring": "root@tcp(172.16.0.2:3306)/shiqu_tools?charset=utf8", "weight": 2},
	1: {"connectstring": "root@tcp(172.16.0.4:3306)/shiqu_tools?charset=utf8", "weight": 4},
	2: {"connectstring": "root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8", "weight": 8},
}

var i int = -1  //表示上一次选择的服务器
var cw int = 0  //表示当前调度的权值
var gcd int = 2 //当前所有权重的最大公约数 比如 2，4，8 的最大公约数为：2

func getDns() string {
	for {
		i = (i + 1) % len(slaveDns)
		if i == 0 {
			cw = cw - gcd
			if cw <= 0 {
				cw = getMaxWeight()
				if cw == 0 {
					return ""
				}
			}
		}

		if weight, _ := slaveDns[i]["weight"].(int); weight >= cw {
			return slaveDns[i]["connectstring"].(string)
		}
	}
}

func getMaxWeight() int {
	max := 0
	for _, v := range slaveDns {
		if weight, _ := v["weight"].(int); weight >= max {
			max = weight
		}
	}

	return max
}

func main() {

	note := map[string]int{}

	s_time := time.Now().Unix()

	for i := 0; i < 20; i++ {
		s := getDns()
		fmt.Println(s)
		if note[s] != 0 {
			note[s]++
		} else {
			note[s] = 1
		}
	}

	e_time := time.Now().Unix()
	fmt.Println("s_time", s_time)
	fmt.Println("e_time", e_time)
	fmt.Println("total time: ", e_time-s_time)

	fmt.Println("--------------------------------------------------")

	for k, v := range note {
		fmt.Println(k, " ", v)
	}
}
output==>
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.4:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.2:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.4:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.4:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.2:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.4:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.4:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.2:3306)/shiqu_tools?charset=utf8
root@tcp(172.16.0.4:3306)/shiqu_tools?charset=utf8
s_time 1465992539
e_time 1465992539
total time:  0
--------------------------------------------------
root@tcp(172.16.0.8:3306)/shiqu_tools?charset=utf8   11
root@tcp(172.16.0.4:3306)/shiqu_tools?charset=utf8   6
root@tcp(172.16.0.2:3306)/shiqu_tools?charset=utf8   3
</pre>
###发送带附件的邮件
<pre>
package main
/*
http://blog.csdn.net/xcl168/article/details/51340272
*/
import (
	"bytes"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

const (
	emlUser = "？？？？？@163.com"
	emlPwd  = "？？？？？？？"
	emlSMTP = "smtp.163.com:25"
)

func main() {

	err := eml()
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("发送成功")
	}

}

// 发普通文本邮件
func eml() error {

	to := "？？？@？？？.com"
	sendTo := strings.Split(to, ";")
	subject := "带附件的邮件"
	mime := bytes.NewBuffer(nil)

	//设置邮件
	mime.WriteString(fmt.Sprintf("From: %s<%s>\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\n", emlUser, emlUser, to, subject))
	mime.WriteString("Content-Description: 这是一封带附件的邮件\r\n")

	//附件
	mime.WriteString("Content-Type: text/plain\r\n")
	mime.WriteString("Content-Description: 附一个Text文件\r\n")
	mime.WriteString("Content-Disposition: attachment; filename=\"test.txt\"\r\n\r\n")
	mime.WriteString("这是写入test.txt文件的内容")
	//发送
	smtpHost, _, err := net.SplitHostPort(emlSMTP)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", emlUser, emlPwd, smtpHost)
	return smtp.SendMail(emlSMTP, auth, emlUser, sendTo, mime.Bytes())
}
output==>
发送成功
</pre>
###rot系列加密

- ROT5：只对数字进行编码，用当前数字往前数的第5个数字替换当前数字，例如当前为0，编码后变成5，当前为1，编码后变成6，以此类推顺序循环。
- ROT13：只对字母进行编码，用当前字母往前数的第13个字母替换当前字母，例如当前为A，编码后变成N，当前为B，编码后变成O，以此类推顺序循环。
- ROT18：这是一个异类，本来没有，它是将ROT5和ROT13组合在一起，为了好称呼，将其命名为ROT18。
- ROT47：对数字、字母、常用符号进行编码，按照它们的ASCII值进行位置替换，用当前字符ASCII值往前数的第47位对应字符替换当前字符，例如当前为小写字母z，编码后变成大写字母K，当前为数字0，编码后变成符号_。
###Goalng常用正则式 
匹配中文字符的正则表达式： [\u4e00-\u9fa5]
 
匹配双字节字符(包括汉字在内)：[^\x00-\xff]
 
匹配空行的正则表达式：\n[\s| ]*\r
 
匹配HTML标记的正则表达式：/<(.*)>.*<\/\1>|<(.*) \/>/
 
匹配首尾空格的正则表达式：(^\s*)|(\s*$)
 
匹配IP地址的正则表达式：/(\d+)\.(\d+)\.(\d+)\.(\d+)/g //
 
匹配Email地址的正则表达式：\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*
 
匹配网址URL的正则表达式：<a href="http://%28/" target="_blank">http://(/</a>[\w-]+\.)+[\w-]+(/[\w- ./?%&=]*)?
 
 sql语句：^(select|drop|delete|create|update|insert).*$
 
 1、非负整数：^\d+$
 
 2、正整数：^[0-9]*[1-9][0-9]*$
 
 3、非正整数：^((-\d+)|(0+))$
 
 4、负整数：^-[0-9]*[1-9][0-9]*$
 
 5、整数：^-?\d+$
 
 6、非负浮点数：^\d+(\.\d+)?$
 
 7、正浮点数：^((0-9)+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*))$
 
 8、非正浮点数：^((-\d+\.\d+)?)|(0+(\.0+)?))$
 
 9、负浮点数：^(-((正浮点数正则式)))$
 
10、英文字符串：^[A-Za-z]+$
 
 11、英文大写串：^[A-Z]+$
 
 12、英文小写串：^[a-z]+$
 
 13、英文字符数字串：^[A-Za-z0-9]+$
 
 14、英数字加下划线串：^\w+$
 
 15、E-mail地址：^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$
 
16、URL：^[a-zA-Z]+://(\w+(-\w+)*)(\.(\w+(-\w+)*))*(\?\s*)?$
或：^http:\/\/[A-Za-z0-9]+\.[A-Za-z0-9]+[\/=\?%\-&_~`@[\]\':+!]*([^<>\"\"])*$
 
17、邮政编码：^[1-9]\d{5}$
 
 18、中文：^[\u0391-\uFFE5]+$
 
 19、电话号码：^((\(\d{2,3}\))|(\d{3}\-))?(\(0\d{2,3}\)|0\d{2,3}-)?[1-9]\d{6,7}(\-\d{1,4})?$
 
 20、手机号码：^((\(\d{2,3}\))|(\d{3}\-))?13\d{9}$
 
 21、双字节字符(包括汉字在内)：^\x00-\xff
 
 22、匹配首尾空格：(^\s*)|(\s*$)（像vbscript那样的trim函数）
 
23、匹配HTML标记：<(.*)>.*<\/\1>|<(.*) \/>
 
 24、匹配空行：\n[\s| ]*\r
 
 25、提取信息中的网络链接：(h|H)(r|R)(e|E)(f|F) *= *('|")?(\w|\\|\/|\.)+('|"| *|>)?
 
 26、提取信息中的邮件地址：\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*
 
 27、提取信息中的图片链接：(s|S)(r|R)(c|C) *= *('|")?(\w|\\|\/|\.)+('|"| *|>)?
 
 28、提取信息中的IP地址：(\d+)\.(\d+)\.(\d+)\.(\d+)
 
 29、提取信息中的中国手机号码：(86)*0*13\d{9}
 
 30、提取信息中的中国固定电话号码：(\(\d{3,4}\)|\d{3,4}-|\s)?\d{8}
 
 31、提取信息中的中国电话号码（包括移动和固定电话）：(\(\d{3,4}\)|\d{3,4}-|\s)?\d{7,14}
 
 32、提取信息中的中国邮政编码：[1-9]{1}(\d+){5}
 
 33、提取信息中的浮点数（即小数）：(-?\d*)\.?\d+
 
 34、提取信息中的任何数字 ：(-?\d*)(\.\d+)?
 
 35、IP：(\d+)\.(\d+)\.(\d+)\.(\d+)
 
 36、电话区号：/^0\d{2,3}$/
 
 37、腾讯QQ号：^[1-9]*[1-9][0-9]*$
 
 38、帐号(字母开头，允许5-16字节，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]{4,15}$
 
 39、中文、英文、数字及下划线：^[\u4e00-\u9fa5_a-zA-Z0-9]+$

###LINUX服务器项目管理
<pre>
cd /home/go/src/项目名称
bee run 
ps -ef |grep 项目名称
kill -9 进程ID
bee run 
nohup ./项目名称 &
</pre>
/*不挂断地启动zcm进程,注意：运行完nohup命令后用exit退出终端，而不是直接点关闭，
因为这样会删除该命令所对应的session，导致nohup对应的进程被通知需要一起shutdown*/

tail -f 文件名

linux tail命令用途是依照要求将指定的文件的最后部分输出到标准设备，通常是终端，
通俗讲来，就是把某个档案文件的最后几行显示到终端上，假设该档案有更新，
tail会自己主动刷新，确保你看到最新的档案内容。

ps 将某个进程显示出来

-A 显示所有进程

-e 效果同上，显示所有进程

-f 显示UID/PPID/C/STIME 栏位


grep 查找

| 管道命令，是指ps与grep同时执行

UID 程序被该 UID 所拥有

PID 就是这个程序的 ID

PPID 则是其上级父程序的ID

C CPU 使用的资源百分比

STIME 系统启动时间


kill

下面是常用的信号：

HUP     1    终端断线

INT     2    中断（同 Ctrl + C）

QUIT    3    退出（同 Ctrl + \）

KILL    9    强制终止  (最常用)

TERM    15    终止

CONT    18    继续（与STOP相反， fg/bg命令）

STOP    19    暂停（同 Ctrl + Z）
###Golang判断电脑系统
runtime.GOOS
runtime.GOROOT()
<pre>
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO run on")
	fmt.Println("Goroot:",runtime.GOROOT())
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println(os)
	}
}
output==>
GO run on
Goroot:D:/go
Windows
</pre>
###Golang实现查找全球ip地址|离线版|没网可查哟
<pre>
package main

import (
	"fmt"
	"github.com/slene/iploc"
	"os"
	"path/filepath"
	. "testing"
)

func init() {
	// replace iplocFilePath to your iploc.dat path
	iplocFilePath, _ := filepath.Abs("../github.com/slene/iploc/iploc.dat")
	// simple set a true param can preload all ipinfo
	// need allocate more memory > 30M
	// and speed can grow up about 40 percent than not preload
	iploc.IpLocInit(iplocFilePath, true)
}

func testIp(ipAddr string) {
	ipinfo, err := iploc.GetIpInfo(ipAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(ipAddr)

	switch ipinfo.Flag {
	case iploc.FLAG_INUSE:
		if ipinfo.Code == "CN" {
			fmt.Println(ipinfo.Code)
			fmt.Println(ipinfo.Country)
			fmt.Println(ipinfo.Region)
			fmt.Println(ipinfo.City)
			fmt.Println(ipinfo.Isp)
		} else {
			fmt.Println(ipinfo.Code)
			fmt.Println(ipinfo.Country)
		}
	case iploc.FLAG_RESERVED:
		fmt.Println(ipinfo.Note)
	case iploc.FLAG_NOTUSE:
		fmt.Println(ipinfo.Note)
	}

	for i := 0; i < 30; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

func testSpeed() {
	r := Benchmark(func(b *B) {
		ips := []string{
			"0.0.0.0",
			"127.0.0.1",
			"169.254.0.1",
			"192.168.1.1",
			"10.0.0.0",
			"255.255.255.255",
			"112.226.155.1",
			"121.18.72.0",
			"6.18.72.0",
			"200.18.72.0",
		}
		for i := 0; i < b.N; i++ {
			for _, ipAddr := range ips {
				iploc.GetIpInfo(ipAddr)
			}
		}
	})
	fmt.Println(r)
	fmt.Printf("10w次查询: %.1f 毫秒\n", float64(r.NsPerOp())/100000000*1000*100000/10)
}

func main() {
	testIp("0.0.0.0")
	testIp("127.0.0.1")
	testIp("169.254.0.1")
	testIp("192.168.1.1")
	testIp("10.0.0.0")
	testIp("255.255.255.255")
	testIp("112.226.155.1")
	testIp("121.18.72.0")

	testSpeed()
}
output==>
0.0.0.0
IANA保留作为特殊地址
------------------------------
127.0.0.1
IANA保留用于本机回环地址
------------------------------
169.254.0.1
IANA保留作为链路本地地址
------------------------------
192.168.1.1
IANA保留用于局域网地址
------------------------------
10.0.0.0
IANA保留用于内部网络地址
------------------------------
255.255.255.255
IANA保留地址
------------------------------
112.226.155.1
CN
中国
山东省
青岛市
联通
------------------------------
121.18.72.0
CN
中国
河北省
保定市
联通
------------------------------
  500000	      2848 ns/op
10w次查询: 284.8 毫秒
</pre>
###Golang的template包
<pre>
package main

import (
	"os"
	"text/template"
)

func main() {
	name := "jason"
	tmpl, err := template.New("AnythingIsOk").Parse("hello,{{.}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		panic(err)
	}
}
output==>
hello,jason
</pre>
###Golang将Unicode转换成字符串string
<pre>
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	str := `\u5bb6\u65cf`
	fmt.Println(u2s(str))
}

func u2s(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}
output==>
家族 <nil>
</pre>
###Golang实现长轮询，实现消息的发送与接收
main.go
<pre>
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var mc *MessageCenter

type Message struct {
	Uid     int
	Message string
}

type MessageCenter struct {
	// 测试 没有加读写锁
	messageList []*Message
	userList    map[int]chan string
}

func NewMessageCenter() *MessageCenter {
	mc := new(MessageCenter)
	mc.messageList = make([]*Message, 0, 100)
	mc.userList = make(map[int]chan string)
	return mc
}

func (mc *MessageCenter) GetMessage(uid int) []string {
	messages := make([]string, 0, 10)
	for i, msg := range mc.messageList {
		if msg == nil {
			continue
		}
		if msg.Uid == uid {
			messages = append(messages, msg.Message)
			// 临时方案 只是测试用 应更换为list
			mc.messageList[i] = nil
		}
	}
	return messages
}

func (mc *MessageCenter) GetMessageChan(uid int) <-chan string {
	messageChan := make(chan string)
	mc.userList[uid] = messageChan
	return messageChan
}

func (mc *MessageCenter) SendMessage(uid int, message string) {
	messageChan, exist := mc.userList[uid]
	if exist {
		messageChan <- message
		return
	}
	// 未考虑同一账号多登陆情况
	mc.messageList = append(mc.messageList, &Message{uid, message})
}

func (mc *MessageCenter) RemoveUser(uid int) {
	_, exist := mc.userList[uid]
	if exist {
		delete(mc.userList, uid)
	}
}

func IndexServer(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "longpoll.html")
}

func SendMessageServer(w http.ResponseWriter, req *http.Request) {
	uid, _ := strconv.Atoi(req.FormValue("uid"))
	message := req.FormValue("message")

	mc.SendMessage(uid, message)

	io.WriteString(w, `{}`)
}

func PollMessageServer(w http.ResponseWriter, req *http.Request) {
	uid, _ := strconv.Atoi(req.FormValue("uid"))

	messages := mc.GetMessage(uid)

	if len(messages) > 0 {
		jsonData, _ := json.Marshal(map[string]interface{}{"status": 0, "messages": messages})
		w.Write(jsonData)
		return
	}

	messageChan := mc.GetMessageChan(uid)

	select {
	case message := <-messageChan:
		jsonData, _ := json.Marshal(map[string]interface{}{"status": 0, "messages": []string{message}})
		w.Write(jsonData)
	case <-time.After(10 * time.Second):
		mc.RemoveUser(uid)
		jsonData, _ := json.Marshal(map[string]interface{}{"status": 1, "messages": nil})
		n, err := w.Write(jsonData)
		fmt.Println(n, err)
	}
}

func main() {
	fmt.Println("http://127.0.0.1:89/")

	mc = NewMessageCenter()

	http.HandleFunc("/", IndexServer)
	http.HandleFunc("/sendmessage", SendMessageServer)
	http.HandleFunc("/pollmessage", PollMessageServer)
	err := http.ListenAndServe(":89", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
</pre>
longpoll.html
<pre>
<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>long-polling</title>
    <style type="text/css">
        .msg {padding: 10px;margin-bottom: 10px;border: 1px solid #ccc;border-radius: 8px;}
    </style>
    <script type="text/javascript" src="http://www.v2ex.com/static/js/jquery.js"></script>
    <script type="text/javascript">
        $(function () {
            $('#uid').val(Date.now() % 10000);
            var setTimeoutId = 0;
            var ajax = null;
            var getmessage = function() {
                var data = {uid:$('#uid').val()};
                ajax = $.getJSON('/pollmessage', data, function(resp) {
                    if (resp.status == 0) {
                        for (var i = 0; i < resp.messages.length; i++) {
                            $('#messagelist').append('<div class="msg">'+resp.messages[i]+'</div>');
                        };
                    }
                    if (setTimeoutId > 0) {
                        setTimeoutId = setTimeout(getmessage, 3000);
                    }
                });
                console.dir(ajax);
            };
            $('#getmessagebtn').click(function(){
                this.disabled = true;
                setTimeoutId = setTimeout(getmessage, 10);
            });
            $('#sendmessagebtn').click(function(){
                var data = {uid:$('#senduid').val(), 'message':$('#message').val()};
                $.post('/sendmessage', data, function(resp){}, 'json');
            });
            $('#stopgetmessagebtn').click(function(){
                clearTimeout(setTimeoutId);
                setTimeoutId = 0
                if (ajax != null) {
                    ajax.abort();
                }
                $('#getmessagebtn').prop('disabled', false);
            });
        });
    </script>
</head>
<body>
Send User ID: <input type="number" id="senduid" /> Message: <input type="text" id="message" /> <button id="sendmessagebtn">发送消息</button>
<hr/>
RecvUser ID: <input type="number" id="uid" /> <button id="getmessagebtn">接收消息</button> <button id="stopgetmessagebtn">停止接收消息</button>
<div id="messagelist"></div>
</body>
</html>
</pre>
###判断channel是否关闭
<pre>
package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	for {
		i, isClose := <-c
		if !isClose {
			fmt.Println("channel Closed")
			break
		} else {
			fmt.Println(i)
		}
	}
}
output==>
1
2
3
channel Closed
</pre>
###Golang认证http
<pre>
package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	auth := req.Header.Get("Authorization")
	if auth == "" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println(auth)

	auths := strings.SplitN(auth, " ", 2)
	if len(auths) != 2 {
		fmt.Println("error")
		return
	}

	authMethod := auths[0]
	authB64 := auths[1]

	switch authMethod {
	case "Basic":
		authstr, err := base64.StdEncoding.DecodeString(authB64)
		if err != nil {
			fmt.Println(err)
			io.WriteString(w, "Unauthorized!\n")
			return
		}
		fmt.Println(string(authstr))

		userPwd := strings.SplitN(string(authstr), ":", 2)
		if len(userPwd) != 2 {
			fmt.Println("error")
			return
		}

		username := userPwd[0]
		password := userPwd[1]

		fmt.Println("Username:", username)
		fmt.Println("Password:", password)
		fmt.Println()

	default:
		fmt.Println("error")
		return
	}

	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
</pre>
###Golang获取上传文件大小
<pre>
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 获取文件大小的接口
type Size interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
	if "POST" == r.Method {
		file, _, err := r.FormFile("userfile")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if statInterface, ok := file.(Stat); ok {
			fileInfo, _ := statInterface.Stat()
			fmt.Fprintf(w, "上传文件的大小为: %d", fileInfo.Size())
		}
		if sizeInterface, ok := file.(Size); ok {
			fmt.Fprintf(w, "上传文件的大小为: %d", sizeInterface.Size())
		}

		return
	}

	// 上传页面
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	html := `
<form enctype="multipart/form-data" action="/hello" method="POST">
    Send this file: <input name="userfile" type="file" />
    <input type="submit" value="Send File" />
</form>
`
	io.WriteString(w, html)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
</pre>
###Golang template自定义函数
<pre>
package main

import (
	"os"
	"text/template"
	"time"
)

type User struct {
	Username, Password string
	RegTime            time.Time
}

func ShowTime(t time.Time, format string) string {
	return t.Format(format)
}

func main() {
	u := User{"dotcoo", "dotcoopwd", time.Now()}    
	//自定义函数
	t, err := template.New("text").Funcs(template.FuncMap{"showtime": ShowTime}).
		Parse(`<p>{{.Username}}|{{.Password}}|{{.RegTime.Format "2006-01-02 15:04:05"}}</p>
<p>{{.Username}}|{{.Password}}|{{showtime .RegTime "2006-01-02 15:04:05"}}</p>
`)
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, u)
}
output==>
<p>dotcoo|dotcoopwd|2016-06-18 16:01:30</p>
<p>dotcoo|dotcoopwd|2016-06-18 16:01:30</p>
</pre>
###Golang计算经纬度之间的距离
<pre>
package main

import (
	"fmt"
	"math"
)

func main() {
	lat1 := 29.490295
	lng1 := 106.486654

	lat2 := 29.615467
	lng2 := 106.581515
	fmt.Println(EarthDistance(lat1, lng1, lat2, lng2))
}

// 返回值的单位为米
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6371000) // 6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))

	return dist * radius
}
output==>
16670.904273268756
</pre>
###Golang版ip2long|long2ip
<pre>
package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
)

func Ip2long(ipstr string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)

	return
}

func Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

//AddrToUint32
func AddrToUint32(addr net.Addr) (uint32, error) {
	var ip net.IP

	switch ipaddr := addr.(type) {
	case *net.IPAddr:
		ip = ipaddr.IP
	case *net.IPNet:
		ip = ipaddr.IP
	case *net.TCPAddr:
		ip = ipaddr.IP
	case *net.UDPAddr:
		ip = ipaddr.IP
	case *net.UnixAddr:
		return 0, errors.New("UnixAddr type not support")
	default:
		return 0, errors.New("addr type not support")
	}
	return binary.BigEndian.Uint32(ip.To4()), nil
}
func main() {
	ip := "12.67.85.145"
	longstr := uint32(205739409)
	fmt.Println("ip2long:",Ip2long(ip))
	fmt.Println("long2ip:",Long2ip(longstr))
}
output==>
ip2long:205739409
long2ip:12.67.85.145
</pre>
###Golang中的net/url包
url.QueryEscape(s)将s进行转码使之可以安全的在URL查询中使用
<pre>
package main

import (
	"fmt"
	"net/url"
)

func main() {
	d := url.QueryEscape("ghty7789<f>>!攻关计划988*gy")
	fmt.Println(d)
}
output==>
ghty7789%3Cf%3E%3E%21%E6%94%BB%E5%85%B3%E8%AE%A1%E5%88%92988%2Agy
</pre>
###Golang实现微信公众平台
<pre>
package main

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Request struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

type Response struct {
	ToUserName   string `xml:"xml>ToUserName"`
	FromUserName string `xml:"xml>FromUserName"`
	CreateTime   string `xml:"xml>CreateTime"`
	MsgType      string `xml:"xml>MsgType"`
	Content      string `xml:"xml>Content"`
	MsgId        int    `xml:"xml>MsgId"`
}

func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func action(w http.ResponseWriter, r *http.Request) {
	postedMsg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	r.Body.Close()
	v := Request{}
	xml.Unmarshal(postedMsg, &v)
	if v.MsgType == "text" {
		v := Request{v.ToUserName, v.FromUserName, v.CreateTime, v.MsgType, v.Content, v.MsgId}
		output, err := xml.MarshalIndent(v, " ", " ")
		if err != nil {
			fmt.Printf("error:%v\n", err)
		}
		fmt.Fprintf(w, string(output))
	} else if v.MsgType == "event" {
		Content := `"欢迎关注
                                我的微信"`
		v := Request{v.ToUserName, v.FromUserName, v.CreateTime, v.MsgType, Content, v.MsgId}
		output, err := xml.MarshalIndent(v, " ", " ")
		if err != nil {
			fmt.Printf("error:%v\n", err)
		}
		fmt.Fprintf(w, string(output))
	}
}

func checkSignature(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var token string = "你的token"
	var signature string = strings.Join(r.Form["signature"], "")
	var timestamp string = strings.Join(r.Form["timestamp"], "")
	var nonce string = strings.Join(r.Form["nonce"], "")
	var echostr string = strings.Join(r.Form["echostr"], "")
	tmps := []string{token, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	tmp := str2sha1(tmpStr)
	if tmp == signature {
		fmt.Fprintf(w, echostr)
	}
}

func main() {
	http.HandleFunc("/check", checkSignature)
	http.HandleFunc("/", action)
	http.ListenAndServe(":8080", nil)
}
</pre>
###Linux下go程序守护进程|保证始终在后台运行
以beego为例：
<pre>
bee run
^C
nohup ./项目名称 &    //将该程序放到后台运行，不是可执行文件
</pre>
打开redis
<pre>
redis-cli
</pre>
linux下执行文件 
<pre>
/usr/test.sh
#! /bin/bash
echo "jason was here"


cd /usr/
. test.sh   //如果没有权限可以看下面

chmod +X ./test.sh   //使该脚本具有执行权限
. test.sh
</pre>
注意：在shell脚本中，声明变量在等号两边不能有空格，比如 myname = "j" 是错误的，必须是 myname="j"，这样才对。在shell脚本中自定义函数的调用方法是不用后面的(),如下:
<pre>
4fun(){
	echo "rrrrrr"
}
4fun        //调用函数不能是4fun()这种形式

. test.sh
输出：
rrrrrr

nl 文件名    //计算行数

source命令通常用于重新执行刚修改的初始化文件，使之立即生效，而不必注销并重新登录
source /etc/profile
</pre>
linux常用统计命令
<pre>
1）统计80端口连接数

netstat -nat|grep -i "80"|wc -l

2）统计httpd协议连接数

ps -ef|grep httpd|wc -l

3）、统计已连接上的，状态为 established

netstat -na|grep ESTABLISHED|wc -l

4）、查出哪个IP地址连接最多，将其封了。

netstat -na|grep ESTABLISHED|awk {print $5}|awk -F： {print $1}|sort|uniq -c|sort -r +0n

netstat -na|grep SYN|awk {print $5}|awk -F： {print $1}|sort|uniq -c|sort -r +0n

1、查看apache当前并发访问数：

netstat -an | grep ESTABLISHED | wc -l

对比httpd.conf中MaxClients的数字差距多少。

2、查看有多少个进程数：

ps aux|grep httpd|wc -l

3、可以使用如下参数查看数据

server-status？auto

ps -ef|grep httpd|wc -l

1388

统计httpd进程数，连个请求会启动一个进程，使用于Apache服务器。

表示Apache能够处理1388个并发请求，这个值Apache可根据负载情况自动调整。

netstat -nat|grep -i "80"|wc -l

4341

netstat -an会打印系统当前网络链接状态，而grep -i "80"是用来提取与80端口有关的连接的，wc -l进行连接数统计。

最终返回的数字就是当前所有80端口的请求总数。

netstat -na|grep ESTABLISHED|wc -l

376

netstat -an会打印系统当前网络链接状态，而grep ESTABLISHED 提取出已建立连接的信息。 然后wc -l统计。

最终返回的数字就是当前所有80端口的已建立连接的总数。

netstat -nat||grep ESTABLISHED|wc - 可查看所有建立连接的详细记录 
</pre>
###Golang实现DDNS客户端
对于动态ip的管理比较不方便，于是乎想到了使用DDNS来解决这个问题。 

主要作用：一，宽带营运商大多只提供动态的IP地址，DDNS可以捕获用户每次变化的IP地址，然后将其与域名相对应，这样其他上网用户就可以通过域名来与用户交流了；二，DDNS可以帮你在自己的公司或家里构建虚拟主机！
<pre>
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const GETIPURL string = "http://ddns.oray.com/checkip"
const UPDATEIP string = "http://%s:%s@ddns.oray.com/ph/update?hostname=%s&myip=%s"

var logger *log.Logger = nil

func getMyIp(url string) (string, int) {
	reqest, err := http.Get(url)
	if err == nil {
		defer reqest.Body.Close()
		b, _ := ioutil.ReadAll(reqest.Body)
		reg := regexp.MustCompile(`\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}`)
		IpTemp := reg.FindString(string(b))
		return IpTemp, 0
	} else {
		return "", 1
	}
}

func updateDDNS(IpAdd string) int {
	url := fmt.Sprintf(UPDATEIP, os.Args[1], os.Args[2], os.Args[3], IpAdd)
	if logger != nil {
		logger.Println(url, os.Getpid())
	}
	reqest, err := http.Get(url)
	if err != nil {
		return 1
	} else {
		defer reqest.Body.Close()
		b, _ := ioutil.ReadAll(reqest.Body)
		spiltB := strings.Split(string(b), " ")
		fmt.Print(spiltB[0])
		if spiltB[0] == "good" {
			return 0
		} else if spiltB[0] == "nochg" {
			return 0
		} else {
			return 1
		}
	}
}

func main() {

	if 4 != len(os.Args) {
		fmt.Printf("the input param: user password url\n")
		return
	}

	logfile, logerr := os.OpenFile("/var/log/ddnsupdate", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	if nil != logerr {
		fmt.Printf("%s\n", logerr.Error())
		return
	}
	defer logfile.Close()

	if os.Getppid() != 1 {
		filePath, _ := filepath.Abs(os.Args[0])
		fmt.Printf("filePath=%s\n", filePath)
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Start()
		return
	}
	logger = log.New(logfile, "", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("Start ddnsupdate", os.Getpid())

	for {
		newIP, errorflg := getMyIp(GETIPURL)
		if errorflg == 0 {
			if 0 == updateDDNS(newIP) {
				logger.Println("updateDDNS OK", os.Getpid())
				break
			} else {
				logger.Println("updateDDNS error", os.Getpid())
			}
		} else {
			logger.Println("getMyIp error", os.Getpid())
		}

		time.Sleep(60 * 60 * time.Second)
	}
	logger.Println("exit ddnsupdate", os.Getpid())
}
</pre>
###Golang实现上传下载|突破百度云4G上传限制
httpserver.go
<pre>
package main
 
import (
        "fmt"
        "html/template"
        "io"
        "net/http"
        "os"
        "path/filepath"
        "regexp"
//      "strconv"
        "time"
)
 
var mux map[string]func(http.ResponseWriter, *http.Request)
 
type Myhandler struct{}
type home struct {
        Title string
}
 
const (
        Template_Dir = "./view/"
        Upload_Dir   = "./upload/"
)
 
func main() {
        server := http.Server{
                Addr:        ":9090",
                Handler:     &Myhandler{},
                ReadTimeout: 10 * time.Second,
        }
        mux = make(map[string]func(http.ResponseWriter, *http.Request))
        mux["/"] = index
        mux["/upload"] = upload
        mux["/file"] = StaticServer
        server.ListenAndServe()
}
 
func (*Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        if h, ok := mux[r.URL.String()]; ok {
                h(w, r)
                return
        }
        if ok, _ := regexp.MatchString("/css/", r.URL.String()); ok {
                http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))).ServeHTTP(w, r)
        } else {
                http.StripPrefix("/", http.FileServer(http.Dir("./upload/"))).ServeHTTP(w, r)
        }
 
}
 
func upload(w http.ResponseWriter, r *http.Request) {
 
        if r.Method == "GET" {
                t, _ := template.ParseFiles(Template_Dir + "file.html")
                t.Execute(w, "上传文件")
        } else {
                r.ParseMultipartForm(32 << 20)
                file, handler, err := r.FormFile("uploadfile")
                if err != nil {
                        fmt.Fprintf(w, "%v", "上传错误")
                        return
                }
                fileext := filepath.Ext(handler.Filename)
                if check(fileext) == false {
                        fmt.Fprintf(w, "%v", "不允许的上传类型")
                        return
                }
                //filename := strconv.FormatInt(time.Now().Unix(), 10) + fileext
                filename := handler.Filename
                f, _ := os.OpenFile(Upload_Dir+filename, os.O_CREATE|os.O_WRONLY, 0660)
                _, err = io.Copy(f, file)
                if err != nil {
                        fmt.Fprintf(w, "%v", "上传失败")
                        return
                }
                filedir, _ := filepath.Abs(Upload_Dir + filename)
                fmt.Fprintf(w, "%v", filename+"上传完成,服务器地址:"+filedir)
        }
}
 
func index(w http.ResponseWriter, r *http.Request) {
        title := home{Title: "首页"}
        t, _ := template.ParseFiles(Template_Dir + "index.html")
        t.Execute(w, title)
}
 
func StaticServer(w http.ResponseWriter, r *http.Request) {
        http.StripPrefix("/file", http.FileServer(http.Dir("./upload/"))).ServeHTTP(w, r)
}
 
func check(name string) bool {
        ext := []string{".exe", ".js", ".png"}
 
        for _, v := range ext {
                if v == name {
                        return false
                }
        }
        return true
}
</pre>
view/file.html 
<pre>
<html>
<head>
    <title>{{.}}</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
  <input type="file" name="uploadfile" />
  <input type="submit" value="upload" />
</form>
</body>
</html>
</pre>
view/index.htm 
<pre>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
        <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
        <title>{{.Title}}</title>
</head>
<body>
        <div>
                <a href="/upload">上传文件</a></p>
                <a href="/file">查看文件</a></p>
        </div>
</body>
</html>
</pre>
htm文件放view里面，要下载的文件放upload里面。 
###在mysql中查询多条不重复记录值并且获得其其他字段的值
<pre>
select id ,account,count(account) from trade_y where account > 18000000000 group by account;
==>
id      account      count(account)  
-----------------------------------
171		18202764528		 2
236		18268846032 	 19
196		18297858940		 6
800		18337393018		 87
1123	18368870825		 17
945		18668410558		 3
1139	18768112617		 18
126		18768176361		 7
-----------------------------------
select id ,account from trade_y where account > 18000000000 group by account;
==>
id      account
-------------------
171		18202764528
236		18268846032
196		18297858940
800		18337393018
1123	18368870825
945		18668410558
1139	18768112617
126		18768176361
-------------------
</pre>
###sql总数
<pre>
select sum(col) from (
select count(*) col  from table1 union all
 select count(*) col  from table2 union all
 select count(*) col  from table3 union  all
 select count(*) col  from table4)d
</pre>
###redis启动与停止脚本|start redis |stop redis
<pre>
//创建启动和停止服务脚本
------- start.sh ---------
#!/bin/bash
/usr/local/webserver/redis/redis-server /usr/local/webserver/redis/conf/redis.conf
------- stop.sh ---------

#!/bin/bash
kill `cat /usr/local/webserver/redis/run/redis.pid`
chmod a+x /usr/local/webserver/redis/start.sh /usr/local/webserver/redis/stop.sh
</pre>
验证redis服务是否成功:
<pre>
netstat -nlpt | grep 6379  //6379是redis默认端口号，同理可通过3306端口看mysql服务是否启动
</pre>
查找redis配置文件|某一文件
<pre>
//第一种:
locate redis.conf
//第二种：假设你忘记了redis.conf这个文件在系统的哪个目录下，甚至在系统的某个地方也不知道，则这是可以使用如下命令 
find / -name redis.conf   //在整个系统中查找redis.conf在中间的位置
//第三种：例如我们知道某个文件包含有redis这个字符串，那么要找到系统中所有包含有这特定字符串的文件是可以实现的，输入：
find / -name '*redis*'    //在整个系统中查找redis字符串在中间的位置

find / -amin -10  //查找在系统中最后10分钟访问的文件
find / -atime -2  //查找在系统中最后48小时访问的文件
find / -empty     //查找在系统中为空的文件或者文件夹
find / -group cat //查找在系统中属于groupcat的文件
find / -mmin -5   //查找在系统中最后5分钟里修改过的文件
find / -mtime -1  //查找在系统中最后24小时里修改过的文件
find / -nouser    //查找在系统中属于作废用户的文件
find / -user fred //查找在系统中属于FRED这个用户的文件

//在某一文件中查找某一特定字符串
find . -name redis.conf | xargs grep "requirepass" //在当前目录中查找redis.conf文件并且在该文件中查找字符串"requirepass"的位置

linux 实用命令运维
//电脑系统CPU核心
cat /proc/cpuinfo |grep -c processor
//TCP连接状态
netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}' 
/*注解：awk（AWK）工作流程是这样的：先执行BEGING，然后读取文件，读入有/n换行符分割的一条记录，然后将记录按指定的域分隔符划分域，填充域，$0则表示所有域,$1表示第一个域,$n表示第n个域,随后开始执行模式所对应的动作action。接着开始读入第二条记录······直到所有的记录都读完，最后执行END操作。*/
//系统运行内存大小
free -m |grep "Mem" | awk '{print $2}'
//按cpu利用率从大到小排列所有进程id等信息
ps -e -o "%C : %p : %z : %a"|sort -nr
//按内存从大到小排列所有进程id等信息
ps -e -o "%C : %p : %z : %a"|sort -k5 -nr
//查看连接某服务端口(比如这里的80端口)最多的的IP地址
netstat -an -t | grep ":80" | grep ESTABLISHED | awk '{printf "%s %s\n",$5,$6}' | sort
//统计服务器下面某种类型的所有文件（比如这里的jpg格式图片）的大小总和
find / -name *.jpg -exec wc -c {} \;|awk '{print $1}'|awk '{a+=$1}END{print a}'
//cpu负载
cat /proc/loadavg
//观察si和so值是否较大
vmstat 1 5
//磁盘空间大小
 df -h
--->比如:
Filesystem               Size  Used Avail Use% Mounted on
/dev/mapper/centos-root   48G  6.3G   42G  14% /
devtmpfs                 1.9G     0  1.9G   0% /dev
tmpfs                    1.9G     0  1.9G   0% /dev/shm
tmpfs                    1.9G  153M  1.8G   9% /run
tmpfs                    1.9G     0  1.9G   0% /sys/fs/cgroup
/dev/sda1                497M   96M  401M  20% /boot
//找出占用空间最多的文件或目录
du -cks * | sort -rn | head -n 10
//磁盘I/O负载
iostat -x 1 2
//网络负载
sar -n DEV
//检查是否有网络错误
netstat -i
//统计进程总数
ps aux | wc -l
//动态观察是否有异常进程出现
top -id 1
//统计系统在线用户人数
who | wc -l
//统计所有打开的文件数目
 lsof | wc -l
收集日志   # logwatch –print   配置/etc/log.d/logwatch.conf，将 Mailto 设置为自己的email 地址，启动mail服务(sendmail或者postfix)，这样就可以每天收到日志报告了。
//linux alias 命令 ，为某一操作设置别名(暂时性有效，长期有效则必须写到 /etc/bashrc 文件里)
alias cdd="cd /home/go/src"     //这边的等号两边不要有空格
</pre>
free命令可以显示Linux系统中空闲的、已用的物理内存及swap内存,及被内核使用的buffer。
<pre>
free -t 		//以总和的形式显示内存的使用信息
free -s 2  //每两秒显示一次
</pre>
在linux系统中，buffers和cached都是缓存，两者有什么区别呢？

为了提高磁盘存取效率, Linux做了一些精心的设计, 除了对dentry进行缓存(用于VFS,加速文件路径名到inode的转换), 还采取了两种主要Cache方式：Buffer Cache和Page Cache。前者针对磁盘块的读写，后者针对文件inode的读写。这些Cache有效缩短了 I/O系统调用(比如read,write,getdents)的时间。

磁盘的操作有逻辑级（文件系统）和物理级（磁盘块），这两种Cache就是分别缓存逻辑和物理级数据的。

Page cache实际上是针对文件系统的，是文件的缓存，在文件层面上的数据会缓存到page cache。文件的逻辑层需要映射到实际的物理磁盘，这种映射关系由文件系统来完成。当page cache的数据需要刷新时，page cache中的数据交给buffer cache，因为Buffer Cache就是缓存磁盘块的。但是这种处理在2.6版本的内核之后就变的很简单了，没有真正意义上的cache操作。

Buffer cache是针对磁盘块的缓存，也就是在没有文件系统的情况下，直接对磁盘进行操作的数据会缓存到buffer cache中，例如，文件系统的元数据都会缓存到buffer cache中。

简单说来，page cache用来缓存文件数据，buffer cache用来缓存磁盘数据。在有文件系统的情况下，对文件操作，那么数据会缓存到page cache，如果直接采用dd等工具对磁盘进行读写，那么数据会缓存到buffer cache。

所以我们看linux,只要不用swap的交换空间,就不用担心自己的内存太少.如果常常swap用很多,可能你就要考虑加物理内存了.这也是linux看内存是否够用的标准.

如果是应用服务器的话，一般只看第二行，+buffers/cache,即对应用程序来说free的内存太少了，也是该考虑优化程序或加内存了。
###简单sql语句，用sql语句修改表结构
<pre>
原表修改：
//增加字段
alter table `tablename` add `tid` int(11) NOT NULL COMMENT '借款类型ID'; 
//删除已有字段
alter table `tablename` drop `tid`;
//修改某一存在字段的内容：（下面以修改变长文本型字段的大小为例）
alter table `tablename` alter `tid` int(1000)
//新建表：
CREATE TABLE `tablename` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL COMMENT '备注内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
</pre>
浅析Golang中的指针:

- &符号的意思是对变量取地址，如：变量a的地址是&a;
- *符号的意思是对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值了。
###Golang错误处理包
<pre>
package err

import "fmt"

var Pkg = "packageName"

type Err struct {
	Pkg string
	Info string
	Prev error
}

func (e *Err) Error() string {
	if e.Prev == nil {
		return fmt.Sprintf("%s: %s", e.Pkg, e.Info)
	}
	return fmt.Sprintf("%s: %s\n%v", e.Pkg, e.Info, e.Prev)
}

func me(err error, format string, args ...interface{}) *Err {
	if len(args) > 0 {
		return &Err{
			Pkg: Pkg,
			Info: fmt.Sprintf(format, args...),
			Prev: err,
		}
	}
	return &Err{
		Pkg: Pkg,
		Info: format,
		Prev: err,
	}
}

func ce(err error, format string, args ...interface{}) {
	if err != nil {
		panic(me(err, format, args...))
	}
}

func ct(err *error) {
	if p := recover(); p != nil {
		if e, ok := p.(error); ok {
			*err = e
		} else {
			panic(p)
		}
	}
}

func oe(e error) error {
	var ret error = e
	for err, ok := ret.(*Err); ok && err.Prev != nil; err, ok = ret.(*Err) {
		ret = err.Prev
	}
	return ret
}
</pre>
###Golang之Float32ToByte、ByteToFloat32、Float64ToByte、ByteToFloat64
<pre>
import (  
    "encoding/binary"  
    "math"  
)  
  
func Float32ToByte(float float32) []byte {  
    bits := math.Float32bits(float)  
    bytes := make([]byte, 4)  
    binary.LittleEndian.PutUint32(bytes, bits)  
  
    return bytes  
}  
  
func ByteToFloat32(bytes []byte) float32 {  
    bits := binary.LittleEndian.Uint32(bytes)  
  
    return math.Float32frombits(bits)  
}  
  
func Float64ToByte(float float64) []byte {  
    bits := math.Float64bits(float)  
    bytes := make([]byte, 8)  
    binary.LittleEndian.PutUint64(bytes, bits)  
  
    return bytes  
}  
  
func ByteToFloat64(bytes []byte) float64 {  
    bits := binary.LittleEndian.Uint64(bytes)  
  
    return math.Float64frombits(bits)  
}  
</pre>
##工作中遇到的问题|记录
<pre>
package main

import (
	"fmt"
)

func main(){
	d := 4.786667
	dd := d * 100
	fmt.Println(dd)
}
output==>
//你猜等于多少?嘿嘿，不是478.6667，而是
478.66669999999993
</pre>
这是golang float64失真问题。
解决方法：使用下面的方法：
<pre>
package main

import (
	"fmt"
	"strconv"
)

func coutfloat(f float64, l int) float64 {
	str1 := fmt.Sprintf("%."+strconv.Itoa(l)+"f", f)

	fre, _ := strconv.ParseFloat(str1, 64)
	return fre
}
</pre>
###Golang 长连接
<pre>
// Package tcpkeepalive implements additional TCP keepalive control beyond what
// is currently offered by the net pkg.
//
// Only Linux >= 2.4, DragonFly, FreeBSD, NetBSD and OS X >= 10.8 are supported
// at this point, but patches for additional platforms are welcome.
//
// See also: http://felixge.de/2014/08/26/tcp-keepalive-with-golang.html
package tcpkeepalive

import (
	"fmt"
	"net"
	"os"
	"syscall"

	"time"
)

// EnableKeepAlive enables TCP keepalive for the given conn, which must be a
// *tcp.TCPConn. The returned Conn allows overwriting the default keepalive
// parameters used by the operating system.
func EnableKeepAlive(conn net.Conn) (*Conn, error) {
	tcp, ok := conn.(*net.TCPConn)
	if !ok {
		return nil, fmt.Errorf("Bad conn type: %T", conn)
	}
	if err := tcp.SetKeepAlive(true); err != nil {
		return nil, err
	}
	file, err := tcp.File()
	if err != nil {
		return nil, err
	}
	fd := int(file.Fd())
	return &Conn{TCPConn: tcp, fd: fd}, nil
}

// Conn adds additional TCP keepalive control to a *net.TCPConn.
type Conn struct {
	*net.TCPConn
	fd int
}

// SetKeepAliveIdle sets the time (in seconds) the connection needs to remain
// idle before TCP starts sending keepalive probes.
func (c *Conn) SetKeepAliveIdle(d time.Duration) error {
	return setIdle(c.fd, secs(d))
}

// SetKeepAliveCount sets the maximum number of keepalive probes TCP should
// send before dropping the connection.
func (c *Conn) SetKeepAliveCount(n int) error {
	return setCount(c.fd, n)
}

// SetKeepAliveInterval sets the time (in seconds) between individual keepalive
// probes.
func (c *Conn) SetKeepAliveInterval(d time.Duration) error {
	return setInterval(c.fd, secs(d))
}

func secs(d time.Duration) int {
	d += (time.Second - time.Nanosecond)
	return int(d.Seconds())
}

// Enable TCP keepalive in non-blocking mode with given settings for
// the connection, which must be a *tcp.TCPConn.
func SetKeepAlive(c net.Conn, idleTime time.Duration, count int, interval time.Duration) (err error) {

	conn, ok := c.(*net.TCPConn)
	if !ok {
		return fmt.Errorf("Bad connection type: %T", c)
	}

	if err := conn.SetKeepAlive(true); err != nil {
		return err
	}

	var f *os.File

	if f, err = conn.File(); err != nil {
		return err
	}
	defer f.Close()

	fd := int(f.Fd())

	if err = setIdle(fd, secs(idleTime)); err != nil {
		return err
	}

	if err = setCount(fd, count); err != nil {
		return err
	}

	if err = setInterval(fd, secs(interval)); err != nil {
		return err
	}

	if err = setNonblock(fd); err != nil {
		return err
	}

	return nil
}

func setNonblock(fd int) error {
	return os.NewSyscallError("setsockopt", syscall.SetNonblock(fd, true))

}
</pre>
###Golang中使用安全证书的tls请求实现
因访问微信退款接口必须使用微信提供的安全证书与CA证书,所以在网上看到一位前辈的实现过程，特此搬运。
<pre>
import (
        "bytes"
        "crypto/tls"
        "crypto/x509"
        "io/ioutil"
        "net/http"
)
 
wechatCertPath = "/path/to/wechat/cert.pem"
wechatKeyPath = "/path/to/wechat/key.pem"
wechatCAPath = "/path/to/wechat/ca.pem"
wechatRefundURL = "https://wechat/refund/url"
 
var _tlsConfig *tls.Config
 
func getTLSConfig() (*tls.Config, error) {
        if _tlsConfig != nil {
                return _tlsConfig, nil
        }
 
        // load cert
        cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
        if err != nil {
                glog.Errorln("load wechat keys fail", err)
                return nil, err
        }
 
        // load root ca
        caData, err := ioutil.ReadFile(wechatCAPath)
        if err != nil {
                glog.Errorln("read wechat ca fail", err)
                return nil, err
        }
        pool := x509.NewCertPool()
        pool.AppendCertsFromPEM(caData)
 
        _tlsConfig = &tls.Config{
                Certificates: []tls.Certificate{cert},
                RootCAs:      pool,
        }
        return _tlsConfig, nil
}
 
func SecurePost(url string, xmlContent []byte) (*http.Response, error) {
        tlsConfig, err := getTLSConfig()
        if err != nil {
                return nil, err
        }
 
        tr := &http.Transport{TLSClientConfig: tlsConfig}
        client := &http.Client{Transport: tr}
        return client.Post(
                wechatRefundURL,
                "text/xml",
                bytes.NewBuffer(xmlContent))
}
</pre>
###Beego发送json数据
<pre>
this.Data["json"] = map[string]interface{}{"state": 1, "message": "发送成功"}
this.ServeJSON()
</pre>
###Golang月份|month
<pre>
package main

import (
	"fmt"
	"time"
)

func main() {
	//Add方法和Sub方法是相反的，获取t0和t1的时间距离d是使用Sub，将t0加d获取t1就是使用Add方法
	var month string
	switch time.Now().Month() {
	case 1:
		month = "一月"
	case 2:
		month = "二月"
	case 3:
		month = "三月"
	case 4:
		month = "四月"
	case 5:
		month = "五月"
	case 6:
		month = "六月"
	case 7:
		month = "七月"
	case 8:
		month = "八月"
	case 9:
		month = "九月"
	case 10:
		month = "十月"
	case 11:
		month = "十一月"
	case 12:
		month = "十二月"
	default:
		month = "default"
	}
	fmt.Println(month)
}
</pre>
###Golang获取今天开始与结束时间戳
<pre>
package main

import (
	"fmt"
	"time"
)

func GetTimeStrSecond(timestr string) int64 {
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
	return end.Unix()
}

func main() {
	// 今天凌晨的时间
	starttoday := time.Now().Format("2006-01-02") + " 00:00:00"
	//今天结束的时间
	endtoday := time.Now().Format("2006-01-02") + " 23:59:59"
	st := GetTimeStrSecond(starttoday)
	et := GetTimeStrSecond(endtoday)
	fmt.Println(starttoday+"的时间戳为：", st)
	fmt.Println(endtoday+"的时间戳为：", et)
}
output==>
2016-07-08 00:00:00的时间戳为： 1467907200
2016-07-08 23:59:59的时间戳为： 1467993599
</pre>
###Golang 获取距离某一时间的剩余秒数
<pre>
package main

import (
	"time"
	"fmt"
)

///////////////获取活动剩余时间的秒数(新的，可以修改时间)//////////////
func GetAtivityTodayLastSecondByEndtime(timestr string) time.Duration {
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
	return time.Duration(end.Unix()-time.Now().Local().Unix()) * time.Second
}

func main() {
	resttime := GetAtivityTodayLastSecondByEndtime("2016-07-20 23:59:59") //剩余的秒数
	fmt.Println(resttime)
}
output==>
128h43m53s
</pre>
###Golang 简单数据结构实现
单链表
<pre>
package main

//链表实现
import (
	"fmt"
)

//定义错误常量
const (
	ERROR = -1000000001
)

//定义元素类型
type Element int64

//定义节点
type LinkNode struct {
	Data Element   //数据域
	Nest *LinkNode //指针域，指向下一个节点
}

//函数接口
type LinkNoder interface {
	Add(head *LinkNode, new *LinkNode)              //后面添加
	Delete(head *LinkNode, index int)               //删除指定index位置元素
	Insert(head *LinkNode, index int, data Element) //在指定index位置插入元素
	GetLength(head *LinkNode) int                   //获取长度
	Search(head *LinkNode, data Element)            //查询元素的位置
	GetData(head *LinkNode, index int) Element      //获取指定index位置的元素
}

//添加 头结点，数据
func Add(head *LinkNode, data Element) {
	point := head //临时指针
	for point.Nest != nil {
		point = point.Nest //移位
	}
	var node LinkNode  //新节点
	point.Nest = &node //赋值
	node.Data = data
	head.Data = Element(GetLength(head)) //打印全部的数据
	if GetLength(head) > 1 {
		Traverse(head)
	}

}

//删除 头结点 index 位置
func Delete(head *LinkNode, index int) Element {
	//判断index合法性
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return ERROR
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Nest //移位
		}
		point.Nest = point.Nest.Nest //赋值
		data := point.Nest.Data
		return data
	}
}

//插入 头结点 index位置 data元素
func Insert(head *LinkNode, index int, data Element) {
	//检验index合法性
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Nest //移位
		}
		var node LinkNode //新节点，赋值
		node.Data = data
		node.Nest = point.Nest
		point.Nest = &node
	}
}

//获取长度 头结点
func GetLength(head *LinkNode) int {
	point := head
	var length int
	for point.Nest != nil {
		length++
		point = point.Nest
	}
	return length
}

//搜索 头结点 data元素
func Search(head *LinkNode, data Element) {
	point := head
	index := 0
	for point.Nest != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			break
		} else {
			index++
			point = point.Nest
			if index > GetLength(head)-1 {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
}

//获取data 头结点 index位置
func GetData(head *LinkNode, index int) Element {
	point := head
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return ERROR
	} else {
		for i := 0; i < index; i++ {
			point = point.Nest
		}
		return point.Data
	}
}

//遍历 头结点
func Traverse(head *LinkNode) {
	point := head.Nest
	for point.Nest != nil {
		fmt.Println(point.Data)
		point = point.Nest
	}
	fmt.Println("Traverse OK!")
}

//主函数测试
func main() {
	var head LinkNode = LinkNode{Data: 0, Nest: nil}
	head.Data = 0
	var nodeArray []Element
	for i := 0; i < 10; i++ {
		nodeArray = append(nodeArray, Element(i+1+i*100))
		Add(&head, nodeArray[i])
	}
	Delete(&head, 3)
	Search(&head, 2032)
	Insert(&head, 23, 10010)
	Traverse(&head)
	fmt.Println("data is", GetData(&head, 6))
	fmt.Println("length:", GetLength(&head))
}
</pre>
线性表
<pre>
package main
 
//线性表的相关算法实现
import (
    "fmt"
)
 
//定义List数据结构
type Element int64
 
const (
    MAX_SIZE = 10 //最大Size
    ERROR    = -1 //出错值
    NULL     = 0  //空值
)
 
type Sqlist struct {
    data   []Element //数据数组
    length int       //当前长度
    size   int       //最大size
}
 
//定义List的Interface,这里写出所有用到的方法
type Lister interface {
    //初始化构造一个线性表
    InitList(sl *Sqlist) bool
    //清空一个线性表
    ClearList(sl *Sqlist) bool
    //判断线性表是否为空
    IsListEmpty(sl *Sqlist) bool
    //判断线性表是否为满
    IsListFull(sl *Sqlist) bool
    //获取线性表长度
    Listlength(sl Sqlist) int
    //根据index获取数据
    GetData(sl Sqlist, index int) Element
    //根据数据返回index
    GetIndex(sl Sqlist, data Element) int
    //在index位置插入元素
    InsertList(sl *Sqlist, index int, data Element) bool
    //删除index的元素
    DeleteList(sl *Sqlist, index int) Element
    //遍历List
    TraverseList(sl *Sqlist) bool
}
 
//新建一个线性表
func InitList() Sqlist {
    var sl Sqlist
    sl.data = make([]Element, MAX_SIZE)
    sl.length = 0
    sl.size = MAX_SIZE
    return sl
}
 
//清空一个线性表
func ClearList(sl *Sqlist) bool {
    for i := 0; i < sl.size; i++ {
        sl.data[i] = NULL //全部都置空
    }
    return true
}
 
//判断线性表是否为空
func IsListEmpty(sl *Sqlist) bool {
    return sl.length == 0
}
 
//判断线性表是否为满
func IsListFull(sl *Sqlist) bool {
    return sl.length == MAX_SIZE
}
 
//获取线性表长度
func Listlength(sl Sqlist) int {
    return sl.length
}
 
//根据index获取数据
func GetData(sl *Sqlist, index int) Element {
    if index < 0 && index > MAX_SIZE {
        fmt.Println("please check index")
        return ERROR
    } else {
        return sl.data[index]
    }
 
}
 
//根据数据返回index
func GetIndex(sl *Sqlist, data Element) int {
    var index int = 0
    for i := 0; i < sl.length; i++ {
        if data == sl.data[i] {
            index = i
            break
        }
    }
    return index
 
}
 
//在index位置插入元素
/**
第一要判断index是否合法，然后判断index的位置，注意移动过程，最后要把length加1
**/
func InsertList(sl *Sqlist, index int, data Element) bool {
    if index < 0 && index > sl.length {
        fmt.Println("please check index")
        return false
    }
    if !IsListFull(sl) {
        if index == 0 && sl.length != 0 {
            for i := sl.length; i < 1; i-- {
                sl.data[i] = sl.data[i-1] //千万注意
            }
        } else if index > 0 && index < sl.length {
            for i := sl.length; i < index; i-- {
                sl.data[i] = sl.data[i-1] //注意这一块儿
            }
        } else if index > sl.length {
            fmt.Println("beyoug length")
            return false
        }
        sl.data[index] = data
        sl.length++
        return true
    } else {
        fmt.Println("list is full")
        return false
    }
}
 
//删除index的元素
/**
第一要判断index是否合法，然后判断index的位置，注意移位的时候要想清楚怎么移动，最后要把length减1
**/
func DeleteList(sl *Sqlist, index int) Element {
    if index < 0 && index > sl.length {
        fmt.Println("please check index")
        return ERROR
    }
    if !IsListEmpty(sl) {
        var data Element = sl.data[index]
        if index == 0 {
            for i := 0; i < sl.length; i++ {
                sl.data[i] = sl.data[i+1] //注意这个
            }
        }
        if index > 0 && index < sl.length {
            for i := index; i < sl.length-1; i++ {
                sl.data[i] = sl.data[i+1] //要注意
            }
        }
        sl.data[sl.length-1] = NULL
        sl.length--
        return data
 
    } else {
        fmt.Println("list is empty!")
        return ERROR
    }
}
 
//遍历List
func TraverseList(sl *Sqlist) bool {
    for i := 0; i < sl.length; i++ {
        fmt.Println(sl.data[i])
    }
    return true
}
func main() {
    list := InitList()
    fmt.Println(list.length)
    for i := 0; i < MAX_SIZE; i++ {
        InsertList(&list, i, Element(i*100))
    }
    TraverseList(&list)
    DeleteList(&list, 4)
    fmt.Println(list.data)
    fmt.Println(GetData(&list, 0))
    fmt.Println(GetIndex(&list, 600))
    ClearList(&list)
}
</pre>
###将数据格式化为[]byte类型
<pre>
data := `{
	"Verifyrealname":"` + ver + `",
	"IdCard":"` + idcard + `"
}`
sodata := []byte(data)
</pre>
###将不同要求的内容放入不同的slice
<pre>
package main

import (
	"fmt"
)

func main() {
	var sli2 []int
	sli1 := []int{1, 2, 3, 4, 5}
	for i := len(sli1) - 1; i >= 0; i-- {
		if i >= 2 {
			sli2 = append(sli2, sli1[i])           // 不符合条件
			sli1 = append(sli1[:i], sli1[i+1:]...) // 符合条件
		}
	}
	fmt.Println("sli1:", sli1)
	fmt.Println("sli2:", sli2)

}
output==>
sli1: [1 2]
sli2: [5 4 3]
</pre>
###SQL小技巧
为避免发生where后没有查询条件而出现错误的时候，一般将"where"改写成"where 1=1"。

为了提高查询效率，一般方法是利用数据库的缓存功能，因此如果 需要使用到类似CURDATE()、NOW()、RAND()或TO_DAYS()等函数的时候必须把它们用变量代替，因为那些函数的返回是会不定的易变的。比如下例：
<pre>
 sql := `select * from table1 where starttime >= CURDATE()  limit 10`    //性能差，没有利用到数据库的查询缓存

 now := time.Now().Format("2016-01-02 15:04:05")
 sql := `select * from table1 where starttime >= '`+now+`' limit 10`     //使用变量来提高性能
</pre>
使用left join进行多表联合查询时候的性能优化建议

1. 多表关联，数据量大的表放到前面，数据量越小的表越放到where后面去，能提高一定速度；
2. 排序问题，非必要不排序，如果要排序要在数据量最小的情况下排序。

拆分大的 DELETE 或 INSERT 语句

delete 与 insert两个操作是会锁表的，表一锁住了，别的操作都进不来了。因此在访问量很大的网站上，对于这种需求需要当心，一般是拆分多次进行。
<pre>
while (1) {
    //每次只做1000条
	o :=orm.NewOrm()
	sql :=`delete from table1 limit 1000`
	_,err := o.Raw(sql)
	if err == nil{
		//没得删的时候退出
		break
	}
	//阶段性停止
	time.Sleep(50000)
}
</pre>

###数据的偏移量
把存储单元的实际地址与其所在段的段地址之间的距离称为段内偏移，也称为“有效地址或偏移量”。 亦： 存储单元的实际地址与其所在段的段地址之间的距离。本质其实就是“实际地址与其所在段的段地址之间的距离” 

更通俗一点讲，内存中存储数据的方式是：一个存储数据的“实际地址”=段首地址+偏移量， 

你也可以这样理解：就像我们现实中的“家庭地址”=“小区地址”+“门牌号” ，上面的“偏移量”就好比“门牌号”。
###Golang mysq日常操作CURD
<pre>
import (
	"github.com/astaxie/beego/orm"
)
o := orm.NewOrm()
//select QueryRaw
err := o.Raw(`select * from activity where is_use = ?`, is_useid).QueryRaw(&ff)

//select QueryRaws
res, err := o.Raw(`select * from activity where is_use = ?`, is_useid).QueryRaws(&ffs)

//update  Exec
res, err := o.Raw(`update activity set is_use = ? where uid=?`, is_uesid, uid).Exec()

//insert Exec
res, err := o.Raw(`insert into activity(is_use) values(?)`, is_useid).Exec()

// delete Exec
res, err := o.Raw(`delete from activity where is_use = ?`, is_useid).Exec()
</pre>
###通过身份证号判断年龄
<pre>
package main

import (
	"fmt"
	"strconv"
	"time"
)

//对字符串进行截取
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func main(){
	cardNo := "3406761999070656453674"
	idstr := Substr(cardNo, 6, 8)
	id, _ := strconv.Atoi(idstr)
	now, _ := strconv.Atoi(time.Now().Format("20060102"))
	bottom := now - 700000
	top := now - 200000
	if id < bottom || id > top {
		result := "本公司不向20周岁以下，70周岁以上的用户提供服务。敬请理解！"
		fmt.Println(result)
		return
	}else{
		fmt.Println("OK")
	}
}
</pre>

###git使用tips
本地分支误删了一个文件，想从remote获取下来(恢复本地)的命令是：
<pre>
git checkout 文件名
</pre>
###mysql的cancat与left|right 与 substring的使用案例
<pre>
SELECT concat('点击量是',click,";",left(author,3)) as newclick FROM `default_news` where click >400;

==>
点击量是433;匿名者
点击量是411;匿名者
点击量是572;匿名者
点击量是413;匿名者
点击量是568;匿名者
点击量是803;匿名者

select  SUBSTRING(summary,1,2)as new from default_news where id = 33;

==> 
何为


select sum(qi)  from (
select count(*) as qi from users_y   union                 //别名要与最外层的保持一致
select count(*) as qi from users_l   union				   //别名要与最外层的保持一致
select count(*) as qi from users_d   union                 //别名要与最外层的保持一致
select count(*) as qi from users_q                         //别名要与最外层的保持一致
) as onetable；

==>
99
</pre>
IPC ==> Inter-Process Communication,进程间通信
###Golang json 
<pre>
package main

import (
	"encoding/json"
	"fmt"
)

type Other struct {
	Host   string
	Port   int
	Fruits []string
}

func main() {

	jsonstr := `{"host":"localhost","port":4566,"fruits":["apple","fff","ere"]}`

	var config Other
	json.Unmarshal([]byte(jsonstr), &config)

	fmt.Println(config)
}

output==>
{localhost 4566 [apple fff ere]}
</pre>
<pre>
func main(){
	today := time.Now().Format("2006-01-02") + " 00:00:00"
	fmt.Println("获取今天凌晨时间戳：", GetTimeStrSecond(today))

	strtime := time.Now()
	fmt.Println("此刻时间戳：", strtime.Local().Unix())
}
</pre>
###Linux 命令行奇用
<pre>
2>&1 &                                            //解释：(1)bash中0，1，2三个数字分别代表STDIN_FILENO、STDOUT_FILENO、STDERR_FILENO，即标准输入（一般是键盘），标准输出（一般是显示屏，准确的说是用户终端控制台），标准错误（出错信息输出）; (2)输入输出可以重定向,有时候会看到如 ls >> 1.txt这类的写法，> 和 >> 的区别在于：> 用于新建而>>用于追加;（3）2>&1就是用来将标准错误2重定向到标准输出1中的。此处1前面的&就是为了让bash将1解释成标准输出而不是文件1。至于最后一个&，则是让bash在后台执行。

find / -name *.java > find.txt 2>&1 &             //解释：将find的命令得到的结果写入 find.txt 文件中并且此操作是在后台进行的，不会受到其他操作的干扰。  

ps xj                                             //解释：守护进程是系统长期运行的后台进程。列出进程的信息，TPGID一栏为-1就是守护进程。
</pre>
###不同数据源访问时间差异
一次内存访问、SSD 硬盘访问和 SATA 硬盘随机访问的时间分别是 ：

几十纳秒，几十微秒，十几毫秒

内存访问速度属于纳秒级别；SSD性能显著优于传统硬盘（HDD）的其中一个原因是其超快的数据访问速度(发出请求和完成读写运行之间的延迟)。SSD的随机数据访问时间为0.1ms或更短，而主流2.5 HDD所用的时间约为10~12ms，甚至更长。就像您在下面对照表中所看到的，SSD的数据访问速度要比HDD快100倍，包括数据搜索时间和延迟。
###Linux远程传输命令scp
<pre>
//将本地 /usr/a/index.html 文件复制到 root@192.168.1.2:/usr/a/文件夹下 (上传)
scp /usr/a/index.html  root@192.168.1.2:/usr/a/

//若 SSH端口不是默认的22，比如，是端口1234 则加-P参数：
scp -P 1234  /usr/a/index.html  root@192.168.1.2:/usr/a/


//将远程 2服务器上的my.cnf文件下载到本地mysql文件夹下 （下载）
scp root@192.168.1.2:/etc/mysql/my.cnf  /etc/mysql

// -r 递归复制整个目录(一般用在需要复制[上传/下载]整个文件夹时使用)
scp -r root@192.168.1.2:/opt/soft/mongodb /opt/soft/
</pre>
###SVN版本控制之回滚
取消对代码的修改分为两种情况：
####第一种情况：改动没有被提交（commit）
这种情况下，使用svn revert就能取消之前的修改。
svn revert用法如下：
<pre>
# svn revert [-R] something
</pre>
其中something可以是（目录或文件的）相对路径也可以是绝对路径。
当something为单个文件时，直接svn revert something就行了；当something为目录时，需要加上参数-R(Recursive,递归)，否则只会将something这个目录的改动。
在这种情况下也可以使用svn update命令来取消对之前的修改，但不建议使用。因为svn update会去连接仓库服务器，耗费时间。
注意：svn revert本身有固有的危险，因为它的目的是放弃未提交的修改。一旦你选择了恢复，Subversion没有方法找回未提交的修改。
####第二种情况：改动已经被提交（commit）
这种情况下，用svn merge命令来进行回滚。 

回滚的操作过程如下： 

1. 保证我们拿到的是最新代码： 
<pre>
     svn update 
     假设最新版本号是28。
</pre> 
2. 然后找出要回滚的确切版本号：
<pre> 
     svn log [something]
     假设根据svn log日志查出要回滚的版本号是25，此处的something可以是文件、目录或整个项目
     如果想要更详细的了解情况，可以使用svn diff -r 28:25 [something]
</pre>
3. 回滚到版本号25：
<pre>
     svn merge -r 28:25 something
</pre>

为了保险起见，再次确认回滚的结果：
<pre>
     svn diff [something]
     发现正确无误，提交。
</pre>
4. 提交回滚：
<pre>
     svn commit -m ”Revert revision from r28 to r25,because of …” 
     提交后版本变成了29。
</pre>
将以上操作总结为三条如下：

1. svn update，svn log，找到最新版本（latest revision）
2. 找到自己想要回滚的版本号（rollbak revision）
3. 用svn merge来回滚： svn merge -r : something
