##Go语言黑魔法--From达达
###达达 · 8 个月前
今天我要教大家一些无用技能，也可以叫它奇技淫巧或者黑魔法。用得好可以提升性能，用得不好就会招来恶魔，嘿嘿。

黑魔法导论

为了让大家在学习了基础黑魔法之后能有所悟，在必要的时候能创造出本文传授之外的属于自己的魔法，这里需要先给大家打好基础。

学习Go语言黑魔法之前，需要先看清Go世界的本质，你才能获得像Neo一样的能力。

在Go语言中，Slice本质是什么呢？是一个reflect.SliceHeader结构体和这个结构体中Data字段所指向的内存。String本质是什么呢？是一个reflect.StringHeader结构体和这个结构体所指向的内存。

在Go语言中，指针的本质是什么呢？是unsafe.Pointer和uintptr。

当你清楚了它们的本质之后，你就可以随意的玩弄它们，嘿嘿嘿。

第一式 - 获得Slice和String的内存数据

让我小试身手，你有一个CGO接口要调用，需要你把一个字符串数据或者字节数组数据从Go这边传递到C那边，比如像这个：mysql/conn.go at master · funny/mysql · GitHub

查了各种教程和文档，它们都告诉你要用C.GoString或C.GoBytes来转换数据。

但是，当你调用这两个函数的时候，发生了什么事情呢？这时候Go复制了一份数据，然后再把新数据的地址传给C，因为Go不想冒任何风险。

你的C程序只是想一次性的用一下这些数据，也不得不做一次数据复制，这对于一个性能癖来说是多麽可怕的一个事实！

这时候我们就需要一个黑魔法，来做到不拷贝数据又能把指针地址传递给C。
<pre>
// returns &s[0], which is not allowed in go
func stringPointer(s string) unsafe.Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(p.Data)
}

// returns &b[0], which is not allowed in go
func bytePointer(b []byte) unsafe.Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return unsafe.Pointer(p.Data)
}
</pre>
以上就是黑魔法第一式，我们先去到Go字符串的指针，它本质上是一个*reflect.StringHeader，但是Go告诉我们这是一个*string，我们告诉Go它同时也是一个unsafe.Pointer，Go说好吧它是，于是你得到了unsafe.Pointer，接着你就躲过了Go的监视，偷偷的把unsafe.Pointer转成了*reflect.StringHeader。

有了*reflect.StringHeader，你很快就取到了Data字段指向的内存地址，它就是Go保护着不想给你看到的隐秘所在，你把这个地址偷偷告诉给了C，于是C就愉快的偷看了Go的隐私。

第二式 - 把[]byte转成string

你肯定要笑，要把[]byte转成string还不简单？Go语言初学者都会的类型转换语法：string(b)。

但是你知道这么做的代价吗？既然我们能随意的玩弄SliceHeader和StringHeader，为什么我们不能造个string给Go呢？Go的内部会不会就是这么做的呢？

先上个实验吧：
<pre>
package labs28

import "testing"
import "unsafe"

func Test_ByteString(t *testing.T) {
	var x = []byte("Hello World!")
	var y = *(*string)(unsafe.Pointer(&x))
	var z = string(x)

	if y != z {
		t.Fail()
	}
}

func Benchmark_Normal(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i ++ {
		_ = string(x)
	}
}

func Benchmark_ByteString(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i ++ {
		_ = *(*string)(unsafe.Pointer(&x))
	}
}
</pre>
这个实验先证明了我们可以用[]byte的数据造个string给Go。接着做了两组Benchmark，分别测试了普通的类型转换和伪造string的效率。

结果如下：

<pre>
$ go test -bench="."
PASS
Benchmark_Normal    20000000            63.4 ns/op
Benchmark_ByteString    2000000000           0.55 ns/op
ok      github.com/idada/go-labs/labs28 2.486s
</pre>
哟西，显然Go这次又为了稳定性做了些复制数据之类的事情了！这让性能癖怎么能忍受！

我现在手头有个[]byte，但是我想用strconv.Atoi()把它转成字面含义对应的整数值，竟然需要发生一次数据拷贝把它转成string，比如像这样：mysql/types.go at master · funny/mysql · GitHub，这实在不能忍啊！

出招：
<pre>
// convert b to string without copy
func byteString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
</pre>
我们取到[]byte的指针，这次Go又告诉你它是*byte不是*string，你告诉它滚犊子这是unsafe.Pointer，Go这下又老实了，接着你很自在的把*byte转成了*string，因为你知道reflect.StringHeader和reflect.SliceHeader的结构体只相差末尾一个字段，两者的内存是对其的，没必要再取Data字段了，直接转吧。

于是，世界终于安宁了，嘿嘿。

第三式 - 结构体和[]byte互转

有一天，你想把一个简单的结构体转成二进制数据保存起来，这时候你想到了encoding/gob和encoding/json，做了一下性能测试，你想到效率有没有可能更高点？

于是你又试了encoding/binady，性能也还可以，但是你还不满意。但是瓶颈在哪里呢？你恍然大悟，最高效的办法就是完全不解析数据也不产生数据啊！

怎么做？是时候使用这个黑魔法了：
<pre>
type MyStruct struct {
	A int
	B int
}

var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}
</pre>
这是个曲折但又熟悉的故事。你造了一个SliceHeader，想把它的Data字段指向你的结构体，但是Go又告诉你不可以，你像往常那样把Go提到一边，你得到了unsafe.Pointer，但是这次Go有不死心，它告诉你Data是uintptr，unsafe.Pointer不是uintptr，你大脚把它踢开，怒吼道：unsafe.Pointer就是uintptr，你少拿这些概念糊弄我，Go屁颠屁颠的跑开了，现在你一马平川的来到了函数的出口，Go竟然已经在哪里等着你了！你上前三下五除二把它踢得远远的，顺利的把手头的SliceHeader转成了[]byte。

过了一阵子，你拿到了一个[]byte，你知道需要把它转成MyStruct来读取其中的数据。Go这时候已经完全不是你的对手了，它已经洗好屁股在函数入口等你，你一行代码就解决了它。

第四式 - 用CGO优化GC

你已经是Go世界的Neo，Go跟本没办法拿你怎么样。但是有一天Go的GC突然抽风了，原来这货是不管对象怎么用的，每次GC都给来一遍人口普查，导致系统暂停时间很长。

可是你是个性能癖，你把一堆数据都放在内存里方便快速访问，你这时候很想再踢Go的屁股，但是你没办法，毕竟你还在Go的世界里，你现在得替它擦屁股了，你似乎看到Go躲在一旁偷笑。

你想到你手头有CGO，可以轻易的用C申请到Go世界外的内存，Go的GC不会扫描这部分内存。

你还想到你可以用unsafe.Pointer将C的指针转成Go的结构体指针。于是一大批常驻内存对象被你用这种方式转成了Go世界的黑户，Go的GC一下子轻松了下来。

但是你手头还有很多Slice，于是你就利用C申请内存给SliceHeader来构造自己的Slice，于是你旗下的Slice纷纷转成了Go世界的黑户，Go的GC终于平静了。

但好景总是不长久，有一天Go世界突然崩溃了，只留下一句话：Segmentation Fault。你一下怂了，怎么段错误了？

经过一个通宵排查，你发现你管辖的黑户对象竟然偷偷的跟Go世界的其它合法居民搞在一起，当Go世界以为某个居民已经消亡时，用GC回收了它的住所，但是你的地下世界却认为它还活着，还继续访问它。

于是你废了一番功夫斩断了所有关联，世界暂时宁静了下来。

但是你已经很累了，这时候你想起一句话：

##为无为，则无不治