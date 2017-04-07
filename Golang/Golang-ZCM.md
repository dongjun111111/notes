#喵
##基础知识
<font color=red>在Golang中不能用string(int)==>string，只能用strconv.Itoa(int) ==>string,同理,将string转成int只能用strconv.Atoi；string强制转换只能用于将切片转成string</font>。

字符串

因为Golan中的字符串是不可变的，所以不会像其他语言那样很容易就修改字符串的内容。但是还是有至少下面两种方式来实现字符串内容的修改。

第一种：转成[]byte类型
<pre>
package main

func main() {
	//Go中字符串是不可变的,所以 var s string = "hello" s[0] = 'c' println(s)报错
	var s string = "hello"
	c := []byte(s)  //将字符串 s 转换成 []byte 类型
	c[0] = 'c'
	s = string(c)  //再转换成 string 类型
	println(s)
}
</pre>
第二种：切片操作
<pre>
package main

import "fmt"

func main() {
	s := "hello"
	s = "c" + s[1:]   //切片操作
	fmt.Println(s)
}
output==>
cello
</pre>
数组  -> 值类型
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
map无序的，可能每次打印的 map 不是相同顺序的；通过 delete 删除 map 元素：
<pre>
package main 

import "fmt"

func main(){
	var a map[int]int = make(map[int]int)
	a[4] = 3
	a[2] = 1
	a[5] = 4
	a[6] = 5
	delete(a,6)      //删除map
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
	fmt.Println("x = ", x)

	x1 := add1(x)

	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x)
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
	fmt.Println("x = ", x)

	x1 := add1(&x)

	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x)
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

//init函数自动调用
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
	Human  //匿名字段，默认Student包含了Human的所有字段
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
匿名的其他类型[]string int等
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
	int    //内置类型作为匿名字段
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
	fmt.Println(x, z, "c iota:", c, "a4 iota:", a4) // c等于2 a4等于3
	fmt.Println("x2 iota:", x2)                     // x2等于2
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
[[1 2 5] [4 6 4]] +++++++ 4
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
	fmt.Println("after insert: ", ee)

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
    "index":    index,  //用于输出数组元素
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
index函数用于输出数组元素：
<pre>
{{index x 1 2 3}}
返回index后面的第一个参数的某个索引对应的元素值，其余的参数为索引值
表示：x[1][2][3]
x必须是一个map、slice或数组
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
- 电子
- 件地址
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
身份证正则：
<pre>
^(\d{15}|^\d{18}|^\d{17}(\d|X|x))$
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
0. beego的basecontroller

func (c *BaseController) Prepare() {
	var abandon bool = false
	if c.Ctx.Input.Method() != "GET" && c.Ctx.Input.Method() != "HEAD" && !c.Ctx.Input.IsUpload() {
		if c.Ctx.Input.Method() == "POST" || c.Ctx.Input.Method() == "PUT" {
			var Res POSTRESULT
			json.Unmarshal(c.Ctx.Input.RequestBody, &Res)
			if strings.TrimSpace(Res.MobileVersion) == ""  {
				abandon = true
			}
		}
	} else {
		if c.Ctx.Input.Method() == "GET" {
			mobileversion := c.GetString("mobileversion")
				if strings.TrimSpace(mobileversion) == ""  {
					abandon = true
				}
		}
	}
	if abandon {
		data := map[string]interface{}{"ret": 403, "err": "亲,请升级APP版本~"}
		c.Data["json"] = data
		c.ServeJSON()
		c.StopRun()
	}
}

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

beego main函数添加共有方法
<pre>
import (
	"github.com/astaxie/beego/plugins/cors"
)

beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"GET", "POST"},
	AllowHeaders:     []string{"Origin"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
}))
</pre>

#####出现这样的错误
<RawSeter.QueryRows> all args must be use ptr slice

原因很可能是：
 QueryRow  ->  QueryRows
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
该方法返回值中会包含两个双引号，需注意！！！
<pre>
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ToString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)    //return string(data)[1:len(data)-1]
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
###Golang将float64转换成string
<pre>
package main

import (
	"fmt"
	"strconv"
)

var s = 12

func init() {
	if true {
		s = 34
	}
}

func main() {
	fmt.Println(s)
	fmt.Println(strconv.FormatFloat(56.78888, 'f', 3, 64))
}
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
	to := strings.Split(user, ";") //多个收件人用;号隔开
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


//启动redis 
redis-server ./redis.conf

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

####mysql 触发器 trigger 

INSERT into table1(id,tradeid) values('0','trigger');

DROP TRIGGER IF EXISTS trigger_on_tab1;
CREATE TRIGGER trigger_on_tab1
AFTER INSERT ON  table1
FOR EACH ROW
BEGIN
      INSERT into table2(id,tradeid) values('0',new.tradeid);
END;
-- 用完删除触发器
DROP TRIGGER trigger_on_tab1;

SHOW TRIGGERS;
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
	Next *LinkNode //指针域，指向下一个节点
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
	for point.Next != nil {
		point = point.Next //移位
	}
	var node LinkNode  //新节点
	point.Next = &node //赋值
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
			point = point.Next //移位
		}
		point.Next = point.Next.Next //赋值
		data := point.Next.Data
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
			point = point.Next //移位
		}
		var node LinkNode //新节点，赋值
		node.Data = data
		node.Next = point.Next
		point.Next = &node
	}
}

//获取长度 头结点
func GetLength(head *LinkNode) int {
	point := head
	var length int
	for point.Next != nil {
		length++
		point = point.Next
	}
	return length
}

//搜索 头结点 data元素
func Search(head *LinkNode, data Element) {
	point := head
	index := 0
	for point.Next != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			break
		} else {
			index++
			point = point.Next
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
			point = point.Next
		}
		return point.Data
	}
}

//遍历 头结点
func Traverse(head *LinkNode) {
	point := head.Next
	for point.Next != nil {
		fmt.Println(point.Data)
		point = point.Next
	}
	fmt.Println("Traverse OK!")
}

//主函数测试
func main() {
	var head LinkNode = LinkNode{Data: 0, Next: nil}
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

###golang将类型转换为string
<pre>
package main

import (
	"encoding/json"
	"fmt"
)

func ToString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func main() {
	var s int
	s = 4545
	fmt.Println(ToString(s))
}
</pre>
###Golang 将标准时间格式转换成时间戳格式
<pre>
func GetTimeStrSecond(timestr string) int64 {
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
	return end.Unix()
}

//毫秒
timestamp := strconv.FormatInt(time.Now().UnixNano()/1000, 10)
fmt.Println(timestamp)
</pre>
###COUNT(*) 与 COUNT(1)
最关键是在你统计个数的时候是否需要考虑到有空值的情况，count(主键)肯定没有空值，但是对于一个没有主键的表或者 count(任意字段)时，count(*)能取出含有空值的所有记录数，count(任意字段)不含空值。
###Golang之md5
<pre>
//md5加密
func mD5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
//数组去重 暂时支持string ,可以自己实现其他几种类型
func RemoveAndEmpty(a []string) (ret []string) {
	sort.Sort(sort.StringSlice(a))
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}
</pre>
###Golang 将数据库中的表信息写入excel文件中
<pre>
func (tagetPath string, fileName string, data interface{}) dbToExcel() {
	if data != nil {
		f, err := os.Create(tagetPath + "\\" + filename + ".xls")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

		w := csv.NewWriter(f)
		//		w.Write([]string{"编号", "姓名", "年龄"})
		//		w.Write([]string{"1", "张三", "23"})
		//		w.Write([]string{"2", "李四", "24"})
		//		w.Write([]string{"3", "王五", "25"})
		//		w.Write([]string{"4", "赵六", "26"})

		for _, v := range rcl {
			w.Write([]string{strconv.Itoa(v.Id), v.ConfigName, v.ConfigKey, v.Desc})

		}
		w.Flush()
	}
}
</pre>
###在post或者get请求中添加自定义参数
<pre>
//发送post请求
func SendPost(url string, body []byte) ([]byte, error) {

	// 为了处理，所有post请求都加入 3个参数
	ss := strings.Replace(string(body), "}", "", -1)
	aa := ss + `,"参数1":` + strconv.Itoa(utils.参数1) + `,"参数2":"` + utils.参数2 + `","参数3":"` + utils.参数3 + `'}`
	body = []byte(aa)
	//==========================
	requestBody, _ := utils.DesBase64Encrypt(body)
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
		responseBody, err := utils.解密(bodyByte)
		return responseBody, err
	} else {
		return nil, err
	}
}

//发送get请求
func SendGet(url string) ([]byte, error) {
	// 为了处理，所有post请求都加入  3个参数
	if strings.Contains(url, "?") {
		url = url + "&参数1=" + strconv.Itoa(utils.参数1) + "&参数2=" + utils.参数2 + "&参数3=" + utils.参数3
	} else {
		url = url + "?参数1=" + strconv.Itoa(utils.参数1) + "&参数2=" + utils.参数2 + "&参数3=" + utils.参数3
	}
	// =======================================
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.Status == "404 Not Found" {
		return nil, errors.New("服务器异常!")
	}
	if bodyByte, err := ioutil.ReadAll(resp.Body); err == nil {
		responseBody, err := utils.解密(bodyByte)
		return responseBody, err
	} else {
		return nil, err
	}
}
</pre>
###Golang将struct转换成map字典
<pre>
// mapStructToMap 将一个结构体所有字段(包括通过组合得来的字段)映射到一个map中
// data:存储字段数据的map
// value:结构体的反射值
func mapStructToMap(data map[interface{}]interface{}, value reflect.Value) {
	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			var fieldValue = value.Field(i)
			if fieldValue.CanInterface() {
				var fieldType = value.Type().Field(i)
				if fieldType.Anonymous {
					//匿名组合字段,进行递归解析
					mapStructToMap(data, fieldValue)
				} else {
					//非匿名字段
					var fieldName = fieldType.Tag.Get("to")
					if fieldName == "" {
						fieldName = fieldType.Name
					}
					data[fieldName] = fieldValue.Interface()
				}
			}
		}
	}
}
</pre>
###Golang官网被墙解决方案
* 修改hosts文件

找到hosts文件，Mac OS X/*nix在/etc/hosts，Windows在C:\WINDOWS\system32\drivers\etc\hosts

增加一行
<pre>
173.194.75.141 golang.org
</pre>
* 本地启动godoc服务
<pre>
godoc -http=:6060
</pre>
通过浏览器访问http://localhost:6060即可。
###Linux 根据Pid获取对应的文件
<pre>
lsof -p "PID" 
</pre>
###[]rune类型拼接
<pre>
package main 

import (
	"fmt"
)
func merge(a, b []rune) []rune {
	c := make([]rune, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}

func main(){
	a := []rune("adobe")
	b := []rune("pdf")
	fmt.Println("merge结果是：", merge(a, b))
}

output==>
merge结果是： [97 100 111 98 101 112 100 102]
string(merge)结果是： adobepdf
</pre>
###nginx报错解决方法之一
手动添加错误日志文件
<pre>
mkdir -p /disk2/logs/error.log
</pre>
###Golang float64数字很大会变成科学计数法的解决方法
在网络传输过程中，过大的float64位的数似乎会变成科学计数法的显示形式，如3.0336e 06,这时候用下面的方法会解决这个问题。
<pre>
res := fmt.Sprintf("%.2f", (originValue))
</pre>
###Golang string rune byte 的关系
在Go当中 string底层是用byte数组存的，并且是不可以改变的。

例如 s:="Go编程" fmt.Println(len(s)) 输出结果应该是8因为中文字符是用3个字节存的。

len(string(rune('编')))的结果是3

如果想要获得我们想要的情况的话，需要先转换为rune切片再使用内置的len函数

fmt.Println(len([]rune(s)))

结果就是4了。

所以用string存储unicode的话，如果有中文，按下标是访问不到的，因为你只能得到一个byte。 要想访问中文的话，还是要用rune切片，这样就能按下表访问。

最直观的区别就是:

- rune 能操作任何字符
- byte 不支持中文的操作
###time.Now().After()用法
<pre>
begin, _ := time.ParseInLocation("2006-01-02 15:04:05", "2015-10-13 00:00:00", time.Local)
if time.Now().After(begin) {
	fmt.Println("yes after")
} else {
	fmt.Println("no")
}
类似的还有是：time.Now().Before(end) 
</pre>
###Linux下Go与Java环境变量
<pre>
GOBIN=/usr/local/go/bin //安装目录,下面文件有 bee  go  godoc  gofmt
JRE_HOME=/home/jdk1.7.0//jre
PATH=/home/jdk1.7//bin:/home/jdk1.7.0//jre/bin:/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin:/root/bin:/root/bin:/usr/local/go/bin
PWD=/home/go/bin
GOARCH=amd64
JAVA_HOME=/home/jdk1.7/
GOROOT=/usr/local/go
GOOS=linux
CLASSPATH=.:/home/jdk1.7//lib:/home/jdk1.7.0//jre/lib:
GOPATH=/home/go
OLDPWD=/home/go
</pre>
###Linux常用命令
<pre>
rpm -qa 						 //查看所有安装包名称
netstat -lntp 				 //查看所有正在监听的端口 
uptime  					 //查看系统运行的时间，在线用户数
netstat -antp  				 //查看所有已经建立的连接
w  							 //查看活跃用户，可以看出现在在线的所有用户 ip 地址以及最后使用的命令
grep MemTotal /proc/meminfo  //查看内存总量
grep MemFree /proc/meminfo   //查看空闲（可使用的）内存量 
uname -a                     //查看内核、操作系统、CPU信息
</pre>
###时间格式显示小坑
在IOS前端无法识别"2016-09-09 12:12:12",改成"2016/09/09 12:12:12"即可。
###Golang时间比较大小 After Before 时间格式化(不同类型)
<pre>
var productStartTime time.Time
if product.StartTime != "" {
	durat := 20
	productStartTime, _ = time.Parse("2006-01-02 15:04:05", product.StartTime)
	product.StartTime = productStartTime.Add(time.Duration(durat) * time.Second).Format("2006-01-02 15:04:05")
}
productS, _ := time.Parse("2006-01-02 15:04:05", product.StartTime)
productNow, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
if productS.After(productNow) {
	fmt.Println("")
} else {
	fmt.Println("")
}
</pre>
####查看linux系统TCP/UDP的端口号:
netstat -tupln
###Xshell下上传与下载的命令
XShell上传与下载文件命令sz,rz
<pre>
sz  test.log   //下载当前目录下文件名为test.log到本地
sz ./* //将当前文件夹下的所有文件发送到本地
rz  上传文件到远程服务器
rz -y 覆盖已存在文件
</pre>
Linux复制cp
<pre>
cp 文件名 目标地址
cp test.log /tmp  //将当前目录下的test,log文件复制到 /tmp 目录下
</pre>
###Linux下更新golang bee beego环境
1. 备份以前的go环境
<pre>
cd /usr/local
bee version 
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
bee version 
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
go get github.com/beego/bee
go version 
bee version 
</pre>
###适用于linux/window下的端口查看命令
<pre>
netstat -ano   //列出所有端口情况
netstat -ano|grep 8080  //看8080端口

系统管理命令:

stat         显示指定文件的详细信息，比ls更详细
who          显示在线登陆用户
whoami       显示当前操作用户
hostname     显示主机名
uname        显示系统简要信息
     -a          显示系统完整信息
top          动态显示当前耗费资源最多进程信息
ps           显示瞬间进程状态 ps aux
     -ef         显示系统常驻进程
du           查看目录大小 du -h /home带有单位显示目录信息
df           查看磁盘大小 df -h 带有单位显示磁盘信息
</pre>
###Bytes包
写
<pre>
func main() {
    // 操作目标
    data := make([]byte, 6)
    buff := bytes.NewWriter(data)
    //往目标中写入1,2,3,目标变为[1,2,3,0,0,0]
    buff.Write([]byte{1, 2, 3})
    fmt.Println(buff.Bytes())//输出buff中写入的内容,应该为[1,2,3]
    //往目标中继续写入数据,变成[1,2,3,4,5,6]
    buff.Write([]byte{4, 5, 6})
    n, err := buff.Write([]byte{7})
    fmt.Println(n, err)//由于目标已经满了,继续写就会造成io.EOF错误啦,n为0
    //写游标重置
    buff.Reset()
    //又能愉快的往目标区写内容啦
    buff.Write([]byte{7, 8, 9})
    fmt.Println(buff.Bytes())//注意,重置的除了游标,buff内原来的内容也清空
    fmt.Println(data)//但是操作区并不会应为Reset而清空,应该输出[7,8,9,4,5,6]
}
</pre>
读
<pre>
func main() {
    //操作目标,这个一般可能会超级大
    data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
    buff := bytes.NewReader(data)
    //这个一般是用来做复用的缓冲区,一般小于data的长度,为了试验效果,假设为4
    x := [4]byte{}
    for {
        n, err := buff.Read(x[:])
        if err != nil&&err == io.EOF {
            break
        }
        fmt.Println("temp", x[:n])
    }
    fmt.Println("remian:", buff.Bytes())//都被读光了,现在应该是空的了
    buff.SeekToBegin()//重新载入一下
    fmt.Println("remian:", buff.Bytes())//可以看到又满了
    fmt.Println(buff.Seek(3,1))
    fmt.Println("remian:", buff.Bytes())//从当前位置,游标跳3,剩余[4,5,6,7,8]
    fmt.Println(buff.Seek(2,0))
    fmt.Println("remian:", buff.Bytes())//从起始位置,游标跳2个,剩余[3,4,5,6,7,8],可以看到,隐含一个重置过程
    fmt.Println(buff.Seek(2,1))
    fmt.Println("remian:", buff.Bytes())//从当前位置,游标跳2,剩余[5,6,7,8]
    fmt.Println(buff.Seek(-3,2))
    fmt.Println("remian:", buff.Bytes())//从末尾开始往前数,游标跳3,剩余[6,7,8]
}
</pre>
###Golang 端口监听/在线聊天
服务端
<pre>
package main

import (
	"fmt"
	"log"
	"net"
)

func startServer() {
	listener, err := net.Listen("tcp", "localhost:7777")
	checkError(err)
	fmt.Println("建立成功")
	for {
		conn, err := listener.Accept()
		checkError(err)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	nameInfo := make([]byte, 512)
	_, err := conn.Read(nameInfo)
	checkError(err)
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		flag := checkError(err)
		if flag == 0 {
			break
		}
		fmt.Println(string(buf)) //打印出来
	}
}

//检查错误
func checkError(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			fmt.Println("用户退出了")
			return 0
		}
		log.Fatal("an error!", err.Error())
		return -1
	}
	return 1
}
func main() {
	//开启服务
	startServer()
}
</pre>
客户端
<pre>
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func connectServer() {
	conn, err := net.Dial("tcp", "localhost:7777")
	checkError(err)
	fmt.Println("连接成功！\n")
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("你是谁？")
	name, _ := inputReader.ReadString('\n')
	trimName := strings.Trim(name, "\r\n")
	conn.Write([]byte(trimName + "接入了\n"))
	for {
		fmt.Println("我们来聊天吧！按quit退出")
		input, _ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\r\n")
		if trimInput == "quit" {
			fmt.Println("再见")
			conn.Write([]byte(trimName + "退出了"))
			return
		}
		_, err = conn.Write([]byte(trimName + " says " + trimInput))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("an error!", err.Error())
	}
}

func main() {
	connectServer()
}
</pre>
UDP-1:
<pre>
//服务端
package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8448,
	})
	if err != nil {
		fmt.Println("监听失败", err)
		return
	}
	defer socket.Close()
	for {
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}
		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n\n", data)
		// 发送数据
		senddata := []byte("hello client!")
		_, err = socket.WriteToUDP(senddata, remoteAddr)
		if err != nil {
			return
			fmt.Println("发送数据失败!", err)
		}
	}
}


//客户端
package main
import (
    "fmt"
    "net"
)
func main() {
    // 创建连接
    socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
        IP:   net.IPv4(192, 168, 1, 103),
        Port: 8080,
    })
    if err != nil {
        fmt.Println("连接失败!", err)
        return
    }
    defer socket.Close()
    // 发送数据
    senddata := []byte("hello server!")
    _, err = socket.Write(senddata)
    if err != nil {
        fmt.Println("发送数据失败!", err)
        return
    }
    // 接收数据
    data := make([]byte, 4096)
    read, remoteAddr, err := socket.ReadFromUDP(data)
    if err != nil {
        fmt.Println("读取数据失败!", err)
        return
    }
    fmt.Println(read, remoteAddr)
    fmt.Printf("%s\n", data)
}
</pre>
UDP-2: 在线聊天
<pre>
//服务端
package main  
  
import (  
    "fmt"  
    "net"  
    "os"  
    "strconv"  
    "time"  
)  
//用户信息  
type User struct {  
     userName string  
     userAddr *net.UDPAddr  
     userListenConn *net.UDPConn  
     chatToConn *net.UDPConn  
}  
  
//服务器监听端口  
const LISTENPORT = 1616  
//缓冲区  
const BUFFSIZE = 1024  
var buff = make([]byte, BUFFSIZE)  
//在线用户  
var onlineUser = make([]User, 0)  
//在线状态判断缓冲区  
var onlineCheckAddr = make([]*net.UDPAddr, 0)  
  
//错误处理  
func HandleError(err error) {  
    if err != nil {  
        fmt.Println(err.Error())  
        os.Exit(2)  
    }  
}  
//消息处理  
func HandleMessage(udpListener *net.UDPConn) {  
    n, addr, err := udpListener.ReadFromUDP(buff)  
    HandleError(err)  
  
    if n > 0 {  
        msg := AnalyzeMessage(buff, n)  
          
        switch msg[0] {  
            //连接信息  
            case "connect  ":  
                //获取昵称+端口  
                userName := msg[1]  
                userListenPort := msg[2]  
                //获取用户ip  
                ip := AnalyzeMessage([]byte(addr.String()), len(addr.String()))  
                //显示登录信息  
                fmt.Println(" 昵称:", userName, " 地址:", ip[0], " 用户监听端口:", userListenPort, " 登录成功！")  
                //创建对用户的连接，用于消息转发  
                userAddr, err := net.ResolveUDPAddr("udp4", ip[0] + ":" + userListenPort)  
                HandleError(err)  
                  
                userConn, err := net.DialUDP("udp4", nil, userAddr)  
                HandleError(err)  
                  
                //因为连接要持续使用，不能在这里关闭连接  
                //defer userConn.Close()  
                //添加到在线用户  
                onlineUser = append(onlineUser, User{userName, addr, userConn, nil})  
                  
            case "online   ":  
                //收到心跳包  
                onlineCheckAddr = append(onlineCheckAddr, addr)  
                  
            case "outline  ":  
                //退出消息，未实现  
            case "chat     ":  
                //会话请求  
                //寻找请求对象  
                index := -1  
                for i := 0; i < len(onlineUser); i++ {  
                    if onlineUser[i].userName == msg[1] {  
                        index = i  
                    }  
                }  
                //将所请求对象的连接添加到请求者中  
                if index != -1 {  
                    nowUser, _ := FindUser(addr)  
                    onlineUser[nowUser].chatToConn = onlineUser[index].userListenConn  
                }  
            case "get      ":  
                //向请求者返回在线用户信息  
                index, _ := FindUser(addr)  
                onlineUser[index].userListenConn.Write([]byte("当前共有" + strconv.Itoa(len(onlineUser)) + "位用户在线"))  
                for i, v := range onlineUser {  
                    onlineUser[index].userListenConn.Write([]byte("" + strconv.Itoa(i + 1) + ":" + v.userName))  
                }  
            default:  
                //消息转发  
                //获取当前用户  
                index, _ := FindUser(addr)  
                //获取时间  
                nowTime := time.Now()  
                nowHour := strconv.Itoa(nowTime.Hour())  
                nowMinute := strconv.Itoa(nowTime.Minute())  
                nowSecond := strconv.Itoa(nowTime.Second())  
                //请求会话对象是否存在  
                if onlineUser[index].chatToConn == nil {  
                    onlineUser[index].userListenConn.Write([]byte("对方不在线"))  
                } else {  
                    onlineUser[index].chatToConn.Write([]byte(onlineUser[index].userName + " " + nowHour + ":" + nowMinute + ":" + nowSecond + "\n" + msg[0]))  
                }  
                  
        }  
    }  
}  
//消息解析，[]byte -> []string  
func AnalyzeMessage(buff []byte, len int) ([]string) {  
    analMsg := make([]string, 0)  
    strNow := ""  
    for i := 0; i < len; i++ {  
        if string(buff[i:i + 1]) == ":" {  
            analMsg = append(analMsg, strNow)  
            strNow = ""  
        } else {  
            strNow += string(buff[i:i + 1])  
        }  
    }  
    analMsg = append(analMsg, strNow)  
    return analMsg  
}  
//寻找用户，返回（位置，是否存在）  
func FindUser(addr *net.UDPAddr) (int, bool) {  
    alreadyhave := false  
    index := -1  
    for i := 0; i < len(onlineUser); i++ {  
          
        if onlineUser[i].userAddr.String() == addr.String() {  
            alreadyhave = true  
            index = i  
            break  
        }  
    }  
    return index, alreadyhave  
}  
//处理用户在线信息（暂时仅作删除用户使用）  
func HandleOnlineMessage(addr *net.UDPAddr, state bool) {  
    index, alreadyhave := FindUser(addr)  
    if state == false {  
        if alreadyhave {  
            onlineUser = append(onlineUser[:index], onlineUser[index + 1:len(onlineUser)]...)   
        }  
    }  
}  
//在线判断，心跳包处理，每5s查看一次所有已在线用户状态  
func OnlineCheck() {  
    for {  
        onlineCheckAddr = make([]*net.UDPAddr, 0)  
        sleepTimer := time.NewTimer(time.Second * 5)  
        <- sleepTimer.C  
        for i := 0; i < len(onlineUser); i++ {  
            haved := false  
            FORIN:for j := 0; j < len(onlineCheckAddr); j++ {  
                if onlineUser[i].userAddr.String() == onlineCheckAddr[j].String() {  
                    haved = true  
                    break FORIN  
                }  
            }  
            if !haved {  
                fmt.Println(onlineUser[i].userAddr.String() + "退出！")  
                HandleOnlineMessage(onlineUser[i].userAddr, false)  
                i--  
            }  
  
        }  
    }  
}  
  
func main() {  
    //监听地址  
    udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:" + strconv.Itoa(LISTENPORT))  
    HandleError(err)  
    //监听连接  
    udpListener, err := net.ListenUDP("udp4", udpAddr)  
    HandleError(err)  
  
    defer udpListener.Close()  
  
    fmt.Println("开始监听：")  
  
    //在线状态判断  
    go OnlineCheck()  
  
    for {  
        //消息处理  
        HandleMessage(udpListener)  
    }  
  
}  

//客户端
package main  
  
import (  
    "fmt"  
    "os"  
    "strconv"  
    "net"  
    "bufio"  
    "math/rand"  
    "time"  
)  
  
//数据包头，标识数据内容  
var reflectString = map[string]string {  
    "连接":       "connect  :",  
    "在线":       "online   :",  
    "聊天":       "chat     :",  
    "在线用户":     "get      :",  
}  
  
//服务器端口  
const CLIENTPORT = 1616  
//数据缓冲区  
const BUFFSIZE = 1024  
var buff = make([]byte, BUFFSIZE)  
  
//错误消息处理  
func HandleError(err error) {  
    if err != nil {  
        fmt.Println(err.Error())  
        os.Exit(2)  
    }  
}  
//发送消息  
func SendMessage(udpConn *net.UDPConn) {  
    scaner := bufio.NewScanner(os.Stdin)  
  
    for scaner.Scan() {  
        if scaner.Text() == "exit" {  
            return  
        }  
        udpConn.Write([]byte(scaner.Text()))  
    }  
}  
//接收消息  
func HandleMessage(udpListener *net.UDPConn) {  
    for {  
        n, _, err := udpListener.ReadFromUDP(buff)  
        HandleError(err)  
  
        if n > 0 {  
            fmt.Println(string(buff[:n]))  
        }  
    }  
}  
/*  
func AnalyzeMessage(buff []byte, len int) ([]string) {  
    analMsg := make([]string, 0)  
    strNow := ""  
    for i := 0; i < len; i++ {  
        if string(buff[i:i + 1]) == ":" {  
            analMsg = append(analMsg, strNow)  
            strNow = ""  
        } else {  
            strNow += string(buff[i:i + 1])  
        }  
    }  
    analMsg = append(analMsg, strNow)  
    return analMsg  
}*/  
//发送心跳包  
func SendOnlineMessage(udpConn *net.UDPConn) {  
    for {  
        //每间隔1s向服务器发送一次在线信息  
        udpConn.Write([]byte(reflectString["在线"]))  
        sleepTimer := time.NewTimer(time.Second)  
        <- sleepTimer.C  
    }  
}  
  
func main() {  
    //判断命令行参数，参数应该为服务器ip  
    if len(os.Args) != 2 {  
        fmt.Println("程序命令行参数错误！")  
        os.Exit(2)  
    }  
    //获取ip  
    host := os.Args[1]  
  
    //udp地址  
    udpAddr, err := net.ResolveUDPAddr("udp4", host + ":" + strconv.Itoa(CLIENTPORT))  
    HandleError(err)  
  
    //udp连接  
    udpConn, err := net.DialUDP("udp4", nil, udpAddr)  
    HandleError(err)  
  
    //本地监听端口  
    newSeed := rand.NewSource(int64(time.Now().Second()))  
    newRand := rand.New(newSeed)  
    randPort := newRand.Intn(30000) + 10000  
  
    //本地监听udp地址  
    udpLocalAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:" + strconv.Itoa(randPort))  
    HandleError(err)  
  
    //本地监听udp连接  
    udpListener, err := net.ListenUDP("udp4", udpLocalAddr)  
    HandleError(err)  
  
    //fmt.Println("监听", randPort, "端口")  
  
    //用户昵称  
    userName := ""  
    fmt.Printf("请输入昵称：")  
    fmt.Scanf("%s", &userName)  
      
    //向服务器发送连接信息（昵称+本地监听端口）  
    udpConn.Write([]byte(reflectString["连接"] + userName + ":" + strconv.Itoa(randPort)))  
  
    //关闭端口  
    defer udpConn.Close()  
    defer udpListener.Close()  
  
    //发送心跳包  
    go SendOnlineMessage(udpConn)  
    //接收消息  
    go HandleMessage(udpListener)  
  
    command := ""  
      
    for {  
        //获取命令  
        fmt.Scanf("%s", &command)  
        switch command {  
            case "chat" :  
                people := ""  
                //fmt.Printf("请输入对方昵称：")  
                fmt.Scanf("%s", &people)  
                //向服务器发送聊天对象信息  
                udpConn.Write([]byte(reflectString["聊天"] + people))  
                //进入会话  
                SendMessage(udpConn)  
                //退出会话  
                fmt.Println("退出与" + people + "的会话")  
            case "get" :  
                //请求在线情况信息  
                udpConn.Write([]byte(reflectString["在线用户"]))  
        }  
    }  
}  
</pre>
Golang TCP长连接
<pre>
/***********************************************************************
1、Notice：
	http的消息处理，是另开goroutine调用的，所以函数中可阻塞
	tcp的消息处理，是在readRoutine中及时调用的，所以函数中不能有阻塞调用
	否则“该条连接”的读会被挂起，c++中的话，整个系统的处理线程都会阻塞掉
2、server端目前是一条连接两个goroutine(readRoutine/writeRoutine)
	假设5k玩家，就有1w个goroutine，太多了
3、msghandler可考虑设计成：不执行逻辑，仅将消息加入buf队列，由一个goroutine来处理
	不过那5k个readRoutine貌似省不了哇，感觉单独一个goroutine处理消息也不会有性能提升
	且增加了风险，若某条消息有阻塞调用，后面的就得等了
***********************************************************************/
package tcp

import (
	"bufio"
	"encoding/binary"
	"gamelog"
	"io"
	"net"
	"runtime"
	"runtime/debug"
	"time"
)

const (
	G_MsgId_Disconnect = 7111
	G_MsgId_Regist     = 7112
)

var (
	G_HandlerMsgMap = map[uint16]func(*TCPConn, []byte){
		G_MsgId_Regist: DoRegistToSvr,
	}
)

type TCPConn struct { //登录时将TCPConn指针写入player中
	conn       net.Conn
	reader     *bufio.Reader //包装conn减少conn.Read的io次数，见【common\net.go】
	writeChan  chan []byte
	isClose    bool
	onNetClose func(*TCPConn)
	Data       interface{}
}

func newTCPConn(conn net.Conn, pendingWriteNum int, callback func(*TCPConn)) *TCPConn {
	tcpConn := new(TCPConn)
	tcpConn.ResetConn(conn)
	tcpConn.onNetClose = callback
	tcpConn.writeChan = make(chan []byte, pendingWriteNum)
	return tcpConn
}

//isClose标记仅在ResetConn、Close中设置，其它地方只读
func (tcpConn *TCPConn) ResetConn(conn net.Conn) {
	tcpConn.conn = conn
	tcpConn.reader = bufio.NewReader(conn)
	tcpConn.isClose = false
}
func (tcpConn *TCPConn) Close() {
	if tcpConn.isClose {
		return
	}
	tcpConn.conn.Close()
	tcpConn.doWrite(nil) //触发writeRoutine结束
	tcpConn.isClose = true

	if tcpConn.onNetClose != nil {
		tcpConn.onNetClose(tcpConn)
	}
}

