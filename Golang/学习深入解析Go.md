#学习<<深入解析Go>>
##介绍
###如何研究Go内部实现
Go是一种编译型语言，它结合了解释型语言的游刃有余，动态类型语言的开发效率，以及静态类型的安全性。它也打算成为现代的，支持网络与多核计算的语言。要满足这些目标，需要解决一些语言上的问题：一个富有表达能力但轻量级的类型系统，并发与垃圾回收机制，严格的依赖规范等等。这些无法通过库或工具解决好，因此Go也就应运而生了。

各种编程语言理解：<br>
编译型语言在程序执行之前，有一个单独的编译过程，将程序翻译成机器语言，以后执行这个程序的时候，就不用再进行翻译了。

解释型语言，是在运行的时候将程序翻译成机器语言，所以运行速度相对于编译型语言要慢。

C/C++ 等都是编译型语言，而Java，C#等都是解释型语言。

虽然Java程序在运行之前也有一个编译过程，但是并不是将程序编译成机器语言，而是将它编译成字节码（可以理解为一个中间语言）。在运行的时候，由JVM将字节码再翻译成机器语言。

注：脚本语言一般都有相应的脚本引擎来解释执行。 他们一般需要解释器才能运行。JAVASCRIPT,ASP,PHP,PERL,Nuva都是脚本语言。C/C++编译、链接后，可形成独立执行的exe文件。
##最基本数据结构
slice不是一个指针，它在栈中是占三个机器字节的。

字符串在Go语言内存模型中用一个2字长的数据结构表示。它包含一个指向字符串存储数据的指针和一个长度数据。因为string类型是不可变的，对于多字符串共享同一个存储数据是安全的。切分操作str[i:j]会得到一个新的2字长结构，一个可能不同的但仍指向同一个字节序列(即上文说的存储数据)的指针和长度数据。这意味着字符串切分可以在不涉及内存分配或复制操作。这使得字符串切分的效率等同于传递下标。

一个slice是一个数组某个部分的引用。在内存中，它是一个包含3个域的结构体：指向slice中第一个元素的指针，slice的长度，以及slice的容量。长度是下标操作的上界，如x[i]中i必须小于长度。容量是分割操作的上界，如x[i:j]中j不能大于容量。

数组的slice并不会实际复制一份数据，它只是创建一个新的数据结构，包含了另外的一个指针，一个长度和一个容量数据。 如同分割一个字符串，分割数组也不涉及复制操作：它只是新建了一个结构来放置一个不同的指针，长度和容量。在例子中，对[]int{2,3,5,7,11}求值操作会创建一个包含五个值的数组，并设置x的属性来描述这个数组。分割表达式x[1:3]并不分配更多的数据：它只是写了一个新的slice结构的属性来引用相同的存储数据。在例子中，长度为2--只有y[0]和y[1]是有效的索引，但是容量为4--y[0:4]是一个有效的分割表达式。

####slice的扩容

其实slice在Go的运行时库中就是一个C语言动态数组的实现。

在对slice进行append等操作时，可能会造成slice的自动扩容。其扩容时的大小增长规则是：

- 如果新的大小是当前大小2倍以上，则大小增长为新大小
- 否则循环以下操作：如果当前大小小于1024，按每次2倍增长，否则每次按当前大小1/4增长。直到增长的大小超过或等于新大小。

Go有两个数据结构创建函数：new和make。基本的区别是new(T)返回一个*T，返回的这个指针可以被隐式地消除引用（图中的黑色箭头）。而make(T, args)返回一个普通的T。通常情况下，T内部有一些隐式的指针（图中的灰色箭头）。一句话，new返回一个指向已清零内存的指针，而make返回一个复杂的结构。

####slice与unsafe.Pointer相互转换
有时候可能需要使用一些比较tricky的技巧，比如利用make弄一块内存了自己管理，或者cgo之类的方式得到的内存，转换为Go类型使用。

从slice中得到一块内存地址是很容易的：
<pre>
s := make([]byte, 200)
ptr := unsafe.Pointer(&s[0])
</pre>
1. 从一个内存指针构造出Go语言的slice结构相对麻烦一些，比如其中一种方式：
<pre>
var ptr unsafe.Pointer
s := ((*[1<<10]byte)(ptr))[:200]
</pre>
先将ptr强制类型转换为另一种指针，一个指向[1<<10]byte数组的指针，这里数组大小其实是假的。然后用slice操作取出这个数组的前200个，于是s就是一个200个元素的slice。

2. 或者这种方式：
<pre>
var ptr unsafe.Pointer
var s1 = struct {
    addr uintptr
    len int
    cap int
}{ptr, length, length}
s := *(*[]byte)(unsafe.Pointer(&s1))
</pre>
把slice的底层结构写出来，将addr，len，cap等字段写进去，将这个结构体赋给s。相比上一种写法，这种更好的地方在于cap更加自然，虽然上面写法中实际上1<<10就是cap。

3. 又或者使用reflect.SliceHeader的方式来构造slice，比较推荐这种做法：
<pre>
var o []byte
sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
sliceHeader.Cap = length
sliceHeader.Len = length
sliceHeader.Data = uintptr(ptr)
</pre>
####map的实现
Go中的map在底层是用哈希表实现的.
####数据结构
哈希表的数据结构中一些关键的域如下所示：
<pre>
struct Hmap
{
    uint8   B;    // 可以容纳2^B个项
    uint16  bucketsize;   // 每个桶的大小

    byte    *buckets;     // 2^B个Buckets的数组
    byte    *oldbuckets;  // 前一个buckets，只有当正在扩容时才不为空
};
</pre>
上面给出的结构体只是Hmap的部分的域。需要注意到的是，这里直接使用的是Bucket的数组，而不是Bucket*指针的数组。这意味着，第一个Bucket和后面溢出链的Bucket分配有些不同。第一个Bucket是用的一段连续的内存空间，而后面溢出链的Bucket的空间是使用mallocgc分配的。

这个hash结构使用的是一个可扩展哈希的算法，由hash值mod当前hash表大小决定某一个值属于哪个桶，而hash表大小是2的指数，即上面结构体中的2^B。每次扩容，会增大到上次大小的两倍。结构体中有一个buckets和一个oldbuckets是用来实现增量扩容的。正常情况下直接使用buckets，而oldbuckets为空。如果当前哈希表正在扩容中，则oldbuckets不为空，并且buckets大小是oldbuckets大小的两倍。

具体的Bucket结构如下所示：
<pre>
struct Bucket
{
    uint8  tophash[BUCKETSIZE]; // hash值的高8位....低位从bucket的array定位到bucket
    Bucket *overflow;           // 溢出桶链表，如果有
    byte   data[1];             // BUCKETSIZE keys followed by BUCKETSIZE values
};
</pre>
其中BUCKETSIZE是用宏定义的8，每个bucket中存放最多8个key/value对, 如果多于8个，那么会申请一个新的bucket，并将它与之前的bucket链起来。

按key的类型不同会采用相应不同的hash算法得到key的hash值。将hash值的低位当作Hmap结构体中buckets数组的index，找到key所在的bucket。将hash的高8位存储在了bucket的tophash中。注意，这里高8位不是用来当作在key/value在bucket内部的offset的，而是作为一个主键，在查找时是对tophash数组的每一项进行顺序匹配的。先比较hash值高位与bucket的tophash[i]是否相等，如果相等则再比较bucket的第i个的key与所给的key是否相等。如果相等，则返回其对应的value，反之，在overflow buckets中按照上述方法继续寻找。

####增量扩容
<b>大家都知道哈希表表就是以空间换时间，访问速度是直接跟填充因子相关的</b>，所以当哈希表太满之后就需要进行扩容。

如果扩容前的哈希表大小为2^B，扩容之后的大小为2^(B+1)，每次扩容都变为原来大小的两倍，哈希表大小始终为2的指数倍，则有(hash mod 2^B)等价于(hash & (2^B-1))。这样可以简化运算，避免了取余操作。

为什么会增量扩容呢？主要是缩短map容器的响应时间。假如我们直接将map用作某个响应实时性要求非常高的web应用存储，如果不采用增量扩容，当map里面存储的元素很多之后，扩容时系统就会卡往，导致较长一段时间内无法响应请求。不过增量扩容本质上还是将总的扩容时间分摊到了每一次哈希操作上面。

扩容会建立一个大小是原来2倍的新的表，将旧的bucket搬到新的表中之后，并不会将旧的bucket从oldbucket中删除，而是加上一个已删除的标记。

正是由于这个工作是逐渐完成的，这样就会导致一部分数据在old table中，一部分在new table中， 所以对于hash table的insert, remove, lookup操作的处理逻辑产生影响。只有当所有的bucket都从旧表移到新表之后，才会将oldbucket释放掉。
####查找过程
1. 根据key计算出hash值。
2. 如果存在old table, 首先在old table中查找，如果找到的bucket已经evacuated，转到步骤3。 反之，返回其对应的value。
3. 在new table中查找对应的value。

