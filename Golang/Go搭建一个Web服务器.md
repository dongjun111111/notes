###http包建立Web服务器
<pre>
package main
import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  
    fmt.Println(r.Form)  
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") 
}

func main() {
    http.HandleFunc("/", sayhelloName) 
    err := http.ListenAndServe(":9090", nil) 
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
</pre>
build之后，然后执行web.exe,这个时候其实已经在9090端口监听http链接请求了,打开http://localhost:9090.

####web服务器制作原理
首先要了解http包执行流程。

-创建Listen Socket, 监听指定的端口, 等待客户端请求到来。
- Listen Socket接受客户端的请求, 得到Client Socket, 接下来通过Client Socket与客户端通信。
- 处理客户端的请求, 首先从Client Socket读取HTTP请求的协议头, 如果是POST方法, 还可能要读取客户端提交的数据, 然后交给相应的handler处理请求, handler处理完毕准备好客户端需要的数据, 通过Client Socket写给客户端。

也就是：

- 如何监听端口？
- 如何接收客户端请求？
- 如何分配handler？
Go是通过一个函数ListenAndServe来处理这些事情的，这个底层其实这样处理的：初始化一个server对象，然后调用了net.Listen("tcp", addr)，也就是底层用TCP协议搭建了一个服务，然后监控我们设置的端口。

###Go的http包详解
Go的http有两个核心功能：Conn、ServeMux。

与我们一般编写的http服务器不同, Go为了实现高并发和高性能, 使用了goroutines来处理Conn的读写事件, 这样每个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。这是Go高效的保证。

Go在等待客户端请求里面是这样写的：
<pre>
c, err := srv.newConn(rw)
if err != nil {
    continue
}
go c.serve()
</pre>
这里我们可以看到客户端的每次请求都会创建一个Conn，这个Conn里面保存了该次请求的信息，然后再传递到对应的handler，该handler中便可以读取到相应的header信息，这样保证了每个请求的独立性。
####简易路由器实现
<pre>
package main
import (
    "fmt"
    "net/http"
)
type MyMux struct {
}
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        sayhelloName(w, r)
        return
    }
    http.NotFound(w, r)
    return
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello myroute!")
}
func main() {
    mux := &MyMux{}
    http.ListenAndServe(":9090", mux)
}
</pre>

###Go代码的执行流程
通过对http包的分析之后，现在让我们来梳理一下整个的代码执行过程。

首先调用Http.HandleFunc

按顺序做了几件事：

1 调用了DefaultServeMux的HandleFunc

2 调用了DefaultServeMux的Handle

3 往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则,其次调用http.ListenAndServe(":9090", nil)

按顺序做了几件事情：

1 实例化Server

2 调用Server的ListenAndServe()

3 调用net.Listen("tcp", addr)监听端口

4 启动一个for循环，在循环体中Accept请求

5 对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()

6 读取每个请求的内容w, err := c.readRequest()

7 判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），handler就设置为DefaultServeMux

8 调用handler的ServeHttp

9 在这个例子中，下面就进入到DefaultServeMux.ServeHttp

10 根据request选择handler，并且进入到这个handler的ServeHTTP
mux.handler(r).ServeHTTP(w, r)

11 选择handler：

A 判断是否有路由能满足这个request（循环遍历ServerMux的muxEntry）

B 如果有路由满足，调用这个路由handler的ServeHttp

C 如果没有路由满足，调用NotFoundHandler的ServeHttp