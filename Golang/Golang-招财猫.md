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