这里一个细节需要注意一下。不认真看可能会以为低位用于定位bucket在数组的index，那么高位就是用于key/valule在bucket内部的offset。事实上高8位不是用作offset的，而是用于加快key的比较的.
<pre>
do { //对每个桶b
    //依次比较桶内的每一项存放的tophash与所求的hash值高位是否相等
    for(i = 0, k = b->data, v = k + h->keysize * BUCKETSIZE; i < BUCKETSIZE; i++, k += h->keysize, v += h->valuesize) {
        if(b->tophash[i] == top) { 
            k2 = IK(h, k);
            t->key->alg->equal(&eq, t->key->size, key, k2);
            if(eq) { //相等的情况下再去做key比较...
                *keyp = k2;
                return IV(h, v);
            }
        }
    }
    b = b->overflow; //b设置为它的下一下溢出链
} while(b != nil);
</pre>

####插入过程分析

1. 根据key算出hash值，进而得出对应的bucket。
2. 如果bucket在old table中，将其重新散列到new table中。
3. 在bucket中，查找空闲的位置，如果已经存在需要插入的key，更新其对应的value。
4. 根据table中元素的个数，判断是否grow table。
5. 如果对应的bucket已经full，重新申请新的bucket作为overbucket。
6. 将key/value pair插入到bucket中。
7. 
这里也有几个细节需要注意一下。

在扩容过程中，oldbucket是被冻结的，查找时会在oldbucket中查找，但不会在oldbucket中插入数据。如果在oldbucket是找到了相应的key，做法是将它迁移到新bucket后加入evalucated标记。并且来会额外的迁移另一个pair。

然后就是只要在某个bucket中找到第一个空位，就会将key/value插入到这个位置。也就是位置位于bucket前面的会覆盖后面的(类似于存储系统设计中做删除时的常用的技巧之一，直接用新数据追加方式写，新版本数据覆盖老版本数据)。找到了相同的key或者找到第一个空位就可以结束遍历了。不过这也意味着做删除时必须完全的遍历bucket所有溢出链，将所有的相同key数据都删除。所以目前map的设计是为插入而优化的，删除效率会比插入低一些。
####map设计中的性能优化
HMap中是Bucket的数组，而不是Bucket指针的数组。好的方面是可以一次分配较大内存，减少了分配次数，避免多次调用mallocgc。但相应的缺点，其一是可扩展哈希的算法并没有发生作用，扩容时会造成对整个数组的值拷贝(如果实现上用Bucket指针的数组就是指针拷贝了，代价小很多)。其二是首个bucket与后面产生了不一致性。这个会使删除逻辑变得复杂一点。比如删除后面的溢出链可以直接删除，而对于首个bucket，要等到evalucated完毕后，整个oldbucket删除时进行。

没有重用设freelist重用删除的结点。作者把这个加了一个TODO的注释，不过想了一下觉得这个做的意义不大。因为一方面，bucket大小并不一致，重用比较麻烦。另一方面，下层存储已经做过内存池的实现了，所以这里不做重用也会在内存分配那一层被重用的，

bucket直接key/value和间接key/value优化。这个优化做得蛮好的。注意看代码会发现，如果key或value小于128字节，则它们的值是直接使用的bucket作为存储的。否则bucket中存储的是指向实际key/value数据的指针，

bucket存8个key/value对。查找时进行顺序比较。第一次发现高位居然不是用作offset，而是用于加快比较的。定位到bucket之后，居然是一个顺序比较的查找过程。后面仔细想了想，觉得还行。由于bucket只有8个，顺序比较下来也不算过分。仍然是O(1)只不过前面系数大一点点罢了。相当于hash到一个小范围之后，在这个小范围内顺序查找。

插入删除的优化。前面已经提过了，插入只要找到相同的key或者第一个空位，bucket中如果存在一个以上的相同key，前面覆盖后面的(只是如果，实际上不会发生)。而删除就需要遍历完所有bucket溢出链了。这样map的设计就是为插入优化的。考虑到一般的应用场景，这个应该算是很合理的。
###nil的语义
代表的是空值的语义。

按照Go语言规范，任何类型在未初始化时都对应一个零值：布尔类型是false，整型是0，字符串是""，而指针，函数，interface，slice，channel和map的零值都是nil。
####interface
一个interface在没有进行初始化时，对应的值是nil。也就是说var v interface{}，此时v就是一个nil。在底层存储上，它是一个空指针。与之不同的情况是，interface值为空。比如：
<pre>
var v *T
var i interface{}
i = v
</pre>
此时i是一个interface，它的值是nil，但它自身不为nil。

Go中的error其实就是一个实现了Error方法的接口：
<pre>
type error interface {
	Error() string
}
</pre>
据此，我们可以自定义一个error：
<pre>
type Error struct {
    errCode uint8
}
func (e *Error) Error() string {
        switch e.errCode {
        case 1:
                return "file not found"
        case 2:
                return "time out"
        case 3:
                return "permission denied"
        default:
                return "unknown error"
         }
}
</pre>
如果我们这样使用它：
<pre>
func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
var e *Error
checkError(e)
</pre>
####string和slice
string的空值是""，它是不能跟nil比较的。即使是空的string，它的大小也是两个机器字长的(ptr,len)。slice也类似，slice的空值是nil,它的空值并不是一个空指针，而是结构体中的指针域为空，空的slice的大小也是三个机器字长的(ptr,len,cap)。
####channel和map

channel跟string或slice有些不同，它在栈上只是一个指针，实际的数据都是由指针所指向的堆上面。

跟channel相关的操作有：初始化/读/写/关闭。channel未初始化值就是nil，未初始化的channel是不能使用的。下面是一些操作规则：

- 读或者写一个nil的channel的操作会永远阻塞。
- 读一个关闭的channel会立刻返回一个channel元素类型的零值。
- 写一个关闭的channel会导致panic。

map也是指针，实际数据在堆中，未初始化的值是nil。
###函数调用协议
理解Go的函数调用协议对于研究其内部实现非常重要。这里将会介绍Go进行函数调用时的内存布局，参数传递和返回值的约定。正如C和汇编都是同一套约定所以能相互调用一样，Go和C以及汇编也是要满足某些约定才能够相互调用。

参见：https://github.com/dongjun111111/blog/issues/24

###defer使用注意点
先来看看几个例子。例1：
<pre>
func f() (result int) {
    defer func() {
        result++
    }()
    return 0
}
</pre>
例2：
<pre>
func f() (r int) {
     t := 5
     defer func() {
       t = t + 5
     }()
     return t
}
</pre>
例3：
<pre>
func f() (r int) {
    defer func(r int) {
          r = r + 5
    }(r)
    return 1
}
</pre>
函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。

defer表达式可能会在设置函数返回值之后，在返回到调用函数之前，修改返回值，使最终的函数返回值与你想象的不一致。

###连续栈 - Golang实现多goroutine的原因
Go语言支持goroutine，每个goroutine需要能够运行，所以它们都有自己的栈。假如每个goroutine分配固定栈大小并且不能增长，太小则会导致溢出，太大又会浪费空间，无法存在许多的goroutine。

为了解决这个问题，goroutine可以初始时只给栈分配很小的空间，然后随着使用过程中的需要自动地增长。这就是为什么Go可以开千千万万个goroutine而不会耗尽内存。
####基本原理
每次执行函数调用时Go的runtime都会进行检测，若当前栈的大小不够用，则会触发“中断”，从当前函数进入到Go的运行时库，Go的运行时库会保存此时的函数上下文环境，然后分配一个新的足够大的栈空间，将旧栈的内容拷贝到新栈中，并做一些设置，使得当函数恢复运行时，函数会在新分配的栈中继续执行，仿佛整个过程都没发生过一样，这个函数会觉得自己使用的是一块大小“无限”的栈空间。
####实现过程
在研究Go的实现细节之前让我们先自己思考一下应该如何实现。第一步肯定要有某种机制检测到当前栈大小不够用了，这个应该是把当前的栈寄存器SP跟栈的可用栈空间的边界进行比较。能够检测到栈大小不够用，就相当于捕捉到了“中断”。

捕获完“中断”，第二步要做的，就应该是进入运行时，保存当前goroutine的上下文。别陷入如何保存上下文的细节，先假如我们把函数栈增长时的上下文保存好了，那下一步就是分配新的栈空间了，我们可以将分配空间想象成就是调用一下malloc而已。

接下来怎么办呢？我们要将旧栈中的内容拷贝到新栈中，然后让函数继续在新栈中运行。这里先暂时忽略旧栈内容拷贝到新栈中的一些技术难点，假设在新栈空间中恢复了“中断”时的上下文，从运行时返回到函数。

函数在新的栈中继续运行了，但是还有个问题：函数如何返回。因为函数返回后栈是要缩小的，否则就会内存浪费空间了，所以还需要在函数返回时处理栈缩小的问题。
####具体细节
如何捕获到函数的栈空间不足