// msgdata must not be modified by other goroutines
func (tcpConn *TCPConn) WriteMsg(msgID uint16, msgdata []byte) {
	msgLen := uint16(len(msgdata))

	msgbuffer := make([]byte, 4+msgLen) //前2字节-msgLen；后2字节-msgID

	binary.LittleEndian.PutUint16(msgbuffer, msgLen)
	binary.LittleEndian.PutUint16(msgbuffer[2:], msgID)

	copy(msgbuffer[4:], msgdata)

	if false == tcpConn.isClose {
		tcpConn.doWrite(msgbuffer)
	}
}
func (tcpConn *TCPConn) doWrite(buf []byte) {
	select {
	case tcpConn.writeChan <- buf: //chan满后再写即阻塞，select进入default分支报错
	default:
		gamelog.Error("doWrite: channel full")
		tcpConn.conn.(*net.TCPConn).SetLinger(0)
		tcpConn.Close()
		// close(tcpConn.writeChan) //client重连chan里的数据得保留，server都是新new的
	}
}
func (tcpConn *TCPConn) writeRoutine() {
	for buf := range tcpConn.writeChan {
		if buf == nil {
			break
		}
		_, err := tcpConn.conn.Write(buf)
		if err != nil {
			gamelog.Error("WriteRoutine error: %s", err.Error())
			break
		}
	}
	tcpConn.Close()
}
func (tcpConn *TCPConn) readRoutine() {
	tcpConn.readLoop()
	tcpConn.Close()

	//通知业务层net断开
	tcpConn.msgDispatcher(G_MsgId_Disconnect, nil)
}
func (tcpConn *TCPConn) readLoop() error {
	var err error
	var msgHeader = make([]byte, 4) //前2字节-msgLen；后2字节-msgID
	var msgID uint16
	var msgLen uint16
	var firstTime bool = true

	for {
		if tcpConn.isClose {
			break
		}

		//TODO：client无需超时限制
		if firstTime == true {
			tcpConn.conn.SetReadDeadline(time.Now().Add(5000 * time.Second)) //首次读，5秒超时
			firstTime = false
		} else {
			tcpConn.conn.SetReadDeadline(time.Time{}) //后面读的就没有超时了
		}

		_, err = io.ReadFull(tcpConn.reader, msgHeader)
		if err != nil {
			gamelog.Error("ReadFull msgHeader error: %s", err.Error())
			return err
		}

		msgLen = binary.LittleEndian.Uint16(msgHeader)
		msgID = binary.LittleEndian.Uint16(msgHeader[2:])
		if msgLen <= 0 || msgLen > 10240 {
			gamelog.Error("ReadProcess Invalid msgLen :%d", msgLen)
			break
		}

		msgData := make([]byte, msgLen)
		_, err = io.ReadFull(tcpConn.reader, msgData)
		if err != nil {
			gamelog.Error("ReadFull msgData error: %s", err.Error())
			return err
		}

		tcpConn.msgDispatcher(msgID, msgData)
	}
	return nil
}
func (tcpConn *TCPConn) msgDispatcher(msgID uint16, pdata []byte) {
	// gamelog.Info("---msgID:%d, dataLen:%d", msgID, len(pdata))
	msghandler, ok := G_HandlerMsgMap[msgID]
	if !ok {
		gamelog.Error("msgid : %d have not a msg handler!!", msgID)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				gamelog.Error("msgID %d Error  %s", msgID, debug.Stack())
			}
		}
	}()
	msghandler(tcpConn, pdata)
}
</pre>
##Golang锁/线程间通讯实践指南
线程间通讯：互斥

锁的问题在哪？

- 最大问题：不易控制

锁

- lock 但忘记unlock 的结果是灾难性的，因为服务器相当于挂了（所有和该锁有关的代码都不能被
执行）！

次要问题：性能杀手

- 锁会导致代码串行化执行
- 但别误会：锁并不特别慢，比线程间通讯其他原语（同步、收发消息）要快很多！比锁快的东西：无锁、原子操作（比锁并不快太多）网上有人用golang 的channel 来实现锁，这很不正确


善用defer 和滥用defer

- 善用defer可以大大降低用锁的心智负担
- 滥用defer可能会导致锁粒度过大

控制锁粒度

- 不要在锁里面执行费时操作，会阻塞服务器，导致其他请求不能及时被响应

读写锁：sync.RWMutex

- 如果一个共享资源(不一定是一个变量，可能是一组变量)，绝大部分情况下是读操作，偶然有写操作
，则非常适合用读写锁。

锁数组：[]sync.Mutex

- 如果一个共享资源，有很强的分区特征，则非常适合用锁数组
- 比如一个网盘服务，网盘不同用户之间的资源彼此完全不相干
<pre>
var mutexs [N]sync.Mutex
mutex := &mutexs[uid % N] // 根据用户id选择锁
mutex.Lock()
defer mutex.Unlock()
</pre>
读写锁与互斥锁使用实例
<pre>
//
// 创建一个文件存放数据,在同一时刻,可能会有多个Goroutine分别进行对此文件的写操作和读操作.
// 每一次写操作都应该向这个文件写入若干个字节的数据,作为一个独立的数据块存在,这意味着写操作之间不能彼此干扰,写入的内容之间也不能出现穿插和混淆的情况
// 每一次读操作都应该从这个文件中读取一个独立完整的数据块.它们读取的数据块不能重复,且需要按顺序读取.
// 例如: 第一个读操作读取了数据块1,第二个操作就应该读取数据块2,第三个读操作则应该读取数据块3,以此类推
// 对于这些读操作是否可以被同时执行,不做要求. 即使同时进行,也应该保持先后顺序.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

//数据文件的接口类型
type DataFile interface {
	// 读取一个数据块
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号
	Rsn() int64
	// 获取最后写入的数据块的序列号
	Wsn() int64
	// 获取数据块的长度
	DataLen() uint32
}

//数据类型
type Data []byte

//数据文件的实现类型
type myDataFile struct {
	f       *os.File     //文件
	fmutex  sync.RWMutex //被用于文件的读写锁
	woffset int64        // 写操作需要用到的偏移量
	roffset int64        // 读操作需要用到的偏移量
	wmutex  sync.Mutex   // 写操作需要用到的互斥锁
	rmutex  sync.Mutex   // 读操作需要用到的互斥锁
	dataLen uint32       //数据块长度
}

//初始化DataFile类型值的函数,返回一个DataFile类型的值
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	//f,err := os.Create(path)
	if err != nil {
		fmt.Println("Fail to find", f, "cServer start Failed")
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}

	df := &myDataFile{
		f:       f,
		dataLen: dataLen,
	}

	return df, nil
}

//获取并更新读偏移量,根据读偏移量从文件中读取一块数据,把该数据块封装成一个Data类型值并将其作为结果值返回

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	var offset int64
	// 读互斥锁定
	df.rmutex.Lock()
	offset = df.roffset
	// 更改偏移量, 当前偏移量+数据块长度
	df.roffset += int64(df.dataLen)
	// 读互斥解锁
	df.rmutex.Unlock()

	//读取一个数据块,最后读取的数据块序列号
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	for {
		//读写锁:读锁定
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			//由于进行写操作的Goroutine比进行读操作的Goroutine少,所以过不了多久读偏移量roffset的值就会大于写偏移量woffset的值
			// 也就是说,读操作很快就没有数据块可读了,这种情况会让df.f.ReadAt方法返回的第二个结果值为代表的非nil且会与io.EOF相等的值
			// 因此不应该把EOF看成错误的边界情况
			// so 在读操作读完数据块,EOF时解锁读操作,并继续循环,尝试获取同一个数据块,直到获取成功为止.
			if err == io.EOF {
				//注意,如果在该for代码块被执行期间,一直让读写所fmutex处于读锁定状态,那么针对它的写操作将永远不会成功.
				//切相应的Goroutine也会被一直阻塞.因为它们是互斥的.
				// so 在每条return & continue 语句的前面加入一个针对该读写锁的读解锁操作
				df.fmutex.RUnlock()
				//注意,出现EOF时可能是很多意外情况,如文件被删除,文件损坏等
				//这里可以考虑把逻辑提交给上层处理.
				continue
			}
		}
		break
	}
	d = bytes
	df.fmutex.RUnlock()
	return
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	//读取并更新写的偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	//写入一个数据块,最后写入数据块的序号
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	df.fmutex.Unlock()
	_, err = df.f.Write(bytes)

	return
}

func (df *myDataFile) Rsn() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) Wsn() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func main() {
	//简单测试下结果
	var dataFile DataFile
	dataFile, _ = NewDataFile("./mutex_2015_1.dat", 10)

	var d = map[int]Data{
		1: []byte("batu_test1"),
		2: []byte("batu_test2"),
		3: []byte("test1_batu"),
	}

	//写入数据
	for i := 1; i < 4; i++ {
		go func(i int) {
			wsn, _ := dataFile.Write(d[i])
			fmt.Println("write i=", i, ",wsn=", wsn, ",success.")
		}(i)
	}

	//读取数据
	for i := 1; i < 4; i++ {
		go func(i int) {
			rsn, d, _ := dataFile.Read()
			fmt.Println("Read i=", i, ",rsn=", rsn, ",data=", d, ",success.")
		}(i)
	}

	time.Sleep(10 * time.Second)
}

//为什么我想对你说hello却说成：hello world？因为你是我的世界!
</pre>
###Golang自用库制作图片
<pre>
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	dd, _ := os.Create("out.gif") //输出为gif文件
	makeimgs(dd)
}

//这里可以输出任意类型数据
func makeimgs(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
</pre>
###Golang将访问URL地址实时打印到网页
<pre>
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8007", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
}
</pre>
升级版/计数
<pre>
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8007", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q   第%d次 \n", r.URL.Path, count)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
</pre>
###Linux du命令
<pre>
du -h //方便阅读的格式显示文件/文件夹使用空间
du -h --max-depth=1 //输出当前目录下各个子目录所使用的空间
</pre>
###Golang sync.Once只会执行一次
<pre>
package main

//整个程序，只会执行onces()方法一次,onced()方法是不会被执行的
import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	time.Sleep(4000)
}
func onces() {
	fmt.Println("onces")
}

func onced() {
	fmt.Println("onced")
}
</pre>
sync.WaitGroup 与 sync.Once
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
	fmt.Printf("%s", robots)
}

// This example fetches several URLs concurrently,
// using a WaitGroup to block until all the fetches are complete.
func ExampleWaitGroup() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.baidu.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			GetDemo(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	fmt.Println("--------------------group wait over ---------------------------")
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
	fmt.Println("once over -------------------------------")
}

func main() {

	ExampleOnce()
	ExampleWaitGroup()
}
</pre>
###Golang flag包的简单使用
<pre>
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
)

//base64加密解密 乱入内容，2333
func Decode(raw []byte) []byte {
	var buf bytes.Buffer
	decoded := make([]byte, 215)
	buf.Write(raw)
	decoder := base64.NewDecoder(base64.StdEncoding, &buf)
	decoder.Read(decoded)
	return decoded
}

func Encode(raw []byte) []byte {
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &encoded)
	encoder.Write(raw)
	encoder.Close()
	return encoded.Bytes()
}

func main() {
	var pi float64
	bpi := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(bpi)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println(pi)

	res := string(Encode([]byte("Jason")))
	fmt.Println(res)
	fmt.Println(string(Decode([]byte(res))))

	//返回一个相应类型的指针
	namePtr := flag.String("name", "Jason", "User's name")
	agePtr := flag.Int("age", 22, "User's age")
	vipPtr := flag.Bool("vip", true, "Is a vip user")
	var svar string
	flag.StringVar(&svar, "you", "总是心太软", "a string var")
	fmt.Println("name:", *namePtr)
	fmt.Println("age:", *agePtr)
	fmt.Println("vip:", *vipPtr)
	fmt.Println("you:", svar)

	//获取从编译时期带进来的参数 如：go run main.go param1 param2 param3
	flag.Parse()
	fmt.Println("all params : ", flag.Args())
}
</pre>
###Golang bufio包
<pre>
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// ReadString 功能同 ReadBytes，只不过返回的是一个字符串
	s := strings.NewReader("Jason Test bufio read from byte to string")
	str := bufio.NewReader(s)
	br, _ := str.ReadString('\n')
	fmt.Println("1--------->", br)

	// Peek 返回缓存的一个切片，该切片引用缓存中前n字节数据
	s2 := strings.NewReader("ABCDEFG")
	br2 := bufio.NewReader(s2)
	b2, _ := br2.Peek(5)
	fmt.Println(string(b2))
	// ReadByte 从 b 中读出一个字节并返回,没有数据则报错
	// ReadRune 从 b 中读出一个 UTF8 编码的字符并返回

	// WriteTo 实现了 io.WriterTo 接口
	s3 := strings.NewReader("S3_CONTENT")
	br3 := bufio.NewReader(s3)
	b3 := bytes.NewBuffer(make([]byte, 0))
	br3.WriteTo(b3)
	fmt.Println(b3) //输出：S3_CONTENT

	//bufio.NewWriter WriteByte Flush
	b4 := bytes.NewBuffer(make([]byte, 0))
	bw4 := bufio.NewWriter(b4)
	bw4.WriteByte('H')
	bw4.WriteByte('e')
	bw4.WriteByte('l')
	bw4.WriteByte('l')
	bw4.WriteByte('o')
	bw4.Flush()
	fmt.Println(b4) //输出：Hello
}
</pre>
###Golang container包
<pre>
package main

/*
由于目前golang 没有提供泛型机制，所以通用容器实现基本和 c 类似，
golang 用 interface{} 做转接， c 用 void * 转接。
*/
import (
	"container/heap"
	"container/ring"
	"fmt"
)

func operation(n, m int) []int {
	var res []int
	ring := ring.New(n)
	ring.Value = 1
	for i, p := 2, ring.Next(); p != ring; i, p = i+1, p.Next() {
		p.Value = i
	}
	h := ring.Prev()
	for h != h.Next() {
		for i := 1; i < m; i++ {
			h = h.Next()
		}
		res = append(res, h.Unlink(1).Value.(int))
	}
	res = append(res, h.Value.(int))
	return res
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {
	fmt.Println(operation(9, 5))
	h := &intHeap{10, 3, 9, 7, 2, 56, 67, 66}
	heap.Init(h)
	heap.Push(h, 1)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
</pre>
###Golang database/sql包
<pre>
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.220:3306)/zcm")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SUCCESS")
	}

	var (
		id         int
		config_key string
	)
	rows, err := db.Query("select id,config_key from config where id =?", 1)
	if err != nil {
		fmt.Println("select failed")
	}
	defer rows.Close()
	defer db.Close()

	for rows.Next() {
		err := rows.Scan(&id, &config_key)
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println(id, config_key)
	}
}
</pre>
###Golang errors包
<pre>
package main

import (
	"errors"
	"fmt"
)

//使用errors包创建自定义错误
type MyError struct{}

func (this *MyError) Error() string {
	return ""
}

func main() {
	var err1 error = errors.New("This is an error")
	fmt.Println(err1.Error())

	err2 := fmt.Errorf("%s", "the error test for fmt.Errorf")
	fmt.Println(err2.Error())
}
</pre>
###Golang text/template包
<pre>
package main

import (
	"os"
	"text/template"
)

func main() {
	type Inventory struct {
		Material string
		Count    int
	}

	NewInventory := Inventory{"Jason", 77}
	muban := `{{.Count}} items are made of {{.Material}}`
	tmpl, err := template.New("test").Parse(muban) //建立一个模板
	//如果将muban文件中内容放入muban.txt文件中，那么等价于tmpl,err := template.ParseFiles("muban.txt")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, NewInventory) //将struct与模板合成，合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}

	/*
	   ParseFiles接受一个字符串，字符串的内容是一个模板文件的路径（绝对路径or相对路径）
	   ParseGlob也差不多，是用正则的方式匹配多个文件
	   假设一个目录里有a.txt b.txt c.txt的话
	   用ParseFiles需要写3行对应3个文件，如果有一万个文件呢？
	   而用ParseGlob只要写成template.ParseGlob("*.txt") 即可
	*/

}
</pre>
###Golang并发分割文件
<pre>
package main
 
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "runtime"
    "strings"
    "sync"
    "time"
)
 
func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}
 
func main() {
    t0 := time.Now()
    file, err := os.Open("test.log")
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }
    defer file.Close() //文件关闭
 
    ch := make(chan string, 10240)
    var wg = new(sync.WaitGroup)
    wg.Add(1)
 
    go writeFile(wg, ch)
 
    //将文件作为一个io.Reader对象进行buffered I/O操作
    //每次读取一行,处理一行
    br := bufio.NewReaderSize(file, 1024*1024*64)
    for {
        line, isPrefix, err := br.ReadLine()
        if err == io.EOF {
            break
        }
        for isPrefix && err == nil {
            println("isPrefix==true")
            var rest []byte
            rest, isPrefix, err = br.ReadLine()
            line = append(line, rest...)
        }
        strLine := string(line)
        ch <- strLine
 
    }
    close(ch)
    wg.Wait()
 
    t := time.Now()
    fmt.Println(t.Sub(t0).Seconds())
}
 
func writeFile(wg *sync.WaitGroup, ch chan string) {
    defer wg.Done()
    var m = make(map[string]*bufio.Writer, 10)
    for line := range ch {
        key := getKey(line)
        if _, ok := m[key]; ok == false {
            file, err := os.Create(key + ".log")
            if err != nil {
                panic(err)
            }
            bw := bufio.NewWriterSize(file, 1024*1024)
            m[key] = bw
            defer func() {
                bw.Flush()
                file.Close()
            }()
        }
        bw := m[key]
        bw.WriteString(line)
    }
}
 
func getKey(s string) (key string) {
    str := strings.Split(s, "//")
    stringKey := strings.Split(str[1], "/")
    //  fmt.Println(stringKey[0])
    key = stringKey[0]
    return
}
</pre>
###Golang去掉文件名的后缀
<pre>
package main

import (
    "fmt"
    "runtime"
    "path"
    "strings"
)

func main() {
    _, fulleFilename, line, _ := runtime.Caller(0)
    fmt.Println(fulleFilename)
    fmt.Println(line)
    var filenameWithSuffix string
    filenameWithSuffix = path.Base(fulleFilename)
    fmt.Println("filenameWithSuffix=", filenameWithSuffix)
    var fileSuffix string
    fileSuffix = path.Ext(filenameWithSuffix)
    fmt.Println("fileSuffix=", fileSuffix)
    var filenameOnly string
    filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
    fmt.Println("filenameOnly=", filenameOnly)
}
</pre>
###Linux中的小技巧Tips
<pre>
ls -lh   //详细列出所有文件的权限信息、用户归属以及大小所占内存[单位不是Kb]等信息
ls -Slh  //按照所占内存空间大小倒序排列 
ls -Slhr //按照所占内存空间大小整序排列
date -s 时间/日期  //修改系统时间
</pre>
###Golang设置Keep_alive（keep_alive）时间
需要设置的是：浏览器与服务器空闲关闭时间，例如30s没有发送请求、返回数据，则服务器关闭与浏览器之间connection。设置了ReadTimeout以后，如果指定时间没有接收到数据，服务器就会关闭客户端连接。
<pre>
package main
import (
    "fmt"
    "net"
    "net/http"
    "io"
    "time"
)
func ListenAndServe(addr string, handler http.Handler, timeout time.Duration) error {
    server := &http.Server{
        Addr:        addr,
        Handler:     handler,
        ReadTimeout: timeout,
    }
    return server.ListenAndServe()
}
func main() {
    addr := "127.0.0.1:6061"
    http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {})
    go ListenAndServe(addr, nil, time.Second*10)
    time.Sleep(time.Second)
    started := time.Now()
    remoteAddr, _ := net.ResolveTCPAddr("tcp4", addr)
    conn, err := net.DialTCP("tcp4", nil, remoteAddr)
    if err != nil {
        panic("failed to connect")
    }
    defer conn.Close()
    _, err = conn.Read(make([]byte, 128))
    if err != io.EOF {
        panic("should return EOF")
    }
    fmt.Printf("time escaped=%s, error=%s\n", time.Now().Sub(started), err)
}
</pre>
###Golang Strings使用
<pre>
package main

//Stringer 是一个可以用字符串描述自己的类型
import (
	"fmt"
)

type Test struct {
	Name string
	QQ   string
}

func (t Test) String() string {
	return fmt.Sprintf("name=%v|qq=%v", t.Name, t.QQ)
}

func main() {
	a := Test{"Jason", "903456967"}
	fmt.Println(a)
}
</pre>
###Beego里面的404情况处理
<pre>
func main() {
	beego.ErrorHandler("404", page_not_found)
	beego.Run()
}
func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
</pre>
###Golang中json流式编解码
<pre>
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

type Dog struct {
    Name string
    Age  int
}

type Car struct {
    Engine string
    Count  int
}

func main() {
    var data = []byte(`
 {"Name":"jemy","Age":26}{"Engine":"Power","Count":4}
`)

    var d Dog
    var c Car

    decoder := json.NewDecoder(bytes.NewReader(data))
    decoder.Decode(&d)
    decoder.Decode(&c)

    fmt.Println(d)
    fmt.Println(c)
}
output==>
{jemy 26}
{Power 4}
</pre>
<pre>
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

type Dog struct {
    Name string
    Age  int
}

type Car struct {
    Engine string
    Count  int
}

func main() {
    var dog = Dog{
        Name: "jemy",
        Age:  26,
    }
    var car = Car{
        Engine: "Power",
        Count:  4,
    }

    buffer := bytes.NewBuffer(nil)
    encoder := json.NewEncoder(buffer)
    if err := encoder.Encode(&dog); err != nil {
        fmt.Println(err)
        return
    }

    if err := encoder.Encode(&car); err != nil {
        fmt.Println(err)
    }

    fmt.Print(buffer.String())
}
output==>
{"Name":"jemy","Age":26}
{"Engine":"Power","Count":4}
</pre>
###Golang并发不安全示例
<pre>
package main

import "fmt"

func main() {
    c := make(map[string]int)

    for i := 0; i < 100; i++ {
        go func() {
            for j := 0; j < 1000000; j++ {
                c[fmt.Sprintf("%d", j)] = j
            }
        }()
    }
}
output==>
fatal error: concurrent map writes

goroutine 19 [running]:
runtime.throw(0x5290a0, 0x15)
        C:/Go/src/runtime/panic.go:547 +0x97 fp=0xc082063e50 sp=0xc082063e38
runtime.mapassign1(0x4cb980, 0xc082037d70, 0xc082063f68, 0xc082063f38)
        C:/Go/src/runtime/hashmap.go:445 +0xb8 fp=0xc082063ef8 sp=0xc082063e50
main.main.func1(0xc082037d70)
        D:/gopath/src/test/t3.go:47 +0x15b fp=0xc082063f98 sp=0xc082063ef8
runtime.goexit()
        C:/Go/src/runtime/asm_amd64.s:1998 +0x1 fp=0xc082063fa0 sp=0xc082063f98
created by main.main
        D:/gopath/src/test/t3.go:49 +0x82
</pre>
###Golang将(字节切片) []byte 转换成 string 类型
<pre> 
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

func main() {
	bytes := [4]byte{1, 2, 3, 4}
	str := convert(bytes[:])
	fmt.Println(str)
}
</pre>
###Golang UDP、TCP[socket编程]
UDP服务端
<pre>
package main

//UDP服务端
import (
	"fmt"
	"net"
	"os"
	"time"
)

/*
先通过net.ResolveUDPAddr创建监听地址
net.ListenUDP创建监听链接
然后通过conn.ReadFromUDP和conn.WriteToUDP收发UDP报文
*/
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func recvUDPMsg(conn *net.UDPConn) {
	var buf [20]byte
	n, raddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println("msg is ", string(buf[0:n]))
	time.Sleep(10 * time.Second)
	_, err = conn.WriteToUDP([]byte("nice to see u"), raddr)
	checkError(err)
}

func main() {
	udp_addr, err := net.ResolveUDPAddr("udp", ":11110")
	checkError(err)

	conn, err := net.ListenUDP("udp", udp_addr)
	defer conn.Close()
	checkError(err)
	recvUDPMsg(conn)
}
</pre>
UDP客户端
<pre>
package main

//UDP客户端
import (
	"fmt"
	"net"
	"os"
)

/*
先通过net.Dial(“udp”, “127.0.0.1:11110”)，建立发送报文至本机11110端口的socket，
然后使用conn.Write和conn.Read收发包，当然conn.ReadFromUDP和conn.WriteToUDP也是可以的
*/
func udp_client() {
	conn, err := net.Dial("udp", "127.0.0.1:11110")
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}
	conn.Write([]byte("Hello Jason!"))
	fmt.Println("send msg")
	var msg [20]byte
	conn.Read(msg[0:])
	fmt.Println("msg is", string(msg[0:10]))
}

func main() {
	udp_client()
}
</pre>
TCP服务端
<pre>
package main

//TCP Socket服务端
import (
	"fmt"
	"io"
	"net"
)

const RECV_BUF_LEN = 1024

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6666") //侦听在 6666 端口
	if err != nil {
		panic("error listening:" + err.Error())
	}
	fmt.Println("Starting the server")

	for {
		conn, err := listener.Accept() //接受连接
		if err != nil {
			panic("Error accept:" + err.Error())
		}
		go EchoServer(conn)
	}
}

func EchoServer(conn net.Conn) {
	buf := make([]byte, RECV_BUF_LEN)
	defer conn.Close()

	for {
		fmt.Println("Accepted the Connection :", conn.RemoteAddr())
		n, err := conn.Read(buf)
		switch err {
		case nil:
			conn.Write(buf[0:n])
		case io.EOF:
			fmt.Printf("Warning: End of data: %s \n", err)
			return
		default:
			fmt.Printf("Error: Reading data : %s \n", err)
			return
		}
	}
}
</pre>
TCP客户端
<pre>
package main

//TCP Socket客户端
import (
	"fmt"
	"net"
	"time"
)

const RECV_BUF_LEN = 1024

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	buf := make([]byte, RECV_BUF_LEN)

	for i := 0; i < 5; i++ {
		//准备要发送的字符串
		msg := fmt.Sprintf("Hello World, %03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("Write Buffer Error:", err.Error())
			break
		}
		fmt.Println("发出的信息:", msg)

		//从服务器端收字符串
		n, err = conn.Read(buf)
		if err != nil {
			println("Read Buffer Error:", err.Error())
			break
		}
		fmt.Println("接收的信息:", string(buf[0:n]))
		//等一秒钟
		time.Sleep(time.Second)
	}
}
</pre>
###Golang代理服务Proxy
<pre>
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", "localhost:1987")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr fail.", err)
		os.Exit(1)
	}

	socket, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("net.DialUDP fail.", err)
		os.Exit(1)
	}
	defer socket.Close()
	r := bufio.NewReader(os.Stdin)
	for {
		switch line, ok := r.ReadString('\n'); true {
		case ok != nil:
			fmt.Printf("bye bye!\n")
			return
		default:
			socket.Write([]byte(line))
			data := make([]byte, 1024)
			_, remoteAddr, err := socket.ReadFromUDP(data)
			if err != nil {
				fmt.Println("error recv data")
				return
			}
			fmt.Printf("from %s:%s\n", remoteAddr.String(), string(data))
		}
	}
}

package main

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	ip   = "127.0.0.1"
	port = 1987
)

const (
	proxy_timeout = 2
	proxy_server  = "localhost:8989"
	msg_length    = 1024
)

type Request struct {
	reqId      int
	reqContent string
	rspChan    chan<- string  // writeonly chan
}

var requestMap map[int]*Request

type Clienter struct {
	client  net.Conn
	isAlive bool
	SendStr chan *Request
	RecvStr chan string
}

func (c *Clienter) Connect() bool {
	if c.isAlive {
		return true
	} else {
		var err error
		c.client, err = net.Dial("tcp", proxy_server)
		if err != nil {
			return false
		}
		c.isAlive = true
		log.Println("connect to " + proxy_server)
	}
	return true
}

func ProxySendLoop(c *Clienter) {

	//store reqId and reqContent
	senddata := make(map[string]string)
	for {
		if !c.isAlive {
			time.Sleep(1 * time.Second)
			c.Connect()
		}
		if c.isAlive {
			req := <-c.SendStr

			//construct request json string
			senddata["reqId"] = strconv.Itoa(req.reqId)
			senddata["reqContent"] = req.reqContent
			sendjson, err := json.Marshal(senddata)
			if err != nil {
				continue
			}

			_, err = c.client.Write([]byte(sendjson))
			if err != nil {
				c.RecvStr <- string("proxy server close...")
				c.client.Close()
				c.isAlive = false
				log.Println("disconnect from " + proxy_server)
				continue
			}
			// log.Println("Write to proxy server: " + string(sendjson))
		}
	}
}

func ProxyRecvLoop(c *Clienter) {
	buf := make([]byte, msg_length)
	recvdata := make(map[string]string, 2)
	for {
		if !c.isAlive {
			time.Sleep(1 * time.Second)
			c.Connect()
		}
		if c.isAlive {
			n, err := c.client.Read(buf)
			if err != nil {
				c.client.Close()
				c.isAlive = false
				log.Println("disconnect from " + proxy_server)
				continue
			}
			//log.Println("Read from proxy server: " + string(buf[0:n]))

			if err := json.Unmarshal(buf[0:n], &recvdata); err == nil {
				reqidstr := recvdata["reqId"]
				if reqid, err := strconv.Atoi(reqidstr); err == nil {
					req, ok := requestMap[reqid]
					if !ok {
						continue
					}
					req.rspChan <- recvdata["resContent"]
				}
				continue
			}
		}
	}
}

func handle(conn *net.UDPConn, remote *net.UDPAddr, id int, tc *Clienter, data []byte) {

	handleProxy := make(chan string)
	request := &Request{reqId: id, rspChan: handleProxy}

	request.reqContent = string(data)

	requestMap[id] = request
	//send to proxy
	select {
	case tc.SendStr <- request:
	case <-time.After(proxy_timeout * time.Second):
		conn.WriteToUDP([]byte("proxy server send timeout."), remote)
	}

	//read from proxy
	select {
	case rspContent := <-handleProxy:
		conn.WriteToUDP([]byte(rspContent), remote)
	case <-time.After(proxy_timeout * time.Second):
		conn.WriteToUDP([]byte("proxy server recv timeout."), remote)
	}
}

func UdpLotusMain(ip string, port int) {
	//start tcp server
	addr, err := net.ResolveUDPAddr("udp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalln("net.ResolveUDPAddr fail.", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln("net.ListenUDP fail.", err)
		//os.Exit(1)
		return
	}
	log.Println("start udp server " + ip + " " + strconv.Itoa(port))
	defer conn.Close()

	//start proxy connect and loop
	var tc Clienter
	tc.SendStr = make(chan *Request, 1000)
	tc.RecvStr = make(chan string)
	tc.Connect()
	go ProxySendLoop(&tc)
	go ProxyRecvLoop(&tc)

	//listen new request
	requestMap = make(map[int]*Request)

	buf := make([]byte, msg_length)
	var id int = 0
	for {
		rlen, remote, err := conn.ReadFromUDP(buf)
		if err == nil {
			id++
			log.Println("connected from " + remote.String())
			go handle(conn, remote, id, &tc, buf[:rlen]) //new thread
		}
	}
}

func main() {
	UdpLotusMain(ip, port)
}
</pre>
###Golang内存分配包
github.com/funny/slab
###image变成[]byte
<pre>
 //将image图片转换成[]byte类型
 buf := new(bytes.Buffer)
 err := jpeg.Encode(buf, new_image, nil)
 send_s3 := buf.Bytes()
</pre>
###nohup的用法
//0、1和2分别表示标准输入、标准输出和标准错误信息输出，可以用来指定需要重定向的标准输入或输出。
<pre>
nohup ./program_name &   //默认输出当前目录下 nohup.out 日志文件
nohup ./program_name >/dev/null 2>log &   //在当前目录下产生只记录错误输出信息的log文件
nohup ./program_name >/dev/null 2>&1 & //不产生任何日志文件
</pre>
###container包
双向链表 list
<pre>
package main

import (
	"container/list"
	"fmt"
)

//list是一个双向链表。该结构具有链表的所有功能
func main() {
	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("---------------------")
	fmt.Println("首部元素的值：", l.Front().Value) //首部元素
	fmt.Println("尾部元素的值：", l.Back().Value)  //尾部元素
	l.InsertAfter(6, l.Front())             //首部元素之后插入一个值为6的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("---------------------")
	l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("---------------------")
	l.MoveToFront(l.Back()) //将尾部元素移动到首部
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("---------------------")
	l.Init()
}
</pre>
ring  优先级队列heap
<pre>
package main

import (
	"container/heap"
	"container/ring"
	"fmt"
)

func josephus(n, m int) []int {
	var res []int
	ring := ring.New(n)
	ring.Value = 1
	for i, p := 2, ring.Next(); p != ring; i, p = i+1, p.Next() {
		p.Value = i
	}
	h := ring.Prev()
	for h != h.Next() {
		for i := 1; i < m; i++ {
			h = h.Next()
		}
		res = append(res, h.Unlink(1).Value.(int))
	}
	res = append(res, h.Value.(int))
	return res
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println("----------循环双向链表-----------")
	fmt.Println(josephus(10, 5))
	fmt.Println("-------二叉堆(binary heap)-------")
	h := &intHeap{11, 12, 13, 14, 15}
	heap.Init(h)
	heap.Push(h, 1) //头部插入一个新的元素
	heap.Pop(h)     //头部弹出第一个元素
	heap.Pop(h)     //头部弹出第一个元素
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
</pre>
stack
<pre>
package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("长度：", stack.Len())
	fmt.Println(stack.Peak().(int))
	fmt.Println(stack.Pop().(int))
}
</pre>
线程安全的优先级队列[排序 唯一值]
<pre>
package main

