#Golang开发环境本地配置（顶，由果子帮忙实现）
go安装目录在C:\go文件夹下面；go的工作目录在D:\gopath文件夹下。现在gopath目录下有pkg与src文件夹。pkg文件夹下是第三方包（一般是github.com文件夹下）的存放地点，src下是第三方包（一般是github.com文件夹下）与主要编码区（例如test文件夹下）。

对应的环境变量是：

- GOPATH:D：\gopath
- GOROOT:C:\go

Path加上

- PATH:c:\go\bin;d:\gopath\bin(如果在最后分号不要加上)

F7 --> 呼出sublime的golang调试控制台
##安装Go-sublime
http://blog.csdn.net/kenkao/article/details/49488833

###net/http
在net/http包中，动态文件的路由和静态文件的路由是分开的，动态文件使用http.HandleFunc进行设置，静态文件就需要使用到http.FileServer
####如何设置Cookie
<pre>
cookie := http.Cookie{Name: "admin_name", Value: rows[0].Str(res.Map("admin_name")), Path: "/"}
http.SetCookie(w,&cookie)
</pre>
####http.FileServer()
文件系统：将本地文件输出到网页。
<pre>
package main
import (
    "net/http"
)
func main() {
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.ListenAndServe(":8123", nil)
}
</pre>
当然用golang写一个文件服务很简单，比如上面的，但是如果想通过localhost:8123/doc（即自定义文件服务器入口）来进入文件目录，则需要
<pre>
http.Handle("/doc",http.StripPrefix("/doc",http.FileServer(http.Dir("./")))) //在浏览器地址栏输入localhost:8123/doc ,显示同上面一样的结果
</pre>
####template包
template包（html/template）实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出。本包提供了和text/template包相同的接口，无论何时当输出是HTML的时候都应使用本包。

####字段操作
Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象，这和Java或者C++中的this类似。

当前对象为struct类型时，对象的字段通过{{.FieldName}}读取，但是需要注意一点：这个字段必须是导出的(字段首字母必须是大写的)，否则在渲染的时候就会报错。这是因为对象的属性要遵循访问修饰符规则，私有属性外部不可访问，所以，会产生错误！

当前对象为Map类型时，对象的字段通过{{.fieldName}}读取，这个字段则没有上述的限制。
####OutputJson()
<pre>
func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
    out := &Result{ret, reason, i}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}
</pre>
####Golang发送Email邮件
<pre>
package main
import (
    "net/smtp"
    "fmt"
    "strings"
)
/*
 *  user : example@example.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */
func SendMail(user, password, host, to, subject, body, mailtype string) error{
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    var content_type string
    if mailtype == "html" {
        content_type = "Content-Type: text/"+ mailtype + "; charset=UTF-8"
    }else{
        content_type = "Content-Type: text/plain" + "; charset=UTF-8"
    }
 
    msg := []byte("To: " + to + "\r\nFrom: " + user + "<"+ user +">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    send_to := strings.Split(to, ";")
    err := smtp.SendMail(host, auth, user, send_to, msg)
    return err
}
func main() {
    user := "xxxx@163.com"
    password := "xxxx"
    host := "smtp.163.com:25"
    to := "xxxx@gmail.com;ssssss@gmail.com"
    subject := "Test send email by golang"
    body := `
    <html>
    <body>
    <h3>
    "Test send email by golang"
    </h3>
    </body>
    </html>
    `
    fmt.Println("send email")
    err := SendMail(user, password, host, to, subject, body, "html")
    if err != nil {
        fmt.Println("send mail error!")
        fmt.Println(err)
    }else{
        fmt.Println("send mail success!")
    }
}
</pre>
####filepath.Walk filepath.Abs
<pre>
package main

import (
	"log"
	"os"
	"fmt"
	"path/filepath"
)
func walkFunc(path string ,info os.FileInfo,err error)error{
	fmt.Println(path)
	return nil
}
func absFunc(){
	abs,err := filepath.Abs("/hello")//检查是否是绝对路径
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(abs)
}
func ReadDirectory(srcDir string) {
    files, _ := filepath.Glob(srcDir + "/[a-Z0-9]")
    fmt.Println(files)
}
func main(){
	filepath.Walk("./",walkFunc)
	absFunc()
}
</pre>
###golang中读写锁sync.RWMutex和互斥锁sync.Mutex区别
golang中sync包实现了两种锁Mutex （互斥锁）和RWMutex（读写锁），其中RWMutex是基于Mutex实现的，只读锁的实现使用类似引用计数器的功能．

- Mutex
定义：互斥锁是传统的并发程序对共享资源进行访问控制的主要手段。<b>互斥锁中Lock()加锁，Unlock()解锁，使用Lock()加锁后，便不能对其重复加锁，直到利用Unlock()对其解锁后，才能再次加锁；如果在使用Unlock()前未加锁，就会引起一个运行错误．</b><br>
适用场景：读写不确定，即读写次数没有明显的区别，并且只允许一个读或者写的场景，所有又称全局锁。<br>
示例:
<pre>
package main

import (
	"time"
	"fmt"
	"sync"
)
func main(){
	var mutex sync.Mutex
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")
	for i:=1;i<4;i++{
		go func(i int){
			fmt.Println("Not lock",i)
			mutex.Lock()
			fmt.Println("Locked",i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)
}
output==>
Lock the lock
The lock is locked
Not lock 1
Not lock 2
Not lock 3
Unlock the lock
Locked 1
</pre>
在需要频繁读，少量写的时候，Mutex的性能比使用channel要高很多，同时还能保证读写同步。
<pre>
package main

import (
	"fmt"
	"runtime"
	"sync"
)
type counter struct{
	mutex sync.Mutex
	x int64
}
func (c *counter) Inc(){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.x++
}
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := counter{}
	var wait sync.WaitGroup
	wait.Add(4)
	for k :=4;k> 0;k--{
		go func(){
			for i :=200000;i>0;i--{
				c.Inc()
			}
			wait.Done()
		}()
	}
	wait.Wait()
	fmt.Println(c.x)
}
output==>
800000
</pre>

- RWMutex
定义：<b>它允许任意读操作同时进行 同一时刻，只允许有一个写操作进行.并且一个写操作被进行过程中，读操作的进行也是不被允许的,读写锁控制下的多个写操作之间都是互斥的,写操作与读操作之间也都是互斥的,多个读操作之间却不存在互斥关系</b>.RWMutex是一个读写锁，该锁可以加多个读锁或者一个写锁。写锁，如果在添加写锁之前已经有其他的读锁和写锁，则lock就会阻塞直到该锁可用，为确保该锁最终可用，已阻塞的 Lock 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定。写锁解锁，如果没有进行写锁定，则就会引起一个运行时错误．注意：写解锁在进行的时候会试图唤醒所有因欲进行读锁定而被阻塞的Goroutine，也就是在所有写锁上锁之前存在的并且被迫停止的读锁将重新开始工作，读解锁在进行的时候只会在已无任何读锁定的情况下试图唤醒一个因欲进行写锁定而被阻塞的Goroutine。<br>
适用场景：经常用于读次数远远多于写次数的场景．<br>
示例：
<pre>
package main
//程序中RUnlock()个数不得多于Rlock()的个数
import (
	"fmt"
	"sync"
)
func main(){
	var g *sync.RWMutex
	g = new(sync.RWMutex)
	g.RLock()
	g.RLock()
	g.RUnlock()
	g.RLock()
	fmt.Println("g")
	g.RUnlock()	
}
output==>
g
</pre>
<pre>
package main
import (
    "fmt"
    "sync"
    "time"
    "os"
    "errors"
    "io"
)
type DataFile interface {
	//读取一个数据块
	Read()(rsn int64,d Data,err error)
	// 写入一个数据块
	Write(d Data)(wsn int64,err error)
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
    f *os.File  //文件
    fmutex sync.RWMutex //被用于文件的读写锁
    woffset int64 // 写操作需要用到的偏移量
    roffset int64 // 读操作需要用到的偏移量
    wmutex sync.Mutex // 写操作需要用到的互斥锁
    rmutex sync.Mutex // 读操作需要用到的互斥锁
    dataLen uint32 //数据块长度
}
//初始化DataFile类型值的函数,返回一个DataFile类型的值
func NewDataFile(path string, dataLen uint32) (DataFile, error){
    f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
    //f,err := os.Create(path)
    if err != nil {
        fmt.Println("Fail to find",f,"cServer start Failed")
        return nil, err
    }
 
    if dataLen == 0 {
        return nil, errors.New("Invalid data length!")
    }
 
    df := &myDataFile{
        f : f,
        dataLen:dataLen,
    }
 
    return df, nil
}
 
//获取并更新读偏移量,根据读偏移量从文件中读取一块数据,把该数据块封装成一个Data类型值并将其作为结果值返回
 
func (df *myDataFile) Read() (rsn int64, d Data, err error){
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
 
func (df *myDataFile) Write(d Data) (wsn int64, err error){
    //读取并更新写的偏移量
    var offset int64
    df.wmutex.Lock()
    offset = df.woffset
    df.woffset += int64(df.dataLen)
    df.wmutex.Unlock()
 
    //写入一个数据块,最后写入数据块的序号
    wsn = offset / int64(df.dataLen)
    var bytes []byte
    if len(d) > int(df.dataLen){
        bytes = d[0:df.dataLen]
    }else{
        bytes = d
    }
    df.fmutex.Lock()
    df.fmutex.Unlock()
    _, err = df.f.Write(bytes)
 
    return
}
 
func (df *myDataFile) Rsn() int64{
    df.rmutex.Lock()
    defer df.rmutex.Unlock()
    return df.roffset / int64(df.dataLen)
}
 
func (df *myDataFile) Wsn() int64{
    df.wmutex.Lock()
    defer df.wmutex.Unlock()
    return df.woffset / int64(df.dataLen)
}
 
func (df *myDataFile) DataLen() uint32 {
    return df.dataLen
}
 
func main(){
    //简单测试下结果
    var dataFile DataFile
    dataFile,_ = NewDataFile("./mutex_2016_3-21.dat", 10)
 
    var d=map[int]Data{
        1:[]byte("batu_test1"),
        2:[]byte("batu_test2"),
        3:[]byte("test1_batu"),
    }
 
    //写入数据
    for i:= 1; i < 4; i++ {
        go func(i int){
            wsn,_ := dataFile.Write(d[i])
            fmt.Println("write i=",i,",wsn=",wsn," ,success.")
        }(i)
    }
 
    //读取数据
    for i:= 1; i < 4; i++ {
        go func(i int){
            rsn,d,_ := dataFile.Read()
            fmt.Println("Read i=",i,",rsn=",rsn,",data=",d," success.")
        }(i)
    }
 
    time.Sleep(10 * time.Second)
}
output==>
write i= 1 ,wsn= 0  ,success.
write i= 2 ,wsn= 1  ,success.
write i= 3 ,wsn= 2  ,success.
Read i= 1 ,rsn= 0 ,data= [98 97 116 117 95 116 101 115 116 49]  success.
Read i= 2 ,rsn= 1 ,data= [98 97 116 117 95 116 101 115 116 50]  success.
Read i= 3 ,rsn= 2 ,data= [116 101 115 116 49 95 98 97 116 117]  success.
</pre>
###Goroutine高并发安全性
多个并发routine对一个共享变量进行操作有两种方法，channel和锁。
这里当然使用channel也能起到原子操作的效果。sync包的atomic和sync的mutex都是锁的方式。
所以说这里其实可以使用channel，mutex，atomic三种方法。


这里有一个经典的例子，介绍关于平时不会出现而在高并发情况下会出现的问题：
<pre>
package main 

import (
	"strings"
	"os"
	"runtime"
	"fmt"
	"math/rand"
	"time"
)
var total_tickets int32 = 20
func sell_tickets(i int){
	for {
		if total_tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(5))*time.Millisecond)
			total_tickets--
			fmt.Println("id:",i,"ticket:",total_tickets)
		}else{
			break
		}
	}
}
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix())
	//产生5个goroutine来卖票
	for i:=0;i<5;i++{
		go sell_tickets(i)
	}
	A:
	var input string
	fmt.Scanln(&input)
	if(strings.ToLower(input) == "!wq"){
		fmt.Println("Exit")
		os.Exit(0)
	}else{
		goto A
	}
	
}
output==>
id: 3 ticket: 19
id: 2 ticket: 18
id: 1 ticket: 17
id: 3 ticket: 16
id: 4 ticket: 15
id: 3 ticket: 14
id: 3 ticket: 13
id: 0 ticket: 12
id: 2 ticket: 11
id: 1 ticket: 10
id: 1 ticket: 8
id: 2 ticket: 9
id: 4 ticket: 7
id: 3 ticket: 6
id: 3 ticket: 5
id: 2 ticket: 4
id: 1 ticket: 3
id: 3 ticket: 2
id: 0 ticket: 1
id: 2 ticket: 0
id: 4 ticket: -1
id: 3 ticket: -2
id: 1 ticket: -3
!wq
Exit
</pre>
上面出现ticket=-1还有-2的情况肯定是不愿意看到的，那并发安全的应该怎么写呢？
当然答案已经有人说出来了，第一种方案是：<br>
是在每个goroutine上加一把锁保证数据同步.
<pre>
package main
import (
	"os"
	"strings"
	"runtime"
	"fmt"
	"math/rand"
	"time"
	"sync"
)
var total_tickets int32 = 20
var mutex = &sync.Mutex{}
func sell_tickets(i int){
	for total_tickets >0 {
		mutex.Lock()
		if total_tickets >0 {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			total_tickets--
			fmt.Println("id:",i,"tickets:",total_tickets)
		}
		mutex.Unlock()
	}
}
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix())
	for i:=0;i<5;i++{
		go sell_tickets(i)
	}
	B:
	var input string
	fmt.Scanln(&input)
	if strings.ToLower(input) == "!q"{
		fmt.Println("Leaved:",total_tickets)
		os.Exit(0)
	}else{
		goto B
	}
}
output==>
id: 0 tickets: 19
id: 0 tickets: 18
id: 0 tickets: 17
id: 0 tickets: 16
id: 0 tickets: 15
id: 0 tickets: 14
id: 0 tickets: 13
id: 0 tickets: 12
id: 0 tickets: 11
id: 0 tickets: 10
id: 0 tickets: 9
id: 0 tickets: 8
id: 0 tickets: 7
id: 0 tickets: 6
id: 0 tickets: 5
id: 0 tickets: 4
id: 0 tickets: 3
id: 0 tickets: 2
id: 0 tickets: 1
id: 0 tickets: 0
!q
Leaved: 0
</pre>
第二种方案是：<br>
原子操作，保证数据同步。<br>
原子操作即是进行过程中不能被中断的操作。针对某个值的原子操作在被进行的过程中，CPU绝不会再去进行其他的针对该值的操作。为了实现这样的严谨性，原子操作仅会由一个独立的CPU指令代表和完成。
<pre>
package main
import (
	"fmt"
	"sync/atomic"
	"time"
)
func main(){
	var cnt uint32 = 0
	for i:=0;i<5;i++{
		go func(){
			//每个goroutine都做20次加1操作
			for i:=0;i<20;i++{
				time.Sleep(time.Millisecond)
				atomic.AddUint32(&cnt,1)
			}
		}()		
	}
	time.Sleep(time.Second)
		cntFinal :=atomic.LoadUint32(&cnt)
		fmt.Println("cnt:",cntFinal)
}
output==>
100
</pre>
atomic Add操作
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
	newI32 := atomic.AddInt32(&i32,3)
	fmt.Println("=====new i32 value=====")
	fmt.Println(i32)  //3
	fmt.Println(newI32)  //3

    var i64 int64
	fmt.Println("=====old i64 value=====")
	fmt.Println(i64)
	newI64 := atomic.AddInt64(&i64,-3)
	fmt.Println("=====new i64 value=====")
	fmt.Println(i64)  //-3
	fmt.Println(newI64) //-3
}
output==>
=====old i32 value=====
0
=====new i32 value=====
3
3
=====old i64 value=====
0
=====new i64 value=====
-3
-3
</pre>
atomic CompareAndSwap操作
<pre>
package main
import (
	"fmt"
	"sync/atomic"
)
var value int32
//不断地尝试原子地更新value的值,直到操作成功为止
func addValueFunc(delta int32){
	//在被操作值被频繁变更的情况下,CAS操作并不那么容易成功，只能利用for循环以进行多次尝试
	for {
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
func main()  {
	fmt.Println("======old value=======")
	fmt.Println(value)
	fmt.Println("======CAS value=======")
	addValueFunc(3)
	fmt.Println(value)
	fmt.Println("======Store value=======")
	atomic.StoreInt32(&value, 10)
	fmt.Println(value)
}
output==>
======old value=======
0
======CAS value=======
3
======Store value=======
10
</pre>
高并发下CAS:
<pre>
package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"fmt"
)
func main(){
	runtime.GOMAXPROCS(1000)
    n := 100000
    wg := new(sync.WaitGroup)
    wg.Add(n)
    
    j := int32(0)
	fmt.Println("开始j的值是：",j)
    for i := 0; i < n; i++{
        go func(){
            if atomic.CompareAndSwapInt32(&j, 0, 1) {
                fmt.Println("j to 1")
				fmt.Println("结束j的值是：",j)
            }
            wg.Done()
        }()
    }
    wg.Wait()

    fmt.Println("Done")
}
output==>
开始j的值是： 0
j to 1
结束j的值是： 1
Done
</pre>
###Golang发送邮件
<pre>
package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

const(
	HOST = "smtp.qq.com"
	SERVER_ADDR = "smtp.qq.com:25"
	USER = "jason@qq.com"
	PASSWORD = "xxxx"
)

type Email struct {
	to string "to"
	subject string "subject"
	msg string "msg"
}

func NewEmail(to,subject,msg string) *Email{
	return &Email{to:to,subject:subject,msg:msg}
}

func SendEmail(email *Email) error{
	auth := smtp.PlainAuth("",USER,PASSWORD,HOST)
	sendTo := strings.Split(email.to,";")
	done := make(chan error,1024)
	
	go func(){
		defer close(done)
		for _,v := range sendTo {
			//warning ; the last \r\n need twice , not only one .
			str := strings.Replace("From:"+USER+"~To :"+v+"~Subject:"+email.subject+"~Content-Type: text/plain;charset=UTF-8~","~","\r\n",-1)+"\r\n"+email.msg
			fmt.Println("Content:",str)
			err := smtp.SendMail(SERVER_ADDR,auth,USER,[]string{v},[]byte(str))
			if err != nil{
				fmt.Println("Send Error:",err)
			}
			done <- err
		}
	}()
		
	for i:=0;i<len(sendTo);i++{
		<- done
	}

	return nil
}

func main() {
	email := NewEmail("jason@qq.com","How Are you","Hello World , I am jason")
	err := SendEmail(email)
	fmt.Println("result:",err)
}
output==>
Content: From:jason@qq.com
To :jason@qq.com
Subject:How Are you
Content-Type: text/plain;charset=UTF-8

Hello World , I am jason
</pre>
又一个UDP案例：
<pre>
package main  
  
import (  
    "fmt"  
    "net"  
    "os"  
)  
  
func main() {  
    addr, err := net.ResolveUDPAddr("udp", ":6000")  
    if err != nil {  
        fmt.Println("net.ResolveUDPAddr fail.", err)  
        os.Exit(1)  
    }  
  
    conn, err := net.ListenUDP("udp", addr)  
    if err != nil {  
        fmt.Println("net.ListenUDP fail.", err)  
        os.Exit(1)  
    }  
    defer conn.Close()  
  
    for {  
        buf := make([]byte, 65535)  
        rlen, remote, err := conn.ReadFromUDP(buf)  
        if err != nil {  
            fmt.Println("conn.ReadFromUDP fail.", err)  
            continue  
        }  
        go handleConnection(conn, remote, buf[:rlen])  
    }  
}  
  
func handleConnection(conn *net.UDPConn, remote *net.UDPAddr, msg []byte) {  
    service_addr, err := net.ResolveUDPAddr("udp", ":6001")  
    if err != nil {  
        fmt.Println("net.ResolveUDPAddr fail.", err)  
        return  
    }  
  
    service_conn, err := net.DialUDP("udp", nil, service_addr)  
    if err != nil {  
        fmt.Println("net.DialUDP fail.", err)  
        return  
    }  
    defer service_conn.Close()  
  
    _, err = service_conn.Write([]byte("request servcie x"))  
    if err != nil {  
        fmt.Println("service_conn.Write fail.", err)  
        return  
    }  
  
    buf := make([]byte, 65535)  
    rlen, err := service_conn.Read(buf)  
    if err != nil {  
        fmt.Println("service_conn.Read fail.", err)  
        return  
    }  
  
    conn.WriteToUDP(buf[:rlen], remote)  
}  
</pre>
###长连接的golang toy_server
<pre>
package main  
  
import (  
    "fmt"  
    "net"  
    "os"  
    "strconv"  
    "time"  
)  
  
type Request struct {  
    isCancel bool  
    reqSeq   int  
    reqPkg   []byte  
    rspChan  chan<- []byte  
}  
  
func main() {  
    addr, err := net.ResolveUDPAddr("udp", ":6000")  
    if err != nil {  
        fmt.Println("net.ResolveUDPAddr fail.", err)  
        os.Exit(1)  
    }  
  
    conn, err := net.ListenUDP("udp", addr)  
    if err != nil {  
        fmt.Println("net.ListenUDP fail.", err)  
        os.Exit(1)  
    }  
    defer conn.Close()  
  
    reqChan := make(chan *Request, 1000)  
    go connHandler(reqChan)  
  
    var seq int = 0  
    for {  
        buf := make([]byte, 1024)  
        rlen, remote, err := conn.ReadFromUDP(buf)  
        if err != nil {  
            fmt.Println("conn.ReadFromUDP fail.", err)  
            continue  
        }  
        seq++  
        go processHandler(conn, remote, buf[:rlen], reqChan, seq)  
    }  
}  
  
func processHandler(conn *net.UDPConn, remote *net.UDPAddr, msg []byte, reqChan chan<- *Request, seq int) {  
    rspChan := make(chan []byte, 1)  
    reqChan <- &Request{false, seq, []byte(strconv.Itoa(seq)), rspChan}  
    select {  
    case rsp := <-rspChan:  
        fmt.Println("recv rsp. rsp=%v", string(rsp))  
    case <-time.After(2 * time.Second):  
        fmt.Println("wait for rsp timeout.")  
        reqChan <- &Request{isCancel: true, reqSeq: seq}  
        conn.WriteToUDP([]byte("wait for rsp timeout."), remote)  
        return  
    }  
  
    conn.WriteToUDP([]byte("all process succ."), remote)  
}  
  
func connHandler(reqChan <-chan *Request) {  
    addr, err := net.ResolveUDPAddr("udp", ":6001")  
    if err != nil {  
        fmt.Println("net.ResolveUDPAddr fail.", err)  
        os.Exit(1)  
    }  
  
    conn, err := net.DialUDP("udp", nil, addr)  
    if err != nil {  
        fmt.Println("net.DialUDP fail.", err)  
        os.Exit(1)  
    }  
    defer conn.Close()  
  
    sendChan := make(chan []byte, 1000)  
    go sendHandler(conn, sendChan)  
  
    recvChan := make(chan []byte, 1000)  
    go recvHandler(conn, recvChan)  
  
    reqMap := make(map[int]*Request)  
    for {  
        select {  
        case req := <-reqChan:  
            if req.isCancel {  
                delete(reqMap, req.reqSeq)  
                fmt.Println("CancelRequest recv. reqSeq=%v", req.reqSeq)  
                continue  
            }  
            reqMap[req.reqSeq] = req  
            sendChan <- req.reqPkg  
            fmt.Println("NormalRequest recv. reqSeq=%d reqPkg=%s", req.reqSeq, string(req.reqPkg))  
        case rsp := <-recvChan:  
            seq, err := strconv.Atoi(string(rsp))  
            if err != nil {  
                fmt.Println("strconv.Atoi fail. err=%v", err)  
                continue  
            }  
            req, ok := reqMap[seq]  
            if !ok {  
                fmt.Println("seq not found. seq=%v", seq)  
                continue  
            }  
            req.rspChan <- rsp  
            fmt.Println("send rsp to client. rsp=%v", string(rsp))  
            delete(reqMap, req.reqSeq)  
        }  
    }  
}  
  
func sendHandler(conn *net.UDPConn, sendChan <-chan []byte) {  
    for data := range sendChan {  
        wlen, err := conn.Write(data)  
        if err != nil || wlen != len(data) {  
            fmt.Println("conn.Write fail.", err)  
            continue  
        }  
        fmt.Println("conn.Write succ. data=%v", string(data))  
    }  
}  
  
func recvHandler(conn *net.UDPConn, recvChan chan<- []byte) {  
    for {  
        buf := make([]byte, 1024)  
        rlen, err := conn.Read(buf)  
        if err != nil || rlen <= 0 {  
            fmt.Println(err)  
            continue  
        }  
        fmt.Println("conn.Read succ. data=%v", string(buf))  
        recvChan <- buf[:rlen]  
    }  
}  
</pre>
###Select
<pre>
package main
import (
	"time"
	"fmt"
)
func main(){
    c1 := make(chan string)
    c2 := make(chan string) 
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
output==>
received one
received two
</pre>
###sync.Once的用法
sync.Once.Do(f func())是一个挺有趣的东西,能保证once只执行一次，无论你是否更换once.Do(xx)这里的方法,这个sync.Once块只会执行一次。
<pre>
package main  
  
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
除了上面的例子，下面一个也是有关于sync.Once的介绍：
<pre>
package main  
  
import (  
    "fmt"  
    "sync"  
    "time"  
)  
  
var counter int = 0  
  
func main() {  
    chls := make([]chan int, 10)  
    for i := 0; i < 10; i++ {  
  
        chls[i] = make(chan int)  
        go addCounter(chls[i])  
  
    }  
  
    for _, val := range chls {  
        counter += <-val  
    }  
  
    fmt.Println("---结果是:", counter)  
  
    //设置一个超时的chan,没有阻塞读超时，写超时，可由程序创建chan来判定是否有超时写入  
    timeout := make(chan bool, 1)  
    subChn := chls[:2]  
    for i := 0; i < 2; i++ {  
        go timeoutAdd(i, subChn[i], timeout)  
    }  
  
    countDonw := 2  
    for {  
        select {  
        case <-subChn[0]:  
            fmt.Println("not time out ")  
            countDonw--  
        case <-timeout:  
            fmt.Println("time out ")  
            countDonw--  
  
        }  
  
        if countDonw <= 0 {  
            break  
        }  
    }  
  
    //关闭channel  
    for idx, ch := range chls {  
        close(ch)  
        _, ok := <-ch  
        if !ok {  
            fmt.Println("close channel ", idx)  
        }  
    }  
  
    // //单向读channel  
    // onedirchl := make(<-chan int)  
    // //我来试试写操作，会有什么现象呢  
    // //go 会报invalid operation: onedirchl <- 1 (send to receive-only type <-chan int)  
    // onedirchl <- close(onedirchl)  
  
    //全局唯一性操作,大爱啊，想想java在做系统初始化只需要执行一次并且是多线程并发情况下的代码怎么写？  
    //Lock ？全局boolean的开关，我的神啊，复杂  
  
    var once sync.Once  
  
    completeChan := []chan bool{make(chan bool, 1), make(chan bool, 1)}  
  
    //注意啊这里一定要传入指针，不然会是once的一个副本  
    go initConifg(&once, func() {  
        fmt.Println("我是第一个初始化的channel!")  
        completeChan[0] <- true  
    })  
  
    go initConifg(&once, func() {  
        fmt.Println("我是第二个完成初始化的channel!")  
        completeChan[1] <- true  
    })  
  
    for _, ch := range completeChan {  
        <-ch  
        close(ch)  
    }  
  
}  
  
func initConifg(once *sync.Once, handler func()) {  
    once.Do(func() {  
        time.Sleep(5e9)  
        fmt.Println("我这是初始化!,我等待了5S完成")  
    })  
  
    handler()  
  
}  
  
func timeoutAdd(index int, chl chan int, timeout chan bool) {  
    if index%2 != 0 {  
        time.Sleep(5e9)  
        fmt.Println("模拟超时了")  
        timeout <- true  
    } else {  
        fmt.Println("正常输出..")  
        chl <- 1  
    }  
  
}  
  
func addCounter(chl chan int) {  
    chl <- 1  
    fmt.Println("countting")  
  
}  
output==>
countting
---结果是: 10
countting
countting
countting
countting
countting
countting
countting
countting
countting
正常输出..
not time out 
模拟超时了
time out 
close channel  0
close channel  1
close channel  2
close channel  3
close channel  4
close channel  5
close channel  6
close channel  7
close channel  8
close channel  9
我这是初始化!,我等待了5S完成
我是第一个初始化的channel!
我是第二个完成初始化的channel!
</pre>
###临时对象池
主角是sync.Pool.我们可以把sync.Pool类型值看作是存放可被重复使用的值的容器。此类容器是自动伸缩的、高效的，同时也是并发安全的。为了描述方便，我们也会把sync.Pool类型的值称为临时对象池，而把存于其中的值称为对象值。

我们在用复合字面量初始化一个临时对象池的时候可以为它唯一的公开字段New赋值。该字段的类型是func() interface{}，即一个函数类型。可以猜到，被赋给字段New的函数会被临时对象池用来创建对象值。不过，实际上，该函数几乎仅在池中无可用对象值的时候才会被调用。

类型sync.Pool有两个公开的方法。一个是Get，另一个是Put。前者的功能是从池中获取一个interface{}类型的值，而后者的作用则是把一个interface{}类型的值放置于池中。

 通过Get方法获取到的值是任意的。如果一个临时对象池的Put方法未被调用过，且它的New字段也未曾被赋予一个非nil的函数值，那么它的Get方法返回的结果值就一定会是nil。我们稍后会讲到，Get方法返回的不一定就是存在于池中的值。不过，如果这个结果值是池中的，那么在该方法返回它之前就一定会把它从池中删除掉。

临时对象池与缓存池很类似，但是它却有着鲜明的特性。
第一个特性是：临时对象池可以把由其中的对象值产生的存储压力进行分摊。更进一步说，它会专门为每一个与操作它的Goroutine相关联的P都生成一个本地池。在临时对象池的Get方法被调用的时候，它一般会先尝试从与本地P对应的那个本地池中获取一个对象值。如果获取失败，它就会试图从其他P的本地池中偷一个对象值并直接返回给调用方。如果依然未果，那它只能把希望寄托于当前的临时对象池的New字段代表的那个对象值生成函数了。注意，这个对象值生成函数产生的对象值永远不会被放置到池中。它会被直接返回给调用方。另一方面，临时对象池的Put方法会把它的参数值存放到与当前P对应的那个本地池中。每个P的本地池中的绝大多数对象值都是被同一个临时对象池中的所有本地池所共享的。也就是说，它们随时可能会被偷走。
第二个突出特性：对垃圾回收友好。垃圾回收的执行一般会使临时对象池中的对象值被全部移除。也就是说，即使我们永远不会显式的从临时对象池取走某一个对象值，该对象值也不会永远待在临时对象池中。它的生命周期取决于垃圾回收任务下一次的执行时间。
<pre>
package main
import (
    "fmt"
    "runtime"
    "runtime/debug"
    "sync"
    "sync/atomic"
)
func main() {
    // 禁用GC，并保证在main函数执行结束前恢复GC
    defer debug.SetGCPercent(debug.SetGCPercent(-1))
    var count int32
    newFunc := func() interface{} {
        return atomic.AddInt32(&count, 1)
    }
    pool := sync.Pool{New: newFunc}
    // New 字段值的作用
    v1 := pool.Get()
    fmt.Printf("v1: %v\n", v1)
    // 临时对象池的存取
    pool.Put(newFunc())
    pool.Put(newFunc())
    pool.Put(newFunc())
    v2 := pool.Get()
    fmt.Printf("v2: %v\n", v2)
    // 垃圾回收对临时对象池的影响
    debug.SetGCPercent(100)
    runtime.GC()
    v3 := pool.Get()
    fmt.Printf("v3: %v\n", v3)
    pool.New = nil
    v4 := pool.Get()
    fmt.Printf("v4: %v\n", v4)
}
output==>
v1: 1
v2: 2
v3: 5
v4: <nil>
</pre>
在把nil赋给pool的New字段之前，即使手动的执行了垃圾回收，我们也是可以从临时对象池获取到一个对象值的。而在这之后，我们却只能取出nil。使用临时对象池的注意点：
首先，我们不能对通过Get方法获取到的对象值有任何假设。到底哪一个值会被取出是完全不确定的。这是因为我们总是不能得知操作临时对象池的Goroutine在哪一时刻会与哪一个P相关联，尤其是在比上述示例更加复杂的程序的运行过程中。在这种情况下，我们也就无从知晓我们放入的对象值会被存放到哪一个本地池中，以及哪一个Goroutine执行的Get方法会返回该对象值。所以，我们给予临时对象池的对象值生成函数所产生的值以及通过调用它的Put方法放入到池中的值都应该是无状态的或者状态一致的。从另一方面说，我们在取出并使用这些值的时候也不应该以其中的任何状态作为先决条件。这一点非常的重要。<br>
第二个需要注意的地方实际上与我们前面讲到的第二个特性紧密相关。临时对象池中的任何对象值都有可能在任何时候被移除掉，并且根本不会通知该池的使用方。这种情况常常会发生在垃圾回收器即将开始回收内存垃圾的时候。如果这时临时对象池中的某个对象值仅被该池引用，那么它还可能会在垃圾回收的时候被回收掉。因此，我们也就不能假设之前放入到临时对象池的某个对象值会一直待在池中，即使我们没有显式的把它从池中取出。甚至一个对象值可以在临时对象池中待多久，我们也无法假设。除非我们像前面的示例那样手动的控制GC的启停。不过，我们并不推荐这种方式。这会带来一些其他问题。
###ticker定时器
为了判断连接是否可用，通常我们会用timer机制来定时检测，这里要用ticker:
<pre>
ticker := time.NewTicker(60 * time.Second)
/*
使用一个60s的ticker，定时去ping，如果ping失败了，证明连接已经断开了，这时候就需要close了*/
for {
    select {
        case <-ticker.C:
            if err := ping(); err != nil {
                close()
            }
    }
}
</pre>
这套机制比较简单，也运行的很好，直到我们的服务器连上了10w+的连接。因为每一个连接都有一个ticker，所以同时会有大量的ticker运行，cpu一直在30%左右徘徊，性能不能让人接受。<br>
其实，我们只需要的是一套高效的超时通知机制。<br>
在go里面，channel是一个很不错的东西，我们可以通过close channel来进行broadcast。
<pre>
ch := make(bool)
/*
启动了10个goroutine，它们都会因为等待ch的数据而block，10s之后close这个channel，那么所有等待该channel的goroutine就会继续往下执行
*/
for i := 0; i < 10; i++ {
    go func() {
        println("begin")
        <-ch
        println("end")
    }
}

time.Sleep(10 * time.Second)

close(ch)
</pre>
通过channel这种close broadcast机制，我们可以非常方便的实现一个timer，timer有一个channel ch，所有需要在某一个时间 “T” 收到通知的goroutine都可以尝试读该ch，当T到达时候，close该ch，那么所有的goroutine都能收到该事件了。
<b>时间轮算法</b>：
<pre>
package timingwheel
//性能很好，转载自siddontang
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
</pre>
###优雅的关闭HTTP服务
go提供了一个ConnState的hook，我们能通过这个来获取到对应的connection，这样在服务结束的时候我们就能够close掉这个connection了。该hook会在如下几种ConnState状态的时候调用。

- StateNew：新的连接，并且马上准备发送请求了
- StateActive：表明一个connection已经接收到一个或者多个字节的请求数据，在 server调用实际的handler之前调用hook。
- StateIdle：表明一个connection已经处理完成一次请求，但因为是keepalived的，所以不会close，继续等待下一次请求。
- StateHijacked：表明外部调用了hijack，最终状态。
- StateClosed：表明connection已经结束掉了，最终状态。
<pre>
s.ConnState = func(conn net.Conn, state http.ConnState) {
    switch state {
    case http.StateNew:
        // 新的连接，计数加1
        s.wg.Add(1)
    case http.StateActive:
        // 有新的请求，从idle conn pool中移除
        s.mu.Lock()
        delete(s.conns, conn.LocalAddr().String())
        s.mu.Unlock()
    case http.StateIdle:
        select {
        case <-s.quit:
            // 如果要关闭了，直接Close，否则加入idle conn pool中。
            conn.Close()
        default:
            s.mu.Lock()
            s.conns[conn.LocalAddr().String()] = conn
            s.mu.Unlock()
        }
    case http.StateHijacked, http.StateClosed:
        // conn已经closed了，计数减一
        s.wg.Done()
    }
</pre>
当结束的时候，会走如下流程：
<pre>
func (s *Server) Close() error {
    // close quit channel, 广播我要结束啦
    close(s.quit)

    // 关闭keepalived，请求返回的时候会带上Close header。客户端就知道要close掉connection了。
    s.SetKeepAlivesEnabled(false)
    s.mu.Lock()

    // close listenser
    if err := s.l.Close(); err != nil {
        return err 
    }

    //将当前idle的connections设置read timeout，便于后续关闭。
    t := time.Now().Add(100 * time.Millisecond)
    for _, c := range s.conns {
        c.SetReadDeadline(t)
    }
    s.conns = make(map[string]net.Conn)
    s.mu.Unlock()

    // 等待所有连接结束
    s.wg.Wait()
    return nil
}
</pre>
通过以上方法，我们能从容的关闭server.
####可变参数args的地址跟实际外部slice的地址一样，用的同一个slice
<pre>
package main
//可变参数args的地址跟实际外部slice的地址一样，用的同一个slice
import (
	"fmt"
)
func t(args ...int){
	fmt.Printf("%p\n",args)
}
func main(){
	a :=[]int{1,2,3}
	b := a[1:]
	t(a...)
	t(b...)
	fmt.Printf("%p\n",a)
	fmt.Printf("%p\n",b)
}
</pre>
###Go使用pprof调试Goroutine
<pre>
package main
import (
    "net/http"
    "runtime/pprof"
)

var quit chan struct{} = make(chan struct{})

func f() {
    <-quit
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")

    p := pprof.Lookup("goroutine")
    p.WriteTo(w, 1)
}

func main() {
    for i := 0; i < 10000; i++ {
        go f()
    }

    http.HandleFunc("/", handler)
    http.ListenAndServe(":11181", nil)
}
//在浏览器访问localhost:11181，可以看到此时goroutine运行状况：
goroutine profile: total 10007
1 @ 0x474b7d 0x474914 0x47099c 0x40112f 0x452e68 0x454604 0x454f91 0x45299e 0x43fd01
#	0x474b7d	runtime/pprof.writeRuntimeProfile+0xdd	c:/go/src/runtime/pprof/pprof.go:540
#	0x474914	runtime/pprof.writeGoroutine+0xa4	c:/go/src/runtime/pprof/pprof.go:502
#	0x47099c	runtime/pprof.(*Profile).WriteTo+0xdc	c:/go/src/runtime/pprof/pprof.go:229
#	0x40112f	main.handler+0xdf			C:/mygo/src/act/main.go:18
#	0x452e68	net/http.HandlerFunc.ServeHTTP+0x48	c:/go/src/net/http/server.go:1265
#	0x454604	net/http.(*ServeMux).ServeHTTP+0x184	c:/go/src/net/http/server.go:1541
#	0x454f91	net/http.serverHandler.ServeHTTP+0x1a1	c:/go/src/net/http/server.go:1703
#	0x45299e	net/http.(*conn).serve+0xb5e		c:/go/src/net/http/server.go:1204
...
</pre>
可以看到，在main.f这个函数中，有10007个goroutine正在执行，符合我们的预期。
转自siddontang.com,感谢原作者:)
###在Go中使用JSON作为主要配置
- why json
主要的原因在于go的json包有一个杀手级别的RawMessage实现。<br>
RawMessage主要是用来告诉go延迟解析用的。当我们定义了某一个字段为RawMessage之后，go就不会解析这段json，这样我们就可以将其推后到各自的子模块单独解析。
假设有一个功能，后台存储可能是redis或者mysql，但是只会使用一个，可能我们会按照如下方式写配置：
<pre>
redis_store : {
    addr : 127.0.0.1
    db : 0
},

mysql_store : {
    addr : 127.0.0.1
    db : test
    password : admin
    user : root
}
store : redis
</pre>
对应的class为
<pre>
type Config struct {
    RedisStore struct {
        Addr string
        DB int
    }

    MysqlStore Struct {
        Addr string
        DB string
        Password string
        User string
    }

    Store string
}
</pre>
如果这时候我们在增加了一种新的store，我们需要在Config文件里面在增加一个新的field，但是实际我们只会使用一种store，并不需要写这么多的配置。
我们可以使用RawMessage来处理：
<pre>
type Config struct {
    Store string
    StoreConfig json.RawMessage
}
</pre>
如果使用redis，对应的配置文件为:
<pre>
store: redis
store_config: {
    addr : 127.0.0.1
    db : 0
}
</pre>
如果使用mysql，对应的配置文件为:
<pre>
store: mysql
store_config: {
    addr : 127.0.0.1
    db : test
    password : admin
    user : root
}
</pre>
go读取配置文件之后，并不会处理RawMessage对应的东西，而是由我们代码自己对应的store模块去处理。这样无论配置文件怎么变动，store模块做了什么变动，都不会影响Config类。
而在各个模块中，我们只需要自己定义相关config，然后可以将RawMessage直接解析映射到该config上面，譬如，对于redis，我们在模块中有如下定义:
<pre>
type RedisConfig config {
    Addr string `json:"addr"`
    DB int `json:"db"`
}

func NewConfig(m json.RawMessage) *RedisConfig {
    c := new(RedisConfig)

    json.Unmarshal(m, c)

    return c
}
</pre>
####json的不足
最大的问题就在于注释，在json中，可不能这样写：
<pre>
{
    //this is a comment
    /*this is a comment*/ 
}
</pre>
但是，我们又不可能不写一点注释来说明配置项是干啥的，所以，通常采用的是引入一个comment字段的方式，譬如：
<pre>
{
    "_comment" : "this is a comment",
    "key" : "value"
}
</pre>
另外，json还需要注意的就是写的时候最后一项不能加上逗号，这样的json会因为格式错误无法解析的。
<pre>
{
    "key" : "value",
}
</pre>
最后那个逗号可是不能要的，但是实际写配置的时候我们可是经常性的随手加上了,需要注意，不要犯这样的错误。
###Go Log模块开发
对于log的level，我们定义如下:
<pre>
const (
    LevelTrace = iota
    LevelDebug
    LevelInfo
    LevelWarn
    LevelError
    LevelFatal
)    
</pre>
相应的，提供如下几个函数:
<pre>
func Trace(format string, v ...interface{}) 
func Debug(format string, v ...interface{}) 
func Info(format string, v ...interface{}) 
func Warn(format string, v ...interface{}) 
func Error(format string, v ...interface{}) 
func Fatal(format string, v ...interface{}) 
</pre>
具体的代码如下：
<pre>
package log

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

//log level, from low to high, more high means more serious
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	Ltime  = 1 << iota //time format "2006/01/02 15:04:05"
	Lfile              //file.go:123
	Llevel             //[Trace|Debug|Info...]
)

var LevelName [6]string = [6]string{"Trace", "Debug", "Info", "Warn", "Error", "Fatal"}

const TimeFormat = "2006/01/02 15:04:05"

const maxBufPoolSize = 16

type atomicInt32 int32

func (i *atomicInt32) Set(n int) {
	atomic.StoreInt32((*int32)(i), int32(n))
}

func (i *atomicInt32) Get() int {
	return int(atomic.LoadInt32((*int32)(i)))
}

type Logger struct {
	level atomicInt32
	flag  int

	hMutex  sync.Mutex
	handler Handler

	quit chan struct{}
	msg  chan []byte

	bufMutex sync.Mutex
	bufs     [][]byte

	wg sync.WaitGroup

	closed atomicInt32
}

//new a logger with specified handler and flag
func New(handler Handler, flag int) *Logger {
	var l = new(Logger)

	l.level.Set(LevelInfo)
	l.handler = handler

	l.flag = flag

	l.quit = make(chan struct{})
	l.closed.Set(0)

	l.msg = make(chan []byte, 1024)

	l.bufs = make([][]byte, 0, 16)

	l.wg.Add(1)
	go l.run()

	return l
}

//new a default logger with specified handler and flag: Ltime|Lfile|Llevel
func NewDefault(handler Handler) *Logger {
	return New(handler, Ltime|Lfile|Llevel)
}

func newStdHandler() *StreamHandler {
	h, _ := NewStreamHandler(os.Stdout)
	return h
}

var std = NewDefault(newStdHandler())

func (l *Logger) run() {
	defer l.wg.Done()
	for {
		select {
		case msg := <-l.msg:
			l.hMutex.Lock()
			l.handler.Write(msg)
			l.hMutex.Unlock()
			l.putBuf(msg)
		case <-l.quit:
			//we must log all msg
			if len(l.msg) == 0 {
				return
			}
		}
	}
}

func (l *Logger) popBuf() []byte {
	l.bufMutex.Lock()
	var buf []byte
	if len(l.bufs) == 0 {
		buf = make([]byte, 0, 1024)
	} else {
		buf = l.bufs[len(l.bufs)-1]
		l.bufs = l.bufs[0 : len(l.bufs)-1]
	}
	l.bufMutex.Unlock()

	return buf
}

func (l *Logger) putBuf(buf []byte) {
	l.bufMutex.Lock()
	if len(l.bufs) < maxBufPoolSize {
		buf = buf[0:0]
		l.bufs = append(l.bufs, buf)
	}
	l.bufMutex.Unlock()
}

func (l *Logger) Close() {
	if l.closed.Get() == 1 {
		return
	}
	l.closed.Set(1)

	close(l.quit)

	l.wg.Wait()

	l.quit = nil

	l.handler.Close()
}

//set log level, any log level less than it will not log
func (l *Logger) SetLevel(level int) {
	l.level.Set(level)
}

// name can be in ["trace", "debug", "info", "warn", "error", "fatal"]
func (l *Logger) SetLevelByName(name string) {
	name = strings.ToLower(name)
	switch name {
	case "trace":
		l.SetLevel(LevelTrace)
	case "debug":
		l.SetLevel(LevelDebug)
	case "info":
		l.SetLevel(LevelInfo)
	case "warn":
		l.SetLevel(LevelWarn)
	case "error":
		l.SetLevel(LevelError)
	case "fatal":
		l.SetLevel(LevelFatal)
	}
}

func (l *Logger) SetHandler(h Handler) {
	if l.closed.Get() == 1 {
		return
	}

	l.hMutex.Lock()
	if l.handler != nil {
		l.handler.Close()
	}
	l.handler = h
	l.hMutex.Unlock()
}

func (l *Logger) Output(callDepth int, level int, s string) {
	if l.closed.Get() == 1 {
		// closed
		return
	}

	if l.level.Get() > level {
		// higher level can be logged
		return
	}

	buf := l.popBuf()

	if l.flag&Ltime > 0 {
		now := time.Now().Format(TimeFormat)
		buf = append(buf, '[')
		buf = append(buf, now...)
		buf = append(buf, "] "...)
	}

	if l.flag&Lfile > 0 {
		_, file, line, ok := runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		} else {
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					file = file[i+1:]
					break
				}
			}
		}

		buf = append(buf, file...)
		buf = append(buf, ':')

		buf = strconv.AppendInt(buf, int64(line), 10)
		buf = append(buf, ' ')
	}

	if l.flag&Llevel > 0 {
		buf = append(buf, '[')
		buf = append(buf, LevelName[level]...)
		buf = append(buf, "] "...)
	}

	buf = append(buf, s...)

	if s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}

	l.msg <- buf
}

//log with Trace level
func (l *Logger) Trace(v ...interface{}) {
	l.Output(2, LevelTrace, fmt.Sprint(v...))
}

//log with Debug level
func (l *Logger) Debug(v ...interface{}) {
	l.Output(2, LevelDebug, fmt.Sprint(v...))
}

//log with info level
func (l *Logger) Info(v ...interface{}) {
	l.Output(2, LevelInfo, fmt.Sprint(v...))
}

//log with warn level
func (l *Logger) Warn(v ...interface{}) {
	l.Output(2, LevelWarn, fmt.Sprint(v...))
}

//log with error level
func (l *Logger) Error(v ...interface{}) {
	l.Output(2, LevelError, fmt.Sprint(v...))
}

//log with fatal level
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(2, LevelFatal, fmt.Sprint(v...))
}

//log with Trace level
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Output(2, LevelTrace, fmt.Sprintf(format, v...))
}

//log with Debug level
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(2, LevelDebug, fmt.Sprintf(format, v...))
}

//log with info level
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(2, LevelInfo, fmt.Sprintf(format, v...))
}

//log with warn level
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(2, LevelWarn, fmt.Sprintf(format, v...))
}

//log with error level
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(2, LevelError, fmt.Sprintf(format, v...))
}

//log with fatal level
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(2, LevelFatal, fmt.Sprintf(format, v...))
}

func SetLevel(level int) {
	std.SetLevel(level)
}

// name can be in ["trace", "debug", "info", "warn", "error", "fatal"]
func SetLevelByName(name string) {
	std.SetLevelByName(name)
}

func SetHandler(h Handler) {
	std.SetHandler(h)
}

func Trace(v ...interface{}) {
	std.Output(2, LevelTrace, fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	std.Output(2, LevelDebug, fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	std.Output(2, LevelInfo, fmt.Sprint(v...))
}

func Warn(v ...interface{}) {
	std.Output(2, LevelWarn, fmt.Sprint(v...))
}

func Error(v ...interface{}) {
	std.Output(2, LevelError, fmt.Sprint(v...))
}

func Fatal(v ...interface{}) {
	std.Output(2, LevelFatal, fmt.Sprint(v...))
}

func Tracef(format string, v ...interface{}) {
	std.Output(2, LevelTrace, fmt.Sprintf(format, v...))
}

func Debugf(format string, v ...interface{}) {
	std.Output(2, LevelDebug, fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	std.Output(2, LevelInfo, fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	std.Output(2, LevelWarn, fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	std.Output(2, LevelError, fmt.Sprintf(format, v...))
}

func Fatalf(format string, v ...interface{}) {
	std.Output(2, LevelFatal, fmt.Sprintf(format, v...))
}
</pre>
###在go中使用linked channels进行数据广播
在go中channels是一个很强大的东西，但是在处理某些事情上面还是有局限的。其中之一就是一对多的通信。channels在多个writer，一个reader的模型下面工作的很好，但是却不能很容易的处理多个reader等待获取一个writer发送的数据的情况。<br>
处理这样的情况，可能的一个go api原型如下：
<pre>
type Broadcaster …
func NewBroadcaster() Broadcaster
func (b Broadcaster) Write(v interface{})
func (b Broadcaster) Listen() chan interface{}
</pre>
broadcast channel通过NewBroadcaster创建，通过Write函数进行数据广播。为了监听这个channel的信息，我们使用Listen，该函数返回一个新的channel去接受Write
发送的数据。
这套解决方案需要一个中间process用来处理所有reader的注册。当调用Listen创建新的channel之后，该channel就被注册，通常该中间process的主循环如下：
<pre>
for {
    select {
        case v := <-inc:
            for _, c := range(listeners) {
                c <- v
            }
        case c := <- registeryc:
            listeners.push(c)
    }
}
</pre>
这是一个通常的做法.但是该process在处理数据广播的时候会阻塞，直到所有的readers读取到值。一个可选的解决方式就是reader的channel是有buffer缓冲的，缓冲大小我们可以按需调节。或者当buffer满的时候我们将数据丢弃。
但是这篇blog并不是介绍上面这套做法的。这篇blog主要提出了另一种实现方式用来实现writer永远不阻塞，一个慢的reader并不会因为writer发送数据太快而要考虑分配太大的buffer。
虽然这么做不会有太高的性能，但是我并不在意，因为我觉得它很酷。我相信我会找到一个很好的使用地方的。
首先是核心的东西：
<pre>
type broadcast struct {
    c chan broadcast
    v interface{}
}
</pre>
这就是我说的linked channel，但是其实是Ouroboros data structure。也就是，这个struct实例在发送到channel的时候包含了自己。
从另一方面来说，如果我有一个chan broadcast类型的数据，那么我就能从中读取一个broadcast b，b.v就是writer发送的任意数据，而b.c，这个原先的chan broadcast，则能够让我重复这个过程。
另一个可能让人困惑的地方在于一个带有缓冲区的channel能够被用来当做一个1对多广播的对象。如果我定义如下的buffered channel：
<pre>
var c = make(chan T, 1)
</pre>
任何试图读取c的process都将阻塞直到有数据写入。
当我们想广播一个数据的时候，我们只是简单的将其写入这个channel，这个值只会被一个reader给获取，但是我们约定，只要读取到了数据，我们立刻将其再次放入该channel，如下：
<pre>
func wait(c chan T) T {
    v := <-c
    c <-v
    return v
}
</pre>
结合上面两个讨论的东西，我们就能够发现如果在broadcast struct里面的channel如果能够按照上面的方式进行处理，我们就能实现一个数据广播。如下：
<pre>
package broadcast

type broadcast struct {
    c   chan broadcast;
    v   interface{};
}

type Broadcaster struct {
    // private fields:
    Listenc chan chan (chan broadcast);
    Sendc   chan<- interface{};
}

type Receiver struct {
    // private fields:
    C chan broadcast;
}

// create a new broadcaster object.
func NewBroadcaster() Broadcaster {
    listenc := make(chan (chan (chan broadcast)));
    sendc := make(chan interface{});
    go func() {
        currc := make(chan broadcast, 1);
        for {
            select {
            case v := <-sendc:
                if v == nil {
                    currc <- broadcast{};
                    return;
                }
                c := make(chan broadcast, 1);
                b := broadcast{c: c, v: v};
                currc <- b;
                currc = c;
            case r := <-listenc:
                r <- currc
            }
        }
    }();
    return Broadcaster{
        Listenc: listenc,
        Sendc: sendc,
    };
}

// start listening to the broadcasts.
func (b Broadcaster) Listen() Receiver {
    c := make(chan chan broadcast, 0);
    b.Listenc <- c;
    return Receiver{<-c};
}

// broadcast a value to all listeners.
func (b Broadcaster) Write(v interface{})   { b.Sendc <- v }

// read a value that has been broadcast,
// waiting until one is available if necessary.
func (r *Receiver) Read() interface{} {
    b := <-r.C;
    v := b.v;
    r.C <- b;
    r.C = b.c;
    return v;
}
</pre>
测试代码如下：
<pre>
func TestBroadcast(t *testing.T) {
    b := NewBroadcaster()

    r := b.Listen()

    b.Write("hello")

    if r.Read().(string) != "hello" {
        t.Fatal("error string")
    }

    r1 := b.Listen()

    b.Write(123)

    if r.Read().(int) != 123 {
        t.Fatal("error int")
    }

    if r1.Read().(int) != 123 {
        t.Fatal("error int")
    }

    b.Write(nil)

    if r.Read() != nil {
        t.Fatal("error nit")
    }

    if r1.Read() != nil {
        t.Fatal("error nit")
    }
}
</pre>
###使用go reflect实现一套简易的rpc框架
在实际项目中，我们经常会碰到服务之间交互的情况，如何方便的与远端服务进行交互，就是一个需要我们考虑的问题。

通常，我们可以采用restful的编程方式，各个服务提供相应的web接口，相互之间通过http方式进行调用。或者采用rpc方式，约定json格式进行数据交互。

在我们的项目中，服务端对用户客户端提供的是restful的接口方式，而在服务器内部，我们则采用rpc方式进行服务之间的交互。

go语言本来就提供了jsonrpc的支持，所以自然开始我们就直接使用jsonrpc。jsonrpc的使用非常简单，对于调用端来说，就如同一个函数调用，如下：
<pre>
args := &Args{7, 8}
reply := new(Reply)
err := client.Call("Arith.Add", args, reply)
</pre>
上面是go jsonrpc自带的一个例子，可以看到，虽然我们通过call(rpcName, inParams, outParams)这样的形式可以很方便的进行rpc的调用，但是跟go实际的函数调用还是稍微有一点区别，对我来说，这么使用总觉得很别扭。
####自己实现
实现一套rpc框架需要考虑server，client以及包协议的问题。
 
- 包协议
我使用了最简单的包头 + 实际数据的做法，包头使用一个4字节的int表示后续数据的长度。而对于实际的rpc数据，我采用的是gob进行打包解包。

为什么选用gob而不是json？主要在于我不想自己做数据类型的转换，在json中，int类型的encode，decode会变成float类型的，如果函数需要的参数是int，json decode之后还需要我们自己根据参数实际的类型进行转换。增加了复杂度。而gob则在encode时候会加上实际的数据类型，这样decode之后我就能直接使用。

而且gob还支持注册自定义的类型，但是为了简单，建议只支持基本的数据类型，因为对于rpc来说，传递复杂的数据类型进行函数调用，我总觉得有点复杂，这在设计上面已经有问题了。

- server
在server需要解决的问题就是rpc函数注册并通过名字能进行该rpc函数调用。而这个通过reflect就能非常方便的实现，一个通过函数名字进行函数调用的例子：
<pre>
func Test(id int) (string, error) {
    return "abc", nil
}

funcmap  = map[string]reflect.Value{}

v := reflect.ValueOf(Test)

funcmap["test_rpc"] = v

args := []reflect.Value{reflect.ValueOf(10)}

funcmap["test_rpc"](args)
</pre>

- client
在client层，我们需要关注在声明一个rpc原型的函数变量之后，如何将其替换成另一个函数进行rpc调用。我们可以通过reflect的MakeFunc函数方便的做到，go自身的例子：
<pre>
swap := func(in []reflect.Value) []reflect.Value {
    return []reflect.Value{in[1], in[0]}
}

 makeSwap := func(fptr interface{}) {
    fn := reflect.ValueOf(fptr).Elem()
    v := reflect.MakeFunc(fn.Type(), swap)
    fn.Set(v)
}

var intSwap func(int, int) (int, int)
makeSwap(&intSwap)
fmt.Println(intSwap(0, 1))
</pre>
MakeFunc的原理在于，根据传入的函数变量的类型，创建一个新的函数，该函数调用的是我们指定的另一个函数。
同时，我们得到传入变量的指针，并用新的函数重新给该变量赋值。

- error处理
因为rpc调用可能会出现其他错误，譬如网络断线，gob encode错误等，client在调用的时候必须得处理这些错误，暴力的作法就是如果是这种内部错误，我们直接panic，但是我觉得太不友好，所以我们约定，所有的rpc函数在最后一个返回值必须是error。这样就是是rpc内部的错误，我们也能够通过error返回。

在注册rpc的时候，我们可以通过判断最后一个返回值是否是interface，同时是否具有Error函数来强制要求必须为error。如下:
<pre>
v := reflect.ValueOf(rpcFunc)

nOut := v.Type().NumOut()

if nOut == 0 || v.Type().Out(nOut-1).Kind() != reflect.Interface {
    err = fmt.Errorf("%s return final output param must be error interface", name)
    return
}

_, b := v.Type().Out(nOut - 1).MethodByName("Error")
if !b {
    err = fmt.Errorf("%s return final output param must be error interface", name)
    return
}
</pre>
但是，如果在MakeFunc里面直接返回error，会出现“reflect: function created by MakeFunc using closure returned wrong type: have *errors.errorString for error”这样的问题，主要在于reflect.Value需要知道我们error的接口类型.
所以，我们通过如下方式对error进行处理，转成相应的reflect.Value
<pre>
	v := reflect.ValueOf(&e).Elem()
</pre>

- nil处理
在实际rpc中，我们可能还会面临参数为nil的问题，如果直接对nil进行reflect.ValueOf，是得不到我们期望的类型的，这时候的Kind是0，reflect压根不能将其正确的转换成函数实际的类型。

当碰到nil的情况，我们只需要根据当前函数参数实际的类型，生成一个Zero Value，就可以很方便的解决这个问题：
假设函数第一个返回值为nil，那么我们这样
<pre>
v := reflect.Zero(fn.Type().Out(0))
</pre>

- 代码
原作者代码在<a href="https://github.com/siddontang/go/tree/master/rpc">感谢作者siddontang</a>
###StructTag 
如果希望手动配置结构体的成员和JSON字段的对应关系，可以在定义结构体的时候给成员打标签：
使用omitempty熟悉，如果该字段为nil或0值（数字0,字符串"",空数组[]等），则打包的JSON结果不会有这个字段。
<pre>
package main

import (
	"encoding/json"
	"fmt"
)
type Message struct {  
    Name string `json:"msg_name"`       // 对应JSON的msg_name  
    Body string `json:"body,omitempty"` // omitempty 如果为空置则忽略字段  
    Time int64  `json:"-"`              // 直接忽略字段  
}  
var m = Message{  
    Name: "Alice",  
    Body: "",  
    Time: 1294706395881547000,  
}  
func main(){
	data, err := json.Marshal(m)  
	if err != nil {  
	    fmt.Printf(err.Error())  
	    return  
	}  
	fmt.Println(string(data)) 
}
output==>
{"msg_name":"Alice"} 
</pre>
<pre>
package main
//go字符切片转成json
import (
	"encoding/json"
	"fmt"
	"os"
)
type ColorGroup struct {  
    ID     int  
    Name   string  
    Colors []string  
}  
func main(){
	group := ColorGroup{  
	    ID:     1,  
	    Name:   "Reds",  
	    Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},  
	}  
	b, err := json.Marshal(group)  
	if err != nil {  
	    fmt.Println("error:", err)  
	}  
	os.Stdout.Write(b)  
}
output==>
{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
</pre>
json转成go变量
<pre>
package main
//json转成go变量
import (
	"encoding/json"
	"fmt"
)
var jsonBlob = []byte(`[  
    {"Name": "Platypus", "Order": "Monotremata"},  
    {"Name": "Quoll",    "Order": "Dasyuromorphia"}  
]`)  
type Animal struct {  
    Name  string  
    Order string  
}  
func main(){
	var animals []Animal  
	err := json.Unmarshal(jsonBlob, &animals)  
	if err != nil {  
	    fmt.Println("error:", err)  
	}  
	fmt.Printf("%+v", animals) 
}
ouptut==>
[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
</pre>
json与结构体
结构体必须是<b>大写字母开头的成员</b>才会被JSON处理到，小写字母开头的成员不会有影响。
<pre>
package main
//json转成go变量
import (
	"encoding/json"
	"fmt"
)
type Message struct {  
    Name  string  
    Body  string  
    Time  int64  
    inner string  
}  
  
var m = Message{  
    Name:  "Alice",  
    Body:  "Hello",  
    Time:  1294706395881547000,  
    inner: "ok",  
}  
func main(){
	b := []byte(`{"nAmE":"Bob","Food":"Pickle", "inner":"changed"}`)  
	err := json.Unmarshal(b, &m)  
	if err != nil {  
	    fmt.Printf(err.Error())  
	    return  
	}  
	fmt.Printf("%v", m)  
}
output==>
{Bob Hello 1294706395881547000 ok}
</pre>
###复用Go内存buffer
为了理解Go的内存管理，分析一些Go运行时代码还是有必要的。Go程序中有两个独立的线程用来标记不再被程序使用的内存（这就是垃圾收集）并在其不再被使用时返还给操作系统（在Go代码中称为收割，scavenging）.

下面是一个小程序，会生成很多内存垃圾，每秒生成一个5MB到10MB的字节数组。它维护了一个20个这样字节数组大小的内存池，随机丢弃内存池中的字节数组。这个程序用来模拟程序中经常发生的场景：程序的各个部分每时每刻都会分配内存，一些分配的内存一直都在使用，大多数分配的内存都不再使用。在一个Go写的网络程序中，在处理网络链接或请求的Go协程里，这种情况很容易发生。常常是这样的，Go协程分配内存块（比如分配一个slices来存储接收的数据），然后就不再使用。随着时间的积累，会有一系列的内存块被正在被处理的网络链接占用，也会有一些累计的来自那些被处理过的链接的内存垃圾。
<pre>
package main
 
import (  
    "fmt"  
    "math/rand"  
    "runtime"  
    "time"
)
 
func makeBuffer() []byte {  
    return make([]byte, rand.Intn(5000000)+5000000)  
}
 
func main() {  
    pool := make([][]byte, 20)
 
    var m runtime.MemStats  
    makes := 0  
    for {  
        b := makeBuffer()  
        makes += 1
        i := rand.Intn(len(pool))
        pool[i] = b
 
        time.Sleep(time.Second)
 
        bytes := 0
 
        for i := 0; i < len(pool); i++ {
            if pool[i] != nil {
                bytes += len(pool[i])
            }
        }
 
        runtime.ReadMemStats(&m)
        fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
            m.HeapIdle, m.HeapReleased, makes)
    }
}
</pre>
这个程序使用runtime.ReadMemStats函数来获取堆大小的信息。这个函数会打印四个值：HeapSys （程序向操作系统请求的内存的字节数），HeapAlloc （当前堆中已经分配的字节数），HeapIdle （堆中未使用的字节数）和HeapReleased （归还给操作系统的字节数）。

Go程序中垃圾收集运行的很频繁（查看GOGC环境变量来理解如何控制GC操作 ）。因此，在运行过程中，堆的大小会随着内存被标记为未使用（这回导致HeapAlloc 和HeapIdle 随之变化）而变化。收割线程只有在内存5分钟都没有使用才会释放内存，因此HeapReleased 并不经常变化。

这类随着请求使用内存在垃圾收集程序中是很常见的（例如，论文Quantifying the Performance of Garbage Collection vs. Explicit Memory Management）。随着程序的运行，堆中未使用的内存又被重新利用，很少会被释放给操作系统。

解决这种问题的一个方法就是在程序中部分地手动管理内存。比如，使用一个管道，可以单独维护一个不再使用字节数组的内存池，当需要新的字节数组时，从内存池中拿（当内存池为空就生成新的字节数组）。

这个程序可以这样重写：
<pre>

package main
 
import (
    "fmt"
    "math/rand"
    "runtime"
    "time"
)
 
func makeBuffer() []byte {
    return make([]byte, rand.Intn(5000000)+5000000)
}
 
func main() {
    pool := make([][]byte, 20)
 
    buffer := make(chan []byte, 5)
 
    var m runtime.MemStats
    makes := 0
    for {
        var b []byte
        select {
        case b = <-buffer:
        default:
            makes += 1
            b = makeBuffer()
        }
 
        i := rand.Intn(len(pool))
        if pool[i] != nil {
            select {
            case buffer <- pool[i]:
                pool[i] = nil
            default:
            }
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
        fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
            m.HeapIdle, m.HeapReleased, makes)
    }
}
</pre>
这种内存复用机制的关键是一个缓存的管道buffer。上面的代码中可以存储5个字节数组。当程序需要一个字节数组时，优先使用select从缓存的管道中去取:
<pre>

select {
    case b = <-buffer:
    default:
        b = makeBuffer()
}
</pre>
select永远不会阻塞因为如果buffer 管道中有字节数组，第一个分支生效，字节数组赋给了 b。如果管道是空的话（也就意味着receive会阻塞），default 分支会执行，并分配了一个新的字节数组。把字节数组放回到管道中使用了类似的无阻塞模式:
<pre>	
select {
    case buffer <- pool[i]:
        pool[i] = nil
    default:
}
</pre>
如果buffer 管道已经满了，往管道里面发送就会阻塞。这种情况下，default分支执行，什么也不做。这种简单的机制可以用来安全的生成一个共享的内存池。由于管道通信对多go协程是安全的，这种机制也可以用于go协程的共享。

实际上，我们在Go程序中使用了类似的技术。下面的代码是真实复用器的简化版。使用一个go协程处理字节数组的生成并在软件中共享给所有的go协程。两个管道get （获取一个新的字节数组）和give （返回字节数组到内存池中）在所有的通信中都被使用。

复用器保存了一个返回的字节数组的链表，间断地丢弃那些时间太久，并不再会被复用（示例代码中，生命周期超过1分钟）的字节数组。这使得程序处理对字符数组的动态需求。
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
这些技术可以在程序员知道内存会被复用而不需要垃圾收集器参与时用来复用内存。它可以显著的减少程序需要内存的大小。并不仅限于字节数组。任何Go类型都可以用类似的行为进行复用。
###Go并发模式：管道和显式取消
Go并发原语使得构建流式数据管道，高效利用I/O和多核变得简单。
管道的阶段：

- 第一步：gen函数,是一个将数字列表转换到一个channel中的函数。Gen函数启动了一个goroutine，将数字发送到channel，并在所有数字都发送完后关闭channel；
<pre>
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
</pre>

- 第二步：sq，从上面的channel接收数字，并返回一个包含所有收到数字的平方的channel。在上游channel关闭后，这个阶段已经往下游发送完所有的结果，然后关闭输出channel：
<pre>
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
</pre>

- 第三步：main函数建立这个管道，并执行第一个阶段，从第二个阶段接收结果并逐个打印，直到channel被关闭。
<pre>
func main() {
    // Set up the pipeline.
    c := gen(2, 3)
    out := sq(c)
 
    // Consume the output.
    fmt.Println(<-out) // 4
    fmt.Println(<-out) // 9
}
</pre>
####扇出扇入（Fan-out, fan-in）
<b>扇出</b><br>
多个函数可以从同一个channel读取数据，直到这个channel关闭，这叫扇出。这是一种多个工作实例分布式地协作以并行利用CPU和I/O的方式。
<b>扇入</b><br>
一个函数可以从多个输入读取并处理数据，直到所有的输入channel都被关闭。这个函数会将所有输入channel导入一个单一的channel。这个单一的channel在所有输入channel都关闭后才会关闭。这叫做扇入.
###用Go语言写HTTP中间件
什么是中间件:在web开发过程中，中间件一般是指应用程序中封装原始信息，添加额外功能的组件。<br>
一个好的中间件拥有单一的功能，可插拔并且是自我约束的。这就意味着你可以在接口的层次上把它放到应用中，并能很好的工作。中间件并不影响你的代码风格，它也不是一个框架，仅仅是你处理请求流程中额外一层罢了。根本不需要重写代码：如果你想用一个中间件，就把它加上应用中；如果你改变主意了，去掉就好了。就这么简单。<br>
可以使用中间件做这些：

- 通过隐藏长度缓解BREACH攻击
- 频率限制
- 屏蔽恶意自动程序
- 提供调试信息
- 添加HSTS, X-Frame-Options头
- 从异常中优雅恢复
- 以及其他等等。
####写一个简单的中间件
写了一个中间件，只允许用户从特定的域（在HTTP的Host头中有域信息）来访问服务器。

定义类型

为了方便，让我们为这个中间件定义一种类型，叫做SingleHost。
<pre>
type SingleHost struct {
    handler     http.Handler
    allowedHost string
}
</pre>
只包含两个字段：

- 封装的Handler。如果是有效的Host访问，我们就调用这个Handler。
- 允许的主机值。

由于我们把字段名小写了，使得该字段只对我们自己的包可见。我们还应该写一个初始化函数
<pre>	
func NewSingleHost(handler http.Handler, allowedHost string) *SingleHost {
    return &SingleHost{handler: handler, allowedHost: allowedHost}
}
</pre>

处理请求

现在才是实际的逻辑。为了实现http.Handler，我们的类型秩序实现一个方法：
<pre>
type Handler interface {
     ServeHTTP(ResponseWriter, *Request)
}
</pre>
具体实现的方法：
<pre>
func (s *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    host := r.Host
    if host == s.allowedHost {
        s.handler.ServeHTTP(w, r)
    } else {
        w.WriteHeader(403)
    }
}
</pre>
ServeHTTP 函数仅仅检查请求中的Host头：

- 如果Host头匹配初始化函数设置的allowedHost ，就调用封装handler的ServeHTTP方法。
- 如果Host头不匹配，就返回403状态码（禁止访问）。

在后一种情况中，封装handler的ServeHTTP方法根本就不会被调用。因此封装的handler根本不会有任何输出，实际上它根本就不知道有这样一个请求到来。

现在我们已经完成了自己的中间件，来把它放到应用中。这次我们不把Handler直接放到net/http服务中，而是先把Handler封装到中间件中。
<pre>
singleHosted = NewSingleHost(myHandler, "example.com")
http.ListenAndServe(":8080", singleHosted)
</pre>

另外一种方法

刚才写的中间件实在是太简单了，只有仅仅15行代码。为了写这样的中间件，引入了一个不太通用的方法。由于Go支持函数第一型和闭包，并且拥有简洁的http.HandlerFunc包装器，我们可以将其实现为一个简单的函数，而不是写一个单独的类型。下面是基于函数的中间件版本。
<pre>
func SingleHost(handler http.Handler, allowedHost string) http.Handler {
    ourFunc := func(w http.ResponseWriter, r *http.Request) {
        host := r.Host
        if host == allowedHost {
            handler.ServeHTTP(w, r)
        } else {
            w.WriteHeader(403)
        }
    }
    return http.HandlerFunc(ourFunc)
}
</pre>

这里我们声明了一个叫做SingleHost的简单函数，接受一个Handler和允许的主机名。在函数内部，我们创建了一个类似之前版本ServeHTTP的函数。这个内部函数其实是一个闭包，所以它可以从SingleHost外部访问。最终，我们通过HandlerFunc把这个函数用作http.Handler。

使用Handler还是定义一个http.Handler类型完全取决于你。对简单的情况而已，一个函数就足够了。但是随着中间件功能的复杂，你应该考虑定义自己的数据结构，把逻辑独立到多个方法中。

实际上，标准库这两种方法都用了。StripPrefix 是一个返回HandlerFunc的函数。虽然TimeoutHandler也是一个函数，但它返回了处理请求的自定义的类型。

>>>总结

这篇文章的目的是吸引Go用户对中间件概念的注意以及展示使用Go写中间件的一些基本组件。尽管Go是一个相对年轻的开发语言，Go拥有非常漂亮的标准HTTP接口。这也是用Go写中间件是个非常简单甚至快乐的过程的原因之一。
###golang中的race检测
在本质上说，goroutine的使用增加了函数的危险系数论go语言中goroutine的使用。比如一个全局变量，如果没有加上锁，我们写一个比较庞大的项目下来，就根本不知道这个变量是不是会引起多个goroutine竞争。下面的是一个案例：
<pre>
package main

import(
    "time"
    "fmt"
    "math/rand"
)

func main() {
    start := time.Now()
    var t *time.Timer
    t = time.AfterFunc(randomDuration(), func() {
        fmt.Println(time.Now().Sub(start))
        t.Reset(randomDuration())
    })
    time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
    return time.Duration(rand.Int63n(1e9))
}
output==>
948.0543ms
1.0330591s
1.7000973s
1.9351107s
2.2231272s
2.7731587s
3.4061949s
3.7382139s
3.9222244s
4.405252s
</pre>
再比如下面的例子：
<pre>
package main
import (
	"time"
	"fmt"
)
func main(){
	a := 1
	go func(){
		a  = 2
	}()
	a = 3
	fmt.Println("a is ",a)
	time.Sleep(2 * time.Second)
}
</pre>
可喜的是，golang在1.1之后引入了竞争检测的概念。我们可以使用go run -race 或者 go build -race 来进行竞争检测。
golang语言内部大概的实现就是同时开启多个goroutine执行同一个命令，并且纪录每个变量的状态。
<pre>
runtime  go run -race race1.go
a is  3
==================
WARNING: DATA RACE
Write by goroutine 5:
  main.func·001()
      /Users/yejianfeng/Documents/workspace/go/src/runtime/race1.go:11 +0x3a

Previous write by main goroutine:
  main.main()
      /Users/yejianfeng/Documents/workspace/go/src/runtime/race1.go:13 +0xe7

Goroutine 5 (running) created at:
  main.main()
      /Users/yejianfeng/Documents/workspace/go/src/runtime/race1.go:12 +0xd7
==================
Found 1 data race(s)
exit status 66
</pre>
这个命令输出了Warning，告诉我们，goroutine5运行到第11行和main goroutine运行到13行的时候触发竞争了。而且goroutine5是在第12行的时候产生的。我们据此可以分析哪里出现了问题。
####Error和Fatal的区别
- Error ： Log() + Fail()  即记录当前错误，记录为失败，但是继续执行
- Fatal ： Log() + FailNow() 即记录当前错误，记录为失败，不继续执行
####linux下获取进程信息是使用/proc/pid/
####获取go的各种路径
1 执行用户当前所在路径：

os.Getwd()

2 执行程序所在路径：

执行程序文件相对路径：

file, _ := exec.LookPath(os.Args[0])
<pre>
package main
import(
        "os"
        "log"
        "os/exec"
        "path"
)
func main() {
        file, _ := os.Getwd()
        log.Println("current path:", file)
        file, _ = exec.LookPath(os.Args[0])
        log.Println("exec path:", file)
        dir,_ := path.Split(file)
        log.Println("exec folder relative path:", dir)
        os.Chdir(dir)
        wd, _ := os.Getwd()
        log.Println("exec folder absolute path:", wd)
}
output==>
2016/03/24 23:56:38 current path: C:\mygo\src\act
2016/03/24 23:56:38 exec path: C:\mygo\src\act\act.exe
2016/03/24 23:56:38 exec folder relative path: 
2016/03/24 23:56:38 exec folder absolute path: C:\mygo\src\act
</pre>
####从文件中json解析
第一种：

使用os.OpenFile直接获取reader，然后再从reader中使用Decoder来解析json
<pre>
package main
 
import (
    "fmt"
    "encoding/json"
    "os")
 
func main() {
    pathToFile := "jsondata.txt"
 
    file, err := os.OpenFile(pathToFile, os.O_RDONLY, 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
 
    configs := make(map[string]map[string][]Service, 0)
    err = json.NewDecoder(file).Decode(&configs)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }}
</pre>
第二种：
<pre>
content, err := ioutil.ReadFile(filepath)
    if err != nil {
        return nil, err
    }
 
    configs := make(map[string]map[string][]Service, 0)
    err = json.Unmarshal(content, configs)
    if err != nil {
        return nil, err
}
</pre>
###继承
<pre>
package main
import(
    "reflect"
)
type A struct {
}
func (self A)Run() {
    c := reflect.ValueOf(self)
    method := c.MethodByName("Test")
    println(method.IsValid())
}
type B struct {
    A
}
func (self B)Test(s string){
    println("b")
}
func main() {
    b := new(B)
    b.Run()
}
output==>
false
</pre>
<pre>
package main
import(
    "reflect"
)
type A struct {
    Parent interface{}
}
func (self A)Run() {
    c := reflect.ValueOf(self.Parent)
    method := c.MethodByName("Test")
    println(method.IsValid())
}
type B struct {
    A
}
func (self B)Test(s string){
    println("b")
}
func (self B)Run(){
    self.A.Run()
}
func main() {
    b := new(B)
    b.A.Parent = b
    b.Run()
}
output==>
true
</pre>
###http客户端
是Get，Post，PostForm三个函数。这三个函数直接实现了http客户端:
<pre>
package main
import (
    "fmt"
    "net/http"
    "io/ioutil"
)
 
func main() {
    response,_ := http.Get("http://www.baidu.com")
    defer response.Body.Close()
    body,_ := ioutil.ReadAll(response.Body)
    fmt.Println(string(body))
}
output==>
<!DOCTYPE html><!--STATUS OK--><html><head><meta http-equiv="content-type" content="text/html;charset=utf-8"><meta http-equiv="X-UA-Compatible" content="IE=Edge"><meta content="always" name="referrer"><meta name="theme-color" content="#2932e1"><link 
...
</pre>
http.Client和http.NewRequest来模拟请求:
<pre>
package main
 
import (
    "net/http"
    "io/ioutil"
    "fmt"
)
 
func main() {
    client := &http.Client{}
    reqest, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
     
    reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
    reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
    reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
    reqest.Header.Set("Cache-Control","max-age=0")
    reqest.Header.Set("Connection","keep-alive")
     
    response,_ := client.Do(reqest)
    if response.StatusCode == 200 {
        body, _ := ioutil.ReadAll(response.Body)
        bodystr := string(body);
        fmt.Println(bodystr)
    }
}
output==>
�
...
</pre>
###time
<pre>
package main
 
import (
    "fmt"
    "time"
)
 
func main() {
    //时间戳
    t := time.Now().Unix()
    fmt.Println(t)
     
    //时间戳到具体显示的转化
    fmt.Println(time.Unix(t, 0).String())
     
    //带纳秒的时间戳
    t = time.Now().UnixNano()
    fmt.Println(t)
    fmt.Println("------------------")
     
    //基本格式化的时间表示
    fmt.Println(time.Now().String())
     
    fmt.Println(time.Now().Format("2006年01月02日"))
 
}
output==>
1458916615
2016-03-25 22:36:55 +0800 +0800
1458916615370188100
------------------
2016-03-25 22:36:55.3701881 +0800 +0800
2016年03月25日
</pre>
输出标准时间
<pre>
package main
import (
    "fmt"
    "time"
)
func main() {
    //格式化字符串为时间
    test, _ := time.Parse("2006-01-02", "2020-03-21")
    //时间增加15秒
    after, _ := time.ParseDuration("15m")
    test = test.Add(after)
    fmt.Println(test)
    //格式化时间为字符串 标准时间格式
    t3 := test.Format("2006-01-02 15:04:05")
    fmt.Println(t3)
}
output==>
2020-03-21 00:15:00 +0000 UTC
2020-03-21 00:15:00
</pre>
输出星期
<pre>
package main
import(
    "fmt"
    "time"
)
func main() {
    //时间戳
    t := time.Now()
    fmt.Println(t.Weekday().String())
 
}
output==>
Friday
</pre>
下面是单核情况下4个goroutine并发
<pre>
package main
import (
	"fmt"
	"time"
)
var c chan int
func ready(w string, sec time.Duration){
	time.Sleep(sec * 1e9)
	fmt.Println(w,"is ready!")
	c <- 1
}
func main(){
	c = make(chan int)
	go ready("Tee",1)
	go ready("Coffee",1)
	go ready("Kele",1)
	go ready("Kele",1)
	fmt.Println("I am waiting")
	<- c
	<- c
	<- c
	<- c
}
output==>
I am waiting
Tee is ready!
Kele is ready!
Kele is ready!
Coffee is ready!
</pre>
###理解Golang中的panic、recover
Panic和Recover我们可以将他们看成是JAVA中的throw和catch.
<pre>
package main
import "fmt"
func main() {
    f()
    fmt.Println("Returned normally from f.")
}
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}
func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
output==>
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f.
</pre>

###条件变量
在Go语言中，sync.Cond类型代表了条件变量。与互斥锁和读写锁不同，简单的声明无法创建出一个可用的条件变量。为了得到这样一个条件变量，我们需要用到sync.NewCond函数。该函数的声明如下：
<pre>
func NewCond(l Locker) *Cond
</pre>
条件变量总是要与互斥量组合使用。因此，sync.NewCond函数的唯一参数是sync.Locker类型的，而具体的参数值既可以是一个互斥锁也可以是一个读写锁。sync.NewCond函数在被调用之后会返回一个*sync.Cond类型的结果值。我们可以调用该值拥有的几个方法来操纵对应的条件变量。

类型*sync.Cond的方法集合中有三个方法，即：Wait方法、Signal方法和Broadcast方法。它们分别代表了等待通知、单发通知和广播通知的操作。

方法Wait会自动的对与该条件变量关联的那个锁进行解锁，并且使调用方所在的Goroutine被阻塞。一旦该方法收到通知，就会试图再次锁定该锁。如果锁定成功，它就会唤醒那个被它阻塞的Goroutine。否则，该方法会等待下一个通知，那个Goroutine也会继续被阻塞。而方法Signal和Broadcast的作用都是发送通知以唤醒正在为此而被阻塞的Goroutine。不同的是，前者的目标只有一个，而后者的目标则是所有。

在Read方法中，我们使用一个for循环来达到重新尝试获取数据块的目的。为此，我们添加了若干条重复的语句、降低了程序的性能，还造成了一个潜在的问题——在某个情况下读写锁fmutex不会被读解锁。为了解决这一系列新生的问题，我们使用代表条件变量的字段rcond。
cond锁定期唤醒锁。cond的主要作用就是获取锁之后，wait()方法会等待一个通知，来进行下一步锁释放等操作，以此控制锁合适释放，释放频率。
案例：
<pre>
package main
import (
        "fmt"
        "sync"
        "time"
)
var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {
        cond.L.Lock() //获取锁
        cond.Wait()//等待通知  暂时阻塞
        fmt.Println(x)
        time.Sleep(time.Second * 1)
        cond.L.Unlock()//释放锁
}
func main() {
        for i := 0; i < 40; i++ {
                go test(i)
        }
        fmt.Println("start all")
        time.Sleep(time.Second * 3)
        fmt.Println("broadcast")
        cond.Signal()   // 下发一个通知给已经获取锁的goroutine
        time.Sleep(time.Second * 3)
        cond.Signal()// 3秒之后 下发一个通知给已经获取锁的goroutine
        time.Sleep(time.Second * 3)
        cond.Broadcast()//3秒之后 下发广播给所有等待的goroutine
        time.Sleep(time.Second * 60)
}
output==>
start all
broadcast
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
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
</pre>
在一个线程等待另一个线程这个场景里，条件变量和channel两种方式的区别。
使用sync.Cond
<pre>
package main
//使用sync.Cond
import (
	"fmt"
	"sync"
)
func main(){
	cv := sync.NewCond(new(sync.Mutex))
	done := false
	go func(){
		cv.L.Lock()
		done = true
		for i:=0;i< 5;i++{
			fmt.Println("I am doing something")
		}
		cv.Signal()
		cv.L.Unlock()
	}()
	//等待事情结束
	cv.L.Lock()
	for !done {
		cv.Wait()
	}
	cv.L.Unlock() //事情已经结束
}
output==>
I am doing something
I am doing something
I am doing something
I am doing something
I am doing something
</pre>
使用channel
<pre>
package main
import (
	"fmt"
)
func main(){
	done :=make(chan struct{})
	go func(){
		for i:=0;i<5;i++{
			fmt.Println("I am doing something")
		}
		close(done)
	}()
	<- done
	//事情已经结束
}
output==>
I am doing something
I am doing something
I am doing something
I am doing something
I am doing something
</pre>
对于很多情况下，我们既可以使用锁机制，也可以使用 Channel 来实现同一个目标，然而实际针对某个特定问题时，可能使用 Channel 会更加方便，但另外一些问题，使用锁机制会更加方便。Channel 和锁机制在 golang 中不是替代和被替代的关系,而是根据实际情况选择最方便的那一个。Use whichever is most expressive and/or most simple.
####channel 使用消息传递实现数据传递
直接使用消息传递实现更新：
<pre>
package main
//用消息传递实现更新
import (
	"fmt"
)
type updateup struct{
	key int
	value string
}
func applyupdate(data map[int]string,op updateup){
	data[op.key] = op.value
}
func main(){
	m :=make(map[int]string)
	m[2] = "Hello"
	
	ch :=make(chan updateup)
	go func(ch chan updateup){
		ch <- updateup{2,"New Value"}
	}(ch)
	applyupdate(m,<- ch)
	fmt.Printf("%s\n",m[2])
	fmt.Println(m)
}
output==>
New Value
map[2:New Value]
</pre>
###Golang构造一个并发安全的字典/Map类型
####基本思路：
Go语言提供的字典类型并不是并发安全的。因此，我们需要使用一些同步方法对它进行扩展。这看起来并不困难。我们只要使用读写锁将针对一个字典类型值的读操作和写操作保护起来就可以了。确实，读写锁应该是我们首先想到的同步工具。不过，我们还不能确定只使用它是否就足够了。不管怎样，让我们先来编写并发安全的字典类型的第一个版本。

- 先确定并发安全的字典类型的行为
我们可以借鉴OrderedMap接口类型的声明并编写出需要在这里声明的接口类型ConcurrentMap。实际上，ConcurrentMap接口类型的方法集合应该是OrderedMap接口类型的方法集合的一个子集。我们只需从OrderedMap中去除那些代表有序Map特有行为的方法声明即可。既然是这样，我何不从这两个自定义的字典接口类型中抽出一个公共接口呢？
<pre>
	// 泛化的Map的接口类型
	type GenericMap interface {
	// 获取给定键值对应的元素值。若没有对应元素值则返回nil
	Get(key interface{}) interface{}
	// 添加键值对，并返回与给定键值对应的旧的元素值。若没有旧元素值则返回(nil, true)
	Put(key interface{}, elem interface{}) (interface{}, bool)
	// 删除与给定键值对应的键值对，并返回旧的元素值。若没有旧元素值则返回nil
	Remove(key interface{}) interface{}
	// 清除所有的键值对
	Clear()
	// 获取键值对的数量
	Len() int
	// 判断是否包含给定的键值
	Contains(key interface{}) bool
	// 获取已排序的键值所组成的切片值
	Keys() []interface{}
	// 获取已排序的元素值所组成的切片值
	Elems() []interface{}
	// 获取已包含的键值对所组成的字典值
	ToMap() map[interface{}]interface{}
	// 获取键的类型
	KeyType() reflect.Type
	// 获取元素的类型
	ElemType() reflect.Type
	}
</pre>
然后，我们把这个名为GenericMap的字典接口类型嵌入到OrderedMap接口类型中，并去掉后者中的已在前者内声明的那些方法。修改后的OrderedMap接口类型如下：
<pre>
	// 有序的Map的接口类型
	type OrderedMap interface {
	GenericMap // 泛化的Map接口
	// 获取第一个键值。若无任何键值对则返回nil
	FirstKey() interface{}
	// 获取最后一个键值。若无任何键值对则返回nil
	LastKey() interface{}
	// 获取由小于键值toKey的键值所对应的键值对组成的OrderedMap类型值
	HeadMap(toKey interface{}) OrderedMap
	// 获取由小于键值toKey且大于等于键值fromKey的键值所对应的键值对组成的OrderedMap类型值
	SubMap(fromKey interface{}, toKey interface{}) OrderedMap
	// 获取由大于等于键值fromKey的键值所对应的键值对组成的OrderedMap类型值
	TailMap(fromKey interface{}) OrderedMap
	}
</pre>
有了GenericMap接口类型之后，我们的ConcurrentMap接口类型的声明就相当简单了。由于后者没有任何特殊的行为，所以我们只要简单地将前者嵌入到后者的声明中即可:
<pre>
type ConcurrentMap interface {
	GenericMap
}
</pre>
下面我们来编写该接口类型的实现类型。我们依然使用一个结构体类型来充当，并把它命名为myConcurrentMap。myConcurrentMap类型的基本结构如下：
<pre>
type myConcurrentMap struct {
    m     map[interface{}]interface{}
	keyType reflect.Type
	elemType reflect.Type
	rwmutex sync.RWMutex
}
</pre>
有了编写myOrderedMap类型（还记得吗？它的指针类型是OrderedMap的实现类型）的经验，写出myConcurrentMap类型的基本结构也是一件比较容易的事情。可以看到，在基本需要之外，我们只为myConcurrentMap类型加入了一个代表了读写锁的rwmutex字段。此外，我们需要为myConcurrentMap类型添加的那些指针方法的实现代码实际上也可以以myOrderedMap类型中的相应方法为蓝本。不过，在实现前者的过程中要注意合理运用同步方法以保证它们的并发安全性。下面，我们就开始编写它们。

首先，我们来看Put、Remove和Clear这几个方法。它们都属于写操作，都会改变myConcurrentMap类型的m字段的值。

方法Put的功能是向myConcurrentMap类型值添加一个键值对。那么，我们在这个操作的前后一定要分别锁定和解锁rwmutex的写锁。Put方法的实现如下：
<pre>
func (cmap *myConcurrentMap) Put(key interface{}, elem interface{}) (interface{}, bool) {
	if !cmap.isAcceptablePair(key, elem) {
	return nil, false
	}
	cmap.rwmutex.Lock()
	defer cmap.rwmutex.Unlock()
	oldElem := cmap.m[key]
	cmap.m[key] = elem
	return oldElem, true
}
</pre>
该实现中的isAcceptablePair方法的功能是检查参数值key和elem是否均不为nil且它们的类型是否均与当前值允许的键类型和元素类型一致。在通过该检查之后，我们就需要对rwmutex进行锁定了。相应的，我们使用defer语句来保证对它的及时解锁。与此类似，我们在Remove和Clear方法的实现中也应该加入相同的操作。

与这些代表着写操作的方法相对应的，是代表读操作的方法。在ConcurrentMap接口类型中，此类方法有Get、Len、Contains、Keys、Elems和ToMap。我们需要分别在这些方法的实现中加入对rwmutex的读锁的锁定和解锁操作。以Get方法为例，我们应该这样来实现它：
<pre>
func (cmap *myConcurrentMap) Get(key interface{}) interface{} {
	cmap.rwmutex.RLock()
	defer cmap.rwmutex.RUnlock()
	return cmap.m[key]
}
</pre>
这里有两点需要特别注意：<br>
我们在使用写锁的时候，要注意方法间的调用关系。比如，一个代表写操作的方法中调用了另一个代表写操作的方法。显然，我们在这两个方法中都会用到读写锁中的写锁。但如果使用不当，我们就会使前者被永远锁住。当然，对于代表写操作的方法调用代表读操作的方法的这种情况来说，也会是这样。请看下面的示例：
<pre>
func (cmap *myConcurrentMap) Remove(key interface{}) interface{} {
	cmap.rwmutex.Lock() 
	defer cmap.rwmutex.Unlock() 
	oldElem := cmap.Get()
	delete(cmap.m, key)
	return oldElem 
}
</pre>
可以看到，我们在Remove方法中调用了Get方法。并且，在这个调用之前，我们已经锁定了rwmutex的写锁。然而，由前面的展示可知，我们在Get方法的开始处对rwmutex的读锁进行了锁定。由于这两个锁定操作之间的互斥性，所以我们一旦调用这个Remove方法就会使当前Goroutine永远陷入阻塞。更严重的是，在这之后，其他Goroutine在调用该*myConcurrentMap类型值的一些方法（涉及到其中的rwmutex字段的读锁或写锁）的时候也会立即被阻塞住。

我们应该避免这种情况的方式。这里有两种解决方案。第一种解决方案是，把Remove方法中的oldElem := cmap.Get()语句与在它前面的那两条语句的位置互换，即变为：
<pre>
	oldElem := cmap.Get() 
	cmap.rwmutex.Lock()
	defer cmap.rwmutex.Unlock()
</pre>
这样可以保证在解锁读锁之后才会去锁定写锁。相比之下，第二种解决方案更加彻底一些，即：消除掉方法间的调用。也就是说，我们需要把oldElem := cmap.Get()语句替换掉。在Get方法中，体现其功能的语句是oldElem := cmap.m[key]。因此，我们把后者作为前者的替代品。若如此，那么我们必须保证该语句出现在对写锁的锁定操作之后。这样，我们才能依然确保其在锁的保护之下。实际上，通过这样的修改，我们升级了Remove方法中的被用来保护从m字段中获取对应元素值的这一操作的锁（由读锁升级至写锁）。

对于rwmutex字段的读锁来说，虽然锁定它的操作之间不是互斥的，但是这些操作与相应的写锁的锁定操作之间却是互斥的。我们在上一条注意事项中已经说明了这一点。因此，为了最小化对写操作的性能的影响，我们应该在锁定读锁之后尽快的对其进行解锁。也就是说，我们要在相关的方法中尽量减少持有读锁的时间。这需要我们综合的考量。

依据前面的示例和注意事项说明，读者可以试着实现Remove、Clear、Len、Contains、Keys、Elems和ToMap方法。它们实现起来并不困难。注意，我们想让*myConcurrentMap类型成为ConcurrentMap接口类型的实现类型。因此，这些方法都必须是myConcurrentMap类型的指针方法。这包括马上要提及的那两个方法。

方法KeyType和ElemType的实现极其简单。我们可以直接分别返回myConcurrentMap类型的keyType字段和elemType字段的值。这两个字段的值应该是在myConcurrentMap类型值的使用方初始化它的时候给出的。

按照惯例，我们理应提供一个可以方便的创建和初始化并发安全的字典值的函数。我们把它命名为NewConcurrentMap，其实现如下：
<pre>
func NewConcurrentMap(keyType, elemType reflect.Type) ConcurrentMap {
	return &amp;amp;myConcurrentMap{
	keyType: keyType,
	elemType: elemType,
	m:       make(map[interface{}]interface{})}
}
</pre>
这个函数并没有什么特别之处。由于myConcurrentMap类型的rwmutex字段并不需要额外的初始化，所以它并没有出现在该函数中的那个复合字面量中。此外，为了遵循面向接口编程的原则，我们把该函数的结果的类型声明为了ConcurrentMap，而不是它的实现类型*myConcurrentMap。如果将来我们编写出了另一个ConcurrentMap接口类型的实现类型，那么就应该考虑调整该函数的名称。比如变更为NewDefaultConcurrentMap，或者其他。
###Go语言内存模型
####名词定义
执行体 - Go里的Goroutine或Java中的Thread
####背景介绍
内存模型的目的是为了定义清楚变量的读写在不同执行体里的可见性。理解内存模型在并发编程中非常重要，因为代码的执行顺序和书写的逻辑顺序并不会完全一致，甚至在编译期间编译器也有可能重排代码以最优化CPU执行, 另外还因为有CPU缓存的存在，内存的数据不一定会及时更新，这样对内存中的同一个变量读和写也不一定和期望一样。

和Java的内存模型规范类似，Go语言也有一个内存模型，相对JMM来说，Go的内存模型比较简单，Go的并发模型是基于CSP（Communicating Sequential Process）的，不同的Goroutine通过一种叫Channel的数据结构来通信；Java的并发模型则基于多线程和共享内存，有较多的概念（violatie, lock, final, construct, thread, atomic等）和场景，当然java.util.concurrent并发工具包大大简化了Java并发编程。

Go内存模型规范了在什么条件下一个Goroutine对某个变量的修改一定对其它Goroutine可见。

####Happens Before
在一个单独的Goroutine里，对变量的读写和代码的书写顺序一致。比如以下的代码:
<pre>
package main
import (
    "log"
)
var a, b, c int
func main() {
    a = 1
    b = 2
    c = a + 2
    log.Println(a, b, c)
}
</pre>
尽管在编译期和执行期，编译器和CPU都有可能重排代码，比如，先执行b=2，再执行a=1，但c=a+2是保证在a=1后执行的。这样最后的执行结果一定是1 2 3，不会是1 2 2。但下面的代码则可能会输出0 0 0，1 2 2, 0 2 3 (b=2比a=1先执行), 1 2 3等各种可能。
<pre>
package main
import (
    "log"
)
var a, b, c int
func main() {
    go func() {
        a = 1
        b = 2
    }()
    go func() {
        c = a + 2
    }()
    log.Println(a, b, c)
}
</pre>
####Happens-before 定义
Happens-before用来指明Go程序里的内存操作的局部顺序。如果一个内存操作事件e1 happens-before e2，则e2 happens-after e1也成立；如果e1不是happens-before e2,也不是happens-after e2，则e1和e2是并发的。

在这个定义之下，如果以下情况满足，则对变量（v）的内存写操作（w）对一个内存读操作（r）来说允许可见的：

r不在w开始之前发生（可以是之后或并发）；
w和r之间没有另一个写操作(w’)发生；
为了保证对变量（v）的一个特定写操作（w）对一个读操作（r）可见，就需要确保w是r唯一允许的写操作，于是如果以下情况满足，则对变量（v）的内存写操作（w）对一个内存读操作（r）来说保证可见的：

w在r开始之前发生；
所有其它对v的写操作只在w之前或r之后发生；
可以看出后一种约定情况比前一种更严格，这种情况要求没有w或r没有其他的并发写操作。

在单个Goroutine里，因为肯定没有并发，上面两种情况是等价的。对变量v的读操作可以读到最近一次写操作的值（这个应该很容易理解）。但在多个Goroutine里如果要访问一个共享变量，我们就必须使用同步工具来建立happens-before条件，来保证对该变量的读操作能读到期望的修改值。

要保证并行执行体对共享变量的顺序访问方法就是用锁。Java和Go在这点上是一致的。

以下是具体的可被利用的Go语言的happens-before规则，从本质上来讲，happens-before规则确定了CPU缓冲和主存的同步时间点（通过内存屏障等指令），从而使得对变量的读写顺序可被确定–也就是我们通常说的“同步”。

####同步方法
初始化<br>

- 如果package p 引用了package q，q的init()方法 happens-before p （Java工程师可以对比一下final变量的happens-before规则）
- main.main()方法 happens-after所有package的init()方法结束。

创建Goroutine<br>
go语句创建新的goroutine happens-before 该goroutine执行（这个应该很容易理解）
<pre>
package main
import (
    "log"
    "time"
)
var a, b, c int
func main() {
    a = 1
    b = 2
    go func() {
        c = a + 2
        log.Println(a, b, c)
    }()
    time.Sleep(1 * time.Second)
}
</pre>
利用这条happens-before，我们可以确定c=a+2是happens-aftera=1和b=2，所以结果输出是可以确定的1 2 3，但如果是下面这样的代码，输出就不确定了，有可能是1 2 3或0 0 2
<pre>
func main() {
    go func() {
        c = a + 2
        log.Println(a, b, c)
    }()
    a = 1
    b = 2
    time.Sleep(1 * time.Second)
}
</pre>
销毁Goroutine<br>
Goroutine的退出并不保证happens-before任何事件。
<pre>
var a string
func hello() {
    go func() { a = "hello" }()
    print(a)
}
</pre>
上面代码因为a="hello" 没有使用同步事件，并不能保证这个赋值被主goroutine可见。事实上，极度优化的Go编译器甚至可以完全删除这行代码go func() { a = "hello" }()。

Goroutine对变量的修改需要让对其它Goroutine可见，除了使用锁来同步外还可以用Channel。

####Channel通信
在Go编程中，Channel是被推荐的执行体间通信的方法，Go的编译器和运行态都会尽力对其优化。

- 对一个Channel的发送操作(send) happens-before 相应Channel的接收操作完成
- 关闭一个Channel happens-before 从该Channel接收到最后的返回值0
- 不带缓冲的Channel的接收操作（receive） happens-before 相应Channel的发送操作完成
<pre>
var c = make(chan int, 10)
var a string
func f() {
    a = "hello, world"
    c <- 0
}
func main() {
    go f()
    <-c
    print(a)
}
</pre>
上述代码可以确保输出hello, world，因为a = "hello, world" happens-before c <- 0，print(a) happens-after <-c， 根据上面的规则1）以及happens-before的可传递性，a = "hello, world" happens-beforeprint(a)。

根据规则2）把c<-0替换成close(c)也能保证输出hello,world，因为关闭操作在<-c接收到0之前发送。
<pre>
var c = make(chan int)
var a string
func f() {
    a = "hello, world"
    <-c
}
func main() {
    go f()
    c <- 0
    print(a)
}
</pre>
根据规则3），因为c是不带缓冲的Channel，a = "hello, world" happens-before <-c happens-before c <- 0 happens-before print(a)， 但如果c是缓冲队列，如定义c = make(chan int, 1), 那结果就不确定了。
####锁
sync 包实现了两种锁数据结构:

- sync.Mutex -> java.util.concurrent.ReentrantLock
- sync.RWMutex -> java.util.concurrent.locks.ReadWriteLock
其happens-before规则和Java的也类似：

- 任何sync.Mutex或sync.RWMutex 变量（l），定义 n < m， 第n次 l.Unlock() happens-before 第m次l.lock()调用返回.
<pre>
var l sync.Mutex
var a string
func f() {
    a = "hello, world"
    l.Unlock()
}
func main() {
    l.Lock()
    go f()
    l.Lock()
    print(a)
}
</pre>
a = "hello, world" happens-before l.Unlock() happens-before 第二个 l.Lock() happens-before print(a)
####Once
sync包还提供了一个安全的初始化工具Once。还记得Java的Singleton设计模式，double-check，甚至triple-check的各种单例初始化方法吗？Go则提供了一个标准的方法。

- once.Do(f)中的f() happens-before 任何多个once.Do(f)调用的返回，且f()有且只有一次调用。
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
上面的代码虽然调用两次doprint()，但实际上setup只会执行一次，并且并发的once.Do(setup)都会等待setup返回后再继续执行。
###初识Golang之Groupcache
Groupcache是使用Go语言编写的缓存及缓存过滤库，作为memcached许多场景下的替代版本。同时，它也基于memcached进行了性能提升。对比老版本memcached，groupcache去掉了缓存有效期及缓存回收机制，随之而来的是通过自动备份来均衡负载。

首先，groupcache与memcached的相似之处：通过key分片，并且通过key来查询响应的peer。

其次，groupcache与memcached的不同之处：

1. 不需要对服务器进行单独的设置，这将大幅度减少部署和配置的工作量。groupcache既是客户端库也是服务器库，并连接到自己的peer上。
2. 具有缓存过滤机制。众所周知，在memcached出现“Sorry，cache miss（缓存丢失）”时，经常会因为不受控制用户数量的请求而导致数据库（或者其它组件）产生“惊群效应（thundering herd）”；groupcache会协调缓存填充，只会将重复调用中的一个放于缓存，而处理结果将发送给所有相同的调用者。
3. 不支持多个版本的值。如果“foo”键对应的值是“bar”，那么键“foo”的值永远都是“bar”。这里既没有缓存的有效期，也没有明确的缓存回收机制，因此同样也没有CAS或者Increment/Decrement。
4. 基于上一点的改变，groupcache就具备了自动备份“超热”项进行多重处理，这就避免了memcached中对某些键值过量访问而造成所在机器CPU或者NIC过载。
5. 当下只支持Go

运行机制

简而言之，groupcache查找一个Get（“foo”）的过程类似下面的情景（机器#5上，它是运行相同代码N台机器集中的一台）：

1.  key“foo”的值是否会因为“过热”而储存在本地内存，如果是，就直接使用
2.  key“foo”的值是否会因为peer #5是其拥有者而储存在本地内存，如果是，就直接使用
3.  首先确定key “fool”是否归属自己N个机器集合的peer中，如果是，就直接加载。如果有其它的调用者介入（通过相同的进程或者是peer的RPC请求，这些请求将会被阻塞，而处理结束后，他们将直接获得相同的结果）。如果不是，将key的所有者RPC到相应的peer。如果RPC失败，那么直接在本地加载（仍然通过备份来应对负载）。

使用情况

groupcache已经在dl.Google.com、Blogger、Google Code、Google Fiber、Google生产监视系统等项目中投入使用。
###Golang1.6新特性
<pre>
package main

import (
    "log"
    "os"
    "text/template"
)

var items = []string{"one", "two", "three"}

func tmplbefore15() {
    var t = template.Must(template.New("tmpl").Parse(`
    <ul>
    {{range . }}
        <li>{{.}}</li>
    {{end }}
    </ul>
    `))

    err := t.Execute(os.Stdout, items)
    if err != nil {
        log.Println("executing template:", err)
    }
}
/*go1.6新特性:{{-和-}}去除action前后的空白字符
func tmplaftergo16() {
    var t = template.Must(template.New("tmpl").Parse(`
    <ul>
    {{range . -}}
        <li>{{.}}</li>
    {{end -}}
    </ul>
    `))

    err := t.Execute(os.Stdout, items)
    if err != nil {
        log.Println("executing template:", err)
    }
}
*/
func main() {
    tmplbefore15()
    //tmplaftergo16()
}
output==>
//这是没有加新特性的效果，出现多余的空白
 <ul>

        <li>one</li>

        <li>two</li>

        <li>three</li>

 </ul>
</pre>
map race detect
Go原生的map类型是goroutine-unsafe的，长久以来，这给很多的gophers带来烦恼。这次go1.6中runtime增加了对并发访问map的检测以降低gopher使用哦map的心智负担。
<pre>
package main
/*
Go原生的map类型是goroutine-unsafe的，长久以来，这给很多Gophers
带来了烦恼。这次Go 1.6中Runtime增加了对并发访问map的检测以降低
gopher们使用map时的心智负担。该程序在go1.5上运行正常，在1.6则会报错。
*/
import "sync"

func main() {
    const workers = 100

    var wg sync.WaitGroup
    wg.Add(workers)
    m := map[int]int{}
    for i := 1; i <= workers; i++ {
        go func(i int) {
            for j := 0; j < i; j++ {
                m[i]++
            }
            wg.Done()
        }(i)
    }
    wg.Wait()
}
</pre>
###Golang经验
####range遍历
常见的range是：
<pre>
for k,v := range data{
	fmt.Println(v)	
}
</pre>
但是，实际上下面这个方法会更快，原因是这里面节省了v的拷贝，速度要比拷贝更快：
<pre>
for k,_ := range data {
    fmt.Println(data[k])
}
</pre>
####reflect反射
除非必要情况下，减少反射可以提升程序的整体性能。
####避免大量重复创建对象
连续小内存分配会导致大量的cpu消耗在scanblock这个函数上；连续make([]T, size)来分配slice还会额外导致cpu消耗在 runtime.memclr函数上。
####interface{}的使用
interface{}提供了golang中的interface类似于java的interface、PHP的interface或C++的纯虚基类。通过这种方式可以提供更快捷的编码。但是这种方式也带来了一些问题，最大的问题还是性能问题。
<pre>
// method 1
a.AA()
// method 2
v.(InterfaceA).AA()
// method 3
switch v.(type) {
case InterfaceA:
    v.(InterfaceA).AA()
}
</pre>
这三组方法性能逐个下降，最好的方式是直接进行类型引用,也就是第一种。
####指针传参效率更高
指针传参会减少对象复制过程，效率更高。具体原因：

- 节省存储，因为不用产生实际参数的函数局部副本
- 减轻函数调用的时间开销，因为不用调用拷贝复制等构造函数
- 允许函数有能力修改实际参数
<pre>
func Call(a *Struct) uint64 {
    return a.Ba
}
</pre>
golang中的指针：
<pre>
package main  
  
import "fmt"  
func main() {  
    var value int = 1  
    //指向int型的指针  
    var pInt *int = &value  
    //打印相关信息  
fmt.Printf("value = %d  pInt = %d  *pInt = %d \n", value, pInt, *pInt)  
  
    //通过指针修改指针指向的值  
    *pInt = 222  
fmt.Printf("value = %d  pInt = %d  *pInt = %d \n", value, pInt, *pInt)  
    //使指针指向别的地址  
    var m int = 123  
    pInt = &m  
    fmt.Printf("value = %d  pInt = %d  *pInt = %d \n", value, pInt, *pInt)  
}  
output==>
value = 1  pInt = 826814767824  *pInt = 1 
value = 222  pInt = 826814767824  *pInt = 222 
value = 222  pInt = 826814767904  *pInt = 123 
</pre>

####Go语言中有两个分配内存的机制
Go语言中有两个分配内存的机制，分别是内建的函数new和make.<br>
new(T)函数是一个分配内存的内建函数，但是不同于其他语言中内建new函数所做的工作，在Go语言中，new只是将内存清零，并没有初始化内存。所以在Go语言中，new(T)所做的工作是为T类型分配了值为零的内存空间并返回其地址，即返回*T。也就是说，new(T)返回一个指向新分配的类型T的零值指针.

make(T, args)函数与new(T)函数的目的不同。make(T, args)仅用于创建切片、map和channel(消息管道)，make(T, args)返回类型T的一个被初始化了的实例。而new（T）返回指向类型T的零值指针。也就是说new函数返回的是*T的未初始化零值指针，而make函数返回的是T的初始化了的实例.

Go语言中出现new和make两个分配内存的函数，并且使用起来还有差异，主要原因是切片、map、channel这三种类型在使用前必须初始化相关的数据结构。例如，切片是一个有三项内容的数据类型，包括指向数据的指针（在一个数组内部进行切片）、长度和容量，在这三项内容被初始化之前，切片值为nil。换句话说：对于切片、map、channel，make(T, args)初始化了其内部的数据结构并为他们准备了将要使用的值.
<pre>
package main  
  
import "fmt"  
  
func main() {  
    //使用new为切片分配内存 但是返回的是零值指针   
    //接着还是用使用make初始化 不必要的使问题复杂化 所以几乎不这样使用  
    var p *[]int = new([]int) //*p = nil  
    fmt.Println(p)        //输出&[]  
fmt.Println(*p)        //输出[]  
  
    //使用make为切片分配内存 在为切片分配内存时一般使用这种方法  
    var v []int = make([]int, 10)  
    fmt.Println(v)        //输出[0 0 0 0 0 0 0 0 0 0 0]  
}  
output==>
&[]
[]
[0 0 0 0 0 0 0 0 0 0]
</pre>
<pre>
package main

import "fmt"

type S map[string][]string

func Summary(param string) (s *S) {
  s = &S{
    "name": []string{param},
    "profession": []string{"Java programmer", "Project Manager"},
    "interest(lang)": []string{"Clojure", "Python", "Go"},
    "focus(project)": []string{"UE", "Agile Methodology", "Software Engineering"},
    "hobby(life)": []string{"Basketball", "Movies", "Travel"},
  }
  return s
}

func main() {
  s := Summary("Harry")
  fmt.Println("Summary(address)地址:",s)
  fmt.Printf("Summary(content)值: %v\n", *s)
}
output==>
Summary(address)地址: &map[focus(project):[UE Agile Methodology Software Engineering] hobby(life):[Basketball Movies Travel] name:[Harry] profession:[Java programmer Project Manager] interest(lang):[Clojure Python Go]]
Summary(content)值: map[profession:[Java programmer Project Manager] interest(lang):[Clojure Python Go] focus(project):[UE Agile Methodology Software Engineering] hobby(life):[Basketball Movies Travel] name:[Harry]]
成功: 进程退出代码 0.
</pre>
###Golang多维map读写操作的问题
关于map：
map中的元素不是变量，因此不能寻址。不能寻址的原因是：map可能会随着元素的增多重新分配更大的内存空间，旧值都会拷贝到新的内存空间，因此之前的地址就会失效。
<pre>
package main 

import "fmt" 
 
func main() { 
    m := make(map[int][2]int) 
    m[0] = [2]int{1, 3} 
    m[0][1] = 2 //错误 
    fmt.Println(m[0][1]) 
} 
output==>
 cannot assign to  m[0][1]
</pre>
因为在map[0]中的元素是array，当你把数组赋值给map时，传递的是
array的数值拷贝，也就是说map[0]中储存的是array的数值拷贝，
当你要修改m[0][1]的值，也就是要修改array[1]中的值时，
是不可寻址的，就是说go不知道array[1]的值对应的内存地址是什么，
所以导致赋值 cannot assign to  m[0][1]错误。

解决方法：

 1. 改为储存指针
 2. 把array改为slice or map这种天然引用类型的值。这样赋值的时候，go就能找到对应的变量储存地址，然后修改它。
<pre>
package main 
 
import "fmt" 
 
func main() { 
    m := make(map[int]map[int]int) 
    m[0] = map[int]int{1:3} 
    m[0][1] = 2 //正确 
    fmt.Println(m[0]) 
} 
output==>
m[1][2]
</pre>
再看一个类似的例子：
<pre>
package main 
 
import "fmt" 
 
type user struct { 
    name string 
    age  int 
    school map[string]school 
} 
 
type school struct { 
    Teacher string 
    Name string 
} 
 
func main() { 
    s := map[string]school{"primarySchool":{Teacher:"李老师", Name:"XX小学"}, "highSchool":{Teacher:"曹老师", Name:"XX中学"}} 
    u := user{name:"张三",age:12,school:s} 
    u.school["highSchool"].Name = "XX第二中学"//错误 
    fmt.Println(u) 
} 
output==>
cannot assign to u.school["highSchool"].Name
</pre>
错误分析：
原因出在user 中的 map[string]school  这里， u.school["highSchool"] 访问到这里都没有问题，问题在于后面的 “ .Name ” ，因为map[string]school 中储存的school是数值拷贝，当要修改school里面的Name时，就发生了不可寻址的错误。

解决方法：

1. 重新覆盖，既然无法单独修改里面的某一项，那就全部都替换掉，u.school["highSchool"] = school{Teacher:"曹老师", Name:"XX第二中学"}
2. 2. 改为储存指针，把map[string]school 改为 map[string]*school，把school的指针存进去，这样go就可以寻址，从而可以修改里面的值
###Gorouter一个轻量级高性能的路由(from[stutostu.com])

- 改善了url正则匹配的，使其匹配更多模式，更加可以自由定制
- 提高了匹配时查找的性能，使用路由的前缀和http方法做hashtable查找，路由再多，查找平均也是o(1)的时间复杂度
<pre>
package goRouter

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// http method
const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

//mime-types
const (
	applicationJson = "application/json"
	applicationXml  = "application/xml"
	textXml         = "text/xml"
)

type Mux struct {
	beforeMatch   http.HandlerFunc
	afterMatch    http.HandlerFunc
	beforeExecute http.HandlerFunc
	afterExecute  http.HandlerFunc
	routes        map[string]map[string][]*route
}

type route struct {
	pattern *regexp.Regexp
	params  []string
	handler http.HandlerFunc
}

var muxInstance = &Mux{
	beforeMatch:   func(rw http.ResponseWriter, req *http.Request) {},
	afterMatch:    func(rw http.ResponseWriter, req *http.Request) {},
	beforeExecute: func(rw http.ResponseWriter, req *http.Request) {},
	afterExecute:  func(rw http.ResponseWriter, req *http.Request) {},
	routes:        make(map[string]map[string][]*route),
}

var config = map[string]string{
	//路由匹配规则
	"matchReg": `^%s$`,
	//默认参数匹配规则
	"defaultParamsReg": `([^/]+)`,
	//查找参数规则
	"findParamsReg": `(:\w+)`,
	//处理没有带正则规则的参数规则
	"processParamsReg": `:\w+`,
	//处理带正则规则的参数规则
	"processParamsWithReg": `:\w+(\(.*?\))`,
}

func init() {
}

func GetMuxInstance() *Mux {
	return muxInstance
}

func (m *Mux) Get(pattern string, handler http.HandlerFunc) {
	m.AddRoute(pattern, handler, GET)
}

func (m *Mux) Post(pattern string, handler http.HandlerFunc) {
	m.AddRoute(pattern, handler, POST)
}

func (m *Mux) Put(pattern string, handler http.HandlerFunc) {
	m.AddRoute(pattern, handler, PUT)
}

func (m *Mux) Delete(pattern string, handler http.HandlerFunc) {
	m.AddRoute(pattern, handler, DELETE)
}

func (m *Mux) AddRoute(pattern string, handler http.HandlerFunc, method string) {
	pattern = strings.TrimRight(pattern, `/`)
	parts := strings.Split(pattern, `/`)
	var prefix []string
	for _, part := range parts {
		if strings.Index(part, ":") == -1 {
			prefix = append(prefix, part)
		} else {
			break
		}
	}

	//找出所有需要匹配的参数
	findParamReg := regexp.MustCompile(config["findParamsReg"])
	params := findParamReg.FindAllString(pattern, -1)

	//先处理带正则规则限定的参数
	replaceReg := regexp.MustCompile(config["processParamsWithReg"])
	pattern = replaceReg.ReplaceAllString(pattern, "$1")

	//没有正则限定的参数，使用默认正则规则来匹配
	replaceReg = regexp.MustCompile(config["processParamsReg"])
	pattern = replaceReg.ReplaceAllString(pattern, config["defaultParamsReg"])

	regex := regexp.MustCompile(fmt.Sprintf(config["matchReg"], pattern))

	if _, exist := m.routes[method]; !exist {
		m.routes[method] = map[string][]*route{}
	}

	prefixUrl := strings.Join(prefix, `/`)
	if _, exist := m.routes[method][prefixUrl]; !exist {
		m.routes[method][prefixUrl] = []*route{}
	}

	m.routes[method][prefixUrl] = append(m.routes[method][prefixUrl], &route{
		pattern: regex,
		handler: handler,
		params:  params,
	})
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.beforeMatch(w, r)

	handler, isMatch := m.match(strings.TrimRight(r.URL.Path, `/`), r)
	if !isMatch {
		http.NotFound(w, r)
		return
	}
	m.afterMatch(w, r)
	m.beforeExecute(w, r)
	handler(w, r)
	m.afterExecute(w, r)
}

func (m *Mux) match(requestPath string, r *http.Request) (http.HandlerFunc, bool) {
	paths := strings.Split(requestPath, `/`)

	if _, ok := m.routes[r.Method]; !ok {
		return nil, false
	}

	var path string
	for i := len(paths); i > 0; i-- {
		path = strings.Join(paths[:i], `/`)
		if routes, ok := m.routes[r.Method][path]; ok {
			for _, route := range routes {
				if !route.pattern.MatchString(requestPath) {
					continue
				}

				//whether need to match the parameters
				if len(route.params) > 0 {
					matches := route.pattern.FindStringSubmatch(requestPath)
					if len(matches) < 2 || len(matches[1:]) != len(route.params) {
						// panic("Parameters do not match")
						return nil, false
					}

					values := r.URL.Query()
					for i, match := range matches[1:] {
						values.Add(route.params[i], match)
					}

					//reassemble query params and add to RawQuery
					r.URL.RawQuery = url.Values(values).Encode() + "&" + r.URL.RawQuery
				}

				return route.handler, true
			}
		}
	}
	return nil, false
}
</pre>
使用案例：
<pre>
package main

import (
    "fmt"
    // import router
    "github.com/Barbery/goRouter"
    "net/http"
)

func main() {
    // get the instance of router
    mux := goRouter.GetMuxInstance()

    // add routes
    // Note: In goRouter, the routes is full match(by default, native router in golang is prefix match).
    mux.Get(`/user/:id(\d+)`, getUser)
    mux.Get(`/user/profile/:id(\d+)\.:format(\w+)`, getUserProfile)
    mux.Post(`/user`, postUser)
    mux.Delete(`/user/:id(\d+)`, deleteUser)
    mux.Put(`/user/:id(\d+)`, putUser)

    // run the serve
    http.ListenAndServe(":8888", mux)
}

// routes handler must be type of http.HandlerFunc
func getUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get(":id")
    w.Write([]byte(fmt.Sprintf(`GET user by id: %s`, id)))
}

func getUserProfile(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    w.Write([]byte(fmt.Sprintf(`GET user profile by id: %s, format: %s`, params[":id"][0], params.Get(":format"))))
}

func postUser(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    w.Write([]byte(fmt.Sprintf(`POST a new user, form data: %s`, fmt.Sprintln(r.PostForm))))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(fmt.Sprintf(`DELETE a user by id: %s`, r.URL.Query().Get(":id"))))
}

func putUser(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    w.Write([]byte(fmt.Sprintf(`UPDATE a user by id: %s, form data: %s`, r.URL.Query().Get(":id"), fmt.Sprintln(r.PostForm))))
}
</pre>
###动态规划问题
详情见：https://github.com/dongjun111111/blog/issues/11
<pre>
package main
import (
"fmt"
)

/*求最小服务时长，每次1单位1单位的切，得到的是最小解*/
func smin(n int32) int32 {
if n&1 == 0 {
return (n / 2) * (n - 1)
}
return (n - 1) / 2 * n
}

/*求每个顾客的时间*/
func serverTime(s, lenght []int32, maxLen int32) {
for i := range lenght {
s[i] = smin(lenght[i])
}
}

/*求二者最大值*/
func maxInt32(a, b int32) int32 {
if a > b {
return a
}
return b
}

func dptz(i, t int32, r, s []int32) int32 {
if i == 0 {
if t >= r[0]+s[0] {
return 1
}
return 0
}
if t >= r[i]+s[i] {
return maxInt32(dptz(i-1, r[i], r, s)+1, dptz(i-1, t, r, s))
}
return dptz(i-1, t, r, s)
}

/*求最后结束时间*/
func endTime(r, s []int32) int32 {
var max, tmp int32 = 0, 0
for i := range r {
tmp = r[i] + s[i]
if max < tmp {
max = tmp
}
}
return max
}

func main() {
//蛋糕长度、先来后到的时间和服务时间
length := []int32{2, 2, 3, 4}
r := []int32{5, 5, 6, 10}
s := make([]int32, 4)
serverTime(s, length, 4)
fmt.Println(dptz(3, endTime(r, s), r, s))
}
output==>
3
</pre>
###Golang之TCP粘包处理
####什么是粘包
一个完成的消息可能会被TCP拆分成多个包进行发送，也有可能把多个小的包封装成一个大的数据包发送，这个就是TCP的拆包和封包问题 
####TCP粘包和拆包产生的原因

- 应用程序写入数据的字节大小大于套接字发送缓冲区的大小
- 进行MSS大小的TCP分段。MSS是最大报文段长度的缩写。MSS是TCP报文段中的数据字段的最大长度。数据字段加上TCP首部才等于整个的TCP报文段。所以MSS并不是TCP报文段的最大长度，而是：MSS=TCP报文段长度-TCP首部长度
- 以太网的payload大于MTU进行IP分片。MTU指：一种通信协议的某一层上面所能通过的最大数据包大小。如果IP层有一个数据包要传，而且数据的长度比链路层的MTU大，那么IP层就会进行分片，把数据包分成托干片，让每一片都不超过MTU。注意，IP分片可以发生在原始发送端主机上，也可以发生在中间路由器上。
####TCP粘包和拆包的解决策略

- 消息定长。例如100字节。
- 在包尾部增加回车或者空格符等特殊字符进行分割，典型的如FTP协议
- 将消息分为消息头和消息尾。
- 其它复杂的协议，如RTMP协议等。
####处理方式
- 发送方在每次发送消息时将消息长度写入一个int32作为包头一并发送出去, 我们称之为Encode
- 接受方则先读取一个int32的长度的消息长度信息, 再根据长度读取相应长的byte数据, 称之为Decode
<pre>
package main

import (
    "bufio"
    "bytes"
    "encoding/binary"
)

func Encode(message string) ([]byte, error) {
    // 读取消息的长度
    var length int32 = int32(len(message))
    var pkg *bytes.Buffer = new(bytes.Buffer)
    // 写入消息头
    err := binary.Write(pkg, binary.LittleEndian, length)
    if err != nil {
        return nil, err
    }
    // 写入消息实体
    err = binary.Write(pkg, binary.LittleEndian, []byte(message))
    if err != nil {
        return nil, err
    }

    return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
    // 读取消息的长度
    lengthByte, _ := reader.Peek(4)
    lengthBuff := bytes.NewBuffer(lengthByte)
    var length int32
    err := binary.Read(lengthBuff, binary.LittleEndian, &length)
    if err != nil {
        return "", err
    }
    if int32(reader.Buffered()) < length+4 {
        return "", err
    }

    // 读取消息真正的内容
    pack := make([]byte, int(4+length))
    _, err = reader.Read(pack)
    if err != nil {
        return "", err
    }
    return string(pack[4:]), nil
}
</pre>
###Golang中slice不支持比较操作的原因
为什么Go语言不支持slice的比较运算呢？第一个原因，slice是引用类型，一个slice甚至可以引用自身。虽然有很多解决办法，但是没有一个是简单有效的。第二个原因，因为slice是间接引用，因此一个slice在不同时间可能包含不同的元素－底层数组的元素可能被修改。只要一个数据类型可以做相等比较，那么就可以用来做map的key,map这种数据结构对key的要求是：如果最开始时key是相等的，那在map的生命周期内，key要一直相等，因此这里key是不可变的。而对于指针或chan这类引用类型，==可以判断两个指针是否引用了想同的对象，是有用的，但是slice的相等测试充满了不确定性，因此，安全的做法是禁止slice之间的比较操作。

唯一例外：slice可以和nil进行比较，例如:
<pre>
if summer == nil { /* ... */ }
</pre>
slice是引用类型，因此它的零值是nil。一个nil slice是没有底层数组的，长度和容量都是0，但是也有非nil的slice,长度和容量也是0，例如[]int{}或make([]int, 3)[3:]。我们可以通过[]int(nil)这种类型转换生成一个[]int类型的nil值。
<pre>
var s []int    // len(s) == 0, s == nil
s = nil        // len(s) == 0, s == nil
s = []int(nil) // len(s) == 0, s == nil
s = []int{}    // len(s) == 0, s != nil
</pre>
由上可知，如果要测试一个slice是否为空，要使用len(s) == 0。除了可以和nil做相等比较外，nil slice的使用和0长度slice的使用方式相同：例如，前文的函数reverse(nil)就是安全的。除非包文档特别说明，否则所有的Go函数都应该以相同的方式对待nil slice和0长度slice(byte包中的部分函数会对nil值slice做特殊处理)。

###图灵机 验证Golang是图灵完备
go作为高级语言，图灵完备是必须的，我们用go来写一个最简单的图灵机。
<pre>
//hi
	++++++++++[>++++++++++<-]>++++.+.
</pre>
上面的仅有7条命令，理论上与任何图灵完备的语言等价。

1. + ==>  使当前数据单元值加1
2. - ==>  使当前数据单元值减1
3. > ==>  下一个单元作为当前数据单元
4. < ==>  上一个单元作为当前数据单元
5. [ ==>  如果当前数据单元的值为0，下一条指令对应的]后
6. ] ==>  如果当前数据单元的值不为0，下一条指令对应的[后
7. . ==>  把当前数据单元的值作为字符输出
但是程序员们使用不同语言的表达能力与效率不同，导致人们总是在探索新的语言的原因。
<pre>
package main

import (
	"fmt"
)
var (
	a [3000]byte
	prog = "++++++++++[>++++++++++<-]>++++.+."
	p,pc int
)
func loop(inc int){
	for i:=inc;i != 0;pc += inc{
		switch prog[pc+inc]{
			case '[':
			i++
			case ']':
			i--
		}
	}
}
func main(){
	for {
		switch prog[pc] {
			case '>':
			p++
			case '<':
			p--
			case '+':
			a[p]++
			case '-':
			a[p]--
			case '.':
			fmt.Print(string(a[p]))
			case '[':
			if a[p] == 0 {
				loop(1)
			}
			case ']':
			if a[p] != 0 {
				loop(-1)
			}
			default:
			fmt.Println("Illegal instruction")
		}
		pc++
		if pc == len(prog){
			return
		}
	}
}
output==>
hi
</pre>
程序一开始的变量a是我们图灵机的数据内存，prog是指令内存，p与pc分别是这两个内存的指针，代表当前数据单元和当前指令。
<pre>
package main

import (
	"os/exec"
	"io"
	"strconv"
	"net/http"
	"log"
	"os"
)
var uniq = make(chan int)
const frontPage =`
<!doctype html>
<html>
<head>

</head>
<body>
<textarea rows="15" cols="40" id="edit" spellcheck="false">
package main
import "fmt"
func main(){
	fmt.Println("hello world")
}
</textarea>
<button onclick="compile();">run</button>
<div id="output"></div>

</body>
</html>
`
func init(){
	go func(){
		for i:=0;;i++{
			uniq <- i
		}
	}()
	if err := os.Chdir(os.TempDir());err != nil{
		log.Fatal(err)
	}
	http.HandleFunc("/",FrontPage)
	http.HandleFunc("/compile",Compile)
	log.Fatal(http.ListenAndServe("127.0.0.1:1234",nil))
}
func FrontPage(w http.ResponseWriter,_ *http.Request){
	w.Write([]byte(frontPage))
}
func err(w http.ResponseWriter,e error) bool{
	if e != nil{
		w.Write([]byte(e.Error()))		
	}
	return true
}
func Compile(w http.ResponseWriter,req *http.Request){
	x := "play_" + strconv.Itoa(<-uniq) + ".go"
	f , e := os.Create(x)
	if err(w,e){
		return
	}
	defer os.Remove(x)
	defer f.Close()
	_,e =io.Copy(f,req.Body)
	if err(w,e){
		return
	}
	f.Close()
	cmd := exec.Command("go","run",x)
	o,e :=cmd.CombinedOutput()
	if err(w,e){
		return
	}
	w.Write(o)
}
func main(){
}
output==>
//127.0.0.1:1234/自己看看
</pre>
###Golang实现常见排序算法
<pre>
package main
import (
    "fmt"
)
func main() {
    //保存需要排序的Slice
    arr := []int{9, 3, 4, 7, 2, 1, 0, 11, 12, 11, 13, 4, 7, 2, 1, 0, 11, 12, 11}
    //实际用于排序的Slice
    list := make([]int, len(arr))

    copy(list, arr)
    BubbleSortX(list)
    fmt.Println("冒泡排序：\t", list)

    copy(list, arr)
    QuickSort(list, 0, len(arr)-1)
    fmt.Println("快速排序：\t", list)

    copy(list, arr) //将arr的数据覆盖到list，重置list
    InsertSort(list)
    fmt.Println("直接插入排序：\t", list)

    copy(list, arr)
    ShellSort(list)
    fmt.Println("希尔排序：\t", list)

    copy(list, arr)
    MergeSort(list)
    fmt.Println("二路归并排序：\t", list)

    copy(list, arr)
    SelectSort(list)
    fmt.Println("简单选择排序：\t", list)

    copy(list, arr)
    HeapSort(list)
    fmt.Println("堆排序：     \t", list)

}

//region 冒泡排序
//1，正宗的冒泡排序
/*
每趟排序过程中通过两两比较，找到第 i 个小（大）的元素，将其往上排。
*/
func BubbleSort(list []int) {
    var temp int // 用来交换的临时数
    var i int
    var j int
    // 要遍历的次数
    for i = 0; i < len(list)-1; i++ {
        // 从后向前依次的比较相邻两个数的大小，遍历一次后，把数组中第i小的数放在第i个位置上
        for j = len(list) - 1; j > i; j-- {
            // 比较相邻的元素，如果前面的数大于后面的数，则交换
            if list[j-1] > list[j] {
                temp = list[j-1]
                list[j-1] = list[j]
                list[j] = temp
            }
        }
    }
}

//2，冒泡排序优化
/*
对冒泡排序常见的改进方法是加入标志性变量exchange，用于标志某一趟排序过程中是否有数据交换。
如果进行某一趟排序时并没有进行数据交换，则说明所有数据已经有序，可立即结束排序，避免不必要的比较过程。
*/
func BubbleSortX(list []int) {
    var exchange bool = false
    var temp int // 用来交换的临时数
    var i int
    var j int
    // 要遍历的次数
    for i = 0; i < len(list)-1; i++ {
        // 从后向前依次的比较相邻两个数的大小，遍历一次后，把数组中第i小的数放在第i个位置上
        for j = len(list) - 1; j > i; j-- {
            // 比较相邻的元素，如果前面的数大于后面的数，则交换
            if list[j-1] > list[j] {
                temp = list[j-1]
                list[j-1] = list[j]
                list[j] = temp
                exchange = true
            }
        }
        if !exchange {
            break
        }
        exchange = false
    }
}

//endregion

//region 快速排序
func division(list []int, left int, right int) int {

    // 以最左边的数(left)为基准
    var base int = list[left]
    for left < right {
        // 从序列右端开始，向左遍历，直到找到小于base的数
        for left < right && list[right] >= base {
            right--
        }
        // 找到了比base小的元素，将这个元素放到最左边的位置
        list[left] = list[right]
        // 从序列左端开始，向右遍历，直到找到大于base的数
        for left < right && list[left] <= base {
            left++
        }
        // 找到了比base大的元素，将这个元素放到最右边的位置
        list[right] = list[left]

    }
    // 最后将base放到left位置。此时，left位置的左侧数值应该都比left小
    // 而left位置的右侧数值应该都比left大。
    list[left] = base //此时left == right
    //fmt.Println("DONE: base:", base, "\tleft:", left, "\tright:", right)
    return left
}

func QuickSort(list []int, left int, right int) {
    // 左下标一定小于右下标，否则就越界了
    if left < right {
        //对数组进行分割，取出下次分割的基准标号
        var base int = division(list, left, right)
        //对“基准标号“左侧的一组数值进行递归的切割，以至于将这些数值完整的排序
        QuickSort(list, left, base-1)
        //对“基准标号“右侧的一组数值进行递归的切割，以至于将这些数值完整的排序
        QuickSort(list, base+1, right)
    }

}

//endregion

//region 直接插入排序
func InsertSort(list []int) {
    var temp int
    var i int
    var j int
    // 第1个数肯定是有序的，从第2个数开始遍历，依次插入有序序列
    for i = 1; i < len(list); i++ {
        temp = list[i] // 取出第i个数，和前i-1个数比较后，插入合适位置
        // 因为前i-1个数都是从小到大的有序序列，所以只要当前比较的数(list[j])比temp大，就把这个数后移一位
        for j = i - 1; j >= 0 && temp < list[j]; j-- {
            list[j+1] = list[j]
        }
        list[j+1] = temp
    }
}

//endregion

//region 希尔排序
func ShellSort(list []int) {
    for gap := (len(list) + 1) / 2; gap >= 1; gap = gap / 2 {
        for i := 0; i < len(list)-gap; i++ {
            InsertSort(list[i:(gap + i + 1)]) //list[i:(gap + i + 1)]表示list索引i到gap+i的元素组成的slice
        }
    }
}

//region

//region 简单选择排序
/*
简单排序处理流程：
（1）从待排序序列中，找到关键字最小的元素；
（2）如果最小元素不是待排序序列的第一个元素，将其和第一个元素互换；
（3）从余下的 N - 1 个元素中，找出关键字最小的元素，重复（1）、（2）步，直到排序结束。
*/
func SelectSort(list []int) {
    var temp int
    var index int
    var i int
    var j int

    // 需要遍历获得最小值的次数
    // 要注意一点，当要排序 N 个数，已经经过 N-1 次遍历后，已经是有序数列
    for i = 0; i < len(list)-1; i++ {
        temp = 0
        index = i // 用来保存最小值得索引
        // 寻找第i个小的数值
        for j = i + 1; j < len(list); j++ {
            if list[index] > list[j] {
                index = j
            }
        }
        // 将找到的第i个小的数值放在第i个位置上
        temp = list[index]
        list[index] = list[i]
        list[i] = temp
    }
}

//endregion

//region 堆排序
func heapAdjust(list []int, parent int, length int) {
    temp := list[parent]  // temp保存当前父节点
    child := 2*parent + 1 // 先获得左孩子

    for child < length {
        // 如果有右孩子结点，并且右孩子结点的值大于左孩子结点，则选取右孩子结点
        if child+1 < length && list[child] < list[child+1] {
            child++
        }

        // 如果父结点的值已经大于孩子结点的值，则直接结束
        if temp >= list[child] {
            break
        }

        // 把孩子结点的值赋给父结点
        list[parent] = list[child]

        // 选取孩子结点的左孩子结点,继续向下筛选
        parent = child
        child = 2*child + 1
    }

    list[parent] = temp
}

func HeapSort(list []int) {
    // 循环建立初始堆
    for i := len(list) / 2; i >= 0; i-- {
        heapAdjust(list, i, len(list)-1)
    }

    // 进行n-1次循环，完成排序
    for i := len(list) - 1; i > 0; i-- {
        // 最后一个元素和第一元素进行交换
        temp := list[i]
        list[i] = list[0]
        list[0] = temp

        // 筛选 R[0] 结点，得到i-1个结点的堆
        heapAdjust(list, 0, i)
    }
}

//endregion

//region 归并排序(二路归并)
func merge(list []int, low int, mid int, high int) {
    var i int = low                  // i是第一段序列的下标
    var j int = mid + 1              // j是第二段序列的下标
    var k int = 0                    // k是临时存放合并序列的下标
    list2 := make([]int, high-low+1) // list2是临时合并序列
    // 扫描第一段和第二段序列，直到有一个扫描结束
    for i <= mid && j <= high {
        // 判断第一段和第二段取出的数哪个更小，将其存入合并序列，并继续向下扫描
        if list[i] <= list[j] {
            list2[k] = list[i]
            i++
            k++
        } else {
            list2[k] = list[j]
            j++
            k++
        }
    }
    // 若第一段序列还没扫描完，将其全部复制到合并序列
    for i <= mid {
        list2[k] = list[i]
        i++
        k++
    }

    // 若第二段序列还没扫描完，将其全部复制到合并序列
    for j <= high {
        list2[k] = list[j]
        j++
        k++
    }
    // 将合并序列复制到原始序列中
    k = 0
    for i = low; i <= high; i++ {
        list[i] = list2[k]
        k++
    }
}

func MergeSort(list []int) {
    for gap := 1; gap < len(list); gap = 2 * gap {
        var i int
        // 归并gap长度的两个相邻子表
        for i = 0; i+2*gap-1 < len(list); i = i + 2*gap {
            merge(list, i, i+gap-1, i+2*gap-1)
        }
        // 余下两个子表，后者长度小于gap
        if i+gap-1 < len(list) {
            merge(list, i, i+gap-1, len(list)-1)
        }
    }
}
output==>
冒泡排序： [0 0 1 1 2 2 3 4 4 7 7 9 11 11 11 11 12 12 13]
快速排序： [0 0 1 1 2 2 3 4 4 7 7 9 11 11 11 11 12 12 13]
直接插入排序： [0 0 1 1 2 2 3 4 4 7 7 9 11 11 11 11 12 12 13]
希尔排序： [0 0 1 1 2 2 3 4 4 7 7 9 11 11 11 11 12 12 13]
二路归并排序： [0 0 1 1 2 2 3 4 4 7 7 9 11 11 11 11 12 12 13]
简单选择排序： [0 0 1 1 2 2 3 4 4 7 7 9 11 11 11 11 12 12 13]
堆排序： [0 0 1 1 2 2 3 4 4 7 7 9 11 11 11 11 12 12 13]
</pre>
####链表排序
<pre>
package main 
import ( 
    "container/list"
    "fmt"
) 
type SortedLinkedList struct { 
    *list.List 
    Limit int
    compareFunc func (old, new interface{}) bool 
} 
func NewSortedLinkedList(limit int, compare func (old, new interface{}) bool) *SortedLinkedList { 
    return &SortedLinkedList{list.New(), limit, compare} 
} 
func (this SortedLinkedList) findInsertPlaceElement(value interface{}) *list.Element { 
    for element := this.Front(); element != nil; element = element.Next() { 
        tempValue := element.Value 
        if this.compareFunc(tempValue, value) { 
            return element 
        } 
    } 
    return nil 
} 
func (this SortedLinkedList) PutOnTop(value interface{}) { 
    if this.List.Len() == 0 { 
        this.PushFront(value) 
        return
    } 
    if this.List.Len() < this.Limit && this.compareFunc(value, this.Back().Value) { 
        this.PushBack(value) 
        return
    } 
    if this.compareFunc(this.List.Front().Value, value) { 
        this.PushFront(value) 
    } else if this.compareFunc(this.List.Back().Value, value) && this.compareFunc(value, this.Front().Value) { 
        element := this.findInsertPlaceElement(value) 
        if element != nil { 
            this.InsertBefore(value, element) 
        } 
    } 
    if this.Len() > this.Limit { 
        this.Remove(this.Back()) 
    } 
}

type WordCount struct { 
    Word  string 
    Count int
} 
func compareValue(old, new interface {}) bool { 
    if new.(WordCount).Count > old.(WordCount).Count { 
        return true
    } 
    return false
} 
func main() { 
    wordCounts := []WordCount{ 
        WordCount{"kate", 87}, 
        WordCount{"herry", 92}, 
        WordCount{"james", 81},
        WordCount{"jason",67},
        WordCount{"jack",97},
        WordCount{"bob",107}}
    var aSortedLinkedList = NewSortedLinkedList(10, compareValue) 
    for _, wordCount := range wordCounts { 
        aSortedLinkedList.PutOnTop(wordCount) 
    } 
    for element := aSortedLinkedList.List.Front(); element != nil; element = element.Next() { 
        fmt.Println(element.Value.(WordCount)) 
    } 
}
output==>
{bob 107}
{jack 97}
{herry 92}
{kate 87}
{james 81}
{jason 67}
</pre>
###Golang之面对对象实现
封装性：
<pre>
package main

import "fmt"

type data struct {
	val int
}
func (p_data* data)set(num int) {

	p_data.val = num
}
func (p_data* data)show() {

	fmt.Println(p_data.val)
}
func main() {
	p_data := &data{4}
	p_data.set(5)
	p_data.show()
}
output==>
5
</pre>
继承特性:
<pre>
package main  
import "fmt"  
type parent struct {
    val int  
}  
type child struct {  
    parent  
    num int  
}  
func main() {  
  
    var c child  
  
    c = child{parent{1}, 2}  
    fmt.Println(c.num)  
    fmt.Println(c.val)  
}  
output==>
1
2
</pre>
多态特性:
<pre>
package main  
import "fmt"  
type act interface {  
 	 write()  
}  
type xiaoming struct {  
}  
type xiaofang struct {  
}  
func (xm *xiaoming) write() {  
    fmt.Println("xiaoming write")  
}  
func (xf *xiaofang) write() {    
    fmt.Println("xiaofang write")  
}  
func main() {  
    var w act;  
    xm := xiaoming{}  
    xf := xiaofang{}  
     
    w = &xm  
    w.write()  
  
    w = &xf  
    w.write()  
}  
output==>
xiaoming write
xiaofang write
</pre>
###多个channel
<pre>
package main  
  
import (
	"fmt"  
	"os"
)
import "time"  
  
func fibonacci(c, quit chan int) { 
    x, y := 1, 1  

    for {  
            select {  

                    case c <- x:  
                            x, y = y, x+y  

                    case <- quit:  
                            fmt.Println("quit")  
                            os.Exit(0)  
            }  
    }  
}  
func show(c, quit chan int) {  
    for i := 0; i < 10; i ++ {  
           fmt.Println(<- c)  
    }  
    quit <- 0  
}  
func main() {  
    data := make(chan int)  
    leave := make(chan int)  
    go show(data, leave)  
    go fibonacci(data, leave)  
    for {  
            time.Sleep(100)  
    }  
} 
output==>
1
1
2
3
5
8
13
21
34
55
quit
</pre>
###输出格式
<pre>
package main

import (
	"math"
	"fmt"
)
func main(){
	fmt.Printf("二进制：%b\n",255)
	fmt.Printf("八进制：%o\n",255)
	fmt.Printf("十六进制：%X\n",255)
	fmt.Printf("十进制：%d\n",255)
	fmt.Printf("浮点数：%f\n",math.Pi)
	fmt.Printf("字符串：%s\n","hi")
}
output==>
二进制：11111111
八进制：377
十六进制：FF
十进制：255
浮点数：3.141593
字符串：hi
</pre>
切片
<pre>
package main

import (
	"fmt"
)
func main(){
	a := [5]int{1,2,3,4,5}
	b := a[2:4]
	fmt.Println(b)
	b = a[:4]
	fmt.Println(b)
}
output==>
[3 4]
[1 2 3 4]
</pre>
<pre>
package main

import (
	"fmt"
)
func main(){
	i := 1
	//精简的for语句
	for i<5{
		fmt.Println(i)
		i++
	}
	/*这一种也是同样效果
	for {
		if i>=5{
			break
		}
		fmt.Println(i)
		i++
	}*/
}
output==>
1
2
3
4
</pre>


string ==>int形式的map：
<pre>
package main
import (
	"fmt"
)
func main(){
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println(m)
	fmt.Println("length is ",len(m))
	v :=m["two"]
	fmt.Println(v)
	delete(m,"two")
	fmt.Println(m)
	m1 :=map[string]int{"one":1,"two":2,"three":3}
	fmt.Println(m1)
	for key,val := range m1 {
		fmt.Println(key,"=>",val)
	}
}
output==>
map[one:1 two:2 three:3]
length is  3
2
map[one:1 three:3]
map[one:1 two:2 three:3]
one => 1
two => 2
three => 3
</pre>
string ==> []string形式的map
<pre>
package main
import (
	"fmt"
)
func main(){
	var abs map[string][]string
	abs =make(map[string][]string)
	abs["two"] = []string{"Two","Three"}
	abs["three"] =[]string{"Ok","No"}
	fmt.Println(abs["three"])
}
output==>
[Ok No]
</pre>
map的range循环输出是随机的，所以不要以为会按照map中元素的顺序排列输出，如下：
<pre>
package main

import (
	_"sort"
	"fmt"
)
func main(){
	blog :=map[string]int{
	  "0unix": 0,
      "1python": 1,
      "2go": 2,
      "3javascript": 3,
      "4testing": 4,
      "5philosophy": 5,
      "6startups": 6,
      "7productivity": 7,
      "8hn": 8,
      "9reddit": 9,
      "10C++": 10,
	}
	 for key, views := range blog {
      fmt.Println("There are", views, "views for", key)
  }
}
output==>
There are 1 views for 1python
There are 7 views for 7productivity
There are 9 views for 9reddit
There are 8 views for 8hn
There are 0 views for 0unix
There are 2 views for 2go
There are 3 views for 3javascript
There are 4 views for 4testing
There are 5 views for 5philosophy
There are 6 views for 6startups
</pre>
那么如何获得依次排列的map呢：
<pre>
package main

import (
	"sort"
	"fmt"
)
func main(){
	blog :=map[string]int{
	  "0unix": 0,
      "1python": 1,
      "2go": 2,
      "3javascript": 3,
      "4testing": 4,
      "5philosophy": 5,
      "6startups": 6,
      "7productivity": 7,
      "8hn": 8,
      "9reddit": 9,
	}
	var keys []string
	for k := range blog {
		keys =append(keys,k)
	}
	sort.Strings(keys)
	 for _,k:= range keys {
      fmt.Println("There are", k, " views for", blog[k])
  }
}
output==>
There are 0unix  views for 0
There are 1python  views for 1
There are 2go  views for 2
There are 3javascript  views for 3
There are 4testing  views for 4
There are 5philosophy  views for 5
There are 6startups  views for 6
There are 7productivity  views for 7
There are 8hn  views for 8
There are 9reddit  views for 9
</pre>



<pre>
package main
import (
	"fmt"
)
func main(){	
	//&取出某变量对应的内存地址，*通过内存地址找到变量的值
	r := 5
	fmt.Println(&r,*&r)
	
}
output==>
0xc0820022d0 5
</pre>
new关键字分配内存的内建函数，但不同于其他语言中同名的new所作的工作，它只是将内存清零，而不是初始化内存。<br>
new(T)为一个类型为T的新项目分配了值为零的存储空间并返回其地址，也就是一个类型为*T的值。用Go的术语来说，就是它返回了一个指向新分配的类型为T的零值的指针。
<pre>
package main

import (
	"fmt"
)
func main(){	
	var p *[]int =new([]int)//为切片结构分配内存
	*p =make([]int,10,10)
	(*p)[2] = 5
	fmt.Println((*p)[2])
	
}
output==>
5
</pre>
函数多返回值，一般情况下，一个是需要返回的值，另一个是错误信息：
<pre>
package main
import (
	"fmt"
)
func multireturn(key string)(int, bool){
	m :=map[string]int{"one":1,"two":2,"three":3}
	var err bool
	var val int
	val,err =m[key]
	return val ,err
}
func main(){
	v,e :=multireturn("two")
	fmt.Println(v,e)
}
output==>
2 true
</pre>
####不定参数
<pre>
package main
import (
	"fmt"
)
func sum(nums ...int){
	//不定参数
	fmt.Println(nums,"")//输出如[1, 2, 3]之类的数组
	total := 0
	for _,num :=range nums {
		total += num
	}
	fmt.Println(total)
}
func main(){
	sum(1,2,4,5,6,6,6)
}
output==>
[1 2 4 5 6 6 6] 
total is : 30
</pre>
####函数闭包
<pre>
package main
//一个斐波拉契数
import (
	"fmt"
)
func next()func() int{
	i,j := 1,1
	return func() int{
		var tmp = i+j
		i,j = j,tmp
		return tmp
	}
}
func main(){
	nextfunc :=next()
	for i:=0;i<5;i++{
		fmt.Println(nextfunc())
	}
}
output==>
2
3
5
8
13
</pre>
####函数递归
<pre>
package main
//递归，适用于有重复使用的方法或者参数
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
####结构体方法
将一个方法作为结构体的属性（在js或者java面对对象中的属性概念）
<pre>
package main

import (
	"fmt"
)
type rect struct {
	width,height int
}
func (r *rect)area() int{
	return r.width * r.height
}
func (r *rect)perimeter()int{
	return 2*(r.width + r.height)
}
func main(){
	r :=rect{width:2,height:4}
	fmt.Println("面积：",r.area())
	fmt.Println("周长：",r.perimeter())
	rp := &r
    fmt.Println("面积: ", rp.area())
    fmt.Println("周长: ", rp.perimeter())
}
output==>
面积： 8
周长： 12
面积:  8
周长:  12
</pre>
####错误处理-Error接口
<pre>
package main

import (
	"errors"
	"fmt"
)
//自定义错误结构
type myerror struct {
	arg int
	errMsg string
}
//实现Error接口
func (e *myerror) Error() string{
	return fmt.Sprintf("%d - %s",e.arg,e.errMsg)
}
func error_test(arg int)(int,error){
	if arg > 0 {
		return -1,errors.New("Bad Arguments - negtive")
	}else if arg >256{
		return -1,&myerror{arg,"Bad Argments - too large"}
	}
	return arg*arg,nil
}
func main(){
	for _,i :=range []int{-1,4,100}{
		if r,e:=error_test(i);e != nil {
			fmt.Println("faild:",e)
		}else{
			fmt.Println("success:",r)
		}
	}
}
output==>
success: 1
faild: Bad Arguments - negtive
faild: Bad Arguments - negtive
</pre>
####错误处理 – Panic/Recover
对于不可恢复的错误，Go提供了一个内建的panic函数，它将创建一个运行时错误并使程序停止（相当暴力）。该函数接收一个任意类型（往往是字符串）作为程序死亡时要打印的东西。当编译器在函数的结尾处检查到一个panic时，就会停止进行常规的return语句检查。
<pre>
package main

import (
	"fmt"
)
func g(i int) {
    if i>1 {
        fmt.Println("Panic!")
        panic(fmt.Sprintf("%v", i))
    }
 
}
 
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
 
    for i := 0; i < 4; i++ {
        fmt.Println("Calling g with ", i)
        g(i)
        fmt.Println("Returned normally from g.")
     }
}
 
func main() {
    f()
    fmt.Println("Returned normally from f.")
}
output==>
Calling g with  0
Returned normally from g.
Calling g with  1
Returned normally from g.
Calling g with  2
Panic!
Recovered in f 2
Returned normally from f.
</pre>
####goroutine
下面的示例包括时间处理，随机数处理，还有goroutine:
<pre>
package main

import (
	"os"
	_"os"
	"math/rand"
	"fmt"
	"time"
)
func routine(name string,delay time.Duration){
	t0 :=time.Now()
	fmt.Println(name,"start at",t0)
	time.Sleep(delay)
	t1 :=time.Now()
	fmt.Println(name,"end at",t1)
	fmt.Println(name,"lasted ",t1.Sub(t0))
}
func main(){
	rand.Seed(time.Now().Unix())
	var name string
	for i:=0;i<3;i++{
		name = fmt.Sprintf("go %d",i) //生成ID
		go routine(name,time.Duration(rand.Intn(5)) * time.Second)
	}
	A:
	var input string
	fmt.Scanln(&input)
	ex := &input
	if *ex =="exit"{
		os.Exit(0)
	}else{
		goto A
	}
}
output==>
go 0 start at 2016-03-31 21:41:13.1557936 +0800 +0800
go 1 start at 2016-03-31 21:41:13.1567936 +0800 +0800
go 2 start at 2016-03-31 21:41:13.1567936 +0800 +0800
go 2 end at 2016-03-31 21:41:13.1567936 +0800 +0800
go 2 lasted  0
go 1 end at 2016-03-31 21:41:16.1569652 +0800 +0800
go 1 lasted  3.0001716s
go 0 end at 2016-03-31 21:41:17.1570224 +0800 +0800
go 0 lasted  4.0012288s
vf
ff
ff
f
exit
</pre>
####goroutine的并发安全性|锁Mutex|mutex
<pre>
package main

import (
	"sync"
	"os"
	"runtime"
	"fmt"
	"math/rand"
	"time"
)
var total_tickets int32 = 10
var mutex = &sync.Mutex{}
func sell_tickets(i int){
	for total_tickets>0 {
		mutex.Lock()
		if total_tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(5))*time.Millisecond)
			total_tickets--
			fmt.Println("id:",i," tickets:",total_tickets)
		}else{
			break
		}
		mutex.Unlock()
	}
}
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	rand.Seed(time.Now().Unix())
	for i:=0;i<5;i++{
		go sell_tickets(i)
	}
	var input string
	fmt.Scanln(&input)
	fmt.Println(total_tickets,"done")
	os.Exit(0)
}
output==>
4
id: 1  tickets: 9
id: 1  tickets: 8
id: 1  tickets: 7
id: 1  tickets: 6
id: 1  tickets: 5
id: 1  tickets: 4
id: 1  tickets: 3
id: 1  tickets: 2
id: 1  tickets: 1
id: 1  tickets: 0
g
0 done
</pre>
####原子操作
10个goroutine，每个会对cnt变量累加20次，所以，最后的cnt应该是200。如果没有atomic的原子操作，那么cnt将有可能得到一个小于200的数。
<pre>
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)
func main(){
	var cnt uint32 = 0
	for i:=0;i<10;i++{
		go func(){
			for i:=0;i<10;i++{
				time.Sleep(time.Millisecond)
				atomic.AddUint32(&cnt,1)
			}
		}()
	}
	time.Sleep(time.Second) //等一秒钟让goroutine完成
	cntFinal := atomic.LoadUint32(&cnt)//取出数据
	fmt.Println("cnt:",cntFinal)
}
output==>
cnt: 100
</pre>
####Channel 信道
Channal就是用来通信的，就像Unix下的管道一样，在Go中是这样使用Channel的。
<pre>
package main

import (
	"fmt"
)
func main(){
	channel :=make(chan string)
	go func(){
		channel <- "hello"
	}()
	msg := <- channel
	fmt.Println(msg)
}
output==>
hello
</pre>
####Channel的阻塞
<pre>
package main

import (
	"time"
	"fmt"
)
func main(){
	channel :=make(chan string)
	go func(){
		channel <- "hello"
		fmt.Println("write \"hello\"done!")
		channel <- "world" //Reader在Sleep，这里在阻塞
		fmt.Println("write\"world\"done!")
		fmt.Println("write go sleep ...")
		time.Sleep(3*time.Second)
		channel <- "channel"
		fmt.Println("write \"channel\"done!")
	}()
	time.Sleep(2*time.Second)
	fmt.Println("Reader wake up ...")
	
	msg :=<-channel
	fmt.Println("Reader:",msg)
	
	msg =<-channel
	fmt.Println("Reader:",msg)
	
	msg =<-channel //Writer在Sleep，这里在阻塞
	fmt.Println("Reader:",msg)
}
output==>
Reader wake up ...
Reader: hello
write "hello"done!
write"world"done!
write go sleep ...
Reader: world
write "channel"done!
Reader: channel
</pre>
####利用channel阻塞实现golang版线程池
<pre>
package main
/*
任务：go语言实现一个线程池，主要功能是：添加total个任务到线程池中，
线程池开启number个线程，每个线程从任务队列中取出一个任务执行，
执行完成后取下一个任务，全部执行完成后回调一个函数。
思路：将任务放到channel里，每个线程不停的从channel中取出任务执行，
并把执行结果写入另一个channel，当得到total个结果后，回调函数。
*/
import (
	"time"
	"io"
	"strings"
	"os"
	"net/http"
	"fmt"
)

type GoroutinePool struct {
      Queue  chan func() error
      Number int
      Total  int
      result         chan error
      finishCallback func()
 }
 
 // 初始化
func (self *GoroutinePool) Init(number int, total int) {
     self.Queue = make(chan func() error, total)
     self.Number = number
     self.Total = total
     self.result = make(chan error, total)
 }
 
 // 开门接客
 func (self *GoroutinePool) Start() {
     // 开启Number个goroutine
     for i := 0; i < self.Number; i++ {
        go func() {
             for {
                task, ok := <-self.Queue
                if !ok {
                     break              
			    }
                err := task()
                self.result <- err
             }         
		}()
     }
 
     // 获得每个work的执行结果
     for j := 0; j < self.Total; j++ {
         res, ok := <-self.result
        if !ok {
             break
        }
        if res != nil {
           fmt.Println(res)   
        }
}    // 所有任务都执行完成，回调函数
    if self.finishCallback != nil {     
	   self.finishCallback()
    }
}
// 关门送客
func (self *GoroutinePool) Stop() {
    close(self.Queue)
    close(self.result)
}

// 添加任务
func (self *GoroutinePool) AddTask(task func() error) {
    self.Queue <- task
}
// 设置结束回调
func (self *GoroutinePool) SetFinishCallback(callback func()) {
    self.finishCallback = callback
 }

func Download_test() {
      urls := []string{
          "http://dlsw.baidu.com/sw-search-sp/soft/44/17448/Baidusd_Setup_4.2.0.7666.1436769697.exe",
          "http://dlsw.baidu.com/sw-search-sp/soft/3a/12350/QQ_V7.4.15197.0_setup.1436951158.exe",
          "http://dlsw.baidu.com/sw-search-sp/soft/9d/14744/ChromeStandalone_V43.0.2357.134_Setup.1436927123.exe",
		}
     pool := new(GoroutinePool)
     pool.Init(3, len(urls))
 
     for i := range urls {
         url := urls[i]
         pool.AddTask(func() error {
             return download(url)
         })
     }
	 isFinish := false
	pool.SetFinishCallback(func() {
		func(isFinish *bool) {
			 *isFinish = true
		}(&isFinish)
	})
	pool.Start()
	 for !isFinish {
         time.Sleep(time.Millisecond * 100)
     }
	 pool.Stop()
     fmt.Println("所有操作已完成！")
}
func download(url string) error {
    fmt.Println("开始下载... ", url)

     sp := strings.Split(url, "/")
     filename := sp[len(sp)-1]
 
     file, err := os.Create("./aa/" + filename)
     if err != nil {
         return err
     }
      res, err := http.Get(url)
     if err != nil {
         return err
     }
 
     length, err := io.Copy(file, res.Body)
     if err != nil {
         return err
     }
 
     fmt.Println("## 下载完成！ ", url, " 文件长度：", length)
     return nil
}
func main(){
	Download_test()
}
output==>
开始下载...  http://dlsw.baidu.com/sw-search-sp/soft/44/17448/Baidusd_Setup_4.2.0.7666.1436769697.exe
开始下载...  http://dlsw.baidu.com/sw-search-sp/soft/3a/12350/QQ_V7.4.15197.0_setup.1436951158.exe
开始下载...  http://dlsw.baidu.com/sw-search-sp/soft/9d/14744/ChromeStandalone_V43.0.2357.134_Setup.1436927123.exe
## 下载完成！  http://dlsw.baidu.com/sw-search-sp/soft/44/17448/Baidusd_Setup_4.2.0.7666.1436769697.exe  文件长度： 28500944
[然后所有文件下载在./aa目录下]
</pre>
####多个Channel的select
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
		time.Sleep(time.Second*1)
		c1 <- "hello"
	}()
	go func(){
		time.Sleep(time.Second*1)
		c2<- "world"
	}()
	for i:=0;i<2;i++{
		select {
			case msg1 :=<- c1:
			fmt.Println("received:",msg1)
			case msg2 :=<- c2:
			fmt.Println("received:",msg2)
		}
	}
}
output==>
received: hello
received: world
</pre>
####Channel的关闭
<pre>
package main
 
import "fmt"
import "time"
import "math/rand"
 
func main() {
 
    channel := make(chan string)
    rand.Seed(time.Now().Unix())
 
    //向channel发送随机个数的message
    go func () {
        cnt := rand.Intn(10)
        fmt.Println("message cnt :", cnt)
        for i:=0; i<cnt; i++{
            channel <- fmt.Sprintf("message-%2d", i)
        }
        close(channel) //关闭Channel
    }()
 
    var more bool = true
    var msg string
    for more {
        select{
        //channel会返回两个值，一个是内容，一个是还有没有内容
        case msg, more = <- channel:
            if more {
                fmt.Println(msg)
            }else{
                fmt.Println("channel closed!")
            }
        }
    }
}
output==>
//结果的一种可能是：
message cnt : 3
message- 0
message- 1
message- 2
channel closed!
</pre>
####定时器
Go语言中可以使用time.NewTimer或time.NewTicker来设置一个定时器，这个定时器会绑定在你的当前channel中，通过channel的阻塞通知机器来通知你的程序。
<pre>
package main

import (
	"fmt"
	"time"
)
func main(){
	timer:=time.NewTimer(2*time.Second)
	<- timer.C
	fmt.Println("timer expired!")
}
output==>
timer expired!
</pre>
<pre>
package main
import (
 "fmt"
 "time"
)
func testTimer1() {
 go func() {
  fmt.Println("test timer1")
 }()
}
func testTimer2() {
 go func() {
  fmt.Println("test timer2")
 }()
}
func timer1() {
 timer1 := time.NewTicker(1 * time.Second)
 for {
  select {
  case <-timer1.C:
   testTimer1()
  }
 }
}
func timer2() {
 timer2 := time.NewTicker(2 * time.Second)
 for {
  select {
  case <-timer2.C:
   testTimer2()
  }
 }
}
func main() {
 go timer1()
 timer2()
}
output==>
test timer1
test timer1
test timer2
test timer1
test timer1
test timer2
test timer1
test timer1
test timer2
test timer1
...
</pre>
Ticker打点器
<pre>
package main
//需要持续通知，通过ticker打点器实现
import (
	"fmt"
	"time"
)
func main(){
	ticker :=time.NewTicker(time.Second)
	for t:=range ticker.C {
		fmt.Println("Tick at",t)
	}
}
output==>
Tick at 2016-04-01 22:27:27.4815005 +0800 +0800
Tick at 2016-04-01 22:27:28.4815577 +0800 +0800
Tick at 2016-04-01 22:27:29.4816149 +0800 +0800
Tick at 2016-04-01 22:27:30.4816721 +0800 +0800
Tick at 2016-04-01 22:27:31.4817293 +0800 +0800
...
</pre>
上面的程序会进入一个死循环中，我们为了实现可控，可以将它放入一个goroutine中：
<pre>
package main
import "time"
import "fmt"
func main() {
 
    ticker := time.NewTicker(time.Second)
 
    go func () {
        for t := range ticker.C {
            fmt.Println(t)
        }
    }()
 
    //设置一个timer，10钞后停掉ticker
    timer := time.NewTimer(10*time.Second)
    <- timer.C
 
    ticker.Stop()
    fmt.Println("timer expired!")
}
output==>
2016-04-01 22:37:01.6023383 +0800 +0800
2016-04-01 22:37:02.6023955 +0800 +0800
2016-04-01 22:37:03.6024527 +0800 +0800
2016-04-01 22:37:04.6025099 +0800 +0800
2016-04-01 22:37:05.6025671 +0800 +0800
2016-04-01 22:37:06.6026243 +0800 +0800
2016-04-01 22:37:07.6026815 +0800 +0800
2016-04-01 22:37:08.6027387 +0800 +0800
2016-04-01 22:37:09.6027959 +0800 +0800
2016-04-01 22:37:10.6028531 +0800 +0800
timer expired!
</pre>
环境变量
<pre>
package main
//遍历环境变量
import (
	"fmt"
	"strings"
	"os"
)
func main(){
	os.Setenv("Jason","dj")//设置环境变量
	for _,env :=range os.Environ(){
		e := strings.Split(env,"=")
		fmt.Println(e[0],"=",e[1])
	}
}
output==>
GOEXE = .exe
Jason = dj
PROCESSOR_ARCHITECTURE = AMD64
COMPUTERNAME = JASON
...
</pre>
<pre>
package main
import "flag"
import "fmt"
 
func main() {
 
    //第一个参数是“参数名”，第二个是“默认值”，第三个是“说明”。返回的是指针
    host := flag.String("host", "baidu.com", "a host name ")
    port := flag.Int("port", 80, "a port number")
    debug := flag.Bool("d", false, "enable/disable debug mode")
 
    //正式开始Parse命令行参数
    flag.Parse()
 
    fmt.Println("host:", *host)
    fmt.Println("port:", *port)
    fmt.Println("debug:", *debug)
}
output==>
host: baidu.com
port: 80
debug: false
用法是：
#指定了参数名后的情况
$ go run flagtest.go -host=localhost -port=22 -d
host: localhost
port: 22
debug: true
</pre>
HTTP Server
<pre>
package main

import (
	"io/ioutil"
	"path/filepath"
	"fmt"
	"net/http"
)
const http_root = "/template/"
func main(){
	http.HandleFunc("/",rootHandler)
	http.HandleFunc("/view/",viewHandler)
	http.HandleFunc("/html/",htmlHandler)
	http.ListenAndServe(":8080",nil)
}
func rootHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"rootHandler:%s\n",r.URL.Path)
	fmt.Fprintf(w,"URL:%S\n",r.URL)
	fmt.Fprintf(w,"Method:%s\n",r.Method)
	fmt.Fprintf(w,"RequestURI:%s\n",r.RequestURI)
	fmt.Fprintf(w,"Proto:%s\n",r.Proto)
	fmt.Fprintf(w,"HOST:%s\n",r.Host)
}
//特别的URL处理
func viewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "viewHandler: %s", r.URL.Path)
}
 
//一个静态网页的服务示例。（在http_root的html目录下）
func htmlHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("htmlHandler: %s\n", r.URL.Path)
     
    filename := http_root + r.URL.Path
    fileext := filepath.Ext(filename)
 
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf("   404 Not Found!\n")
        w.WriteHeader(http.StatusNotFound)
        return
    }
     
    var contype string
    switch fileext {
        case ".html", "htm":
            contype = "text/html"
        case ".css":
            contype = "text/css"
        case ".js":
            contype = "application/javascript"
        case ".png":
            contype = "image/png"
        case ".jpg", ".jpeg":
            contype = "image/jpeg"
        case ".gif":
            contype = "image/gif"
        default: 
            contype = "text/plain"
    }
    fmt.Printf("ext %s, ct = %s\n", fileext, contype)
     
    w.Header().Set("Content-Type", contype)
    fmt.Fprintf(w, "%s", content)   
}

</pre>
###函数式编程
当我们说起函数式编程来说，我们会看到如下函数式编程的长相：

- （三大特性）：
	1. immutable data 不可变数据：像Clojure一样，默认上变量是不可变的，如果你要改变变量，你需要把变量copy出去修改。这样一来，可以让你的程序少很多Bug。因为，程序中的状态不好维护，在并发的时候更不好维护。（你可以试想一下如果你的程序有个复杂的状态，当以后别人改你代码的时候，是很容易出bug的，在并行中这样的问题就更多了）
	2. first class functions：这个技术可以让你的函数就像变量一样来使用。也就是说，你的函数可以像变量一样被创建，修改，并当成变量一样传递，返回或是在函数中嵌套函数。这个有点像Javascript的Prototype（参看Javascript的面向对象编程）
	3. 尾递归优化：我们知道递归的害处，那就是如果递归很深的话，stack受不了，并会导致性能大幅度下降。所以，我们使用尾递归优化技术——每次递归时都会重用stack，这样一来能够提升性能，当然，这需要语言或编译器的支持。Python就不支持。

- 函数式编程的几个技术
	1. map & reduce ：这个技术不用多说了，函数式编程最常见的技术就是对一个集合做Map和Reduce操作。这比起过程式的语言来说，在代码上要更容易阅读。（传统过程式的语言需要使用for/while循环，然后在各种变量中把数据倒过来倒过去的）这个很像C++中的STL中的foreach，find_if，count_if之流的函数的玩法。
	2. pipeline：这个技术的意思是，把函数实例成一个一个的action，然后，把一组action放到一个数组或是列表中，然后把数据传给这个action list，数据就像一个pipeline一样顺序地被各个函数所操作，最终得到我们想要的结果。
	3. recursing 递归 ：递归最大的好处就简化代码，他可以把一个复杂的问题用很简单的代码描述出来。注意：递归的精髓是描述问题，而这正是函数式编程的精髓。
	4. currying：把一个函数的多个参数分解成多个函数， 然后把函数多层封装起来，每层函数都返回一个函数去接收下一个参数这样，可以简化函数的多个参数。在C++中，这个很像STL中的bind_1st或是bind2nd。
	5. higher order function 高阶函数：所谓高阶函数就是函数当参数，把传入的函数做一个封装，然后返回这个封装函数。现象上就是函数传进传出，就像面向对象对象满天飞一样。

- 还有函数式的一些好处
	1. parallelization 并行：所谓并行的意思就是在并行环境下，各个线程之间不需要同步或互斥。
	2. lazy evaluation 惰性求值：这个需要编译器的支持。表达式不在它被绑定到变量之后就立即求值，而是在该值被取用的时候求值，也就是说，语句如x:=expression; (把一个表达式的结果赋值给一个变量)明显的调用这个表达式被计算并把结果放置到 x 中，但是先不管实际在 x 中的是什么，直到通过后面的表达式中到 x 的引用而有了对它的值的需求的时候，而后面表达式自身的求值也可以被延迟，最终为了生成让外界看到的某个符号而计算这个快速增长的依赖树。
	3. determinism 确定性：所谓确定性的意思就是像数学那样 f(x) = y ，这个函数无论在什么场景下，都会得到同样的结果，这个我们称之为函数的确定性。而不是像程序中的很多函数那样，同一个参数，却会在不同的场景下计算出不同的结果。所谓不同的场景的意思就是我们的函数会根据一些运行中的状态信息的不同而发生变化。

- 通过实例走近函数式编程：
首先是一个非函数式：
<pre>
var cnt int
func increment(){
    cnt++;
}
</pre>
接着是一个函数式：
<pre>
func increment(cnt int){
    return cnt+1;
}
</pre>
通过上面的比较可以得出：<b>不依赖于外部的数据，而且也不改变外部数据的值，而是返回一个新的值给你</b>。<br>
还有一个特点：<b>把函数当成变量来用，关注于描述问题而不是怎么实现，这样可以让代码更易读</b>。
###带有表单处理的web服务器
main.go
<pre>
package main

import (
	"html/template"
	"fmt"
	"net/http"
)
func hi(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello wolrd")
}
func login(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
		t,_:=template.ParseFiles("login.gtpl")
		t.Execute(w,nil)
	}else{
		r.ParseForm()
		username := r.Form["username"]
		password := r.Form["password"]
		user := ""
		pass := ""
		for _,s :=range username{
			user +=s
		}
		for _,s :=range password{
			pass +=s
		}
		fmt.Fprintf(w,"username:",user)
		fmt.Fprintf(w,"password:",pass)
		//fmt.Println("username:",r.Form["username"])
		//fmt.Println("password:",r.Form["password"])
	}
}
func main(){
	http.HandleFunc("/",hi)
	http.HandleFunc("/login",login)
	http.ListenAndServe(":9099",nil)
}
</pre>
login.gtpl
<pre>
<html>  
<head>  
<title> </title>  
</head>  
  
<body>  
<form action="http://127.0.0.1:9099/login" method="post">  
        user: <input type="text" name ="username">  
        pass: <input type="password" name="password">  
        <input type="submit" value="login">  
</form>  
</body>  
</html> 
</pre>
Golang中的字符串
<pre>
package main

import (
	"fmt"
)
func main(){
	str := "高贵不发"
	for i,s:=range str {
		fmt.Println(i,"unicode(",s,") string=",string(s))
	}
}
output==>
0 unicode( 39640 ) string= 高
3 unicode( 36149 ) string= 贵
6 unicode( 19981 ) string= 不
9 unicode( 21457 ) string= 发
</pre>
###strings包
Golang字符串处理函数
<pre>
package main

import (
	"strings"
	"fmt"
)
func main(){
	str := "hello,董军"
	sli :=[]string{"H","I"}
	fmt.Println(strings.Count(str,"l"))//统计字符串指定字符的个数
	fmt.Println(strings.Contains(str,"he"))//存在true,反之false
	fmt.Println(strings.Contains("", "")) //注意    
	fmt.Println(strings.ContainsAny("", ""))// false
	fmt.Println(strings.ContainsAny(str,"e&o")) //true检测是否同时存在多个元素
	fmt.Println(strings.ContainsRune(str,'董'))//第二个是字符
	fmt.Println(strings.EqualFold("GO","go"))//大小写忽略的情况下判断两个字符串是否相等 
	fmt.Println(strings.HasPrefix(str,"he"))//判断字符串前缀是否是指定字符
	fmt.Println(strings.HasSuffix(str,"军"))//判断字符串后缀是否是指定字符
	fmt.Println(strings.Index(str,"军"))//指定字符第一次出现的位置|中文占2个字符
	fmt.Println(strings.LastIndex(str, "l")) //指定字符最后一次出现的位置
	fmt.Println(strings.IndexAny(str, "军"))//指定字符的位置|中文占2个字符
	fmt.Println(strings.IndexRune(str, '军'))//指定字符的位置|中文占2个字符
	fmt.Println(strings.Join(sli,""))//将数组或切片按照指定连接符连接成字符串
	fmt.Println("ba" + strings.Repeat("na", 2))//重复指定次数的字符串
	fmt.Println(strings.Replace(str,"l","L",1))//替换指定字符指定个数|heLlo,董军
	fmt.Println(strings.Replace(str,"l","L",2))//heLLo,董军
	fmt.Println(strings.Split(str,",")) //[hello 董军]
	fmt.Println(strings.ToLower(str))//hello,董军
	fmt.Println(strings.ToUpper(str))//HELLO,董军HELLO,董军
	fmt.Println(strings.Trim("!!!Achtung!!!", "!"))//删除左右指定字符|Achtung|同理有trimLeft,trimRight
	fmt.Println(strings.TrimSpace(" \t\n Achtung \n\t\r\n"))//删除空格Achtung
}
output==>
2
true
true
false
true
true
true
true
true
9
3
9
9
HI
banana
heLlo,董军
heLLo,董军
[hello 董军]
hello,董军
HELLO,董军
Achtung
Achtung
</pre>
###回调函数
####什么是是回调
其实回调就是一种利用函数指针进行函数调用的过程.  

为什么要用回调呢?比如我要写一个子模块给你用,   来接收远程socket发来的命令.当我接收到命令后,   需要调用你的主模块的函数,   来进行相应的处理.但是我不知道你要用哪个函数来处理这个命令,     我也不知道你的主模块是什么.cpp或者.h,   或者说,   我根本不用关心你在主模块里怎么处理它,   也不应该关心用什么函数处理它......   怎么办?使用回调!

对回调函数的一种理解是：<br>
使用回调函数实际上就是在调用某个函数（通常是API函数）时，将自己的一个函数（这个函数为回调函数）的地址作为参数传递给那个函数。而那个函数在需要的时候，利用传递的地址调用回调函数，这时你可以利用这个机会在回调函数中处理消息或完成一定的操作。

还有人这样理解回调：<br>
回调函数，就是由你自己写的。你需要调用另外一个函数，而这个函数的其中一个参数，就是你的这个回调函数名。这样，系统在必要的时候，就会调用你写的回调函数，这样你就可以在回调函数里完成你要做的事。

Golang回调函数示例：
<pre>
package main

import (
	"fmt"
)
type testStruct struct {}
func (object *testStruct)test(msg string){
	fmt.Println(msg)
}
type callBack func(msg string)
func calBackTest(backfunc callBack){
	backfunc("jason")
}
func main(){
	object := new(testStruct)
	calBackTest(object.test)
}
output==>
jason
</pre>
另一个例子：
<pre>
package main

import (
	"fmt"
)
func CallBack(f func(int) int){
	fmt.Println(f)
	f(32)
}
func main(){
	CallBack(func(m int)int{
		fmt.Println(m)
		return m
	})
}
output==>
0x401110
32
</pre>
####Golang处理支付宝的回调
<pre>
package main

import (
    "crypto"
    "crypto/rsa"
    "crypto/sha1"
    "crypto/x509"
    "encoding/base64"
    "encoding/hex"
    "encoding/pem"
    "fmt"
    "io"
    "net/url"
    "sort"
)

const (
    //支付宝公钥
    ALIPAY_PUBLIC_KEY = `-----BEGIN PUBLIC KEY-----  
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCnxj/9qwVfgoUh/y2W89L6BkRA
FljhNhgPdyPuBV64bfQNN1PjbCzkIM6qRdKBoLPXmKKMiFYnkd6rAoprih3/PrQE
B/VsW8OoM8fxn67UDYuyBTqA23MML9q1+ilIZwBC2AQ2UBVOrFXfFl75p6/B5Ksi
NG9zpgmLCUYuLkxpLQIDAQAB 
-----END PUBLIC KEY-----
`
)
func main() {
    //这是我从支付宝callback的http body部分读取出来的一段参数列表。
    paramerStr := `discount=0.00&payment_type=1&subject=%E7%BC%B4%E7%BA%B3%E4%BF%9D%E8%AF%81%E9%87%91&trade_no=2015122121001004460085270336&buyer_email=xxaqch%40163.com&gmt_create=2015-12-21+13%3A13%3A28&notify_type=trade_status_sync&quantity=1&out_trade_no=a378c684be7a4f99be1bf3b56e6d38fd&seller_id=2088121529348920&notify_time=2015-12-21+13%3A17%3A45&body=%E7%BC%B4%E7%BA%B3%E4%BF%9D%E8%AF%81%E9%87%91&trade_status=TRADE_SUCCESS&is_total_fee_adjust=N&total_fee=0.01&gmt_payment=2015-12-21+13%3A13%3A28&seller_email=172886370%40qq.com&price=0.01&buyer_id=2088002578894463&notify_id=5104b719303162e2b79d577aeaa5494jjs&use_coupon=N&sign_type=RSA&sign=YeshUpQO1GsR4KxQtAlPzdlqKUMlTfEunQmwmNI%2BMJ1T2qzd9WuA6bkoHYMM8BpHxtp5mnFM3rXlfgETVsQcNIiqwCCn1401J%2FubOkLi2O%2Fmta2KLxUcmssQ0OnkFIMjjNQuU9N3eIC1Z6SzDkocK092w%2Ff3un4bxkIfILgdRr0%3D`

    //调用url.ParseQuery来获取到参数列表，url.ParseQuery还会自动做url safe decode
    values_m, _err := url.ParseQuery(string(paramerStr))
    if _err != nil {
        fmt.Println("error parse parameter, reason:", _err)
        return
    }
    var m map[string]interface{}
    m = make(map[string]interface{}, 0)

    for k, v := range values_m {
        if k == "sign" || k == "sign_type" { //不要'sign'和'sign_type'
            continue
        }
        m[k] = v[0]
    }

    sign := values_m["sign"][0]
    fmt.Println("Parsed Sign:", []byte(sign))

    //获取要进行计算哈希的sign string
    strPreSign, _err := genAlipaySignString(m)
    if _err != nil {
        fmt.Println("error get sign string, reason:", _err)
        return
    }

    fmt.Println("Presign string:", strPreSign)

    //进行rsa verify
    pass, _err := RSAVerify([]byte(strPreSign), []byte(sign))

    if pass {
        fmt.Println("verify sig pass.")
    } else {
        fmt.Println("verify sig not pass. error:", _err)
    }
}
/***************************************************************
*函数目的：获得从参数列表拼接而成的待签名字符串
*mapBody：是我们从HTTP request body parse出来的参数的一个map
*返回值：sign是拼接好排序后的待签名字串。
***************************************************************/
func genAlipaySignString(mapBody map[string]interface{}) (sign string, err error) {
    sorted_keys := make([]string, 0)
    for k, _ := range mapBody {
        sorted_keys = append(sorted_keys, k)
    }
    sort.Strings(sorted_keys)
    var signStrings string

    index := 0
    for _, k := range sorted_keys {
        fmt.Println("k=", k, "v =", mapBody[k])
        value := fmt.Sprintf("%v", mapBody[k])
        if value != "" {
            signStrings = signStrings + k + "=" + value
        }
        //最后一项后面不要&
        if index < len(sorted_keys)-1 {
            signStrings = signStrings + "&"
        }
        index++
    }

    return signStrings, nil
}

/***************************************************************
*RSA签名验证
*src:待验证的字串，sign:支付宝返回的签名
*pass:返回true表示验证通过
*err :当pass返回false时，err是出错的原因
****************************************************************/
func RSAVerify(src []byte, sign []byte) (pass bool, err error) {
    //步骤1，加载RSA的公钥
    block, _ := pem.Decode([]byte(ALIPAY_PUBLIC_KEY))
    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        fmt.Printf("Failed to parse RSA public key: %s\n", err)
        return
    }
    rsaPub, _ := pub.(*rsa.PublicKey)

    //步骤2，计算代签名字串的SHA1哈希
    t := sha1.New()
    io.WriteString(t, string(src))
    digest := t.Sum(nil)

    //步骤3，base64 decode,必须步骤，支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
    data, _ := base64.StdEncoding.DecodeString(string(sign))

    hexSig := hex.EncodeToString(data)
    fmt.Printf("base decoder: %v, %v\n", string(sign), hexSig)

    //步骤4，调用rsa包的VerifyPKCS1v15验证签名有效性
    err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
    if err != nil {
        fmt.Println("Verify sig error, reason: ", err)
        return false, err
    }

    return true, nil
}
output==>
Parsed Sign: [89 101 115 104 85 112 81 79 49 71 115 82 52 75 120 81 116 65 108 80 122 100 108 113 75 85 77 108 84 102 69 117 110 81 109 119 109 78 73 43 77 74 49 84 50 113 122 100 57 87 117 65 54 98 107 111 72 89 77 77 56 66 112 72 120 116 112 53 109 110 70 77 51 114 88 108 102 103 69 84 86 115 81 99 78 73 105 113 119 67 67 110 49 52 48 49 74 47 117 98 79 107 76 105 50 79 47 109 116 97 50 75 76 120 85 99 109 115 115 81 48 79 110 107 70 73 77 106 106 78 81 117 85 57 78 51 101 73 67 49 90 54 83 122 68 107 111 99 75 48 57 50 119 47 102 51 117 110 52 98 120 107 73 102 73 76 103 100 82 114 48 61]
k= body v = 缴纳保证金
k= buyer_email v = xxaqch@163.com
k= buyer_id v = 2088002578894463
k= discount v = 0.00
k= gmt_create v = 2015-12-21 13:13:28
k= gmt_payment v = 2015-12-21 13:13:28
k= is_total_fee_adjust v = N
k= notify_id v = 5104b719303162e2b79d577aeaa5494jjs
k= notify_time v = 2015-12-21 13:17:45
k= notify_type v = trade_status_sync
k= out_trade_no v = a378c684be7a4f99be1bf3b56e6d38fd
k= payment_type v = 1
k= price v = 0.01
k= quantity v = 1
k= seller_email v = 172886370@qq.com
k= seller_id v = 2088121529348920
k= subject v = 缴纳保证金
k= total_fee v = 0.01
k= trade_no v = 2015122121001004460085270336
k= trade_status v = TRADE_SUCCESS
k= use_coupon v = N
Presign string: body=缴纳保证金&buyer_email=xxaqch@163.com&buyer_id=2088002578894463&discount=0.00&gmt_create=2015-12-21 13:13:28&gmt_payment=2015-12-21 13:13:28&is_total_fee_adjust=N&notify_id=5104b719303162e2b79d577aeaa5494jjs&notify_time=2015-12-21 13:17:45&notify_type=trade_status_sync&out_trade_no=a378c684be7a4f99be1bf3b56e6d38fd&payment_type=1&price=0.01&quantity=1&seller_email=172886370@qq.com&seller_id=2088121529348920&subject=缴纳保证金&total_fee=0.01&trade_no=2015122121001004460085270336&trade_status=TRADE_SUCCESS&use_coupon=N
base decoder: YeshUpQO1GsR4KxQtAlPzdlqKUMlTfEunQmwmNI+MJ1T2qzd9WuA6bkoHYMM8BpHxtp5mnFM3rXlfgETVsQcNIiqwCCn1401J/ubOkLi2O/mta2KLxUcmssQ0OnkFIMjjNQuU9N3eIC1Z6SzDkocK092w/f3un4bxkIfILgdRr0=, 61eb2152940ed46b11e0ac50b4094fcdd96a2943254df12e9d09b098d23e309d53daacddf56b80e9b9281d830cf01a47c6da799a714cdeb5e57e011356c41c3488aac020a7d78d3527fb9b3a42e2d8efe6b5ad8a2f151c9acb10d0e9e41483238cd42e53d3777880b567a4b30e4a1c2b4f76c3f7f7ba7e1bc6421f20b81d46bd
verify sig pass.
</pre>
####Golang显示本机IP
<pre>
package main
//显示本机IP代码
import (
	"fmt"
	"net"
)

func main(){
	addrs ,err :=net.InterfaceAddrs()
	if err != nil{
		panic(err)
	}
	for _,addr :=range addrs{
		fmt.Println(addr.String())
	}
}
output==>
0.0.0.0
0.0.0.0
192.168.1.110
192.168.209.1
192.168.171.1
</pre>
Golang版long2ip与ip2long
<pre>
package main

import (
	"fmt"
	"strconv"
	"regexp"
)
func ip2long(ipstr string)(ip uint32){
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg,err :=regexp.Compile(r)
	if err != nil{
		return 
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil{
		return
	}
	ip1,_:=strconv.Atoi(ips[1])
	ip2,_:=strconv.Atoi(ips[2])
	ip3,_:=strconv.Atoi(ips[3])
	ip4,_:=strconv.Atoi(ips[4])
	if ip1>255||ip2>255||ip3>255||ip4>255{
		return
	}
	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
    ip += uint32(ip3 * 0x100)
    ip += uint32(ip4)
	return 
}
func long2ip(ip uint32) string{
	return fmt.Sprintf("%d.%d.%d.%d",ip>>24,ip<<8>>24,ip<<16>>24,ip<<24>>24)
}
func main(){
}
</pre>
####Golang字符串截取函数substr
<pre>
package main

import (
	"fmt"
)
func substr(str string,start,length int)string{
	rs :=[]rune(str)
	rl :=len(rs)
	end := 0
	
	if start <0 {
		start = rl -1 +start
	}
	end = start +length
	if start >end {
		start,end =end,start
	}
	if start <0 {
		start = 0
	}
	if start >rl{
		start = rl
	}
	if end <0{
		end = 0
	}
	if end >rl{
		end =rl
	}
	return string(rs[start:end])
}
func main(){
	str := "hello,jason"
	fmt.Println(substr(str,6,5))
}
output==>
jason
</pre>
####Golang编写一个守护进程
<pre>
package main
     
import (
        "log"
        "os"
        "os/exec"
        "time"
)
     
func main() {
        lf, err := os.OpenFile("angel.txt", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0600)
        if err != nil {
                os.Exit(1)
        }
        defer lf.Close()
     
        // 日志
        l := log.New(lf, "", os.O_APPEND)
     
        for {
                cmd := exec.Command("/usr/local/bin/node", "/*****.js")
                err := cmd.Start()
                if err != nil {
                        l.Printf("%s 启动命令失败", time.Now().Format("2006-01-02 15:04:05"), err)
     
                        time.Sleep(time.Second * 5)
                        continue
                }
                l.Printf("%s 进程启动", time.Now().Format("2006-01-02 15:04:05"), err)
                err = cmd.Wait()
                l.Printf("%s 进程退出", time.Now().Format("2006-01-02 15:04:05"), err)
     
                time.Sleep(time.Second * 1)
        }
}
 
</pre>
####Golang计算两个经度和纬度之间的距离
<pre>
package main

import (
	"fmt"
	"math"
)
//Golang计算两个经度和纬度之间的距离
func main(){
	lat1 := 29.490295
	lng1 := 106.486654
	lat2 := 29.615467
	lng2 := 106.581515
	fmt.Println(EarthDistance(lat1,lng1,lat2,lng2))
}
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
    var radius float64 = 6371000 // 6378137
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
####Golang使用时间作为种子生成随机数
<pre>
package main

import (
	"fmt"
	"time"
	"math/rand"
)
func main(){
	r :=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<10;i++{
		fmt.Println(r.Intn(100))
	}
}
output==>
16
42
33
51
33
83
20
95
56
64
</pre>
####Golang通过生日算出年龄、生肖和星座
<pre>
package main
 
import (
  "fmt"
  "time"
)
 
func GetTimeFromStrDate(date string) (year, month, day int) {
	const shortForm = "2006-01-02"
	d, err := time.Parse(shortForm, date)
	if err != nil {
		fmt.Println("出生日期解析错误！")
		return 0,0,0
	}
	year = d.Year()
	month = int(d.Month())
	day = d.Day()
	return
}
 
func GetZodiac(year int) (zodiac string) {
	if year <= 0 {
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}
	if x == 0 {
		zodiac = "牛"
	}
	if x == 11 || x == -1 {
		zodiac = "虎"
	}
	if x == 10 || x == -2 {
		zodiac = "兔"
	}
	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}
 
func GetAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}
 
func GetConstellation(month, day int) (star string) {
	if month <= 0 || month >= 13 {
		star = "-1"
	}
	if day <= 0 || day >= 32 {
		star = "-1"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "水瓶座"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "双鱼座"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "白羊座"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "金牛座"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "双子座"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "巨蟹座"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "狮子座"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "处女座"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "天秤座"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "天蝎座"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "射手座"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "魔蝎座"
	}
 
	return star
}
 
func main() {
	y, m, d := GetTimeFromStrDate("1992-03-21")
	fmt.Println(GetAge(y))
	fmt.Println(GetConstellation(m, d))
	fmt.Println(GetZodiac(y))
}
output==>
24
白羊座
猴
</pre>
####Golang获取本机mac地址
<pre>
package main

import (
	"fmt"
	"net"
)
func mac(){
	interfaces,err :=net.Interfaces()
	if err != nil {
		panic("Poor soul,here is what you got:"+err.Error())
	}
	for _,inter:=range interfaces{
		mac :=inter.HardwareAddr
		fmt.Println("MAC:",mac)
	}
}
func main(){
	mac()
}
output==>
MAC: e6:46:19:57:5c:a2
MAC: c4:46:19:57:5c:a2
MAC: 88:ae:1d:3c:38:fc
MAC: 00:50:56:c0:00:01
MAC: 00:50:56:c0:00:08
</pre>
####Golang获取IP地址
<pre>
package main
import (
	"net"
	"fmt"
	"strings"
)
func main() {
	conn, err := net.Dial("udp","google.com:80")
 	if err != nil {
	fmt.Println(err.Error())
	return
}
defer conn.Close()
	fmt.Println(strings.Split(conn.LocalAddr().String(),":")[0])
}
output==>
192.168.1.110
</pre>
####Golang计算两个时间的时间差
<pre>
package main
 
import (
    "fmt"
    "time"
)
 
func main() {
    //Add方法和Sub方法是相反的，获取t0和t1的时间距离d是使用Sub，将t0加d获取t1就是使用Add方法
    k := time.Now()
    //一天之前
    d, _ := time.ParseDuration("-24h")
    fmt.Println(k.Add(d))
 
    //一周之前
    fmt.Println(k.Add(d * 7))
 
    //一月之前
    fmt.Println(k.Add(d * 30))
 
}
output==>
2016-04-03 13:09:04.0291987 +0800 +0800
2016-03-28 13:09:04.0291987 +0800 +0800
2016-03-05 13:09:04.0291987 +0800 +0800
</pre>
####Golang获取系统盘符
<pre>
package main
 
import (
    "fmt"
    . "strconv"
    "syscall"
)
 
func GetLogicalDrives() []string {
    kernel32 := syscall.MustLoadDLL("kernel32.dll")
    GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
    n, _, _ := GetLogicalDrives.Call()
    s := FormatInt(int64(n), 2)
 
    var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：", "V：", "W：", "X：", "Y：", "Z："}
    temp := drives_all[0:len(s)]
 
    var d []string
    for i, v := range s {
 
        if v == 49 {
            l := len(s) - i - 1
            d = append(d, temp[l])
        }
    }
 
    var drives []string
    for i, v := range d {
        drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
    }
    return drives
 
}
 
func main() {
    fmt.Println(GetLogicalDrives())
}
output==>
[C: D: E: F: G: H: J: K: Z：]
</pre>
####Golang通过luhn算法验证信用卡卡号是否有效
<pre>
package main
  
import (
    "fmt"
    "strings"
)
  
const input = `49927398716
49927398717
1234567812345678
1234567812345670`
  
var t = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
  
func luhn(s string) bool {
    odd := len(s) & 1
    var sum int
    for i, c := range s {
        if c < '0' || c > '9' {
            return false
        }
        if i&1 == odd {
            sum += t[c-'0']
        } else {
            sum += int(c - '0')
        }
    }
    return sum%10 == 0
}
  
func main() {
    for _, s := range strings.Split(input, "\n") {
        fmt.Println(s, luhn(s))
    }
}
output==>
49927398716 true
49927398717 false
1234567812345678 false
1234567812345670 true
</pre>
####Golang进行文件分割
<pre>
package main
import (
    "flag"
    "fmt"
    "io"
    "os"
)
 
import "strconv"
 
var infile *string = flag.String("f", "Null", "please input a file name or dir.")
var size *string = flag.String("s", "0", "please input a dst file size.")
 
 
func SplitFile(file *os.File, size int) {
    finfo, err := file.Stat()
    if err != nil {
        fmt.Println("get file info failed:", file, size)
    }
 
    fmt.Println(finfo, size)
 
    //每次最多拷贝1m
    bufsize := 1024 * 1024
    if size < bufsize {
        bufsize = size
    }
 
    buf := make([]byte, bufsize)
 
    num := (int(finfo.Size()) + size - 1) / size
    fmt.Println(num, len(buf))
 
    for i := 0; i < num; i++ {
        copylen := 0
        newfilename := finfo.Name() + strconv.Itoa(i)
        newfile, err1 := os.Create(newfilename)
        if err1 != nil {
            fmt.Println("failed to create file", newfilename)
        } else {
            fmt.Println("create file:", newfilename)
        }
 
        for copylen < size {
            n, err2 := file.Read(buf)
            if err2 != nil && err2 != io.EOF {
                fmt.Println(err2, "failed to read from:", file)
                break
            }
 
            if n <= 0 {
                break
            }
 
            //写文件
            w_buf := buf[:n]
            newfile.Write(w_buf)
            copylen += n
        }
    }
 
    return
}
 
func main() {
    flag.Parse()
 
    if *infile == "Null" {
        fmt.Println("no file to input")
        return
    }
 
    file, err := os.Open(*infile)
    if err != nil {
        fmt.Println("failed to open:", *infile)
    }
 
    defer file.Close()
 
    size, _ := strconv.Atoi(*size)
 
    SplitFile(file, size*1024)
 
}
</pre>
####GolangMessageBox示例
创建一个GUI应用
<pre>
package main
 
import (
       "syscall"
       "unsafe"
       "fmt"
)
 
 
func abort(funcname string, err int) {
       panic(funcname + " failed: " + syscall.Errno(err).Error())
}
 
var (
       kernel32, _ = syscall.LoadLibrary("kernel32.dll")
       getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")
        
       user32, _ = syscall.LoadLibrary("user32.dll")
       messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)
 
 
const (
       MB_OK                      = 0x00000000
       MB_OKCANCEL                = 0x00000001
       MB_ABORTRETRYIGNORE        = 0x00000002
       MB_YESNOCANCEL             = 0x00000003
       MB_YESNO                   = 0x00000004
       MB_RETRYCANCEL             = 0x00000005
       MB_CANCELTRYCONTINUE       = 0x00000006
       MB_ICONHAND                = 0x00000010
       MB_ICONQUESTION            = 0x00000020
       MB_ICONEXCLAMATION         = 0x00000030
       MB_ICONASTERISK            = 0x00000040
       MB_USERICON                = 0x00000080
       MB_ICONWARNING             = MB_ICONEXCLAMATION
       MB_ICONERROR               = MB_ICONHAND
       MB_ICONINFORMATION         = MB_ICONASTERISK
       MB_ICONSTOP                = MB_ICONHAND
 
       MB_DEFBUTTON1              = 0x00000000
       MB_DEFBUTTON2              = 0x00000100
       MB_DEFBUTTON3              = 0x00000200
       MB_DEFBUTTON4              = 0x00000300
)
 
func MessageBox(caption, text string, style uintptr) (result int) {
       // var hwnd HWND
       ret, _, callErr := syscall.Syscall6(uintptr(messageBox), 4,
               0, // HWND
               uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))), // Text
               uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))), // Caption
               style, // type
               0,
               0)
       if callErr != 0 {
               abort("Call MessageBox", int(callErr))
       }
       result = int(ret)
       return
}
 
func main() {
       defer syscall.FreeLibrary(kernel32)
       defer syscall.FreeLibrary(user32)
        
       fmt.Printf("Retern: %d\n", MessageBox("Done Title", "This test is Done.", MB_YESNOCANCEL))
}
 
func init() {
       fmt.Print("Starting Up\n")
}
</pre>
####Golang重写文件
<pre>
package  main
import "fmt"
import "os"
//重写文件，覆盖
func main() {
    fileName := "test.dat"
    dstFile,err := os.Create(fileName)
    if err!=nil{
        fmt.Println(err.Error())   
        return
    }  
 
    defer dstFile.Close()
    s:="hello world"
    dstFile.WriteString(s + "\n")
 
}
</pre>
####Golang指针运算
<pre>
package main
/*
Go语言的语法上是不支持指针运算的，所有指针都在可控的一个范围内
使用，没有C语言的*void然后随意转换指针类型这样的东西。最近在
思考Go如何操作共享内存，共享内存就需要把指针转成不同类型或者
对指针进行运算再获取数据。
*/
import "fmt"
import "unsafe"
 
type Data struct {
    Col1 byte
    Col2 int
    Col3 string
    Col4 int
}
 
func main() {
    var v Data
 
    fmt.Println(unsafe.Sizeof(v))
 
    fmt.Println("**************")
 
    fmt.Println(unsafe.Alignof(v.Col1))
    fmt.Println(unsafe.Alignof(v.Col2))
    fmt.Println(unsafe.Alignof(v.Col3))
    fmt.Println(unsafe.Alignof(v.Col4))
 
    fmt.Println("**************")
 
    fmt.Println(unsafe.Offsetof(v.Col1))
    fmt.Println(unsafe.Offsetof(v.Col2))
    fmt.Println(unsafe.Offsetof(v.Col3))
    fmt.Println(unsafe.Offsetof(v.Col4))
 
    fmt.Println("**************")
 
    v.Col1 = 98
    v.Col2 = 77
    v.Col3 = "1234567890abcdef"
    v.Col4 = 23
 
    fmt.Println(unsafe.Sizeof(v))
 
    fmt.Println("**************")
 
    x := unsafe.Pointer(&v)
 
    fmt.Println(*(*byte)(x))
    fmt.Println(*(*int)(unsafe.Pointer(uintptr(x) + unsafe.Offsetof(v.Col2))))
    fmt.Println(*(*string)(unsafe.Pointer(uintptr(x) + unsafe.Offsetof(v.Col3))))
    fmt.Println(*(*int)(unsafe.Pointer(uintptr(x) + unsafe.Offsetof(v.Col4))))
}
output==>
40
**************
1
8
8
8
**************
0
8
16
32
**************
40
**************
98
77
1234567890abcdef
23
</pre>
####Golang生成缩略图
<pre>
package main
   
import (
    "fmt"
    "os"
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
)
   
func main() {
    f1, err := os.Open("1.jpg")
    if err != nil {
        panic(err)
    }
    defer f1.Close()
   
    f2, err := os.Open("2.jpg")
    if err != nil {
        panic(err)
    }
    defer f2.Close()
   
    f3, err := os.Create("3.jpg")
    if err != nil {
        panic(err)
    }
    defer f3.Close()
   
    m1, err := jpeg.Decode(f1)
    if err != nil {
        panic(err)
    }
    bounds := m1.Bounds()
   
    m2, err := jpeg.Decode(f2)
    if err != nil {
        panic(err)
    }
   
    m := image.NewRGBA(bounds)
    white := color.RGBA{255, 255, 255, 255}
    draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src)
    draw.Draw(m, bounds, m1, image.ZP, draw.Src)
    draw.Draw(m, image.Rect(100, 200, 300, 600), m2, image.Pt(250, 60), draw.Src)
   
    err = jpeg.Encode(f3, m, &jpeg.Options{90})
    if err != nil {
        panic(err)
    }
   
    fmt.Println("ok")
}
</pre>
###Golang实现数据结构-堆栈
####栈
container/list
<pre>
package main

import (
	"fmt"
	"container/list"
)
func main(){
	// 生成队列
	l :=list.New()
	//入队，入栈
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	//队首元素出队
	lf :=l.Front() //队首元素
	l.Remove(lf)
	fmt.Println(lf.Value)
	//队尾元素出栈
	lb :=l.Back()  //队尾元素
	l.Remove(lb)
	fmt.Println(lb.Value)
	
}
output==>
1
3
</pre>
####堆
container/heap
<pre>
package main

import (
	"sort"
	"container/heap"
	"fmt"
)
type IntHeap []int
func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i,j int) bool{return h[i] < h[j]}
func (h IntHeap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}
func (h *IntHeap) Push(x interface{}){
	*h =append(*h,x.(int))
}
func (h *IntHeap) Pop() interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	*h =old[0:n-1]
	return x
}
func main(){
	h :=&IntHeap{100,46,5,6,4,44,5,6,7,56,55}
	fmt.Println("Heap:\n",*h)
	heap.Init(h)
	fmt.Println("最小值:\n",(*h)[0])
	fmt.Println("Heap sort:")
	//for(Pop)依次输出最小值,则相当于执行了HeapSort
	for h.Len() >0{
		fmt.Print(heap.Pop(h)," ")
	}
	fmt.Println()
	//增加一个新值
	fmt.Println("Push(h,3),然后再看看堆：")
	heap.Push(h,3)
	for h.Len()>0{
		fmt.Print(heap.Pop(h),"\n")
	}
	fmt.Println("使用sort.Sort把h2排序：")
	h2 :=IntHeap{100,455,7,1,445,787,67,5,4,55,6,7,787,54,65}
	sort.Sort(h2)
	for _,v :=range h2{
		fmt.Print(v," ")
	}
}
output==>
Heap:
 [100 46 5 6 4 44 5 6 7 56 55]
最小值:
 4
Heap sort:
4 5 5 6 6 7 44 46 55 56 100 
Push(h,3),然后再看看堆：
3
使用sort.Sort把h2排序:
1 4 5 6 7 7 54 55 65 67 100 445 455 787 787 
</pre>
上例分析：<br>
自定义的类,实现相关接口后,交由heap.Init()去构建堆.从堆中Pop()后,数据就被从heap中移除了.升降序由Less()来决定.自定义类也可以直接用Sort来排序,因为实现了相关接口.
####自定义实现堆的相关方法
两相对比之后发现,container/heap 竞然,真的,仅且只是封装了一个堆而已.<br>
把那些不确定的,各种需要定制的东西,都交给了用户去实现.<br>
它仅仅只负责最核心的堆这部份的东西.<br>
这样基础库清爽了,使用堆时也不会再感到缚手缚脚了.
<pre>
package main

import (
	"fmt"
)

var(
  heap = []int{100,16,4,8,70,2,36,22,5,12}	
)


func main(){
	fmt.Println("\n数组:")
	Print(heap)

    MakeHeap()
    fmt.Println("\n构建树后:")
	Print(heap)

	fmt.Println("\n增加 90,30,1 :")
	Push(90)
	Push(30)
	Push(1)	
	Print(heap)	
	
	n := Pop()
	fmt.Println("\nPop出最小值(",n,")后:")
	Print(heap)	

	fmt.Println("\nRemove()掉idx为3即值",heap[3-1],"后:")
	Remove(3)	
	Print(heap)	

	fmt.Println("\nHeapSort()后:")
	HeapSort()
	fmt.Println(heap)	

}


func Print(arr []int){		
	for _,v := range arr {
		fmt.Printf("%d ",v)
	}
}

//构建堆
func MakeHeap(){	
	n := len(heap)	
	for i := n/2-1 ;i >= 0;i--{		
		down(i, n)
	}
}


//由父节点至子节点依次建堆
//parent      : i
//left child  : 2*i+1
//right child : 2*i+2
func down(i,n int) {
	//构建最小堆,父小于两子节点值
	for {	

		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}

		//找出两个节点中最小的(less: a<b)
		j := j1 // left child		
		if j2 := j1 + 1; j2 < n && !Less(j1, j2) {
			j = j2 // = 2*i + 2  // right child
		}

		//然后与父节点i比较,如果父节点小于这个子节点最小值,则break,否则swap
		if !Less(j, i) {
			break
		}
		Swap(i, j)
		i = j	
	}
}


//由子节点至父节点依次重建堆
func up(j int)  {
	
	for {
		i := (j - 1) / 2 // parent      
    
		if i == j || !Less(j, i) { 
			//less(子,父) !Less(9,5) == true 
			//父节点小于子节点,符合最小堆条件,break
			break
		}
		//子节点比父节点小,互换
		Swap(i, j)
		j = i
	}
}

func Push(x interface{}){
	heap = append(heap, x.(int))
	up(len(heap)-1)
	return 
}

func Pop() interface{} {
	n := len(heap) - 1
	Swap(0, n)
	down(0, n)

	old :=heap
	n = len(old)
	x := old[n-1]
	heap = old[0 : n-1]
	return x
}

func Remove(i int) interface{} {
	n := len(heap) - 1
	if n != i {
		Swap(i, n)
		down(i, n)
		up(i)
	}
	return Pop()
}

func Less(a,b int)bool{
	return heap[a] < heap[b]
}
func Swap(a,b int){
	heap[a],heap[b] = heap[b],heap[a]
}
func HeapSort(){
	//升序 Less(heap[a] > heap[b])	//最大堆
	//降序 Less(heap[a] < heap[b])	//最小堆
	for i := len(heap)-1 ;i > 0;i--{	
		//移除顶部元素到数组末尾,然后剩下的重建堆,依次循环
		Swap(0, i)
		down(0, i)
	}
}
output==>
数组:
100 16 4 8 70 2 36 22 5 12 
构建树后:
2 5 4 8 12 100 36 22 16 70 
增加 90,30,1 :
1 5 2 8 12 4 36 22 16 70 90 100 30 
Pop出最小值( 1 )后:
2 5 4 8 12 30 36 22 16 70 90 100 
Remove()掉idx为3即值 4 后:
4 5 8 16 12 30 36 22 100 70 90 
HeapSort()后:
[100 90 70 36 30 22 16 12 8 5 4]
</pre>
####Golang异常打印堆栈错误信息
<pre>
package main

import (
	"runtime/debug"
)
func main(){
	defer func(){
		if err := recover();err !=nil{
			debug.PrintStack()
		}
	}()
	value := 11
	zero := 0
	value =value/zero
	
}
output==>
C:/mygo/src/act/main.go:9 (0x4010be)
	func.001: debug.PrintStack()
c:/go/src/runtime/asm_amd64.s:401 (0x436b0c)
	call16: CALLFN(·call16, 16)
c:/go/src/runtime/panic.go:387 (0x40fd1f)
	gopanic: reflectcall(unsafe.Pointer(d.fn), deferArgs(d), uint32(d.siz), uint32(d.siz))
c:/go/src/runtime/panic.go:24 (0x40eef5)
	panicdivide: panic(divideError)
c:/go/src/runtime/os_windows.go:47 (0x40edc2)
	sigpanic: panicdivide()
C:/mygo/src/act/main.go:14 (0x401048)
	main: value =value/zero
c:/go/src/runtime/proc.go:63 (0x4118da)
	main: main_main()
c:/go/src/runtime/asm_amd64.s:2232 (0x438ae1)
	goexit: 
</pre>
###Golang之JSON序列化
JSON序列化时,Go语言序列化会自动对一些特殊字符会作编码处理。<br>
字符串编码为json字符串。角括号"<"和">"会转义为"\u003c"和"\u003e"以避免某些浏览器吧json输出错误理解为HTML。基于同样的原因，"&"转义为"\u0026"。用Golang开发这类接口的时候需要注意这些。
<pre>
package main  
  
import (  
    "bytes"  
    "encoding/json"  
    "fmt"  
    "time"  
)  
  
type Query struct {  
    AppID     string `json:"AppID"`  
    Timestamp int64  `json:"Timestamp"`  
    Package   string `json:"Package"`  
}  
  
func main() {  
    MarshalDemo()  
}  
  
func MarshalDemo() {  
    v := &Query{}  
    v.AppID = "testid"  
    v.Timestamp = time.Now().Unix()  
    v.Package = "xxcents=100&bank=666"  
  
    data, _ := json.Marshal(v)  
    fmt.Println("Marshal:", string(data))  
  
    data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)  
    data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)  
    data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)  
    data = bytes.Replace(data, []byte("\\u003d"), []byte("="), -1)  
  
    fmt.Println("处理后:", string(data))  
}  
 
output==>
Marshal: {"AppID":"testid","Timestamp":1459765387,"Package":"xxcents=100\u0026bank=666"}
处理后: {"AppID":"testid","Timestamp":1459765387,"Package":"xxcents=100&bank=666"}
</pre>
###七牛文件hash值算法
<pre>
package main
import (
	"fmt"
	"io"
	"os"
	"bytes"
	"crypto/sha1"
	"encoding/base64"
)

const (
	BLOCK_BITS = 22 // Indicate that the blocksize is 4M
	BLOCK_SIZE = 1 << BLOCK_BITS
)

func BlockCount(fsize int64) int {

	return int((fsize + (BLOCK_SIZE-1)) >> BLOCK_BITS)
}

func CalSha1(b []byte, r io.Reader) ([]byte, error) {

	h := sha1.New()
	_, err := io.Copy(h, r)
	if err != nil {
		return nil, err
	}
	return h.Sum(b), nil
}

func GetEtag(filename string) (etag string, err error) {

	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return
	}

	fsize := fi.Size()
	blockCnt := BlockCount(fsize)
	sha1Buf := make([]byte, 0, 21)

	if blockCnt <= 1 { // file size <= 4M
		sha1Buf = append(sha1Buf, 0x16)
		sha1Buf, err = CalSha1(sha1Buf, f) 
		if err != nil {
			return
		}
	} else { // file size > 4M
		sha1Buf = append(sha1Buf, 0x96)
		sha1BlockBuf := make([]byte, 0, blockCnt * 20)
		for i := 0; i < blockCnt; i ++ {
			body := io.LimitReader(f, BLOCK_SIZE)
			sha1BlockBuf, err = CalSha1(sha1BlockBuf, body)
			if err != nil {
				return
			}
		}
		sha1Buf, _ = CalSha1(sha1Buf, bytes.NewReader(sha1BlockBuf))
	}
	etag = base64.URLEncoding.EncodeToString(sha1Buf)
	return
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, `Usage: qetag <filename>`)
		return
	}
	etag, err := GetEtag(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(etag)
}
</pre>
##Golang批量替换和转移目录程序UpUpUp
使用Json文件作为配置文件，遍历指定目录(包含子目录)，对于指定扩展名的文件， 查找并替换文件内容中的指定字符串，并将其输出到新的目录(包含子目录）下。原文件内容不变。至于其它非指定的文件，也一并复制一份到新目录下。
<pre>
//配置文件 flag.json
 {  
"sourcedir":"C:\\mygo\\src\\act\\aa\\",  
"destdir":"C:\\mygo\\src\\act\\bb\\",  
"fileext":[".go",".conf"],  
"replacewhere":[{  
        "findwhat":"parseFile",  
        "replacewith":"----parseFile----"  
    },  
    {  
        "findwhat":"172.18.1.101",  
        "replacewith":"192.168.1.101"  
    }]  
}  
</pre> 
<pre>
//主程序 main.go
package main  
  
import (  
    "bufio"  
    "encoding/json"  
    "fmt"  
    "io"  
    "io/ioutil"  
    "os"  
    "path/filepath"  
    "runtime"  
    "strings"  
    "sync"  
    "time"  
)  
  
const (  
    flagFile = "flag.json"  
)  
  
type RWhere struct {  
    FindWhat    string `json:"findwhat"`  
    ReplaceWith string `json:"replacewith"`  
}  
  
type ReplaceConf struct {  
    SourceDir     string   `json:"sourcedir"`  
    DestDir       string   `json:"destdir"`  
    FileExtension []string `json:"fileext"`  
    ReplaceWhere  []RWhere `json:"replacewhere"`  
    CaseSensitive bool     `json:"casesensitive,omitempty"`  
}  
  
var repConf ReplaceConf  
var repReplacer *strings.Replacer  
var extFileNum, otherFileNum int  
var maxGoroutines int  
  
func init() {  
    maxGoroutines = 10  
}  
  
func main() {  
    now := time.Now()  
    runtime.GOMAXPROCS(runtime.NumCPU())  
  
    parseJsonFile()  
  
    findSourceFiles(repConf.SourceDir)  
  
    end_time := time.Now()  
    var dur_time time.Duration = end_time.Sub(now)  
    fmt.Printf("elapsed %f seconds\n", dur_time.Seconds())  
    fmt.Println("处理统计")  
    fmt.Println("  处理指定类型文件:", extFileNum)  
    fmt.Println("  处理其它文件:", otherFileNum)  
}  
  
func findSourceFiles(dirname string) {  
    waiter := &sync.WaitGroup{}  
    fmt.Println("dirname:", dirname)  
    filepath.Walk(dirname, sourceWalkFunc(waiter))  
    waiter.Wait()  
}  
  
func sourceWalkFunc(waiter *sync.WaitGroup) func(string, os.FileInfo, error) error {  
    return func(path string, info os.FileInfo, err error) error {  
  
        if err == nil && info.Size() > 0 && !info.IsDir() {  
            if runtime.NumGoroutine() > maxGoroutines {  
                parseFile(path, nil)  
            } else {  
                waiter.Add(1)  
                go parseFile(path, func() { waiter.Done() })  
            }  
        } else {  
            fmt.Println("[sourceWalkFunc] err:", err)  
        }  
        return nil  
    }  
}  
  
func parseFile(currfile string, done func()) {  
    if done != nil {  
        defer done()  
    }  
  
    //这地方要注意，配置要对。  
    destFile := strings.Replace(currfile, repConf.SourceDir, repConf.DestDir, -1)  
    if destFile == currfile {  
        panic("[parseFile] ERROR 没有替换对. SourceDir与DestDir配置出问题了。请检查Json配置.")  
    }  
    destDir := filepath.Dir(destFile)  
    if _, er := os.Stat(destDir); os.IsNotExist(er) {  
        if err := os.MkdirAll(destDir, 0700); err != nil {  
            fmt.Println("[parseFile] MkdirAll ", destDir)  
            panic(err)  
        }  
    }  
    fmt.Println("[parseFile] 源文件:", currfile)  
    fmt.Println("[parseFile] 目标文件:", destFile)  
    /////////////////////////////////////////////////  
  
    oldFile, err := os.Open(currfile)  
    if err != nil {  
        fmt.Println("[parseFile] Failed to open the input file ", oldFile)  
        return  
    }  
    defer oldFile.Close()  
  
    newFile, err := os.Create(destFile)  
    if err != nil {  
        panic(err)  
    }  
    defer newFile.Close()  
  
    f1 := func(ext string) bool {  
        for _, e := range repConf.FileExtension {  
            if ext == e {  
                return true  
            }  
        }  
        return false  
    }  
  
    if f1(filepath.Ext(currfile)) {  
        copyRepFile(newFile, oldFile)  
        extFileNum++  
    } else {  
        if _, err := io.Copy(newFile, oldFile); err != nil {  
            panic(err)  
        }  
        otherFileNum++  
    }  
}  
  
func copyRepFile(newFile, oldFile *os.File) {  
    br := bufio.NewReader(oldFile)  
    bw := bufio.NewWriter(newFile)  
  
    for {  
        row, err1 := br.ReadString(byte('\n'))  
        if err1 != nil {  
            break  
        }  
  
        str := string(row)  
        if str == "" {  
            continue  
        }  
  
        ret := repReplacer.Replace(str)  
        //fmt.Println("[copyRepFile] str:", str)  
        //fmt.Println("[copyRepFile] ret:", ret)  
        if _, err := bw.WriteString(ret); err != nil {  
            panic(err)  
        }  
    }  
    bw.Flush()  
}  
  
func parseJsonFile() {  
    f, err := os.Open(flagFile)  
    if err != nil {  
        panic("[parseJsonFile] open failed!")  
    }  
    defer f.Close()  
  
    j, err := ioutil.ReadAll(f)  
    if err != nil {  
        panic("[parseJsonFile] ReadAll failed!")  
    }  
  
    err = json.Unmarshal(j, &repConf)  
    if err != nil {  
        fmt.Println("[parseJsonFile] json err:", err)  
        panic("[parseJsonFile] Unmarshal failed!")  
    }  
  
    fmt.Println(" ------------------------------------------------------")  
    fmt.Println(" 源目录:", repConf.SourceDir)  
    fmt.Println(" 目标目录:", repConf.DestDir)  
    fmt.Println(" 仅包含的指定扩展名的文件:", repConf.FileExtension)  
    for _, e := range repConf.FileExtension {  
        fmt.Println(" 文件扩展名:", e)  
    }  
  
    arr := make([]string, 0, 1)  
    for _, v := range repConf.ReplaceWhere {  
        fmt.Println(" 原文本:", v.FindWhat, " 替换为:", v.ReplaceWith)  
        arr = append(arr, v.FindWhat)  
        arr = append(arr, v.ReplaceWith)  
    }  
    repReplacer = strings.NewReplacer(arr...)  
    fmt.Println(" ------------------------------------------------------")  
  
    if repConf.SourceDir == "" || repConf.DestDir == "" {  
        panic("[parseJsonFile] 目录设置不对!")  
    }  
}
output==>
 源目录: C:\mygo\src\act\aa\
 目标目录: C:\mygo\src\act\bb\
 仅包含的指定扩展名的文件: [.go .conf]
 文件扩展名: .go
 文件扩展名: .conf
 原文本: parseFile  替换为: ----parseFile----
 原文本: 172.18.1.101  替换为: 192.168.1.101
 ------------------------------------------------------
dirname: C:\mygo\src\act\aa\
[sourceWalkFunc] err: <nil>
[sourceWalkFunc] err: <nil>
[sourceWalkFunc] err: <nil>
[parseFile] 源文件: C:\mygo\src\act\aa\Baidusd_Setup_4.2.0.7666.1436769697.exe
[parseFile] 目标文件: C:\mygo\src\act\bb\Baidusd_Setup_4.2.0.7666.1436769697.exe
[parseFile] 源文件: C:\mygo\src\act\aa\ChromeStandalone_V43.0.2357.134_Setup.1436927123.exe
[parseFile] 目标文件: C:\mygo\src\act\bb\ChromeStandalone_V43.0.2357.134_Setup.1436927123.exe
[parseFile] 源文件: C:\mygo\src\act\aa\QQ_V7.4.15197.0_setup.1436951158.exe
[parseFile] 目标文件: C:\mygo\src\act\bb\QQ_V7.4.15197.0_setup.1436951158.exe
elapsed 6.946397 seconds
处理统计
  处理指定类型文件: 0
  处理其它文件: 3
</pre>
###Golang实现CRC32与Adler32算法
<pre>
package main  
// 校验算法(ADLER32/CRC32)例子  
import (  
    "fmt"  
    "hash/adler32"  
    "hash/crc32"  
)  
  
var ADLER32 int = 0  
var CRC32 int = 1  
  
func main() {  
    for _, v := range []string{"aaaaaaaaaa", "3333sdfsdffsdffsd", "234esrewr234324", `An Adler-32 checksum is obtained by calculating two 16-bit checksums A and B and concatenating their bits into a 32-bit integer. A is the sum of all bytes in the stream plus one, and B is the sum of the individual values of A from each step.  
                    At the beginning of an Adler-32 run, A is initialized to 1, B to 0. The sums are done modulo 65521 (the largest prime number smaller than 216). The bytes are stored in network order (big endian), B occupying the two most significant bytes.  
                    The function may be expressed as  
                    A = 1 + D1 + D2 + ... + Dn (mod 65521)  
                     B = (1 + D1) + (1 + D1 + D2) + ... + (1 + D1 + D2 + ... + Dn) (mod 65521)  
                       = n×D1 + (n−1)×D2 + (n−2)×D3 + ... + Dn + n (mod 65521)  
                     Adler-32(D) = B × 65536 + A  
                    where D is the string of bytes for which the checksum is to be calculated, and n is the length of D.`} {  
        calc(ADLER32, []byte(v))  
        calc(CRC32, []byte(v))  
    }  
}  
  
func calc(t int, b []byte) {  
    var ret uint32  
    if ADLER32 == t {  
        ret = adler32.Checksum([]byte(b))  
        fmt.Printf("ADLER32 %15d  : %s...  \n", ret, string(b[:5]))  
    } else if CRC32 == t {  
        ret = crc32.ChecksumIEEE([]byte(b))  
        fmt.Printf("CRC32   %15d  : %s...  \n", ret, string(b[:5]))  
    } else {  
        return  
    }  
}  
output==>
ADLER32       350290891  : aaaaa...  
CRC32        1276235248  : aaaaa...  
ADLER32       839517735  : 3333s...  
CRC32        4210835749  : 3333s...  
ADLER32       622527588  : 234es...  
CRC32        1997398613  : 234es...  
ADLER32      4130281146  : An Ad...  
CRC32        3088706129  : An Ad...  
</pre>
###Golang热更新配置参数
在不停止程序的情况下，通过发送USR1或USR2等信号量，触发运行中程序的参数更新处理。当然还可以通过处理如kill等信号量，让程序正确的处理退出操作。
<pre>
//热更新配置参数    
package main  

import (  
	"fmt"  
	"os"  
	"os/signal"  
	"syscall"  
	"time"  
)  

var gConfig string  

func main() {  
	quit := make(chan bool)  
	readConfig()  
	go signals(quit)  
	go displayConfig(quit)  
	EXIT:  
	for {  
	    select {  
	    case <-quit:  
	        break EXIT  
	    default:  
	    }  
	}  
	fmt.Println("[main()]  exit")  
}  

func signals(q chan bool) bool {  
	sigs := make(chan os.Signal)  
	defer close(sigs)  
	EXIT:  
	for {  
	    signal.Notify(sigs, syscall.SIGQUIT,  
	        syscall.SIGTERM,  
	        syscall.SIGINT,  
	        syscall.SIGUSR1,  
	        syscall.SIGUSR2)  
	
	    sig := <-sigs  
	
	    switch sig {  
	    case syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT:  
	        fmt.Println("[signals()] Interrupt...")  
	        break EXIT  
	    case syscall.SIGUSR1:  
	        fmt.Println("[signals()] syscall.SIGUSR1...")  
	        updateConfig()  
	    case syscall.SIGUSR2:  
	        fmt.Println("[signals()] syscall.SIGUSR2...")  
	        //updateVersion()  
	    default:  
	        break EXIT  
	    }  
	}  
	q <- true  
	return true  
}  

func readConfig() {  
	gConfig = "init"  
	fmt.Println("[readConfig()] ", gConfig)  
}  

func updateConfig() {  
	gConfig = "update"  
	fmt.Println("[updateConfig()] ", gConfig)  
}  

func displayConfig(quit chan bool) {  
	for {  
	    select {  
	    case <-quit:  
	        fmt.Println("[displayConfig()] exit")  
	        return  
	    default:  
	    }  
	    fmt.Println("[displayConfig()] Config:", gConfig)  
	    time.Sleep(time.Second * 2)  
	}  
}  
</pre>
###flag参数解析
<pre>
package main  
//常见用法
import (  
    "flag"  
    "fmt"  
    "os"      
)  
  
var (  
     levelFlag = flag.Int("level", 0, "级别")    
     bnFlag int    
)  
  
func init() {  
     flag.IntVar(&bnFlag, "bn", 3, "份数")    
}  
  
func main() {  
  
    flag.Parse()    
    count := len(os.Args)  
    fmt.Println("参数总个数:",count)  
  
    fmt.Println("参数详情:")  
    for i := 0 ; i < count ;i++{  
        fmt.Println(i,":",os.Args[i])  
    }  
     
    fmt.Println("\n参数值:")  
    fmt.Println("级别:", *levelFlag)  
    fmt.Println("份数:", bnFlag)  
}  
output==>
参数总个数: 1
参数详情:
0 : C:\mygo\src\act\act.exe

参数值:
级别: 0
份数: 3
</pre>
<pre>
package main  
  
import (  
    "flag"  
    "fmt"  
    "os"  
    "time"  
)  
  
var (  
 
    flagSet = flag.NewFlagSet(os.Args[0],flag.ExitOnError)   
    verFlag = flagSet.String("ver", "", "version")  
    xtimeFlag  = flagSet.Duration("time", 10*time.Minute, "time Duration")  
  
    addrFlag = StringArray{}  
)  
  
func init() {  
    flagSet.Var(&addrFlag, "a", "b")  
}  
  
func main() {  
    fmt.Println("os.Args[0]:", os.Args[0])  
    flagSet.Parse(os.Args[1:]) //flagSet.Parse(os.Args[0:])  
  
    fmt.Println("当前命令行参数类型个数:", flagSet.NFlag())    
     for i := 0; i != flagSet.NArg(); i++ {    
        fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))    
    }    
  
    fmt.Println("\n参数值:")  
    fmt.Println("ver:", *verFlag)  
    fmt.Println("xtimeFlag:", *xtimeFlag)  
    fmt.Println("addrFlag:",addrFlag.String())  
  
    for i,param := range flag.Args(){  
        fmt.Printf("---#%d :%s\n",i,param)  
    }  
}  
  
  
type StringArray []string  
  
func (s *StringArray) String() string {  
    return fmt.Sprint([]string(*s))  
}  
  
func (s *StringArray) Set(value string) error {  
    *s = append(*s, value)  
    return nil  
} 
output==>
os.Args[0]: C:\mygo\src\act\act.exe
当前命令行参数类型个数: 0
参数值:
ver: 
xtimeFlag: 10m0s
addrFlag: []
</pre>
###JSON Restful API
<pre>
package main  
  
//简单的JSON Restful API演示(服务端)  
  
import (  
    "encoding/json"  
    "fmt"  
    "net/http"  
    "time"  
)  
  
type Item struct {  
    Seq    int  
    Result map[string]int  
}  
  
type Message struct {  
    Dept    string  
    Subject string  
    Time    int64  
    Detail  []Item  
}  
  
func getJson() ([]byte, error) {  
    pass := make(map[string]int)  
    pass["x"] = 50  
    pass["c"] = 60  
    item1 := Item{100, pass}  
  
    reject := make(map[string]int)  
    reject["l"] = 11  
    reject["d"] = 20  
    item2 := Item{200, reject}  
  
    detail := []Item{item1, item2}  
    m := Message{"IT", "KPI", time.Now().Unix(), detail}  
    return json.MarshalIndent(m, "", "")  
}  
  
func handler(w http.ResponseWriter, r *http.Request) {  
    resp, err := getJson()  
    if err != nil {  
        panic(err)  
    }  
    fmt.Fprintf(w, string(resp))  
}  
  
func main() {  
    http.HandleFunc("/", handler)  
    http.ListenAndServe("localhost:8085", nil)  
}  
//浏览器输入：localhost:8085/
{
"Dept": "IT",
"Subject": "KPI",
"Time": 1459777585,
"Detail": [
{
"Seq": 100,
"Result": {
"c": 60,
"x": 50
}
},
{
"Seq": 200,
"Result": {
"d": 20,
"l": 11
}
}
]
}
</pre>
再用Golang写一个测试程序，来接收localhost:8085传来的json数据，如果没有则报错。
<pre>
package main  
  
//简单的JSON Restful API演示(调用端)  
 
import (  
    "encoding/json"  
    "fmt"  
    "io/ioutil"  
    "net/http"  
    "time"  
)  
  
type Item struct {  
    Seq    int  
    Result map[string]int  
}  
  
type Message struct {  
    Dept    string  
    Subject string  
    Time    int64  
    Detail  []Item  
}  
  
func main() {  
    url := "http://localhost:8085"  
    ret, err := http.Get(url)  
  
    if err != nil {  
        panic(err)  
    }  
    defer ret.Body.Close()  
  
    body, err := ioutil.ReadAll(ret.Body)  
    if err != nil {  
        panic(err)  
    }  
  
    var msg Message  
    err = json.Unmarshal(body, &msg)  
    if err != nil {  
        panic(err)  
    }  
  
    strTime := time.Unix(msg.Time, 0).Format("2006-01-02 15:04:05")  
    fmt.Println("Dept:", msg.Dept)  
    fmt.Println("Subject:", msg.Subject)  
    fmt.Println("Time:", strTime, "\n", msg.Detail)  
} 
</pre>
###Golang自定义错误等级
满屏的error处理会是个悲剧，也不利于对错误进行区分处理。建议在项目中多用自定义错误，再对错误集中处理。
<pre>
package main  
  
//error处理方式演示  

import "fmt"  
import "errors"  
  
func main() {  
  
    errType(test0())  
    errType(test1(" test1 "))  
    errType(test2(500))  
    errType(test3(" test3 "))  
    errType(test4(" test4 "))  
}  
  
type Error1 struct {  
    arg    int  
    errMsg string  
}  
  
func (e *Error1) Error() string {  
    return fmt.Sprintf("%s", e.errMsg)  
}  
  
type Error2 struct {  
    arg    string  
    errMsg string  
}  
  
func (e *Error2) Error() string {  
    return fmt.Sprintf("%s", e.errMsg)  
}  
  
func test0() error {  
    return errors.New("errors.New() - test0()")  
}  
  
func test1(arg string) error {  
    return fmt.Errorf("fmt.Errorf() - test1()")  
}  
  
func test2(arg int) *Error1 {  
    return &Error1{arg, "Error1{} - test2()"}  
}  
  
func test3(arg string) error {  
    return &Error2{arg, "Error2{} - test3()"}  
}  
  
func test4(arg string) *Error2 {  
    return &Error2{arg, "Error2{} - test4() "}  
}  
  
func errType(err interface{}) {  
    switch e := err.(type) {  
    case *Error1:  
        fmt.Println("Type:Error1 ", e)  
    case *Error2:  
        fmt.Println("Type:Error2 ", e)  
    case error:  
        fmt.Println("Type:error ", e)  
    default:  
        fmt.Println("Type:default ", e)  
    }  
}  
output==>
Type:error  errors.New() - test0()
Type:error  fmt.Errorf() - test1()
Type:Error1  Error1{} - test2()
Type:Error2  Error2{} - test3()
Type:Error2  Error2{} - test4() 
</pre>
###信号量与定时器
<pre>
    package main  
      
    //信号量与定时器  
      
    import "fmt"  
    import "os"  
    import "os/signal"  
    import "time"  
      
    func main() {  
      
        sigs := make(chan os.Signal, 1)  
        done := make(chan bool, 1)  
      
        signal.Notify(sigs, os.Interrupt, os.Kill)  
      
        go func() {  
            sig := <-sigs  
            switch sig {  
            case os.Interrupt:  
                fmt.Println("signal: Interrupt")  
            case os.Kill:  
                fmt.Println("signal: Kill")  
            default:  
                fmt.Println("signal: Others")  
            }  
            done <- true  
        }()  
      
        fmt.Println("awaiting signal")  
      
        //main()....  
        go JobTicker(done)  
        <-done  
        close(done)  
        //app.Exit()  
        fmt.Println("exiting")  
    }  
      
    func JobTicker(done <-chan bool) {  
        ticker := time.NewTicker(time.Second)  
        defer ticker.Stop()  
      
        for {  
            select {  
            case <-done:  
                return  
            case <-ticker.C:  
                fmt.Println("ready......")  
            }  
        }  
    } 

output==>
awaiting signal
ready......
ready......
ready......
ready......
ready......
//听说按ctrl+c停止，反正我在win7下测试失败。。。 
</pre>
##Golang一致性哈希
 一致性哈希可用于解决服务器均衡问题.Golang简单实现了一致性哈希，并加入了权重，可采用合适的权重配合算法使用。
<pre>
package main  
  
//一致性哈希(Consistent Hashing)  
  
import (  
    "fmt"  
    "hash/crc32"  
    "sort"  
    "strconv"  
    "sync"  
)  
  
const DEFAULT_REPLICAS = 160  
  
type HashRing []uint32  
  
func (c HashRing) Len() int {  
    return len(c)  
}  
  
func (c HashRing) Less(i, j int) bool {  
    return c[i] < c[j]  
}  
  
func (c HashRing) Swap(i, j int) {  
    c[i], c[j] = c[j], c[i]  
}  
  
type Node struct {  
    Id       int  
    Ip       string  
    Port     int  
    HostName string  
    Weight   int  
}  
  
func NewNode(id int, ip string, port int, name string, weight int) *Node {  
    return &Node{  
        Id:       id,  
        Ip:       ip,  
        Port:     port,  
        HostName: name,  
        Weight:   weight,  
    }  
}  
  
type Consistent struct {  
    Nodes     map[uint32]Node  
    numReps   int  
    Resources map[int]bool  
    ring      HashRing  
    sync.RWMutex  
}  
  
func NewConsistent() *Consistent {  
    return &Consistent{  
        Nodes:     make(map[uint32]Node),  
        numReps:   DEFAULT_REPLICAS,  
        Resources: make(map[int]bool),  
        ring:      HashRing{},  
    }  
}  
  
func (c *Consistent) Add(node *Node) bool {  
    c.Lock()  
    defer c.Unlock()  
  
    if _, ok := c.Resources[node.Id]; ok {  
        return false  
    }  
  
    count := c.numReps * node.Weight  
    for i := 0; i < count; i++ {  
        str := c.joinStr(i, node)  
        c.Nodes[c.hashStr(str)] = *(node)  
    }  
    c.Resources[node.Id] = true  
    c.sortHashRing()  
    return true  
}  
  
func (c *Consistent) sortHashRing() {  
    c.ring = HashRing{}  
    for k := range c.Nodes {  
        c.ring = append(c.ring, k)  
    }  
    sort.Sort(c.ring)  
}  
  
func (c *Consistent) joinStr(i int, node *Node) string {  
    return node.Ip + "*" + strconv.Itoa(node.Weight) +  
        "-" + strconv.Itoa(i) +  
        "-" + strconv.Itoa(node.Id)  
}  
  
// MurMurHash算法 :https://github.com/spaolacci/murmur3  
func (c *Consistent) hashStr(key string) uint32 {  
    return crc32.ChecksumIEEE([]byte(key))  
}  
  
func (c *Consistent) Get(key string) Node {  
    c.RLock()  
    defer c.RUnlock()  
  
    hash := c.hashStr(key)  
    i := c.search(hash)  
  
    return c.Nodes[c.ring[i]]  
}  
  
func (c *Consistent) search(hash uint32) int {  
  
    i := sort.Search(len(c.ring), func(i int) bool { return c.ring[i] >= hash })  
    if i < len(c.ring) {  
        if i == len(c.ring)-1 {  
            return 0  
        } else {  
            return i  
        }  
    } else {  
        return len(c.ring) - 1  
    }  
}  
  
func (c *Consistent) Remove(node *Node) {  
    c.Lock()  
    defer c.Unlock()  
  
    if _, ok := c.Resources[node.Id]; !ok {  
        return  
    }  
  
    delete(c.Resources, node.Id)  
  
    count := c.numReps * node.Weight  
    for i := 0; i < count; i++ {  
        str := c.joinStr(i, node)  
        delete(c.Nodes, c.hashStr(str))  
    }  
    c.sortHashRing()  
}  
  
func main() {  
  
    cHashRing := NewConsistent()  
  
    for i := 0; i < 10; i++ {  
        si := fmt.Sprintf("%d", i)  
        cHashRing.Add(NewNode(i, "172.18.1."+si, 8080, "host_"+si, 1))  
    }  
  
    for k, v := range cHashRing.Nodes {  
        fmt.Println("Hash:", k, " IP:", v.Ip)  
    }  
  
    ipMap := make(map[string]int, 0)  
    for i := 0; i < 1000; i++ {  
        si := fmt.Sprintf("key%d", i)  
        k := cHashRing.Get(si)  
        if _, ok := ipMap[k.Ip]; ok {  
            ipMap[k.Ip] += 1  
        } else {  
            ipMap[k.Ip] = 1  
        }  
    }  
  
    for k, v := range ipMap {  
        fmt.Println("Node IP:", k, " count:", v)  
    }  
}  
output==>
Hash: 639784443  IP: 172.18.1.2
Hash: 3682251441  IP: 172.18.1.3
Hash: 2670724558  IP: 172.18.1.5
Hash: 156374651  IP: 172.18.1.7
Hash: 2948123901  IP: 172.18.1.8
Hash: 1109271580  IP: 172.18.1.9
Hash: 575030648  IP: 172.18.1.0
Hash: 2620094433  IP: 172.18.1.2
Hash: 3640551135  IP: 172.18.1.3
Hash: 685271154  IP: 172.18.1.0
Hash: 3131212918  IP: 172.18.1.3
Hash: 4254065535  IP: 172.18.1.4
Hash: 2154741956  IP: 172.18.1.6
Hash: 3956006673  IP: 172.18.1.7
Hash: 2890459947  IP: 172.18.1.7
Hash: 2795648409  IP: 172.18.1.8
Hash: 3155536303  IP: 172.18.1.0
Hash: 239008809  IP: 172.18.1.1
Hash: 389506850  IP: 172.18.1.2
Hash: 2857473473  IP: 172.18.1.3
Hash: 414058768  IP: 172.18.1.8
Hash: 1753815575  IP: 172.18.1.0
Hash: 812018069  IP: 172.18.1.1
Hash: 931573112  IP: 172.18.1.3
Hash: 2791843655  IP: 172.18.1.4
Hash: 4126764660  IP: 172.18.1.6
Hash: 485364041  IP: 172.18.1.7
Hash: 2068186773  IP: 172.18.1.8
Hash: 4168187427  IP: 172.18.1.8
Hash: 4286109439  IP: 172.18.1.8
Hash: 194091380  IP: 172.18.1.2
Hash: 404072331  IP: 172.18.1.2
Hash: 662027471  IP: 172.18.1.3
Hash: 3414795014  IP: 172.18.1.3
Hash: 3063625456  IP: 172.18.1.4
Hash: 312588205  IP: 172.18.1.6
Hash: 1321327999  IP: 172.18.1.6
Hash: 1237874083  IP: 172.18.1.6
Hash: 832557970  IP: 172.18.1.7
Hash: 4264509387  IP: 172.18.1.9
Hash: 293620475  IP: 172.18.1.1
Hash: 2467890461  IP: 172.18.1.6
Hash: 4226432972  IP: 172.18.1.6
Hash: 4011579018  IP: 172.18.1.8
Hash: 3476688146  IP: 172.18.1.9
Hash: 282567631  IP: 172.18.1.0
Hash: 3323219705  IP: 172.18.1.0
Hash: 1832288372  IP: 172.18.1.2
Hash: 3531821005  IP: 172.18.1.2
Hash: 880676919  IP: 172.18.1.3
Hash: 110337953  IP: 172.18.1.3
Hash: 2339799632  IP: 172.18.1.4
Hash: 3704903351  IP: 172.18.1.5
Hash: 346226969  IP: 172.18.1.6
Hash: 3936106021  IP: 172.18.1.6
Hash: 1476052166  IP: 172.18.1.7
Hash: 3754499237  IP: 172.18.1.9
Hash: 3141391875  IP: 172.18.1.9
Hash: 2628019189  IP: 172.18.1.1
Hash: 178924839  IP: 172.18.1.4
Hash: 2306095520  IP: 172.18.1.6
Hash: 4075865768  IP: 172.18.1.6
Hash: 3247911638  IP: 172.18.1.7
Hash: 2955680553  IP: 172.18.1.7
Hash: 2354541122  IP: 172.18.1.1
Hash: 2883981101  IP: 172.18.1.1
Hash: 2199573799  IP: 172.18.1.1
Hash: 3944577709  IP: 172.18.1.2
Hash: 919684172  IP: 172.18.1.2
Hash: 3791902495  IP: 172.18.1.6
Hash: 1391277483  IP: 172.18.1.9
Hash: 2117970591  IP: 172.18.1.1
Hash: 3095544330  IP: 172.18.1.2
Hash: 2207710104  IP: 172.18.1.5
Hash: 2513618934  IP: 172.18.1.6
Hash: 957402518  IP: 172.18.1.6
Hash: 1449515992  IP: 172.18.1.8
Hash: 1984168000  IP: 172.18.1.9
Hash: 3099132525  IP: 172.18.1.9
...
</pre>
###sync.Once
Golang中的sync.Once，用于实现"只执行一次"的功能。
<pre>
package main  
  
import (  
    "fmt"  
    "sync"  
    "time"  
)  
  
var once sync.Once  
var Gid int  
  
func setup() {  
    Gid++  
    fmt.Println("Called once")  
}  
  
func doprint() {  
    once.Do(setup)  
    fmt.Println("doprint()...")  
}  
  
func main() {  
  
    go doprint()  
    go doprint()  
    go doprint()  
    go doprint()  
  
    time.Sleep(time.Second)  
    fmt.Println("Gid:", Gid)  
}  
output==>
Called once
doprint()...
doprint()...
doprint()...
doprint()...
Gid: 1
</pre>
Golang的sync.Once后面不允许直接传参数，但可以通过以下方法来变通。 
<pre>
package main  
  
import (  
    "fmt"  
    "sync"  
    "time"  
)  
  
var once sync.Once  
var Gid int  
  
func doprint(parm string) {  
    setup := func() {  
        Gid++  
        fmt.Println("Called once! parm:", parm)  
    }  
    once.Do(setup)  
    fmt.Println("doprint()...")  
}  
  
func main() {  
  
    go doprint("1")  
    go doprint("2")  
    go doprint("3")  
    go doprint("4")  
  
    time.Sleep(time.Second)  
    fmt.Println("Gid:", Gid)  
}  
output==>
Called once! parm: 1
doprint()...
doprint()...
doprint()...
doprint()...
Gid: 1
</pre>
###缓存淘汰算法|LRU算法
LRU（Least recently used，最近最少使用）算法根据数据的历史访问记录来进行淘汰数据，其核心思想是“如果数据最近被访问过，那么将来被访问的几率也更高”。<br>
内存管理的一种页面置换算法，对于在内存中但又不用的数据块（内存块）叫做LRU，操作系统会根据哪些数据属于LRU而将其移出内存而腾出空间来加载另外的数据。
####实现
1. 新数据插入到链表头部；
2. 每当缓存命中（即缓存数据被访问），则将数据移到链表头部；
3. 当链表满的时候，将链表尾部的数据丢弃。
####分析
【命中率】

当存在热点数据时，LRU的效率很好，但偶发性的、周期性的批量操作会导致LRU命中率急剧下降，缓存污染情况比较严重。

【复杂度】

实现简单

【代价】

命中时需要遍历链表，找到命中的数据块索引，然后需要将数据移到头部。
<pre>
package main  
  
  
//LRU Cache  
  
import (  
    "fmt"   
	"container/list"  
    "errors" 
)  
  
type CacheNode struct {  
    Key,Value interface{}     
}  
  
func (cnode *CacheNode)NewCacheNode(k,v interface{})*CacheNode{  
    return &CacheNode{k,v}  
}  
  
type LRUCache struct {  
    Capacity int      
    dlist *list.List  
    cacheMap map[interface{}]*list.Element  
}  
  
func NewLRUCache(cap int)(*LRUCache){  
    return &LRUCache{  
                Capacity:cap,  
                dlist: list.New(),  
                cacheMap: make(map[interface{}]*list.Element)}  
}  
  
func (lru *LRUCache)Size()(int){  
    return lru.dlist.Len()  
}  
  
func (lru *LRUCache)Set(k,v interface{})(error){  
  
    if lru.dlist == nil {  
        return errors.New("LRUCache结构体未初始化.")         
    }  
  
    if pElement,ok := lru.cacheMap[k]; ok {       
        lru.dlist.MoveToFront(pElement)  
        pElement.Value.(*CacheNode).Value = v  
        return nil  
    }  
  
    newElement := lru.dlist.PushFront( &CacheNode{k,v} )  
    lru.cacheMap[k] = newElement  
  
    if lru.dlist.Len() > lru.Capacity {        
        //移掉最后一个  
        lastElement := lru.dlist.Back()  
        if lastElement == nil {  
            return nil  
        }  
        cacheNode := lastElement.Value.(*CacheNode)  
        delete(lru.cacheMap,cacheNode.Key)  
        lru.dlist.Remove(lastElement)  
    }  
    return nil  
}  
  
  
func (lru *LRUCache)Get(k interface{})(v interface{},ret bool,err error){  
  
    if lru.cacheMap == nil {  
        return v,false,errors.New("LRUCache结构体未初始化.")         
    }  
  
    if pElement,ok := lru.cacheMap[k]; ok {       
        lru.dlist.MoveToFront(pElement)       
        return pElement.Value.(*CacheNode).Value,true,nil  
    }  
    return v,false,nil  
}  
  
  
func (lru *LRUCache)Remove(k interface{})(bool){  
  
    if lru.cacheMap == nil {  
        return false  
    }  
  
    if pElement,ok := lru.cacheMap[k]; ok {  
        cacheNode := pElement.Value.(*CacheNode)  
        delete(lru.cacheMap,cacheNode.Key)        
        lru.dlist.Remove(pElement)  
        return true  
    }  
    return false  
}  

func main(){  
  
    lru := NewLRUCache(3)  
  
    lru.Set(10,"value1")  
    lru.Set(20,"value2")  
    lru.Set(30,"value3")  
    lru.Set(10,"value4")  
    lru.Set(50,"value5")  
  
    fmt.Println("LRU Size:",lru.Size())  
    v,ret,_ := lru.Get(30)  
    if ret  {  
        fmt.Println("Get(30) : ",v)  
    }  
  
    if lru.Remove(30) {  
        fmt.Println("Remove(30) : true ")  
    }else{  
        fmt.Println("Remove(30) : false ")  
    }  
    fmt.Println("LRU Size:",lru.Size())  
}  
output==>
LRU Size: 3
Get(30) :  value3
Remove(30) : true 
LRU Size: 2
</pre>
###Golang实现二叉查找树
<pre>
package main  
  
//Binary Search Trees  
  
import (  
    "fmt"  
    "math/rand"  
)  
  
func main() {  
  
    t := New(10, 1)  
  
    if Search(t, 6) {  
        fmt.Println("Search(6) true")  
    } else {  
        fmt.Println("Search(6) false")  
    }  
    Print(t)  
  
    if Delete(t, 6) {  
        fmt.Println("Delete(6) true")  
    } else {  
        fmt.Println("Delete(6) false")  
    }  
    Print(t)  
  
    if Delete(t, 9) {  
        fmt.Println("Delete(9) true")  
    } else {  
        fmt.Println("Delete(9) false")  
    }  
    Print(t)  
  
    min, foundMin := GetMin(t)  
    if foundMin {  
        fmt.Println("GetMin() =", min)  
    }  
  
    max, foundMax := GetMax(t)  
    if foundMax {  
        fmt.Println("GetMax() =", max)  
    }  
  
    t2 := New(100, 1)  
    fmt.Println(Compare(t2, New(100, 1)), " Compare() Same Contents")  
    fmt.Println(Compare(t2, New(99, 1)), " Compare() Differing Sizes")  
  
}  
  
type Tree struct {  
    Left  *Tree  
    Value int  
    Right *Tree  
}  
  
func New(n, k int) *Tree {  
    var t *Tree  
    for _, v := range rand.Perm(n) {  
        t = Insert(t, (1+v)*k)  
    }  
    return t  
}  
  
func Insert(t *Tree, v int) *Tree {  
    if t == nil {  
        return &Tree{nil, v, nil}  
    }  
    if v < t.Value {  
        t.Left = Insert(t.Left, v)  
        return t  
    }  
    t.Right = Insert(t.Right, v)  
    return t  
}  
  
//中序遍历  
func Print(t *Tree) { //Recursive  
    if t == nil {  
        return  
    }  
    Print(t.Left)  
    fmt.Println("node:", t.Value)  
    Print(t.Right)  
}  
  
func Search(t *Tree, v int) bool {  
  
    if t == nil {  
        return false  
    }  
    switch {  
    case v == t.Value:  
        return true  
    case v < t.Value:  
        return Search(t.Left, v)  
    case v > t.Value:  
        return Search(t.Right, v)  
    }  
    return false  
}  
  
func GetMin(t *Tree) (int, bool) {  
    if t == nil {  
        return -1, false  
    }  
  
    for {  
        if t.Left != nil {  
            t = t.Left  
        } else {  
            return t.Value, true  
        }  
    }  
}  
  
func GetMax(t *Tree) (int, bool) {  
    if t == nil {  
        return -1, false  
    }  
    for {  
        if t.Right != nil {  
            t = t.Right  
        } else {  
            return t.Value, true  
        }  
    }  
}  
  
func Delete(t *Tree, v int) bool {  
    if t == nil {  
        return false  
    }  
  
    parent := t  
    found := false  
    for {  
        if t == nil {  
            break  
        }  
        if v == t.Value {  
            found = true  
            break  
        }  
  
        parent = t  
        if v < t.Value { //left  
            t = t.Left  
        } else {  
            t = t.Right  
        }  
    }  
  
    if found == false {  
        return false  
    }  
    return deleteNode(parent, t)  
}  
  
func deleteNode(parent, t *Tree) bool {  
    if t.Left == nil && t.Right == nil {  
        fmt.Println("delete() 左右树都为空 ")  
        if parent.Left == t {  
            parent.Left = nil  
        } else if parent.Right == t {  
            parent.Right = nil  
        }  
        t = nil  
        return true  
    }  
  
    if t.Right == nil { //右树为空  
        fmt.Println("delete() 右树为空 ")  
        parent.Left = t.Left.Left  
        parent.Value = t.Left.Value  
        parent.Right = t.Left.Right  
        t.Left = nil  
        t = nil  
        return true  
    }  
  
    if t.Left == nil { //左树为空  
        fmt.Println("delete() 左树为空 ")  
        parent.Left = t.Right.Left  
        parent.Value = t.Right.Value  
        parent.Right = t.Right.Right  
        t.Right = nil  
        t = nil  
        return true  
    }  
  
    fmt.Println("delete() 左右树都不为空 ")  
    previous := t  
    //找到左子节点的最右叶节点，将其值替换至被删除节点  
    //然后将这个最右叶节点清除，所以说，为了维持树，  
    //这种情况下，这个最右叶节点才是真正被删除的节点  
    next := t.Left  
    for {  
        if next.Right == nil {  
            break  
        }  
        previous = next  
        next = next.Right  
    }  
  
    t.Value = next.Value  
    if previous.Left == next {  
        previous.Left = next.Left  
    } else {  
        previous.Right = next.Right  
    }  
    next.Left = nil  
    next.Right = nil  
    next = nil  
    return true  
}  
  
// Walk traverses a tree depth-first,  
// sending each Value on a channel.  
func Walk(t *Tree, ch chan int) {  
    if t == nil {  
        return  
    }  
    Walk(t.Left, ch)  
    ch <- t.Value  
    Walk(t.Right, ch)  
}  
  
// Walker launches Walk in a new goroutine,  
// and returns a read-only channel of values.  
func Walker(t *Tree) <-chan int {  
    ch := make(chan int)  
    go func() {  
        Walk(t, ch)  
        close(ch)  
    }()  
    return ch  
}  
  
// Compare reads values from two Walkers  
// that run simultaneously, and returns true  
// if t1 and t2 have the same contents.  
func Compare(t1, t2 *Tree) bool {  
    c1, c2 := Walker(t1), Walker(t2)  
    for {  
        v1, ok1 := <-c1  
        v2, ok2 := <-c2  
        if !ok1 || !ok2 {  
            return ok1 == ok2  
        }  
        if v1 != v2 {  
            break  
        }  
    }  
    return false  
}  
output==>
Search(6) true
node: 1
node: 2
node: 3
node: 4
node: 5
node: 6
node: 7
node: 8
node: 9
node: 10
delete() 左右树都为空 
Delete(6) true
node: 1
node: 2
node: 3
node: 4
node: 5
node: 7
node: 8
node: 9
node: 10
delete() 右树为空 
Delete(9) true
node: 1
node: 2
node: 3
node: 4
node: 5
node: 8
node: 10
GetMin() = 1
GetMax() = 10
true  Compare() Same Contents
false  Compare() Differing Sizes
</pre>
###Golang实现位图排序(bitmap) 
原理：Golang提供了byte类型，一个byte对应8个位，所以转换一下就可以实现位图了。
<pre>
package main  
   
import (  
   "fmt"  
)  
 
func main() {  
   arrInt32 := [...]uint32{5, 4, 2, 1, 3, 17, 13}  
 
   var arrMax uint32 = 20  
   bit := NewBitmap(arrMax)  
 
   for _, v := range arrInt32 {  
       bit.Set(v)  
   }  
 
   fmt.Println("排序后:")  
   for i := uint32(0); i < arrMax; i++ {  
       if k := bit.Test(i); k == 1 {  
           fmt.Printf("%d ", i)  
       }  
   }  
}  
 
const (  
   BitSize = 8 //一个字节8位  
)  
 
type Bitmap struct {  
   BitArray  []byte  
   ArraySize uint32  
}  
 
func NewBitmap(max uint32) *Bitmap {  
   var r uint32  
   switch {  
   case max <= BitSize:  
       r = 1  
   default:  
       r = max / BitSize  
       if max%BitSize != 0 {  
           r += 1  
       }  
   }  
 
   fmt.Println("数组大小:", r)  
   return &Bitmap{BitArray: make([]byte, r), ArraySize: r}  
}  
 
func (bitmap *Bitmap) Set(i uint32) {  
   idx, pos := bitmap.calc(i)  
   bitmap.BitArray[idx] |= 1 << pos  
   fmt.Println("set()  value=", i, " idx=", idx, " pos=", pos, ByteToBinaryString(bitmap.BitArray[idx]))  
}  
 
func (bitmap *Bitmap) Test(i uint32) byte {  
   idx, pos := bitmap.calc(i)  
   return bitmap.BitArray[idx] >> pos & 1  
}  
 
func (bitmap *Bitmap) Clear(i uint32) {  
   idx, pos := bitmap.calc(i)  
   bitmap.BitArray[idx] &^= 1 << pos  
}  
 
func (bitmap *Bitmap) calc(i uint32) (idx, pos uint32) {  
 
   idx = i >> 3 //相当于i / 8,即字节位置  
   if idx >= bitmap.ArraySize {  
       panic("数组越界.")  
       return  
   }  
   pos = i % BitSize //位位置  
   return  
}  
 
//ByteToBinaryString函数来源:  
// Go语言版byte变量的二进制字符串表示  
// http://www.sharejs.com/codes/go/4357  
func ByteToBinaryString(data byte) (str string) {  
   var a byte  
   for i := 0; i < 8; i++ {  
       a = data  
       data <<= 1  
       data >>= 1  
 
       switch a {  
       case data:  
           str += "0"  
       default:  
           str += "1"  
       }  
 
       data <<= 1  
   }  
   return str  
}  
output==>
数组大小: 3
set()  value= 5  idx= 0  pos= 5 00100000
set()  value= 4  idx= 0  pos= 4 00110000
set()  value= 2  idx= 0  pos= 2 00110100
set()  value= 1  idx= 0  pos= 1 00110110
set()  value= 3  idx= 0  pos= 3 00111110
set()  value= 17  idx= 2  pos= 1 00000010
set()  value= 13  idx= 1  pos= 5 00100000
排序后:
1 2 3 4 5 13 17
</pre>
###Golang实现Rabin-Karp算法
字符串匹配。<br>
为什么是16777619：字符串哈希，会经常用到FNV哈希算法。FNV哈希算法如下：将字符串看作是字符串长度的整数，这个数的进制是一个质数。计算出来结果之后，按照哈希的范围求余数，结果就是哈希结果。继续看Golang的代码，字符串字串匹配用的是无符号32位整数，那就是32位长度，自然，质数就需要选16777619了。结果会按照32位最大的整数求余，在这里，因为是将结果存在uint32里面的，所以超出范围的会被丢弃，也可以认为是求余操作。
<pre>
package main   
  
import (  
    "fmt"  
    "unicode/utf8"     
)  
  
  
func main(){  
    count := Count("9876520210520","520")  
    fmt.Println("count==",count)  
}  
  
  
//primeRK is the prime base used in Rabin-Karp algorithm.  
//primeRK相当于进制  
//本例中,只用到0-9这10个数字,即所有字符的总个数为10,所以定为10  
//源码中是16777619,即相当于16777619进制  
//The magic is in the interesting relationship between the special prime   
 //16777619 (2^24 + 403) and 2^32 and 2^8.   
const primeRK = 10 // 16777619   
  
// hashStr returns the hash and the appropriate multiplicative  
// factor for use in Rabin-Karp algorithm.  
func hashStr(sep string) (uint32, uint32) {  
    hash := uint32(0)  
    charcode := [...]uint32{5,2,0}   
  
    for i := 0; i < len(sep); i++ {  
        //hash = hash*primeRK + uint32(sep[i])  
        hash = hash*primeRK + charcode[i]   
    }  
  
    //即相当于千位->百位->十位,得到乘数因子(pow),本例中的520,得到的pow是1000  
    var pow, sq uint32 = 1, primeRK  
    for i := len(sep); i > 0; i >>= 1 { //len(sep)=3 i>>{1,0} sq:{10,100}  
        if i&1 != 0 {   
            pow *= sq  
        }  
        sq *= sq  
    }  
    /* 
    var pow uint32 = 1   
    for i := len(sep); i > 0; i-- {       
        pow *= primeRK       
    } 
    */  
    fmt.Println("hashStr() sep:",sep," hash:",hash," pow:",pow)  
    return hash, pow  
}  
  
  
// Count counts the number of non-overlapping instances of sep in s.  
func Count(s, sep string) int {  
    fmt.Println("Count() s:",s," sep:",sep)  
  
    n := 0  
    // special cases  
    switch {  
    case len(sep) == 0: //seq为空,返回总数加1  
        return utf8.RuneCountInString(s) + 1  
    case len(sep) == 1: //seq为单个字符,直接遍历比较即可  
        // special case worth making fast  
        c := sep[0]  
        for i := 0; i < len(s); i++ {  
            if s[i] == c {  
                n++  
            }  
        }  
        return n  
    case len(sep) > len(s):  
        return 0  
    case len(sep) == len(s):  
        if sep == s {  
            return 1  
        }  
        return 0  
    }  
    // Rabin-Karp search  
    hashsep, pow := hashStr(sep)   
  
    lastmatch := 0 //最后一次匹配的位置  
    charcode := [...]uint32{9,8,7,6,5,2,0,2,1,0,5,2,0} //对应字符串"9876520210520"  
  
  
    //验证s字符串 0 - len(sep)是不是匹配的  
    h := uint32(0)  
    for i := 0; i < len(sep); i++ {   
        //h = h*primeRK + uint32(s[i])  
        h = h*primeRK +  charcode[i]   
    }  
  
    //如初始s的len(seq)内容是匹配的,n++, lastmatch指向len(seq)位置   
    if h == hashsep && s[:len(sep)] == sep {  
        n++  
        lastmatch = len(sep)  
    }  
  
    for i := len(sep); i < len(s); {   
  
        fmt.Println("\na h ==",h )  
        h *= primeRK  
  
        //加上新的  
        //h += uint32(s[i])   
        h += charcode[i]   
        fmt.Println("b h ==",h )  
  
        // 去掉旧的  
        //h -= pow * uint32(s[i-len(sep)])    
        h -= pow * charcode[i-len(sep)]  
        fmt.Println("c h ==",h )          
        i++  
  
        if h == hashsep && lastmatch <= i-len(sep) && s[i-len(sep):i] == sep {         
            n++  
            lastmatch = i         
            fmt.Println("found n==",n ," lastmatch==",lastmatch)      
  
        }  
    }  
    return n  
}  
output==>
Count() s: 9876520210520  sep: 520
hashStr() sep: 520  hash: 520  pow: 1000

a h == 987
b h == 9876
c h == 876

a h == 876
b h == 8765
c h == 765

a h == 765
b h == 7652
c h == 652

a h == 652
b h == 6520
c h == 520
found n== 1  lastmatch== 7

a h == 520
b h == 5202
c h == 202

a h == 202
b h == 2021
c h == 21

a h == 21
b h == 210
c h == 210

a h == 210
b h == 2105
c h == 105

a h == 105
b h == 1052
c h == 52

a h == 52
b h == 520
c h == 520
found n== 2  lastmatch== 13
count== 2
</pre>
###strings.NewReplacer,replacer.Replace()
对传入参数,能依优先级替换,并能处理中文字符串参数.
<pre>
package main   
        
import (  
    "fmt"  
    "strings"  
)  
  
func main(){  
  
   patterns := []string{    
            "y","25",  
            "中","国",  
            "中工","家伙",  
        }    
          
    replacer := strings.NewReplacer(patterns...)  
  
    format := "中(国)--中工(家伙)"  
    strfmt := replacer.Replace(format)      
    NewReplacer(patterns...);  
    fmt.Println("\nmain() replacer.Replace old=",format)  
    fmt.Println("main() replacer.Replace new=",strfmt)  
}  
 
func NewReplacer(oldnew ...string){  
  
   r :=  makeGenericReplacer(oldnew)  
  
   val,keylen,found := r.lookup("中",true)  
   fmt.Println("\nNewReplacer() 中   val:",val," keylen:",keylen," found:",found)  
  
   val,keylen,found = r.lookup("中工",true)  
   fmt.Println("NewReplacer() 中工 val:",val," keylen:",keylen," found:",found)  
  
   val,keylen,found = r.lookup("y",false)  
   fmt.Println("NewReplacer() y    val:",val," keylen:",keylen," found:",found)  

}  
  
  
type genericReplacer struct {  
    root trieNode  //一个字典树  
    // tableSize is the size of a trie node's lookup table. It is the number  
    // of unique key bytes.  
    tableSize int  
    // mapping maps from key bytes to a dense index for trieNode.table.  
    mapping [256]byte    
}  
  
func makeGenericReplacer(oldnew []string) *genericReplacer {  
    r := new(genericReplacer)  
    // Find each byte used, then assign them each an index.  
    for i := 0; i < len(oldnew); i += 2 { //步长2. 第一个为pattern   
        key := oldnew[i]  
        fmt.Println("\nmakeGenericReplacer() for key=",key)  
  
        //key[j]=utf8存储汉字的三个编码位置中的一个如228,则将其对应位置设置为1  
        //即 r.mapping[228] = 1  
        for j := 0; j < len(key); j++ {  
            r.mapping[key[j]] = 1     
            fmt.Println("makeGenericReplacer() key[",j,"]=",key[j])  
        }  
    }  
  
    for _, b := range r.mapping {   
        r.tableSize += int(b)    
    }  
    fmt.Println("makeGenericReplacer()  r.tableSize=",r.tableSize)  
   
    var index byte  
    for i, b := range r.mapping {  
        if b == 0 {  
            r.mapping[i] = byte(r.tableSize)  
        } else {  
            //依数组字符编码位置,建立索引  
            r.mapping[i] = index  
            fmt.Println("makeGenericReplacer()  r.mapping[",i,"] =",r.mapping[i] )   
            index++  
        }  
    }  
    // Ensure root node uses a lookup table (for performance).  
    r.root.table = make([]*trieNode, r.tableSize)   
      
    //将key,val放入字典树,注意priority=len(oldnew)-i,即越数组前面的,值越大.级别越高  
    for i := 0; i < len(oldnew); i += 2 {  
        r.root.add(oldnew[i], oldnew[i+1], len(oldnew)-i, r)   
    }  
    return r  
}  
  
type trieNode struct {  
    value string  
    priority int  
  
    prefix string  
    next   *trieNode  
    table []*trieNode   
}  
  
func (t *trieNode) add(key, val string, priority int, r *genericReplacer) {  
     fmt.Println("trieNode->add() val=",val," key=",key)  
     if key == "" {  
        if t.priority == 0 {  
            t.value = val  
            t.priority = priority  
            fmt.Println("trieNode->add() t.priority==",priority)  
        }  
        return  
    }  
  
    if t.prefix != "" { //处理已有前缀的node     
        // Need to split the prefix among multiple nodes.  
        var n int // length of the longest common prefix  
        for ; n < len(t.prefix) && n < len(key); n++ { //prefix与key的比较  
            if t.prefix[n] != key[n] {  
                break  
            }  
        }  
        if n == len(t.prefix) {  //相同,继续放下面  
            t.next.add(key[n:], val, priority, r)  
        } else if n == 0 { //没一个相同  
            // First byte differs, start a new lookup table here. Looking up  
            // what is currently t.prefix[0] will lead to prefixNode, and  
            // looking up key[0] will lead to keyNode.  
            var prefixNode *trieNode  
            if len(t.prefix) == 1 {  //如果prefix只是一个字节的字符编码,则挂在节点下面  
                prefixNode = t.next  
            } else {                    //如果不是,将余下的新建一个trie树  
                prefixNode = &trieNode{  
                    prefix: t.prefix[1:],  
                    next:   t.next,  
                }  
            }  
            keyNode := new(trieNode)  
            t.table = make([]*trieNode, r.tableSize) //lookup()中的if node.table != nil   
  
            t.table[r.mapping[t.prefix[0]]] = prefixNode   
            t.table[r.mapping[key[0]]] = keyNode      
            t.prefix = ""  
            t.next = nil  
            keyNode.add(key[1:], val, priority, r)   
        } else {  
            // Insert new node after the common section of the prefix.  
            next := &trieNode{  
                prefix: t.prefix[n:],  
                next:   t.next,  
            }  
            t.prefix = t.prefix[:n]  
            t.next = next  
            next.add(key[n:], val, priority, r)  
        }  
    } else if t.table != nil {  
        // Insert into existing table.  
        m := r.mapping[key[0]]  
        if t.table[m] == nil {  
            t.table[m] = new(trieNode)  
        }  
        t.table[m].add(key[1:], val, priority, r) //构建树        
    } else {    
        t.prefix = key  
        t.next = new(trieNode)  
        t.next.add("", val, priority, r)  
    }  
}  
  
func (r *genericReplacer) lookup(s string, ignoreRoot bool) (val string, keylen int,found bool) {  
    // Iterate down the trie to the end, and grab the value and keylen with  
    // the highest priority.  
    bestPriority := 0  
    node := &r.root  
    n := 0  
  
    for node != nil {  
         if node.priority > bestPriority && !(ignoreRoot && node == &r.root) {  
            bestPriority = node.priority  
            val = node.value  
            keylen = n  
            found = true  
        }  
  
        if s == "" {  
            break  
        }  
  
        if node.table != nil {  
            index := r.mapping[s[0]]  
            if int(index) == r.tableSize { //字符编码第一个字节就没在table中,中断查找  
                break  
            }  
            node = node.table[index]   
            s = s[1:]  
            n++  
        } else if node.prefix != "" && HasPrefix(s, node.prefix) {   
            //字符编码非第一个字节的节点会保留key在prefix中,所以通过分析prefix来继续找其它字节  
            n += len(node.prefix)  
            s = s[len(node.prefix):]  
            node = node.next //继续找相同prefix以外其它字符  
        } else {  
            break  
        }  
    }  
    return  
}  
// HasPrefix tests whether the string s begins with prefix.  
func HasPrefix(s, prefix string) bool {  
    return len(s) >= len(prefix) && s[0:len(prefix)] == prefix  
}  
output==>
makeGenericReplacer() for key= y
makeGenericReplacer() key[ 0 ]= 121

makeGenericReplacer() for key= 中
makeGenericReplacer() key[ 0 ]= 228
makeGenericReplacer() key[ 1 ]= 184
makeGenericReplacer() key[ 2 ]= 173

makeGenericReplacer() for key= 中工
makeGenericReplacer() key[ 0 ]= 228
makeGenericReplacer() key[ 1 ]= 184
makeGenericReplacer() key[ 2 ]= 173
makeGenericReplacer() key[ 3 ]= 229
makeGenericReplacer() key[ 4 ]= 183
makeGenericReplacer() key[ 5 ]= 165
makeGenericReplacer()  r.tableSize= 7
makeGenericReplacer()  r.mapping[ 121 ] = 0
makeGenericReplacer()  r.mapping[ 165 ] = 1
makeGenericReplacer()  r.mapping[ 173 ] = 2
makeGenericReplacer()  r.mapping[ 183 ] = 3
makeGenericReplacer()  r.mapping[ 184 ] = 4
makeGenericReplacer()  r.mapping[ 228 ] = 5
makeGenericReplacer()  r.mapping[ 229 ] = 6
trieNode->add() val= 25  key= y
trieNode->add() val= 25  key= 
trieNode->add() t.priority== 6
trieNode->add() val= 国  key= 中
trieNode->add() val= 国  key= ��
trieNode->add() val= 国  key= 
trieNode->add() t.priority== 4
trieNode->add() val= 家伙  key= 中工
trieNode->add() val= 家伙  key= ��工
trieNode->add() val= 家伙  key= 工
trieNode->add() val= 家伙  key= 
trieNode->add() t.priority== 2

NewReplacer() 中   val: 国  keylen: 3  found: true
NewReplacer() 中工 val: 国  keylen: 3  found: true
NewReplacer() y    val: 25  keylen: 1  found: true

main() replacer.Replace old= 中(国)--中工(家伙)
main() replacer.Replace new= 国(国)--国工(家伙)
</pre>
##非常全面的Golang时间处理函数
<pre>
package main   
  
import (  
    "fmt"  
    "time"  
    "sort" 
	"strings" 
)  
func GoStdTime()string{
	return "2006-01-02 15:04:05"
}

func GoStdUnixDate()string{
    return "Mon Jan _2 15:04:05 MST 2006"
}

func GoStdRubyDate()string{
    return "Mon Jan 02 15:04:05 -0700 2006"
}

func GetTmStr(tm time.Time,format string)(string){
	 patterns := []string{	 		
    		"y","2006",    		
    		"m","01",
    		"d","02",

    		"Y","2006",
    		"M","01",
    		"D","02",

    		"h","03",	//12小时制
    		"H","15",	//24小时制

    		"i","04",
    		"s","05",

    		"t","pm",
    		"T","PM",
    	 }    
    return convStr(tm,format,patterns)
}

func GetTmShortStr(tm time.Time,format string)(string){
		patterns := []string{		
    		"y","06",
    		"m","1",
    		"d","2",

    		"Y","06",
    		"M","1",
    		"D","2",

    		"h","3",  //12小时制
    		"H","15", //24小时制

    		"i","4",
    		"s","5",

    		"t","pm",
    		"T","PM",
    	 }

    return convStr(tm,format,patterns)
}


func convStr(tm time.Time,format string,patterns []string)(string){
	replacer := strings.NewReplacer(patterns...)
    strfmt := replacer.Replace(format)
    return tm.Format(strfmt)
}

func GetLocaltimeStr()(string){
	now := time.Now().Local()
	year,mon,day := now.Date()
	hour,min,sec := now.Clock()
	zone,_ := now.Zone()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d %s",year,mon,day,hour,min,sec,zone)
}

func GetGmtimeStr()(string){
	now := time.Now()
	year,mon,day := now.UTC().Date()
	hour,min,sec := now.UTC().Clock()
	zone,_ := now.UTC().Zone()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d %s",year,mon,day,hour,min,sec,zone)
}

func GetUnixTimeStr(ut int64,format string)(string){
    t := time.Unix(ut,0)
    return GetTmStr(t,format)
}

func GetUnixTimeShortStr(ut int64,format string)(string){
    t := time.Unix(ut,0)
    return GetTmShortStr(t,format)
}

func Greatest(arr []time.Time)(time.Time){
    var temp time.Time
    for _,at := range arr {
        if temp.Before(at) {
            temp = at
        }
    }
    return temp;
}


type TimeSlice []time.Time

func (s TimeSlice) Len() int {
     return len(s) 
 }

func (s TimeSlice) Swap(i, j int) {
     s[i], s[j] = s[j], s[i] 
 }

func (s TimeSlice) Less(i, j int) bool {
    if s[i].IsZero() {
        return false
    }
    if s[j].IsZero() {
        return true
    }
    return s[i].Before(s[j])
}
/*上面的可以单独封装成一个方法*/


func main(){  
    t := time.Now();  
    //alter session set nls_date_format='yyyy-mm-dd hh24:mi:ss';  
    //select to_date('2014-06-09 18:04:06','yyyy-MM-dd HH24:mi:ss') as dt from dual;  
    fmt.Println("\n演示时间 => ",GetTmShortStr(t,"y-m-d H:i:s a"))  
  
    //返回当前是一年中的第几天  
    //select to_char(sysdate,'ddd'),sysdate from dual;  
    yd := t.YearDay();  
    fmt.Println("一年中的第几天: ",yd)         
  
    //一年中的第几周  
    year,week := t.ISOWeek()  
    fmt.Println("一年中的第几周: ",year," | ",week)         
  
    //当前是周几  
    //select to_char(sysdate,'day') from dual;  
    //select to_char(sysdate,'day','NLS_DATE_LANGUAGE = American') from dual;     
    fmt.Println("当前是周几: ",t.Weekday().String())  
  
    //字符串转成time.Time  
    //alter session set nls_date_format='yyyy-mm-dd hh:mi:ss';  
    //select to_date('14-06-09 6:04:06','yy-MM-dd hh:mi:ss') as dt from dual;      
    tt,er := time.Parse(GoStdTime(),"2014-06-09 16:04:06")  
    if(er != nil){  
        fmt.Println("字符串转时间: parse error!")  
    }else{  
        fmt.Println("字符串转时间: ",tt.String())  
    }  
  
  
    fmt.Println("\n演示时间 => ",GetTmStr(t,"y-m-d h:i:s"))  
     
    ta := t.AddDate(1,0,0)     
    fmt.Println("增加一年 => ",GetTmStr(ta,"y-m-d"))  
   
    ta = t.AddDate(0,1,0)  
    fmt.Println("增加一月 => ",GetTmStr(ta,"y-m-d"))  
  
    //select sysdate,sysdate + interval '1' day from dual;  
    ta = t.AddDate(0,0,1) //18  
    fmt.Println("增加一天 => ",GetTmStr(ta,"y-m-d"))  
  
    durdm,_ := time.ParseDuration("432h")  
    ta = t.Add(durdm)  
    fmt.Println("增加18天(18*24=432h) => ",GetTmStr(ta,"y-m-d"))  
   
    //select sysdate,sysdate - interval '7' hour from dual;  
    dur,_ := time.ParseDuration("-2h")  
    ta = t.Add(dur)  
    fmt.Println("减去二小时 => ",GetTmStr(ta,"y-m-d h:i:s"))  
  
    //select sysdate,sysdate - interval '7' MINUTE from dual;  
    durmi,_ := time.ParseDuration("-7m")  
    ta = t.Add(durmi)  
    fmt.Println("减去7分钟 => ",GetTmStr(ta,"y-m-d h:i:s"))  
  
    //select sysdate,sysdate - interval '10' second from dual;  
    durs,_ := time.ParseDuration("-10s")  
    ta = t.Add(durs)  
    fmt.Println("减去10秒 => ",GetTmStr(ta,"y-m-d h:i:s"))  
  
    ttr,er := time.Parse(GoStdTime(),"2014-06-09 16:58:06")  
    if(er != nil){  
        fmt.Println("字符串转时间: 转换失败!")  
    }else{  
        fmt.Println("字符串转时间: ",ttr.String())  
    }  
  
    //alter session set nls_date_format='yyyy-mm-dd hh24:mi:ss';  
    //select trunc(to_date('2014-06-09 16:58:06','yyyy-mm-dd hh24:mi:ss'),'mi') as dt from dual;   
    // SQL => 2014-06-09 16:58:00  
    // Truncate =>  2014-06-09 16:50:00  
    durtr,_ := time.ParseDuration("10m")  
    ta = ttr.Truncate(durtr)  
    fmt.Println("Truncate => ",GetTmStr(ta,"y-m-d H:i:s"))  
  
    //select round(to_date('2014-06-09 16:58:06','yyyy-mm-dd hh24:mi:ss'),'mi') as dt from dual;   
    // SQL => 2014-06-09 16:58:00  
    // Round =>  2014-06-09 17:00:00  
    ta = ttr.Round(durtr)  
    fmt.Println("Round => ",GetTmStr(ta,"y-m-d H:i:s"))  
  
    //日期比较  
    tar1,_ := time.Parse(GoStdTime(),"2014-06-09 19:38:36")  
    tar2,_ := time.Parse(GoStdTime(),"2015-01-14 17:08:26")  
    if tar1.After(tar2) {  
        fmt.Println("tar1 > tar2")  
    }else if tar1.Before(tar2) {  
        fmt.Println("tar1 < tar2")  
    }else{  
        fmt.Println("tar1 = tar2")  
    }  
    tar3,_ := time.Parse(GoStdTime(),"2000-07-19 15:58:16")  
  
    //日期列表中最晚日期  
    // select greatest('2014-06-09','2015-01-14','2000-07-19') from dual;      
    var arr TimeSlice  
    arr = []time.Time{tar1,tar2,tar3}  
    temp := Greatest(arr)  
    fmt.Println("日期列表中最晚日期 => ",GetTmStr(temp,"y-m-d"))      
  
    //日期数组从早至晚排序  
    fmt.Println("\n日期数组从早至晚排序")  
    sort.Sort(arr)  
    for _,at := range arr {  
         fmt.Println("Sort => ",GetTmStr(at,"y-m-d H:i:s"))  
    }  
}  
output==>
演示时间 =>  16-4-4 22:58:29 a
一年中的第几天:  95
一年中的第几周:  2016  |  14
当前是周几:  Monday
字符串转时间:  2014-06-09 16:04:06 +0000 UTC

演示时间 =>  2016-04-04 10:58:29
增加一年 =>  2017-04-04
增加一月 =>  2016-05-04
增加一天 =>  2016-04-05
增加18天(18*24=432h) =>  2016-04-22
减去二小时 =>  2016-04-04 08:58:29
减去7分钟 =>  2016-04-04 10:51:29
减去10秒 =>  2016-04-04 10:58:19
字符串转时间:  2014-06-09 16:58:06 +0000 UTC
Truncate =>  2014-06-09 16:50:00
Round =>  2014-06-09 17:00:00
tar1 < tar2
日期列表中最晚日期 =>  2015-01-14

日期数组从早至晚排序
Sort =>  2000-07-19 15:58:16
Sort =>  2014-06-09 19:38:36
Sort =>  2015-01-14 17:08:26
</pre>
###linux下Golang改变终端字体颜色
<pre>
package main   
  
import (
	"fmt"
	"runtime"
)

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func Black(str string) string {
	return textColor(TextBlack, str)
}

func Red(str string) string {
	return textColor(TextRed, str)
}

func Green(str string) string {
	return textColor(TextGreen, str)
}

func Yellow(str string) string {
	return textColor(TextYellow, str)
}

func Blue(str string) string {
	return textColor(TextBlue, str)
}

func Magenta(str string) string {
	return textColor(TextMagenta, str)
}

func Cyan(str string) string {
	return textColor(TextCyan, str)
}

func White(str string) string {
	return textColor(TextWhite, str)
}

func textColor(color int, str string) string {
	if IsWindows() {
		return str
	}

	switch color {
	case TextBlack:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlack, str)
	case TextRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
	case TextGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
	case TextYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
	case TextBlue:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlue, str)
	case TextMagenta:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextMagenta, str)
	case TextCyan:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextCyan, str)
	case TextWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
	default:
		return str
	}
}

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}

func main() {  
    fmt.Println(Black("Black()"))  
    fmt.Println(Red("Red()"))  
    fmt.Println(Green("Green()"))  
    fmt.Println(Yellow("Yellow()"))  
    fmt.Println(Blue("Blue()"))  
    fmt.Println(Magenta("Magenta()"))  
    fmt.Println(Cyan("Cyan()"))  
    fmt.Println(White("White()"))  
}  
</pre>
###sync.WaitGroup|select发送任务命令到各个客户端
sync.WaitGroup是Golang提供的一种简单的同步方法集合。它有三个方法.

- Add() 添加计数,数目可以为一个,也可以为多个。
- Done() 减掉一个计数,如果计数不为0,则Wait()会阻塞在那,直到全部为0
- Wait() 等待计数为0.
用一个for{}，让用户在作业执行过程中，可以输入指令退出执行。并且，在退出过程中，各个IP也会作相关的取消处理。以保证不会因
强制中断而出现一些不必要的麻烦。
<pre>
package main  
  
import (  
    "bufio"  
    "fmt"  
    "os"  
    "sync"  
    "time"  
)  
  
var waitGrp sync.WaitGroup  
  
func main() {  
  
    ch := make(chan bool)  
  
    go schedule(ch)  
  
    r := bufio.NewReader(os.Stdin)  
    for {  
        time.Sleep(time.Second)  
  
        fmt.Print("Command:> ")  
        ln, _, _ := r.ReadLine()  
        cmd := string(ln)  
  
        if "q" == cmd || "quit" == cmd {  
            close(ch)  
            break  
        } else {  
            fmt.Println(" = cmd = ", cmd, "\n")  
        }  
    }  
  
    waitGrp.Wait()  
    fmt.Println("main() end.")  
}  
  
func schedule(ch chan bool) {  
  
    for _, ip := range []string{"ip1", "ip2"} {  
        waitGrp.Add(1)  
  
        go doJobs(ip, ch)  
        fmt.Println("schedule() IP = ", ip)  
    }  
    fmt.Println("schedule() end.")  
    return  
}  
  
func doJobs(ip string, ch chan bool) {  
  
    defer waitGrp.Done()  
  
    for i := 0; i < 10; i++ {  
  
        select {  
        case <-ch:  
            fmt.Println("doJobs() ", ip, "=>Job Cancel......")  
            return  
        default:  
        }  
  
        fmt.Println("doJobs()...... ", ip, " for:", i)  
        time.Sleep(time.Second)  
    }  
}  
output==>
schedule() IP =  ip1
schedule() IP =  ip2
schedule() end.
doJobs()......  ip1  for: 0
doJobs()......  ip2  for: 0
Command:> doJobs()......  ip2  for: 1
doJobs()......  ip1  for: 1
doJobs()......  ip2  for: 2
doJobs()......  ip1  for: 2
doJobs()......  ip2  for: 3
doJobs()......  ip1  for: 3
doJobs()......  ip2  for: 4
doJobs()......  ip1  for: 4
doJobs()......  ip2  for: 5
doJobs()......  ip1  for: 5
doJobs()......  ip2  for: 6
doJobs()......  ip1  for: 6
doJobs()......  ip2  for: 7
doJobs()......  ip1  for: 7
doJobs()......  ip2  for: 8
doJobs()......  ip1  for: 8
doJobs()......  ip2  for: 9
doJobs()......  ip1  for: 9
q
main() end.
</pre>
###获取中文字符串正确长度的方法
<pre>
package main

import (
	"unicode/utf8"
	"fmt"
)
func main(){
	a :="jason"
	b :="中韩国"
	fmt.Println(utf8.RuneCountInString(a),len(a))
	fmt.Println(utf8.RuneCountInString(b),len(b))
}
output==>
5 5
3 9
</pre>
可以看出，在对有中文的字符串进行计算长度的时候，len()没有utf8.RuneCountInString()来的准确。

####Golang的一个ORM库
地址：https://github.com/donnie4w/gdao/blob/master/gdao.go

###Golang动态调用方法
<pre>
package main

import (
	"reflect"
	"fmt"
)
type YourT struct {
}
func (y *YourT)MethodBar(){
	fmt.Println("MethodBar called")
}
type YourT2 struct{
}
func (y *YourT2)MethodFoo(i int,oo string){
	fmt.Println("MethodFoo called",i,oo)
}
//调用
func InvokeObjectMethod(object interface{},methodName string,args ...interface{}){
	inputs :=make([]reflect.Value,len(args))
	for i,_:=range args{
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}
func main(){
	InvokeObjectMethod(new(YourT2),"MethodFoo",10,"abc")
	InvokeObjectMethod(new(YourT),"MethodBar")
}
output==>
MethodFoo called 10 abc
MethodBar called
</pre> 
##理解Goroutine|深入理解goroutine
<pre>
package main
import (
	"fmt"
	"runtime"
)
 
func say(s string) {
	for i := 0; i < 2; i++ {
 
		if s == "hello" {
 
			fmt.Println("~~ hello")
		} else {
 
			fmt.Println("~~ world")
		}
 
		runtime.Gosched()
		fmt.Println(s)
 
		if s == "hello" {
 
			fmt.Println("2~~ hello")
		} else {
 
			fmt.Println("2~~ world")
		}
 
	}
}
 
func main() {
	go say("world")
 
	say("hello")
 
}
output==>
~~ hello
~~ world
hello
2~~ hello
~~ hello
world
2~~ world
~~ world
hello
2~~ hello
</pre>
代码解析：

1、启goroutine

2、主线程继续执行say(“hello”)

3、主线程输出 ~~hello

4、主线程遇到runtime.Goshed，切换CPU去执行goroutine——say(“world”)

5、输出 ~~world

6、goroutine遇到runtime.Goshed，切换CPU去执行主线程

7、主线程继续向下执行输出 hello，及2~~hello

8、主线程第一次for循环结束，将i++，并输出 ~~hello

9、主线程遇到runtime.Goshed，切换CPU去执行goroutine

10、输出 world及2~~world，第一次for循环结束，输出~~world

11、goroutine再次遇到runtime.Goshed，切换CPU去执行主线程

12、主线程输出hello及2~~hello，将i++已经>2，主线程结束循环退出

###自己实现一个serverHTTP方法
<pre>
package main

import (
    "io"
    "net/http"
)

type a struct{}

func (*a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello Jason.")
}

func main() {
    http.ListenAndServe(":8080", &a{})//第2个参数需要实现Hander的struct，a满足
}
//浏览器打印
hello Jason.
</pre>
解析：<br>
当http.ListenAndServe(":8080", &a{})后，开始等待有访问请求;一旦有访问请求过来，http包帮我们处理了一系列动作后，最后他会去调用a的ServeHTTP这个方法，并把自己已经处理好http.ResponseWriter,*http.Request传进去；而a的ServeHTTP这个方法，拿到*http.ResponseWriter后，并往里面写东西，客户端的网页就显示出来了.
<pre>
package main

import (
	"io"
	"net/http"
)
type a struct{}

//这里ServeHTTP必须写成ServeHTTP,类似serverHTTP的都是错误
func (*a) ServeHTTP(w http.ResponseWriter,r *http.Request){
	path := r.URL.String()
	io.WriteString(w,path)
}
func main(){
	http.ListenAndServe(":8080",&a{})
}
//
访问http://localhost:8080/fff
/fff
</pre>
<pre>
package main

import (
	"io"
	"net/http"
)
type a struct{}

//这里ServeHTTP必须写成ServeHTTP,类似serverHTTP的都是错误
func (*a) ServeHTTP(w http.ResponseWriter,r *http.Request){
	path := r.UserAgent()
	io.WriteString(w,path)
}
func main(){
	http.ListenAndServe(":8080",&a{})
}
//访问http://localhost:8080
Mozilla/5.0 (Windows NT 6.1; WOW64; rv:45.0) Gecko/20100101 Firefox/45.0
</pre>
简单网站雏形
<pre>
package main

import (
	"io"
	"net/http"
)
type a struct{}
func (*a) ServeHTTP(w http.ResponseWriter,r *http.Request){
	path := r.URL.String()
	switch path {
		case "/":
		io.WriteString(w,"Jason Index")
		case "/abc":
		io.WriteString(w,"ABC")
	}
}
func main(){
	http.ListenAndServe(":8089",&a{})
}
//
访问：localhost:8089
Jason index
访问：localhost:8089/abc
ABC
</pre>
上面的switch实现一个类似路由的功能，但是一旦是一个稍微有一点规模的网站，这种做法效率非常低。有问题，自然有对策，ServeMux就登场了。<br>
ServeMux大致作用是，他有一张map表，map里的key记录的是r.URL.String()，而value记录的是一个方法，这个方法和ServeHTTP是一样的，这个方法有一个别名，叫HandlerFunc.ServeMux还有一个方法名字是Handle，他是用来注册HandlerFunc 的ServeMux还有另一个方法名字是ServeHTTP，这样ServeMux是实现Handler接口的，否者无法当http.ListenAndServe的第二个参数传输.
<pre>
//ServeMux实现Golang规则路由
package main

import (
	"io"
	"net/http"
)
type a struct{}
func (*a)ServeHTTP(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello")
}
func main(){
	mux := http.NewServeMux() //新建一个ServeMux
	mux.Handle("/hello",&a{})//注册路由，把"/hello"注册给a这个实现Handler接口的struct，注册到map表中
	http.ListenAndServe(":8089",mux) //)第二个参数是mux
}
//访问 localhost:8089
404 page not found
//访问 localhost:8089/hello
Hello
</pre>
上文解析：<br>
运行时，因为第二个参数是mux，所以http会调用mux的ServeHTTP方法。
ServeHTTP方法执行时，会检查map表（表里有一条数据，key是“/hello”，value是&b{}的ServeHTTP方法）.<br>
如果用户访问/h的话，mux因为匹配上了，mux的ServeHTTP方法会去调用&b{}的 ServeHTTP方法，从而打印hello.<br>
如果用户访问/abc的话，mux因为没有匹配上，从而打印404 page not found.

通过观察上面的示例，我们可以发现struct a 仅仅是为了装一个ServeHTTP而存在，所以可以将struct a 省略掉，直接用过HandleFunc来实现。如下：
<pre>
package main

import (
	"io"
	"net/http"
)
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/abc",func(w http.ResponseWriter,r *http.Request){
		io.WriteString(w,"ABC")
	})
	mux.HandleFunc("/hi",func(w http.ResponseWriter,r *http.Request){
		io.WriteString(w,"hi")
	})
	mux.HandleFunc("/jason",func(w http.ResponseWriter,r *http.Request){
		io.WriteString(w,"I am Jason")
	})
	mux.HandleFunc("/xwq",xwq)
	http.ListenAndServe(":8089",mux)
}
func xwq(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Xwq")
}

//访问可以得到结果：）
</pre>
###Golang指针符号的*和&
理论

&符号的意思是对变量取地址，如：变量a的地址是&a
*符号的意思是对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值

注意点：两个符号抵消顺序

*&可以在任何时间抵消掉，但&*不可以被抵消的，因为顺序不对.

###Golang中bytes.buffer
定义

bytes.buffer是一个缓冲byte类型的缓冲器，这个缓冲器里存放着都是byte.

下面用一个例子解释创建与写入一个缓冲器：<br>
使用Write方法写入，将一个byte类型的slice写入缓冲器尾部
<pre>
package main

import (
	"fmt"
	"bytes"
)

func main(){
	s := []byte("world")
	buf :=bytes.NewBufferString("hello") //创建
	fmt.Println(buf.String())//将buffer里面的数据转成string
	buf.Write(s) //将s这个slice写到buffer的尾部
	fmt.Println(buf.String())
}
output==>
hello
helloworld
</pre>
使用WriteString方法写入，将一个字符串写入到缓冲器的尾部
<pre>
package main

import (
	"fmt"
	"bytes"
)
func main(){
	s :="jason"
	buf := bytes.NewBufferString("hi")
	buf.WriteString(s)
	fmt.Println(buf.String())
	
}
output==>
hijason
</pre>
使用WriteByte方法写入,将一个byte类型的数据写入到缓冲器的尾部
<pre>
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var s byte = '!'
    buf := bytes.NewBufferString("hello")
    fmt.Println(buf.String())  //buf.String()方法是吧buf里的内容转成string，以便于打印
    buf.WriteByte(s) //将s这个string写到buf的尾部
    fmt.Println(buf.String())  //打印 hello!
}
output==>
hello
hello!
</pre>
使用WriteRune方法，将一个rune类型的数据放到缓冲器的尾部
<pre>
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var s rune = '好'
    buf := bytes.NewBufferString("hello")
    fmt.Println(buf.String())  //buf.String()方法是吧buf里的内容转成string，以便于打印
    buf.WriteRune(s) //将s这个string写到buf的尾部
    fmt.Println(buf.String())  //打印 hello好
}
output==>
hello
hello好
</pre>
使用WriteTo方法，将一个缓冲器的数据写到w里，w是实现io.Writer的，比如os.File就是实现io.Writer,可以将缓冲器的内容导出到文件
<pre>
package main

import (
	"fmt"
	"bytes"
	"os"
)
func main(){
	file,_ :=os.Create("jason.txt")
	buf :=bytes.NewBufferString("Json")
	buf.WriteTo(file)//这一步已经把buffer内容导出到jason.txt
	//fmt.Fprintf(file,buf.String())效果同WriteTo	
}
</pre>
###读出缓冲器
使用Read方法
<pre>
package main
import(
    "fmt"
    "bytes"
)

func main() {
    s1:=[]byte("hello")//声明一个slice为s1
    buff:=bytes.NewBuffer(s1)//new一个缓冲器buff，里面存着hello这5个byte
    s2:=[]byte(" world")//声明另一个slice为s2
    buff.Write(s2)//把s2写入添加到buff缓冲器内
    fmt.Println(buff.String()) 
	s3 :=make([]byte,5)//声明一个空的slice,长度为5
	buff.Read(s3)//将buffer内容读入到s3中，当然只读了5个rune过来,同时原buffer会从头部减少5个字符
	fmt.Println(string(s3))
	fmt.Println(buff.String())//头部开始的5个字符被读出，剩下 world
}
output==>
hello world
hello
 world
</pre>
使用ReadByte方法，返回缓冲器头部的第一个byte，缓冲器头部第一个byte被拿掉
<pre>
package main

import (
    "bytes"
    "fmt"
)

func main() {
    buf := bytes.NewBufferString("hello")
    fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，>以便于打印
    b, _ := buf.ReadByte()    //读取第一个byte，赋值给b
    fmt.Println(buf.String()) //打印 ello，缓冲器头部第一个h被拿掉
    fmt.Println(string(b))    //打印 h
}
output==>
hello
ello
h
</pre>
使用ReadRune方法，返回缓冲器头部的第一个rune，缓冲器头部第一个rune被拿掉
<pre>
package main

import (
    "bytes"
    "fmt"
)

func main() {
    buf := bytes.NewBufferString("好hello")
    fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，>以便于打印
    b, n, _ := buf.ReadRune() //读取第一个rune，赋值给b
    fmt.Println(buf.String()) //打印 hello
    fmt.Println(string(b))    //打印中文字： 好，缓冲器头部第一个“好”被拿掉
    fmt.Println(n)            //打印3，“好”作为utf8储存占3个byte
    b, n, _ = buf.ReadRune()  //再读取第一个rune，赋值给b
    fmt.Println(buf.String()) //打印 ello
    fmt.Println(string(b))    //打印h，缓冲器头部第一个h被拿掉
    fmt.Println(n)            //打印 1，“h”作为utf8储存占1个byte
}
output==>
好hello
hello
好
3
ello
h
1
</pre>
ReadBytes和ReadByte根本就不是一回事，ReadBytes需要一个byte作为分隔符，读的时候从缓冲器里找第一个出现的分隔符（delim），找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回，作为byte类型的slice，返回后，缓冲器也会空掉一部分.
<pre>
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var d byte = 'e' //分隔符为e
    buf := bytes.NewBufferString("hello")
    fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，以便于打印
    b, _ := buf.ReadBytes(d)  //读到分隔符，并返回给b
    fmt.Println(buf.String()) //打印 llo，缓冲器被取走一些数据
    fmt.Println(string(b))    //打印 he，找到e了，将缓冲器从头开始，到e的内容都返回给b
}
output==>
hello
llo
he
</pre>
ReadBytes和ReadString基本就是一回事.ReadBytes需要一个byte作为分隔符，读的时候从缓冲器里找第一个出现的分隔符（delim），找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回，作为字符串，返回后，缓冲器也会空掉一部分.
<pre>
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var d byte = 'e' //分隔符为e
    buf := bytes.NewBufferString("hello")
    fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，以便于打印
    b, _ := buf.ReadString(d)  //读到分隔符，并返回给b
    fmt.Println(buf.String()) //打印 llo，缓冲器被取走一些数据
    fmt.Println(b)    //打印 he，找到e了，将缓冲器从头开始，到e的内容都返回给b
}
output==>
hello
llo
he
</pre>
###读入缓冲器（缓冲器变大）
跟WriteTo相对的就是这个ReadForm.从一个实现io.Reader接口的r，把r里的内容读到缓冲器里，n返回读的数量.可以把文件里面的内容写入buffer中。
<pre>
package main

import (
    "bytes"
    "fmt"
    "os"
)

func main() {
    file, _ := os.Open("test.txt")  //test.txt的内容是“world”
    buf := bytes.NewBufferString("hello ")
    buf.ReadFrom(file)              //将text.txt内容追加到缓冲器的尾部
    fmt.Println(buf.String())    //打印“hello world”
}
output==>
hello world
</pre>
###从缓冲器取出（缓冲器变小）
使用Next方法，返回前n个byte，成为slice返回，原缓冲器变小
<pre>
package main

import (
    "bytes"
    "fmt"
)

func main() {
    buf := bytes.NewBufferString("hello")
    fmt.Println(buf.String())
    b := buf.Next(2)   //重头开始，取2个
    fmt.Println(buf.String())  //变小了
    fmt.Println(string(b))   //打印he
}
output==>
hello
llo
he
</pre>
###关于GOlang的slice
因为slice是按值传递的,所有常规修改方式会失效，然后我们强行改成按引用传递，就是传递slice的地址作为参数，这时的修改就是起作用的。所以，对于修改slice的操作要注意，有可能会修改失败。

##Golang的垃圾回收机制理解与优化
下面说一下我对GC的理解。为了保证程序内内存的连续，Golang会申请一大块内存（有时候，只写一个hello, world可能监控内存可能都会发现占用内存比想象中的大）。当用户的程序申请的内存大于之前预申请的内存时，runtime会进行一次GC，并且将GC的阈值翻倍。也就是说，之前是超过10M时进行GC，那么下一次GC就是超过20M才进行。此外，runtime还支持定时GC。我们内存升高的原因，目前看来就是访问量过大，数据库访问的时候导致GC阈值变大，回收频率变低。而且在回收方面，Golang采用了一种拖延症策略，即使是被释放的内存，runtime也不会立刻把内存还给系统。这就导致了内存降不下来，一种内存泄漏的假象。

Golang在GC的时候会发生Stop the world，整个程序会暂停，然后去标记整个内存里面可以被回收的变量，标记完之后恢复程序执行，最后异步得去回收内存。一般这个过程会达到20ms。标记可回收变量的时间取决于临时变量的个数。临时变量数量越多，扫描时间会越长。

所以目前GC的优化方式原则就是尽可能少的声明临时变量：

- 局部变量尽量复用；
- 如果局部变量过多，可以把这些变量放到一个大结构体里面，这样扫描的时候可以只扫描一个变量，回收掉它包含的很多内存；
- 
Golang目前一直在优化GC，目前整体效果来看和大Java差不多，但是稳定性上面来看，还是不行。一般压力测试第一波上面效果极差。

为确保Go程序稳定性，所以先解决的是内存泄漏问题，主要靠memprof来定位问题，接着是进一步提高性能，主要靠cpuprof和自己做的一些统计信息来定位问题。

调优性能的过程中我从cpuprof的结果发现发现gc的scanblock调用占用的cpu竟然有40%多，于是我开始搞各种对象重用和尽量避免不必要的对象创建，效果显著，CPU占用降到了10%多。

但我还是挺不甘心的，想继续优化看看。网上找资料时看到GOGCTRACE这个环境变量可以开启gc调试信息的打印，于是我就在内网测试服开启了，每当go执行gc时就会打印一行信息，内容是gc执行时间和回收前后的对象数量变化。

我惊奇的发现一次gc要20多毫秒，我们服务器请求处理时间平均才33微秒，差了一个量级别呢。

于是我开始关心起gc执行时间这个数值，它到底是一个恒定值呢？还是更数据多少有关呢？

我带着疑问在外网玩家测试的服务器也开启了gc追踪，结果更让我冒冷汗了，gc执行时间竟然达到300多毫秒。go的gc是固定每两分钟执行一次，每次执行都是暂停整个程序的，300多毫秒应该足以导致可感受到的响应延迟。

所以缩短gc执行时间就变得非常必要。从哪里入手呢？首先，可以推断gc执行时间跟数据量是相关的，内网数据少外网数据多。其次，gc追踪信息把对象数量当成重点数据来输出，估计扫描是按对象扫描的，所以对象多扫描时间长，对象少扫描时间短。

于是我便开始着手降低对象数量，一开始我尝试用cgo来解决问题，由c申请和释放内存，这部分c创建的对象就不会被gc扫描了。

但是实践下来发现cgo会导致原有的内存数据操作出些诡异问题，例如一个对象明明初始化了，但还是读到非预期的数据。另外还会引起go运行时报申请内存死锁的错误，我反复读了go申请内存的代码，跟我直接用c的malloc完全都没关联，实在是很诡异。

我只好暂时放弃cgo的方案，另外想了个法子。一个玩家有很多数据，如果把非活跃玩家的数据序列化成一个字节数组，就等于把多个对象压缩成了一个，这样就可以大量减少对象数量。

我按这个思路用快速改了一版代码，放到外网实际测试，对象数量从几百万降至几十万，gc扫描时间降至二十几微秒。

效果不错，但是要用玩家数据时要反序列化，这个消耗太大，还需要再想办法。

于是我索性把内存数据都改为结构体和切片存放，之前用的是对象和单向链表，所以一条数据就会有一个对象对应，改为结构体和结构体切片，就等于把多个对象数据缩减下来。

结果如预期的一样，内存多消耗了一些，但是对象数量少了一个量级。

其实项目之初我就担心过这样的情况，那时候到处问人，对象多了会不会增加gc负担，导致gc时间过长，结果没得到答案。

现在我填过这个坑了，可以确定的说，会。大家就不要再往这个坑跳了。

如果go的gc聪明一点，把老对象和新对象区别处理，至少在我这个应用场景可以减少不必要的扫描，如果gc可以异步进行不暂停程序，我才不在乎那几百毫秒的执行时间呢。

但是也不能完全怪go不完善，如果一开始我早点知道用GOGCTRACE来观测，就可以比较早点发现问题从而比较根本的解决问题。但是既然用了，项目也上了，没办法大改，只能见招拆招了。

总结以下几点给打算用go开发项目或已经在用go开发项目的朋友：
1、尽早的用memprof、cpuprof、GCTRACE来观察程序。
2、关注请求处理时间，特别是开发新功能的时候，有助于发现设计上的问题。
3、尽量避免频繁创建对象(&abc{}、new(abc{})、make())，在频繁调用的地方可以做对象重用。
4、尽量不要用go管理大量对象，内存数据库可以完全用c实现好通过cgo来调用。

示例1，数据结构的重构过程：

最初的数据结构类似这样
<pre>
// 玩家数据表的集合
type tables struct {
        tableA *tableA
        tableB *tableB
        tableC *tableC
        // ...... 此处省略一大堆表
}

// 每个玩家只会有一条tableA记录
type tableA struct {
        fieldA int
        fieldB string
}

// 每个玩家有多条tableB记录
type tableB struct {
        xxoo int
        ooxx int
        next *tableB  // 指向下一条记录
}

// 每个玩家只有一条tableC记录
type tableC struct {
        id int
        value int64
}

</pre>
最初的设计会导致每个玩家有一个tables对象，每个tables对象里面有一堆类似tableA和tableC这样的一对一的数据，也有一堆类似tableB这样的一对多的数据。

假设有1万个玩家，每个玩家都有一条tableA和一条tableC的数据，又各有10条tableB的数据，那么将总的产生1w (tables) + 1w (tableA) + 1w (tableC) + 10w (tableB)的对象。

而实际项目中，表数量会有大几十，一对多和一对一的表参半，对象数量随玩家数量的增长倍数显而易见。

为什么一开始这样设计？

1. 因为有的表可能没有记录，用对象的形式可以用 == nil 来判断是否有记录
2. 一对多的表可以动态增加和删除记录，所以设计成链表
3. 省内存，没数据就是没数据，有数据才有对象

改造后的设计：
<pre>
// 玩家数据表的集合
type tables struct {
        tableA tableA
        tableB []tableB
        tableC tableC
        // ...... 此处省略一大堆表
}

// 每个玩家只会有一条tableA记录
type tableA struct {
        _is_nil bool
        fieldA int
        fieldB string
}

// 每个玩家有多条tableB记录
type tableB struct {
        _is_nil bool
        xxoo int
        ooxx int
}
// 每个玩家只有一条tableC记录
type tableC struct {
        _is_nil bool
        id int
        value int64
} 
</pre>

一对一表用结构体，一对多表用slice，每个表都加一个_is_nil的字段，用来表示当前的数据是否是有用的数据。

这样修改的结果就是，一万个玩家，产生的对象总量是 1w (tables) + 1w ([]tablesB)，跟之前的设计差别很明显。

但是slice不会收缩，而结构体则是一开始就占了内存，所以修改后会导致内存消耗增大。

<h4>其他人的一些经验分享</h4>

如果要逼近有着同样良好设计的C/C++的高性能，需要严格注意频繁分配小块内存的情况。

尤其是对于cpu密集型的应用来说，否则gc的运行会导致大量的cpu消耗在scanblock这个函数上。如果你是通过 make([]T, size)来分配slice，还会有相当一部分的cpu会消耗在 runtime.memclr函数上。

因此，最好使用sync.Pool (since go1.3) 或者利用chan 来写一个对象池。这样能够有效减少gc的负担，将cpu更合理地利用起来。
###xml弹幕文件转换为ass字幕文件
<pre>
// 将bilibili的xml弹幕文件转换为ass字幕文件。
// xml文件中，弹幕的格式如下：
// <d p="32.066,1,25,16777215,1409046965,0,017d3f58,579516441">地板好评</d>
// p的属性为时间、弹幕类型、字体大小、字体颜色、创建时间、？、创建者ID、弹幕ID。
// p的属性中，后4项对ass字幕无用，舍弃。被<d>和</d>包围的是弹幕文本。
// 只处理右往左、上现隐、下现隐三种类型的普通弹幕。
package main
 
import (
    "fmt"
    "io"
    "io/ioutil"
    "math"
    "os"
    "regexp"
    "sort"
    "strconv"
    "strings"
)
 
// ass文件的头部
const header = `[Script Info]
ScriptType: v4.00+
Collisions: Normal
playResX: 640
playResY: 360
 
[V4+ Styles]
Format: Name, Fontname, Fontsize, primaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: Default, Microsoft YaHei, 28, &H00FFFFFF, &H00FFFFFF, &H00000000, &H00000000, 0, 0, 0, 0, 100, 100, 0.00, 0.00, 1, 1, 0, 2, 10, 10, 10, 0
 
[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`
 
// 正则匹配获取弹幕原始信息
var line = regexp.MustCompile(`<d\sp="([\d\.]+),([145]),(\d+),(\d+),\d+,\d+,\w+,\d+">([^<>]+?)</d>`)
 
// 用来保管弹幕的信息
type Danmu struct {
    text  string
    time  float64
    kind  byte
    size  int
    color int
}
 
// 使[]Danmu实现sort.Interface接口，以便排序
type Danmus []Danmu
 
func (d Danmus) Len() int {
    return len(d)
}
func (d Danmus) Less(i, j int) bool {
    return d[i].time < d[j].time
}
func (d Danmus) Swap(i, j int) {
    d[i], d[j] = d[j], d[i]
}
 
// 将正则匹配到的数据填写入Danmu类型里
func fill(d *Danmu, s [][]byte) {
    d.time, _ = strconv.ParseFloat(string(s[1]), 64)
    d.kind = s[2][0] - '0'
    d.size, _ = strconv.Atoi(string(s[3]))
    bgr, _ := strconv.Atoi(string(s[4]))
    d.color = ((bgr >> 16) & 255) | (bgr & (255 << 8)) | ((bgr & 255) << 16)
    d.text = string(s[5])
}
 
// 返回文本的长度，假设ascii字符都是0.5个字长，其余都是1个字长
func length(s string) float64 {
    l := 0.0
    for _, r := range s {
        if r < 127 {
            l += 0.5
        } else {
            l += 1
        }
    }
    return l
}
 
// 生成时间点的ass格式表示：`0:00:00.00`
func timespot(f float64) string {
    h, f := math.Modf(f / 3600)
    m, f := math.Modf(f * 60)
    return fmt.Sprintf("%d:%02d:%05.2f", int(h), int(m), f*60)
}
 
// 读取文件并获取其中的弹幕
func open(name string) ([]Danmu, error) {
    data, err := ioutil.ReadFile(name)
    if err != nil {
        return nil, err
    }
    dan := line.FindAllSubmatch(data, -1)
    ans := make([]Danmu, len(dan))
    for i := len(dan) - 1; i >= 0; i-- {
        fill(&ans[i], dan[i])
    }
    return ans, nil
}
 
// 将弹幕排布并写入w，采用的简单的固定移速、最小重叠排布算法
func save(w io.Writer, dans []Danmu) {
    p1 := make([]float64, 36)
    p2 := make([]float64, 36)
    p3 := make([]float64, 36)
    t := 0
    max := func(x []float64) float64 {
        i := x[0]
        for _, j := range x[1:] {
            if i < j {
                i = j
            }
        }
        return i
    }
    set := func(x []float64, f float64) {
        for i, _ := range x {
            x[i] = f
        }
    }
    find := func(p []float64, f float64, i, d int) int {
        i = (i/d + 1) * d % 36
        m, k := f+10000, 0
        for j := 0; j < 36; j += d {
            t := (i + j) % 36
            if n := max(p[t : t+d]); n <= f {
                k = t
                break
            } else if m > n {
                k = t
                m = n
            }
        }
        return k
    }
    for _, dan := range dans {
        s, l := "", length(dan.text)
        if l == 0 {
            continue
        }
        switch {
        case dan.size < 25:
            dan.size, l, s = 2, l*18, "\\fs18"
        case dan.size == 25:
            dan.size, l = 3, l*28
        case dan.size > 25:
            dan.size, l, s = 4, l*38, "\\fs38"
        }
        if dan.color != 0x00FFFFFF {
            s += fmt.Sprintf("\\c&H%06X", dan.color)
        }
        switch dan.kind {
        case 1: // 右往左
            t := find(p1, dan.time, t, dan.size)
            set(p1[t:t+dan.size], dan.time+8)
            h := (t+dan.size)*10 - 1
            s += fmt.Sprintf("\\move(%d,%d,%d,%d)", 640+int(l/2), h, -int(l/2), h)
            fmt.Fprintf(w, "Dialogue: 1,%s,%s,Default,,0000,0000,0000,,{%s}%s\n",
                timespot(dan.time+0),
                timespot(dan.time+8), s, dan.text)
        case 4: // 下现隐
            j := find(p2, dan.time, 35, dan.size)
            set(p2[j:j+dan.size], dan.time+4)
            s += fmt.Sprintf("\\pos(%d,%d)", 320, (36-j)*10-1)
            fmt.Fprintf(w, "Dialogue: 2,%s,%s,Default,,0000,0000,0000,,{%s}%s\n",
                timespot(dan.time+0),
                timespot(dan.time+4), s, dan.text)
        case 5: // 上现隐
            j := find(p3, dan.time, 35, dan.size)
            set(p3[j:j+dan.size], dan.time+4)
            s += fmt.Sprintf("\\pos(%d,%d)", 320, (j+dan.size)*10-1)
            fmt.Fprintf(w, "Dialogue: 3,%s,%s,Default,,0000,0000,0000,,{%s}%s\n",
                timespot(dan.time+0),
                timespot(dan.time+4), s, dan.text)
        }
    }
}
 
// 主函数，实现了命令行
func main() {
    if len(os.Args) <= 1 {
        os.Exit(0)
    }
    for _, name := range os.Args[1:] {
        dans, err := open(name)
        if err != nil {
            os.Exit(1)
        }
        if n := strings.LastIndex(name, "."); n != -1 {
            name = name[:n]
        }
        name += ".ass"
        file, err := os.Create(name)
        if err != nil {
            os.Exit(2)
        }
        file.WriteString(header)
        sort.Sort(Danmus(dans))
        save(file, dans)
        file.Close()
    }
}
</pre>

###map映射到struct
将Golang的字典格式映射到结构体中。
<pre>
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type User struct {
	Name string
	Age  int8
	Date time.Time
}

func main() {

	data := make(map[string]interface{})
	data["Name"] = "张三"
	data["Age"] = 26
	data["Date"] = "2015-09-29 00:00:00"

	result := &User{}
	err := FillStruct(data, result)
	fmt.Println(err, fmt.Sprintf("%+v", *result))
}

//用map填充结构
func FillStruct(data map[string]interface{}, obj interface{}) error {
	for k, v := range data {
		err := SetField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

//用map的值替换结构的值
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	var err error
	if structFieldType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}

	structFieldValue.Set(val)
	return nil
}

//类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}
	//else if .......增加其他一些类型的转换
	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
output==>
<nil> {Name:张三 Age:26 Date:2015-09-29 00:00:00 +0800 +0800}
</pre>
###Win32Api
<pre>
package win32api

import (
	"fmt"
	"unsafe"
)

// #define WIN32_LEAN_AND_MEAN
// #include <windows.h>
import "C"
import "syscall"

func GetCurrentDirectory() string {
	if bufLen := C.GetCurrentDirectoryW(0, nil); bufLen != 0 {
		buf := make([]uint16, bufLen)
		if bufLen := C.GetCurrentDirectoryW(bufLen, (*C.WCHAR)(&buf[0])); bufLen != 0 {
			return syscall.UTF16ToString(buf)
		}
	}
	return ""
}

func SetWallPaper(file_path string) error {
	path := []byte(file_path)
	result := int(C.SystemParametersInfo(C.SPI_SETDESKWALLPAPER, 0, C.PVOID(unsafe.Pointer(&path[0])), C.SPIF_UPDATEINIFILE))
	if result != C.TRUE {
		return fmt.Errorf("", C.GetLastError())
	}
	return nil
}
</pre>
###获取差值最小的那个数
<pre>
package main
 
import (
     
    "fmt"
)
 
func main(){
    arr:=[]int{12,16,29,34,39,43,55,64,71,89,90,9}
    zuijin:=get_zuijin(88,arr)
    fmt.Println(zuijin)
}
 
func get_zuijin(this int,arr []int) int{
    min:=0 
    if this==arr[0]{
        return arr[0]
    }else if this>arr[0]{
        min = this-arr[0]
    }else if this<arr[0]{
        min = arr[0]-this
    }
     
    for _,v:=range arr{
        if v==this{
            return v    
        }else if v>this{
            if min>v-this{
                min = v-this   
            }
        }else if v<this{
            if min>this-v{
                min = this-v    
            }
        }
    }
     
    for _,v:=range arr{
        if this+min == v{
            return v    
        }else if this-min == v{
            return v    
        }    
    }
    return min
}
output==>
89
</pre>
###Golang写文件读文件
<pre>
package main
 //追加文件
import (
    "io/ioutil"
    "os"
 
    "fmt"
)
 //这里若a.txt不存在会报错
func main() {
    file_write("hello world!\n", "a.txt")
    content := file_read("a.txt")
    fmt.Println(content)
}
 
func file_read(path string) string {
    fi, err := os.Open(path)
    if err != nil {
		panic(err)

    }
    defer fi.Close()
    fd, err := ioutil.ReadAll(fi)
    fmt.Println(err)
    return string(fd)
}
 
func file_write(context, file string) {
    f, err := os.OpenFile(file, os.O_APPEND, 0644)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    f.WriteString(context)
}
output==>
<nil>
hello world!
hello world!
hello world!
hello world!
</pre>
###Golang利用http.Client post数据
<pre>
package main
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
)
 
func main() {
    v := url.Values{}
    v.Set("huifu", "hello world")
    body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
    client := &http.Client{}
    req, _ := http.NewRequest("POST", "http://192.168.2.83:8080/bingqinggongxiang/test2", body)
 
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	 //这个一定要加，不加form的值post不过去，被坑了两小时
    fmt.Printf("%+v\n", req) //看下发送的结构
 
    resp, err := client.Do(req) //发送
    defer resp.Body.Close()     //一定要关闭resp.Body
    data, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(data), err)
}
output==>
&{Method:POST URL:http://192.168.2.83:8080/bingqinggongxiang/test2 Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Content-Type:[application/x-www-form-urlencoded; param=value]] Body:{Reader:0xc082002640} ContentLength:0 TransferEncoding:[] Close:false Host:192.168.2.83:8080 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr: RequestURI: TLS:<nil>}
</pre>
###Golang生成csv文件
<pre>
package main
 
import (
    "encoding/csv"
    "os"
)
 
func main() {
    f, err := os.Create("test.csv")//创建文件
    if err != nil {
        panic(err)
    }
    defer f.Close()
 
    f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
 
    w := csv.NewWriter(f)//创建一个新的写入文件流
    data := [][]string{
        {"1", "中国", "23"},
        {"2", "美国", "23"},
        {"3", "bb", "23"},
        {"4", "bb", "23"},
        {"5", "bb", "23"},
    }
    w.WriteAll(data)//写入数据
    w.Flush()
}

</pre>
###Golang UDP
传统方法：
<pre>
package main  
  
import (  
    "fmt"  
    "net"  
    "os"  
)  
  
func main() {  
    addr, err := net.ResolveUDPAddr("udp", ":6000")  
    if err != nil {  
        fmt.Println("net.ResolveUDPAddr fail.", err)  
        os.Exit(1)  
    }  
  
    conn, err := net.ListenUDP("udp", addr)  
    if err != nil {  
        fmt.Println("net.ListenUDP fail.", err)  
        os.Exit(1)  
    }  
    defer conn.Close()  
  
    for {  
        buf := make([]byte, 65535)  
        rlen, remote, err := conn.ReadFromUDP(buf)  
        if err != nil {  
            fmt.Println("conn.ReadFromUDP fail.", err)  
            continue  
        }  
        go handleConnection(conn, remote, buf[:rlen])  
    }  
}  
  
func handleConnection(conn *net.UDPConn, remote *net.UDPAddr, msg []byte) {  
    service_addr, err := net.ResolveUDPAddr("udp", ":6001")  
    if err != nil {  
        fmt.Println("net.ResolveUDPAddr fail.", err)  
        return  
    }  
  
    service_conn, err := net.DialUDP("udp", nil, service_addr)  
    if err != nil {  
        fmt.Println("net.DialUDP fail.", err)  
        return  
    }  
    defer service_conn.Close()  
  
    _, err = service_conn.Write([]byte("request servcie x"))  
    if err != nil {  
        fmt.Println("service_conn.Write fail.", err)  
        return  
    }  
  
    buf := make([]byte, 65535)  
    rlen, err := service_conn.Read(buf)  
    if err != nil {  
        fmt.Println("service_conn.Read fail.", err)  
        return  
    }  
  
    conn.WriteToUDP(buf[:rlen], remote)  
}  
</pre>
下面的解决了传统方法的两个问题：

1. 延时(Latency)：Server与后端Service之间采用短链接通信，对于UDP类无连接方式影响不大，但是对于TCP类有连接方式，开销还是比较客观的，增加了请求的响应延时
2. 并发(Concurrency)：16位的端口号数量有限，如果每次后端交互都需要新建连接，理论上来说，同时请求后端Service的Goroutine数量无法超过65535这个硬性限制，在如今这个动辄“十万”“百万”高并发时代，最高6w并发貌似不太拿得出手

使用过多线程并发模型的同学应该已经注意到，这两个问题在多线程模型中同样存在，只是不如Golang如此突出：创建的线程数量一般是受控的，不会达到端口上限，但是Goer显然不能满足于这个量级的并发度。

改进方法：
<pre>
package main  
 /*
抽取出独立的通信代理（conn-proxy），代理本身处理所有的
网络通信细节（连接管理，数据收发等），具体的
process-goroutine通过channel与communication-proxy
进行交互（提交请求，等待响应等）
*/
import (  
    "fmt"  
    "net"  
    "os"  
    "strconv"  
    "time"  
)  
  
type Request struct {  
    isCancel bool  
    reqSeq   int  
    reqPkg   []byte  
    rspChan  chan<- []byte  
}  
  
func main() {  
    addr, err := net.ResolveUDPAddr("udp", ":6000")  
    if err != nil {  
        fmt.Println("net.ResolveUDPAddr fail.", err)  
        os.Exit(1)  
    }  
  
    conn, err := net.ListenUDP("udp", addr)  
    if err != nil {  
        fmt.Println("net.ListenUDP fail.", err)  
        os.Exit(1)  
    }  
    defer conn.Close()  
  
    reqChan := make(chan *Request, 1000)  
    go connHandler(reqChan)  
  
    var seq int = 0  
    for {  
        buf := make([]byte, 1024)  
        rlen, remote, err := conn.ReadFromUDP(buf)  
        if err != nil {  
            fmt.Println("conn.ReadFromUDP fail.", err)  
            continue  
        }  
        seq++  
        go processHandler(conn, remote, buf[:rlen], reqChan, seq)  
    }  
}  
  
func processHandler(conn *net.UDPConn, remote *net.UDPAddr, msg []byte, reqChan chan<- *Request, seq int) {  
    rspChan := make(chan []byte, 1)  
    reqChan <- &Request{false, seq, []byte(strconv.Itoa(seq)), rspChan}  
    select {  
    case rsp := <-rspChan:  
        fmt.Println("recv rsp. rsp=%v", string(rsp))  
    case <-time.After(2 * time.Second):  
        fmt.Println("wait for rsp timeout.")  
        reqChan <- &Request{isCancel: true, reqSeq: seq}  
        conn.WriteToUDP([]byte("wait for rsp timeout."), remote)  
        return  
    }  
  
    conn.WriteToUDP([]byte("all process succ."), remote)  
}  
  
func connHandler(reqChan <-chan *Request) {  
    addr, err := net.ResolveUDPAddr("udp", ":6001")  
    if err != nil {  
        fmt.Println("net.ResolveUDPAddr fail.", err)  
        os.Exit(1)  
    }  
  
    conn, err := net.DialUDP("udp", nil, addr)  
    if err != nil {  
        fmt.Println("net.DialUDP fail.", err)  
        os.Exit(1)  
    }  
    defer conn.Close()  
  
    sendChan := make(chan []byte, 1000)  
    go sendHandler(conn, sendChan)  
  
    recvChan := make(chan []byte, 1000)  
    go recvHandler(conn, recvChan)  
  
    reqMap := make(map[int]*Request)  
    for {  
        select {  
        case req := <-reqChan:  
            if req.isCancel {  
                delete(reqMap, req.reqSeq)  
                fmt.Println("CancelRequest recv. reqSeq=%v", req.reqSeq)  
                continue  
            }  
            reqMap[req.reqSeq] = req  
            sendChan <- req.reqPkg  
            fmt.Println("NormalRequest recv. reqSeq=%d reqPkg=%s", req.reqSeq, string(req.reqPkg))  
        case rsp := <-recvChan:  
            seq, err := strconv.Atoi(string(rsp))  
            if err != nil {  
                fmt.Println("strconv.Atoi fail. err=%v", err)  
                continue  
            }  
            req, ok := reqMap[seq]  
            if !ok {  
                fmt.Println("seq not found. seq=%v", seq)  
                continue  
            }  
            req.rspChan <- rsp  
            fmt.Println("send rsp to client. rsp=%v", string(rsp))  
            delete(reqMap, req.reqSeq)  
        }  
    }  
}  
  
func sendHandler(conn *net.UDPConn, sendChan <-chan []byte) {  
    for data := range sendChan {  
        wlen, err := conn.Write(data)  
        if err != nil || wlen != len(data) {  
            fmt.Println("conn.Write fail.", err)  
            continue  
        }  
        fmt.Println("conn.Write succ. data=%v", string(data))  
    }  
}  
  
func recvHandler(conn *net.UDPConn, recvChan chan<- []byte) {  
    for {  
        buf := make([]byte, 1024)  
        rlen, err := conn.Read(buf)  
        if err != nil || rlen <= 0 {  
            fmt.Println(err)  
            continue  
        }  
        fmt.Println("conn.Read succ. data=%v", string(buf))  
        recvChan <- buf[:rlen]  
    }  
}  
</pre>
###使用channel控制并发
<pre>
package main

import (
	"fmt"
)
var quit chan int
func foo(id int){
	fmt.Println(id)
	quit <- 0
}
func main(){
	count := 10
	quit = make(chan int)
	for i:=0;i<count;i++{
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
###关于死锁
会发生死锁的几种情况：

- 只在单一的goroutine里操作无缓冲信道，一定死锁。比如你只在main函数里操作信道
<pre>
package main
import "fmt"
func main(){
	ch := make(chan int)
	ch <- 1
	fmt.Println("This is Great")
}
</pre>
- 主线等ch1中的数据流出，ch1等ch2的数据流出，但是ch2等待数据流入，两个goroutine都在等，也就是死锁
<pre>
package main
import "fmt"
var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)

func say(s string) {
    fmt.Println(s)
    ch1 <- <- ch2 // ch1 等待 ch2流出的数据
}

func main() {
    go say("hello")
    <- ch1  // 堵塞主线
}
</pre>
- 非缓冲信道上如果发生流入无流出，或者流出无流入，导致发生死锁（除了channel操作没有执行的情况）
<pre>
package main

func main(){
	c, quit := make(chan int), make(chan int)
	go func() {
	  c <- 1  // c通道的数据没有被其他goroutine读取走，堵塞当前goroutine
	  quit <- 0 // quit始终没有办法写入数据
	}()
	<- quit // quit 等待数据的写
}
</pre>
对于上面的情况，有一个反例，可以补充说明一下，在只有channel单向操作的情况下，程序依然可以正常执行。
<pre>
package main
func main(){
	c := make(chan int)
	go func(){
		c<- 1
	}()
}
</pre>
解析：main又没等待其它goroutine，自己先跑完了， 所以没有数据流入c信道，一共执行了一个main, 并且没有发生阻塞，所以没有死锁错误。

死锁解决方法：

- 很简单，把没有取走的数据取走，没放入的数据放入,因为无缓冲channel不能存储数据；
- 将无缓冲channel变成缓冲channel，保证cap值大于等于channel里面将要处理的数据量。缓冲信道是先进先出的，我们可以把缓冲信道看作为一个线程安全的队列.类似于Python中的队列Queue.
###Go非侵入式接口<-Go语言编程
要想了解Golang在接口方面独特的魅力，那还必须了解其他高级语言接口设计的理念。下面是现在常见的接口实现方式：
<pre>
//抽象接口
interface IFly{
	virtual void Fly() = 0;
};
//实现类
class Bird:public IFly{
	public :Bird(){}
	virtual ~Bird(){}
	public :void Fly(){
		//以鸟的方式飞行
	}
}
void main(){
	IFly* pFly = new Bird();
	pFly->Fly();
	delete pFly;
}
</pre>
显然，在实现一个接口之前必须先定义该接口，并且类型与接口紧密绑定，即接口的修改会影响到所有实现了该接口的类型，而Golang的接口体系则避免了这类问题：
<pre>
type Bird struct{
	...
}
func (b *Bird) Fly(){
	//以鸟的方式飞行
}
</pre>
我们在实现Bird类型时完全没有任何IFly的信息。我们可以在另外一个地方定义这个IFly()接口：
<pre>
type IFly() interface{
	Fly()
}
</pre>
这两者目前看起来完全没有关系，现在我们如何使用它们：
<pre>
func main(){
	var fly IFly() = new(Bird)
	fly.Fly()
}
</pre>
可以看出，虽然Bird类型与接口实现的时候，没有声明与接口IFly()的关系，但接口与类可以直接转换，甚至接口的定义都不用再类型声明之前，这种比较松散的对应关系可以大幅降低因为接口调整而导致的大量代码调整工作。
####并发编程<- Go语言编程
Go语言实现了CSP(Communicating Sequential Process)模型来作为goroutine间的推荐通信方式。在CSP模型中，一个并发系统由若干并行运行的顺序进程组成，每个进程不能对其他进程的变量赋值。进程之间只能通过一对通信原语实现协作。Go语言用channel这个概念轻巧地实现了CSP模型。channel的使用方比较接近Unix系统中的管道概念，可以方便地进行跨goroutine的通信。

另外，由于一个进程内创建的所有goroutine运行在同一个内存地址空间中， 因此如果不同的goroutine不得不去访问共享的内存变量。访问前应该先获取相应的读写锁。Go中的sync包提供了完备的读写锁功能。

一个使用goroutine与channel进行并行计算的示例：
<pre>
package main
import "fmt"
func sum(values []int, resultChan chan int){
	sum := 0
	for _,value := range values {
		sum +=value
	}
	resultChan <- sum //将计算结果写入resultChan中
}
func main(){
	values := []int{1,2,3,4,5,6,7,8,9,10}
	go sum(values[:len(values)/2],resultChan)
	go sum(values[len(values)/2:],resultChan)
	sum1,sum2 := <-resultChan,<- resultChan 
	fmt.Println("Result：",sum1,sum2,sum1+sum2)
}
</pre>
####GDB调试<- Go语言编程
不用设置什么编译选项，Go语言编译的二进制程序直接支持GDB调试，比如之前用go build test.go编译出来的可执行文test,就可以使用下面命调试模式运行:
<pre>
	gdb test
</pre>
需要注意的是，Go编译器生成的调试信息格式为DWARFv3，只要版本高于7.1的GDB应该都支持它。
###关于Golang内存管理(复用内存Buffer)
Go程序中有两个独立的线程，用来标记不再被程序使用的内存（这就是垃圾收集），并在其不再被使用时返还给操作系统（在Go代码中称为收割，scavenging）。

下面的一个小程序会不断产生内存垃圾，每秒生成一个5MB到10MB的字节数组。它维护了一个20个这样的字节数组大小的内存池，随机丢弃内存池中的字节数组。这个程序用来模拟程序中经常发生的场景：程序的各个部分每时每刻都会分配内存，一些分配的内存一直都在使用，大多数分配的内存都不再使用。在一个Go写的网络程序中，在处理网络链接或请求的Go协程里，这种情况很容易发生。常常是这样的，Go协程分配内存块（比如分配一个slices来存储接收的数据），然后就不再使用。随着时间的积累，会有一系列的内存块被正在被处理的网络链接占用，也会有一些累计的来自那些被处理过的链接的内存垃圾。
<pre>
package main
import (
	"fmt"
	"time"
	"runtime"
	"math/rand"
)
func makeBuffer() []byte{
	return make([]byte,rand.Intn(5000000)+5000000)
}
func main(){
	pool :=make([][]byte,20)
	var m runtime.MemStats
	makes := 0
	for {
		b := makeBuffer()
		makes += 1
		i := rand.Intn(len(pool))
		pool[i ]= b
		time.Sleep(time.Second)
		bytes := 0
		for i :=0;i<len(pool);i++{
			if pool[i] !=nil {
				bytes +=len(pool[i])
			}
		}
		runtime.ReadMemStats(&m)
		fmt.Println(m.HeapSys,bytes,m.HeapAlloc,m.HeapIdle,m.HeapReleased,makes)
	}
}
</pre>
可以观察到，电脑的物理内存使用量不断上升。

这个程序使用runtime.ReadMemStats函数来获取堆大小的信息。这个函数会打印四个值：HeapSys （程序向操作系统请求的内存的字节数），HeapAlloc （当前堆中已经分配的字节数），HeapIdle （堆中未使用的字节数）和HeapReleased （归还给操作系统的字节数）。

Go程序中垃圾收集运行的很频繁（查看GOGC环境变量来理解如何控制GC操作 ）。因此，在运行过程中，堆的大小会随着内存被标记为未使用（这将导致HeapAlloc 和HeapIdle 随之变化）而变化。收割线程只有在内存5分钟都没有使用才会释放内存，因此HeapReleased 并不经常变化。

随着程序的运行，堆中未使用的内存又被重新利用，很少会被释放给操作系统。

解决这种问题的一个方法就是在程序中部分地手动管理内存。比如，使用一个管道，可以单独维护一个不再使用字节数组的内存池，当需要新的字节数组时，从内存池中拿（当内存池为空就生成新的字节数组）。
<pre>
package main
 
import (
    "fmt"
    "math/rand"
    "runtime"
    "time"
)
 
func makeBuffer() []byte {
    return make([]byte, rand.Intn(5000000)+5000000)
}
 
func main() {
    pool := make([][]byte, 20)
 
    buffer := make(chan []byte, 5)
 
    var m runtime.MemStats
    makes := 0
    for {
        var b []byte
        select {
        case b = <-buffer:
        default:
            makes += 1
            b = makeBuffer()
        }
 
        i := rand.Intn(len(pool))
        if pool[i] != nil {
            select {
            case buffer <- pool[i]:
                pool[i] = nil
            default:
            }
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
        fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
            m.HeapIdle, m.HeapReleased, makes)
    }
}
</pre>
内存池中的内存和从操作系统请求的内存很接近。垃圾收集器也基本不做什么。堆中只有很少量的未使用内存最终返还给操作系统。

这种内存复用机制的关键是一个缓存的管道buffer。上面的代码中可以存储5个字节数组。当程序需要一个字节数组时，优先使用select从缓存的管道中去取:
<pre>
select {
    case b = <-buffer:
    default:
        b = makeBuffer()
}
</pre>
select永远不会阻塞因为如果buffer 管道中有字节数组，第一个分支生效，字节数组赋给了 b。如果管道是空的话（也就意味着receive会阻塞），default 分支会执行，并分配了一个新的字节数组。

把字节数组放回到管道中使用了类似的无阻塞模式:
<pre>
select {
    case buffer <- pool[i]:
        pool[i] = nil
    default:
}
</pre>
如果buffer 管道已经满了，往管道里面发送就会阻塞。这种情况下，default分支执行，什么也不做。这种简单的机制可以用来安全的生成一个共享的内存池。由于管道通信对多go协程是安全的，这种机制也可以用于go协程的共享。

实际上，我们在Go程序中使用了类似的技术。下面的代码是真实复用器的简化版。使用一个go协程处理字节数组的生成并在软件中共享给所有的go协程。两个管道get （获取一个新的字节数组）和give （返回字节数组到内存池中）在所有的通信中都被使用。

复用器保存了一个返回的字节数组的链表，间断地丢弃那些时间太久，并不再会被复用（示例代码中，生命周期超过1分钟）的字节数组。这使得程序处理对字符数组的动态需求。
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
这些技术可以在程序员知道内存会被复用而不需要垃圾收集器参与时用来复用内存。它可以显著的减少程序需要内存的大小。并不仅限于字节数组。任何Go类型都可以用类似的行为进行复用。
###Golang实战经验总结
####变量作用域
重名的变量，由于作用域的不同导致错误。比如：
<pre>
var err error
for i:=0;i<3;i++{
	socket,err := getSocket()
	if err !=nil{
		continue
	}
	socket.Write(...)
	socket.Read(...)
}
return err
</pre>
上面的代码循环内的err和循环外的err不是同一个，即使出现网络异常，外面的err任然是nil.下面这样做才是正确的。
<pre>
var err error
var socket Socket
for i:=0;i<3;i++{
	socket,err := getSocket()
	if err != nil {
		continue
	}
	socket.Write(...)
	socket.Read(...)
}
return err
</pre>
####先检查错误才能defer
如果不先检查错误，defer的程序可能panic,错误示例如下：
<pre>
f,err := os.Open(filename)
defer f.Close()//如果文件打开失败，此处会panic
if err !=nil{
	return err
}
</pre>
正确应该是
<pre>
f,err := os.Open(filename)
if err !=nil{
	return err
}
defer f.Close()
</pre>
####defer在函数调用结束后才会被调用
由于defer在函数调用结束后才会被调用，因此在循环中使用defer会影响性能
<pre>
func Process(){
	for i := 0;i<10;i++{
		defer fmt.Println(i)
	}
	fmt.Println("end")
}
output==>
end 
987654321
</pre>
####Golang中的map顺序是随机
<pre>
package main

import (
	"fmt"
)
func main(){
	numbers := map[string]int{
		"one":1,
		"two":2,
		"three":3,
		"four":4,
	}
	for i:=0;i<2;i++{
		for k,v:=range numbers{
			fmt.Println(k,"=",v)
		}
	}
}
output==>
three = 3
four = 4
one = 1
two = 2
</pre>
####Golang之racket版本
<pre>
package main

import (
	"fmt"
)
func 我是(x string,n int)[]string{
	if n == 0{
		return make([]string,0)
	}else{
		return append(我是(x,n-1),x)
	}
}
func main(){
	三 := 3
	猪头 := "猪头"
	s :=我是(猪头,三)
	fmt.Println(s)
}
output==>
[猪头 猪头 猪头]
</pre>
###GIF动画
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
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

func main() {
    // The sequence of images is deterministic unless we seed
    // the pseudo-random number generator using the current time.
    // Thanks to Randall McPherson for pointing out the omission.
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
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
    gif.EncodeAll(out, &anim) //NOTE: ignoring encoding errors
}
</pre>

##goroutine背后的系统知识
这篇文章初衷是希望能为比较缺少系统编程背景的Web开发人员介绍一下goroutine背后的系统知识。

1. 操作系统与运行库
2. 并发与并行 (Concurrency and Parallelism)
3. 线程的调度
4. 并发编程框架
5. goroutine

一 、 操作系统与运行库

对于普通的电脑用户来说，能理解应用程序是运行在操作系统之上就足够了，可对于开发者，我们还需要了解我们写的程序是如何在操作系统之上运行起来的，操作系统如何为应用程序提供服务，这样我们才能分清楚哪些服务是操作系统提供的，而哪些服务是由我们所使用的语言的运行库提供的。

除了内存管理、文件管理、进程管理、外设管理等等内部模块以外，操作系统还提供了许多外部接口供应用程序使用，这些接口就是所谓的“系统调用”。从DOS时代开始，系统调用就是通过软中断的形式来提供，也就是著名的INT 21，程序把需要调用的功能编号放入AH寄存器，把参数放入其他指定的寄存器，然后调用INT 21，中断返回后，程序从指定的寄存器(通常是AL)里取得返回值。这样的做法一直到奔腾2也就是P6出来之前都没有变，譬如windows通过INT 2E提供系统调用，Linux则是INT 80，只不过后来的寄存器比以前大一些，而且可能再多一层跳转表查询。后来，Intel和AMD分别提供了效率更高的SYSENTER/SYSEXIT和SYSCALL/SYSRET指令来代替之前的中断方式，略过了耗时的特权级别检查以及寄存器压栈出栈的操作，直接完成从RING 3代码段到RING 0代码段的转换。

系统调用都提供什么功能呢？用操作系统的名字加上对应的中断编号到谷歌上一查就可以得到完整的列表 (Windows, Linux)，这个列表就是操作系统和应用程序之间沟通的协议，如果需要超出此协议的功能，我们就只能在自己的代码里去实现，譬如，对于内存管理，操作系统只提供进程级别的内存段的管理，譬如Windows的virtualmemory系列，或是Linux的brk，操作系统不会去在乎应用程序如何为新建对象分配内存，或是如何做垃圾回收，这些都需要应用程序自己去实现。如果超出此协议的功能无法自己实现，那我们就说该操作系统不支持该功能，举个例子，Linux在2.6之前是不支持多线程的，无论如何在程序里模拟，我们都无法做出多个可以同时运行的并符合POSIX 1003.1c语义标准的调度单元。

可是，我们写程序并不需要去调用中断或是SYSCALL指令，这是因为操作系统提供了一层封装，在Windows上，它是NTDLL.DLL，也就是常说的Native API，我们不但不需要去直接调用INT 2E或SYSCALL，准确的说，我们不能直接去调用INT 2E或SYSCALL，因为Windows并没有公开其调用规范，直接使用INT 2E或SYSCALL无法保证未来的兼容性。在Linux上则没有这个问题，系统调用的列表都是公开的，而且Linus非常看重兼容性，不会去做任何更改，glibc里甚至专门提供了syscall(2)来方便用户直接用编号调用，不过，为了解决glibc和内核之间不同版本兼容性带来的麻烦，以及为了提高某些调用的效率(譬如__NR_ gettimeofday)，Linux上还是对部分系统调用做了一层封装，就是VDSO (早期叫linux-gate.so)。

可是，我们写程序也很少直接调用NTDLL或者VDSO，而是通过更上一层的封装，这一层处理了参数准备和返回值格式转换、以及出错处理和错误代码转换，这就是我们所使用语言的运行库，对于C语言，Linux上是glibc，Windows上是kernel32(或调用msvcrt)，对于其他语言，譬如Java，则是JRE，这些“其他语言”的运行库通常最终还是调用glibc或kernel32。

"运行库"这个词其实不止包括用于和编译后的目标执行程序进行链接的库文件，也包括了脚本语言或字节码解释型语言的运行环境，譬如Python，C#的CLR，Java的JRE。

对系统调用的封装只是运行库的很小一部分功能，运行库通常还提供了诸如字符串处理、数学计算、常用数据结构容器等等不需要操作系统支持的功能，同时，运行库也会对操作系统支持的功能提供更易用更高级的封装，譬如带缓存和格式的IO、线程池。

所以，在我们说"某某语言新增了某某功能"的时候，通常是这么几种可能：

1. 支持新的语义或语法，从而便于我们描述和解决问题。譬如Java的泛型、Annotation、lambda表达式。
2. 提供了新的工具或类库，减少了我们开发的代码量。譬如Python 2.7的argparse.
3. 对系统调用有了更良好更全面的封装，使我们可以做到以前在这个语言环境里做不到或很难做到的事情。譬如Java NIO。

但任何一门语言，包括其运行库和运行环境，都不可能创造出操作系统不支持的功能，Go语言也是这样，不管它的特性描述看起来多么炫丽，那必然都是其他语言也可以做到的，只不过Go提供了更方便更清晰的语义和支持，提高了开发的效率。

二、并发与并行 (Concurrency and Parallelism)

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


三、 线程的调度

上一节主要说的是并发和并行的概念，而线程是最直观的并发的实现，这一节我们主要说操作系统如何让多个线程并发的执行，当然在多CPU的时候，也就是并行的执行。我们不讨论进程，进程的意义是“隔离的执行环境”，而不是“单独的执行序列”。

我们首先需要理解IA-32 CPU的指令控制方式，这样才能理解如何在多个指令序列(也就是逻辑控制流)之间进行切换。CPU通过CS:EIP寄存器的值确定下一条指令的位置，但是CPU并不允许直接使用MOV指令来更改EIP的值，必须通过JMP系列指令、CALL/RET指令、或INT中断指令来实现代码的跳转；在指令序列间切换的时候，除了更改EIP之外，我们还要保证代码可能会使用到的各个寄存器的值，尤其是栈指针SS:ESP，以及EFLAGS标志位等，都能够恢复到目标指令序列上次执行到这个位置时候的状态。

线程是操作系统对外提供的服务，应用程序可以通过系统调用让操作系统启动线程，并负责随后的线程调度和切换。我们先考虑单颗单核CPU，操作系统内核与应用程序其实是也是在共享同一个CPU，当EIP在应用程序代码段的时候，内核并没有控制权，内核并不是一个进程或线程，内核只是以实模式运行的，代码段权限为RING 0的内存中的程序，只有当产生中断或是应用程序呼叫系统调用的时候，控制权才转移到内核，在内核里，所有代码都在同一个地址空间，为了给不同的线程提供服务，内核会为每一个线程建立一个内核堆栈，这是线程切换的关键。通常，内核会在时钟中断里或系统调用返回前(考虑到性能，通常是在不频繁发生的系统调用返回前)，对整个系统的线程进行调度，计算当前线程的剩余时间片，如果需要切换，就在“可运行”的线程队列里计算优先级，选出目标线程后，则保存当前线程的运行环境，并恢复目标线程的运行环境，其中最重要的，就是切换堆栈指针ESP，然后再把EIP指向目标线程上次被移出CPU时的指令。Linux内核在实现线程切换时，耍了个花枪，它并不是直接JMP，而是先把ESP切换为目标线程的内核栈，把目标线程的代码地址压栈，然后JMP到__switch_to()，相当于伪造了一个CALL __switch_to()指令，然后，在__switch_to()的最后使用RET指令返回，这样就把栈里的目标线程的代码地址放入了EIP，接下来CPU就开始执行目标线程的代码了，其实也就是上次停在switch_to这个宏展开的地方。

这里需要补充几点：(1) 虽然IA-32提供了TSS (Task State Segment)，试图简化操作系统进行线程调度的流程，但由于其效率低下，而且并不是通用标准，不利于移植，所以主流操作系统都没有去利用TSS。更严格的说，其实还是用了TSS，因为只有通过TSS才能把堆栈切换到内核堆栈指针SS0:ESP0，但除此之外的TSS的功能就完全没有被使用了。(2) 线程从用户态进入内核的时候，相关的寄存器以及用户态代码段的EIP已经保存了一次，所以，在上面所说的内核态线程切换时，需要保存和恢复的内容并不多。(3) 以上描述的都是抢占式(preemptively)的调度方式，内核以及其中的硬件驱动也会在等待外部资源可用的时候主动调用schedule()，用户态的代码也可以通过sched_yield()系统调用主动发起调度，让出CPU。

现在我们一台普通的PC或服务里通常都有多颗CPU (physical package)，每颗CPU又有多个核 (processor core)，每个核又可以支持超线程 (two logical processors for each core)，也就是逻辑处理器。每个逻辑处理器都有自己的一套完整的寄存器，其中包括了CS:EIP和SS:ESP，从而，以操作系统和应用的角度来看，每个逻辑处理器都是一个单独的流水线。在多处理器的情况下，线程切换的原理和流程其实和单处理器时是基本一致的，内核代码只有一份，当某个CPU上发生时钟中断或是系统调用时，该CPU的CS:EIP和控制权又回到了内核，内核根据调度策略的结果进行线程切换。但在这个时候，如果我们的程序用线程实现了并发，那么操作系统可以使我们的程序在多个CPU上实现并行。

这里也需要补充两点：(1) 多核的场景里，各个核之间并不是完全对等的，譬如在同一个核上的两个超线程是共享L1/L2缓存的；在有NUMA支持的场景里，每个核访问内存不同区域的延迟是不一样的；所以，多核场景里的线程调度又引入了“调度域”(scheduling domains)的概念，但这不影响我们理解线程切换机制。(2) 多核的场景下，中断发给哪个CPU？软中断(包括除以0，缺页异常，INT指令)自然是在触发该中断的CPU上产生，而硬中断则又分两种情况，一种是每个CPU自己产生的中断，譬如时钟，这是每个CPU处理自己的，还有一种是外部中断，譬如IO，可以通过APIC来指定其送给哪个CPU；因为调度程序只能控制当前的CPU，所以，如果IO中断没有进行均匀的分配的话，那么和IO相关的线程就只能在某些CPU上运行，导致CPU负载不均，进而影响整个系统的效率。

四、并发编程框架

以上大概介绍了一个用多线程来实现并发的程序是如何被操作系统调度以及并行执行(在有多个逻辑处理器时)，同时大家也可以看到，代码片段或者说逻辑控制流的调度和切换其实并不神秘，理论上，我们也可以不依赖操作系统和其提供的线程，在自己程序的代码段里定义多个片段，然后在我们自己程序里对其进行调度和切换。

为了描述方便，我们接下来把"代码片段"称为"任务"。

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

五、goroutine

Go语言通过goroutine提供了目前为止所有(我所了解的)语言里对于并发编程的最清晰最直接的支持，Go语言的文档里对其特性也描述的非常全面甚至超过了，在这里，基于我们上面的系统知识介绍，列举一下goroutine的特性，算是小结：

- (1) goroutine是Go语言运行库的功能，不是操作系统提供的功能，goroutine不是用线程实现的。具体可参见Go语言源码里的pkg/runtime/proc.c
- (2) goroutine就是一段代码，一个函数入口，以及在堆上为其分配的一个堆栈。所以它非常廉价，我们可以很轻松的创建上万个goroutine，但它们并不是被操作系统所调度执行
- (3) 除了被系统调用阻塞的线程外，Go运行库最多会启动$GOMAXPROCS个线程来运行goroutine
- (4) goroutine是协作式调度的，如果goroutine会执行很长时间，而且不是通过等待读取或写入channel的数据来同步的话，就需要主动调用Gosched()来让出CPU
- (5) 和所有其他并发框架里的协程一样，goroutine里所谓“无锁”的优点只在单线程下有效，如果$GOMAXPROCS > 1并且协程间需要通信，Go运行库会负责加锁保护数据，这也是为什么sieve.go这样的例子在多CPU多线程时反而更慢的原因
- (6) Web等服务端程序要处理的请求从本质上来讲是并行处理的问题，每个请求基本独立，互不依赖，几乎没有数据交互，这不是一个并发编程的模型，而并发编程框架只是解决了其语义表述的复杂性，并不是从根本上提高处理的效率，也许是并发连接和并发编程的英文都是concurrent吧，很容易产生“并发编程框架和coroutine可以高效处理大量并发连接”的误解。
- (7) Go语言运行库封装了异步IO，所以可以写出貌似并发数很多的服务端，可即使我们通过调整$GOMAXPROCS来充分利用多核CPU并行处理，其效率也不如我们利用IO事件驱动设计的、按照事务类型划分好合适比例的线程池。在响应时间上，协作式调度是硬伤。
- (8) goroutine最大的价值是其实现了并发协程和实际并行执行的线程的映射以及动态扩展，随着其运行库的不断发展和完善，其性能一定会越来越好，尤其是在CPU核数越来越多的未来，终有一天我们会为了代码的简洁和可维护性而放弃那一点点性能的差别。
###Golang图片处理与base64格式转换
 1. 图片文件的读写。 2. 图片在go缓存中如何与base64互相转换 3. 图片裁剪

本文中，为了方便查看，去掉所有错误判断
base64 -> file
<pre>
ddd, _ := base64.StdEncoding.DecodeString(datasource) //成图片文件并把文件写入到buffer
err2 := ioutil.WriteFile("./output.jpg", ddd, 0666)   //buffer输出到jpg文件中（不做处理，直接写到文件）
datasource base64 string
</pre>
base64 -> buffer
<pre>
ddd, _ := base64.StdEncoding.DecodeString(datasource) //成图片文件并把文件写入到buffer
bbb := bytes.NewBuffer(ddd)                           // 必须加一个buffer 不然没有read方法就会报错
转换成buffer之后里面就有Reader方法了。才能被图片API decode
</pre>
buffer-> ImageBuff（图片裁剪,代码接上面）
<pre>
m, _, _ := image.Decode(bbb)                                       // 图片文件解码
rgbImg := m.(*image.YCbCr)
subImg := rgbImg.SubImage(image.Rect(0, 0, 200, 200)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
</pre>
img -> file(代码接上面)
<pre>
f, _ := os.Create("test.jpg")     //创建文件
defer f.Close()                   //关闭文件
jpeg.Encode(f, subImg, nil)       //写入文件
</pre>
img -> base64(代码接上面)
<pre>
emptyBuff := bytes.NewBuffer(nil)                  //开辟一个新的空buff
jpeg.Encode(emptyBuff, subImg, nil)                //img写入到buff
dist := make([]byte, 50000)                        //开辟存储空间
base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
fmt.Println(string(dist))                          //输出图片base64(type = []byte)
_ = ioutil.WriteFile("./base64pic.txt", dist, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
</pre>
imgFile -> base64
<pre>
ff, _ := ioutil.ReadFile("output2.jpg")               //我还是喜欢用这个快速读文件
bufstore := make([]byte, 5000000)                     //数据缓存
base64.StdEncoding.Encode(bufstore, ff)               // 文件转base64
_ = ioutil.WriteFile("./output2.jpg.txt", dist, 0666) //直接写入到文件就ok完活了。
</pre>
大概就是这些代码基本上一些小网站都够用。 缩放什么的可以先靠前端。后端有个裁剪就够了。


###Golang GC调优总结经验
总结为以下几点：

- 尽早的用memprof、cpuprof、GCTRACE来观察程序。
- 关注请求处理时间，特别是开发新功能的时候，有助于发现设计上的问题。
- 尽量避免频繁创建对象(&abc{}、new(abc{})、make())，在频繁调用的地方可以做对象重用。
- 尽量不要用go管理大量对象，内存数据库可以完全用c实现好通过cgo来调用。

###Golang健全的代码风格与检查工具
go fmt这个命令，统一了代码风格，带来的影响就是别人写的代码感觉也是自己写的一样。还有golint，可以按照go team的风格和要求来写代码。还有go vet可以用来检查一些在GO中很隐蔽的坑。

###Go语言编程的并发安全
Go的内存模型对于并发安全有两种保护措施。 一种是通过加锁来保护，另一种是通过channel（只有channel是并发安全的）来保护。前者没什么好说的，后者其实就是一个线程安全的队列。 

可能很多人都听说过一个高逼格的词叫【无锁队列】。都一听到加锁就觉得很low，其实对于大部分程序来说。 根本不需要那些高逼格的技术，该加锁就加锁。一次加锁的耗时差不多是在几十纳秒， 而一次网络IO都是在毫秒级别以上的。它们根本不是一个量级。特别是在现在云计算时代，大部分人一辈子都遇不到因为加锁成为性能瓶颈的应用场景。

Go语言编程中， 当有多个goroutine并发操作同一个变量时，除非是全都是只读操作， 否则就得【加锁】或者【使用channel】来保证并发安全。不要觉得加锁麻烦，但是它能保证并发安全。

###获取Golang版本
<pre>
package main

import (
	"runtime"
	"fmt"
)
func main(){
	fmt.Println("Golang版本:",runtime.Version())
}
</pre>
###Golang循环for（From yushuangqi.com）
for 是 Go 中唯一的循环结构。这里有 for 循环的三个基本使用方式。最常用的方式，带单个循环条件。经典的初始化/条件/后续形式 for 循环。不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环。
<pre>
// `for` 是 Go 中唯一的循环结构。这里有 `for` 循环
// 的三个基本使用方式。

package main

import "fmt"

func main() {

	// 最常用的方式，带单个循环条件。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始化/条件/后续形式 `for` 循环。
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带条件的 `for` 循环将一直执行，直到在循环体内使用
	// 了 `break` 或者 `return` 来跳出循环。
	for {
		fmt.Println("loop")
		break
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
</pre>
###Golang 数组 (From  yushuangqi.com)
在 Go 中，数组 是一个固定长度的数列。

这里我们创建了一个数组 a 来存放刚好 5 个 int。元素的类型和长度都是数组类型的一部分。数组默认是零值的，对于 int 数组来说也就是 0。

我们可以使用 array[index] = value 语法来设置数组指定位置的值，或者用 array[index] 得到值。

使用内置函数 len 返回数组的长度

使用这个语法在一行内初始化一个数组

数组的存储类型是单一的，但是你可以组合这些数据来构造多维的数据结构。
<pre>
// 在 Go 中，_数组_ 是一个固定长度的数列。

package main

import "fmt"

func main() {

	// 这里我们创建了一个数组 `a` 来存放刚好 5 个 `int`。
	// 元素的类型和长度都是数组类型的一部分。数组默认是
	// 零值的，对于 `int` 数组来说也就是 `0`。
	var a [5]int
	fmt.Println("emp:", a)

	// 我们可以使用 `array[index] = value` 语法来设置数组
	// 指定位置的值，或者用 `array[index]` 得到值。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 使用内置函数 `len` 返回数组的长度
	fmt.Println("len:", len(a))

	// 使用这个语法在一行内初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 数组的存储类型是单一的，但是你可以组合这些数据
	// 来构造多维的数据结构。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
output==>
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]
</pre>
###Golang 切片 (From  yushuangqi.com)
使用append时，如果slice的容量(即对应array的长度)不够，go会创建一个新的array(长度通常为之前的两倍)以容纳新添加的数据，所有旧的array数据都会被拷贝到新的array里。需要频繁使用append时，需要考虑到其效率问题。
<pre>
// _Slice_ 是 Go 中一个关键的数据类型，是一个比数组更
// 加强大的序列接口

package main

import "fmt"

func main() {

	// 不想数组，slice 的类型仅有它所包含的元素决定（不像
	// 数组中还需要元素的个数）。要创建一个长度非零的空
	// slice，需要使用内建的方法 `make`。这里我们创建了一
	// 个长度为3的 `string` 类型 slice（初始化为零值）。
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 我们可以和数组一下设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 如你所料，`len` 返回 slice 的长度
	fmt.Println("len:", len(s))

	// 作为基本操作的补充，slice 支持比数组更多的操作。
	// 其中一个是内建的 `append`，它返回一个包含了一个
	// 或者多个新值的 slice。注意我们接受返回由 append
	// 返回的新的 slice 值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice 也可以被 `copy`。这里我们创建一个空的和 `s` 有
	// 相同长度的 slice `c`，并且将 `s` 复制给 `c`。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice 支持通过 `slice[low:high]` 语法进行“切片”操
	// 作。例如，这里得到一个包含元素 `s[2]`, `s[3]`,
	// `s[4]` 的 slice。
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这个 slice 从 `s[0]` 到（但是包含）`s[5]`。
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个 slice 从（包含）`s[2]` 到 slice 的后一个值。
	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们可以在一行代码中申明并初始化一个 slice 变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不
	// 同，这和多位数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
output=>
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]
</pre>
Golang中的slice与C++中的vector十分类型。下面是vector的简单认识。

vector是STL中最常见的容器，它是一种顺序容器，支持随机访问。vector是一块连续分配的内存，从数据安排的角度来讲，和数组极其相似，不同的地方就是：数组是静态分配空间，一旦分配了空间的大小，就不可再改变了；而vector是动态分配空间，随着元素的不断插入，它会按照自身的一套机制不断扩充自身的容量。

vector的扩充机制：按照容器现在容量的一倍进行增长。vector容器分配的是一块连续的内存空间，每次容器的增长，并不是在原有连续的内存空间后再进行简单的叠加，而是重新申请一块更大的新内存，并把现有容器中的元素逐个复制过去，然后销毁旧的内存。这时原有指向旧内存空间的迭代器已经失效，所以当操作容器时，迭代器要及时更新。
###Golang 字典(From yushuangqi.com)
map 是 Go 内置关联数据类型（在一些其他的语言中称为哈希 或者字典 ）。

要创建一个空 map，需要使用内建的 make:make(map[key-type]val-type).

使用典型的 make[key] = val 语法来设置键值对。

使用例如 Println 来打印一个 map 将会输出所有的键值对。

使用 name[key] 来获取一个键的值

当对一个 map 调用内建的 len 时，返回的是键值对数目

内建的 delete 可以从一个 map 中移除键值对

当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义。

你也可以通过这个语法在同一行申明和初始化一个新的map。
<pre>
package main

import (
	"fmt"
)
func main(){
	m :=make(map[string]int)
	m["k1"] = 5
	m["k2"] = 4
	fmt.Println("map",m)
	v :=m["k2"]
	fmt.Println(v)
	delete(m,"k2")
	v2 := m["k2"]
	fmt.Println(v)
	fmt.Println(v2)
	value,isexist :=m["k1"]
	fmt.Println(value,isexist)
}
output==>
map map[k1:5 k2:4]
4
4
0
5 true
</pre>
###Golang 可变参数函数（From yushuangqi.com）
可变参数函数。可以用任意数量的参数调用。例如，fmt.Println 是一个常见的变参函数。

这个函数使用任意数目的 int 作为参数。

变参函数使用常规的调用方式，除了参数比较特殊。

如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)。
<pre>
package main

import (
	"fmt"
)
func sum(nums ...int){
	fmt.Println(nums," ")
	total := 0
	for _,num := range nums{
		total +=num
	}
	fmt.Println(total)
}
func main(){
	sum(4,5,6,7,8,5)
	nums := []int{3,6,5,44,5,3}
	sum(nums...)
}
output==>
[4 5 6 7 8 5]  
35
[3 6 5 44 5 3]  
66
</pre>
###Golang 闭包（From yushuangqi.com）
Go 支持通过闭包来使用匿名函数。匿名函数在你想定义一个不需要命名的内联函数时是很实用的。

这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。这个返回的函数使用闭包的方式隐藏变量 i。

我们调用 intSeq 函数，将返回值（也是一个函数）赋给nextInt。这个函数的值包含了自己的值 i，这样在每次调用 nextInt 是都会更新 i 的值。

通过多次调用 nextInt 来看看闭包的效果。
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
	nextInts := intSeq()
	fmt.Println(nextInts())
}
output==>
1
2
1
</pre>
Golang中的struct和数组一样，也是值类型。

Channel与锁都可以实现并发安全，但是Channel和锁谁轻量? 一句话告诉你: Channel本身用锁实现的. 因此在迫不得已时, 还是尽量减少使用Channel, 但Channel属于语言层支持, 适度使用, 可以改善代码可读写
###Golang 原子计数器atomic(From yushuangqi.com)
Go 中最主要的状态管理方式是通过通道间的沟通来完成的，我们在工作池的例子中碰到过，但是还是有一些其他的方法来管理状态的。这里我们将看看如何使用 sync/atomic包在多个 Go 协程中进行 原子计数 。

我们将使用一个无符号整形数来表示（永远是正整数）这个计数器。

为了模拟并发更新，我们启动 50 个 Go 协程，对计数器每隔 1ms （译者注：应为非准确时间）进行一次加一操作。

使用 AddUint64 来让计数器自动增加，使用& 语法来给出 ops 的内存地址。

允许其它 Go 协程的执行

等待一秒，让 ops 的自加操作执行一会。

为了在计数器还在被其它 Go 协程更新时，安全的使用它，我们通过 LoadUint64 将当前值得拷贝提取到 opsFinal中。和上面一样，我们需要给这个函数所取值的内存地址 &ops
<pre>
// Go 中最主要的状态管理方式是通过通道间的沟通来完成的，我们
// 在[工作池](../worker-pools/)的例子中碰到过，但是还是有一
// 些其他的方法来管理状态的。这里我们将看看如何使用 `sync/atomic`
// 包在多个 Go 协程中进行 _原子计数_ 。

package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

	// 我们将使用一个无符号整形数来表示（永远是正整数）这个计数器。
	var ops uint64 = 0

	// 为了模拟并发更新，我们启动 50 个 Go 协程，对计数
	// 器每隔 1ms （译者注：应为非准确时间）进行一次加一操作。
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用 `AddUint64` 来让计数器自动增加，使用
				// `&` 语法来给出 `ops` 的内存地址。
				atomic.AddUint64(&ops, 1)

				// 允许其它 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}

	// 等待一秒，让 ops 的自加操作执行一会。
	time.Sleep(time.Second)

	// 为了在计数器还在被其它 Go 协程更新时，安全的使用它，
	// 我们通过 `LoadUint64` 将当前值得拷贝提取到 `opsFinal`
	// 中。和上面一样，我们需要给这个函数所取值的内存地址 `&ops`
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
output==>
ops: 1354264
</pre>
###Golang URL解析（From yushuangqi.com）
<pre>
// URL 提供了一个[统一资源定位方式](http://adam.heroku.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/)。
// 这里了解一下 Go 中是如何解析 URL 的。

package main

import "fmt"
import "net/url"
import "strings"

func main() {

	// 我们将解析这个 URL 示例，它包含了一个 scheme，
	// 认证信息，主机名，端口，路径，查询参数和片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 直接访问 scheme。
	fmt.Println(u.Scheme)

	// `User` 包含了所有的认证信息，这里调用 `Username`
	// 和 `Password` 来获取独立值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` 同时包括主机名和端口信息，如过端口存在的话，
	// 使用 `strings.Split()` 从 `Host` 中手动提取端口。
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	// 这里我们提出路径和查询片段信息。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 要得到字符串中的 `k=v` 这种格式的查询参数，可以使
	// 用 `RawQuery` 函数。你也可以将查询参数解析为一个
	// map。已解析的查询参数 map 以查询字符串为键，对应
	// 值字符串切片为值，所以如何只想得到一个键对应的第
	// 一个值，将索引位置设置为 `[0]` 就行了。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
output==>
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v
</pre>
###Golang 退出 （From yushuangqi.com）
<pre>
// 使用 `os.Exit` 来立即进行带给定状态的退出。

package main

import "fmt"
import "os"

func main() {

	// 当使用 `os.Exit` 时 `defer` 将_不会_ 执行，所以这里的 `fmt.Println`
	// 将永远不会被调用。
	defer fmt.Println("!")

	// 退出并且退出状态为 3。
	os.Exit(3)
}

// 注意，不像例如 C 语言，Go 不使用在 `main` 中返回一个整
// 数来指明退出状态。如果你想以非零状态退出，那么你就要
// 使用 `os.Exit`。
</pre>
###Golang SHA1散列（From yushuangqi.com）
SHA1 散列经常用生成二进制文件或者文本块的短标识。例如，git 版本控制系统大量的使用 SHA1 来标识受版本控制的文件和目录。这里是 Go中如何进行 SHA1 散列计算的例子。

Go 在多个 crypto/* 包中实现了一系列散列函数。

产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。

写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。

这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。

SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
<pre>
// [_SHA1 散列_](http://en.wikipedia.org/wiki/SHA-1)经常用
// 生成二进制文件或者文本块的短标识。例如，[git 版本控制系统](http://git-scm.com/)
// 大量的使用 SHA1 来标识受版本控制的文件和目录。这里是 Go
// 中如何进行 SHA1 散列计算的例子。

package main

// Go 在多个 `crypto/*` 包中实现了一系列散列函数。
import "crypto/sha1"
import "fmt"

func main() {
	s := " string"

	// 产生一个散列值得方式是 `sha1.New()`，`sha1.Write(bytes)`，
	// 然后 `sha1.Sum([]byte{})`。这里我们从一个新的散列开始。
	h := sha1.New()

	// 写入要处理的字节。如果是一个字符串，需要使用
	// `[]byte(s)` 来强制转换成字节数组。
	h.Write([]byte(s))

	// 这个用来得到最终的散列值的字符切片。`Sum` 的参数可以
	// 用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)

	// SHA1 值经常以 16 进制输出，例如在 git commit 中。使用
	// `%x` 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
output==>
 string
afa6a745c81b4be08a6681769fc3ff2a1f03e9fe
</pre>
###Golang对结构体中单个数据加锁
下面是给结构体x中的a加锁
<pre>
type x struct {
    a int
    lock_a sync.Mutex
    b int
}

func (p *x) Geta() int {
    p.lock_a.Lock()
    defer p.lock_a.Unlock()
    return p.a
}
</pre>
###GO在一个没有返回值的方法中增加slice的长度,并且能根据下标进行修改
<pre>
package main

import (
    "fmt"
)

func main() {
    b := make([]string, 0, 1000)
    apppppend(&b)
    fmt.Println(b)
}

func apppppend(b *[]string) {
    *b = append(*b, "first_element")  // 可以进行append
    *b = append(*b, "second_element") // 可以进行append
    (*b)[0] = "1st_element"           // []的优先级比*高,所以通过下标修改必须这样写

    // 但是总的来说..不能有返回值的函数在golang里面实在是太奇葩了
    // 建议别这样
}
output==>
[1st_element second_element]
</pre>

##让 go get 显示进度
看了下golang的源码 src/cmd/go 下是go命令的源码, 其中, get.go是go get命令的代码, build.go 是go build的代码.

刚开始走了点弯路, 想着改变get.go来显示进度, 无果之后想了下, go get 其实就是调用git , hg, svn的命令从仓库中下载的, 由此思路找到vcs.go(vcs全称为version control system), 果然这里面包含了调用git, hg, svn的命令. 问题迎刃而解:

- 修改git clone命令, 添加 --progress选项, 使其输出进度
- 修改cmd.Run()执行的地方, 使其将输出定位到标准输出流上

一、 修改git clone命令, 找到如下代码, 在createdCmd修改为 clone --progress {repo} {dir}

其它命令hg,svn...添加进度方法类似
<pre>
// vcsGit describes how to use Git.
var vcsGit = &vcsCmd{
	name: "Git",
	cmd:  "git",

	createCmd:   "clone {repo} {dir}", // 此处修改为 clone --progress {repo} {dir}
	downloadCmd: "pull --ff-only"
}
</pre>
二、 重定向输出流

找到run1()方法, 在 cmd.Stderr = &buf 下添加两行, 如:
<pre>
var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	cmd.Stdout = os.Stdout // 重定向标准输出
	cmd.Stderr = os.Stderr // 重定向标准输出
	err = cmd.Run()
</pre>
Ok,搞定,接下来执行golang源码src下的all.bash重新编译golang,编译要些时间, 编译完后使用go get试试。

###Golang热更新
首先强类型的golang自己没有从语言层面支持热更新，也就是说大家可以理解为golang自身不支持热更新。不过有第三方的库让golang支持热更新，比如：https://github.com/rcrowley/goagain与https://github.com/facebookgo/grace，这两个都是star在1k以上的，可用性稳定性应该不错（自己还没有尝试使用过-.-）。当然还有人提出使用C的方式来支持热更新。具体是通过编译成so共享库文件（so 为共享库,是shared object,用于Linux下动态链接文件,和window下动态链接库文件dll差不多。特点：ELF格式文件，共享库（动态库），类似于DLL。节约资源，加快速度，代码升级简化），例如：

主程序：
<pre>
package main
/*
#include <dlfcn.h>
#cgo LDFLAGS: -ldl

void (*foo)(int);

void set_foo(void* fn) {
	foo = fn;
}

void call_foo(int i) {
	foo(i);
}
*/
import "C"
import "fmt"

func main() {
	n := 0
	var bar string
	for {
		hd := C.dlopen(C.CString(fmt.Sprintf("./foo-%d.so", n)), C.RTLD_LAZY)
		if hd == nil {
			panic("dlopen")
		}
		foo := C.dlsym(hd, C.CString("foo"))
		if foo == nil {
			panic("dlsym")
		}
		fmt.Printf("%v\n", foo)
		C.set_foo(foo)
		C.call_foo(42)
		fmt.Scanf("%s", &bar)
		n++
	}
}
</pre>
so源码：
<pre>
package main

import "fmt"
import "C"

func main() {
}

//export foo
func foo(i C.int) {
	fmt.Printf("%d-2\n", i)
}
</pre>
用 go build -buildmode=c-shared -o foo-1.so mod.go 编译。需要golang编译器版本>=1.5。这是借助C的机制来实现的，go的execution modes文档提到会有go原生的plugin模式。不过这种的可行性有待考究。

还有一种做法是可以将服务端微服务化；这样每个服务的重启成本很低，reload数据库到内存的时间成本就会更低。另外为服务化后，可以针对不同的服务是否有必要热更新，结合脚本或其它方法实现（如:游戏运维的活动服务需要频繁变更）。而像一些基本的如用户，游戏逻辑等接口设计灵活一点的情况下；是完全没必要热更新的；每次版本变更停服重启就ok。

nginx是支持热升级的，可以用老进程服务先前链接的链接，使用新进程服务新的链接，即在不停止服务的情况下完成系统的升级与运行参数修改。那么热升级和热编译是不同的概念，热编译是通过监控文件的变化重新编译，然后重启进程。那么也可以用golang模仿nginx的方式来实现热更新。

根据以上的思路，谢大总结出了一套他自己实现beego热更新的方法,思路如下：
<pre>
 热升级的原理基本上就是：主进程fork一个进程，然后子进程exec相应的程序。那么这个过程中发生了什么呢？我们知道进程fork之后会把主进程的所有句柄、数据和堆栈、但是里面所有的句柄存在一个叫做CloseOnExec，也就是执行exec的时候，copy的所有的句柄都被关闭了，除非特别申明，而我们期望的是子进程能够复用主进程的net.Listener的句柄。一个进程一旦调用exec类函数，它本身就"死亡"了，系统把代码段替换成新的程序的代码，废弃原有的数据段和堆栈段，并为新程序分配新的数据段与堆栈段，唯一留下的，就是进程号，也就是说，对系统而言，还是同一个进程，不过已经是另一个程序了。

那么我们要做的：
第一步就是让子进程继承主进程的这个句柄，我们可以通过os.StartProcess的参数来附加Files，把需要继承的句柄写在里面。

第二步就是我们希望子进程能够从这个句柄启动监听，还好Go里面支持net.FileListener，直接从句柄来监听，但是我们需要子进程知道这个FD，所以在启动子进程的时候我们设置了一个环境变量设置这个FD。

第三步就是我们期望老的链接继续服务完，而新的链接采用新的进程，这里面有两个细节，第一就是老的链接继续服务，那么我们怎么有老链接存在？所以我们必须每次接收一个链接记录一下，这样我们就知道还存在没有服务完的链接，第二就是怎么让老进程停止接收数据，让新进程接收数据呢？大家都监听在同一个端口，理论上是随机来接收的，所以这里我们只要关闭老的链接的接收就行，这样就会使得在l.Accept的时候报错。

演示地址：http://my.oschina.net/astaxie/blog/136364 
</pre>
到底golang是不是一定要热更新功能，最后用达达来观点来总结一下。

没有热更新的确没那么方便，但是也没那么可怕。

原因：

1. 需要临时重启更新就运营公告，如果实际较长就适当发放补偿。
2. Go加载数据到内存的速度也比之前快很多，重启压力也没想象的那么大。
3. 强类型语法在编译器提前排除了很多之前要到线上运行时才能发现的问题，所以BUG率低了。

所以没有热更新也顺利跑下来了。

不过以上只能做为参考，不同项目需求不一样，还是得结合实际情况来判断。

热更新肯定是可以做的，方案挺多，数据驱动、内嵌脚本或者无状态进程都可行，只是花多大代价换多少回报的问题。

如果评估下来觉得热更新必做不可，那么用再大代价也得做，这是项目存亡问题。


如果不是必须的，那就需要评估性价比了。

做热更新、换编程语言或者换服务端架构所花的代价，换来的产品在运营、运维或开发各方面的效率提升，是否划算。 

参考链接：https://www.zhihu.com/question/31912663/answer/53872820

###侵入式与非侵入式理解
假设大家都想要把用户代码塞到一个框架里。侵入式的做法就是要求用户代码“知道”框架的代码，表现为用户代码需要继承框架提供的类。非侵入式则不需要用户代码引入框架代码的信息，从类的编写者角度来看，察觉不到框架的存在。

- 侵入式让用户代码产生对框架的依赖，这些代码不能在框架外使用，不利于代码的复用。但侵入式可以使用户跟框架更好的结合，更容易更充分的利用框架提供的功能;
- 非侵入式的代码则没有过多的依赖，可以很方便的迁移到其他地方。但是与用户代码互动的方式可能就比较复杂。

###Golang压力测试工具
https://github.com/wg/wrk

###Golang GC优化
go没有像jvm那样多的可以调整的参数，并且不是分代回收。优化gc的方式仅仅只能是通过优化程序。但go有一个优势：有真正的array（而仅仅是an array of referece）。go的gc算法是mark and sweep，array对此是友好的：整个array一次性被处理。可以用一个array用open addressing的方式实现map，以此优化gc，也会减少内存的使用。
根据前面的知识，应对GC抖动的策略是，减少对象数，用海量array代替海量struct。

###获取golang goroutine的id
原理：

利用runtime.Stack的堆栈信息。runtime.Stack(buf []byte, all bool) int会
将当前的堆栈信息写入到一个slice中，堆栈的第一行为goroutine #### [.....
其中####就是当前的gororutine id。

需要注意的是，获取堆栈信息会影响性能，所以建议你在debug的时候才用它。

<pre>
package main
import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)
/*
利用runtime.Stack的堆栈信息。runtime.Stack(buf []byte, all bool) int会
将当前的堆栈信息写入到一个slice中，堆栈的第一行为goroutine #### [.....
其中####就是当前的gororutine id。
需要注意的是，获取堆栈信息会影响性能，所以建议你在debug的时候才用它。
*/
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
func main() {
	fmt.Println("main", GoID())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, "Goroutine ID:",GoID())
		}()
	}
	wg.Wait()
}
</pre>

##Golang实现一个速度快内存占用小的一致性哈希算法
哈希算法应该满足的4个适应条件：：Balance(均衡)、Monotonicity(单调性)、Spread(分散性)、Load(负载)。

在分布式缓存系统中使用一致性哈希算法时，某个节点的添加和移除不会重新分配全部的缓存，而只会影响小部分的缓存系统，如果均衡性做的好的话，当添加一个节点时，会均匀地从其它节点移一部分缓存到新的节点上；当删除一个节点的时候，这个节点上的缓存会均匀地分配到其它活着的节点上。

一致性哈希缓存还被扩展到分布式存储系统上。数据被分成一组Shard ,每个Shard由一个节点管理，当需要扩容时，我们可以添加新的节点，然后将其它Shard的一部分数据移动到这个节点上。比如我们有10个Shard的分布式存储系统，当前存储了120个数据，每个Shard存储了12个数据。当扩容成12个Shard时，我们从每个Shard上拿走2个数据，存入到新的两个Shard上,这样每个Shard都存储了10个数据，而整个过程中我们只移动了20/120=1/6的数据。

Karger 一致性哈希算法将每个节点(bucket)关联一个圆环上的一些随机点，对于一个键值，将其映射到圆环中的一个点上，然后按照顺时针方向找到第一个关联bucket的点，将值放入到这个bucke中。因此你需要存储一组bucket和它们的关联点，当bucket以及每个bucket的关联点很多的时候，你就需要多一点的内存来记录它。这个你经常在网上看到的介绍一致性哈希的算法(有些文章将节点均匀地分布在环上，比如节点1节点2节点3节点4节点1节点2节点3节点4……， 这是不对的，在这种情况下节点2挂掉后它上面的缓存全部转移给节点3了)。

其它的一致性算法还有Rendezvous hashing, 计算一个key应该放入到哪个bucket时，它使用哈希函数h(key,bucket)计算每个候选bucket的值，然后返回值最大的bucket。buckets比较多的时候耗时也较长，有人也提出了一些改进的方法，比如将bucket组织成tree的结构，但是在reblance的时候花费时间又长了。

Java程序员熟悉的Memcached的客户端Spymemcached、Xmemcached以及Folsom都提供了Ketama算法。其实Ketama算法最早于2007年用c 实现(libketama)，很多其它语言也实现了相同的算法，它是基于Karger 一致性哈希算法实现：

建立一组服务器的列表 (如: 1.2.3.4:11211, 5.6.7.8:11211, 9.8.7.6:11211)
为每个服务器节点计算一二百个哈希值
从概念上讲，这些数值被放入一个环上(continuum). (想象一个刻度为 0 到 2^32的时钟，这个时钟上就会散落着一些数字)
每一个数字关联一个服务器，所以服务器出现在这个环上的一些点上，它们是哈希分布的
为了找个一个Key应该放入哪个服务器，先哈希你的key，得到一个无符号整数, 沿着圆环找到和它相邻的最大的数，这个数对应的服务器就是被选择的服务器
对于靠近 2^32的 key, 因为没有超过它的数字点，按照圆环的原理，选择圆环中的第一个服务器。
以上两种算法可以处理节点增加和移除的情况。对于分布式存储系统，当一个节点失效时，我们并不期望它被移除，而是使用备份节点替换它，或者将它恢复起来，因为我们不期望丢掉它上面的数据。对于这种情况(节点可以扩容，但是不会移除节点)，Google的 John Lamping, Eric Veach提供一个高效的几乎不占用持久内存的算法：Jump Consistent Hash。

参考链接：http://www.udpwork.com/item/15346.html
<pre>
package main
import "fmt"
func JumpHash(key uint64, buckets int) int {
	var b, j int64	
	if buckets <= 0 {
			buckets = 1
	}
	for j < int64(buckets) {
			b = j		
			key = key*2862933555777941757 + 1
			j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}	
	return int(b)
}

func main() {
	buckets := make(map[int]int, 10)	
	count := 10	
	for i := uint64(0); i < 120000; i++ {	
		b := JumpHash(i, count)		
		buckets[b] = buckets[b] + 1
	}	
	fmt.Printf("buckets: %v\n", buckets)//add two buckets
	count = 12	
	for i := uint64(0); i < 120000; i++ {	
		oldBucket := JumpHash(i, count-2)	
		newBucket := JumpHash(i, count)//如果对象需要移动到新的bucket中,则首先从原来的bucket删除，再移动
		if oldBucket != newBucket {		
			buckets[oldBucket] = buckets[oldBucket] - 1			
			buckets[newBucket] = buckets[newBucket] + 1		
		}	
	}	
	fmt.Printf("buckets after add two servers: %v\n", buckets)
}
output==>
buckets: map[1:12001 7:12071 2:12012 3:11997 0:11992 6:11989 8:11908 4:12009 9:12054 5:11967]
buckets after add two servers: map[2:10024 3:10003 1:9997 7:10086 8:9950 4:10016 9:10028 5:9971 10:9973 11:9967 0:9998 6:9987]	
</pre>
因为Jump consistent hash算法不使用节点挂掉，如果你真的有这种需求，比如你要做一个缓存系统，你可以考虑使用ketama算法，或者对Jump consistent hash算法改造一下：节点挂掉时我们不移除节点，只是标记这个节点不可用。当选择节点时，如果选择的节点不可用，则再一次Hash，尝试选择另外一个节点，比如下面的算法将key加1再进行选择。
<pre>
/*
算法有一点问题，就是没有设定重试的测试，如果所有的节点都挂掉，则会进入死循环，所以最好设置一下重试次数(递归次数)，超过n次还没有选择到则返回失败。
*/
func JumpHash(key uint64, buckets int, checkAlive func(int) bool) int {
	var b, j int64 = -1, 0	
	if buckets <= 0 {
		buckets = 1
	}	
	for j < int64(buckets) {	
		b = j		
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))	
	}	
	if checkAlive != nil && !checkAlive(int(b)) {
		return JumpHash(key+1, buckets, checkAlive) //最好设置深度，避免key+1一直返回当掉的服务器
	}	
	return int(b)
}
</pre>
###Golang一段典型的请求超时退出代码
<pre>
func DoSomething() {
   done := make(chan error)
   go func() {
      done <- DoThing()
   }()

   select {
   case <-time.After(time.Second*10):
      Close()//中止DoThing()的执行，比如关闭网络连接
      return fmt.Errorf("Timeout")
   case err := <-done:
      if err != nil {
         err = fmt.Errorf("call failed, err %v", err)
      }
      return err
   }
}
</pre>
###Golang更新第三方包
将第三方包升级到最新版本，直接go get -u github.com/xxx/xxx。
###Golang一致性哈希库consistent
stathat.com/c/consistent是一个一致性哈希库。一致性哈希是为了解决在分布式系统中，数据存取时选择哪一个具体节点的问题。

比如，系统中有五个节点，大量用户信息分别存在不同的节点上，具体到某一个用户，其信息应该确定的存在一个节点上，不能两次请求，分别去不同的节点上取数据。最简单的思路，可以拿用户ID和节点数求余数，比如用户ID是 1、6、11、16的在第一个节点上，2、7、12、17的在第二个节点上，依此类推。

但是，如果系统中某一个节点坏掉了，变成4个了。如果再按4求余的话，会导致大量数据需要重新初始化。比如用户6，原来在第1个节点上，坏掉一个以后6%4=2，用户数据跑到第2个节点上去了。

如果系统中增加了新的节点，同样也会导致这个问题。
<pre>
package main

import (
	"fmt"
	_ "github.com/astaxie/beego"
	_ "github.com/huichen/sego"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gopkg.in/redis.v3"
	"stathat.com/c/consistent"
)

/*
func main() {
	db, err := gorm.Open("mysql", "root:123456@/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err.Error())
	}
} */
func main() {
	cons := consistent.New()
	cons.Add("cacheA")
	cons.Add("cacheB")
	cons.Add("cacheC")

	server1, err := cons.Get("user_1")
	server2, err := cons.Get("user_2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("server1:", server1) //输出 server1: cacheC
	fmt.Println("server2:", server2) //输出 server2: cacheA

	fmt.Println()

	//user_1在cacheA上，把cacheA删掉后看下效果
	cons.Remove("cacheA")
	server1, err = cons.Get("user_1")
	server2, err = cons.Get("user_2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("server1:", server1) //输出 server1: cacheC,和删除之前一样，在同一个server上
	fmt.Println("server2:", server2) //输出 server2: cacheB,换到另一个server了
}
output==>
server1: cacheC
server2: cacheA
	
server1: cacheC
server2: cacheB
</pre>
##Golang session Session
https://github.com/gorilla/sessions
##Golang删除slice中元素
<pre>
package main

import (
    "fmt"
)

//删除slice元素函数
func remove(s []string, i int) []string {
    return append(s[:i], s[i+1:]...)
}

func main() {
    s := []string{"a", "b", "c"}
    fmt.Println(s)
    s = remove(s, 1)
    fmt.Println(s)
}
</pre>
##Golang获取系统信息|MAC|IP|GOROOT|CPU|操作系统|架构|有线网络连接|无线网络连接
<pre>
package main

import (
    "fmt"
    "net"
    "runtime"
)

func main() {
    //操作系统
    fmt.Println("GOOS:", runtime.GOOS)
    //架构
    fmt.Println("GOARCH:", runtime.GOARCH)
    //GOROOT
    fmt.Println("GOROOT:", runtime.GOROOT())
    //go版本
    fmt.Println("Version:", runtime.Version())
    //cpu数
    fmt.Println("NumCPU:", runtime.NumCPU())
    fmt.Println()

    //MAC和IP地址
    interfaces, err := net.Interfaces()
    if err != nil {
        panic("Poor soul, here is what you got: " + err.Error())
    }
    for _, inter := range interfaces {
        fmt.Println(inter.Name, inter.HardwareAddr)

        addrs, _ := inter.Addrs()
        for _, addr := range addrs {
            fmt.Println("  ", addr.String())
        }
    }
}
output==>
GOOS: windows
GOARCH: amd64
GOROOT: C:\go
Version: go1.6.1
NumCPU: 4

无线网络连接 15 26:46:19:57:5c:a2
   fe80::e4e5:fbfe:7dad:8576/64
   192.168.191.1/24
无线网络连接 c4:46:19:57:5c:a2
   fe80::d1a3:6c2b:4ecc:f643/64
   10.110.1.186/16
VMware Network Adapter VMnet1 00:50:56:c0:00:01
   fe80::2970:5fe:d4f1:e60b/64
   192.168.209.1/24
VMware Network Adapter VMnet8 00:50:56:c0:00:08
   fe80::945a:8b61:f8a7:cb8e/64
   192.168.171.1/24
Loopback Pseudo-Interface 1 
   ::1/128
   127.0.0.1/8
isatap.wireless 00:00:00:00:00:00:00:e0
   fe80::5efe:a6e:1ba/128
本地连接* 10 00:00:00:00:00:00:00:e0
   fe80::e0:0:0:0/64
isatap.{C86C9321-0E23-416F-A979-B0A1103AB743} 00:00:00:00:00:00:00:e0
isatap.{99DBEBBA-0E54-488B-91B5-21352C7C5B2F} 00:00:00:00:00:00:00:e0
isatap.{BBBEE258-288E-4756-9DC0-E23D3819D958} 00:00:00:00:00:00:00:e0
</pre>
##Golang重试机制的实现
<pre>
package main

import (
	"errors"
	"fmt"
	_ "github.com/astaxie/beego"
	_ "github.com/huichen/sego"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/matryer/try"
	_ "gopkg.in/redis.v3"
	"math/rand"
	_ "stathat.com/c/consistent"
	"time"
)

/*
func main() {
	db, err := gorm.Open("mysql", "root:123456@/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err.Error())
	}
} */
func main() {
	var value string

	err := try.Do(func(attempt int) (bool, error) {
		var err error
		value, err = SomeFunction()
		if err != nil {
			fmt.Println("Run error - ", err)
		} else {
			fmt.Println("Run ok - ", value)
		}

		return attempt < 5, err // 重试5次
	})

	if err != nil {
		fmt.Println("error:", err)
	}
}

func SomeFunction() (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//如果生成随机数大于90，返回成功
	if r.Intn(100) > 90 {
		return "ok", nil
	} else {
		return "", errors.New("network error")
	}
}
output==>
Run error -  network error
Run error -  network error
Run error -  network error
Run error -  network error
Run error -  network error
error: network error
</pre>
##Golang md5和sha1加密算法
md5和sha1都是一种hash算法，加密后不可解密。
sha算法家族中，除了sha1还有SHA-224、SHA-256、SHA-384，和SHA-512，SHA后面的数字表示摘要长度，越长安全性越高，同时计算时消耗也越大。
在go中这几种加密方法的使用很简单：
<pre>
package main

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
)

func main() {
    data := []byte("hello world")
    fmt.Println("%x", sha1.Sum(data))
    fmt.Println("%x", sha256.Sum224(data))
    fmt.Println("%x", sha256.Sum256(data))
    fmt.Println("%x", sha512.Sum384(data))
    fmt.Println("%x", sha512.Sum512(data))
    fmt.Println("%x", md5.Sum(data))
}
output==>
%x [42 174 108 53 201 79 207 180 21 219 233 95 64 139 156 233 30 232 70 237]
%x [47 5 71 127 194 75 180 250 239 216 101 23 21 109 175 222 206 196 91 138 211 207 37 34 165 99 88 43]
%x [185 77 39 185 147 77 62 8 165 46 82 215 218 125 171 250 196 132 239 227 122 83 128 238 144 136 247 172 226 239 205 233]
%x [253 189 142 117 166 127 41 247 1 164 224 64 56 94 46 35 152 99 3 234 16 35 146 17 175 144 127 203 184 53 120 179 228 23 203 113 206 100 110 253 8 25 221 140 8 141 225 189]
%x [48 158 204 72 156 18 214 235 76 196 15 80 201 2 242 180 208 237 119 238 81 26 124 122 155 205 60 168 109 76 216 111 152 157 211 91 197 255 73 150 112 218 52 37 91 69 176 207 216 48 232 31 96 93 207 125 197 84 46 147 174 156 215 111]
%x [94 182 59 187 224 30 238 208 147 203 34 187 143 90 205 195]
</pre>
##Golang获取系统用户信息|Go user包
Go提供了os/user包，用来查询系统用户的信息。
<pre>
package main

import (
    "fmt"
    "os/user"
)

func main() {
    me, _ := user.Current()
    fmt.Println("My Uid : ", me.Uid)
    fmt.Println("My Username : ", me.Username)
    fmt.Println("My Gid : ", me.Gid)
    fmt.Println("My HomeDir : ", me.HomeDir)
    fmt.Println("My Name : ", me.Name)
}
output==>
My Uid :  S-1-5-21-2855060091-2234719249-1014910425-500
My Username :  JASON\Administrator
My Gid :  S-1-5-21-2855060091-2234719249-1014910425-513
My HomeDir :  C:\Users\Administrator
My Name :  
</pre>
###可变函数
<pre>
package main
import "fmt"
func Greeting(who ...string) {
    //接收到who是一个数组，可以用for遍历。
    for _, name := range who {
        fmt.Println(name)
    }
}
func main() {
    Greeting("Hello:", "tom", "mike", "jesse")
}
output==>
Hello:
tom
mike
jesse
</pre>
##Golang database/sql数据库连接类
官方实现的database/sql包中的DB和Stmt是协程安全的，因为内部实现是连接池。
##Golang 高性能分布式数据库influxDB
开源分布式的时序、事件和指标数据库，无需外部依赖。其设计目标是实现分布式和水平伸缩扩展。https://github.com/influxdata/influxdb
##某一位知友自己用Golang实现的k_v数据库
https://github.com/male110/SimpleDb（代码量不多，适合拿来学习模仿）
##golang多协程操作同一全局变量
race检测
##制作APP接口
通过restful规范，按照团队约定返回加密数据。
##一个查询手机号码归属地,运营商,区号信息的Golang库
通过手机号码（中国大陆）查询归属地、运营商等信息
<pre>
package main

import (
	_ "errors"
	"fmt"
	_ "github.com/astaxie/beego"
	_ "github.com/huichen/sego"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/matryer/try"
	"github.com/zheng-ji/gophone"
	_ "gopkg.in/redis.v3"
	_ "math/rand"
	_ "stathat.com/c/consistent"
	_ "time"
)

func main() {
	//打印所有关于该号码的信息
	pr, err := gophone.Find("18855971036")
	if err == nil {
		fmt.Println(pr)
	}
	/*也可以单独获取该号码各个属性
	pr, err = gophone.Find("18855971036")
	if err == nil {
		fmt.Println(pr.PhoneNum)
		fmt.Println(pr.Province)
		fmt.Println(pr.AreaZone)
		fmt.Println(pr.City)
		fmt.Println(pr.ZipCode)
	}
	*/
}
output==>
PhoneNum: 18855971036
AreaZone: 0559
CardType: 移动虚拟运营商
City: 黄山
ZipCode: 242700
Province: 安徽
</pre>
##Golang实现ring buffer缓存池
Ring buffer算法优点：高内存使用率，在缓冲buffer内存模型中，不太容易发生内存越界、悬空指针等 bug ，出了问题也容易在内存级别分析调试。做出来的系统容易保持健壮。
<pre>
package main

type Ringbuf struct {
    buf         []byte
    start, size int
}
 
func New(size int) *Ringbuf {
    return &Ringbuf{make([]byte, size), 0, 0}
}
 
func (r *Ringbuf) Write(b []byte) {
    for len(b) > 0 {
        start := (r.start + r.size) % len(r.buf)
        n := copy(r.buf[start:], b)
        b = b[n:] //golang就是要好好运用切片
 
        if r.size >= len(r.buf) {
            if n <= len(r.buf) {
                r.start += n
                if r.start >= len(r.buf) {
                    r.start = 0
                }
            } else {
                r.start = 0
            }
        }
        r.size += n
        // Size can't exceed the capacity
        if r.size > cap(r.buf) {
            r.size = cap(r.buf)
        }
    }
}
 
func (r *Ringbuf) Read(b []byte) int {
    read := 0
    size := r.size
    start := r.start
    for len(b) > 0 && size > 0 {
        end := start + size
        if end > len(r.buf) {
            end = len(r.buf)
        }
        n := copy(b, r.buf[start:end])
        size -= n
        read += n
        b = b[n:]
 
        start = (start + n) % len(r.buf)
    }
    return read
}
 
func (r *Ringbuf) Size() int {
    return r.size
}

func main(){
	
}
</pre>
##gogoprotobuf(https://github.com/gogo/protobuf/)
磁盘上存储所有的数据都使用了protobuf,然而我们并没有使用Google官方的protobuf类库，我们强烈推荐使用一个叫做gogoprotobuf的第三方包。

gogoprotobuf遵循了很多我们上面提到的关于避免不必要的内存分配的原则。尤其是，它允许将数据编码到一个后端使用数组的字节切片以避免多次内存分配。此外，非空注解允许你直接嵌入消息而无需额外的内存分配开销，这在始终需要嵌入消息时是非常有用的。

最后一点优化是，较基于反射进行编码和解编码的Google标准protobuf类库，gogoprotobuf使用编码和解编码协程提供了不错的性能改善。
##Golang探测局域网里面的设备
<pre>
package main

import (
	"errors"
	"fmt"
	"github.com/franela/goreq"
	"github.com/j-keck/arping"
	"log"
	"net"
	"net/url"
	"strings"
	"time"
)

func main() {
	fing := new(Fing)
	fing.Detect()
	fing.Show()
}

type Fing struct {
	Devices []*Device
}

type Device struct {
	Ip     string
	Mac    string
	Vendor string
	Type   int
}

func NewDevice(ip, mac, vendor string, t int) *Device {
	device := new(Device)
	device.Ip = ip
	device.Mac = mac
	device.Vendor = vendor
	device.Type = t
	return device
}

func (this *Fing) Detect() {
	// Get own IP
	ip, ownmac, err := ExternalIP()
	if err != nil {
		log.Println(err)
		return
	}

	vendor, err := Vendor(ownmac)
	if err != nil {
		log.Println(err)
		return
	}
	this.Devices = append(this.Devices, NewDevice(ip, ownmac, vendor, TYPE_OWN_DEVICE))

	ipFormat := ip[:strings.LastIndex(ip, ".")+1] + "%d"
	for i := 1; i <= 27; i++ {
		nextIp := fmt.Sprintf(ipFormat, i)
		if nextIp != ip {
			hwAddr, duration, err := Mac(nextIp)
			if err == arping.ErrTimeout {
				log.Printf("IP %s is offline.\n", nextIp)
			} else if err != nil {
				log.Printf("IP %s :%s\n", nextIp, err.Error())
			} else {
				log.Printf("%s (%s) %d usec\n", nextIp, hwAddr, duration/1000)
				vendor, err := Vendor(hwAddr.String())
				if err != nil {
					log.Println(err)
					return
				}
				this.Devices = append(this.Devices, NewDevice(nextIp, hwAddr.String(), vendor, TYPE_OTHER_DEVICE))
			}
		}
	}

}

func (this *Fing) Show() {
	fmt.Printf("%3s|%15s|%17s|%20s|%4s\n", "#", "IP", "MAC", "VENDOR", "TYPE")
	for i, device := range this.Devices {
		fmt.Printf("%3d|%15s|%17s|%20s|%4s\n", i, device.Ip, device.Mac, device.Vendor, this.showType(device.Type))
	}
}

func (this *Fing) showType(t int) string {
	switch t {
	case TYPE_OWN_DEVICE:
		return "OWN"
	}
	return ""
}

const (
	TYPE_OWN_DEVICE = iota
	TYPE_OTHER_DEVICE
)

func Vendor(mac string) (string, error) {
	macs := strings.Split(mac, ":")
	if len(macs) != 6 {
		return "", fmt.Errorf("MAC Error: %s", mac)
	}
	mac = strings.Join(macs[0:3], "-")
	form := url.Values{}
	form.Add("x", mac)
	form.Add("submit2", "Search!")
	res, err := goreq.Request{
		Method:      "POST",
		Uri:         "http://standards.ieee.org/cgi-bin/ouisearch",
		ContentType: "application/x-www-form-urlencoded",
		UserAgent:   "Cyeam",
		ShowDebug:   true,
		Body:        form.Encode(),
	}.Do()
	if err != nil {
		return "", err
	}
	body, err := res.Body.ToString()
	if err != nil {
		return "", err
	}
	vendor := body[strings.Index(body, strings.ToUpper(mac))+len(mac):]
	vendor = strings.TrimLeft(vendor, "</b>   (hex)")
	vendor = strings.TrimSpace(vendor)
	return strings.Split(vendor, "\n")[0], nil
}

func Mac(ip string) (net.HardwareAddr, time.Duration, error) {
	dstIP := net.ParseIP(ip)
	return arping.Ping(dstIP)
}

func ExternalIP() (string, string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), iface.HardwareAddr.String(), nil
		}
	}
	return "", "", errors.New("are you connected to the network?")
}
outout==>
             IP|              MAC|              VENDOR|TYPE 
	
	2016/04/25 16:53:11 POST /cgi-bin/ouisearch HTTP/1.1
	Host: standards.ieee.org
	Content-Type: application/x-www-form-urlencoded
	User-Agent: Cyeam
	
	submit2=Search%21&x=26-46-19
	2016/04/25 16:53:12 gzip: invalid checksum
</pre>
###Golang实现插入排序
<pre>
package main

import "fmt"

func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

type BySortIndex []int

func (a BySortIndex) Len() int      { return len(a) }
func (a BySortIndex) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySortIndex) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	test0 := []int{49, 38, 65, 97, 76, 13, 27, 49}
	insertionSort(BySortIndex(test0), 0, len(test0))
	fmt.Println(test0)
}

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
output==>
[13 27 38 49 49 65 76 97]
</pre>
###Golang实现反转字符串|字符串反转
<pre>
package main

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func main() {
	a := "Hello, 世界"
	println(a)
	println(Reverse(a))
}
outout==>
Hello, 世界
界世 ,olleH
</pre>
##Golang自己实现ToUpper函数|转换成大写
把一个字符串中的字符从小写转为大写
<pre>
package main

import "fmt"

const MaxASCII = '\u007F'

func toUpper(r rune) rune {
	if r <= MaxASCII {
		if 'a' <= r && r <= 'z' {
			r -= 'a' - 'A'
		}
		return r
	}
	return r
}

func ToUpper(s []rune) (res []rune) {
	for i := 0; i < len(s); i++ {
		res = append(res, toUpper(s[i]))
	}
	return res
}

func main() {
	a := "Hello, 世界"
	fmt.Println(string(ToUpper([]rune(a))))

}
output==>
HELLO, 世界
</pre>
##Golang限制协程的最大开启数
<pre>
package main
 
import (
    "fmt"
    "strconv"
    "time"
)
 
var (
    maxRoutineNum = 10
)
 
// 模拟下载页面的方法
func download(url string, ch chan int) {
    fmt.Println("download from ", url)
    // 休眠两秒模拟下载页面
    time.Sleep(2 * 1e9)
    // 下载完成则从ch推出数据
    <-ch
}
 
func main() {
    ch := make(chan int, maxRoutineNum)
 
    urls := [100]string{}
    for i := 0; i < 100; i++ {
        urls[i] = "url" + strconv.Itoa(i)
    }
    for i := 0; i < len(urls); i++ {
        // 开启下载协程前往ch塞一个数据
        // 如果ch满了则会处于阻塞，从而达到限制最大协程的功能
        ch <- 1
        go download(urls[i], ch)
    }
 
    // 休眠一下
    for {
        time.Sleep(1 * 1e9)
    }
}
</pre>
##Golang中zip压缩和解压文件
zip压缩文件
<pre>
package main
 
import (
    "archive/zip"
    "bytes"
    "log"
    "os"
)
 
func main() {
    // 创建一个缓冲区用来保存压缩文件内容
    buf := new(bytes.Buffer)
 
    // 创建一个压缩文档
    w := zip.NewWriter(buf)
 
    // 将文件加入压缩文档
    var files = []struct {
        Name, Body string
    }{
        {"readme.txt", "This archive contains some text files."},
        {"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
        {"todo.txt", "Get animal handling licence.\nWrite more examples."},
    }
    for _, file := range files {
        f, err := w.Create(file.Name)
        if err != nil {
            log.Fatal(err)
        }
        _, err = f.Write([]byte(file.Body))
        if err != nil {
            log.Fatal(err)
        }
    }
 
    // 关闭压缩文档
    err := w.Close()
    if err != nil {
        log.Fatal(err)
    }
 
    // 将压缩文档内容写入文件
    f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    buf.WriteTo(f)
}
</pre>
zip解压文件(针对file.zip文件)
<pre>
package main
 
import (
    "archive/zip"
    "fmt"
    "io"
    "log"
    "os"
)
 
func main() {
    // 打开一个zip格式文件
    r, err := zip.OpenReader("file.zip")
    if err != nil {
        log.Fatal(err)
    }
    defer r.Close()
 
    // 迭代压缩文件中的文件，打印出文件中的内容
    for _, f := range r.File {
        fmt.Printf("文件名 %s:\n", f.Name)
        rc, err := f.Open()
        if err != nil {
            log.Fatal(err)
        }
        _, err = io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64))
        if err != nil {
            log.Fatal(err)
        }
        rc.Close()
        fmt.Println()
    }
 
}
output==>
文件名 readme.txt:
This archive contains some text files.
文件名 gopher.txt:
Gopher names:
George
Geoffrey
Gonzo
文件名 todo.txt:
Get animal handling licence.
Write more examples.
</pre>
##Golang语言接口开发——不确定JSON数据结构的解析
在公司主要做接口的开发，会经常遇到接口对接的情况。有的时候，同一个请求返回的JSON数据格式并不一样。如果是正常，则可能只返回一个status字段，说明正常；如果中间出错，除了在status字段里面说明错误类型，还会通过error_message附带错误详细信息。比如要给用户加积分，如果加分失败，还会附带用户id等信息。那么，请求一个接口可能的返回值就是不确定的。

我最初就是定义两个结构体，我处理的数据都共有一个字段status，如果能够解析并且status表示操作成功，那么用封装成功内容的结构体解析；否则，用封装失败的结构体解析。这就是传说中的DIRTY HACK。。。

后来，偶然发现封装正确的结构体也会解析错误的字符串，当然，只会解析共有字段。那么，这个问题就好解决多了。把两个结构体放到一起即可，如果没有该字段，就不会被解析放入值。也就是说，未被解析的变量放的是默认值。
<pre>
package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Status       int    `json:"status"`
	Message      string `json:"message"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func main() {
	json_str0 := `{"status":0,"message":"success"}`
	json_str1 := `{"status":1,"error_code":5,"error_message":"error"}`

	res0 := Result{}
	res1 := Result{}

	err0 := json.Unmarshal([]byte(json_str0), &res0)
	err1 := json.Unmarshal([]byte(json_str1), &res1)

	fmt.Println(res0, err0)
	fmt.Println(res1, err1)
}
output==>
{0 success 0 } <nil>
{1  5 error} <nil>
</pre>
###Golang控制客户端连接数量
<pre>
package main

import (
    "io"
    "log"
    "net/http"
)    

func maxClients(h http.Handler, n int) http.Handler {
     sema := make(chan struct{}, n)

     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
         sema <- struct{}{}
         defer func() { <-sema }()

         h.ServeHTTP(w, r)
     })
}

func main() {
     handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
         res := getExpensiveResource()
         io.WriteString(w, res.String())
     })

     http.Handle("/", maxClients(handler, 10))

     log.Fatal(http.ListenAndServe(":8080", nil))
}
//控制客户端连接数量
</pre>
###让Golang程序达到平滑重启的第三方库
https://github.com/rcrowley/goagain（stars>1k）

基本原理：

goagain会监控2个系统信号，一个为SIGTERM，接收到这个信号，程序就停止运行。另一个信号为SIGUSR2，接收到这个信号的行为是，当前进程，也就是父进程会新建一个子进程，然后把父进程的pid保存到一个名为GOAGAIN_PPID的环境变量；子进程启动的时候会检索GOAGAIN_PPID这个变量，来判断程序是否要重启，通过这个变量来关闭父进程，来达到平滑重启的效果。
###Golang通过遗传算法找出相近的图片的第三方库
https://github.com/armhold/polygen
###Golang使用多核
<pre>
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go go1()
	go go2()
	time.Sleep(20 * time.Millisecond)
}

func go1() {
	for {
		fmt.Print(1)
		runtime.Gosched()
	}
}

func go2() {
	for {
		fmt.Print(2)
		runtime.Gosched()
	}
}
output==>
212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212111212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212112121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121212121
</pre>
###Golang爬虫第三方库
- https://github.com/PuerkitoBio/goquery
- 分布式、中文手册爬虫 https://github.com/henrylee2cn/pholcus
###Golang实现的mysql读写分离分表分库|MySQL proxy
Go开发高性能MySQL Proxy项目，kingshard在满足基本的读写分离的功能上，致力于简化MySQL分库分表操作；能够让DBA通过kingshard轻松平滑地实现MySQL数据库扩容。 kingshard的性能大约是直连MySQL性能的80%以上。

https://github.com/flike/kingshard