Go语言和C不同，不是使用栈指针寄存器和栈基址寄存器确定函数的栈的。在Go的运行时库中，每个goroutine对应一个结构体G，大致相当于进程控制块的概念。这个结构体中存了stackbase和stackguard，用于确定这个goroutine使用的栈空间信息。每个Go函数调用的前几条指令，先比较栈指针寄存器跟g->stackguard，检测是否发生栈溢出。如果栈指针寄存器值超越了stackguard就需要扩展栈空间。

为了加深理解，下面让我们跟踪一下代码，并看看实际生成的汇编吧。首先写一个test.go文件，内容如下：
<pre>
package main
func main() {
    main()
}
</pre>
然后生成汇编文件：
<pre>
go tool 6g -S test.go
</pre>
可以看以输出是：
<pre>
000000 00000 (test.go:3)    TEXT    "".main+0(SB),$0-0
000000 00000 (test.go:3)    MOVQ    (TLS),CX
0x0009 00009 (test.go:3)    CMPQ    SP,(CX)
0x000c 00012 (test.go:3)    JHI    ,21
0x000e 00014 (test.go:3)    CALL    ,runtime.morestack00_noctxt(SB)
0x0013 00019 (test.go:3)    JMP    ,0
0x0015 00021 (test.go:3)    NOP    ,
</pre>
让我们好好看一下这些指令。(TLS)取到的是结构体G的第一个域，也就是g->stackguard地址，将它赋值给CX。然后CX地址的值与SP进行比较，如果SP大于g->stackguard了，则会调用runtime.morestack00函数。这几条指令的作用就是检测栈是否溢出。

不过并不是所有函数在链接时都会插入这种指令。如果你读源代码，可能会发现#pragma textflag 7，或者在汇编函数中看到TEXT reuntime.exit(SB),7,$0，这种函数就是不会检测栈溢出的。这个是编译标记，控制是否生成栈溢出检测指令。

runtime.morestack是用汇编实现的，做的事情大致是将一些信息存在M结构体中，这些信息包括当前栈桢，参数，当前函数调用，函数返回地址（两个返回地址，一个是runtime.morestack的函数地址，一个是f的返回地址）。通过这些信息可以把新栈和旧栈链起来。
<pre>
void runtime.morestack() {
    if(g == g0) {
        panic();
    } else {
        m->morebuf.gobuf_pc = getCallerCallerPC();
        void *SP = getCallerSP();
        m->morebuf.gobuf_sp = SP;
        m->moreargp = SP;
        m->morebuf.gobuf_g = g;
        m->morepc = getCallerPC();

        void *g0 = m->g0;
        g = g0;
        setSP(g0->g_sched.gobuf_sp);
        runtime.newstack();
    }
}
</pre>
需要注意的就是newstack是切换到m->g0的栈中去调用的。m->g0是调度器栈，go的运行时库的调度器使用的都是m->g0。
####旧栈数据复制到新栈
runtime.morestack会调用于runtime.newstack，newstack做的事情很好理解：分配一个足够大的新的空间，将旧的栈中的数据复制到新的栈中，进行适当的修饰，伪装成调用过runtime.lessstack的样子（这样当函数返回时就会调用于runtime.lessstack再次进入runtime中做一些栈收缩的处理）。

这里有一个技术难点：旧栈数据复制到新栈的过程，要考虑指针失效问题。

比如有某个指针，引用了旧栈中的地址，如果仅仅是将旧栈内容搬到新栈中，那么该指针就失效了，因为旧栈已被释放，应该修改这个指针让它指向新栈的对应地址。考虑如下代码：
<pre>
func f1() {
    var a A
    f(&a)
}
func f2(a *A) {
    // modify a
}
</pre>
如果在f2中发生了栈增长，此时分配更大的空间作为新栈，并将旧栈内容拷贝到新栈中，仅仅这样是不够的，因为f2中的a还是指向旧栈中的f1的，所以必须调整。

Go实现了精确的垃圾回收，运行时知道每一块内存对应的对象的类型信息。在复制之后，会进行指针的调整。具体做法是，对当前栈帧之前的每一个栈帧，对其中的每一个指针，检测指针指向的地址，如果指向地址是落在旧栈范围内的，则将它加上一个偏移使它指向新栈的相应地址。这个偏移值等于新栈基地址减旧栈基地址。

runtime.lessstack比较简单，它其实就是切换到m->g0栈之后调用runtime.oldstack函数。这时之前保存的那个Stktop结构体是时候发挥作用了，从上面可以找到旧栈空间的SP和PC等信息，通过runtime.gogo跳转过去，整个过程就完成了。
<pre>
gp = m->curg; //当前g
top = (Stktop*)gp->stackbase; //取得Stktop结构体
label = top->gobuf; //从结构体中取出Gobuf
runtime·gogo(&label, cret); //通过Gobuf恢复上下文
</pre>
####小结
使用分段栈的函数头几个指令检测%esp和stackguard，调用于runtime.morestack
runtime.more函数的主要功能是保存当前的栈的一些信息，然后转换成调度器的栈了调用runtime.newstack
runtime.newstack函数的主要功能是分配空间，装饰此空间，将旧的frame和arg弄到新空间
使用gogocall的方式切换到新分配的栈，gogocall使用的JMP返回到被中断的函数
继续执行遇到RET指令时会返回到runtime.less，less做的事情跟more相反，它要准备好从newstack到old　stack
整个过程有点像一次中断，中断处理时保存当时的现场，弄个新的栈，中断恢复时恢复到新栈中运行。栈的收缩是垃圾回收的过程中实现的．当检测到栈只使用了不到1/4时，栈缩小为原来的1/2.
###Golang闭包
看下面的例子就知道了：
<pre>
package main
//go闭包
import (
	"fmt"
)
func f(i int) func() int{
	return func() int{
		i++
		return i
	}
}
func main(){
	c1 := f(0)
	fmt.Println(c1())
	fmt.Println(c1())
	c2 :=f(5)
	fmt.Println(c2())
	fmt.Println(c2())
}
output==>
1
2
6
7
</pre>

闭包的用途

闭包可以用在许多地方。它的最大用处有两个，一个是前面提到的可以读取函数内部的变量，另一个就是让这些变量的值始终保持在内存中（看看上面的c1两次输出结果）。

使用闭包的注意点

1. 由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包
2. 闭包会在父函数外部，改变父函数内部变量的值。所以，如果你把父函数当作对象（object）使用，把闭包当作它的公用方法（Public Method），把内部变量当作它的私有属性（private value），这时一定要小心，不要随便改变父函数内部变量的值。

闭包小结

- Go语言支持闭包
- Go语言能通过escape analyze识别出变量的作用域，自动将变量在堆上分配。将闭包环境变量在堆上分配是Go实现闭包的基础。
- 返回闭包时并不是单纯返回一个结构体，而是返回了一个结构体，记录下函数返回地址和引用的环境中的变量地址。
###设计与演化
为了理解goroutine的本质，这里将从最基本的线程池讲起，谈谈Go调度设计背后的故事，讲清楚它为什么是这样子。
####线程池
把每个工作线程叫worker的话，每条线程运行一个worker，每个worker做的事情就是不停地从队列中取出任务并执行：
<pre>
while(!empty(queue)) {
    q = get(queue); //从任务队列中取一个(涉及加锁等)
    q->callback(); //执行该任务
}
</pre>
当然，这是最简单的情形，但是一个很明显的问题就是一个进入callback之后，就失去了控制权。因为没有一个调度器层的东西，一个任务可以执行很长很长时间一直占用的worker线程，或者阻塞于io之类的。
也许用Go语言表述会更地道一些。好吧，那么让我们用Go语言来描述。假设我们有一些“任务”，任务是一个可运行的东西，也就是只要满足Run函数，它就是一个任务。所以我们就把这个任务叫作接口G吧。
<pre>
type G interface {
    Run() 
}
</pre>
我们有一个全局的任务队列，里面包含很多可运行的任务。线程池的各个线程从全局的任务队列中取任务时，显然是需要并发保护的，所以有下面这个结构体：
<pre>
type Sched struct {
    allg  []G
    lock    *sync.Mutex
}
</pre>
以及它的变量
<pre>
	var sched Sched
</pre>
每条线程是一个worker，这里我们给worker换个名字，就把它叫M吧。前面已经说过了，worker做的事情就是不停的去任务队列中取一个任务出来执行。于是用Go语言大概可以写成这样子：
<pre>
func M() {
    for {
        sched.lock.Lock()    //互斥地从就绪G队列中取一个g出来运行
        if sched.allg > 0 {
            g := sched.allg[0]
            sched.allg = sched.allg[1:]
            sched.lock.Unlock()
            g.Run()        //运行它
        } else {
            sched.lock.Unlock()
        }
    }
}
</pre>
接下来，将整个系统启动：
<pre>
for i:=0; i<GOMAXPROCS; i++ {
    go M()
}
</pre>
假定我们有一个满足G接口的main，然后它在自己的Run中不断地将新的任务挂到sched.allg中，这个线程池+任务队列的系统模型就会一直运行下去。