import (
	"container/heap"
	"fmt"
	"sort"
	"sync"
)

// An Item is something we manage in a priority queue.
type Item struct {
	Key      interface{} //The unique key of the item.
	Value    interface{} // The value of the item; arbitrary.
	Priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}
type ItemSlice struct {
	items    []*Item
	itemsMap map[interface{}]*Item
}

func (s ItemSlice) Len() int { return len(s.items) }
func (s ItemSlice) Less(i, j int) bool {
	return s.items[i].Priority < s.items[j].Priority
}
func (s ItemSlice) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
	s.items[i].Index = i
	s.items[j].Index = j
	if s.itemsMap != nil {
		s.itemsMap[s.items[i].Key] = s.items[i]
		s.itemsMap[s.items[j].Key] = s.items[j]
	}
}
func (s *ItemSlice) Push(x interface{}) {
	n := len(s.items)
	item := x.(*Item)
	item.Index = n
	s.items = append(s.items, item)
	s.itemsMap[item.Key] = item
}
func (s *ItemSlice) Pop() interface{} {
	old := s.items
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	delete(s.itemsMap, item.Key)
	s.items = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (s *ItemSlice) update(key interface{}, value interface{}, priority int) {
	item := s.itemByKey(key)
	if item != nil {
		s.updateItem(item, value, priority)
	}
}

// update modifies the priority and value of an Item in the queue.
func (s *ItemSlice) updateItem(item *Item, value interface{}, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(s, item.Index)
}
func (s *ItemSlice) itemByKey(key interface{}) *Item {
	if item, found := s.itemsMap[key]; found {
		return item
	}
	return nil
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	slice   ItemSlice
	maxSize int
	mutex   sync.RWMutex
}

func (pq *PriorityQueue) Init(maxSize int) {
	pq.slice.items = make([]*Item, 0, pq.maxSize)
	pq.slice.itemsMap = make(map[interface{}]*Item)
	pq.maxSize = maxSize
}
func (pq PriorityQueue) Len() int {
	pq.mutex.RLock()
	size := pq.slice.Len()
	pq.mutex.RUnlock()
	return size
}
func (pq *PriorityQueue) minItem() *Item {
	len := pq.slice.Len()
	if len == 0 {
		return nil
	}
	return pq.slice.items[0]
}
func (pq *PriorityQueue) MinItem() *Item {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()
	return pq.minItem()
}
func (pq *PriorityQueue) PushItem(key, value interface{}, priority int) (bPushed bool) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	size := pq.slice.Len()
	item := pq.slice.itemByKey(key)
	if size > 0 && item != nil {
		pq.slice.updateItem(item, value, priority)
		return true
	}
	item = &Item{
		Value:    value,
		Key:      key,
		Priority: priority,
		Index:    -1,
	}
	if pq.maxSize <= 0 || size < pq.maxSize {
		heap.Push(&(pq.slice), item)
		return true
	}
	min := pq.minItem()
	if min.Priority >= priority {
		return false
	}
	heap.Pop(&(pq.slice))
	heap.Push(&(pq.slice), item)
	return true
}
func (pq PriorityQueue) GetQueue() []interface{} {
	items := pq.GetQueueItems()
	values := make([]interface{}, len(items))
	for i := 0; i < len(items); i++ {
		values[i] = items[i].Value
	}
	return values
}
func (pq PriorityQueue) GetQueueItems() []*Item {
	size := pq.Len()
	if size == 0 {
		return []*Item{}
	}
	s := ItemSlice{}
	s.items = make([]*Item, size)
	pq.mutex.RLock()
	for i := 0; i < size; i++ {
		s.items[i] = &Item{
			Value:    pq.slice.items[i].Value,
			Priority: pq.slice.items[i].Priority,
		}
	}
	pq.mutex.RUnlock()
	sort.Sort(sort.Reverse(s))
	return s.items
}
func main() {
	var Real PriorityQueue
	Real.Init(100)
	//第一个为key[唯一，否则会被覆盖]，第二为value,第三个是排序字段[大的排前面]
	Real.PushItem("1", "queue1", 1)
	Real.PushItem("2", "queue2", 2)
	Real.PushItem("2", "queue3", 3)
	Real.PushItem("3", "queue4", 4)
	for i := 0; i < Real.Len(); i++ {
		fmt.Println(Real.GetQueue()[i].(string))
	}
}
output==>
queue4
queue3
queue1
</pre>
###Golang RPC远程调用实战
服务端
<pre>
package main

 import (
 	"log"
 	"net"
 	"net/http"
 	"net/rpc"
 )

 type Echo int

 func (t *Echo) Hi(args string, reply *string) error {
 	*reply = "Jason's server return:" + args
 	return nil
 }

 func main() {
 	rpc.Register(new(Echo))
 	rpc.HandleHTTP()
 	l, e := net.Listen("tcp", ":1234")
 	if e != nil {
 		log.Fatal("listen error:", e)
 	}
 	http.Serve(l, nil)
 }
</pre>
客户端
<pre>
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var args = "-----------"
	var reply string
	//Echo.Hi将要调用的函数名，args传过去的参数[非必须]，&replay服务器返回结果
	err = client.Call("Echo.Hi", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(reply)
}
output=>
Jason's server return:-----------
</pre>
###Golang 1.7 SSA编译器 反射 reflect
<pre>
package main

/*
由于Golang的SSA的编译器，变得非常聪明了，因此会把使用反射reflect.StringHeader，reflect.SliceHeader返回值中的uintptr指向的内存块，当成了没有被使用的内存块回收了。
解决方法：
一是尽量不要过分追求性能，使用反射reflect和unsafe包内的函数。这样能避免一些诡异的、很难分析的bug出现。
如果非要使用反射reflect和unsafe包内的函数，请注意一定要使用runtime.KeepAlive告诉SSA编译器，在指定的代码段内，不要回收内存块。
*/
import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func SimpleCrc(ptr uintptr, size int) int {
	ret := 0
	maxPtr := ptr + uintptr(size)
	for ptr < maxPtr {
		b := *(*byte)(unsafe.Pointer(ptr))
		ret += int(b)
		ptr++
	}
	return ret
}

//模拟申请内存，触发Gc回收内存
func Allocation(size int) {
	var free []byte
	free = make([]byte, size)
	if len(free) == 0 {
		panic("Allocation Error")
	}
}

func SliceCrcTest(slice []byte, N int) (ret int) {
	newSlice := []byte(string(slice))                       //获取独立内存
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&newSlice)) //反射切片结构
	ptr, size := uintptr(sh.Data), sh.Len                   //获取地址尺寸
	runtime.GC()                                            //强制内存回收
	for i := 0; i < N; i++ {
		ret = SimpleCrc(ptr, size) //计算crc校验码
		Allocation(size)           //模拟申请内存，触发Gc回收内存
	}

	//runtime.KeepAlive(newSlice) //本行一旦注释后结果不再是1665，取消注释节正确
	return
}

func StringCrcTest(str string, N int) (ret int) {
	newStr := string([]byte(str))                          //获取独立内存
	runtime.SetFinalizer(&newStr, func(x *string) {})      //设置回收事件
	sh := (*reflect.StringHeader)(unsafe.Pointer(&newStr)) //反射字符串结构
	ptr, size := uintptr(sh.Data), sh.Len                  //获取地址尺寸
	runtime.GC()                                           //强制内存回收
	for i := 0; i < N; i++ {
		ret = SimpleCrc(ptr, size) //计算crc校验码
		Allocation(size)           //模拟申请内存，触发Gc回收内存
	}

	//runtime.KeepAlive(newStr) //本行一旦注释后结果不再是1665，取消注释节正确
	return
}

func main() {
	var B = []byte("1234567890-1234567890-1234567890") //Crc的值为：1665
	var S = string(B)                                  //生成字符串
	N := 1000000                                       //循环执行1,000,000次
	fmt.Printf("SimpleCrc(\"%s\") = %v\n", B, SliceCrcTest(B, N))
	fmt.Printf("SimpleCrc(\"%s\") = %v\n", B, StringCrcTest(S, N))
}
</pre>
###Golang 工作线程池
<pre>
package main

import "fmt"
import "time"

// 使用goroutine  开启大小为3的线程池
// 其中1个channel为执行做通信，1个对结果进行保存

// 创建的worker
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
func main() {
	// 创建channel
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 3个worker作为一个pool
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送9个jobs，然后关闭
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后收集结果
	for a := 1; a <= 9; a++ {
		<-results
	}
}
</pre>
###Golang 跳出for-select死循环方法优化
<pre>
package main

import "fmt"
import "time"

/*
func main() {

L:
    for {
        select {
        case <-time.After(time.Second):
            fmt.Println("hello")
        default:
            break L
        }
    }

    fmt.Println("ending")
}
如你所见，需要联合break使用标签。这有其用途，不过我不喜欢。这个例子中的 for 循环看起来很小，但是通常它们会更大，而判断break的条件也更为冗长。
如果需要退出循环，我会将 for-select 封装到函数中:
*/
func foo() {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("hello")
		default:
			fmt.Println("default")
			return
		}
	}
}

func main() {
	foo()
	fmt.Println("ending")
}
</pre>
###Golang 自定义类型
<pre>
package main

import (
	"fmt"
	"strings"
)

/*
将 slice 或 map 定义成自定义类型可以让代码维护起来更加容易。假设有一个Server类型和一个返回服务器列表的函数,增加自定义筛选条件。
*/

type Server struct {
	Name string
}

type Servers []Server

// ListServers 返回服务器列表
func ListServers() Servers {
	return []Server{
		{Name: "Server1"},
		{Name: "Server2"},
		{Name: "Foo1"},
		{Name: "Foo2"},
	}
}

func (s Servers) Filter(name string) Servers {
	filtered := make(Servers, 0)

	for _, server := range s {
		if strings.Contains(server.Name, name) {
			filtered = append(filtered, server)
		}

	}

	return filtered
}

//弹性扩展
// func (s Servers) Check()
// func (s Servers) AddRecord()
// func (s Servers) Len()

func main() {
	servers := ListServers()
	servers = servers.Filter("Foo")
	fmt.Printf("servers %+v\n", servers)
}
</pre>
###Golang context 上下文
<pre>
package main

import (
	"fmt"
	"sync"
)

/*
withContext 封装函数
自定义上下文，让代码更优雅
P对象(processor) 代表cpu，M(work thread)代表工作线程，G对象（goroutine) 上下文切换
*/

var Mutex sync.Mutex

func withMutexContext(fn func()) {
	//可添加更多公用方法属性
	Mutex.Lock()
	defer Mutex.Unlock()
	fn()
	fmt.Println("--------已执行上下文--------")
}

func foo() {
	withMutexContext(func() {
		fmt.Println("--->foo()")
	})
}

func bar() {
	withMutexContext(func() {
		fmt.Println("--->bar()")
	})
}
func main() {
	foo()
	bar()
}
</pre>
###Golang 线程安全【并发安全】的map setter getter
<pre>
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const defaultShardCunt uint8 = 32
const seed uint32 = 131

type syncMap struct {
	items map[string]interface{}
	sync.RWMutex
}

type SyncMap struct {
	shardCount uint8
	shards     []*syncMap
}

func New() *SyncMap {
	return NewWithShard(defaultShardCunt)
}

func NewWithShard(shardCount uint8) *SyncMap {
	m := new(SyncMap)
	m.shardCount = shardCount
	m.shards = make([]*syncMap, m.shardCount)
	for i, _ := range m.shards {
		m.shards[i] = &syncMap{items: make(map[string]interface{})}
	}
	return m
}
func bkdrHash(str string) uint32 {
	var h uint32

	for _, c := range str {
		h = h*seed + uint32(c)
	}

	return h
}

func (m *SyncMap) locate(key string) *syncMap {
	return m.shards[bkdrHash(key)&uint32((m.shardCount-1))]
}

func (m *SyncMap) Get(key string) (value interface{}, ok bool) {
	shard := m.locate(key)
	shard.RLock()
	value, ok = shard.items[key]
	shard.RUnlock()
	return
}

func (m *SyncMap) Set(key string, value interface{}) {
	shard := m.locate(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

func (m *SyncMap) Delete(key string) {
	shard := m.locate(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}
func (m *SyncMap) Has(key string) bool {
	_, ok := m.Get(key)
	return ok
}
func (m *SyncMap) Size() int {
	size := 0
	for _, shard := range m.shards {
		shard.RLock()
		size += len(shard.items)
		shard.RUnlock()
	}
	return size
}

func (m *SyncMap) Flush() int {
	size := 0
	for _, shard := range m.shards {
		shard.Lock()
		size += len(shard.items)
		shard.items = make(map[string]interface{})
		shard.Unlock()
	}
	return size
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	test := New()
	test.Set("test", 1)
	test.Set("jason", 2)
	test.Flush()
	if test.Has("test") {
		fmt.Println(test.Get("test"))
	} else {
		fmt.Println("Nodata!")
	}
}
</pre>
###Golang 使用io.Pipe
客户端

主要是来解决当发送的数据量很大的时候出现严重影响性能的问题。
<pre>
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	cli := http.Client{}

	msg := struct {
		Name, Addr string
		Price      float64
	}{
		Name:  "hello",
		Addr:  "beijing",
		Price: 123.56,
	}
	r, w := io.Pipe()
	// 注意这边的逻辑！！
	go func() {
		defer func() {
			time.Sleep(time.Second * 2)
			log.Println("encode完成")
			// 只有这里关闭了，Post方法才会返回
			w.Close()
		}()
		log.Println("管道准备输出")
		// 只有Post开始读取数据，这里才开始encode，并传输
		err := json.NewEncoder(w).Encode(msg)
		log.Println("管道输出数据完毕")
		if err != nil {
			log.Fatalln("encode json failed:", err)
		}
	}()
	time.Sleep(time.Second * 1)
	log.Println("开始从管道读取数据")
	resp, err := cli.Post("http://localhost:9999/json", "application/json", r)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("POST传输完成")

	body := resp.Body
	defer body.Close()

	if body_bytes, err := ioutil.ReadAll(body); err == nil {
		log.Println("response:", string(body_bytes))
	} else {
		log.Fatalln(err)
	}
}
</pre> 
服务端[便于调试]
<pre>
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile)
	println("----->服务器启动成功！")
}

func main() {
	http.HandleFunc("/json", handleJson)
	http.ListenAndServe(":9999", nil)
}

func handleJson(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body := req.Body
		defer body.Close()
		body_bytes, err := ioutil.ReadAll(body)
		if err != nil {
			log.Println(err)
			resp.Write([]byte(err.Error()))
			return
		}
		j := map[string]interface{}{}
		if err := json.Unmarshal(body_bytes, &j); err != nil {
			log.Println(err)
			resp.Write([]byte(err.Error()))
			return
		}
		println("收到的信息：", string(body_bytes))
		resp.Write(body_bytes)
	} else {
		resp.Write([]byte("请使用post方法!"))
	}
}
</pre>
###哈希适用于高并发系统的重要原因之一
但凡大数据处理，高并发系统，必言哈希，随机插入，时间复杂度O(1)，随便查询，时间复杂度O(1)，除了耗费点空间以外，几乎没什么缺点了，在现在这个内存廉价的时代，哈希表变成了一个高并发系统的标配。
###Golang hash+list 哈希表+链表 LRU算法 缓存算法
LRU : Least Recently Used
核心思想是：如果数据最近被访问过，那么将来被访问的几率也更高.
<pre>
package main

//LRU算法
/*
1.key记录在map
2.对于set/get添加或命中的元素移到链表头
3.如总个数大于Cache容量(cap),则将最末的元素移除.
*/
import (
	"container/list"
	"errors"
	"fmt"
)

type CacheNode struct {
	Key, Value interface{}
}

func (cnode *CacheNode) NewCacheNode(k, v interface{}) *CacheNode {
	return &CacheNode{k, v}
}

type LRUCache struct {
	Capacity int
	dlist    *list.List
	cacheMap map[interface{}]*list.Element
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Capacity: cap,
		dlist:    list.New(),
		cacheMap: make(map[interface{}]*list.Element),
	}
}

func (lru *LRUCache) Size() int {
	return lru.dlist.Len()
}

func (lru *LRUCache) Set(k, v interface{}) error {
	if lru.dlist == nil {
		return errors.New("LRUCache结构体未初始化")
	}
	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pElement)
		pElement.Value.(*CacheNode).Value = v
		return nil
	}
	newElement := lru.dlist.PushFront(&CacheNode{k, v})
	lru.cacheMap[k] = newElement
	if lru.dlist.Len() > lru.Capacity {
		//移除最后一个
		lastElement := lru.dlist.Back()
		if lastElement == nil {
			return nil
		}
		CacheNode := lastElement.Value.(*CacheNode)
		delete(lru.cacheMap, CacheNode.Key)
		lru.dlist.Remove(lastElement)
	}
	return nil
}

func (lru *LRUCache) Get(k interface{}) (v interface{}, ret bool, err error) {
	if lru.cacheMap == nil {
		return v, false, errors.New("LRUCache结构体未初始化")
	}
	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pElement)
		return pElement.Value.(*CacheNode).Value, true, nil
	}
	return v, false, nil
}

func (lru *LRUCache) Remove(k interface{}) bool {
	if lru.cacheMap == nil {
		return false
	}
	if pElement, ok := lru.cacheMap[k]; ok {
		CacheNode := pElement.Value.(*CacheNode)
		delete(lru.cacheMap, CacheNode.Key)
		lru.dlist.Remove(pElement)
		return true
	}
	return false
}

func main() {
	LRU := NewLRUCache(3)
	//jason1 将被移除 先进先出 list
	LRU.Set("jason1", "T1")
	LRU.Set("jason2", "T2")
	LRU.Set("jason3", "T3")
	LRU.Set("jason4", "T4")
	fmt.Println("长度：", LRU.Size())
	fmt.Println(LRU.Get("jason1"))
	fmt.Println(LRU.Get("jason4"))
}
</pre>
###Golang实现长连接
<pre>
package main

/*
Golang实现长连接思路：
创建一个套接字对象, 指定其IP以及端口；
开始监听套接字指定的端口；
如有新的客户端连接请求, 则建立一个goroutine, 在goroutine中, 读取客户端消息, 并转发回去, 直到客户端断开连接；
主进程继续监听端口。
*/
import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println(string(message))
		msg := time.Now().String() + "\n"
		b := []byte(msg)
		conn.Write(b)
	}
}
</pre>
###Golang内存池
fork from https://github.com/wyh267
<pre>
package main

import (
	"fmt"
	"time"
)

var MAX_NODE_COUNT int = 1000

type Node struct {
	key   string
	value string
	hash1 uint64
	hash2 uint64
	Next  *Node
}

func NewNode() *Node {
	node := &Node{hash1: 0, hash2: 0, Next: nil}
	return node
}

// SMemPool 内存池对象
type SMemPool struct {
	nodechanGet   chan *Node
	nodechanGiven chan *Node
	nodeList      []Node
	freeList      []*Node
}

func (pool *SMemPool) makeNodeList() error {

	pool.nodeList = make([]Node, MAX_NODE_COUNT)

	return nil
}

func NewMemPool() *SMemPool {

	this := &SMemPool{nodechanGet: make(chan *Node, MAX_NODE_COUNT),
		nodechanGiven: make(chan *Node, MAX_NODE_COUNT),
		nodeList:      make([]Node, MAX_NODE_COUNT),
		freeList:      make([]*Node, 0)}

	return this
}

func (this *SMemPool) Alloc() *Node {

	return <-this.nodechanGet

}

func (this *SMemPool) Free(node *Node) error {

	this.nodechanGiven <- node

	return nil

}

func (this *SMemPool) InitMemPool() error {

	//初始化node

	go func() {
		//q := new(list.List)
		q := make([]Node, MAX_NODE_COUNT)
		for {

			if len(q) == 0 {
				q = append(q, make([]Node, MAX_NODE_COUNT)...)
			}
			e := q[0]
			timeout := time.NewTimer(time.Second)
			select {
			case b := <-this.nodechanGiven:
				timeout.Stop()
				fmt.Printf("Free Buffer...\n")
				//b=b[:MAX_DOCID_LEN]
				q = append(q, *b)
			case this.nodechanGet <- &e:
				timeout.Stop()
				q = q[1:]
				//fmt.Printf("Alloc Buffer...\n")
				//q.Remove(e)

			case <-timeout.C:
				fmt.Printf("remove Buffer...\n")

			}

		}

	}()
	return nil
}
func main() {
	mempool := NewMemPool()
	mempool.InitMemPool()

	for i := 0; i < 100; i++ {
		n := mempool.Alloc()
		mempool.Free(n)
	}
}
</pre>
###Mesos与K8S适用场景
- Mesos更适合做跨DC的资源管理，对于大数据领域或大量存在短任务，可以采用Mesos+上层调度器来解决大数据的资源池化调度问题；
- K8S更适合当应用的集群管理，它解决大规模应用部署的问题，而它的集群的热升级，动态伸缩，负载均衡，服务发现等特性可以让你的应用的更可靠。
###Golang  byte to string 
<pre>
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
</pre>
###Golang 简单回调函数
<pre>
package main 

import (
	"fmt"
)

type Callback func(str string)

func jasonFunc(str string){
	fmt.Println("callback started: "+str)
}

func TestCallBack(str string,callback Callback){
	if str ==""{
	callback(jasonFunc)
	}
	fmt.Println("not callback yet")
}
func main(){
	TestCallBack("",jasonFunc)
	TestCallBack("jason",jasonFunc)
}
</pre>
###Golang icmp协议
<pre>
package main

/*作用类似于Ping
主要用于在主机与路由器之间传递控制信息，包括报告错误、交换受限控制和状态信息等。当遇到IP数据无法访问目标、IP路由器无法按当前的传输速率转发数据包等情况时，会自动发送ICMP消息。ICMP报文在IP帧结构的首部协议类型字段（Protocol 8bit)的值=1.
*/
import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : ", os.Args[0], "host")
		os.Exit(0)
	}
	service := os.Args[1]

	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)

	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37
	len := 8
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)
	_, err = conn.Write(msg[0:len])
	checkError(err)

	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s\n", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}
</pre>
###Golang slice append&preappend
<pre>
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a,b,c"
	myslice := strings.Split(str, ",")
	myslice = append(myslice, "1")
	myslice = append([]string{"2"}, myslice...)
	fmt.Println(myslice)
}
</pre>
###Golang 网页生成图片
<pre>
package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

var root = flag.String("root", ".", "file system path")

func main() {
	http.HandleFunc("/blue/", blueHandler)
	http.HandleFunc("/red/", redHandler)
	http.Handle("/", http.FileServer(http.Dir(*root)))
	log.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func blueHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

func redHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{255, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	writeImageWithTemplate(w, &img)
}

var ImageTemplate string = `<!DOCTYPE html>
<html lang="en"><head></head>
<body><img src="data:image/jpg;base64,{{.Image}}"></body>`

// Writeimagewithtemplate encodes an image 'img' in jpeg format and writes it into ResponseWriter using a template.
func writeImageWithTemplate(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	if tmpl, err := template.New("image").Parse(ImageTemplate); err != nil {
		log.Println("unable to parse image template.")
	} else {
		data := map[string]interface{}{"Image": str}
		if err = tmpl.Execute(w, data); err != nil {
			log.Println("unable to execute template.")
		}
	}
}

// writeImage encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
</pre>
###简单的KV缓存
<pre>
package hicache

import "sync"

type Cache struct {
	hash map[string]interface{}
	lock sync.RWMutex
}

func New() *Cache {
	var c Cache
	c.hash = make(map[string]interface{})
	return &c
}

func (c *Cache) Set(k string, v interface{}) {
	c.lock.Lock()
	c.hash[k] = v
	c.lock.Unlock()
}

func (c *Cache) Get(k string) (interface{}, bool) {
	c.lock.RLock()
	v, ok := c.hash[k]
	c.lock.RUnlock()
	return v, ok
}

func (c *Cache) Count() int {
	return len(c.hash)
}

func (c *Cache) Del(k string) {
	c.lock.Lock()
	delete(c.hash, k)
	c.lock.Unlock()
}

func (c *Cache) Flush() {
	c.lock.Lock()
	c.hash = make(map[string]interface{})
	c.lock.Unlock()
}

func (c *Cache) Incr(k string, n int) int {
	v, ok := c.Get(k)
	new_v := n
	if ok {
		switch v.(type) {
		case int:
			new_v += v.(int)
		}
	}
	c.Set(k, new_v)
	return new_v
}
</pre>
###Golang 实现定时任务
<pre>
package main

import (
	"fmt"
	"time"
)

const INTERVAL_PERIOD time.Duration = 24 * time.Hour

const HOUR_TO_TICK int = 20
const MINUTE_TO_TICK int = 00
const SECOND_TO_TICK int = 00

func main() {
	ticker := updateTicker()
	for {
		<-ticker.C
		fmt.Println(time.Now(), "- just ticked")
		ticker = updateTicker()
	}
}

func updateTicker() *time.Ticker {
	nextTick := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
	if !nextTick.After(time.Now()) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	fmt.Println(nextTick, "- next tick")
	diff := nextTick.Sub(time.Now())
	return time.NewTicker(diff)
}
</pre>
###Golang 判断是否有汉字的方法 - unicode包
<pre>
package main

import (
	"fmt"
	"regexp"
	"unicode"
)

//判断是否有汉字的两个方法
func main() {
	str := "中文文文文"
	var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]{3,8}$")
	fmt.Println(hzRegexp.MatchString(str))
	fmt.Println(IsChineseChar(str))
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}
</pre>
###Golang 跳表 algorithm
<pre>
package main

import (
	"fmt"
	"math/rand"
)

const SKIPLIST_MAXLEVEL = 32 //8
const SKIPLIST_P = 4

type Node struct {
	Forward []Node
	Value   interface{}
}

func NewNode(v interface{}, level int) *Node {
	return &Node{Value: v, Forward: make([]Node, level)}
}

type SkipList struct {
	Header *Node
	Level  int
}

func NewSkipList() *SkipList {
	return &SkipList{Level: 1, Header: NewNode(0, SKIPLIST_MAXLEVEL)}
}

func (skipList *SkipList) Insert(key int) {

	update := make(map[int]*Node)
	node := skipList.Header

	for i := skipList.Level - 1; i >= 0; i-- {
		for {
			if node.Forward[i].Value != nil && node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
			} else {
				break
			}
		}
		update[i] = node
	}

	level := skipList.Random_level()
	if level > skipList.Level {
		for i := skipList.Level; i < level; i++ {
			update[i] = skipList.Header
		}
		skipList.Level = level
	}

	newNode := NewNode(key, level)

	for i := 0; i < level; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = *newNode
	}

}

func (skipList *SkipList) Random_level() int {

	level := 1
	for (rand.Int31()&0xFFFF)%SKIPLIST_P == 0 {
		level += 1
	}
	if level < SKIPLIST_MAXLEVEL {
		return level
	} else {
		return SKIPLIST_MAXLEVEL
	}
}

func (skipList *SkipList) PrintSkipList() {

	fmt.Println("\nSkipList-------------------------------------------")
	for i := SKIPLIST_MAXLEVEL - 1; i >= 0; i-- {

		fmt.Println("level:", i)
		node := skipList.Header.Forward[i]
		for {
			if node.Value != nil {
				fmt.Printf("%d ", node.Value.(int))
				node = node.Forward[i]
			} else {
				break
			}
		}
		fmt.Println("\n--------------------------------------------------------")
	} //end for

	fmt.Println("Current MaxLevel:", skipList.Level)
}

func (skipList *SkipList) Search(key int) *Node {

	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {

		fmt.Println("\n Search() Level=", i)
		for {
			if node.Forward[i].Value == nil {
				break
			}

			fmt.Printf("  %d ", node.Forward[i].Value)
			if node.Forward[i].Value.(int) == key {
				//fmt.Println("\nFound level=", i, " key=", key)
				return &node.Forward[i]
			}

			if node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				break
			}
		} //end for find

	} //end level
	return nil
}

func (skipList *SkipList) Remove(key int) {

	update := make(map[int]*Node)
	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {

		for {

			if node.Forward[i].Value == nil {
				break
			}

			if node.Forward[i].Value.(int) == key {
				fmt.Println("Remove() level=", i, " key=", key)
				update[i] = node
				break
			}

			if node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				break
			}

		} //end for find

	} //end level

	for i, v := range update {
		if v == skipList.Header {
			skipList.Level--
		}
		v.Forward[i] = v.Forward[i].Forward[i]
	}
}

func main() {
	skiplist := NewSkipList()
	for i := 0; i < 100; i++ {
		skiplist.Insert(i)
	}
	skiplist.PrintSkipList()
	skiplist.Search(78)
}
</pre>
###Golang 获取到今日的剩余时间
<pre>
func GetTodayLastSecond() time.Duration {
	today := GetToday() + " 23:59:59"
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", today, time.Local)
	return time.Duration(end.Unix()-time.Now().Local().Unix()) * time.Second
}
</pre>
###Golang amqp
公司大神写的Golang amqp连接工具 心跳重连
<pre>
package zcmmq

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

type Callback func(data []byte)

var (
	conn          *amqp.Connection
	channel       *amqp.Channel
	mqurl         string //链接字符串
	disconnet     bool   //是否失去连接
	isclose       bool
	receives      map[string]chan struct{}
	heartexchange = "heartbeat"
)

func Init(mq_url, log_db_url string) {
	isclose = false
	mqurl = mq_url
	err := connect()
	if err != nil {
		fmt.Println(mq_url, "can't connection to "+mq_url+" server.err:"+err.Error())
	} else {
		fmt.Println("连接成功")
	}
	receives = make(map[string]chan struct{})
	registerDataBase(log_db_url)
	go heartbeat()
}

func Push(exchange, routingkey string, data []byte) error {
	if channel == nil {
		err := connect()
		if err != nil {
			addErrorRecord(exchange, routingkey, string(data), err.Error())
			return err
		} else if channel == nil {
			addErrorRecord(exchange, routingkey, string(data), "channel/connection is not open")
			return errors.New("channel/connection is not open")
		}
	}

	fmt.Println("push", routingkey, exchange, string(data))
	err := channel.Publish(exchange, routingkey, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         data,
		DeliveryMode: 2,
	})
	if err != nil {
		failOnErr(err, "错误")
		addErrorRecord(exchange, routingkey, string(data), err.Error())
	}
	return err
}

func Receive(queueName string, callback Callback) {
	if channel == nil {
		connect()
	}
	if channel == nil {
		if receive1, ok := receives[queueName]; ok {
			<-receive1
			Receive(queueName, callback)
		}
		return
	}
	msgs, err := channel.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		failOnErr(err, "receive init err")
		return
	}
	receive := make(chan struct{})
	receives[queueName] = receive
	go func() {
		for d := range msgs {
			callback(d.Body)
		}
		disconnet = true
		if !isclose {
			//等待重连
			<-receive
			Receive(queueName, callback)
		} else {
			fmt.Println(queueName + " receive exit.")
		}
	}()
}

func connect() error {
	var err error
	conn, err = amqp.Dial(mqurl)
	if err != nil {
		failOnErr(err, "connecttion err")
		return err
	}
	channel, err = conn.Channel()
	if err != nil {
		failOnErr(err, "open channel err")
		return err
	}
	return nil
}

func Close() {
	isclose = true
	channel.Close()
	conn.Close()
}

func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

//心跳包 检测连接健康状况 自动重连
func heartbeat() {

	for {
		if channel == nil {
			connect()
		}
		if channel == nil {
			sleep()
			continue
		}
		err := channel.Publish(heartexchange, "heartbeat", false, false, amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte("heartbeat"),
			DeliveryMode: 2,
		})
		if err != nil {
			failOnErr(err, "send heartbeat err")
			connect()
			disconnet = true
		} else if disconnet {
			disconnet = false
			for _, r := range receives {
				r <- struct{}{}
			}
		}
		sleep()
	}
}

//心跳包休眠时间
func sleep() {
	time.Sleep(30 * time.Second)
}
</pre>
###Golang zero time 
<pre>
res,_ := time.ParseInLocation("0000-00-00 00:00:00", "0000-00-00 00:00:00", time.Local)
fmt.Println(res)
output==>
0000-01-01 00:00:00 +0800 CST
</pre>
###Golang 程序设计之用户登录状态监测
将下面的方法作为全局调用的基本方法
<pre>
//检查用户的登录信息
func (this *BaseController) CheckUsersLoginStatus() {
	account := this.GetSession("account")
	loginresult := make(map[string]interface{})
	defer func() {
		this.Data["json"] = loginresult
		this.ServeJSON()
	}()
	if account == nil {
		loginresult["islogin"] = false
	} else {
		loginresult["islogin"] = true
		loginresult["account"] = account.(string)
		//统计用户未读消息
		uid := this.GetSession("loginid").(int)
		uidstr := strconv.Itoa(uid)
		//统计方法
		loginresult["msg"] = 消息条数统计结果
	}
}
</pre>
###Golang net/url
<pre>
package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlparam = url.Values{}
	urlparam.Add("jason", "1")
	urlparam.Add("jason2", "2")
	urlparam.Add("jason3", "3")
	urlparam.Add("jason4", "4")
	var params = urlparam.Encode()
	fmt.Println("www.baidu.com?" + params)
}
output==>
www.baidu.com?jason=1&jason2=2&jason3=3&jason4=4
</pre>
###Golang 批量压缩图片与加水印
仅仅支持jpg格式
<pre>
package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd("data/")
	fmt.Println("OK!")
}

