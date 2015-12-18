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