可以看到，这里在代码取中故意地用Go语言中的G，M，甚至包括GOMAXPROCS等取名字。其实本质上，Go语言的调度层无非就是这样一个工作模式的：几条物理线程，不停地取goroutine运行。
####系统调用
上面的情形太简单了，就是工作线程不停地取goroutine运行，这个还不能称之为调度。调度之所以为调度，是因为有一些复杂的控制机制，比如哪个goroutine应该被运行，它应该运行多久，什么时候将它换出来。用前面的代码来说明Go的调度会有一些小问题。Run函数会一直执行，在它结束之前不会返回到调用器层面。那么假设上面的任务中Run进入到一个阻塞的系统调用了，那么M也就跟着一起阻塞了，实际工作的线程就少了一个，无法充分利用CPU。

一个简单的解决办法是在进入系统调用之前再制造一个M出来干活，这样就填补了这个进入系统调用的M的空缺，始终保证有GOMAXPROCS个工作线程在干活了。
<pre>
func entersyscall() {
    go M()
}
</pre>
那么出系统调用时怎么办呢？如果让M接着干活，岂不超过了GOMAXPROCS个线程了？所以这个M不能再干活了，要限制干活的M个数为GOMAXPROCS个，多了则让它们闲置(物理线程比CPU多很多就没意义了，让它们相互抢CPU反而会降低利用率)。
<pre>
func exitsyscall() {
    if len(allm) >= GOMAXPROCS {
        sched.lock.Lock()
        sched.allg = append(sched.allg, g)    //把g放回到队列中
        sched.lock.Unlock()
        time.Sleep()    //这个M不再干活
    }
}
</pre>
其实这个也很好理解，就像线程池做负载调节一样，当任务队列很长后，忙不过来了，则再开几条线程出来。而如果任务队列为空了，则可以释放一些线程。
####协程与保存上下文

大家都知道阻塞于系统调用，会白白浪费CPU。而使用异步事件或回调的思维方式又十分反人类。上面的模型既然这么简单明了，为什么不这么用呢？其实上面的东西看上去简单，但实现起来确不那么容易。

将一个正在执行的任务yield出去，再在某个时刻再弄回来继续运行，这就涉及到一个麻烦的问题，即保存和恢复运行时的上下文环境。

在此先引入协程的概念。协程是轻量级的线程，它相对线程的优势就在于协程非常轻量级，进行切换以及保存上下文环境代价非常的小。协程的具体的实现方式有多种，上面就是其中一种基于线程池的实现方式。每个协程是一个任务，可以保存和恢复任务运行时的上下文环境。

协程一类的东西一般会提供类似yield的函数。协程运行到一定时候就主动调用yield放弃自己的执行，把自己再次放回到任务队列中等待下一次调用时机等等。

其实Go语言中的goroutine就是协程。每个结构体G中有一个sched域就是用于保存自己上下文的。这样，这种goroutine就可以被换出去，再换进来。这种上下文保存在用户态完成，不必陷入到内核，非常的轻量，速度很快。保存的信息很少，只有当前的PC,SP等少量信息。只是由于要优化，所以代码看上去更复杂一些，比如要重用内存空间所以会有gfree和mhead之类的东西。
###抢占式调度
Goroutine本来是设计为协程形式，但是随着调度器的实现越来越成熟，Go在1.2版中开始引入比较初级的抢占式调度。
####从一个bug说起
Go在设计之初并没考虑将goroutine设计成抢占式的。用户负责让各个goroutine交互合作完成任务。一个goroutine只有在涉及到加锁，读写通道或者主动让出CPU等操作时才会触发切换。

垃圾回收器是需要stop the world的。如果垃圾回收器想要运行了，那么它必须先通知其它的goroutine合作停下来，这会造成较长时间的等待时间。考虑一种很极端的情况，所有的goroutine都停下来了，只有其中一个没有停，那么垃圾回收就会一直等待着没有停的那一个。

抢占式调度可以解决这种问题，在抢占式情况下，如果一个goroutine运行时间过长，它就会被剥夺运行权。
####总体思路

引入抢占式调度，会对最初的设计产生比较大的影响，Go还只是引入了一些很初级的抢占，并没有像操作系统调度那么复杂，没有对goroutine分时间片，设置优先级等。

只有长时间阻塞于系统调用，或者运行了较长时间才会被抢占。runtime会在后台有一个检测线程，它会检测这些情况，并通知goroutine执行调度。

目前并没有直接在后台的检测线程中做处理调度器相关逻辑，只是相当于给goroutine加了一个“标记”，然后在它进入函数时才会触发调度。这么做应该是出于对现有代码的修改最小的考虑。
####sysmon

前面讲Go程序的初始化过程中有提到过，runtime开了一条后台线程，运行一个sysmon函数。这个函数会周期性地做epoll操作，同时它还会检测每个P是否运行了较长时间。

如果检测到某个P状态处于Psyscall超过了一个sysmon的时间周期(20us)，并且还有其它可运行的任务，则切换P。

如果检测到某个P的状态为Prunning，并且它已经运行了超过10ms，则会将P的当前的G的stackguard设置为StackPreempt。这个操作其实是相当于加上一个标记，通知这个G在合适时机进行调度。

目前这里只是尽最大努力送达，但并不保证收到消息的goroutine一定会执行调度让出运行权。
####morestack的修改

前面说的，将stackguard设置为StackPreempt实际上是一个比较trick的代码。我们知道Go会在每个函数入口处比较当前的栈寄存器值和stackguard值来决定是否触发morestack函数。

将stackguard设置为StackPreempt作用是进入函数时必定触发morestack，然后在morestack中再引发调度。

看一下StackPreempt的定义，它是大于任何实际的栈寄存器的值的：
<pre>
// 0xfffffade in hex.
#define StackPreempt ((uint64)-1314)
</pre>
然后在morestack中加了一小段代码，如果发现stackguard为StackPreempt，则相当于调用runtime.Gosched。

所以，到目前为止Go的抢占式调度还是很初级的，比如一个goroutine运行了很久，但是它并没有调用另一个函数，则它不会被抢占。当然，一个运行很久却不调用函数的代码并不是多数情况。
### 内存管理
内存管理是非常重要的一个话题。关于编程语言是否应该支持垃圾回收就有个搞笑的争论，一派人认为，内存管理太重要了，而手动管理麻烦且容易出错，所以我们应该交给机器去管理。另一派人则认为，内存管理太重要了！所以如果交给机器管理我不能放心。争论归争论，但不管哪一派，大家对内存管理重要性的认同都是勿庸质疑的。

Go是一门带垃圾回收的语言，Go语言中有指针，却没有C中那么灵活的指针操作。大多数情况下是不需要用户自己去管理内存的，但是理解Go语言是如何做内存管理对于写出优秀的程序是大有帮助的。

接下来将从两个方面来看Go中的内存管理机制，一个方面是内存池，另一个方面是垃圾回收。
###垃圾回收
Go语言中使用的垃圾回收使用的是标记清扫算法。进行垃圾回收时会stoptheworld。不过，在当前1.3版本中，实现了精确的垃圾回收和并行的垃圾回收，大大地提高了垃圾回收的速度，进行垃圾回收时系统并不会长时间卡住。
####标记清扫算法
标记清扫算法是一个很基础的垃圾回收算法，该算法中有一个标记初始的root区域，以及一个受控堆区。root区域主要是程序运行到当前时刻的栈和全局数据区域。在受控堆区中，很多数据是程序以后不需要用到的，这类数据就可以被当作垃圾回收了。判断一个对象是否为垃圾，就是看从root区域的对象是否有直接或间接的引用到这个对象。如果没有任何对象引用到它，则说明它没有被使用，因此可以安全地当作垃圾回收掉。

标记清扫算法分为两阶段：标记阶段和清扫阶段。标记阶段，从root区域出发，扫描所有root区域的对象直接或间接引用到的对象，将这些对上全部加上标记。在回收阶段，扫描整个堆区，对所有无标记的对象进行回收。
####位图标记和内存布局
既然垃圾回收算法要求给对象加上垃圾回收的标记，显然是需要有标记位的。一般的做法会将对象结构体中加上一个标记域，一些优化的做法会利用对象指针的低位进行标记，这都只是些奇技淫巧罢了。Go没有这么做，它的对象和C的结构体对象完全一致，使用的是非侵入式的标记位，我们看看它是怎么实现的。

堆区域对应了一个标记位图区域，堆中每个字(不是byte，而是word)都会在标记位区域中有对应的标记位。每个机器字(32位或64位)会对应4位的标记位。因此，64位系统中相当于每个标记位图的字节对应16个堆中的字节。