// 执行操作
func cmd(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("目录" + file.Name())
			cmd(path + file.Name() + "/")
		} else {
			if strings.Contains(strings.ToLower(file.Name()), ".jpg") {
				// 随机名称
				to := path + random_name() + ".jpg"

				origin := path + file.Name()

				fmt.Println("正在处理" + origin + ">>>" + to)

				cmd_resize(origin, 2048, 0, to)

				defer os.Remove(origin)
			}

		}
	}
}

// 改变大小
func cmd_resize(file string, width uint, height uint, to string) {
	// 打开图片并解码
	file_origin, _ := os.Open(file)
	origin, _ := jpeg.Decode(file_origin)
	defer file_origin.Close()

	canvas := resize.Resize(width, height, origin, resize.Lanczos3)

	file_out, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer file_out.Close()

	jpeg.Encode(file_out, canvas, &jpeg.Options{80})

	// cmd_watermark(to, strings.Replace(to, ".jpg", "@big.jpg", 1))
	cmd_thumbnail(to, 480, 360, strings.Replace(to, ".jpg", "@small.jpg", 1))
}

func cmd_thumbnail(file string, width uint, height uint, to string) {
	// 打开图片并解码
	file_origin, _ := os.Open(file)
	origin, _ := jpeg.Decode(file_origin)
	defer file_origin.Close()

	canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer file_out.Close()

	jpeg.Encode(file_out, canvas, &jpeg.Options{80})
}

// 水印
func cmd_watermark(file string, to string) {
	// 打开图片并解码
	file_origin, _ := os.Open(file)
	origin, _ := jpeg.Decode(file_origin)
	defer file_origin.Close()

	// 打开水印图并解码
	file_watermark, _ := os.Open("watermark.png")
	watermark, _ := png.Decode(file_watermark)
	defer file_watermark.Close()

	//原始图界限
	origin_size := origin.Bounds()

	//创建新图层
	canvas := image.NewNRGBA(origin_size)
	// 贴原始图
	draw.Draw(canvas, origin_size, origin, image.ZP, draw.Src)
	// 贴水印图
	draw.Draw(canvas, watermark.Bounds().Add(image.Pt(30, 30)), watermark, image.ZP, draw.Over)

	//生成新图片
	create_image, _ := os.Create(to)
	jpeg.Encode(create_image, canvas, &jpeg.Options{95})
	defer create_image.Close()
}
// 随机生成文件名
func random_name() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Int())
}
</pre>
###redis缓存
<pre>
keys * //看所有缓存key
flushdb //清除所有缓存
set authen_off 1 //设置不用验证身份
</pre>
###defer位置之坑
defer需要放到方法的最前面，不然调用不到！！！
<pre>
func test(){
	defer func(){
	...
	}()
	...
}
</pre>
###Linux 
<pre>
du -shx *   //列出所有文件夹硬盘占用|物理内存
df -Th       //显示文件总大小|物理内存、有没有已用空间非常大的tmpfs文件系统
free -m     //查看内存占用|虚拟内存   free -m -s 3 每3秒显示一次
ps -aux | sort -k4nr | head -5 //使用内存最多的5个进程
ps -aux | sort -k3nr | head -5 //使用CPU最多的5个进程
echo 1 > /proc/sys/vm/drop_caches   //清理内存命令|虚拟内存  echo 3 > ...  //清除所有缓存

vi /etc/sysconfig/iptables //编辑防火墙配置文件
service iptables restart   //最后重启防火墙使配置生效

/*
 buffers是为块设备设计的缓冲。比如磁盘读写，把分散的写操作集中进行，减少磁盘I/O，从而提高系统性能。
 cached是缓存读取过的内容，下次再读时，如果在缓存中命中，则直接从缓存读取，否则读取磁盘。由于缓存空间有限，过一段时间以后没用的缓存会被移动到swap里面，所以有时看到物理内存还有很多，swap就被利用了。
*/
　 系统
　　# uname -a # 查看内核/操作系统/CPU信息
　 　# head -n 1 /etc/issue # 查看操作系统版本
　　# cat /proc/cpuinfo # 查看CPU信息
　 　# hostname # 查看计算机名
　　# lspci -tv # 列出所有PCI设备
　　# lsusb -tv # 列出所有USB设备
　　# lsmod # 列出加载的内核模块
　　# env # 查看环境变量
　　资源
　 　# free -m # 查看内存使用量和交换区使用量
　　# df -h # 查看各分区使用情况
　　# du -sh # 查看指定目录的大小
　　# grep MemTotal /proc/meminfo # 查看内存总量
　　# grep MemFree /proc/meminfo # 查看空闲内存量
　　# uptime # 查看系统运行时间、用户数、负载
　　# cat /proc/loadavg # 查看系统负载
　　磁盘和分区
　　# mount | column -t # 查看挂接的分区状态
　　# fdisk -l # 查看所有分区
　　# swapon -s # 查看所有交换分区
　　# hdparm -i /dev/hda # 查看磁盘参数(仅适用于IDE设备)
　　# dmesg | grep IDE # 查看启动时IDE设备检测状况
　　网络
　　# ifconfig # 查看所有网络接口的属性
　　# iptables -L # 查看防火墙设置
　　# route -n # 查看路由表
　　# netstat -lntp # 查看所有监听端口
　　# netstat -antp # 查看所有已经建立的连接
　　# netstat -s # 查看网络统计信息
   用户
　　# w # 查看活动用户
　　# id # 查看指定用户信息
　　# last # 查看用户登录日志
　　# cut -d: -f1 /etc/passwd # 查看系统所有用户
　　# cut -d: -f1 /etc/group # 查看系统所有组
　　# crontab -l # 查看当前用户的计划任务
    # chmod +wx filename   filename目录增加权限给当前用户
	#chmod 777 filename 让所有用户对该目录【不涉及内部文件夹】有读写执行权限
	#chmod -R 777 filename 让所有用户可对该目录内所有的文件和文件夹及子文件夹具备读写执行的权限
	#useradd 添加用户
	#useradd -d /usr/sam -m sam 创建了一个用户sam,其中-d和-m选项用来为登录名sam产生一个主目录/usr/sam(/usr为默认的用户主目录所在的父目录)
	#userdel 删除用户
	#usermod 修改用户属性
	#passwd 选项 用户名 | -l 锁定口令，即禁用账号；-u  口令解锁；-d 使账号无口令；-f  强迫用户下次登录时修改口令
	#如果当前是超级用户，如果是超级用户，可以用下列形式指定任何用户的口令：passwd 任意用户的用户名
	#为用户指定空口令：passwd -d sam ，下次可以免密登录
	#groupadd  新建用户组
　　服务
　　# chkconfig --list # 列出所有系统服务
　  # chkconfig --list | grep on # 列出所有启动的系统服务
　　程序
　  # rpm -qa # 查看所有安装的软件包

sync //迫使缓冲块数据立即写盘，完成内存缓存区（buffers cache）有效数据向外设的存储
使用原理：
在linux系统中，为了加快数据的读取速度，默认情况下，某些数据将不会直接写入硬盘，而是先暂存内存中，如果一个数据被重复写，这样速度一定快，但存在一个问题，万一重新启动，或者是关机，或者是不正常断电的情况下，由于数据还没来得及存入硬盘，会造成数据更新不正常，这时需要命令sync进行数据的写入，即#sync，在内存中尚未更新的的数据会写入硬盘中。当然正常情况下，关闭系统时会自动进行内存数据于硬盘数据的同步检测，保证硬盘数据在关闭系统时是最新的。
swapoff -a;swapon -a  //清理swap可以一定程度上减轻系统卡顿问题，一般执行此命令之前先执行sync命令；
或者--->[注意：请选择业务低峰期进行操作] 关闭swap: date && swapoff -a &&   ;打开swap: swapon -a 
 ----------------------
 第一步：sync      执行命令，完成内存缓存区（buffers cache）有效数据向外设的存储;
 第二步：echo 3 > /proc/sys/vm/drop_caches 修改内核对内存的管理（主要是内存数据的清理）;
 第三步：free -m  查看内存失败结果；
 第四步：echo 0 > /proc/sys/vm/drop_caches 还原内核内存管理机制。
 ----------------------
-buffers/cache=used-buffers-cached，这个是应用程序真实使用的内存大小
+buffers/cache=free+buffers+cached，这个是服务器真实还可利用的内存大小
//-----清除ARP缓存
 arp -n|awk '/^[1-9]/ {print "arp -d "$1}' | sh
//----清除192.168.0.0网段的所有缓存
for((ip=2;ip<255;ip++));do arp -d 192.168.0.$ip &>/dev/null;done
注意：以上均需要root权限，尤其是最后一个，如果不再root下执行，则改为：
arp -n|awk '/^[1-9]/ {print "arp -d "$1}' | sudo sh
</pre>
###Mysql
<pre>
SHOW FULL PROCESSLIST; //展示所有链接到本数据库的所有进程
show status;
</pre>
###Golang hmac
<pre>
package main

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
)

//生成Hmac
func MakeHmac(str string) string {
	var hash = hmac.New(md5.New, []byte(YBPRIVATEKEYHMAC))
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func main(){
	fmt.Println(MakeHmac("str"))
}
</pre>
MYSQL :
ctrl + H  //历史日志  
###linux查看已连接的可用WIFI
<pre>
//列出可用的WiFi
nmcli dev wifi
//连接
nmcli dev wifi connect [SSID] password 
</pre>
MySQL中看一个字段a是否为空正确的判断方法是：
<pre>
a = "" or isnull(a)

----------端口访问权限设置----------
//只打开22端口
iptables -A INPUT -p tcp –dport 22 -j ACCEPT
iptables -A OUTPUT -p tcp –sport 22 -j ACCEPT
//保存设置
service iptables save
//重启服务
service iptables restart


或者是

nc -lp 22 &
----------------------
一、查看哪些端口被打开  netstat -anp
二、关闭端口号:iptables -A INPUT -p tcp --drop 端口号-j DROP
　　iptables -A OUTPUT -p tcp --dport 端口号-j DROP
三、打开端口号：iptables -A INPUT -ptcp --dport  端口号-j ACCEPT

//查看线程的栈跟踪
pstack <PID>  //
example: 
pstack 3114  //bash
#0  0x00007fb8661fe0ec in waitpid () from /lib64/libc.so.6
#1  0x00000000004406d4 in waitchld.isra.10 ()
#2  0x000000000044198c in wait_for ()
#3  0x00000000004337ee in execute_command_internal ()
#4  0x0000000000433a1e in execute_command ()
#5  0x000000000041e205 in reader_loop ()
#6  0x000000000041c88e in main ()

//以树状图显示进程间的关系
pstree    //有的需要下载

pstree -p <PID>

/*
strace常用来跟踪进程执行时的系统调用和所接收的信号。 在Linux世界，进程不能直接访问硬件设备，当进程需要访问硬件设备(比如读取磁盘文件，接收网络数据等等)时，必须由用户态模式切换至内核态模式，通 过系统调用访问硬件设备。strace可以跟踪到一个进程产生的系统调用,包括参数，返回值，执行消耗的时间。
*/
yum -y install strace 

strace   

example : strace cat /dev/null 

//对连接的IP按连接数量进行排序
netstat -ntu | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -n
//查看TCP连接状态并对每种状态进行统计
netstat -ant | awk '{print $NF}' | grep -v '[a-z]' | sort | uniq -c
//查看80端口连接数最多的20个IP
netstat -anlp|grep 80|grep tcp|awk '{print $5}'|awk -F: '{print $1}'|sort|uniq -c|sort -nr|head -n20
//查找较多time_wait连接
netstat -n|grep TIME_WAIT|awk '{print $5}'|sort|uniq -c|sort -rn|head -n20
//找查较多的SYN连接
netstat -an | grep SYN | awk '{print $5}' | awk -F: '{print $1}' | sort | uniq -c | sort -nr | more
</pre>
###Golang 时间 日期操作 week
<pre>
package main 

import (
	"fmt"
	"time"
)

func main(){
	var w = map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 7: "Sunday"}
	time_now := time.Now()
	week := time_now.Weekday().String()
	date_now := time_now.Format("2006-01-02")
	check11, _ := time.ParseInLocation("2006-01-02 15:04:05", date_now+" 00:00:00", time.Local)
	check12, _ := time.ParseInLocation("2006-01-02 15:04:05", date_now+" 06:00:00", time.Local)
	check21, _ := time.ParseInLocation("2006-01-02 15:04:05", date_now+" 18:00:00", time.Local)
	check22, _ := time.ParseInLocation("2006-01-02 15:04:05", date_now+" 23:59:59", time.Local)
	if week != w[6] {
		fmt.Println(week)
	}
	fmt.Println(check11, "\n", check12, "\n", check21, "\n", check22)
}
</pre>
###Golang 定时器timer优化
<pre>
package main

import (
	"os"
	"time"
)

func main() {
	c := make(chan int, 100)
	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
			time.Sleep(time.Second)
		}
		os.Exit(0)
	}()

	//方法一：只会产生一个定时器对象
	/*维护一个全局单一的定时器，每次操作前调整一下定时器的超时时间，
	从而避免每次循环都生成新的定时器对象*/
	timer := time.NewTimer(time.Second)
	for {
		timer.Reset(time.Second)
		select {
		case n := <-c:
			println(n)
		case <-timer.C:
		}
	}

	//方法二：会产生多个定时器对象，数量会根据i值而定，消耗性能，而且是常用做法
	for {
		select {
		case n := <-c:
			println(n)
		case <-timeAfter(time.Second * 2):
		}
	}
}

func timeAfter(d time.Duration) chan int {
	q := make(chan int, 1)
	time.AfterFunc(d, func() {
		q <- 1
		println("run")
	})
	return q
}
</pre>
###Golnag 去除byte、string空格函数
<pre>
package main

import (
	"bytes"
	"strings"
)

var newlineBytes = []byte{'\n'}

func TrimSpaceByte(src []byte) []byte {
	bytesArr := bytes.Split(src, newlineBytes)
	for i := 0; i < len(bytesArr); i++ {
		bytesArr[i] = bytes.TrimSpace(bytesArr[i])
	}

	return bytes.Join(bytesArr, nil)
}

func TrimSpaceString(src string) string {
	strs := strings.Split(src, "\n")
	for i := 0; i < len(strs); i++ {
		strs[i] = strings.TrimSpace(strs[i])
	}
	return strings.Join(strs, "")
}

func main() {
	str := "    -----Ki----"
	println(TrimSpaceString(str))
}
</pre>
###Golang XML、MAP互相转化
<pre>
package util

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
)

// DecodeXMLToMap decodes xml reading from io.Reader and returns the first-level sub-node key-value set,
// if the first-level sub-node contains child nodes, skip it.
func DecodeXMLToMap(r io.Reader) (m map[string]string, err error) {
	m = make(map[string]string)
	var (
		decoder = xml.NewDecoder(r)
		depth   = 0
		token   xml.Token
		key     string
		value   bytes.Buffer
	)
	for {
		token, err = decoder.Token()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}

		switch v := token.(type) {
		case xml.StartElement:
			depth++
			switch depth {
			case 2:
				key = v.Name.Local
				value.Reset()
			case 3:
				if err = decoder.Skip(); err != nil {
					return
				}
				depth--
				key = "" // key == "" indicates that the node with depth==2 has children
			}
		case xml.CharData:
			if depth == 2 && key != "" {
				value.Write(v)
			}
		case xml.EndElement:
			if depth == 2 && key != "" {
				m[key] = value.String()
			}
			depth--
		}
	}
}

// EncodeXMLFromMap encodes map[string]string to io.Writer with xml format.
//  NOTE: This function requires the rootname argument and the keys of m (type map[string]string) argument
//  are legitimate xml name string that does not contain the required escape character!
func EncodeXMLFromMap(w io.Writer, m map[string]string, rootname string) (err error) {
	switch v := w.(type) {
	case *bytes.Buffer:
		bufw := v
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return nil

	case *bufio.Writer:
		bufw := v
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return bufw.Flush()

	default:
		bufw := bufio.NewWriterSize(w, 256)
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return bufw.Flush()
	}
}
</pre>
###Golang check validate 
golang验证 手机 邮箱 中文名 昵称 
<pre>
package check

import (
	"regexp"
)

const (
	// 中国大陆手机号码正则匹配, 不是那么太精细
	// 只要是 13,14,15,17,18 开头的 11 位数字就认为是中国手机号
	chinaMobilePattern = `^1[34578][0-9]{9}$`
	// 用户昵称的正则匹配, 合法的字符有 0-9, A-Z, a-z, _, 汉字
	// 字符 '_' 只能出现在中间且不能重复, 如 "__"
	nicknamePattern = `^[a-z0-9A-Z\p{Han}]+(_[a-z0-9A-Z\p{Han}]+)*?$`
	// 用户名的正则匹配, 合法的字符有 0-9, A-Z, a-z, _
	// 第一个字母不能为 _, 0-9
	// 最后一个字母不能为 _, 且 _ 不能连续
	usernamePattern = `^[a-zA-Z][a-z0-9A-Z]*(_[a-z0-9A-Z]+)*?$`
	// 电子邮箱的正则匹配, 考虑到各个网站的 mail 要求不一样, 这里匹配比较宽松
	// 邮箱用户名可以包含 0-9, A-Z, a-z, -, _, .
	// 开头字母不能是 -, _, .
	// 结尾字母不能是 -, _, .
	// -, _, . 这三个连接字母任意两个不能连续, 如不能出现 --, __, .., -_, -., _.
	// 邮箱的域名可以包含 0-9, A-Z, a-z, -
	// 连接字符 - 只能出现在中间, 不能连续, 如不能 --
	// 支持多级域名, x@y.z, x@y.z.w, x@x.y.z.w.e
	mailPattern = `^[a-z0-9A-Z]+([\-_\.][a-z0-9A-Z]+)*@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)*?\.)+[a-zA-Z]{2,4}$`

	chineseNamePattern   = "^\\p{Han}+(\u00B7\\p{Han}+)*?$"
	chineseNameExPattern = "^\\p{Han}+([\u00B7\u2022\u2027\u30FB\u002E\u0387\u16EB\u2219\u22C5\uFF65\u05BC]\\p{Han}+)*?$"
)

var (
	chinaMobileRegexp   = regexp.MustCompile(chinaMobilePattern)
	nicknameRegexp      = regexp.MustCompile(nicknamePattern)
	usernameRegexp      = regexp.MustCompile(usernamePattern)
	mailRegexp          = regexp.MustCompile(mailPattern)
	chineseNameRegexp   = regexp.MustCompile(chineseNamePattern)
	chineseNameExRegexp = regexp.MustCompile(chineseNameExPattern)
)

// 检验是否为合法的中国手机号, 不是那么太精细
// 只要是 13,14,15,18 开头的 11 位数字就认为是中国手机号
func IsChinaMobile(b []byte) bool {
	if len(b) != 11 {
		return false
	}
	return chinaMobileRegexp.Match(b)
}

// 同 func IsChinaMobile(b []byte) bool
func IsChinaMobileString(s string) bool {
	if len(s) != 11 {
		return false
	}
	return chinaMobileRegexp.MatchString(s)
}

// 检验是否为合法的昵称, 合法的字符有 0-9, A-Z, a-z, _, 汉字
// 字符 '_' 只能出现在中间且不能重复, 如 "__"
func IsNickname(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	return nicknameRegexp.Match(b)
}

// 同 func IsNickname(b []byte) bool
func IsNicknameString(s string) bool {
	if len(s) == 0 {
		return false
	}
	return nicknameRegexp.MatchString(s)
}

// 检验是否为合法的用户名, 合法的字符有 0-9, A-Z, a-z, _
// 第一个字母不能为 _, 0-9
// 最后一个字母不能为 _, 且 _ 不能连续
func IsUserName(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	return usernameRegexp.Match(b)
}

// 同 func IsName(b []byte) bool
func IsUserNameString(s string) bool {
	if len(s) == 0 {
		return false
	}
	return usernameRegexp.MatchString(s)
}

// 检验是否为合法的电子邮箱, 考虑到各个网站的 mail 要求不一样, 这里匹配比较宽松
// 邮箱用户名可以包含 0-9, A-Z, a-z, -, _, .
// 开头字母不能是 -, _, .
// 结尾字母不能是 -, _, .
// -, _, . 这三个连接字母任意两个不能连续, 如不能出现 --, __, .., -_, -., _.
// 邮箱的域名可以包含 0-9, A-Z, a-z, -
// 连接字符 - 只能出现在中间, 不能连续, 如不能 --
// 支持多级域名, x@y.z, x@y.z.w, x@x.y.z.w.e
func IsMail(b []byte) bool {
	if len(b) < 6 { // x@x.xx
		return false
	}
	return mailRegexp.Match(b)
}

// 同 func IsMail(b []byte) bool
func IsMailString(s string) bool {
	if len(s) < 6 { // x@x.xx
		return false
	}
	return mailRegexp.MatchString(s)
}

// IsChineseName 检验是否为有效的中文姓名(比如 张三, 李四, 张三·李四)
func IsChineseName(b []byte) bool {
	return chineseNameRegexp.Match(b)
}

// 同 IsChineseName(b []byte) bool
func IsChineseNameString(s string) bool {
	return chineseNameRegexp.MatchString(s)
}

// IsChineseNameEx 检验是否为有效的中文姓名(比如 张三, 李四, 张三·李四),
// 主要功能和 IsChineseName 相同, 但是如果姓名中包含不规范的间隔符, 会自动修正为正确的间隔符 '\u00B7', 并返回正确的结果.
func IsChineseNameEx(b []byte) ([]byte, bool) {
	if chineseNameRegexp.Match(b) {
		return b, true
	}
	if !chineseNameExRegexp.Match(b) {
		return b, false
	}
	list := []rune(string(b))
	for i := 0; i < len(list); i++ {
		switch list[i] {
		case '\u2022', '\u2027', '\u30FB', '\u002E', '\u0387', '\u16EB', '\u2219', '\u22C5', '\uFF65', '\u05BC':
			list[i] = '\u00B7'
		}
	}
	return []byte(string(list)), true
}

// 同 IsChineseNameEx(b []byte) ([]byte, bool)
func IsChineseNameStringEx(s string) (string, bool) {
	if chineseNameRegexp.MatchString(s) {
		return s, true
	}
	if !chineseNameExRegexp.MatchString(s) {
		return s, false
	}
	list := []rune(s)
	for i := 0; i < len(list); i++ {
		switch list[i] {
		case '\u2022', '\u2027', '\u30FB', '\u002E', '\u0387', '\u16EB', '\u2219', '\u22C5', '\uFF65', '\u05BC':
			list[i] = '\u00B7'
		}
	}
	return string(list), true
}
</pre>
###Golang 获取本机ipv4列表
<pre>
package main

import (
	"net"
)

// IPv4List 获取本机的 ipv4 列表.
func IPv4List() ([]net.IP, error) {
	itfs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var (
		itf      net.Interface
		addrs    []net.Addr
		addr     net.Addr
		ipNet    *net.IPNet
		ok       bool
		ipv4     net.IP
		ipv4List []net.IP
	)
	for _, itf = range itfs {
		if itf.Flags&net.FlagUp == 0 {
			continue
		}
		addrs, err = itf.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr = range addrs {
			ipNet, ok = addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}
			ipv4 = ipNet.IP.To4()
			if ipv4 == nil {
				continue
			}
			ipv4List = append(ipv4List, ipv4)
		}
	}
	return ipv4List, nil
}

func main() {
	res, _ := IPv4List()
	for i := 0; i < len(res); i++ {
		println(res[i].String())
	}
}
</pre>
###Golang crypto/subtle
<pre>
package main

//安全比较string byte 的方法
import (
	"crypto/subtle"
)

func SecureCompareByte(given, actual []byte) bool {
	if subtle.ConstantTimeEq(int32(len(given)), int32(len(actual))) == 1 {
		if subtle.ConstantTimeCompare(given, actual) == 1 {
			return true
		}
		return false
	}
	// Securely compare actual to itself to keep constant time, but always return false
	if subtle.ConstantTimeCompare(actual, actual) == 1 {
		return false
	}
	return false
}

func SecureCompareString(given, actual string) bool {
	// The following code is incorrect:
	// return SecureCompare([]byte(given), []byte(actual))

	if subtle.ConstantTimeEq(int32(len(given)), int32(len(actual))) == 1 {
		if subtle.ConstantTimeCompare([]byte(given), []byte(actual)) == 1 {
			return true
		}
		return false
	}
	// Securely compare actual to itself to keep constant time, but always return false
	if subtle.ConstantTimeCompare([]byte(actual), []byte(actual)) == 1 {
		return false
	}
	return false
}

func main() {
	println(SecureCompareString("4554", "4554"))
}
</pre>
###Golang生成唯一token
<pre>
//生成密码aes密码 16位密码
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
</pre>
###Golang根据身份证获取生日
<pre>
// 根据身份证号获得出生日期
func GetBrithDate(idcard string) string {
	l := len(idcard)
	var s string
	if l == 15 {
		s = "19" + Substr(idcard, 6, 2) + "-" +Substr(idcard, 8, 2) + "-" + Substr(idcard, 10, 2)
		return s
	}
	if l == 18 {
		s = Substr(idcard, 6, 4) + "-" + Substr(idcard, 10, 2) + "-" + Substr(idcard, 12, 2)
		return 
	}
	return GetToday()
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

func GetToday() string {
	today := time.Now().Format("2006-01-02")
	return today
}
</pre>
###Golang简单的防止SQL注入
<pre>
func SqlDefend() error {
regexpstr := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(regexpstr)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	examplestr := "存在 sel44ectdeclare"
	return re.MatchString(examplestr) 
}
</pre>
###Golang 常用tips
<pre>
//加密电子邮箱为星号 将邮箱地址的部分变成*号
func EncryptEmail(email string) string {
	length := len(email)
	if 0 == length {
		return ""
	}
	strs := strings.Split(email, "@")
	return sToS(strs[0], '*', 1, len(strs[0])-2) + "@" + strs[1]
}

//加密姓名 将姓名的部分变成*号
func EncryptName(name string) string {
	r := []rune(name)
	r[0] = '*'
	return string(r)
}
</pre>
###Golang sync.Once
<pre>
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var one sync.Once //只会调用一次
	first := func() {
		fmt.Println("first times")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		j := i
		go func(int) { //int 是 j 的类型
			one.Do(first)
			fmt.Println(j)
			done <- true
		}(j)
	}
	<-done
	time.Sleep(2e9)
}
</pre>
###Golang http.PostForm
<pre>
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	m := url.Values{}
	m.Add("name1", "value1")
	m.Add("name2", "value2")
	resp, err := http.PostForm("https://www.baidu.com", m)
	if err != nil {
		fmt.Println("err")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err2")
	}
	fmt.Println(string(body))
}
</pre>
###Golang rsa RSA加密与解密方法
<pre>
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 加密
func RsaEncrypt(publicKey []byte, origData []byte) ([]byte, error) {
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
func RsaDecrypt(privateKey []byte, ciphertext []byte) ([]byte, error) {
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

func main() {
	res, _ := RsaEncrypt([]byte(`Hello`), []byte(`this is original data`))
	println("RSA-数据-" + string(res))
}
</pre>
###Golang通过身份证号码获取性别
<pre>
func WhichSexByIdcard(idcard string) string {
	var sexs = [2]string{"F", "M"} //或者 [2]string{"女","男"}
	length := len(idcard)
	if length == 18 {
		sex, _ := strconv.Atoi(string(idcard[16]))
		return sexs[sex%2]
	} else if length == 15 {
		sex, _ := strconv.Atoi(string(idcard[14]))
		return sexs[sex%2]
	}
	return "M"
}
</pre>
###Golang slice 通用插入方法
<pre>
package main

import (
	"fmt"
	"reflect"
)

//slice的通用插入方法
func Insert(slice interface{}, pos int, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	v = reflect.Append(v, reflect.ValueOf(value))
	reflect.Copy(v.Slice(pos+1, v.Len()), v.Slice(pos, v.Len()))
	v.Index(pos).Set(reflect.ValueOf(value))
	return v.Interface()
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(Insert(a, 3, 3))
}
</pre>
###Golang图片base64
<pre>
package main

import (
	"encoding/base64"
	//"io/ioutil"
	"os"
)

func main() {
	//读原图片
	ff, _ := os.Open("qr.png")
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩  sourcestring 图片base64数据
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	// //写入临时文件
	// ioutil.WriteFile("qr.png.txt", []byte(sourcestring), 0667)
	// //读取临时文件
	// cc, _ := ioutil.ReadFile("qr.png.txt")
	//解码
	dist, _ := base64.StdEncoding.DecodeString(sourcestring)
	//写入新文件，生成新图片
	f, _ := os.OpenFile("qb.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
}
</pre>
###Golang md5加密
<pre>
func MD5keyforreal(pars map[string]interface{}) string {
	md5key := "original_string"
	parsJson, _ := json.Marshal(pars) //转json
	sign_src := string(parsJson) + md5key
	h := md5.New()
	h.Write([]byte(sign_src))
	return hex.EncodeToString(h.Sum(nil))
}
</pre>
###MYSQL常用时间函数
<pre>
select date(NOW());
select year(NOW());
select month(NOW());
select day(NOW());
select hour(NOW());

2016-12-01
2016
12
1
12
</pre>
###Golang database/sql 使用官方包
<pre>
//声明一个全局的db对象，并进行初始化。
var db *sql.DB

func init() {
    db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/example?charset=utf8")
    db.SetMaxOpenConns(2000)
    db.SetMaxIdleConns(1000)
    db.SetConnMaxLifetime(300*time.Second)
    db.Ping()
}

//使用golang官方sql包查询表并且以字典形式输出结果
func SelectQuery(dbdriver, dbconnecturl, selectsqlquery string) map[interface{}]interface{} {
	db, err := sql.Open(dbdriver, dbconnecturl)
	//db.SetMaxIdleConns(N)设置最大空闲连接数
	//db.SetMaxOpenConns(N)设置最大打开连接数
	//db.SetConnMaxLifetime(d)超时时间
 	//example=>db.SetMaxOpenConns(2000)
    //example=>db.SetMaxIdleConns(1000)
	//通常，mysql的最大连接数默认是100, 最大可以达到16384。要考虑操作系统支持的最大并发线程数等等因素
	defer db.Close()
	rows, err := db.Query(selectsqlquery)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[interface{}]interface{})
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}
	return record
}

//使用golang官方的database/sql包 逐行打印
func SelectQueryByRow() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.221:3306)/zcm?charset=utf8&loc=Asia%2FShanghai")
	defer db.Close()
	// 执行sql语句
	rows, err := db.Query("SELECT * FROM config limit 3")
	if err != nil {
		panic(err.Error())
	}
	// 获取一条数据中所有列的名称
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	// 生成一个切片来接收值
	values := make([]sql.RawBytes, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var datacount int
	// 获取所有行
	for rows.Next() {
		datacount++
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		// 打印出每列信息(string)
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ":", value) //每条数据的 字段名:值
		}
		fmt.Println("***************************************************")
	}
	fmt.Println(datacount) //数据总条数
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}

