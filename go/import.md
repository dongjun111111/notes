##import详解
我们在写Go代码的时候经常用到import这个命令用来导入包文件，而我们经常看到的方式参考如下：
<pre>
 import(
    "fmt"
)
</pre>
然后我们代码里面可以通过如下的方式调用
<pre>
 fmt.Println("hello world")
</pre>
上面这个fmt是Go语言的标准库，其实是去GOROOT环境变量指定目录下去加载该模块，当然Go的import还支持如下两种方式来加载自己写的模块：

* 相对路径
<pre>
import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import
</pre>
* 绝对路径
<pre>
import “shorturl/model”//加载gopath/src/shorturl/model模块
</pre>
上面展示了一些import常用的几种方式，但是还有一些特殊的import，让很多新手很费解，下面我们来一一讲解一下到底是怎么一回事

* 点操作<br>
我们有时候会看到如下的方式导入包
<pre>
 import(
    . "fmt"
)
</pre>
这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("helloworld")可以省略的写成Println("helloworld")

* 别名操作<br>
别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
<pre>
 import(
    f "fmt"
)
</pre>
别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("helloworld")

*  _ 操作<br>
这个操作经常是让很多人费解的一个操作符，请看下面这个import
<pre>
 import (
    "database/sql"
    _ "github.com/ziutek/mymysql/godrv"
)
</pre>
_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。


