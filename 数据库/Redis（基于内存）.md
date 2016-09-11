####这里介绍的是window下面的Redis使用体验
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
常用命令是lpush,rpush,lpop,rpop,lrange等。插入分为左右两种情况，分别对应lpush与rpush,因为两者的效果类似，为求简单直观，下面大部分以lpush为主进行介绍。 <br>
应用场景：<br>
Redis list的应用场景非常多，也是Redis最重要的数据结构之一，比如twitter的关注列表，粉丝列表等都可以用Redis的list结构来实现，比较好理解，这里不再重复。

s
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
127.0.0.1:6379 >del group *                     //删除group组下所有元素
<integer>1
127.0.0.1:6379 >llen group                      //输出数组group长度（元素个数）
<integer>0
127.0.0.1:6379 >lpush group 1
<integer>1
127.0.0.1:6379 >lpush group 2
<integer>1
127.0.0.1:6379 >lpop group                     //1将从group组中移除
"1"
</pre>
实现方式：
Redis list的实现为一个双向链表，即可以支持反向查找和遍历，更方便操作，不过带来了部分额外的内存开销，Redis内部的很多实现，包括发送缓冲队列等也都是用的这个数据结构。

4. Set
常用命令是sadd,spop,smembers,sunion等等.<br>
应用场景：<br>
Redis set对外提供的功能与list类似是一个列表的功能，特殊之处在于set是可以自动排重的，当你需要存储一个列表数据，又不希望出现重复数据时，set是一个很好的选择，并且set提供了判断某个成员是否在一个set集合内的重要接口，这个也是list所不能提供的。<br>
实现方式：<br>
set 的内部实现是一个 value永远为null的HashMap，实际就是通过计算hash的方式来快速排重的，这也是set能提供判断一个成员是否在集合内的原因。
<pre>
127.0.0.1:6379 >sadd bbs 23
<integer>1
127.0.0.1:6379 >sadd bbs 23
<integer>0                           //添加失败，已经存在
127.0.0.1:6379 >sadd bbs 12
<integer>1
127.0.0.1:6379 >smembers bbs        //显示所有元素
1>"23"
2>"12"
127.0.0.1:6379 >del bbs              //删除所有元素
127.0.0.1:6379 >sadd bbs 1
<integer>1
127.0.0.1:6379 >sadd bbs 2
<integer>1
127.0.0.1:6379 >sadd bbs 3
<integer>1
127.0.0.1:6379 >spop bbs             //随机删除一个元素
"2"
127.0.0.1:6379 >sadd bbb 77
<integer>1
127.0.0.1:6379 >sunion bbs bbb       //返回两个集合的元素
1>"1"
2>"2"
3>"77"
</pre>

5. Sorted set
常用命令：zadd,zrange,zrem,zcard等<br>
使用场景：<br>
Redis sorted set的使用场景与set类似，区别是set不是自动有序的，而sorted set可以通过用户额外提供一个优先级(score)的参数来为成员排序，并且是插入有序的，即自动排序。当你需要一个有序的并且不重复的集合列表，那么可以选择sorted set数据结构，比如twitter 的public timeline可以以发表时间作为score来存储，这样获取时就是自动按时间排好序的。<br>
实现方式：<br>
Redis sorted set的内部使用HashMap和跳跃表(SkipList)来保证数据的存储和有序，HashMap里放的是成员到score的映射，而跳跃表里存放的是所有的成员，排序依据是HashMap里存的score,使用跳跃表的结构可以获得比较高的查找效率，并且在实现上比较简单。<br>
<pre>
127.0.0.1:6379 >zadd group 1 1
<integer>1
127.0.0.1:6379 >zadd group 2 2
<integer>1
127.0.0.1:6379 >zadd group 3 3
<integer>1
127.0.0.1:6379 >zrange group 0 10
1>"1"
2>"2"
3>"3"
127.0.0.1:6379 >zadd group 4 gy
<integer>1
127.0.0.1:6379 >zrem group gy       //移除一个元素
<integer>1
127.0.0.1:6379 >zrange group 0 10    //遍历一个集合
1>"1"
2>"2"
3>"3"
127.0.0.1:6379 >zcard group       //返回元素个数
<integer>3
</pre>


##Redis与Memcached
Redis是一个高性能的key-value存储系统，和Memcached类似，它支持存储的value类型相对更多，包括string（字符串）、list（链表）、set（集合）和zset（有序集合）。与memcached一样，为了保证效率，数据都是缓存在内存中，区别的是Redis会周期性的把更新的数据写入磁盘或者把修改操作写入追加的记录文件，并且在此基础上实现了主从同步。 

Redis的出现，很大程度补偿了memcached这类key/value存储的不足，在部分场合可以对关系数据库起到很好的补充作用。

##Redis处理过期的Key
redis如何清理数据库中过期的键？它分为两种：

- 惰性删除：当你去操作一个键（譬如 get name），redis首先会检查这个键是否关联了一个超时时间，如果有，则检查是否超时，若超时则返回空，否则返回相应的值；
- 定时删除：redis中有个时间事件，它会清理数据库中已经过期的键（redis会限定该操作占用的时间，避免阻塞客户端的请求）

一般情况下我们要手动操作（调用）会生成缓存的方法来实现过期的key会被及时清除。这也是使用Redis中的惰性删除的属性。