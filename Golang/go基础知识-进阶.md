###net/http
在net/http包中，动态文件的路由和静态文件的路由是分开的，动态文件使用http.HandleFunc进行设置，静态文件就需要使用到http.FileServer
####如何设置cookie
<pre>
cookie := http.Cookie{Name: "admin_name", Value: rows[0].Str(res.Map("admin_name")), Path: "/"}
http.SetCookie(w, &cookie)
</pre>
####http.FileServer()
文件系统。将本地文件输出到网页
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
http.Handle("/doc",http.StripPrefix("/doc",http.FileServer(http.Dir("./"))))   //在浏览器地址栏输入localhost:8123/doc ,显示同上面一样的结果
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
####Golang发送email邮件
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
###条件变量
在Go语言中，sync.Cond类型代表了条件变量。与互斥锁和读写锁不同，简单的声明无法创建出一个可用的条件变量。为了得到这样一个条件变量，我们需要用到sync.NewCond函数。该函数的声明如下：
<pre>
func NewCond(l Locker) *Cond
</pre>
条件变量总是要与互斥量组合使用。因此，sync.NewCond函数的唯一参数是sync.Locker类型的，而具体的参数值既可以是一个互斥锁也可以是一个读写锁。sync.NewCond函数在被调用之后会返回一个*sync.Cond类型的结果值。我们可以调用该值拥有的几个方法来操纵对应的条件变量。

类型*sync.Cond的方法集合中有三个方法，即：Wait方法、Signal方法和Broadcast方法。它们分别代表了等待通知、单发通知和广播通知的操作。

方法Wait会自动的对与该条件变量关联的那个锁进行解锁，并且使调用方所在的Goroutine被阻塞。一旦该方法收到通知，就会试图再次锁定该锁。如果锁定成功，它就会唤醒那个被它阻塞的Goroutine。否则，该方法会等待下一个通知，那个Goroutine也会继续被阻塞。而方法Signal和Broadcast的作用都是发送通知以唤醒正在为此而被阻塞的Goroutine。不同的是，前者的目标只有一个，而后者的目标则是所有。

在Read方法中，我们使用一个for循环来达到重新尝试获取数据块的目的。为此，我们添加了若干条重复的语句、降低了程序的性能，还造成了一个潜在的问题——在某个情况下读写锁fmutex不会被读解锁。为了解决这一系列新生的问题，我们使用代表条件变量的字段rcond。
案例：



