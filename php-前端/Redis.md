####这里介绍的是window下面的Redis使用体验。
Redis最为常用的数据类型主要有以下五种：

* String
* Hash
* List
* Set
* Sorted set

1. String
set与get是使用频率最高的命令了。分别是设置（新增或者是修改）与取出（显示）一个字符串键值对。<br>
在redis-cli.exe客户端中，键入:
<pre>
127.0.0.1:6379 > set name Jason
ok 
127.0.0.1:6379 > get name
"Jason"
</pre>
mget一次取出多个键值
<pre>
127.0.0.1:6379 > set name1 Jason1
ok
127.0.0.1:6379 > set name2 Jason2
ok
127.0.0.1:6379 > mget name1 name2
1>Jason1
2>Jason2
</pre>
decr与incr在当值是int型时，分别是减1与加1的效果。这里不做演示。

2. Hash
常用命令是hget,hset,hgetall等。
<pre>
127.0.0.1:6379 > hset hash name jason
ok
127.0.0.1:6379 > hset hash age 34
ok
127.0.0.1:6379 > hset hash email 454@qq.com
ok
127.0.0.1:6379 > hget hash age
"34"
127.0.0.1:6379 > hgetall hash
1>"name"
2>"jason"
3>"age"
4>"34"
5>"email"
6>"454@qq.com"
</pre>