虽然是一个堆字节对应4位标记位，但标记位图区域的内存布局并不是按4位一组，而是16个堆字节为一组，将它们的标记位信息打包存储的。每组64位的标记位图从上到下依次包括：
<pre>
16位的 特殊位 标记位
16位的 垃圾回收 标记位
16位的 无指针/块边界 的标记位
16位的 已分配 标记位
</pre>
这样设计使得对一个类型的相应的位进行遍历很容易。

前面提到堆区域和堆地址的标记位图区域是分开存储的，其实它们是以mheap.arena_start地址为边界，向上是实际使用的堆地址空间，向下则是标记位图区域。以64位系统为例，计算堆中某个地址的标记位的公式如下：
<pre>
偏移 = 地址 - mheap.arena_start
标记位地址 = mheap.arena_start - 偏移/16 - 1
移位 = 偏移 % 16
标记位 = *标记位地址 >> 移位
</pre>
然后就可以通过 (标记位 & 垃圾回收标记位),(标记位 & 分配位),等来测试相应的位。其中已分配的标记为1<<0,无指针/块边界是1<<16,垃圾回收的标记位为1<<32,特殊位1<<48
####精确的垃圾回收
像C这种不支持垃圾回收的语言，其实还是有些垃圾回收的库可以使用的。这类库一般也是用的标记清扫算法实现的，但是它们都是保守的垃圾回收。为什么叫“保守”的垃圾回收呢？之所以叫“保守”是因为它们没办法获取对象类型信息，因此只能保守地假设地址区间中每个字都是指针。

无法获取对象的类型信息会造成什么问题呢？这里举两个例子来说明。先看第一个例子，假设某个结构体中是不包含指针成员的，那么对该结构体成员进行垃圾回收时，其实是不必要递归地标记结构体的成员的。但是由于没有类型信息，我们并不知道这个结构体成员不包含指针，因此我们只能对结构体的每个字节递归地标记下去，这显然会浪费很多时间。这个例子说明精确的垃圾回收可以减少不必要的扫描，提高标记过程的速度。

再看另一个例子，假设堆中有一个long的变量，它的值是8860225560。但是我们不知道它的类型是long，所以在进行垃圾回收时会把个当作指针处理，这个指针引用到了0x2101c5018位置。假设0x2101c5018碰巧有某个对象，那么这个对象就无法被释放了，即使实际上已经没任何地方使用它。这个例子说明，保守的垃圾回收某些情况下会出现垃圾无法被回收。虽然不会造成大的问题，但总是让人很不爽，都是没有类型信息惹的祸。

现在好了，Go在1.1版本中开始支持精确的垃圾回收。精确的垃圾回收首先需要的就是类型信息，上一节中讲过MSpan结构体，类型信息是存储在MSpan中的。从一个地址计算它所属的MSpan，公式如下：
<pre>
页号 = (地址 - mheap.arena_start) >> 页大小
MSpan = mheap->map[页号]
</pre>
接下来通过MSpan->type可以得到分配块的类型。这是一个MType的结构体：
<pre>
   struct MTypes
    {
        byte    compression;    // one of MTypes_*
        bool    sysalloc;    // whether (void*)data is from runtime·SysAlloc
        uintptr    data;
    };
</pre>
MTypes描述MSpan里分配的块的类型，其中compression域描述数据的布局。它的取值为MTypes_Empty，MTypes_Single，MTypes_Words，MTypes_Bytes四个中的一种。
<pre>
MTypes_Empty:
    所有的块都是free的，或者这个分配块的类型信息不可用。这种情况下data域是无意义的。
MTypes_Single:
    这个MSpan只包含一个块，data域存放类型信息，sysalloc域无意义
MTypes_Words:
    这个MSpan包含多个块(块的种类多于7)。这时data指向一个数组[NumBlocks]uintptr,，数组里每个元素存放相应块的类型信息
MTypes_Bytes:
    这个MSpan中包含最多7种不同类型的块。这时data域指下面这个结构体
    struct {
        type  [8]uintptr       // type[0] is always 0
        index [NumBlocks]byte
    }
    第i个块的类型是data.type[data.index[i]]
</pre>
###垃圾回收
目前Go中垃圾回收的核心函数是scanblock，源代码在文件runtime/mgc0.c中。这个函数非常难读，单个函数写了足足500多行。上面有两个大的循环，外层循环作用是扫描整个内存块区域，将类型信息提取出来，得到其中的gc域。内层的大循环是实现一个状态机，解析执行类型信息中gc域的指令码。

先说说上一节留的疑问吧。MType中的数据其实是类型信息，但它是用uintptr表示，而不是Type结构体的指针，这是一个优化的小技巧。由于内存分配是机器字节对齐的，所以地址就只用到了高位，低位是用不到的。于是低位可以利用起来存储一些额外的信息。这里的uintptr中高位存放的是Type结构体的指针，低位用来存放类型。通过
<pre>
    t = (Type*)(type & ~(uintptr)(PtrSize-1));
</pre>
就可以从uintptr得到Type结构体指针，而通过
<pre>
	type & (PtrSize-1)
</pre>
就可以得到类型。这里的类型有TypeInfo_SingleObject，TypeInfo_Array，TypeInfo_Map，TypeInfo_Chan几种。
####基本的标记过程
从最简单的开始看，基本的标记过程，有一个不带任何优化的标记的实现，对应于函数debug_scanblock。

debug_scanblock函数是递归实现的，单线程的，更简单更慢的scanblock版本。该函数接收的参数分别是一个指针表示要扫描的地址，以及字节数。
<pre>
首先要将传入的地址，按机器字节大小对齐。
然后对待扫描区域的每个地址：
找到它所属的MSpan，将地址转换为MSpan里的对象地址。
根据对象的地址，找到对应的标记位图里的标记位。
判断标记位，如果是未分配则跳过。否则加上特殊位标记(debug_scanblock中用特殊位代码的mark位)完成标记。
判断标记位中标记了无指针标记位，如果没有，则要递归地调用debug_scanblock。
</pre>
这个递归版本的标记算法还是很容易理解的。其中涉及的细节在上节中已经说过了，比如任意给定一个地址，找到它的标记位信息。很明显这里仅仅使用了一个无指针位，并没有精确的垃圾回收。
####并行的垃圾回收
Go在这个版本中不仅实现了精确的垃圾回收，而且实现了并行的垃圾回收。标记算法本质上就是一个树的遍历过程，上面实现的是一个递归版本。

并行的垃圾回收需要做的第一步，就是先将算法做成非递归的。非递归版本的树的遍历需要用到一个队列。树的非递归遍历的伪代码大致是：
<pre>
根结点进队
while(队列不空) {
    出队
    访问
    将子结点进队
}
</pre>
第二步是使上面的代码能够并行地工作，显然这时是需要一个线程安全的队列的。假设有这样一个队列，那么上面代码就能够工作了。但是，如果不加任何优化，这里的队列的并行访问非常地频繁，对这个队列加锁代价会非常高，即使是使用CAS操作也会大大降低效率。

所以，第三步要做的就是优化上面队列的数据结构。事实上，Go中并没有使用这样一个队列，为了优化，它通过三个数据结构共同来完成这个队列的功能，这三个数据结构分别是PtrTarget数组，Workbuf，lfstack。

先说Workbuf吧。听名字就知道，这个结构体的意思是工作缓冲区，里面存放的是一个数组，数组中的每个元素都是一个待处理的结点，也就是一个Obj指针。这个对象本身是已经标记了的，这个对象直接或间接引用到的对象，都是应该被标记的，它们不会被当作垃圾回收掉。Workbuf是比较大的，一般是N个内存页的大小(目前是2页，也就是8K)。

PtrTarget数组也是一个缓冲区，相当于一个intermediate buffer，跟Workbuf有一点点的区别。第一，它比Workbuf小很多，大概只有32或64个元素的数组。第二，Workbuf中的对象全部是已经标记过的，而PtrTarget中的元素可能是标记的，也可能是没标记的。第三，PtrTarget里面的元素是指针而不是对象，指针是指向任意地址的，而对象是对齐到正确地址的。从一个指针变为一个对象要经过一次变换，上一节中有讲过具体细节。

垃圾回收过程中，会有一个从PtrTarget数组冲刷到Workbuf缓冲区的过程。对应于源代码中的flushptrbuf函数，这个函数作用就是对PtrTaget数组中的所有元素，如果该地址是mark了的，则将它移到Workbuf中。标记过程形成了一个环，在环的一边，对Workbuf中的对象，会将它们可能引用的区域全部放到PtrTarget中记录下来。在环的另一边，又会将PtrTarget中确定需要标记的地址刷到Workbuf中。这个过程一轮一轮地进行，推动非递归版本的树的遍历过程，也就是前面伪代码中的出队，访问，子结点进队的过程。

