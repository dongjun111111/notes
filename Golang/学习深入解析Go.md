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





####乱入内容 - C语言中的malloc
extern void *malloc(unsigned int num_bytes);

头文件：#include <malloc.h> 或 #include <alloc.h> (注意：alloc.h 与 malloc.h 的内容是完全一致的。)

功能：分配长度为num_bytes字节的内存块

说明：如果分配成功则返回指向被分配内存的指针，否则返回空指针NULL。

当内存不再使用时，应使用free()函数将内存块释放。


void *malloc(int size);

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