//mysql 事务
tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}
defer tx.Rollback()
stmt, err := tx.Prepare("INSERT INTO foo VALUES (?)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close() // danger!
for i := 0; i < 10; i++ {
    _, err = stmt.Exec(i)
    if err != nil {
        log.Fatal(err)
    }
}
err = tx.Commit()
if err != nil {
    log.Fatal(err)
}
</pre>
###Golang pool 临时对象池   GC
<pre>
package main

import (
	"log"
	//"runtime"
	"sync"
)

func main() {
	//临时对象池 ，GC会销毁所有值
	var pool = &sync.Pool{New: func() interface{} { return "This is pool auto reproduce an init value" }}
	val1 := "Hello World--1"
	val2 := "Hello World--2"
	//放入2个
	pool.Put(val1)
	pool.Put(val2)
	//runtime.GC()  //WARNING : GC一旦被调用，pool值全部消失...
	//取出4个，后2个自动调用New方法产生
	log.Println(pool.Get())
	log.Println(pool.Get())
	log.Println(pool.Get())
	log.Println(pool.Get())
}
</pre>
###Golang 工作线程工具 workqueue
<pre>
package workqueue

type Queue struct {
	Jobs    chan string
	done    chan bool
	workers chan chan int
}

func (q *Queue) worker(id int, callback func(string, int)) (done chan int) {
	done = make(chan int)

	go func() {
	work:
		for {
			select {
			case <-q.done:
				break work
			case j := <-q.Jobs:
				callback(j, id)
			}
		}
		done <- id
		close(done)
	}()
	return done
}

func (q *Queue) Init(size int, workers int, callback func(string, int)) {

	q.Jobs = make(chan string, size)
	q.done = make(chan bool)
	q.workers = make(chan chan int, workers)
	for w := 1; w <= workers; w++ {
		q.workers <- q.worker(w, callback)
	}
	close(q.workers)
}

func (q *Queue) Run() {

	// Wait for workers to be halted
	for w := range q.workers {
		<-w
	}
	// Nothing should still be mindlessly adding jobs
	close(q.Jobs)
}

// Allow the queueue to be drained after it is closed
func (q *Queue) Drain(callback func(string)) {
	for j := range q.Jobs {
		callback(j)
	}
}

func (q *Queue) Close() {
	close(q.done)
}

/*------------------------使用-----------------------------*/

// A real worker would be parsing a web page or crunching numbers
func workerFunc(job string, workerId int) {
	fmt.Println("worker", workerId, "processing job", job)
	//time.Sleep(1 * time.Second)
	//fmt.Println("worker", workerId, "saving job", job)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	jobQueueSize := 100                 //队列容量
	numberOfWorkers := runtime.NumCPU() //工作线程数

	queue := goworkqueue.Queue{}
	queue.Init(jobQueueSize, numberOfWorkers, workerFunc)

	// Pretend we suddenly need to stop the workers.
	// This might be a SIGTERM or perhaps the workerFunc() called queue.Close()
	go func() {
		time.Sleep(1 * time.Second)
		queue.Close()
		fmt.Println("ABORT!")
	}()

	// We can optionally prefill the work queue
	for j := 0; j <= 9999; j++ {
		queue.Jobs <- fmt.Sprintf("Job %d", j)
	}

	// Blocks until queue.Close()
	queue.Run()

	// Optional, callback for emptying the queue *if* anything remains
	queue.Drain(func(job string) {
		fmt.Printf("'%s' wasn't finished\n", job)
	})
}
</pre>
###Beego MyAutoRouter UrlMapping 自动路由 性能会有所改善？
<pre>
router:
func init() {
//=======================================解决自动路由bug====================================
	exceptMethod := []string{"GetInt8", "GetInt16", "GetInt32", "GetInt64", "GetFiles",
		"XSRFToken", "CheckXSRFCookie", "XSRFFormHTML", "HandlerFunc", "Mapping", "URLMapping", "renderTemplate", "URLFor"}
	for _, fn := range exceptMethod {
		beego.ExceptMethodAppend(fn)
	}

	beego.MyAutoRouter(&controller_path_dirname.Controller_name{})
}



Controller_name:
func (c *Controller_name) URLMapping() {
	c.Mapping("examplle", c.example)     
}
</pre>
###Beego 获取config信息 app.conf
<pre>
func init() {
	DEBUG, _ = beego.AppConfig.Bool("debug")
	var config map[string]string
	var err error
	if DEBUG {
		config, err = beego.AppConfig.GetSection("debug")
	} else {
		config, err = beego.AppConfig.GetSection("release")
	}
	beego.Info("Release模式:", !DEBUG)
	if err != nil {
		panic(errors.New("配置文件读取错误 " + err.Error()))
	}
	REDIS_URI = config["redis_url"]
	MYSQL_URI = config["mysql_url"]
	BEEGO_CACHE = config["beego_cache"]
	Rc, Re = cache.NewCache("redis", BEEGO_CACHE)
}
</pre>
###Golang 给数值加上分隔逗号
<pre>
//给数值类型加上逗号,
func AddDouhao(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".")
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".")
}
func main(){
	fmt.Println(AddDouhao("566768778"))
}
</pre>
###Golang pdf 方法
<pre>
package pdf

import (
	"errors"
	"github.com/signintech/gopdf"
	"log"
)

type TASKPdfConfig struct {
	PAGEWIDE      float64
	PAGEHEIGHT    float64
	SIDE          float64
	TOP           float64
	BOTTOM        float64
	TABLELINESIZE float64
	LINESIZE      float64
}

type TASKPdf struct {
	gopdf.GoPdf
	Config *TASKPdfConfig
}

const FONTNAME = "TASK"                   //字体文件名称
const FONTPATH = "./pdf/ttf/SIMYOU.TTF"   //字体文件位置
const STAMPPATH = "./pdf/ttf/example.jpg" //叠加图片位置
const DEFAULTFONTSIZE = 10                //中文字，高和宽等于这个值，英文为一半
// const PAGEWIDE = 595.28    //595.28, 841.89 = A4
// const PAGEHEIGHT = 841.89
// const SIDE = 50
// const TOP = 90
// const BOTTOM = 90

var (
	defaultBr          float64
	defaultLineHight   float64
	defaultLineFontNum int
)

func NewTASKPdf(config *TASKPdfConfig) *TASKPdf {
	if config == nil {
		config = &TASKPdfConfig{PAGEWIDE: 595.28, PAGEHEIGHT: 841.89, SIDE: 50, TOP: 90, BOTTOM: 90, LINESIZE: 1, TABLELINESIZE: 0.5}
	}
	defaultLineHight = float64(float64(DEFAULTFONTSIZE) + float64(DEFAULTFONTSIZE)/5)
	defaultBr = float64(float64(DEFAULTFONTSIZE) + float64(DEFAULTFONTSIZE)/2)
	num := (config.PAGEWIDE - config.SIDE*2) / DEFAULTFONTSIZE
	defaultLineFontNum = int(num)
	pdf := &TASKPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: config.PAGEWIDE, H: config.PAGEHEIGHT}})
	pdf.SetTopMargin(config.TOP)
	pdf.SetLeftMargin(config.SIDE)
	pdf.SetLineWidth(config.SIDE)
	pdf.Config = config
	//设置默认字体
	var err error
	err = pdf.AddTTFFont(FONTNAME, FONTPATH)
	HandleError("load font error", err)
	err = pdf.SetFont(FONTNAME, "", DEFAULTFONTSIZE)
	HandleError("set font error", err)
	//设置线
	pdf.SetLineWidth(1)

	return pdf
}

func (pdf *TASKPdf) DoBR(height float64) *TASKPdf {
	if pdf.GetY() > pdf.Config.PAGEHEIGHT-pdf.Config.BOTTOM-DEFAULTFONTSIZE {
		pdf.AddPage()
	}
	pdf.Br(height)
	return pdf
}

func (pdf *TASKPdf) DoDefaultBR() *TASKPdf {
	if pdf.GetY() > pdf.Config.PAGEHEIGHT-pdf.Config.BOTTOM-DEFAULTFONTSIZE {
		pdf.AddPage()
	}
	pdf.Br(defaultBr)
	return pdf
}

func (pdf *TASKPdf) WriteDefaultLine() *TASKPdf {
	x, y := pdf.GetX(), pdf.GetY()
	// pdf.SetLineWidth(pdf.Config.TABLELINESIZE)
	pdf.Line(x, y, pdf.Config.PAGEWIDE-pdf.Config.SIDE, y)
	// pdf.SetY(y + 5)
	// pdf.SetLineWidth(pdf.Config.LINESIZE)
	return pdf
}
func (pdf *TASKPdf) WriteVerticalLine(height float64) *TASKPdf {
	x, y := pdf.GetX(), pdf.GetY()
	// pdf.SetLineWidth(pdf.Config.TABLELINESIZE)
	pdf.Line(x, y, x, y+height)
	// pdf.SetLineWidth(pdf.Config.LINESIZE)
	return pdf
}
func (pdf *TASKPdf) WriteRightVerticalLine(height float64) *TASKPdf {
	y := pdf.GetY()
	// pdf.SetLineWidth(pdf.Config.TABLELINESIZE)
	pdf.Line(pdf.Config.PAGEWIDE-pdf.Config.SIDE, y, pdf.Config.PAGEWIDE-pdf.Config.SIDE, y+height)
	// pdf.SetLineWidth(pdf.Config.LINESIZE)
	return pdf
}
func (pdf *TASKPdf) ResetX() *TASKPdf {
	pdf.SetX(pdf.Config.SIDE)
	return pdf
}

func (pdf *TASKPdf) AddY(size float64) *TASKPdf {
	y := pdf.GetY()
	pdf.SetY(y + size)
	return pdf
}
func (pdf *TASKPdf) AddX(size float64) *TASKPdf {
	x := pdf.GetX()
	pdf.SetX(x + size)
	return pdf
}
func (pdf *TASKPdf) WriteInCenter(text string, fontSize int) *TASKPdf {
	//1.设置字体
	pdf.SetFontSize(fontSize)
	wide, _ := pdf.MeasureTextWidth(text)
	//2.居中定位,写
	// offset := (PAGEWIDE - pdf.GetX() - wide) / 2
	offset := (pdf.Config.PAGEWIDE - wide) / 2
	pdf.SetX(offset)
	pdf.Cell(nil, text)
	pdf.DoBR(GetBR(fontSize))
	if fontSize != DEFAULTFONTSIZE {
		//3.字体设置回默认的
		pdf.SetDefaultFontSize()
	}
	return pdf
}

//右对齐
func (pdf *TASKPdf) WriteInRight(text string) *TASKPdf {

	wide, _ := pdf.MeasureTextWidth(text)
	//2.居中定位,写
	// offset := (PAGEWIDE - pdf.GetX() - wide) / 2
	offset := pdf.Config.PAGEWIDE - pdf.Config.SIDE - wide
	pdf.SetX(offset)
	pdf.Cell(nil, text)
	pdf.DoBR(GetBR(DEFAULTFONTSIZE))
	return pdf
}

func (pdf *TASKPdf) WriteAnyPlace(percent float64, text string, fontSize int) *TASKPdf {
	if fontSize == DEFAULTFONTSIZE {
		pdf.SetX(pdf.Config.PAGEWIDE * percent)
		pdf.Cell(nil, text)
	} else {
		pdf.SetFontSize(fontSize)

		pdf.SetX(pdf.Config.PAGEWIDE * percent)
		pdf.Cell(nil, text)

		pdf.SetDefaultFontSize()
	}
	return pdf
}

func (pdf *TASKPdf) WriteWithLine(text string) *TASKPdf {
	x, y := pdf.GetX(), pdf.GetY()
	// log.Println("x,y:", x, y)
	pdf.Cell(nil, text)
	wide, _ := pdf.MeasureTextWidth(text)
	pdf.Line(x, y+defaultLineHight, x+wide, y+defaultLineHight)
	// log.Println("wide:", wide)
	// log.Println(x, y+defaultLineHight, x+wide, y+defaultLineHight)
	//划线后x轴值不是线的末尾，设置回去
	pdf.SetX(x + wide)
	return pdf
}

func (pdf *TASKPdf) WritePassage(text string) *TASKPdf {
	length := Strlen(text)
	wideFlag := pdf.Config.PAGEWIDE - pdf.Config.SIDE*2
	var i int
	var offset int
	var flag bool
	for i = 0; i <= length-defaultLineFontNum; i += defaultLineFontNum {
		offset = 0
		flag = true
		for flag {
			if i+defaultLineFontNum+offset > length {
				flag = false
			}
			strline := Substr(text, i, defaultLineFontNum+offset)
			wide, _ := pdf.MeasureTextWidth(strline)
			if wide < wideFlag {
				offset++
			} else {
				flag = false
			}
		}
		pdf.Write(Substr(text, i, defaultLineFontNum+offset))
		pdf.DoDefaultBR()
		i += offset
	}
	if i < length {
		pdf.Write(Substr(text, i, length-i))
	}
	return pdf
}

func (pdf *TASKPdf) Write(text string) *TASKPdf {
	pdf.Cell(nil, text)
	return pdf
}

func (pdf *TASKPdf) ZAddPage() *TASKPdf {
	pdf.AddPage()
	return pdf
	// pdf.ZImage(BACKGROUNDPATH, 35, 0)
	// log.Println(pdf)
	// pdf.Image(BACKGROUNDPATH, 35, 0, nil)
}
func (pdf *TASKPdf) ZImage(image string, x float64, y float64) *TASKPdf {
	pdf.Image(image, x, y, nil)
	return pdf
}
func (pdf *TASKPdf) Out(filename string) *TASKPdf {
	pdf.WritePdf(filename)
	return pdf
}

func (pdf *TASKPdf) WriteWithColor(text string, r uint8, g uint8, b uint8) *TASKPdf {
	pdf.SetTextColor(r, g, b)
	pdf.Cell(nil, text)
	//颜色设置回去
	pdf.SetTextColor(0, 0, 0)
	pdf.SetGrayFill(0)
	return pdf
}

func (pdf *TASKPdf) SetFontSize(fontSize int) *TASKPdf {
	err := pdf.SetFont(FONTNAME, "", fontSize)
	HandleError("set font error", err)
	return pdf
}

func (pdf *TASKPdf) SetDefaultFontSize() *TASKPdf {
	err := pdf.SetFont(FONTNAME, "", DEFAULTFONTSIZE)
	HandleError("set font error", err)
	return pdf
}

func HandleError(prefix string, err error) bool {
	if err != nil {
		log.Fatalln(err.Error())
		panic(errors.New(prefix))
		return true
	} else {
		return false
	}
}

func (pdf *TASKPdf) DrawCommonForm(x []float64, y []float64, value [][]string) *TASKPdf {
	lenX := len(x)
	lenY := len(y)
	//画横线
	for i := 0; i < lenY; i++ {
		pdf.Line(x[0], y[i], x[lenX-1], y[i])
	}
	//画竖线
	for i := 0; i < lenX; i++ {
		pdf.Line(x[i], y[0], x[i], y[lenY-1])
	}
	//填值，暂不支持居中
	for i := 0; i < lenY-1; i++ {
		for j := 0; j < lenX-1; j++ {
			pdf.SetX(x[j] + 5)
			pdf.SetY(y[i] + 5)
			pdf.Cell(nil, value[i][j])
		}

	}
	return pdf
}

//==============test=================================
//详见func (pdf *TASKPdf) WriteWithLine
func (pdf *TASKPdf) WriteWithLineFont(text string, font int) *TASKPdf {
	x, y := pdf.GetX(), pdf.GetY()

	pdf.Cell(nil, text)
	wide, _ := pdf.MeasureTextWidth(text)
	lineHight := float64(float64(font) + float64(font)/5)
	pdf.Line(x, y+lineHight, x+wide, y+lineHight)

	pdf.SetX(x + wide)
	return pdf
}

//在固定长度length里写text，起始位置为start
func (pdf *TASKPdf) WriteInFixedlength(text string, length int, start int) *TASKPdf {
	//去除逻辑错误
	if length < 0 {
		length = 10
	}
	if start > length || start < 0 {
		start = 0
	}
	//获取视觉长度
	lentext := LenOfSee(text)
	if lentext+start > length {
		text = GetBlank(start) + text[:length-start]
	} else {
		text = GetBlank(start) + text + GetBlank(length-(lentext+start))
	}

	return pdf.Write(text)
}

//固定长度中间写划线
func (pdf *TASKPdf) WriteInCenterWithLineFixedlength(text string, length int, font int) *TASKPdf {
	lentext := LenOfSee(text)
	//去除逻辑错误
	if length < 1 {
		length = 1
	}
	if font < 0 {
		font = 10
	}
	if lentext > length {
		text = text[:length]
	} else {
		//前后缀空格
		prefix := GetBlank((length - lentext) / 2)
		suffix := GetBlank((length - lentext) - (length-lentext)/2)
		text = prefix + text + suffix
	}
	return pdf.WriteWithLineFont(text, font)
}

//
func (pdf *TASKPdf) WriteInCenterFixedlength(text string, length int) *TASKPdf {
	lentext := LenOfSee(text)
	//去除逻辑错误
	if length < 1 {
		length = 1
	}
	if lentext > length {
		text = text[:length]
		lentext = length
	} else {
		//前后缀空格
		prefix := GetBlank((length - lentext) / 2)
		suffix := GetBlank((length - lentext) - (length-lentext)/2)
		text = prefix + text + suffix
	}
	return pdf.Write(text)
}

//自动分段(自定宽度，前缀长度，行间距，所有内容)
func (pdf *TASKPdf) WritePassageAnyWidth(wideFlag int, preLength float64, lineSpace int, text string) *TASKPdf {
	r := []rune(text)
	length := len(r)
	if length < wideFlag {
		return pdf.AddX(preLength).Write(text).DoBR(GetBR(lineSpace))
	}
	var i int

	for length > wideFlag {
		length -= wideFlag
		pdf.AddX(preLength).Write(string(r[i : i+wideFlag])).DoBR(GetBR(lineSpace))
		i += wideFlag
	}
	return pdf.AddX(preLength).Write(string(r[i:])).DoBR(GetBR(lineSpace))
}

//------common方法---------
func LenOfSee(str string) int {
	rs := []rune(str)
	return (len(rs) + len(str)) / 2
}

//获取i长度的空格串
func GetBlank(i int) string {
	blank := ""
	if i == 0 {
		return blank
	}
	shi := i / 10
	ge := i % 10
	if shi > 0 {
		for num := 0; num < shi; num++ {
			blank += "          "
		}
	}
	if ge > 0 {
		for num := 0; num < ge; num++ {
			blank += " "
		}
	}

	return blank
}

func GetBR(fontSize int) float64 {
	return float64(float64(fontSize) + float64(fontSize)/2)
}

func Strlen(s string) int {
	rs := []rune(s)
	rl := len(rs)
	return rl
}
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
</pre>
###Golang 发送带有附件的邮件
<pre>
package main

import (
	"github.com/scorredoira/email"
	"log"
	"net/mail"
	"net/smtp"
)

func main() {
	m := email.NewMessage("This is Title", "this is content")
	m.From = mail.Address{Name: "Jason", Address: "example@qq.com"}
	m.To = []string{"example@163.com"}
	err := m.Attach("qr.png")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("附件添加成功！")
	}
	err = email.Send("smtp.qq.com:25", smtp.PlainAuth("", "example@qq.com", "pwd", "smtp.qq.com"), m)
	if err == nil {
		log.Println("邮件发送成功！")
	} else {
		log.Println(err.Error())
	}
}
</pre>
###Golang 断点续传 
<pre>
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	// 最大线程数量
	MaxThread = 5
	// 缓冲区大小
	CacheSize = 1024
)

// 创建新的文件下载
//
// 如果 size <= 0 则自动获取文件大小
func NewFileDl(url string, file *os.File, size int64) (*FileDl, error) {
	if size <= 0 {
		// 获取文件信息
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		size = resp.ContentLength
	}

	f := &FileDl{
		Url:  url,
		Size: size,
		File: file,
	}

	return f, nil
}

type FileDl struct {
	Url  string   // 下载地址
	Size int64    // 文件大小
	File *os.File // 要写入的文件

	BlockList []Block // 用于记录未下载的文件块起始位置

	onStart  func()
	onPause  func()
	onResume func()
	onDelete func()
	onFinish func()
	onError  func(int, error)

	paused bool
	status Status
}

// 开始下载
func (f *FileDl) Start() {
	go func() {
		if f.Size <= 0 {
			f.BlockList = append(f.BlockList, Block{0, -1})
		} else {
			blockSize := f.Size / int64(MaxThread)
			var begin int64
			// 数据平均分配给各个线程
			for i := 0; i < MaxThread; i++ {
				var end = (int64(i) + 1) * blockSize
				f.BlockList = append(f.BlockList, Block{begin, end})
				begin = end + 1
			}
			// 将余出数据分配给最后一个线程
			f.BlockList[MaxThread-1].End += f.Size - f.BlockList[MaxThread-1].End
		}

		f.touch(f.onStart)
		// 开始下载
		err := f.download()
		if err != nil {
			f.touchOnError(0, err)
			return
		}
	}()
}

func (f *FileDl) download() error {
	f.startGetSpeeds()

	ok := make(chan bool, MaxThread)
	for i := range f.BlockList {
		go func(id int) {
			defer func() {
				ok <- true
			}()

			for {
				err := f.downloadBlock(id)
				if err != nil {
					f.touchOnError(0, err)
					// 重新下载
					continue
				}
				break
			}
		}(i)
	}

	for i := 0; i < MaxThread; i++ {
		<-ok
	}
	// 检查是否为暂停导致的“下载完成”
	if f.paused {
		f.touch(f.onPause)
		return nil
	}
	f.paused = true
	f.touch(f.onFinish)

	return nil
}

// 文件块下载器
// 根据线程ID获取下载块的起始位置
func (f *FileDl) downloadBlock(id int) error {
	request, err := http.NewRequest("GET", f.Url, nil)
	if err != nil {
		return err
	}
	begin := f.BlockList[id].Begin
	end := f.BlockList[id].End
	if end != -1 {
		request.Header.Set(
			"Range",
			"bytes="+strconv.FormatInt(begin, 10)+"-"+strconv.FormatInt(end, 10),
		)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var buf = make([]byte, CacheSize)
	for {
		if f.paused == true {
			// 下载暂停
			return nil
		}

		n, e := resp.Body.Read(buf)

		bufSize := int64(len(buf[:n]))
		if end != -1 {
			// 检查下载的大小是否超出需要下载的大小
			// 这里End+1是因为http的Range的end是包括在需要下载的数据内的
			// 比如 0-1 的长度其实是2，所以这里end需要+1
			needSize := f.BlockList[id].End + 1 - f.BlockList[id].Begin
			if bufSize > needSize {
				// 数据大小不正常
				// 一般是因为网络环境不好导致
				// 比如用中国电信下载国外文件

				// 设置数据大小来去掉多余数据
				// 并结束这个线程的下载
				bufSize = needSize
				n = int(needSize)
				e = io.EOF
			}
		}
		// 将缓冲数据写入硬盘
		f.File.WriteAt(buf[:n], f.BlockList[id].Begin)

		// 更新已下载大小
		f.status.Downloaded += bufSize
		f.BlockList[id].Begin += bufSize

		if e != nil {
			if e == io.EOF {
				// 数据已经下载完毕
				return nil
			}
			return e
		}
	}

	return nil
}

func (f *FileDl) startGetSpeeds() {
	go func() {
		var old = f.status.Downloaded
		for {
			if f.paused {
				f.status.Speeds = 0
				return
			}
			time.Sleep(time.Second * 1)
			f.status.Speeds = f.status.Downloaded - old
			old = f.status.Downloaded
		}
	}()
}

// 获取下载统计信息
func (f FileDl) GetStatus() Status {
	return f.status
}

// 暂停下载
func (f *FileDl) Pause() {
	f.paused = true
}

// 继续下载
func (f *FileDl) Resume() {
	f.paused = false
	go func() {
		if f.BlockList == nil {
			f.touchOnError(0, errors.New("BlockList == nil, can not get block info"))
			return
		}

		f.touch(f.onResume)
		err := f.download()
		if err != nil {
			f.touchOnError(0, err)
			return
		}
	}()
}

// 任务开始时触发的事件
func (f *FileDl) OnStart(fn func()) {
	f.onStart = fn
}

// 任务暂停时触发的事件
func (f *FileDl) OnPause(fn func()) {
	f.onPause = fn
}

// 任务继续时触发的事件
func (f *FileDl) OnResume(fn func()) {
	f.onResume = fn
}

// 任务完成时触发的事件
func (f *FileDl) OnFinish(fn func()) {
	f.onFinish = fn
}

// 任务出错时触发的事件
//
// errCode为错误码，errStr为错误描述
func (f *FileDl) OnError(fn func(int, error)) {
	f.onError = fn
}

// 用于触发事件
func (f FileDl) touch(fn func()) {
	if fn != nil {
		go fn()
	}
}

// 触发Error事件
func (f FileDl) touchOnError(errCode int, err error) {
	if f.onError != nil {
		go f.onError(errCode, err)
	}
}

type Status struct {
	Downloaded int64
	Speeds     int64
}

type Block struct {
	Begin int64 `json:"begin"`
	End   int64 `json:"end"`
}

//断点续传
func main() {
	filename := "http://packages.linuxdeepin.com/ubuntu/dists/devel/main/binary-amd64/Packages.gz"
	index := strings.LastIndex(filename, ".")
	downloadname := filename[index+1:]
	file, err := os.Create("./tmpfile" + "." + downloadname)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	fileDl, err := NewFileDl(filename, file, -1)
	if err != nil {
		log.Println(err)
	}

	var exit = make(chan bool)
	var resume = make(chan bool)
	var pause bool
	var wg sync.WaitGroup
	wg.Add(1)
	fileDl.OnStart(func() {
		fmt.Println("download started")
		format := "\033[2K\r%v/%v [%s] %v byte/s %v"
		for {
			status := fileDl.GetStatus()
			var i = float64(status.Downloaded) / float64(fileDl.Size) * 50
			h := strings.Repeat("=", int(i)) + strings.Repeat(" ", 50-int(i))

			select {
			case <-exit:
				fmt.Printf(format, status.Downloaded, fileDl.Size, h, 0, "[FINISH]")
				fmt.Println("\ndownload finished")
				wg.Done()
			default:
				if !pause {
					time.Sleep(time.Second * 1)
					fmt.Printf(format, status.Downloaded, fileDl.Size, h, status.Speeds, "[DOWNLOADING]")
					os.Stdout.Sync()
				} else {
					fmt.Printf(format, status.Downloaded, fileDl.Size, h, 0, "[PAUSE]")
					os.Stdout.Sync()
					<-resume
					pause = false
				}
			}
		}
	})

	fileDl.OnPause(func() {
		pause = true
	})

	fileDl.OnResume(func() {
		resume <- true
	})

	fileDl.OnFinish(func() {
		exit <- true
	})

	fileDl.OnError(func(errCode int, err error) {
		log.Println(errCode, err)
	})

	fmt.Printf("%+v\n", fileDl)

	fileDl.Start()
	time.Sleep(time.Second * 2)
	fileDl.Pause()
	time.Sleep(time.Second * 3)
	fileDl.Resume()
	wg.Wait()
}
</pre>
###Golang 非线程安全的优先级队列
<pre>
package main

import (
	"container/heap"
	"fmt"
)

type Interface interface {
	Less(other interface{}) bool
}

type sorter []Interface

// Implement heap.Interface: Push, Pop, Len, Less, Swap
func (s *sorter) Push(x interface{}) {
	*s = append(*s, x.(Interface))
}

func (s *sorter) Pop() interface{} {
	n := len(*s)
	if n > 0 {
		x := (*s)[n-1]
		*s = (*s)[0 : n-1]
		return x
	}
	return nil
}

func (s *sorter) Len() int {
	return len(*s)
}

func (s *sorter) Less(i, j int) bool {
	return (*s)[i].Less((*s)[j])
}

func (s *sorter) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

// Define priority queue struct
type PriorityQueue struct {
	s *sorter
}

func New() *PriorityQueue {
	q := &PriorityQueue{s: new(sorter)}
	heap.Init(q.s)
	return q
}

func (q *PriorityQueue) Push(x Interface) {
	heap.Push(q.s, x)
}

func (q *PriorityQueue) Pop() Interface {
	return heap.Pop(q.s).(Interface)
}

func (q *PriorityQueue) Top() Interface {
	if len(*q.s) > 0 {
		return (*q.s)[0].(Interface)
	}
	return nil
}

func (q *PriorityQueue) Fix(x Interface, i int) {
	(*q.s)[i] = x
	heap.Fix(q.s, i)
}

func (q *PriorityQueue) Remove(i int) Interface {
	return heap.Remove(q.s, i).(Interface)
}

func (q *PriorityQueue) Len() int {
	return q.s.Len()
}

/************测试************/
type Node struct {
	priority int
	value    int
}

func (this *Node) Less(other interface{}) bool {
	return this.priority < other.(*Node).priority
}

func main() {
	q := New()

	q.Push(&Node{priority: 8, value: 1}) //权重 , 值
	q.Push(&Node{priority: 7, value: 2})
	q.Push(&Node{priority: 9, value: 3})

	x := q.Top().(*Node)
	fmt.Println(x.priority, x.value)

	for q.Len() > 0 {
		x = q.Pop().(*Node)
		fmt.Println(x.priority, x.value)
	}
}
</pre>
###Golang sync.Cond   锁的条件变量
<pre>
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Cond有三个方法：Wait，Signal，Broadcast。
Wait添加一个计数，也就是添加一个阻塞的goroutine;
Signal解除一个goroutine的阻塞，计数减一;
Broadcast接触所有wait goroutine的阻塞;

那外部传入的Locker，是对wait，Signal，Broadcast进行保护。防止发送信号的时候，不会有新的goroutine进入wait。在wait逻辑完成前，不会有新的事件发生。

注意：在调用Signal，Broadcast之前，应确保目标进入Wait阻塞状态。
*/

func main() {
	wait := sync.WaitGroup{}
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)
	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wait.Done()
			wait.Add(1)
			cond.L.Lock()
			fmt.Println("Waiting start...")
			cond.Wait()
			fmt.Println("Waiting end...")
			cond.L.Unlock()
			fmt.Println("Goroutine run. Number:", i)
		}(i)
	}
	time.Sleep(time.Second)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	wait.Wait()
}
</pre>
更加简单的可以是
<pre>
package main

import "sync"

func main() {
	cv := sync.NewCond(new(sync.Mutex))
	done := false
	go func() {
		cv.L.Lock()
		//处理事情
		done = true
		cv.Signal()
		cv.L.Unlock()
	}()
	//等待事情结束
	cv.L.Lock()
	for !done {
		cv.Wait()
	}
	cv.L.Unlock()
	// 事情已经结束
}
</pre>
###Golang 线程安全链表
<pre>
package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	data *list.List
}

func NewQueue() *Queue {
	q := new(Queue)
	q.data = list.New()
	return q
}

func (q *Queue) push(v interface{}) {
	defer lock.Unlock()
	lock.Lock()
	q.data.PushFront(v)
}

func (q *Queue) pop() interface{} {
	defer lock.Unlock()
	lock.Lock()
	iter := q.data.Back()
	v := iter.Value
	q.data.Remove(iter)
	return v
}

func (q *Queue) dump() {
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}

var lock sync.Mutex

func main() {
	q := NewQueue()
	go func() {
		q.push("one")
	}()
	go func() {
		q.push("four")
	}()
	q.push("two")
	q.push("three")
	v := q.pop()
	fmt.Println("pop v:", v)
	fmt.Println("......")
	time.Sleep(1 * time.Second)
	q.dump()
}
</pre>
###Linux 常用命令
<pre>
//linux 查看端口占用的项目 
netstat -ntpl 
--->
tcp        0      0 0.0.0.0:10050           0.0.0.0:*               LISTEN      4783/zabbix_agentd  
tcp        0      0 0.0.0.0:11810           0.0.0.0:*               LISTEN      2131/sequoiadb(1181 
tcp        0      0 0.0.0.0:11780           0.0.0.0:*               LISTEN      2132/sdbom(11780)   
tcp        0      0 0.0.0.0:11814           0.0.0.0:*               LISTEN      2131/sequoiadb(1181 
tcp6       0      0 :::6379                 :::*                    LISTEN      2946/redis-server * 
tcp6       0      0 :::80                   :::*                    LISTEN      26379/./zcm         
tcp6       0      0 :::9009                 :::*                    LISTEN      1987/java           
tcp6       0      0 :::22                   :::*                    LISTEN      1112/sshd  


//查看开机启动项
 cat /etc/rc.d/rc.local

/防止rc.local重复运行        【放入rc.local文件中】
touch /var/lock/subsys/local

//redis启动
1-->   /usr/local/redis/src/redis-server /etc/redis.conf
2-->   /home/soft/redis-2.8.17/src/redis-server /home/soft/redis-2.8.17/redis.conf

//启动svn服务 
svnserve  -d -r /home/svn/svndata      | svnserve -d -r [仓库目录]

//在当前目录下搜索文件内含有某字符串【大小写敏感】的文件 
find . -type f |xargs grep 'rdb'   

//在当前目录下搜索文件内含有某字符串【忽略大小写】的文件
find . -type f -name '*.sh' | xargs grep -i 'your_string'

//在当前目录下根据文件名模糊搜索文件
find . -name 'rdb.*'

//启动mysqld服务
service mysqld start

//启动nginx服务
/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf

//关闭nginx服务
/usr/local/nginx/sbin/nginx -s stop

//mysql 更新语句

UPDATE table   
SET title =( CASE title
WHEN instr(LOWER(title), '最大化') <= 0 THEN
	title
WHEN instr(LOWER(title), '最终') <= 0 THEN
	title
WHEN instr(LOWER(title), '最新') <= 0 THEN
	title
ELSE
	REPLACE(title,'最','')
END )
WHERE
	id = 1;

//mysql if else

select *,if(actvity_name='活动1',"男","女") as ssva from activity where actvity_name != ""

touch 命令
linux的touch命令不常用，一般在使用make的时候可能会用到，用来修改文件时间戳，或者新建一个不存在的文件

Linux系统本地自启动脚本常用：
touch /var/local
/etc/init.d/mysqld start
/etc/init.d/nginx start
/etc/init.d/php-fpm start
service vsftpd start      //安全的ftp服务器软件
/usr/local/redis/src/redis-server /etc/redis.conf 
svnserve -d -r /svn
service zabbix-agent start   // 监控软件
/home/tomcat/tomcat_9.0/bin/catalina.sh start   //web services

重启脚本：
<pre>
#!/bin/sh 

port=80
#关闭watch.sh
echo -e "closing watch.sh..."
watchpid=$(ps -ef | grep -v grep | grep watch.sh | awk '{print $2}' | tail -n 1)
if [ -n "$watchpid" ]
then
    kill -9 $watchpid
    echo -e "kill watch.sh by pid:"$watchpid
fi


# 编译
echo -e "building zcm..."
cd $GOPATH/src/zcm
go build


# 关闭zcm
echo -e "closing zcm..."
zcmpid=$(lsof -i:$port | awk '{print $2}' | tail -n 1) 
if [ -n "$zcmpid" ]
then
    kill -9 $zcmpid
    echo -e "kill zcm which listening on 80 by pid:"$zcmpid
fi

# 运行
nohup ./zcm &

while [ -z "$(lsof -i:$port |awk '{print $2}' | tail -n 1)" ]
do
   sleep 0.5
done

newzcmpid=$(lsof -i:$port |awk '{print $2}' | tail -n 1) 
echo "Finish boot, zcm new pid: "$newzcmpid


# 启动watch.sh
echo -e "start watch.sh"
nohup ../watch.sh &


# 查看zcm输出
tail -f nohup.out
</pre>

</pre>
###Golang goroutine pool
<pre>
package main

import (
  "fmt"
  "runtime"
  "time"

  "github.com/ivpusic/grpool"
)

func main() {
  // number of workers, and size of job queue
  pool := grpool.NewPool(100, 50)

  // release resources used by pool
  defer pool.Release()

  // submit one or more jobs to pool
  for i := 0; i < 10; i++ {
    count := i

    pool.JobQueue <- func() {
      fmt.Printf("I am worker! Number %d\n", count)
    }
  }

  // dummy wait until jobs are finished
  time.Sleep(1 * time.Second)
}
<pre>
###Golang 获取当前文件执行的行数
<pre>
package main 

import (
	"runtime"
	"sync"
)

type worker struct {
	Func func()
}

func main() {
	var wg sync.WaitGroup

	channels := make(chan worker, 10)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ch := range channels {
				//reflect.ValueOf(ch.Func).Call(ch.Args)
				//
				ch.Func()
			}
		}()
	}

	for i := 0; i < 100; i++ {
		j := i
		wk := worker{
			Func: func() {
				fmt.Println(j + j)
				_, file, line, ok := runtime.Caller(1)
				if ok {
					fmt.Println("filename:", file, " line:", line)
				}
			},
		}
		channels <- wk
	}
	close(channels)
	wg.Wait()
}
</pre>
###Beego框架工作问题记录
<pre>
//如果是web应用，则在main.go文件中必须import _ "*/routers"
	import _"*/routers"