另一个数据结构是lfstack，这个名字的意思是lock free栈。其实它是被用作了一个无锁的链表，链表结点是以Workbuf为单位的。并行垃圾回收中，多条线程会从这个链表中取数据，每次以一个Workbuf为工作单位。同时，标记的过程中也会产生Workbuf结点放到链中。lfstack保证了对这个链的并发访问的安全性。由于现在链表结点是以Workbuf为单位的，所以保证整体的性能，lfstack的底层代码是用CAS操作实现的。

经过第三步中数据结构上的拆解，整个并行垃圾回收的架构已经呼之欲出了，这就是标记扫描的核心函数scanblock。这个函数是在多线程下并行安全的。

那么，最后一步，多线程并行。整个的gc是以runtime.gc函数为入口的，它实际调用的是gc。进入gc函数后会先stoptheworld，接着添加标记的root区域。然后会设置markroot和sweepspan的并行任务。运行mark的任务，扫描块，运行sweep的任务，最后starttheworld并切换出去。

有一个ParFor的数据结构。在gc函数中调用了
<pre>
 runtime·parforsetup(work.markfor, work.nproc, work.nroot, nil, false, markroot);
    runtime·parforsetup(work.sweepfor, work.nproc, runtime·mheap->nspan, nil, true, sweepspan);
</pre>
是设置好回调函数让线程去执行markroot和sweepspan函数。垃圾回收时会stoptheworld，其它goroutine会对发起stoptheworld做出响应，调用runtime.gchelper，这个函数会调用scanblock帮助标记过程。也会并行地做markroot和sweepspan的过程。
<pre>
  void
    runtime·gchelper(void)
    {
        gchelperstart();

        // parallel mark for over gc roots
        runtime·parfordo(work.markfor);

        // help other threads scan secondary blocks
        scanblock(nil, nil, 0, true);

        if(DebugMark) {
            // wait while the main thread executes mark(debug_scanblock)
            while(runtime·atomicload(&work.debugmarkdone) == 0)
                runtime·usleep(10);
        }

        runtime·parfordo(work.sweepfor);
        bufferList[m->helpgc].busy = 0;
        if(runtime·xadd(&work.ndone, +1) == work.nproc-1)
            runtime·notewakeup(&work.alldone);
    }
</pre>
其中并行时也有实现工作流窃取的概念，多个worker同时去工作缓存中取数据出来处理，如果自己的任务做完了，就会从其它的任务中“偷”一些过来执行。
####垃圾回收的时机
垃圾回收的触发是由一个gcpercent的变量控制的，当新分配的内存占已在使用中的内存的比例超过gcprecent时就会触发。比如，gcpercent=100，当前使用了4M的内存，那么当内存分配到达8M时就会再次gc。如果回收完毕后，内存的使用量为5M，那么下次回收的时机则是内存分配达到10M的时候。也就是说，并不是内存分配越多，垃圾回收频率越高，这个算法使得垃圾回收的频率比较稳定，适合应用的场景。

gcpercent的值是通过环境变量GOGC获取的，如果不设置这个环境变量，默认值是100。如果将它设置成off，则是关闭垃圾回收。
###高级数据结构实现--channel数据结构
Go语言channel是first-class的，意味着它可以被存储到变量中，可以作为参数传递给函数，也可以作为函数的返回值返回。作为Go语言的核心特征之一，虽然channel看上去很高端，但是其实channel仅仅就是一个数据结构而已，结构体定义如下：
<pre>
struct    Hchan
{
    uintgo    qcount;            // 队列q中的总数据数量
    uintgo    dataqsiz;        // 环形队列q的数据大小
    uint16    elemsize;
    bool    closed;
    uint8    elemalign;
    Alg*    elemalg;        // interface for element type
    uintgo    sendx;            // 发送index
    uintgo    recvx;            // 接收index
    WaitQ    recvq;            // 因recv而阻塞的等待队列
    WaitQ    sendq;            // 因send而阻塞的等待队列
    Lock;
};
</pre>
让我们来看一个Hchan这个结构体。其中一个核心的部分是存放channel数据的环形队列，由qcount和elemsize分别指定了队列的容量和当前使用量。dataqsize是队列的大小。elemalg是元素操作的一个Alg结构体，记录下元素的操作，如copy函数，equal函数，hash函数等。

可能会有人疑惑，结构体中只看到了队列大小相关的域，并没有看到存放数据的域啊？如果是带缓冲区的chan，则缓冲区数据实际上是紧接着Hchan结构体中分配的。
<pre>
c = (Hchan*)runtime.mal(n + hint*elem->size);
</pre>
另一个重要部分就是recvq和sendq两个链表，一个是因读这个通道而导致阻塞的goroutine，另一个是因为写这个通道而阻塞的goroutine。如果一个goroutine阻塞于channel了，那么它就被挂在recvq或sendq中。WaitQ是链表的定义，包含一个头结点和一个尾结点：
<pre>
struct    WaitQ
{
    SudoG*    first;
    SudoG*    last;
};
</pre>
队列中的每个成员是一个SudoG结构体变量。
<pre>
struct    SudoG
{
    G*    g;        // g and selgen constitute
    uint32    selgen;        // a weak pointer to g
    SudoG*    link;
    int64    releasetime;
    byte*    elem;        // data element
};
</pre>
该结构中主要的就是一个g和一个elem。elem用于存储goroutine的数据。读通道时，数据会从Hchan的队列中拷贝到SudoG的elem域。写通道时，数据则是由SudoG的elem域拷贝到Hchan的队列中。
####读写channel操作
先看写channel的操作，基本的写channel操作，在底层运行时库中对应的是一个runtime.chansend函数。
<pre>
c <- v
</pre>
在运行时库中会执行：
<pre>
void runtime·chansend(ChanType *t, Hchan *c, byte *ep, bool *pres, void *pc)
</pre>
其中c就是channel，ep是取变量v的地址。这里的传值约定是调用者负责分配好ep的空间，仅需要简单的取变量地址就够了。pres参数是在select中的通道操作使用的。

这个函数首先会区分是同步还是异步。同步是指chan是不带缓冲区的，因此可能写阻塞，而异步是指chan带缓冲区，只有缓冲区满才阻塞。

在同步的情况下，由于channel本身是不带数据缓存的，这时首先会查看Hchan结构体中的recvq链表时否为空，即是否有因为读该管道而阻塞的goroutine。如果有则可以正常写channel，否则操作会阻塞。

recvq不为空的情况下，将一个SudoG结构体出队列，将传给通道的数据(函数参数ep)拷贝到SudoG结构体中的elem域，并将SudoG中的g放到就绪队列中，状态置为ready，然后函数返回。

如果recvq为空，否则要将当前goroutine阻塞。此时将一个SudoG结构体，挂到通道的sendq链表中，这个SudoG中的elem域是参数eq，SudoG中的g是当前的goroutine。当前goroutine会被设置为waiting状态并挂到等待队列中。

在异步的情况，如果缓冲区满了，也是要将当前goroutine和数据一起作为SudoG结构体挂在sendq队列中，表示因写channel而阻塞。否则也是先看有没有recvq链表是否为空，有就唤醒。

跟同步不同的是在channel缓冲区不满的情况，这里不会阻塞写者，而是将数据放到channel的缓冲区中，调用者返回。

读channel的操作也是类似的，对应的函数是runtime.chansend。一个是收一个是发，基本的过程都是差不多的。

需要注意的是几种特殊情况下的通道操作--空通道和关闭的通道。

空通道是指将一个channel赋值为nil，或者定义后不调用make进行初始化。按照Go语言的语言规范，读写空通道是永远阻塞的。其实在函数runtime.chansend和runtime.chanrecv开头就有判断这类情况，如果发现参数c是空的，则直接将当前的goroutine放到等待队列，状态设置为waiting。

读一个关闭的通道，永远不会阻塞，会返回一个通道数据类型的零值。这个实现也很简单，将零值复制到调用函数的参数ep中。写一个关闭的通道，则会panic。关闭一个空通
道，也会导致panic。
###select的实现
select-case中的chan操作编译成了if-else。比如：
<pre>
select {
case v = <-c:
        ...foo
default:
        ...bar
}
</pre>
会被编译为:
<pre>
if selectnbrecv(&v, c) {
        ...foo
} else {
        ...bar
}
</pre>
类似地
<pre>
select {
case v, ok = <-c:
    ... foo
default:
    ... bar
}
</pre>
会被编译为:
<pre>
if c != nil && selectnbrecv2(&v, &ok, c) {
    ... foo
} else {
    ... bar
}
</pre>
接下来就是看一下selectnbrecv相关的函数了。其实没有任何特殊的魔法，这些函数只是简单地调用runtime.chanrecv函数，只不过设置了一个参数，告诉当runtime.chanrecv函数，当不能完成操作时不要阻塞，而是返回失败。也就是说，所有的select操作其实都仅仅是被换成了if-else判断，底层调用的不阻塞的通道操作函数。

在Go的语言规范中，select中的case的执行顺序是随机的，而不像switch中的case那样一条一条的顺序执行。那么，如何实现随机呢？

