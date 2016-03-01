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
Hash实现方式：
上面已经说到Redis Hash对应Value内部实际就是一个HashMap，实际这里会有2种不同实现，这个Hash的成员比较少时Redis为了节省内存会采用类似一维数组的方式来紧凑存储，而不会采用真正的HashMap结构，对应的value redisObject的encoding为zipmap,当成员数量增大时会自动转成真正的HashMap,此时encoding为ht。

3. List
常用命令是lpush,rpush,lpop,rpop,lrange等。插入分为左右两种情况，分别对应lpush与rpush,因为两者的效果类似，为求简单直观，下面大部分以lpush为主进行介绍。 
<pre>
127.0.0.1:6379 >lpush group1 34
<integer> 1
127.0.0.1:6379 >lpush group1 yes
<integer> 1
127.0.0.1:6379 >lpush group1 56
<integer> 1
127.0.0.1:6379 >lrange group1 0 10             //将0-10位的元素遍历出来
1>"34"
2>"yes"
3>"56"
127.0.0.1:6379 >del group *                    //删除group组下所有元素
<integer>1
127.0.0.1:6379 >llen group                      //输出数组group长度（元素个数）
<integer>0
127.0.0.1:6379 >lpush group 1
<integer>1
127.0.0.1:6379 >lpush group 2
<integer>1
127.0.0.1:6379 >lpop group                  //1将从group组中移除
"1"
</pre>