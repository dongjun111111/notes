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