select和case关键字使用了下面的结构体：
<pre>
struct    Scase
{
    SudoG    sg;            // must be first member (cast to Scase)
    Hchan*    chan;        // chan
    byte*    pc;            // return pc
    uint16    kind;
    uint16    so;            // vararg of selected bool
    bool*    receivedp;    // pointer to received bool (recv2)
};

struct    Select
{
    uint16    tcase;            // 总的scase[]数量
    uint16    ncase;            // 当前填充了的scase[]数量
    uint16*    pollorder;        // case的poll次序
    Hchan**    lockorder;        // channel的锁住的次序
    Scase    scase[1];        // 每个case会在结构体里有一个Scase，顺序是按出现的次序
};
</pre>
每个select都对应一个Select结构体。在Select数据结构中有个Scase数组，记录下了每一个case，而Scase中包含了Hchan。然后pollorder数组将元素随机排列，这样就可以将Scase乱序了。
###interface
interface是Go语言中最成功的设计之一，空的interface可以被当作“鸭子”类型使用，它使得Go这样的静态语言拥有了一定的动态性，但却又不损失静态语言在类型安全方面拥有的编译时检查的优势。

依赖于接口而不是实现，优先使用组合而不是继承，这是程序抽象的基本原则。但是长久以来以C++为代表的“面向对象”语言曲解了这些原则，让人们走入了误区。为什么要将方法和数据绑死？为什么要有多重继承这么变态的设计？面向对象中最强调的应该是对象间的消息传递，却为什么被演绎成了封装继承和多态。面向对象是否实现程序程序抽象的合理途径，又或者是因为它存在我们就认为它合理了。历史原因，中间出现了太多的错误。不管怎么样，Go的interface给我们打开了一扇新的窗。

那么，Go中的interface在底层是如何实现的呢？

interface实际上就是一个结构体，包含两个成员。其中一个成员是指向具体数据的指针，另一个成员中包含了类型信息。空接口和带方法的接口略有不同，下面分别是空接口和带方法的接口是使用的数据结构：
<pre>
struct Eface
{
    Type*    type;
    void*    data;
};
struct Iface
{
    Itab*    tab;
    void*    data;
};
</pre>
先看Eface，它是interface{}底层使用的数据结构。数据域中包含了一个void*指针，和一个类型结构体的指针。interface{}扮演的角色跟C语言中的void*是差不多的，Go中的任何对象都可以表示为interface{}。不同之处在于，interface{}中有类型信息，于是可以实现反射。

类型信息的结构体定义如下：
<pre>
struct Type
{
    uintptr size;
    uint32 hash;
    uint8 _unused;
    uint8 align;
    uint8 fieldAlign;
    uint8 kind;
    Alg *alg;
    void *gc;
    String *string;
    UncommonType *x;
    Type *ptrto;
};
</pre>
其实在前面我们已经见过它了。精确的垃圾回收中，就是依赖Type结构体中的gc域的。不同类型数据的类型信息结构体并不完全一致，Type是类型信息结构体中公共的部分，其中size描述类型的大小，hash数据的hash值，align是对齐，fieldAlgin是这个数据嵌入结构体时的对齐，kind是一个枚举值，每种类型对应了一个编号。alg是一个函数指针的数组，存储了hash/equal/print/copy四个函数操作。UncommonType是指向一个函数指针的数组，收集了这个类型的实现的所有方法。

在reflect包中有个KindOf函数，返回一个interface{}的Type，其实该函数就是简单的取Eface中的Type域。

Iface和Eface略有不同，它是带方法的interface底层使用的数据结构。data域同样是指向原始数据的，而Itab的结构如下：
<pre>
struct    Itab
{
    InterfaceType*    inter;
    Type*    type;
    Itab*    link;
    int32    bad;
    int32    unused;
    void    (*fun[])(void);
};
</pre>
Itab中不仅存储了Type信息，而且还多了一个方法表fun[]。一个Iface中的具体类型中实现的方法会被拷贝到Itab的fun数组中。
具体类型向接口类型赋值

将具体类型数据赋值给interface{}这样的抽象类型，中间会涉及到类型转换操作。从接口类型转换为具体类型(也就是反射)，也涉及到了类型转换。这个转换过程中做了哪些操作呢？先看将具体类型转换为接口类型。如果是转换成空接口，这个过程比较简单，就是返回一个Eface，将Eface中的data指针指向原型数据，type指针会指向数据的Type结构体。

将某个类型数据转换为带方法的接口时，会复杂一些。中间涉及了一道检测，该类型必须要实现了接口中声明的所有方法才可以进行转换。这个检测是在编译过程中做的，我们可以做个测试：
<pre>
type I interface {
    String()
}
var a int = 5
var b I = a
</pre>
编译会报错：
<pre>
cannot use a (type int) as type I in assignment:
    int does not implement I (missing String method)
</pre>
说明具体类型转换为带方法的接口类型是在编译过程中进行检测的。

那么这个检测是如何实现的呢？在runtime下找到了iface.c文件，应该是早期版本是在运行时检测留下的，其中有一个itab函数就是判断某个类型是否实现了某个接口，如果是则返回一个Itab结构体。

类型转换时的检测就是比较具体类型的方法表和接口类型的方法表，看具体类型是实现了接口类型所声明的所有的方法。还记得Type结构体中是有个UncommonType字段的，里面有张方法表，类型所实现的方法都在里面。而在Itab中有个InterfaceType字段，这个字段中也有一张方法表，就是这个接口所要求的方法。这两处方法表都是排序过的，只需要一遍顺序扫描进行比较，应该可以知道Type中否实现了接口中声明的所有方法。最后还会将Type方法表中的函数指针，拷贝到Itab的fun字段中。

这里提到了三个方法表，有点容易把人搞晕，所以要解释一下。

Type的UncommonType中有一个方法表，某个具体类型实现的所有方法都会被收集到这张表中。reflect包中的Method和MethodByName方法都是通过查询这张表实现的。表中的每一项是一个Method，其数据结构如下：
<pre>
struct Method
{
    String *name;
    String *pkgPath;
    Type    *mtyp;
    Type *typ;
    void (*ifn)(void);
    void (*tfn)(void);
};
</pre>
Iface的Itab的InterfaceType中也有一张方法表，这张方法表中是接口所声明的方法。其中每一项是一个IMethod，数据结构如下：
<pre>
struct IMethod
{
    String *name;
    String *pkgPath;
    Type *type;
};
</pre>
跟上面的Method结构体对比可以发现，这里是只有声明没有实现的。

Iface中的Itab的func域也是一张方法表，这张表中的每一项就是一个函数指针，也就是只有实现没有声明。

类型转换时的检测就是看Type中的方法表是否包含了InterfaceType的方法表中的所有方法，并把Type方法表中的实现部分拷到Itab的func那张表中。
###reflect

reflect就是给定一个接口类型的数据，得到它的具体类型的类型信息，它的Value等。reflect包中的TypeOf和ValueOf函数分别做这个事情。还有像：
<pre>
    v, ok := i.(T)
</pre>
这样的语法，也是判断一个接口i的具体类型是否为类型T，如果是则将其值返回给v。这跟上面的类型转换一样，也会检测转换是否合法。不过这里的检测是在运行时执行的。在runtime下的iface.c文件中，有一系统的assetX2X函数，比如runtime.assetE2T，runtime.assetI2T等等。这个实现起来比较简单，只需要比较Iface中的Itab的type是否与给定Type为同一个。





####乱入内容 - C语言中的malloc|常用分配内存函数
extern void *malloc(unsigned int num_bytes);

头文件：#include <malloc.h> 或 #include <alloc.h> (注意：alloc.h 与 malloc.h 的内容是完全一致的。)

功能：分配长度为num_bytes字节的内存块

说明：如果分配成功则返回指向被分配内存的指针，否则返回空指针NULL。

当内存不再使用时，应使用free()函数将内存块释放。
<pre>
	void *malloc(int size);
</pre>
说明：malloc 向系统申请分配指定size个字节的内存空间。返回类型是 void* 类型。void* 表示未确定类型的指针。C,C++规定，void* 类型可以强制转换为任何其它类型的指针。

示例：
<pre>
#include<stdio.h>  
#include<malloc.h>  
int main()  
{  
    char *p;  
   
    p=(char *)malloc(100);  
    if(p)  
        printf("Memory Allocated at: %x/n",p);  
    else  
        printf("Not Enough Memory!/n");  
    free(p);  
    return 0;  
}  
</pre>
总结：

