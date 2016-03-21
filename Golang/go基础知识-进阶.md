###net/http
在net/http包中，动态文件的路由和静态文件的路由是分开的，动态文件使用http.HandleFunc进行设置，静态文件就需要使用到http.FileServer
####如何设置cookie
<pre>
cookie := http.Cookie{Name: "admin_name", Value: rows[0].Str(res.Map("admin_name")), Path: "/"}
http.SetCookie(w, &cookie)
</pre>
####http.FileServer()
文件系统显示本地文件在网页上。
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
-RWMutex
定义：<b>它允许任意读操作同时进行 同一时刻，只允许有一个写操作进行.并且一个写操作被进行过程中，读操作的进行也是不被允许的,读写锁控制下的多个写操作之间都是互斥的,写操作与读操作之间也都是互斥的,多个读操作之间却不存在互斥关系</b>.RWMutex是一个读写锁，该锁可以加多个读锁或者一个写锁。写锁，如果在添加写锁之前已经有其他的读锁和写锁，则lock就会阻塞直到该锁可用，为确保该锁最终可用，已阻塞的 Lock 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定。写锁解锁，如果没有进行写锁定，则就会引起一个运行时错误．<br>
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

