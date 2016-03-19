###net/http
在net/http包中，动态文件的路由和静态文件的路由是分开的，动态文件使用http.HandleFunc进行设置，静态文件就需要使用到http.FileServer
####如何设置cookie
<pre>
cookie := http.Cookie{Name: "admin_name", Value: rows[0].Str(res.Map("admin_name")), Path: "/"}
http.SetCookie(w, &cookie)
</pre>