malloc()函数其实就在内存中找一片指定大小的空间，然后将这个空间的首地址范围给一个指针变量，这里的指针变量可以是一个单独的指针，也可以是一个数组的首地址，这要看malloc()函数中参数size的具体内容。我们这里malloc分配的内存空间在逻辑上连续的，而在物理上可以连续也可以不连续。对于我们程序员来说，我们关注的是逻辑上的连续，因为操作系统会帮我们安排内存分配，所以我们使用起来就可以当做是连续的。
####I/O操作
简单来说，就是指磁盘的输入和输出。
读/写IO，最为常见说法，读IO，就是发指令，从磁盘读取某段扇区的内容。指令一般是通知磁盘开始扇区位置，然后给出需要从这个初始扇区往后读取的连续扇区个数，同时给出动作是读，还是写。磁盘收到这条指令，就会按照指令的要求，读或者写数据。控制器发出的这种指令＋数据，就是一次IO，读或者写。
####为什么mysql使用B+tree数据结构来实现索引
评价一个数据结构作为索引的优劣最重要的指标就是在查找过程中磁盘I/O操作次数的渐进复杂度。换句话说，索引的结构组织要尽量减少查找过程中磁盘I/O的存取次数。下面先介绍内存和磁盘存取原理，然后再结合这些原理分析B-/+Tree作为索引的效率。

主存存取原理 >>>

目前计算机使用的主存基本都是随机读写存储器（RAM），现代RAM的结构和存取原理比较复杂，这里本文抛却具体差别，抽象出一个十分简单的存取模型来说明RAM的工作原理。

从抽象角度看，主存是一系列的存储单元组成的矩阵，每个存储单元存储固定大小的数据。每个存储单元有唯一的地址，现代主存的编址规则比较复杂，这里将其简化成一个二维地址：通过一个行地址和一个列地址可以唯一定位到一个存储单元。图5展示了一个4 x 4的主存模型。

主存的存取过程如下：

当系统需要读取主存时，则将地址信号放到地址总线上传给主存，主存读到地址信号后，解析信号并定位到指定存储单元，然后将此存储单元数据放到数据总线上，供其它部件读取。

写主存的过程类似，系统将要写入单元地址和数据分别放在地址总线和数据总线上，主存读取两个总线的内容，做相应的写操作。

这里可以看出，主存存取的时间仅与存取次数呈线性关系，因为不存在机械操作，两次存取的数据的“距离”不会对时间有任何影响，例如，先取A0再取A1和先取A0再取D3的时间消耗是一样的。

磁盘存取原理 >>>

上文说过，索引一般以文件形式存储在磁盘上，索引检索需要磁盘I/O操作。与主存不同，磁盘I/O存在机械运动耗费，因此磁盘I/O的时间消耗是巨大的。

一个磁盘由大小相同且同轴的圆形盘片组成，磁盘可以转动（各个磁盘必须同步转动）。在磁盘的一侧有磁头支架，磁头支架固定了一组磁头，每个磁头负责存取一个磁盘的内容。磁头不能转动，但是可以沿磁盘半径方向运动（实际是斜切向运动），每个磁头同一时刻也必须是同轴的，即从正上方向下看，所有磁头任何时候都是重叠的（不过目前已经有多磁头独立技术，可不受此限制）

盘片被划分成一系列同心环，圆心是盘片中心，每个同心环叫做一个磁道，所有半径相同的磁道组成一个柱面。磁道被沿半径线划分成一个个小的段，每个段叫做一个扇区，每个扇区是磁盘的最小存储单元。为了简单起见，我们下面假设磁盘只有一个盘片和一个磁头。

当需要从磁盘读取数据时，系统会将数据逻辑地址传给磁盘，磁盘的控制电路按照寻址逻辑将逻辑地址翻译成物理地址，即确定要读的数据在哪个磁道，哪个扇区。为了读取这个扇区的数据，需要将磁头放到这个扇区上方，为了实现这一点，磁头需要移动对准相应磁道，这个过程叫做寻道，所耗费时间叫做寻道时间，然后磁盘旋转将目标扇区旋转到磁头下，这个过程耗费的时间叫做旋转时间。

局部性原理与磁盘预读 >>>

由于存储介质的特性，磁盘本身存取就比主存慢很多，再加上机械运动耗费，磁盘的存取速度往往是主存的几百分分之一，因此为了提高效率，要尽量减少磁盘I/O。为了达到这个目的，磁盘往往不是严格按需读取，而是每次都会预读，即使只需要一个字节，磁盘也会从这个位置开始，顺序向后读取一定长度的数据放入内存。这样做的理论依据是计算机科学中著名的局部性原理：

当一个数据被用到时，其附近的数据也通常会马上被使用。

程序运行期间所需要的数据通常比较集中。

由于磁盘顺序读取的效率很高（不需要寻道时间，只需很少的旋转时间），因此对于具有局部性的程序来说，预读可以提高I/O效率。

预读的长度一般为页（page）的整倍数。页是计算机管理存储器的逻辑块，硬件及操作系统往往将主存和磁盘存储区分割为连续的大小相等的块，每个存储块称为一页（在许多操作系统中，页得大小通常为4k），主存和磁盘以页为单位交换数据。当程序要读取的数据不在主存中时，会触发一个缺页异常，此时系统会向磁盘发出读盘信号，磁盘会找到数据的起始位置并向后连续读取一页或几页载入内存中，然后异常返回，程序继续运行。

B-/+Tree索引的性能分析 >>>

到这里终于可以分析B-/+Tree索引的性能了。

上文说过一般使用磁盘I/O次数评价索引结构的优劣。先从B-Tree分析，根据B-Tree的定义，可知检索一次最多需要访问h个节点。数据库系统的设计者巧妙利用了磁盘预读原理，将一个节点的大小设为等于一个页，这样每个节点只需要一次I/O就可以完全载入。为了达到这个目的，在实际实现B-Tree还需要使用如下技巧：

每次新建节点时，直接申请一个页的空间，这样就保证一个节点物理上也存储在一个页里，加之计算机存储分配都是按页对齐的，就实现了一个node只需一次I/O。

B-Tree中一次检索最多需要h-1次I/O（根节点常驻内存），渐进复杂度为O(h)=O(logdN)。一般实际应用中，出度d是非常大的数字，通常超过100，因此h非常小（通常不超过3）。

综上所述，用B-Tree作为索引结构效率是非常高的。

而红黑树这种结构，h明显要深的多。由于逻辑上很近的节点（父子）物理上可能很远，无法利用局部性，所以红黑树的I/O渐进复杂度也为O(h)，效率明显比B-Tree差很多。
####nio简介
nio 是 java New IO 的简称，在 jdk1.4 里提供的新 api 。 Sun 官方标榜的特性如下：
–     为所有的原始类型提供 (Buffer) 缓存支持。
–     字符集编码解码解决方案。
–     Channel ：一个新的原始 I/O 抽象。
–     支持锁和内存映射文件的文件访问接口。
–     提供多路 (non-bloking) 非阻塞式的高伸缩性网络 I/O 。
本文将围绕这几个特性进行学习和介绍。
4.   Buffer&Chanel
Channel 和 buffer 是 NIO 是两个最基本的数据类型抽象。
Buffer:
–        是一块连续的内存块。
–        是 NIO 数据读或写的中转地。
Channel:
–        数据的源头或者数据的目的地
–        用于向 buffer 提供数据或者读取 buffer 数据 ,buffer 对象的唯一接口。
–         异步 I/O 支持
####io与nio
1.io是面向流的，也就是读取数据的时候是从流上逐个读取，所以数据不能进行整体以为，没有缓冲区;nio是面向缓冲区的，数据是存储在缓冲区中，读取数据是在缓冲区中进行，所以进行数据的偏移操作更加方便
2，io是阻塞的，当一个线程操作io时如果当前没有数据可读，那么线程阻塞，nio由于是对通道操作io，所以是非阻塞，当一个通道无数据可读，可切换通道处理其他io
3，nio有selecter选择器，就是线程通过选择器可以选择多个通道，而io只能处理一个
<pre>
package sample;  
  
import java.io.FileInputStream;  
import java.io.FileOutputStream;  
import java.nio.ByteBuffer;  
import java.nio.channels.FileChannel;  
  
public class CopyFile {  
    public static void main(String[] args) throws Exception {  
        String infile = "C:\\copy.sql";  
        String outfile = "C:\\copy.txt";  
        // 获取源文件和目标文件的输入输出流  
        FileInputStream fin = new FileInputStream(infile);  
        FileOutputStream fout = new FileOutputStream(outfile);  
        // 获取输入输出通道  
        FileChannel fcin = fin.getChannel();  
        FileChannel fcout = fout.getChannel();  
        // 创建缓冲区  
        ByteBuffer buffer = ByteBuffer.allocate(1024);  
        while (true) {  
            // clear方法重设缓冲区，使它可以接受读入的数据  
            buffer.clear();  
            // 从输入通道中将数据读到缓冲区  
            int r = fcin.read(buffer);  
            // read方法返回读取的字节数，可能为零，如果该通道已到达流的末尾，则返回-1  
            if (r == -1) {  
                break;  
            }  
            // flip方法让缓冲区可以将新读入的数据写入另一个通道  
            buffer.flip();  
            // 从输出通道中将数据写入缓冲区  
            fcout.write(buffer);  
        }  
    }  
} 
</pre>