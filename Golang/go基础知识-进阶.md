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
不定参数
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
函数闭包：
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
函数递归
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
结构体方法
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
错误处理-Error接口
<pre>

</pre>