</pre>
###Golang 并发下载
<pre>
package main

//HTTP的并发下载
import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"mime"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	DEFAULT_DOWNLOAD_BLOCK int64 = 4096
)

type GoGet struct {
	Url           string
	Cnt           int
	DownloadBlock int64
	CostomCnt     int
	Latch         int
	Header        http.Header
	MediaType     string
	MediaParams   map[string]string
	FilePath      string // 包括路径和文件名
	GetClient     *http.Client
	ContentLength int64
	DownloadRange [][]int64
	File          *os.File
	TempFiles     []*os.File
	WG            sync.WaitGroup
}

func NewGoGet() *GoGet {
	get := new(GoGet)
	get.FilePath = "./"
	get.GetClient = new(http.Client)

	flag.Parse()
	get.Url = *urlFlag
	get.DownloadBlock = DEFAULT_DOWNLOAD_BLOCK

	return get
}

var urlFlag = flag.String("u", "http://7b1h1l.com1.z0.glb.clouddn.com/bryce.jpg", "Fetch file url")

// var cntFlag = flag.Int("c", 1, "Fetch concurrently counts")

func main() {
	get := NewGoGet()

	download_start := time.Now()

	req, err := http.NewRequest("HEAD", get.Url, nil)
	resp, err := get.GetClient.Do(req)
	get.Header = resp.Header
	if err != nil {
		log.Panicf("Get %s error %v.\n", get.Url, err)
	}
	get.MediaType, get.MediaParams, _ = mime.ParseMediaType(get.Header.Get("Content-Disposition"))
	get.ContentLength = resp.ContentLength
	get.Cnt = int(math.Ceil(float64(get.ContentLength / get.DownloadBlock)))
	if strings.HasSuffix(get.FilePath, "/") {
		get.FilePath += get.MediaParams["filename"]
	}
	get.File, err = os.Create(get.FilePath)
	if err != nil {
		log.Panicf("Create file %s error %v.\n", get.FilePath, err)
	}
	log.Printf("Get %s MediaType:%s, Filename:%s, Size %d.\n", get.Url, get.MediaType, get.MediaParams["filename"], get.ContentLength)
	if get.Header.Get("Accept-Ranges") != "" {
		log.Printf("Server %s support Range by %s.\n", get.Header.Get("Server"), get.Header.Get("Accept-Ranges"))
	} else {
		log.Printf("Server %s doesn't support Range.\n", get.Header.Get("Server"))
	}

	log.Printf("Start to download %s with %d thread.\n", get.MediaParams["filename"], get.Cnt)
	var range_start int64 = 0
	for i := 0; i < get.Cnt; i++ {
		if i != get.Cnt-1 {
			get.DownloadRange = append(get.DownloadRange, []int64{range_start, range_start + get.DownloadBlock - 1})
		} else {
			// 最后一块
			get.DownloadRange = append(get.DownloadRange, []int64{range_start, get.ContentLength - 1})
		}
		range_start += get.DownloadBlock
	}
	// Check if the download has paused.
	for i := 0; i < len(get.DownloadRange); i++ {
		range_i := fmt.Sprintf("%d-%d", get.DownloadRange[i][0], get.DownloadRange[i][1])
		temp_file, err := os.OpenFile(get.FilePath+"."+range_i, os.O_RDONLY|os.O_APPEND, 0)
		if err != nil {
			temp_file, _ = os.Create(get.FilePath + "." + range_i)
		} else {
			fi, err := temp_file.Stat()
			if err == nil {
				get.DownloadRange[i][0] += fi.Size()
			}
		}
		get.TempFiles = append(get.TempFiles, temp_file)
	}

	go get.Watch()
	get.Latch = get.Cnt
	for i, _ := range get.DownloadRange {
		get.WG.Add(1)
		go get.Download(i)
	}

	get.WG.Wait()

	for i := 0; i < len(get.TempFiles); i++ {
		temp_file, _ := os.Open(get.TempFiles[i].Name())
		cnt, err := io.Copy(get.File, temp_file)
		if cnt <= 0 || err != nil {
			log.Printf("Download #%d error %v.\n", i, err)
		}
		temp_file.Close()
	}
	get.File.Close()
	log.Printf("Download complete and store file %s with %v.\n", get.FilePath, time.Now().Sub(download_start))
	defer func() {
		for i := 0; i < len(get.TempFiles); i++ {
			err := os.Remove(get.TempFiles[i].Name())
			if err != nil {
				log.Printf("Remove temp file %s error %v.\n", get.TempFiles[i].Name(), err)
			} else {
				log.Printf("Remove temp file %s.\n", get.TempFiles[i].Name())
			}
		}
	}()
}

func (get *GoGet) Download(i int) {
	defer get.WG.Done()
	if get.DownloadRange[i][0] > get.DownloadRange[i][1] {
		return
	}
	range_i := fmt.Sprintf("%d-%d", get.DownloadRange[i][0], get.DownloadRange[i][1])
	log.Printf("Download #%d bytes %s.\n", i, range_i)

	defer get.TempFiles[i].Close()

	req, err := http.NewRequest("GET", get.Url, nil)
	req.Header.Set("Range", "bytes="+range_i)
	resp, err := get.GetClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Download #%d error %v.\n", i, err)
	} else {
		cnt, err := io.Copy(get.TempFiles[i], resp.Body)
		if cnt == int64(get.DownloadRange[i][1]-get.DownloadRange[i][0]+1) {
			log.Printf("Download #%d complete.\n", i)
		} else {
			req_dump, _ := httputil.DumpRequest(req, false)
			resp_dump, _ := httputil.DumpResponse(resp, true)
			log.Panicf("Download error %d %v, expect %d-%d, but got %d.\nRequest: %s\nResponse: %s\n", resp.StatusCode, err, get.DownloadRange[i][0], get.DownloadRange[i][1], cnt, string(req_dump), string(resp_dump))
		}
	}
}

func (get *GoGet) Watch() {
	fmt.Printf("[===============>]\n")
}
</pre>
###Linux nohup 
<pre>
	//将项目运行的日志名称改成日期格式
	nohup ./项目名称 >  `date +%Y-%m-%d`.out 2>&1 &
	//每秒执行一次
	/bin/sleep 1    #每秒执行一次
	

/**********************************************************************/备份数据


	//备份数据库
	#!/bin/bash
	Date=`date +%Y%m%d`
	Begin=`date +"%Y年%m月%d日 %H:%M:%S"`
	LogFile=10分钟.log
	#mysqldump -h主机名称 -u用户名 -p密码 数据库名  表名  > 写入的文件名称
	mysqldump -h00.00.00.00  -uUSERNAME -pPASSWORD  数据库名称   表一  表二  表三  表四 > /home/log/$Date.table
	#备份mysql数据库
	Last=`date +"%Y年%m月%d日 %H:%M:%S"`
	echo 开始:$Begin 结束:$Last  success  >> /home/log/$LogFile
	
	
	#ftp自动备份 
	#lcd   ftp切换本地目录
	#cd    ftp切换远程目录
	#user  用户名  密码
	ftp -ivn  00.00.00.00 <<EOF
	user  USERNAME  PASSWORD
	lcd /home/log 
	cd /mysql-copy/table
	put $Date.table
	bye
	EOF	
	exit

/**********************************************************************/备份数据

	#!/bin/bash
	Date=`date +%Y%m%d`
	Begin=`date +"%Y年%m月%d日 %H:%M:%S"`
	LogFile=15分钟exam_log
	mysqldump -h00.00.00.00  -uUSERNAME -pPASSWORD  exam_log > /home/log/$Date-exam_log.sql
	cd /home/log
	tar -zcvf $Date-exam_log.sql.tar.gz $Date-exam_log.sql
	#备份数据库
	Last=`date +"%Y年%m月%d日 %H:%M:%S"`
	echo 开始:$Begin 结束:$Last  success  >> /home/log/$LogFile


	ftp -ivn  00.00.00.00 <<EOF
	user  USERNAME  PASSWORD
	bin
	lcd /home/log
	cd mysql-copy/log
	put $Date-exam_log.sql.tar.gz
	bye
	EOF
	exit
	
	find /home/log/ -amin +360  -exec rm -rvf {} \;

/**********************************************************************/ 进程守护,项目停止自动重启

	#!/bin/sh
	PRO_NAME="项目名称"
	while true;
	do
	NUM=`ps aux | grep PRO_NAME | grep -v grep |wc -l`
	if [ "${NUM}" -lt 1 ]
	then
	echo "${PRO_NAME} was killed"
	cd /home/项目目录
	./项目名称
	else
	echo "ok"
	fi
	sleep 3 
	done
</pre>
###Golang sync.WaitGroup 
<pre>
package main

import (
	"log"
	"sync"
)

// WaitGroup 同步的是 goroutine
// wg 给拷贝传递到了 goroutine 中，导致只有 Add 操作，其实 Done 操作是在 wg 的副本执行的
func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			log.Printf("i:%d", i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	log.Println("exit")
}
</pre>
###Mysql 获取此前30天的信息
<pre>
	SELECT * from (
	SELECT * from exam_1 where state= 'YES' and create_date >= now()-interval 30 DAY  UNION ALL
	SELECT * from exam_2 where state= 'YES' and create_date >= now()-interval 30 DAY  UNION ALL
	SELECT * from exam_3 where state= 'YES' and create_date >= now()-interval 30 DAY )exams 
	ORDER BY exams.create_date ASC
</pre>
###Mysql  在整个数据库中搜索某一个字段
<pre>
-- 查找mysql数据库中所有包含特定名字的字段所在的表 
select * from INFORMATION_SCHEMA.columns where COLUMN_NAME Like '%config%';

-- group by * having ~~~
select * from table group by id having count(id)>1
</pre> 
###Golang 对结构体进行排序 第三方库 赞
https://github.com/as/structslice
<pre>
//structslice.go
//////////////////////////////////////
package structslice

import (
	"fmt"
	"reflect"
	"sort"
)

const (
	errFieldFmt = "structslice: The field '%s' doesn't exist in struct '%s'"
)

// structSlice is a control structure for implementing abstract operations
// on a slice of structs.
type structSlice struct {
	Value          reflect.Value
	SortFieldIndex int
}

// sortByIndex sorts the slice of structs by the field index 'i'
func sortByIndex(v interface{}, i int) {
	s := attach(v)
	s.SortFieldIndex = i
	sort.Sort(s)
}

// sortStableByIndex is like sortByIndex, except it performs a stable sort.
// Because it performs a stable sort, it accepts a variadic number of sort
// keys. Sorting is done for every key in the order that the key is passed in
// to the function.
func sortStableByIndex(v interface{}, i ...int) {
	if len(i) == 0 {
		return
	}

	s := attach(v)
	for _, v := range i {
		s.SortFieldIndex = v
		sort.Stable(s)
	}
}

// attach binds to the slice of structs and returns a structSlice object
// for executing sorting operations on the slice elements. Attach panics
// if the underlying interface, v, is not a slice of structs.
func attach(v interface{}) *structSlice {
	//panicf panics with a pre-formatted error string
	panicf := func(f string, s ...interface{}) {
		panic(fmt.Sprintf("structslice: input must be a slice of structs. %s", fmt.Sprintf(f, s)))
	}

	s := new(structSlice)
	s.Value = reflect.ValueOf(v)

	vtype := reflect.TypeOf(v)
	// Test one: Panics if the v interface isn't a slice
	if vtype.Kind() != reflect.Slice {
		panicf("expected: [slice], actual: %s\n", vtype.Kind())
	}

	// Test two: Panics if the elements of v are not structs
	if vtype.Elem().Kind() != reflect.Struct {
		panicf("expected: [slice struct], actual: %v\n", vtype.Kind(), vtype.Elem().Kind())
	}

	return s
}

// Less satisfies the sort.Interface type in the go standard library
func (s structSlice) Less(i, j int) bool {
	it := s.Value.Index(i).Type().Field(s.SortFieldIndex)
	jt := s.Value.Index(j).Type().Field(s.SortFieldIndex)

	if it.Type.Kind() != jt.Type.Kind() {
		panic(fmt.Sprintf("structSlice.Less(): Type mismatch %s != %s", it.Type.Name(), jt.Type.Name()))
	}

	iv := s.Value.Index(i).Field(s.SortFieldIndex).Interface()
	jv := s.Value.Index(j).Field(s.SortFieldIndex).Interface()

	switch t := iv.(type) {
	case string:
		return t < jv.(string)
	case bool:
		return t && !jv.(bool)
	case int:
		return t < jv.(int)
	case int32:
		return t < jv.(int32)
	case int64:
		return t < jv.(int64)
	case float64:
		return t < jv.(float64)
	case float32:
		return t < jv.(float32)
	case Stringer:
		return t.String() < jv.(Stringer).String()
	case Comparer:
		return t.Less(jv.(Comparer))
	}

	return false
}

// Len satisfies the sort.Interface type in the go standard library
func (s structSlice) Len() int {
	return s.Value.Len()
}

// Swap satisfies the sort.Interface type in the go standard library
func (s structSlice) Swap(i, j int) {
	v := s.Value
	tmp := v.Index(i).Interface()
	v.Index(i).Set(v.Index(j))
	v.Index(j).Set(reflect.ValueOf(tmp))
}

// Index returns the value given by the index of the struct slice
func (s structSlice) Index(i int) reflect.Value {
	return s.Value.Index(i)
}



///**************************************************///



//interfaces.go
/////////////////////////////////


package structslice

import (
	"fmt"
	"reflect"
)

// Comparer is an interface for types that can compare themselves to each other.
type Comparer interface {
	Less(Comparer) bool
}

// Stringer is an interface for types that have a string representation
type Stringer interface {
	String() string
}

// SortByName sorts the slice of structs by the field name given by 'n'
func SortByName(v interface{}, n string) error {
	s := attach(v)
	fmt.Println("s.Value type is", reflect.TypeOf(s.Value.Interface()))
	f, ok := s.Value.Index(0).Type().FieldByName(n)

	if !ok {
		return fmt.Errorf(errFieldFmt, n, s.Value.Index(0).Type())
	}

	sortByIndex(v, f.Index[0])
	return nil
}

// SortStableByName is like SortByName, except it performs a stable sort.
// Because it performs a stable sort, it accepts a variadic number of sort
// keys. Sorting is done for every key in the order that the key is passed in
// to the function.
func SortStableByName(v interface{}, n ...string) error {
	if len(n) == 0 {
		return nil
	}
	s := attach(v)

	keys := make([]int, len(n))
	for i, v := range n {
		f, ok := s.Value.Index(0).Type().FieldByName(v)
		if !ok {
			return fmt.Errorf(errFieldFmt, v, s.Value.Index(0).Type())
		}
		keys[i] = f.Index[0]
	}

	sortStableByIndex(v, keys...)

	return nil
}

//使用 
 internalDB := Employees{
        Employee{1, 95000, Name{"Jake", "M", "Anderson"}},
        Employee{5, 45000, Name{"Hunter", "L", "Alice"}},
        Employee{6, 345000, Name{"Steinberg", "F", "Charles"}},
        Employee{2, 108000, Name{"Williams", "L", "Bill"}},
        Employee{4, 190000, Name{"Morgan", "A", "Janice"}},
        Employee{3, 108000, Name{"Williams", "L", "Will"}},
        Employee{5, 145000, Name{"Steinberg", "L", "Alice"}},
    }

    fmt.Println(internalDB)
    ss.SortByName(internalDB, "ID")
  	fmt.Println(internalDB)
</pre>
###Golang 错误延迟调用
<pre>
defer func() {
	if err := recover(); err != nil {
		beego.Emergency("[ERROR]", err)
		//RestartThisServe()
	}
}()
</pre>
###Golang 嵌套结构体
<pre>
type ResultStruct struct {
	Identitytype    int
	Cardlist      []struct {
		Phone     string
		Code      string
	}
	Error_code string
	Error_msg  string
}


//RSA 加解密
//  RSA加密 PKCS8
func RsaEncrypt2(origData, USERS_publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey2)
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

// RSA解密 PKCS8
func RsaDecrypt2(ciphertext []byte) ([]byte, error) {

	block, _ := pem.Decode(USERS_privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	priv := privInterface.(*rsa.PrivateKey)
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}



****************************************************

package tool

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	ran "math/rand"
	"os"
	"reflect"
	"sort"
	"strconv"
	"time"
	"github.com/astaxie/beego"
)

const key = `fvck`

//对字符串进行MD5哈希
func Md5Encrypt(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串进行SHA1哈希
func ShaEncrypt(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//base64 加密
func Base64Encrypt(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

//base64 解密
func Base64Decrypt(data string) []byte {
	resdata, _ := base64.StdEncoding.DecodeString(data)
	return resdata
}

//  RSA加密 PKCS8
func RsaEncrypt2(origData, publicKey2 []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey2)
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

// RSA解密 PKCS8
func RsaDecrypt2(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	priv := privInterface.(*rsa.PrivateKey)
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// 从文件中读取数据
func ReadFileContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		beego.Emergency(file, err)
	}
	file1, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeType)
	if err != nil {
		beego.Emergency("从文件读取数据错误！第1步:" + err.Error())
		return "", err
	}
	defer file1.Close()
	// 往创建的文件中写入字符
	//_, err = file1.WriteString("aaaaa\r\nbbbbb\r\ncccccc")
	//if err != nil {
	//    panic(err)
	//}
	// A. 使用 bufio按行读取文件
	//br := bufio.NewReader(file1)
	//for {
	//    line, err := br.ReadString('\r')
	//    if err == io.EOF {
	//        fmt.Println("eof")
	//        break
	//    } else {
	//        fmt.Printf("%v", line)
	//    }
	//}

	// B. 使用ioutil读取文件所有内容
	b, err := ioutil.ReadAll(file1)
	if err != nil {
		beego.Emergency("从文件读取数据错误！第2步:" + err.Error())
		return "", err
	}
	//	fmt.Printf("%v", string(b))
	//	time.Sleep(3 * time.Second)
	return string(b), nil
}

//struct转换成map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// aes 解密 ECB  PKCS5
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := codec.NewECBDecrypter(block)
	blockMode.CryptBlocks(crypted, crypted)
	crypted = PKCS5UnPadding(crypted)
	return crypted, nil

}

//生成密码aes密码 16位密码
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

//aes加密 ECB  PKCS5
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := codec.NewECBEncrypter(block)

	crypted := origData
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	//	crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
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

//rsa 生成待签名串
func GenSignDataYbByRsa(m map[string]interface{}) string {
	sorted_keys := make([]string, 0)
	for k, _ := range m {
		sorted_keys = append(sorted_keys, k)
	}
	// sort 'string' key in increasing order
	sort.Strings(sorted_keys)
	var valuestr string
	for _, k := range sorted_keys {
		//	fmt.Printf("k=%v, v=%v\n", k, m[k])
		valuestr += fmt.Sprint(m[k])
	}
	//	beego.Info(valuestr)
	return valuestr
}

// 产生签名：使用PKCS8 进行加密签名（私钥）
func SignYb(data []byte) (sig []byte, err error) {
	hashFunc := crypto.SHA1
	h := hashFunc.New()
	h.Write(data)
	digest := h.Sum(nil)

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("privateKey key error")
	}
	//pubInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	pubInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PrivateKey)

	bytes, err := rsa.SignPKCS1v15(nil, pub, hashFunc, digest)
	if err != nil {
		panic(err)
	}
	return bytes, err
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
###Golang 字符串/字节比较方法
<pre>
package main

import (
	"crypto/subtle"
)

func SecureCompare(given, actual []byte) bool {
	if subtle.ConstantTimeEq(int32(len(given)), int32(len(actual))) == 1 {
		if subtle.ConstantTimeCompare(given, actual) == 1 {
			return true
		}
		return false
	}
	// Securely compare actual to itself to keep constant time, but always return false
	if subtle.ConstantTimeCompare(actual, actual) == 1 {
		return false
	}
	return false
}

func SecureCompareString(given, actual string) bool {
	// The following code is incorrect:
	// return SecureCompare([]byte(given), []byte(actual))

	if subtle.ConstantTimeEq(int32(len(given)), int32(len(actual))) == 1 {
		if subtle.ConstantTimeCompare([]byte(given), []byte(actual)) == 1 {
			return true
		}
		return false
	}
	// Securely compare actual to itself to keep constant time, but always return false
	if subtle.ConstantTimeCompare([]byte(actual), []byte(actual)) == 1 {
		return false
	}
	return false
}

func main() {
	res := SecureCompareString("jason", "Jason")
	println(res)
}
</pre>
###Golang 去除空格与换行
<pre>
package main 

import (
	"bytes"
	"strings"
)

var newlineBytes = []byte{'\n'}

// 去掉 src 开头和结尾的空白, 如果 src 包括换行, 去掉换行和这个换行符两边的空白
//  NOTE: 根据 '\n' 来分行的, 某些系统或软件用 '\r' 来分行, 则不能正常工作.
func TrimSpace(src []byte) []byte {
	bytesArr := bytes.Split(src, newlineBytes)
	for i := 0; i < len(bytesArr); i++ {
		bytesArr[i] = bytes.TrimSpace(bytesArr[i])
	}
	return bytes.Join(bytesArr, nil)
}

// 去掉 src 开头和结尾的空白, 如果 src 包括换行, 去掉换行和这个换行符两边的空白
//  NOTE: 根据 '\n' 来分行的, 某些系统或软件用 '\r' 来分行, 则不能正常工作.
func TrimSpaceString(src string) string {
	strs := strings.Split(src, "\n")
	for i := 0; i < len(strs); i++ {
		strs[i] = strings.TrimSpace(strs[i])
	}
	return strings.Join(strs, "")
}
</pre>
###Golang xml工具
<pre>
package xml 

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
)

// DecodeXMLToMap decodes xml reading from io.Reader and returns the first-level sub-node key-value set,
// if the first-level sub-node contains child nodes, skip it.
func DecodeXMLToMap(r io.Reader) (m map[string]string, err error) {
	m = make(map[string]string)
	var (
		decoder = xml.NewDecoder(r)
		depth   = 0
		token   xml.Token
		key     string
		value   bytes.Buffer
	)
	for {
		token, err = decoder.Token()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}

		switch v := token.(type) {
		case xml.StartElement:
			depth++
			switch depth {
			case 2:
				key = v.Name.Local
				value.Reset()
			case 3:
				if err = decoder.Skip(); err != nil {
					return
				}
				depth--
				key = "" // key == "" indicates that the node with depth==2 has children
			}
		case xml.CharData:
			if depth == 2 && key != "" {
				value.Write(v)
			}
		case xml.EndElement:
			if depth == 2 && key != "" {
				m[key] = value.String()
			}
			depth--
		}
	}
}

// EncodeXMLFromMap encodes map[string]string to io.Writer with xml format.
//  NOTE: This function requires the rootname argument and the keys of m (type map[string]string) argument
//  are legitimate xml name string that does not contain the required escape character!
func EncodeXMLFromMap(w io.Writer, m map[string]string, rootname string) (err error) {
	switch v := w.(type) {
	case *bytes.Buffer:
		bufw := v
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return nil

	case *bufio.Writer:
		bufw := v
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return bufw.Flush()

	default:
		bufw := bufio.NewWriterSize(w, 256)
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return bufw.Flush()
	}
}
</pre>
###Golang orm事务
<pre>
//删除账号 mysql事务
func DeleteAccount(account string) error {
	o := orm.NewOrm()
	sql1 := "delete from table1 where account_id=(SELECT id FROM table2 where account=?)"
	sql2 := "delete from table2 where account=?"
	o.Begin()
	_, err1 := o.Raw(sql1, account).Exec()
	if err1 == nil {
		_, err2 := o.Raw(sql2, account).Exec()
		if err2 == nil {
			o.Commit()
		} else {
			return err2
		}
	}
	o.Rollback()
	return err1
}
</pre>
###Golang 图片操作 图片文件不同格式数据的转换
<pre>
base64 -> file

ddd, _ := base64.StdEncoding.DecodeString(datasource) //成图片文件并把文件写入到buffer
err2 := ioutil.WriteFile("./output.jpg", ddd, 0666)   //buffer输出到jpg文件中（不做处理，直接写到文件）
datasource base64 string

base64 -> buffer

ddd, _ := base64.StdEncoding.DecodeString(datasource) //成图片文件并把文件写入到buffer
bbb := bytes.NewBuffer(ddd)                           // 必须加一个buffer 不然没有read方法就会报错
转换成buffer之后里面就有Reader方法了。才能被图片API decode

buffer-> ImageBuff（图片裁剪,代码接上面）

m, _, _ := image.Decode(bbb)                                       // 图片文件解码
rgbImg := m.(*image.YCbCr)
subImg := rgbImg.SubImage(image.Rect(0, 0, 200, 200)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
img -> file(代码接上面)

f, _ := os.Create("test.jpg")     //创建文件
defer f.Close()                   //关闭文件
jpeg.Encode(f, subImg, nil)       //写入文件
img -> base64(代码接上面)

emptyBuff := bytes.NewBuffer(nil)                  //开辟一个新的空buff
jpeg.Encode(emptyBuff, subImg, nil)                //img写入到buff
dist := make([]byte, 50000)                        //开辟存储空间
base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
fmt.Println(string(dist))                          //输出图片base64(type = []byte)
_ = ioutil.WriteFile("./base64pic.txt", dist, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
imgFile -> base64

ff, _ := ioutil.ReadFile("output2.jpg")               //我还是喜欢用这个快速读文件
bufstore := make([]byte, 5000000)                     //数据缓存
base64.StdEncoding.Encode(bufstore, ff)               // 文件转base64
_ = ioutil.WriteFile("./output2.jpg.txt", dist, 0666) //直接写入到文件就ok完活了。
</pre>
###Golang sha1
<pre>
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func sha1string(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
func main() {
	fmt.Println(sha1string("jaSon"))
}
</pre>
###Golang 时间戳格式化
<pre>
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))
}
</pre>
###Golang 获取当月的第一天
<pre>
package main

import (
	"fmt"
	"strconv"
	"time"
)

var month = map[string]string{"January": "01", "February": "02", "March": "03", "April": "04", "May": "05", "June": "06", "July": "07", "August": "08", "September": "09", "October": "10", "November": "11", "December": "12"}

//获取当月的第一天
func GetCurrentMonth() string {
	year := strconv.Itoa(time.Now().Year())
	m := month[time.Now().Month().String()]
	return year + "-" + m + "-" + "01"
}
func main() {
	fmt.Println(GetCurrentMonth())
}
</pre>
###Golang 微信公众号
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
		Content := `"欢迎关注 我的微信"`
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
###Golang 重写println方法
<pre>
package main

import (
	"log"
)

func Ldefault() {
	log.Println("这是默认的格式\n")
}

func Ldate() {
	log.SetFlags(log.Ldate)
	log.Println("这是输出日期格式\n")
}

func Ltime() {
	log.SetFlags(log.Ltime)
	log.Println("这是输出时间格式\n")
}

func Lmicroseconds() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("这是输出微秒格式\n")
}

func Llongfile() {
	log.SetFlags(log.Llongfile)
	log.Println("这是输出路径+文件名+行号格式\n")
}

func Lshortfile() {
	log.SetFlags(log.Lshortfile)
	log.Println("这是输出文件名+行号格式\n")
}

func LUTC() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.Println("这是输出 使用标准的UTC时间格式 格式\n")
}
func main() {
	Ldefault()
	Ldate()
	Ltime()
	Lmicroseconds()
	Llongfile()
	Lshortfile()
	LUTC()
}
</pre>
###Golang最简单的http转发 翻墙程序 赞赞赞【http://www.flysnow.org/2016/12/24/golang-http-proxy.html】
<pre>
package main

//将该文件放到国外的服务器上，在自己机器上配置好HTTP代理，就可以到处访问了
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Panic(err)
	}
	for {
		client, err := l.Accept()
		if strings.Contains(client.RemoteAddr().String(), "60.191.37.251") {
			if err != nil {
				log.Panic(err)
			}
			log.Println("Start Working ---->", client.RemoteAddr().String())
			go handleClientRequest(client)
		}

	}
}
func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}
	if hostPortURL.Opaque == "443" { //https访问
		address = hostPortURL.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}
	//获得了请求的host和port，就开始拨号吧
	fmt.Println(address)
	if address != ":80" && address != ":443" {
		server, err := net.Dial("tcp", address)
		if err != nil {
			log.Println(err)
			return
		}
		if method == "CONNECT" {
			fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n")
		} else {
			server.Write(b[:n])
		}
		//进行转发
		go io.Copy(server, client)
		io.Copy(client, server)
	}
}
</pre>
###Golang socket5 代理
<pre>
package main

//简易版本的Socket5代理
import (
	"io"
	"log"
	"net"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}
}
func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	if b[0] == 0x05 { //只处理Socket5协议
		//客户端回应：Socket服务端不需要验证方式
		client.Write([]byte{0x05, 0x00})
		n, err = client.Read(b[:])
		var host, port string
		switch b[3] {
		case 0x01: //IP V4
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03: //域名
			host = string(b[5 : n-2]) //b[4]表示域名的长度
		case 0x04: //IP V6
			host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
		}
		port = strconv.Itoa(int(b[n-2]<<8) | int(b[n-1]))
		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			log.Println(err)
			return
		}
		defer server.Close()
		client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //响应客户端连接成功
		//进行转发
		go io.Copy(server, client)
		io.Copy(client, server)
	}
}
</pre>
###Golang 读取大内存文件效率比较
<pre>
package main

import (
    "bufio"
    "crypto/sha1"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "runtime"
    "testing"
)

const (
    fname    = "./poe.tgz"
    hashSum  = "0bd099caa5ac2be51075085daea9cbac89d8a6be"
)

func copyFile() {
    f, err := os.Open(fname)
    if err != nil {
        panic(err)
        return
    }
    defer f.Close()

    h := sha1.New()
    _, err = io.Copy(h, f)

    if err != nil {
        panic(err)
        return
    }
    if fmt.Sprintf("%x", h.Sum(nil)) != hashSum {
        panic(fmt.Errorf("not equal"))
    }
}
//bufio.Reader+io.Copy都是最优的方式
func readAll() {         //最优
    runtime.GC()
    f, err := os.Open(fname)
    if err != nil {
        panic(err)
        return
    }
    defer f.Close()

    h := sha1.New()
    bs, err := ioutil.ReadAll(f)
    if err != nil {
        panic(err)
        return
    }

    _, err = h.Write(bs)
    if err != nil {
        panic(err)
        return
    }
    if fmt.Sprintf("%x", h.Sum(nil)) != hashSum {
        panic(fmt.Errorf("not equal"))
    }
}

func bufioRead() {
    f, err := os.Open(fname)
    if err != nil {
        panic(err)
        return
    }
    defer f.Close()

    br := bufio.NewReader(f)

    h := sha1.New()
    _, err = io.Copy(h, br)

    if err != nil {
        panic(err)
        return
    }
    if fmt.Sprintf("%x", h.Sum(nil)) != hashSum {
        panic(fmt.Errorf("not equal"))
    }
}

func BenchmarkCopy(b *testing.B) {
    for i := 0; i < b.N; i++ {
        copyFile()
    }
}

func BenchmarkBufio(b *testing.B) {
    for i := 0; i < b.N; i++ {
        bufioRead()
    }
}

func BenchmarkReadall(b *testing.B) {
    for i := 0; i < b.N; i++ {
        readAll()
    }
}
</pre>
###Golang 内存复用，减少内存占用
<pre>
package main

import (
    "container/list"
    "fmt"
    "math/rand"
    "runtime"
    "time"
)

var makes int
var frees int

func makeBuffer() []byte {
    makes += 1
    return make([]byte, rand.Intn(5000000)+5000000)
}

