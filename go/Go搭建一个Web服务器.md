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