type queued struct {
    when time.Time
    slice []byte
}

func makeRecycler() (get, give chan []byte) {
    get = make(chan []byte)
    give = make(chan []byte)

    go func() {
        q := new(list.List)
        for {
            if q.Len() == 0 {
                q.PushFront(queued{when: time.Now(), slice: makeBuffer()})
            }

            e := q.Front()

            timeout := time.NewTimer(time.Minute)
            select {
            case b := <-give:
                timeout.Stop()
                q.PushFront(queued{when: time.Now(), slice: b})

           case get <- e.Value.(queued).slice:
               timeout.Stop()
               q.Remove(e)

           case <-timeout.C:
               e := q.Front()
               for e != nil {
                   n := e.Next()
                   if time.Since(e.Value.(queued).when) > time.Minute {
                       q.Remove(e)
                       e.Value = nil
                   }
                   e = n
               }
           }
       }

    }()

    return
}

func main() {
    pool := make([][]byte, 20)

    get, give := makeRecycler()

    var m runtime.MemStats
    for {
        b := <-get
        i := rand.Intn(len(pool))
        if pool[i] != nil {
            give <- pool[i]
        }

        pool[i] = b

        time.Sleep(time.Second)

        bytes := 0
        for i := 0; i < len(pool); i++ {
            if pool[i] != nil {
                bytes += len(pool[i])
            }
        }

        runtime.ReadMemStats(&m)
        fmt.Printf("%d,%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc
             m.HeapIdle, m.HeapReleased, makes, frees)
    }
}
</pre>
###Golang GC优化 GC调优
<pre>
package main

/*
GC优化
优化gc的方式仅仅只能是通过优化程序。但go有一个优势：
有真正的array（而仅仅是an array of referece）。
go的gc算法是mark and sweep，array对此是友好的：整个array一次性被处理。
可以用一个array用open addressing的方式实现map，
以此优化gc（也会减少内存的使用，后面可以看到）
*/

import (
	"fmt"
	//"math/rand"
)

type DealTiny struct {
	Dealid    int32
	Classid   int32
	Mttypeid  int32
	Bizacctid int32
	Isonline  bool
	Geocnt    int32
}

type DealMap struct {
	table   []DealTiny
	buckets int
	size    int
}

const SIZE = 50000

// round 到最近的2的倍数
func minBuckets(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func hashInt32(x int) int {
	x = ((x >> 16) ^ x) * 0x45d9f3b
	x = ((x >> 16) ^ x) * 0x45d9f3b
	x = ((x >> 16) ^ x)
	return x
}

func NewDealMap(maxsize int) *DealMap {
	buckets := minBuckets(maxsize)
	return &DealMap{size: 0, buckets: buckets, table: make([]DealTiny, buckets)}
}

// TODO rehash策略
func (m *DealMap) Put(d DealTiny) {
	num_probes, bucket_count_minus_one := 0, m.buckets-1
	bucknum := hashInt32(int(d.Dealid)) & bucket_count_minus_one
	for {
		if m.table[bucknum].Dealid == 0 { // insert, 不支持放入ID为0的Deal
			m.size += 1
			m.table[bucknum] = d
			return
		}
		if m.table[bucknum].Dealid == d.Dealid { // update
			m.table[bucknum] = d
			return
		}
		num_probes += 1 // Open addressing with Linear probing
		bucknum = (bucknum + num_probes) & bucket_count_minus_one
	}
}

func (m *DealMap) Get(id int32) (DealTiny, bool) {
	num_probes, bucket_count_minus_one := 0, m.buckets-1
	bucknum := hashInt32(int(id)) & bucket_count_minus_one
	for {
		if m.table[bucknum].Dealid == id {
			return m.table[bucknum], true
		}
		if m.table[bucknum].Dealid == 0 {
			return m.table[bucknum], false
		}
		num_probes += 1
		bucknum = (bucknum + num_probes) & bucket_count_minus_one
	}
}

func main() {
	dm := NewDealMap(SIZE)
	for i := 0; i < SIZE; i++ {
		dm.Put(DealTiny{Dealid: int32(i), Classid: int32(200)})
	}
	dealres, boo := dm.Get(1)
	fmt.Println(dealres, boo)
}
</pre>
###Nginx 小技巧
负载均衡设置
<pre>
upstream wishome_backend {
    server mawenbao.com:9001;
    server mawenbao.com:9002;
}

server {
    server_name test.mawenbao.com;

    location / {
        proxy_set_header X-Real-Ip $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://wishome_backend;
    }
}
</pre>
通过上面的配置，访问test.mawenbao.com的请求将被平均分配到9001和9002两个端口。

proxy_set_header设置的两个http头X-Real-Ip和X-Forwarded-For用于记录访问者的原始ip地址，其中X-Real-Ip只是一个ip，而X-Forwarded-For是一系列逗号分割的ip列表，第一个是访问者的ip，其后都是转发服务器的ip地址。

访问txt文件时提示下载

txt文件的MIME类型为text/plain，使用浏览器访问时默认行为是直接在浏览器中显示。如果需要将默认行为改为直接下载，可以在nginx配置文件中添加如下规则即可。
<pre>
location ~* \.(txt) {
  add_header Content-Disposition "attachment";
}
<pre>
对于某些特殊的文件，如果在访问时需要直接在浏览器上显示文件内容，则可使用如下规则，以gpg的asc加密文件为例。
<pre>
location ~* \.(asc) {
  default_type text/plain;
  add_header Content-Disposition "inline";
}
</pre>
###Golang TCP交互封装
<pre>
package main

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync"
	"time"

	"fmt"
	"os"
)

//////////////////TCP交互封装//////////////////////

//支持的最大消息长度
const maxLength int = 1<<32 - 1 // 4294967295

var (
	rHeadBytes = [4]byte{0, 0, 0, 0}
	wHeadBytes = [4]byte{0, 0, 0, 0}
	errMsgRead = errors.New("Message read length error")
	errHeadLen = errors.New("Message head length error")
	errMsgLen  = errors.New("Message length is no longer in normal range")
)
var connPool sync.Pool

//从对象池中获取一个对象,不存在则申明
func Newconnection(conn net.Conn) Conn {
	c := connPool.Get()
	if cnt, ok := c.(*connection); ok {
		cnt.rwc = conn
		return cnt
	}
	return &connection{rlen: 0, rwc: conn}
}

type Conn interface {
	Read() (r io.Reader, size int, err error)
	Write(p []byte) (n int, err error)
	Writer(size int, r io.Reader) (n int64, err error)
	RemoteAddr() net.Addr
	LocalAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	Close() (err error)
}

//定义一个结构体,用来封装Conn.
type connection struct {
	rlen  int        //消息长度
	rwc   net.Conn   //原始的网络链接
	rlock sync.Mutex //Conn读锁
	wlock sync.Mutex //Conn写锁
}

//此方法用来读取头部消息.
func (self *connection) rhead() error {
	n, err := self.rwc.Read(rHeadBytes[:])
	if n != 4 || err != nil {
		if err != nil {
			return err
		}
		return errHeadLen
	}
	self.rlen = int(binary.BigEndian.Uint32(rHeadBytes[:]))
	return nil
}

//此方法用来发送头消息
func (self *connection) whead(l int) error {
	if l <= 0 || l > maxLength {
		return errMsgLen
	}
	binary.BigEndian.PutUint32(wHeadBytes[:], uint32(l))
	_, err := self.rwc.Write(wHeadBytes[:])
	return err
}

//头部消息解析之后.返回一个io.Reader借口.用来读取远端发送过来的数据
//封装成limitRead对象.来实现ioReader接口
func (self *connection) Read() (r io.Reader, size int, err error) {
	self.rlock.Lock()
	if err = self.rhead(); err != nil {
		self.rlock.Unlock()
		return
	}
	size = self.rlen
	r = limitRead{r: io.LimitReader(self.rwc, int64(size)), unlock: self.rlock.Unlock}
	return
}

//发送消息前先调用whead函数,来发送头部信息,然后发送body
func (self *connection) Write(p []byte) (n int, err error) {
	self.wlock.Lock()
	err = self.whead(len(p))
	if err != nil {
		self.wlock.Unlock()
		return
	}
	n, err = self.rwc.Write(p)
	self.wlock.Unlock()
	return
}

//发送一个流.必须指定流的长度
func (self *connection) Writer(size int, r io.Reader) (n int64, err error) {
	self.wlock.Lock()
	err = self.whead(int(size))
	if err != nil {
		self.wlock.Unlock()
		return
	}
	n, err = io.CopyN(self.rwc, r, int64(size))
	self.wlock.Unlock()
	return
}

func (self *connection) RemoteAddr() net.Addr {
	return self.rwc.RemoteAddr()
}

func (self *connection) LocalAddr() net.Addr {
	return self.rwc.LocalAddr()
}

func (self *connection) SetDeadline(t time.Time) error {
	return self.rwc.SetDeadline(t)
}

func (self *connection) SetReadDeadline(t time.Time) error {
	return self.rwc.SetReadDeadline(t)
}

func (self *connection) SetWriteDeadline(t time.Time) error {
	return self.rwc.SetWriteDeadline(t)
}

func (self *connection) Close() (err error) {
	err = self.rwc.Close()
	self.rlen = 0
	connPool.Put(self)
	return
}

type limitRead struct {
	r      io.Reader
	unlock func()
}

func (self limitRead) Read(p []byte) (n int, err error) {
	n, err = self.r.Read(p)
	if err != nil {
		self.unlock()
	}
	return n, err
}

//////////////////测试//////////////////////
func Dial() {
	conn, err := net.Dial("tcp", "127.0.0.1:1789")
	if err != nil {
		fmt.Println(err)
		return
	}
	c := Newconnection(conn)
	defer c.Close()
	c.Write([]byte("Test"))
	c.Write([]byte("Test"))

	r, size, err := c.Read()
	if err != nil {
		fmt.Println(err, size)
		return
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}
}

func Listener(proto, addr string) {
	lis, err := net.Listen(proto, addr)
	if err != nil {
		panic("Listen port error:" + err.Error())
		return
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			time.Sleep(1e7)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	c := Newconnection(conn)
	msgchan := make(chan struct{})
	defer c.Close()
	go func(ch chan struct{}) {
		<-msgchan
		f, _ := os.Open("t2.go")
		defer f.Close()
		info, _ := f.Stat()
		c.Writer(int(info.Size()), f)
		c.Close()
	}(msgchan)
	for {
		r, size, err := c.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		n, err := io.Copy(os.Stdout, r)
		if err != nil || n != int64(size) {
			if err == io.EOF {
				continue
			}
			fmt.Println("读取数据失败:", err)
			return
		}
		time.Sleep(2e9)
		msgchan <- struct{}{}
	}
}

func main() {
	go Listener("tcp", ":1789")
	time.Sleep(10e9)
	Dial()
}
</pre>
###Golang 基于Gob的通讯包
<pre>
package main

////////////基于Gob的tcp通讯用的包/////////////////////

import (
	"encoding/gob"
	"errors"
	"net"
	"reflect"
	"sync"
	"unsafe"

	"fmt"
	"time"
)

type message struct {
	Type  string
	value reflect.Value
}

func (self *message) Recovery() {
	putPointer(self.value)
	putMsg(self)
}

func (self *message) Interface() interface{} {
	return self.value.Interface()
}

/* 声明一个消息池用来重用对象 */

var msgPool sync.Pool

func getMsg() *message {
	if msg, ok := msgPool.Get().(*message); ok {
		return msg
	}
	return new(message)
}

func putMsg(msg *message) {
	msgPool.Put(msg)
}

type gobConnection struct {
	rwc   net.Conn
	enc   *gob.Encoder
	dec   *gob.Decoder
	rlock sync.Mutex
	wlock sync.Mutex
}

type GobConnection interface {
	Read() (msg *message, err error)
	Write(msg interface{}) (err error)
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

var gobPool sync.Pool

func NewGobConnection(conn net.Conn) GobConnection {
	if gcn, ok := gobPool.Get().(*gobConnection); ok {
		gcn.rwc = conn
		gcn.enc = gob.NewEncoder(conn)
		gcn.dec = gob.NewDecoder(conn)
		return gcn
	}
	return &gobConnection{rwc: conn, enc: gob.NewEncoder(conn), dec: gob.NewDecoder(conn)}
}

type msgStruct struct {
	StructName string
}

var (
	rheadMsg = msgStruct{}
	wheadMsg = msgStruct{}
)

func (self *gobConnection) Read() (msg *message, err error) {
	self.rlock.Lock()
	defer self.rlock.Unlock()

	err = self.dec.Decode(&rheadMsg)
	if err != nil {
		return
	}
	var typ reflect.Type
	typ, err = GetMsgType(rheadMsg.StructName)
	if err != nil {
		return
	}
	msg = getMsg()
	msg.Type = rheadMsg.StructName
	var value = getPointer(typ)
	err = self.dec.DecodeValue(value)
	if err != nil {
		msg.Recovery()
		return
	}
	msg.value = value
	return
}

func (self *gobConnection) Write(msg interface{}) (err error) {
	self.wlock.Lock()
	value := reflect.ValueOf(msg)
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Ptr {
		wheadMsg.StructName = value.Elem().Type().String()
	} else {
		wheadMsg.StructName = value.Type().String()
	}
	err = self.enc.Encode(wheadMsg)
	if err != nil {
		self.wlock.Unlock()
		return
	}
	err = self.enc.EncodeValue(value)
	self.wlock.Unlock()
	return
}

func (self *gobConnection) Close() error {
	self.enc = nil
	self.dec = nil
	err := self.rwc.Close()
	gobPool.Put(self)
	return err
}

func (self *gobConnection) LocalAddr() net.Addr {
	return self.rwc.LocalAddr()
}

func (self *gobConnection) RemoteAddr() net.Addr {
	return self.rwc.RemoteAddr()
}

/* 通过指定类型申请一个定长的内存. */

var (
	lock   sync.Mutex
	ptrMap = make(map[string]*sync.Pool)
)

func getPointer(typ reflect.Type) reflect.Value {
	p, ok := ptrMap[typ.String()]
	if ok {
		if value, ok := p.Get().(reflect.Value); ok {
			return value
		}
		return reflect.New(typ)
	}
	lock.Lock()
	ptrMap[typ.String()] = new(sync.Pool)
	lock.Unlock()
	return reflect.New(typ)
}

func putPointer(value reflect.Value) {
	elem := value.Elem().Type()
	p, ok := ptrMap[elem.String()]
	if !ok {
		lock.Lock()
		p = new(sync.Pool)
		ptrMap[elem.String()] = p
		lock.Unlock()
	}
	ClearData(elem.Size(), unsafe.Pointer(value.Pointer()))
	p.Put(value)
}

/* 使用此包进行数据发送之前必须将类型注册.否则接收方无法解包 */

var (
	typeMap   = make(map[string]reflect.Type)
	Errortype = errors.New("type not register")
)

func GetMsgType(name string) (reflect.Type, error) {
	typ, ok := typeMap[name]
	if ok {
		return typ, nil
	}
	return nil, Errortype
}

func GetMsgAllType() []string {
	list := make([]string, 0, len(typeMap))
	for name, _ := range typeMap {
		list = append(list, name)
	}
	return list
}

func RegisterType(typ reflect.Type) {
	typeMap[typ.String()] = typ
}

func DeleteType(name string) {
	delete(typeMap, name)
}

/* 清除固定长度的内存数据,使用方法是:指定内存开始地址和长度. 请勿随便使用.使用不当可能会清除有效数据 */

func ClearData(size uintptr, ptr unsafe.Pointer) {
	var temptr uintptr = uintptr(ptr)
	var step uintptr = 1
	for {
		if size <= 0 {
			break
		}
		switch {
		case 1 <= size && size < 8:
			step = 1
		case 8 <= size && size < 32:
			step = 8
		case 32 <= size && size < 64:
			step = 32
		case size >= 64:
			step = 64
		}
		clearData(step, unsafe.Pointer(temptr))
		temptr += step
		size -= step
	}
}

func clearData(size uintptr, ptr unsafe.Pointer) {
	switch size {
	case 1:
		*(*[1]byte)(ptr) = [1]byte{}
	case 8:
		*(*[8]byte)(ptr) = [8]byte{}
	case 32:
		*(*[32]byte)(ptr) = [32]byte{}
	case 64:
		*(*[64]byte)(ptr) = [64]byte{}
	}
}

//////////////////测试///////////////////////

type Info struct {
	Name string
	Age  int
	Job  string
	Hob  []string
}

type Test struct {
	Date    int
	Login   string
	Path    string
	Servers float64
	List    []string
	Dir     string
	Stream  bool
}

//初始化要发送的类型
func init() {
	go InitListen("tcp", ":2789")
	time.Sleep(1e9)
	RegisterType(reflect.TypeOf(Info{}))
	RegisterType(reflect.TypeOf(Test{}))
}
func main() {
	Test_rw()
	now := time.Now().Unix()
	Benchmark_rw()
	fmt.Println(time.Now().Unix())
	fmt.Println(now)
}
func Test_rw() {
	Dail("tcp", "127.0.0.1:2789", 1)
}

func Benchmark_rw() {
	Dail("tcp", "127.0.0.1:2789", 10000)
}

//创建tcp监听的端口
func InitListen(proto, addr string) {
	lis, err := net.Listen(proto, addr)
	if err != nil {
		fmt.Println("listen error,", err.Error())
		return
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("接入错误:", err)
			continue
		}
		go handle(conn)
	}
}

//链接处理逻辑
func handle(conn net.Conn) {
	con := NewGobConnection(conn)
	defer con.Close()
	for {
		msg, err := con.Read()
		if err != nil {
			fmt.Println(con.RemoteAddr())
			fmt.Println("服务端ReadError:", err)
			return
		}
		err = con.Write(msg.Interface())
		if err != nil {
			fmt.Println("服务端WriteError:", err)
			msg.Recovery()
			return
		}
		msg.Recovery()
	}
}

//创建连接.
func Dail(proto, addr string, count int) {
	con, err := net.Dial(proto, addr)
	if err != nil {
		fmt.Println("客户端连接错误:", err)
		return
	}
	conn := NewGobConnection(con)
	defer conn.Close()

	for i := 0; i < count; i++ {
		err = conn.Write(Info{"testing", 25, "IT", []string{"backetball", "football"}})
		if err != nil {
			fmt.Println("客户端WriteError:", err)
			return
		}
		msg, err := conn.Read()
		if err != nil {
			fmt.Println("客户端ReadError:", err)
			return
		}
		fmt.Println(msg, msg.Interface())
		msg.Recovery()
	}
}
</pre>
###Golang Context简单用法-类似sync.WaitGroup
<pre>
package main

import (
	"context"
	"fmt"
	"time"
)

////////////////Context简单用法-类似sync.WaitGroup//////////////////

func main() {
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	t, ok := ctx.Deadline()
	if ok {
		fmt.Println(time.Now())
		fmt.Println(t.String())
	}
	go func(ctx context.Context) {
		fmt.Println(ctx.Value("Test"))
		<-ctx.Done()
		fmt.Println(ctx.Err())
	}(ctx)

	if ctx.Err() == nil {
		time.Sleep(6e9)
	}
	if ctx.Err() != nil {
		fmt.Println("已经退出了")
	}
	cancelFunc()
}
</pre>
###Golang Golang的TLS通信,证书文件使用
<pre>
package main

//////////////Golang的TLS通信,证书文件使用////////////////

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

func main() {
	info := CertInformation{Country: []string{"中国"}, Organization: []string{"游戏蜗牛"},
		OrganizationalUnit: []string{"blog.csdn.net/fyxichen"}, EmailAddress: []string{"czxichen@163.com"},
		StreetAddress: []string{"中心大道西171"}, Province: []string{"江苏省·工业园区"}, SubjectKeyId: []byte{1, 2, 3, 4, 5, 6},
		Certificate: "client.crt", PrivateKey: "client.key", ROOTCertificate: "server.pem", ROOTPrivateKey: "server.key"}
	err := CreateCerts(info)
	if err != nil {
		println(err.Error())
	}
}

type CertInformation struct {
	Country            []string
	Organization       []string
	OrganizationalUnit []string
	EmailAddress       []string
	Province           []string
	StreetAddress      []string
	SubjectKeyId       []byte
	Certificate        string
	PrivateKey         string
	ROOTCertificate    string
	ROOTPrivateKey     string
}

func CreateCerts(info CertInformation) error {

	var rootPrivateKey *rsa.PrivateKey
	var rootcertificate *x509.Certificate
	var err error
	ca := newCertificate(info)
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)

	//读取根证书
	if info.Certificate != "" && info.PrivateKey != "" {
		rootcertificate, rootPrivateKey, err = parseCerts(info.ROOTCertificate, info.ROOTPrivateKey)
		if os.IsNotExist(err) {
			rootPrivateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
			rootcertificate = ca
			ca_b, err := x509.CreateCertificate(rand.Reader, rootcertificate, rootcertificate, &rootPrivateKey.PublicKey, rootPrivateKey)
			if err != nil {
				return err
			}
			err = write(info.ROOTCertificate, "CERTIFICATE", ca_b)
			if err != nil {
				return err
			}
			priv_b := x509.MarshalPKCS1PrivateKey(rootPrivateKey)
			err = write(info.ROOTPrivateKey, "PRIVATE KEY", priv_b)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	} else {
		rootcertificate = ca
		rootPrivateKey = priv
	}

	ca_b, err := x509.CreateCertificate(rand.Reader, ca, rootcertificate, &priv.PublicKey, rootPrivateKey)
	if err != nil {
		return err
	}
	err = write(info.Certificate, "CERTIFICATE", ca_b)
	if err != nil {
		return err
	}

	priv_b := x509.MarshalPKCS1PrivateKey(priv)
	err = write(info.PrivateKey, "PRIVATE KEY", priv_b)
	if err != nil {
		return err
	}

	return nil
}

func write(filename, Type string, p []byte) error {
	File, err := os.Create(filename)
	defer File.Close()
	if err != nil {
		return err
	}
	var b *pem.Block = &pem.Block{Bytes: p, Type: Type}
	err = pem.Encode(File, b)
	if err != nil {
		return err
	}
	return nil
}

func parseCerts(ROOTCertificate, ROOTPrivateKey string) (*x509.Certificate, *rsa.PrivateKey, error) {
	var rootPrivateKey *rsa.PrivateKey
	var rootcertificate *x509.Certificate

	buf, err := ioutil.ReadFile(ROOTCertificate)
	if err != nil {
		return nil, nil, err
	}
	p := &pem.Block{}
	p, buf = pem.Decode(buf)
	rootcertificate, err = x509.ParseCertificate(p.Bytes)
	if err != nil {
		return nil, nil, err
	}
	buf, err = ioutil.ReadFile(ROOTPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	p, buf = pem.Decode(buf)
	rootPrivateKey, err = x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return nil, nil, err
	}
	return rootcertificate, rootPrivateKey, nil
}

func newCertificate(info CertInformation) *x509.Certificate {
	return &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Country:            info.Country,
			Organization:       info.Organization,
			OrganizationalUnit: info.OrganizationalUnit,
			Province:           info.Province,
			StreetAddress:      info.StreetAddress,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		SubjectKeyId:          info.SubjectKeyId,
		BasicConstraintsValid: true,
		IsCA:           true,
		ExtKeyUsage:    []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:       x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		EmailAddresses: info.EmailAddress,
	}
}
</pre>
###Golang https通讯
<pre>
package main

import (
	"crypto/rand"
	"crypto/tls"
	"net"
)

///////////////tls 链接通信https//////////////////

func Servertls(addr, crt, key string) (net.Listener, error) {
	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, err
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader
	listener, err := tls.Listen("tcp", addr, &config)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func ClientTls(addr, crt, key string) (net.Conn, error) {
	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, err
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", addr, &config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
</pre>
###Golang http上传图片
<pre>
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":1789", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintln(w, "upload ok!")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tpl))
}

const tpl = `<html>
<head>
<title>上传文件</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
 <input type="file" name="uploadfile" />
 <input type="hidden" name="token" value="{...{.}...}"/>
 <input type="submit" value="upload" />
</form>
</body>
</html>`
</pre>
###Golang cgi CGI 公共网关接口
<pre>
package main

import (
	"log"
	"net/http"
	"net/http/cgi"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := new(cgi.Handler)
		handler.Path = "D:/gopath/src/test" + r.URL.Path //需要访问的文件地址，该文件必须是可以执行的
		log.Println(handler.Path)
		handler.Dir = "D:/gopath/src/test/"
		handler.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8989", nil))
}
</pre>
###Golang  TCP 聊天程序
<pre>
package main

//Server
import (
	"fmt"
	"net"
)

const (
	ip   = "localhost"
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

package main

//Client
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
###Golang  接口型函数
<pre>
package main

import (
	"fmt"
)

type Handler interface {
	Do(k, v interface{})
}
type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}
func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}
func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁.\n", k, v)
}
func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	EachFunc(persons, selfInfo)

}
</pre>
###Golang AES CBC模式加解密程序
<pre>
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

///////////加密代码 start/////////////
func Encrypt(plantText, key []byte) ([]byte, error) { //key的位数是16
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	plantText = PKCS7Padding(plantText, block.BlockSize())
	blockModel := cipher.NewCBCEncrypter(block, key)
	ciphertext := make([]byte, len(plantText))
	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

///////////加密代码 end/////////////

///////////解密代码 start///////////
func Decrypt(ciphertext, key []byte) ([]byte, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, keyBytes)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

//////////解密代码 end////////////

func main() {
	res, _ := Encrypt([]byte("ddd"), []byte("1000000000000000"))
	finalres, _ := Decrypt(res, []byte("1000000000000000"))
	println(string(finalres))
}
</pre>
###Golang之io.Pipe
<pre>
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	cli := http.Client{}

	msg := struct {
		Name, Addr string
		Price      float64
	}{
		Name:  "hello",
		Addr:  "beijing",
		Price: 123.56,
	}
	r, w := io.Pipe()
	// 注意这边的逻辑！！
	go func() {
		defer func() {
			time.Sleep(time.Second * 2)
			log.Println("encode完成")
			// 只有这里关闭了，Post方法才会返回
			w.Close()
		}()
		log.Println("管道准备输出")
		// 只有Post开始读取数据，这里才开始encode，并传输
		err := json.NewEncoder(w).Encode(msg)
		log.Println("管道输出数据完毕")
		if err != nil {
			log.Fatalln("encode json failed:", err)
		}
	}()
	time.Sleep(time.Second * 1)
	log.Println("开始从管道读取数据")
	resp, err := cli.Post("http://localhost:9999/json", "application/json", r)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("POST传输完成")

	body := resp.Body
	defer body.Close()

	if body_bytes, err := ioutil.ReadAll(body); err == nil {
		log.Println("response:", string(body_bytes))
	} else {
		log.Fatalln(err)
	}
}
</pre>
###Golang 
<pre>
package main

//golang 允许并发最多10万的线程
//处理图片的能力C是golang的1000倍
import (
	"fmt"
	"sync"
	"time"
)

func closure1() func() int {
	i := 0
	return func() int {
		i++ //该匿名函数引用了closure1函数中的i变量故该匿名函数与i变量形成闭包
		return i
	}
}

func main() {
	o := make(chan int)
	c := make(chan int)

	go func() {
		for {
			select {
			case a := <-c: //监听c管道只要一有数据进来 就打印出来
				fmt.Println(a)
			//这里After返回 <-chan Time 也就是监听 <-chan Time这个管道
			//如果超过5秒钟 如果select一直未收到消息 那么 就会给<-chan Time通道发送一个消息
			//每隔5秒就会发送一次
			case <-time.After(time.Second):
				o <- 0
				break //仅仅是跳出select循环并未跳出for循环
			}
		}
	}()
	//100000协程
	for i := 0; i < 100000; i++ {
		c <- i
	}
	<-o //接收消息

	fmt.Println("闭包：", closure1())

	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)
	fmt.Println(a, b)
}

// 要理解这个事儿首先得了解操作系统是怎么玩线程的。一个线程就是一个栈加一堆资源。操作系统一会让cpu跑线程A，一会让cpu跑线程B，靠A和B的栈来保存A和B的执行状态。每个线程都有他自己的栈。
// 但是线程又老贵了，花不起那个钱，所以go发明了goroutine。大致就是说给每个goroutine弄一个分配在heap里面的栈来模拟线程栈。比方说有3个goroutine，A,B,C，就在heap上弄三个栈出来。然后Go让一个单线程的scheduler开始跑他们仨。相当于 { A(); B(); C() }，连续的，串行的跑。
// 和操作系统不太一样的是，操作系统可以随时随地把你线程停掉，切换到另一个线程。这个单线程的scheduler没那个能力啊，他就是user space的一段朴素的代码，他跑着A的时候控制权是在A的代码里面的。A自己不退出谁也没办法。
// 所以A跑一小段后需要主动说，老大（scheduler），我不想跑了，帮我把我的所有的状态保存在我自己的栈上面，让我歇一会吧。这时候你可以看做A返回了。A返回了B就可以跑了，然后B跑一小段说，跑够了，保存状态，返回，然后C再跑。C跑一段也返回了。
// 这样跑完{A(); B(); C()}之后，我们发现，好像他们都只跑了一小段啊。所以外面要包一个循环，大致是：
// goroutine_list = [A, B, C]
// while(goroutine):
//   for goroutine in goroutine_list:
//     r = goroutine()
//     if r.finished():
//       goroutine_list.remove(r)
// 比如跑完一圈A，B，C之后谁也没执行完，那么就在回到A执行一次。由于我们把A的栈保存在了HEAP里，这时候可以把A的栈复制粘贴会系统栈里（我很确定真实情况不是这么玩的，会意就行），然后再调用A，这时候由于A是跑到一半自己说跳出来的，所以会从刚刚跳出来的地方继续执行。比如A的内部大致上是这样
// def A:
//   上次跑到的地方 = 找到上次跑哪儿了
//   读取所有临时变量
//   goto 上次跑到的地方
//   a = 1
//   print("do something")
//   go.scheduler.保存程序指针 // 设置"这次跑哪儿了"
//   go.scheduler.保存临时变量们
//   go.scheduler.跑够了_换人 //相当于return
//   print("do something again")
//   print(a)
// 第一次跑A，由于这是第一次，会打印do something，然后保存临时变量a，并保存跑到的地方，然后返回。再跑一次A，他会找到上次返回的地方的下一句，然后恢复临时变量a，然后接着跑，会打印“do something again"和1

// 所以你看出来了，这个关键就在于每个goroutine跑一跑就要让一让。一般支持这种玩意（叫做coroutine）的语言都是让每个coroutine自己说，我跑够了，换人。goroutine比较文艺的地方就在于，他可以来帮你判断啥时候“跑够了”。

// 其中有一大半就是靠的你说的“异步并发”。go把每一个能异步并发的操作，像你说的文件访问啦，网络访问啦之类的都包包好，包成一个看似朴素的而且是同步的“方法”，比如string readFile（我瞎举得例子）。但是神奇的地方在于，这个方法里其实会调用“异步并发”的操作，比如某操作系统提供的asyncReadFile。你也知道，这种异步方法都是很快返回的。
// 所以你自己在某个goroutine里写了
// string s = go.file.readFile("/root")

// 其实go偷偷在里面执行了某操作系统的API asyncReadFIle。跑起来之后呢，这个方法就会说，我当前所在的goroutine跑够啦，把刚刚跑的那个异步操作的结果保存下下，换人：
// // 实际上
// handler h = someOS.asyncReadFile("/root") //很快返回一个handler
// while (!h.finishedAsyncReadFile()): //很快返回Y/N
//   go.scheduler.保存现状()
//   go.scheduler.跑够了_换人() // 相当于return，不过下次会从这里的下一句开始执行
// string s = h.getResultFromAsyncRead()

// 然后scheduler就换下一个goroutine跑了。等下次再跑回刚才那个goroutine的时候，他就看看，说那个asyncReadFile到底执行完没有啊，如果没有，就再换个人吧。如果执行完了，那就把结果拿出来，该干嘛干嘛。所以你看似写了个同步的操作，已经被go替换成异步操作了。

// 还有另外一种情况是，某个goroutine执行了某个不能异步调用的会blocking的系统调用，这个时候goroutine就没法玩那种异步调用的把戏了。他会把你挪到一个真正的线程里让你在那个县城里等着，他接茬去跑别的goroutine。比如A这么定义
// def A:
//   print("do something")
//   go.os.InvokeSomeReallyHeavyAndBlockingSystemCall()
//   print("do something 2")
// golang 会帮你转成
// def 真实的A:
//   print("do something")
//   Thread t = new Thread( () => {
//     SomeReallyHeavyAndBlockingSystemCall();
//   })
//   t.start()
//   while !t.finished():
//     go.scheduler.保存现状
//     go.scheduler.跑够了_换人
//   print("finished")
// 所以真实的A还是不会blocking，还是可以跟别的小伙伴(goroutine)愉快地玩耍（轮流往复的被执行），但他其实已经占了一个真是的系统线程了。

// 当然会有一种情况就是A完全没有调用任何可能的“异步并发”的操作，也没有调用任何的同步的系统调用，而是一个劲的用CPU做运算（比如用个死循环调用a++）。在早期的go里，这个A就把整个程序block住了。后面新版本的go好像会有一些处理办法，比如如果你A里面call了任意一个别的函数的话，就有一定几率被踢下去换人。好像也可以自己主动说我要换人的，可以去查查新的go的spec
// 另外，请不要在意语言细节，技术细节,会意即可。 
</pre>
###Linux 查看端口被哪个程序占用
<pre>
// lsof -i :port_number |grep "LISTEN"
lsof -i :8080   //查看8080端口被哪个程序占用

//linux mysql 更改用户[root]密码
mysql> UPDATE mysql.user SET password =password('newpwd') WHERE user='root';
mysql> FLUSH PRIVILEGES;

//使用密码登录
mysql>mysql -u root -p       //以root用户身份登录
mysql>ENTER Password

//允许远程IP连接mysql数据库 
mysql>GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '111111' WITH GRANT OPTION;
mysql>FLUSH PRIVILEGES;
上句话的意思就是使用root在任意一台计算机上面以密码“111111”来连接，你如果在远程计算机上面使用密码“123”是无法连接的，包括你在本地使用mysql -uroot -p 密码为111111也无法连接。
当然执行上面一句SQL我们还需要FLUSH下缓存区，使之生效。
  
//以M为单位显示磁盘使用量和占用率
df -m

//列出home目录下所有文件或目录占用的大小，以KB作为计量单位
du -k /home
</pre>
###Golang json解析float64类型数据 bool类型数据
<pre>
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Record struct {
	A json.Number `json:"a"`
	B json.Number `json:"b"`
}

func main() {
	var (
		data   = `{"a":0.058,"b":2.060}`
		record Record
	)

	dec := json.NewDecoder(strings.NewReader(data))
	dec.UseNumber()
	if err := dec.Decode(&record); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("a = %s ; b = %s \n", record.A, record.B)
}

/////////////////////////////

package main

import (
	"encoding/json"
	"log"
)

type A struct {
	B *bool `json:"b"`
}

func main() {
	var a A
	str := `{"b":true}`
	json.Unmarshal([]byte(str), &a)
	if a.B != nil {
		log.Println(*a.B)
	}
}
</pre>
###Golang安全并发取不重复的随机字符串,自定义长度
<pre>
package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	b := make([]byte, 45)    //自定义的长度
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		fmt.Errorf("Could not successfully read from the system CSPRNG.")
	}
	fmt.Println(hex.EncodeToString(b))
}
</pre> 
###Golang拼接string方法性能比较
<pre>
package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

var way map[int]string

func benchmarkStringFunction(n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer

	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch index {
		case 0: // fmt.Sprintf
			s = fmt.Sprintf("%s[%s]", s, v)
		case 1: // string +
			s = s + "[" + v + "]"
		case 2: // strings.Join
			s = strings.Join([]string{s, "[", v, "]"}, "")
		case 3: // stable bytes.Buffer
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}

	}
	d = time.Since(t0)
	if index == 3 {
		s = buf.String()
	}
	fmt.Printf("string len: %d\t", len(s))
	fmt.Printf("time of [%s]=\t %v\n", way[index], d)
	return d
}

func main() {
	way = make(map[int]string, 5)
	way[0] = "fmt.Sprintf"
	way[1] = "+"
	way[2] = "strings.Join"
	way[3] = "bytes.Buffer"

	k := 4
	d := [5]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkStringFunction(10000, i)
	}
}

package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	strs = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}
)

func TestStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(strs, "")
	}
}

func TestStringsPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < len(strs); j++ {
			s += strs[j]
		}
	}
}

func TestBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		for j := 0; j < len(strs); j++ {
			b.WriteString(strs[j])
		}
	}
}

func main() {
	fmt.Println("strings.Join:")
	fmt.Println(testing.Benchmark(TestStringsJoin))
	fmt.Println("bytes.Buffer:")
	fmt.Println(testing.Benchmark(TestBytesBuffer))
	fmt.Println("+:")
	fmt.Println(testing.Benchmark(TestStringsPlus))
}
</pre>
###Golang 通过身份证号计算年龄/年纪
<pre>
func GetAgeByIdCard(idCard string) int {
	// 仅支持18位身份证
	var mapmou = map[string]int{"January": 1, "february": 2, "March": 3, "April": 4, "May": 5, "June": 6, "July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12}
	now := time.Now()
	now_year := now.Year()                 // 年
	now_mo := mapmou[now.Month().String()] // 月
	now_day := now.Day()                   // 日
	idcard_year, _ := strconv.Atoi(Substr(idCard, 6, 4)) // 年
	idcard_mo, _ := strconv.Atoi(Substr(idCard, 10, 2))  // 月
	idcard_day, _ := strconv.Atoi(Substr(idCard, 12, 2)) // 日
	age := now_year - idcard_year // 如果计算虚岁需这样：age := now_year - idcard_year+1
	if now_year < idcard_year {
		age = 0
	} else {
		if now_mo < idcard_mo {
			age = age - 1
		} else if now_mo == idcard_mo{
			if now_day < idcard_day {
				age = age - 1
			}
		}
	}
	return age
}
</pre>
###Golang json 解析key为数字的数据
<pre>
type Examp struct {
    Num1 string `json:"1"`
	Num2 string `json:"2"`
}
</pre>
###启动svn服务 
<pre>
svnserve -d -r  /home/svndata
</pre>
###Golang 用缓存限制每日提交次数
<pre>
func Limit10TimesPerDay(key string) int {
	count := 10  //限制10次
	if Rc.IsExist(key) {
		count, _ = Rc.RedisInt(key)
	}
	count--
	Rc.Put(key, count, GetTodayLastSecond())
	return count
}

func GetTodayLastSecond() time.Duration {
	today := GetToday() + " 23:59:59"
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", today, time.Local)
	return time.Duration(end.Unix()-time.Now().Local().Unix()) * time.Second
}
</pre>
###Golang json解析时候的转义问题
<pre>
type Query struct {
	AppID     string `json:"AppID"`
	Timestamp int64  `json:"Timestamp"`
	Package   string `json:"Package"`
}

func MarshalDemo() {
	v := &Query{}
	v.AppID = "testid"
	v.Timestamp = time.Now().Unix()
	v.Package = "xxcents=100&bank=666"

	data, _ := json.Marshal(v)
	fmt.Println("Golang在解析JSON时需注意的:处理前Marshal:", string(data))

	data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)
	data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)
	data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)
	data = bytes.Replace(data, []byte("\\u003d"), []byte("="), -1)

	fmt.Println("Golang在解析JSON时需注意的:处理后Marshal:", string(data))
}
output==>
Golang在解析JSON时需注意的:处理前Marshal: {"AppID":"testid","Timestamp":1484568598,"Package":"xxcents=100\u0026bank=666"}
Golang在解析JSON时需注意的:处理后Marshal: {"AppID":"testid","Timestamp":1484568598,"Package":"xxcents=100&bank=666"}
</pre>
### Beego框架JSON默认加密
<pre>
//BaseController.go

package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	if c.Ctx.Input.Method() != "GET" && c.Ctx.Input.Method() != "HEAD" && !c.Ctx.Input.IsUpload() {
		c.Ctx.Input.RequestBody = utils.DesBase64Decrypt(c.Ctx.Input.RequestBody)
	}
}

func (c *BaseController) ServeJSON(encoding ...bool) {
	var (
		hasIndent   = true
		hasEncoding = false
	)
	if beego.BConfig.RunMode == beego.PROD  {
		hasIndent = false
	}
	if len(encoding) > 0 && encoding[0] == true {
		hasEncoding = true
	}
	c.JSON(c.Data["json"], hasIndent, hasEncoding)
}

// json writes json to response body.
// if coding is true , it converts utf-8 to \u0000 type.
func (c *BaseController) JSON(data interface{}, hasIndent bool, coding bool) error {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	var content []byte
	var err error
	if hasIndent {
		content, err = json.MarshalIndent(data, "", "  ")
	} else {
		content, err = json.Marshal(data)
	}
	if err != nil {
		http.Error(c.Ctx.Output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	if coding {
		content = []byte(utils.StringsToJSON(string(content)))
	}
	return c.Ctx.Output.Body(自定义加密方法(content))
}
</pre>
###Golang 顺序存储 线性表
<pre>
package main

//顺序存储线性表
import (
	"fmt"
)

const MAXSIZE = 20

type List struct {
	Element [MAXSIZE]int
	Length  int
}

func (l *List) InitList(value int, position int) {
	l.Element[position] = value
	l.Length++
}

func (l *List) Insert(value, position int) bool {
	if position < 0 || position >= MAXSIZE || l.Length >= MAXSIZE {
		return false
	}
	if position < l.Length {
		for k := l.Length - 1; k >= position; k-- {
			l.Element[k+1] = l.Element[k]
		}
		l.Element[position] = value
		l.Length++
		return true
	} else {
		l.Element[l.Length] = value
		l.Length++
		return true
	}
}

func (l *List) Delete(position int) bool {
	if position < 0 || position > l.Length || position >= MAXSIZE {
		return false
	}
	for ; position < l.Length-1; position++ {
		l.Element[position] = l.Element[position+1]
	}
	l.Element[l.Length-1] = 0
	return true
}

func main() {
	var L List
	i := 0
	b := 1
	for i < 15 {
		L.InitList(b, i)
		i++
		b++
	}
	L.Insert(100, 16)
	L.Insert(200, 17)
	fmt.Println(L.Length)
	fmt.Println(L.Element)
}
</pre>
###Golang 生成token方法
<pre>
//生成token
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

func main(){
	print(GetRandom(64))
}
</pre>
###Golang 访问http /https post方法
<pre>
package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

// 访问https协议的接口
func HttpsPost(url, params string) ([]byte, error) {
	body := ioutil.NopCloser(strings.NewReader(params))
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	resp, err := client.Do(req) 
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() 
	return ioutil.ReadAll(resp.Body)
}

func HttpPost(url, params string) ([]byte, error) {
	body := ioutil.NopCloser(strings.NewReader(params))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req) 
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() 

	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}
</pre>
###Golang 压缩静态文件工具
<pre>
package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
    //需要压缩的文件路径
	paths := []string{`/home/go/src/p_temp/static/js`, `/home/go/src/pragram/static`}
	ExampleNewWatcher(paths)
}

var eventTime = make(map[string]int64)

func ExampleNewWatcher(paths []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					iscompress := true
					if !checkIfWatchExt(event.Name) {
						continue
					}
					mt := getFileModTime(event.Name)
					if t := eventTime[event.Name]; mt == t {
						//fmt.Println("[SKIP] # %s #", event.String())
						iscompress = false
					}
					//fmt.Println("file:", event.Name, iscompress, event.Op, fsnotify.Write, event.Op&fsnotify.Write)
					if iscompress {
						fmt.Println("modified file:", event.Name)
						var err1 error
						if strings.HasSuffix(event.Name, ".js") {
							icmd := exec.Command("uglifyjs", event.Name, "-m", "-o", event.Name)
							err1 = icmd.Run()
						} else if strings.HasSuffix(event.Name, ".css") {
							icmd := exec.Command("cleancss", "-o", event.Name, event.Name)
							err1 = icmd.Run()
						}

						if err1 != nil {
							fmt.Println(err1.Error())
						} else {
							eventTime[event.Name] = time.Now().Unix()
						}
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()
	for _, path := range paths {
		fmt.Println("[TRAC] Directory( %s )", path)
		err = watcher.Add(path)
		if err != nil {
			fmt.Println("[ERRO] Fail to watch directory[ %s ]", err.Error())
			//os.Exit(2)
		}
	}
	<-done
}

var watchExts = []string{".js", ".css"}

// checkIfWatchExt returns true if the name HasSuffix <watch_ext>.
func checkIfWatchExt(name string) bool {
	for _, s := range watchExts {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

// getFileModTime retuens unix timestamp of `os.File.ModTime` by given path.
func getFileModTime(path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	fmt.Println(path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("[ERRO] Fail to open file[ %s ]", err.Error())
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		fmt.Println("[ERRO] Fail to get file information[ %s ]", err.Error())
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
</pre>
###Golang 时间格式化到毫秒
<pre>
package main 

import (
	"fmt"
	"time"
)
func main(){
	now := time.Now().Format("2006-01-02 15:04:05.99999")
	fmt.Println(now)
}
</pre>
###Golang 循环更新数据库数据
<pre>
func UpdateUsersUrid() {
	o := orm.NewOrm()
	type Id int
	var res []Id
	for {
		sql0 = `select id from table where id> limit 10000`
		o.Raw(sql0).QueryRows(&res)
		if len(res) > 0 {
			for i := 0; i < len(res); i++ {
				urid := uuid.NewUUID().Hex()
				sql := `update users_l set urid =? where id =? limit 1`
				_, err := o.Raw(sql, urid, res[i]).Exec()
				if err != nil {
					println("第" + strconv.Itoa(i+1) + "条数据出错啦~")
				} else {
					println("第" + strconv.Itoa(i+1) + "条数据OK~")
				}
			}
		} else {
			return
		}
	}
	email.SendEmail("更新数据", "所有都完成啦", "admin@qq.com")
}
</pre>
###Golang byte处理 
<pre>
package packet

import (
	"encoding/binary"
	"fmt"
)

func IsPow2(size uint32) bool{
	return (size&(size-1)) == 0
}

func SizeofPow2(size uint32) uint32{
	if IsPow2(size){
		return size
	}
	size = size -1
	size = size-1
	size = size | (size>>1)
	size = size | (size>>2)
	size = size | (size>>4)
	size = size | (size>>8)
	size = size | (size>>16)
	return size + 1
}

func GetPow2(size uint32) uint32{
	var pow2 uint32 = 0
	if !IsPow2(size) {
		size = (size << 1)
	}
	for size > 1 {
		pow2++
	}
	return pow2
}
const (
	Max_bufsize  uint32 = 32000 
	Max_string_len  uint32 = 32000
	Max_bin_len  uint32 = 32000
)

type ByteBuffer struct {
	buffer []byte
	datasize uint32
	capacity uint32
}


var (
	ErrMaxDataSlotsExceeded     = fmt.Errorf("bytebuffer: Max Buffer Size Exceeded")
	ErrInvaildData              = fmt.Errorf("bytebuffer: Invaild Data")
)


func NewBufferByBytes(bytes []byte,datasize uint32)(*ByteBuffer){
	return &ByteBuffer{buffer:bytes,datasize:datasize,capacity:(uint32)(cap(bytes))}
}


func NewByteBuffer(size uint32)(*ByteBuffer){
	if size == 0 {
		size = 64
	}else{
		size = SizeofPow2(size)
	}
	return &ByteBuffer{buffer:make([]byte,size),datasize:0,capacity:size}
}

func (this *ByteBuffer) Clone() (*ByteBuffer){
	b := make([]byte,this.capacity)
	copy(b[0:],this.buffer[:this.capacity])
	return &ByteBuffer{buffer:b,datasize:this.datasize,capacity:this.capacity}
}

func (this *ByteBuffer) Bytes()([]byte){
	return this.buffer
}

func (this *ByteBuffer) Len()(uint32){
	return this.datasize
}

func (this *ByteBuffer) Cap()(uint32){
	return this.capacity
}

func (this *ByteBuffer) expand(newsize uint32)(error){
	newsize = SizeofPow2(newsize)
	if newsize > Max_bufsize {
		return ErrMaxDataSlotsExceeded
	}
	//allocate new buffer
	tmpbuf := make([]byte,newsize)
	//copy data
	copy(tmpbuf[0:], this.buffer[:this.datasize])
	//replace buffer
	this.buffer = tmpbuf
	this.capacity = newsize
	return nil
}

func (this *ByteBuffer) buffer_check(idx,size uint32)(error){
	if idx+size > this.capacity {
		err := this.expand(idx+size)
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *ByteBuffer) SetByte(idx uint32,value byte){
	this.buffer[idx] = value
}

func (this *ByteBuffer) PutByte(idx uint32,value byte)(error){
	err := this.buffer_check(idx,1)
	if err != nil {
		return err
	}
	this.buffer[idx] = value
	this.datasize += 1
	return nil
}

func (this *ByteBuffer) SetUint16(idx uint32,value uint16) {
	binary.BigEndian.PutUint16(this.buffer[idx:idx+2],value)
}

func (this *ByteBuffer) PutUint16(idx uint32,value uint16)(error){
	err := this.buffer_check(idx,2)
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint16(this.buffer[idx:idx+2],value)
	this.datasize += 2
	return nil
}

func (this *ByteBuffer) SetUint32(idx uint32,value uint32){
	binary.BigEndian.PutUint32(this.buffer[idx:idx+4],value)
}

func (this *ByteBuffer) PutUint32(idx uint32,value uint32)(error){
	err := this.buffer_check(idx,4)//(uint32)(unsafe.Sizeof(value)))
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint32(this.buffer[idx:idx+4],value)
	this.datasize += 4
	return nil
}


func (this *ByteBuffer) PutString(idx uint32,value string)(error){
	sizeneed := (uint32)(4)
	sizeneed += (uint32)(len(value))
	err := this.buffer_check(idx,sizeneed)
	if err != nil {
		return err
	}

	//first put string len
	this.PutUint32(idx,(uint32)(len(value)))
	
	idx += 4
	//second put string
	copy(this.buffer[idx:],value[:len(value)])
	this.datasize += (uint32)(len(value))
	return nil
}

func (this *ByteBuffer) PutBinary(idx uint32,value []byte)(error){
	sizeneed := (uint32)(4)
	sizeneed += (uint32)(len(value))
	err := this.buffer_check(idx,sizeneed)
	if err != nil {
		return err
	}

	//first put bin len
	this.PutUint32(idx,(uint32)(len(value)))
	idx += 4
	//second put bin
	copy(this.buffer[idx:],value[:len(value)])
	this.datasize += (uint32)(len(value))
	return nil
}

func (this *ByteBuffer) PutRawBinary(value []byte)(error){
	sizeneed := (uint32)(len(value))
	err := this.buffer_check(uint32(this.datasize),sizeneed)
	if err != nil {
		return err
	}
	//second put bin
	copy(this.buffer[this.datasize:],value[:len(value)])
	this.datasize += (uint32)(len(value))
	return nil
}

func (this *ByteBuffer) Uint16(idx uint32)(ret uint16,err error){
	if idx + 2 > this.datasize {
		ret = 0
		err = ErrInvaildData
		return
	}
	ret = binary.BigEndian.Uint16(this.buffer[idx:idx+2])
	err = nil
	return
}

func (this *ByteBuffer) Uint32(idx uint32)(ret uint32,err error){
	if idx + 4 > this.datasize {
		ret = 0
		err = ErrInvaildData
		return
	}
	ret = binary.BigEndian.Uint32(this.buffer[idx:idx+4])
	err = nil
	return
}

func (this *ByteBuffer) String(idx uint32)(ret string,err error){
	if idx + 4 > this.datasize {
		err = ErrInvaildData
		return
	}
	var bin []byte
	bin,err = this.Binary(idx)
	if err != nil {
		return
	}
	ret = string(bin)
	return
}

func (this *ByteBuffer) Binary(idx uint32)(ret []byte,err error){
	if idx + 4 > this.datasize {
		err = ErrInvaildData
		return
	}
	var bin_len uint32
	//read bin len
	bin_len,err = this.Uint32(idx)
	if err != nil {
		return
	}
	idx += 4
	if idx + bin_len > this.datasize {
		err = ErrInvaildData
		return
	}
	err = nil
	ret = this.buffer[idx:idx+bin_len]
	return
}
</pre>
###Golang []byte to int64
<pre>
package main 

import "encoding/binary"

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8) // int64 is 8 byte
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buf))
}
</pre>
###json 基础知识
json 中的所有数字类型的数据都是 float64 类型
###Golang RSA 公钥与私钥的格式转换
<pre>
package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/astaxie/beego"
	"net"
)

// LocalIP 获取本机IP
func LocalIP() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return ""
}

func Bytes2RSAPrivateKey(priKey []byte) interface{} {
	block, _ := pem.Decode(priKey)
	if block == nil {
		fmt.Println("Sign private key decode error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	return privateKey
}

func Bytes2RSAPublicKey(pubKey []byte) interface{} {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		fmt.Println("Sign pubilc key decode error")
	}
	pubilcKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	return pubilcKey
}

func main() {
	fmt.Println(LocalIP())
	beego.Emergency(Bytes2RSAPrivateKey(PRIVATE_KEY).(*rsa.PrivateKey))
	beego.Emergency(Bytes2RSAPrivateKey(Public_KEY).(*rsa.PublicKey))
}
</pre>
###Golang hmac sha1 加密
<pre>
package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha1"
	"io"
)
func main(){
 	h := sha1.New()
	io.WriteString(h, "jason")
	fmt.Printf("%x\n", h.Sum(nil))

	key := []byte("1212")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("jason"))
	fmt.Printf("%x\n", mac.Sum(nil))
}
</pre>
###Beego 定时任务
<pre>
package main 

import(
	"github.com/astaxie/beego/toolbox"
)
func TaskInfo() error {
	beego.Emergency("this is task......")
	return nil
}

func DoTask() {
	// 秒 分 时 天 月 年
	todotk := toolbox.NewTask("todotk", "0 00 14 * * *", TaskInfo)
	toolbox.AddTask("todotk", todotk)
	toolbox.StartTask()
}

func main(){
	DoTask()
}
</pre>
###Golang csv excel 
<pre>
func EXCEL() bool {
	f, err := os.Create("D://TestCSV.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write([]string{"Order_number", "Trade_number", "Account", "Pname", "Pay_money", "Create_date", "Capital"})
	for i := 0; i < 10; i++ {
		w.Write([]string{"Order_number", "Trade_number", "Account", "Pname", "Pay_money", "Create_date", "Capital"})
	}
	w.Flush()
	return true
}
</pre>
###Redis 基本操作
<pre>
redis-server ./redis.conf
#如果更改了端口，使用`redis-cli`客户端连接时，也需要指定端口，例如：
redis-cli -p 6380
#检测6379端口是否在监听
netstat -lntp | grep 6379
#使用客户端
redis-cli shutdown
</pre>
###解决‘Linux提示命令找不到’的问题
如果新装的系统，运行一些很正常的诸如：shutdown，fdisk的命令时，悍然提示：bash:command not found。那么 

首先就要考虑root 的$PATH里是否已经包含了这些环境变量。 

主要是这四个：/bin ,/usr/bin,/sbin,/usr/sbin。 

四个主要存放的东东： 

./bin: 

bin为binary的简写主要放置一些系统的必备执行档例如:cat、cp、chmod df、dmesg、gzip、kill、ls、mkdir、more、mount、rm、su、tar等。 

/usr/bin: 
主要放置一些应用软体工具的必备执行档例如c++、g++、gcc、chdrv、diff、dig、du、eject、elm、free、gnome*、 gzip、htpasswd、kfm、ktop、last、less、locale、m4、make、man、mcopy、ncftp、newaliases、nslookup passwd、quota、smb*、wget等。 

/sbin: 
主要放置一些系统管理的必备程式例如:cfdisk、dhcpcd、dump、e2fsck、fdisk、halt、ifconfig、ifup、 ifdown、init、insmod、lilo、lsmod、mke2fs、modprobe、quotacheck、reboot、rmmod、 runlevel、shutdown等。 

/usr/sbin: 
放置一些网路管理的必备程式例如:dhcpd、httpd、imap、in.*d、inetd、lpd、named、netconfig、nmbd、samba、sendmail、squid、swap、tcpd、tcpdump等。 
###Golang xls excel表格生成实例
<pre>
package controllers

import (
	"github.com/tealeg/xlsx"
)

//生成excel表格导出
func ToExcel(){
	RES, _ := models.FUNC(uid, condition, "")
	if Tradelist1 != nil {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	for i := 0; i <= len(RES); i++ {
		row = sheet.AddRow()
		if i == 0 { //创建表
			cell = row.AddCell()
			cell.Value = "序号"
			cell = row.AddCell()
			cell.Value = "部门"
			cell = row.AddCell()
			cell.Value = "业务主管"
			cell = row.AddCell()
			cell.Value = "客户专员"
			cell = row.AddCell()
			cell.Value = "客户姓名"
			cell = row.AddCell()
			cell.Value = "性别"
			cell = row.AddCell()
			cell.Value = "身份证号"
			cell = row.AddCell()
			cell.Value = "联系电话"
			cell = row.AddCell()
			cell.Value = "产品名称"
			cell = row.AddCell()
			cell.Value = "合同编号"
			cell = row.AddCell()
			cell.Value = "投资日期"
			cell = row.AddCell()
			cell.Value = "结算日期"
			cell = row.AddCell()
			cell.Value = "投资额度/元"
			cell = row.AddCell()
			cell.Value = "投资期限/天"
		} else {
			cell = row.AddCell()
			cell.Value = strconv.Itoa(i)
			cell = row.AddCell()
			orgname, _ := models.FUNC(UID[i-1].Uid)
			if orgname != "" {
				cell.Value = orgname
			}
			directorName, salemanName, _ := models.FUNC(Tradelist1[i-1].UID)
			cell = row.AddCell()
			cell.Value = directorName
			cell = row.AddCell()
			cell.Value = salemanName
			cell = row.AddCell()
			cell.Value = FUNC[i-1].FUNC
			cell = row.AddCell()
			if Tradelist1[i-1].Sex == "F" {
				cell.Value = "女"
			} else {
				cell.Value = "男"
			}
			cell = row.AddCell()
			cell.Value = Tradelist1[i-1].Id_card
			cell = row.AddCell()
			cell.Value = Tradelist1[i-1].Account
			cell = row.AddCell()
			cell.Value = Tradelist1[i-1].Pname
			cell = row.AddCell()
			TadeView, _ := models.FUNC(FUNC[i-1].Uid, UID[i-1].Id)
			if TadeView != nil {
				cell.Value = TadeView.Serial_number
			}
			cell = row.AddCell()
			cell.Value = Tradelist1[i-1].Create_date
			cell = row.AddCell()
			cell.Value = Tradelist1[i-1].Endtime
			cell = row.AddCell()
			cell.Value = Tradelist1[i-1].Capital
			cell = row.AddCell()
			cell.Value = strconv.Itoa(Tradelist1[i-1].Period)
		}
	}

	excelPath := "./static/deriveExcel/" + userid + ".xls"
	err = file.Save(excelPath)
	if err != nil {
		this.Data["json"] = ""
		fmt.Printf(err.Error())
	} else {
		this.Data["json"] = excelPath
		fmt.Println("insert success!")
	}
}
</pre>
###Golang 异或加密
<pre>
package main

import (
	"fmt"
	"strconv"
)

//简单的加密字符串方法-> 异或加密算法
var XorKey []byte = []byte{0xB2, 0x09, 0xBB, 0x55, 0x93, 0x6D, 0x44, 0x47}

type Xor struct {
}
type m interface {
	enc(src string) string
	dec(src string) string
}

func (a *Xor) enc(src string) string {
	var result string
	j := 0
	s := ""
	bt := []rune(src)
	for i := 0; i < len(bt); i++ {
		s = strconv.FormatInt(int64(byte(bt[i])^XorKey[j]), 16)
		if len(s) == 1 {
			s = "0" + s
		}
		result = result + (s)
		j = (j + 1) % 8
	}
	return result
}

func (a *Xor) dec(src string) string {
	var result string
	var s int64
	j := 0
	bt := []rune(src)
	for i := 0; i < len(src)/2; i++ {
		s, _ = strconv.ParseInt(string(bt[i*2:i*2+2]), 16, 0)
		result = result + string(byte(s)^XorKey[j])
		j = (j + 1) % 8
	}
	return result
}
func main() {
	xor := Xor{}
	fmt.Println(xor.enc("jason"))
	fmt.Println(xor.dec("d868c83afd"))
}
</pre>
###延时执行方法
<pre>
package main

import (
	"log"
	"time"
)

const (
	TIME_OUT_RUN_OK  int = 1 //运行完成
	TIME_OUT_RUN_OUT int = 0 //超时退出
)

var (
	_TIME_OUT_RUN_DEBUG = false
)

func SetDebug(debug_mode bool) {
	//设置调试模式:true调试模式
	_TIME_OUT_RUN_DEBUG = debug_mode
	return
}

func TimeOutRunApp(timeout uint64, callback func(v ...interface{}), v ...interface{}) {
	var single chan int = make(chan int, 1)

	//开启线程运行(参数必须采用v...,这样才是全参数传递,否则会出现[]interface{}问题)
	go timeOutRunApp(single, callback, v...)

	//阻塞等待运行结果或超时
	select {
	case <-single:
		//运行完成退出
		if _TIME_OUT_RUN_DEBUG {
			log.Println("run ok!")
		}
		close(single)

		return
	case <-time.After(time.Millisecond /*ms*/ * time.Duration(timeout)):
		//超时退出
		if _TIME_OUT_RUN_DEBUG {
			log.Println("run timeout!")
		}
		close(single)

		return
	}
}
func timeOutRunApp(single chan int, callback func(v ...interface{}), v ...interface{}) {
	callback(v...)
	single <- TIME_OUT_RUN_OK

	return
}

//=================测试样例
func Run(v ...interface{}) {
	//用于变量类型转换
	var convert interface{}

	//第一个参数
	convert = v[0]
	var args0 string = convert.(string)

	//第二个参数
	convert = v[1]
	var args1 string = convert.(string)

	//循环执行,flag=false:演示超时的情况,flag=true:演示非超时的情况
	var count uint32
	var flag bool = false
	for {
		time.Sleep(time.Millisecond * 1000)
		log.Println(args0, args1, count)
		count++

		if flag {
			return
		}
	}
	return
}

func TestTimeOutRun() {
	//设置调用超时5s,传入回调方法类App,及相关参数
	TimeOutRunApp(5000, Run, "利用interface实现超时执行某方法:", "count=")
}

func main() {
	TestTimeOutRun()
}
</pre>
###Golang  RSA加密 解密 生成公钥文件
<pre>
package rsa

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "flag"
    //"log"
    //"os"
    "bytes"
    "errors"
)

//创建公钥与私钥
func GenRsaKey(bits int) (string, string, error) {
    // 生成私钥文件
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return "", "", err
    }
    derStream := x509.MarshalPKCS1PrivateKey(privateKey)
    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: derStream,
    }

    w := bytes.NewBuffer([]byte(nil)) //bufio.NewWriter()
    err = pem.Encode(w, block)
    if err != nil {
        return "", "", err
    }
    prikey := string(w.Bytes())
    // 生成公钥文件
    publicKey := &privateKey.PublicKey
    derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        return "", "", err
    }
    block = &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: derPkix,
    }
    w2 := bytes.NewBuffer(nil) //bufio.NewWriter()
    err = pem.Encode(w2, block)
    if err != nil {
        return "", "", err
    }
    pubkey := string(w2.Bytes())
    return pubkey, prikey, nil
}

// 加密
func RsaEncrypt(origData, publicKey []byte) ([]byte, error) {
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
func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
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

func main(){
    //生成密钥
    var bits int
    flag.IntVar(&bits, "b", 1024, "密钥长度，默认为1024位")
    pub, pri, err := GenRsaKey(bits)
    if err != nil {
        t.Fatal("密钥文件生成失败！")
    }
    println(pub)
    t.Log("密钥文件生成成功！")

    source := "admin"

    //加密
    ds, err := RsaEncrypt([]byte(source), []byte(pub))
    if err != nil {
        t.Fatal(err.Error())
    }
    s1 := string(ds)
    //解密
    es, err := RsaDecrypt(ds, []byte(pri))
    if err != nil {
        t.Fatal(err.Error())
    }
    s2 := string(es)
    println("解密成功！")
}
</pre>
###接口验签
<pre>
验签：顾名思义，就是对接口的使用者身份的验证。对于常规的GET/POST/PUT请求中，GET是最容易暴露使用者信息的一种请求方式。如果在使用中，接口使用者的信息被任何一个第三人获知也是非常危险的，因此可以将使用者的唯一标示加动态时间戳经过MD5多次混合加密的方式进行封装数据，以在保证安全的前提下进行验签。
</pre>
###Golang aes+base64
<pre>
package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// aes + base64 加密的KEY 16位
const key_aes = 1234567887654321`


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
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//手机号加密，若解码出错，返回空字符串
func EncCode(acc string) (account string) {
	defer func() {
		err := recover()
		if err != nil {
			account = ""
		}
	}()
	accbyt := []byte(acc)
	accbyt = DesBase64Encrypt(accbyt)
	account = string(accbyt)
	return
}

func OriginAesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func OriginAesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

//aes3 + base64 encrypt
func AesBase64Encrypt(origData []byte) []byte {
	result, err := OriginAesEncrypt(origData, []byte(key_aes))
	if err != nil {
		panic(err)
	}
	return []byte(base64.StdEncoding.EncodeToString(result))
}

func AesBase64Decrypt(crypted []byte) []byte {
	result, _ := base64.StdEncoding.DecodeString(string(crypted))
	origData, err := OriginAesDecrypt(result, []byte(key_aes))
	if err != nil {
		panic(err)
	}
	return origData
}
</pre>
###Beego router 设置不同api地址
<pre>
// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"zcm_hwt/controllers"
	"zcm_hwt/controllers/api"
)

func init() {
	ROUTER1()
	ROUTER2()
}

func ROUTER1() {
	ns := beego.NewNamespace("/x1",

		beego.NSNamespace("/activity",
			beego.NSInclude(
				&controllers.ActivityController{},
			),
		),
	beego.AddNamespace(ns)
}

func ROUTER2() {
	ns2 := beego.NewNamespace("/cgi-bin",
		beego.NSNamespace("/api/users", 
			beego.NSInclude(
				&api.HwtsController{},
			),
		),
	)
	beego.AddNamespace(ns2)
}
</pre>