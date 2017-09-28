### 零拷贝
零拷贝（zero copy）技术概述
什么是零拷贝？
简单一点来说，零拷贝就是一种避免 CPU 将数据从一块存储拷贝到另外一块存储的技术。针对操作系统中的设备驱动程序、文件系统以及网络协议堆栈而出现的各种零拷贝技术极大地提升了特定应用程序的性能，并且使得这些应用程序可以更加有效地利用系统资源。这种性能的提升就是通过在数据拷贝进行的同时，允许 CPU 执行其他的任务来实现的。零拷贝技术可以减少数据拷贝和共享总线操作的次数，消除传输数据在存储器之间不必要的中间拷贝次数，从而有效地提高数据传输效率。而且，零拷贝技术减少了用户应用程序地址空间和操作系统内核地址空间之间因为上下文切换而带来的开销。进行大量的数据拷贝操作其实是一件简单的任务，从操作系统的角度来说，如果 CPU 一直被占用着去执行这项简单的任务，那么这将会是很浪费资源的；如果有其他比较简单的系统部件可以代劳这件事情，从而使得 CPU 解脱出来可以做别的事情，那么系统资源的利用则会更加有效。综上所述，零拷贝技术的目标可以概括如下：
避免数据拷贝
避免操作系统内核缓冲区之间进行数据拷贝操作。
避免操作系统内核和用户应用程序地址空间这两者之间进行数据拷贝操作。
用户应用程序可以避开操作系统直接访问硬件存储。
数据传输尽量让 DMA 来做。
将多种操作结合在一起
避免不必要的系统调用和上下文切换。
需要拷贝的数据可以先被缓存起来。
对数据进行处理尽量让硬件来做。
前文提到过，对于高速网络来说，零拷贝技术是非常重要的。这是因为高速网络的网络链接能力与 CPU 的处理能力接近，甚至会超过 CPU 的处理能力。如果是这样的话，那么 CPU 就有可能需要花费几乎所有的时间去拷贝要传输的数据，而没有能力再去做别的事情，这就产生了性能瓶颈，限制了通讯速率，从而降低了网络链接的能力。一般来说，一个 CPU 时钟周期可以处理一位的数据。举例来说，一个 1 GHz 的处理器可以对 1Gbit/s 的网络链接进行传统的数据拷贝操作，但是如果是 10 Gbit/s 的网络，那么对于相同的处理器来说，零拷贝技术就变得非常重要了。对于超过 1 Gbit/s 的网络链接来说，零拷贝技术在超级计算机集群以及大型的商业数据中心中都有所应用。然而，随着信息技术的发展，1 Gbit/s，10 Gbit/s 以及 100 Gbit/s 的网络会越来越普及，那么零拷贝技术也会变得越来越普及，这是因为网络链接的处理能力比 CPU 的处理能力的增长要快得多。传统的数据拷贝受限于传统的操作系统或者通信协议，这就限制了数据传输性能。零拷贝技术通过减少数据拷贝次数，简化协议处理的层次，在应用程序和网络之间提供更快的数据传输方法，从而可以有效地降低通信延迟，提高网络吞吐率。零拷贝技术是实现主机或者路由器等设备高速网络接口的主要技术之一。
现代的 CPU 和存储体系结构提供了很多相关的功能来减少或避免 I/O 操作过程中产生的不必要的 CPU 数据拷贝操作，但是，CPU 和存储体系结构的这种优势经常被过高估计。存储体系结构的复杂性以及网络协议中必需的数据传输可能会产生问题，有时甚至会导致零拷贝这种技术的优点完全丧失。

### 说Kafka是下一代分布式消息系统的原因
kafka提倡使用拉模式，并且可以对消息重复消费，看起来不符合传统queue的思想，但却提供了额外的好处，比如：某模块更新到产线发现有bug，需要将上线以来的消息全部重新消费，即消息回溯。

kafka是高并发型的消息队列，但这是有前提条件的。条件是topic要定义多个partition，将压力分担到各个partition上。topic是逻辑概念，partition是物理存在各个broker，以此达到负载均衡的目的。要注意的是，各个partition可以独立消费，各partition间的消息是无法保证顺序性的，顺序只存在同一partition。以我的经验看，无论哪种MQ，要严格保证顺序，都要付出昂贵的代价，因此弱化顺序是有必要的。

kafka的另一个特性是高可用。放眼目前业界数据层的高可用解决方案，采用的无非都是两种：冗余数据和共享存储。后者以价格昂贵著称，比如SAN，给土豪公司玩的。在党中央构建节约性社会的号召下，我建议使用前者。冗余数据最常见的便是日志复制，kafka的道理也一样。由一组节点组成leader，follower组成小的cluster，由zookeeper做协调(Paxos算法)。leader，follower的比例和数量可配置，一般为1:2。在写入的时候, follower会不断复制leader的数据，leader挂掉后会从follwer中选举新的leader。

kafka使用了零拷贝技术来优化性能，直接发送磁盘的数据到socket。此为其极为取巧的设计和亮点。

### MySQL 几个基本操作
<pre>
1. 创建用户oldboy，使之可以管理数据库oldboy
mysql>grant all on oldboy.* to oldboy@'localhost' identified by '123';

2. 查看创建的用户oldboy拥有哪些权限
mysql>show grants for oldboy@'localhost'\G;

3. 查看当前数据库里有哪些用户
mysql>select user,host from mysql.user;

4.
delete是逻辑删除表中的数据，一列一列的删除表中数据，速度比较慢
mysql> delete from test;
truncate是物理删除表中的数据，一次性全部都给清空表中数据，速度很快
mysql> truncate table test;


-- binlog 文件位置
show variables like 'log_bin_basename';

-- 当前binlog文件
show master status;

-- 清空现有的所用binlog;
reset master;

-- 查看当前MySQL配置   log_bin ON则表明已经开启二进制日志binlog

show variables like '%bin%';


show variables like '%char%';  
set character_set_server=utf8;
-- 配置文件my.ini


--  利用bin_log恢复数据

/usr/bin/mysqlbinlog  --no-defaults  mysql-bin.000034 --start-datetime='2017-07-17 00:00:00' --stop-datetime='2017-07-17 14:00:00'  > binlogtest.sql;
</pre>
 grep用法  -v 不包含  -E 多个条件联合
<pre>
cat bak.sql |grep -v  "INSERT INTO `log`" | grep -E "INSERT"


-- EXISTS的用法
SELECT * FROM activity  WHERE NOT EXISTS  (SELECT * FROM advises WHERE id=-1);
SELECT * FROM activity  WHERE  EXISTS  (SELECT * FROM advises WHERE id=-1);
</pre>

查看所有操作记录】Git记录着你输入的每一条指令！键入查看自己的每一次提交：
git reflog
你会发现，版本号就在这里：
然后键入：
git reset --hard 版本号

gitk 图形化界面显示git内容

创建本地分支
git branch test  
把分支推到远程分支【创建远程分支】
git push origin test
删除本地分支   
git branch -d xxxxx
删除远程分支  
git branch -r -d origin/branch-name  
git push origin :branch-name  

### 数据链路
在数据通信网中，按一种链路协议的技术要求连接两个或多个数据站的电信设施，称为数据链路，简称数据链。数据链路(data link) 除了物理线路外，还必须有通信协议来控制这些数据的传输。若把实现这些协议的硬件和软件加到链路上，就构成了数据链路。

MySQL数据库开启数据库链路：

1. 开启federated引擎  ||federated  联合的；联邦的；
　　
windows下在my.ini中加入federated，即可开启;
　　
linux中，需要编译时加入选项，再在my.ini中加入federated，方可开启。

### Mysql常用命令
schema 计划，图表  [ˈski:mə]
<pre>
select TABLE_NAME from INFORMATION_SCHEMA.columns where COLUMN_NAME  like 'contract_code';

SELECT * FROM `performance_schema`.users;   //查询该数据库的用户以及连接数
SELECT * FROM `performance_schema`.accounts;
SELECT * FROM `performance_schema`.hosts;  //查询该数据库的使用者IP以及连接数


show index from overdue;
show columns from overdue;
ALTER TABLE `overdue` ADD INDEX idx_contract_code( `contract_code` ); -- 添加查询索引
ALTER TABLE `overdue` ADD UNIQUE (`code`) ; -- 添加唯一索引
ALTER TABLE `overdue` ADD PRIMARY KEY ( `code` );-- 添加主键索引


 慢查询日志开启：

在配置文件my.cnf或my.ini中在[mysqld]一行下面加入两个配置参数

log-slow-queries=/data/mysqldata/slow-query.log           

long_query_time=2                                                                 

注：log-slow-queries参数为慢查询日志存放的位置，一般这个目录要有mysql的运行帐号的可写权限，一般都将这个目录设置为mysql的数据存放目录；

long_query_time=2中的2表示查询超过两秒才记录；

在my.cnf或者my.ini中添加log-queries-not-using-indexes参数，表示记录下没有使用索引的查询。

log-slow-queries=/data/mysqldata/slow-query.log           

long_query_time=10                                                               

log-queries-not-using-indexes    

//查看所有链接的详细信息
show full processlist;
</pre>
### 并发与并行
并发与并行都可以是多线程，就看这些线程能不能同时被（多个）cpu执行，如果可以就说明是并行，而并发是多个线程被（一个）cpu轮流切换着执行。

在并发程序中可以同时拥有两个或者多个线程。这意味着，如果程序在单核处理器上运行，那么这两个线程将交替地换入或者换出内存。这些线程是同时“存在”的——每个线程都处于执行过程中的某个状态。如果程序能够并行执行，那么就一定是运行在多核处理器上。此时，程序中的每个线程都将分配到一个独立的处理器核上，因此可以同时运行。


并发就是指代码逻辑上可以并行，有并行的潜力，但是不一定当前是真的以物理并行的方式运行【物理上只有一个cpu在执行】

并发指的是代码的性质，并行指的是物理运行状态 【物理上有多个cpu在执行】


Redis:
redis中list数据结构：入队列时用lpush，拿数据时用brpop。

String——字符串
Hash——字典
List——列表
Set——集合 //Set 就是一个集合，集合的概念就是一堆不重复值的组合  应用：1.共同好友、二度好友 2.利用唯一性，可以统计访问网站的所有独立 IP 3.好友推荐的时候，根据 tag 求交集，大于某个 threshold 就可以推荐
Sorted Set——有序集合  //Sorted Sets是将 Set 中的元素增加了一个权重参数 score，使得集合中的元素能够按 score 进行有序排列

### 闭包 closure
子函数可以访问父函数的局部变量

### MySQL 严格模式
查看MySQL的SQL模式
mysql> select @@sql_mode;
+----------------------------------------------------------------+
| @@sql_mode                                                     |
+----------------------------------------------------------------+
| STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION |
+----------------------------------------------------------------+

临时开启：
set sql_mode="STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION";

检查是否生效，执行sql：select @@sql_mode;

永久开启：

通过配置文件修改：linux找my.cnf文件；sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES。

严格模式主要用以下场景：

1. 不支持对not null字段插入null值
2. 不支持对自增长字段插入”值
3. 不支持text字段有默认值


### Golang 性能调优
<pre>
package main

import (
"fmt"
"log"
"os"
"runtime"
"runtime/debug"
"runtime/pprof"
"strconv"
"sync/atomic"
"syscall"
"time"
)


var heapProfileCounter int32
var startTime = time.Now()
var pid int

func init() {
pid = os.Getpid()
}

func StartCPUProfile() {
f, err := os.Create("cpu-" + strconv.Itoa(pid) + ".pprof")
if err != nil {
log.Fatal(err)
}
pprof.StartCPUProfile(f)
}


func StopCPUProfile() {
pprof.StopCPUProfile()
}

func StartBlockProfile(rate int) {
runtime.SetBlockProfileRate(rate)
}

func StopBlockProfile() {
filename := "block-" + strconv.Itoa(pid) + ".pprof"
f, err := os.Create(filename)
if err != nil {
log.Fatal(err)
}
if err = pprof.Lookup("block").WriteTo(f, 0); err != nil {
log.Fatalf(" can't write %s: %s", filename, err)
}
f.Close()
}


func SetMemProfileRate(rate int) {
runtime.MemProfileRate = rate
}

func GC() {
runtime.GC()
}

func DumpHeap() {
filename := "heap-" + strconv.Itoa(pid) + "-" + strconv.Itoa(int(atomic.AddInt32(&heapProfileCounter, 1))) + ".pprof"
f, err := os.Create(filename)
if err != nil {
fmt.Fprintf(os.Stderr, "testing: %s", err)
return
}
if err = pprof.WriteHeapProfile(f); err != nil {
fmt.Fprintf(os.Stderr, "testing: can't write %s: %s", filename, err)
}
f.Close()
}

func showSystemStat(interval time.Duration, count int) {
usage1 := &syscall.Rusage{}
var lastUtime int64
var lastStime int64
counter := 0
for {
//http://man7.org/linux/man-pages/man3/vtimes.3.html
syscall.Getrusage(syscall.RUSAGE_SELF, usage1)
utime := usage1.Utime.Sec*1000000000 + usage1.Utime.Usec
stime := usage1.Stime.Sec*1000000000 + usage1.Stime.Usec
userCPUUtil := float64(utime-lastUtime) * 100 / float64(interval)
sysCPUUtil := float64(stime-lastStime) * 100 / float64(interval)
memUtil := usage1.Maxrss * 1024
lastUtime = utime
lastStime = stime
if counter > 0 {
fmt.Printf("cpu: %3.2f%% us  %3.2f%% sy, mem:%s \n", userCPUUtil, sysCPUUtil, toH(uint64(memUtil)))
}
counter += 1
if count >= 1 && count < counter {
return
}
time.Sleep(interval)
}
}
func ShowSystemStat(seconds int) {
go func() {
interval := time.Duration(seconds) * time.Second
showSystemStat(interval, 0)
}()
}
func PrintSystemStats() {
interval := time.Duration(1) * time.Second
showSystemStat(interval, 1)
}

func ShowGCStat() {
go func() {
var numGC int64
interval := time.Duration(100) * time.Millisecond
gcstats := &debug.GCStats{PauseQuantiles: make([]time.Duration, 100)}
memStats := &runtime.MemStats{}
for {
debug.ReadGCStats(gcstats)
if gcstats.NumGC > numGC {
runtime.ReadMemStats(memStats)
printGC(memStats, gcstats)
numGC = gcstats.NumGC
}
time.Sleep(interval)
}
}()
}

func PrintGCSummary() {
memStats := &runtime.MemStats{}
runtime.ReadMemStats(memStats)
gcstats := &debug.GCStats{PauseQuantiles: make([]time.Duration, 100)}
debug.ReadGCStats(gcstats)
printGC(memStats, gcstats)
}
func printGC(memStats *runtime.MemStats, gcstats *debug.GCStats) {
if gcstats.NumGC > 0 {
lastPause := gcstats.Pause[0]
elapsed := time.Now().Sub(startTime)
overhead := float64(gcstats.PauseTotal) / float64(elapsed) * 100
allocatedRate := float64(memStats.TotalAlloc) / elapsed.Seconds()
fmt.Printf("NumGC:%d Pause:%s Pause(Avg):%s Overhead:%3.2f%% Alloc:%s Sys:%s Alloc(Rate):%s/s Histogram:%s %s %s \n",
gcstats.NumGC,
toS(lastPause),
toS(avg(gcstats.Pause)),
overhead,
toH(memStats.Alloc),
toH(memStats.Sys),
toH(uint64(allocatedRate)),
toS(gcstats.PauseQuantiles[94]),
toS(gcstats.PauseQuantiles[98]),
toS(gcstats.PauseQuantiles[99]))
} else {
// while GC has disabled
elapsed := time.Now().Sub(startTime)
allocatedRate := float64(memStats.TotalAlloc) / elapsed.Seconds()
fmt.Printf("Alloc:%s Sys:%s Alloc(Rate):%s/s\n",
toH(memStats.Alloc),
toH(memStats.Sys),
toH(uint64(allocatedRate)))
}
}
func avg(items []time.Duration) time.Duration {
var sum time.Duration
for _, item := range items {
sum += item
}
return time.Duration(int64(sum) / int64(len(items)))
}
// human readable format
func toH(bytes uint64) string {
switch {
case bytes < 1024:
return fmt.Sprintf("�", bytes)
case bytes < 1024*1024:
return fmt.Sprintf("%.2fK", float64(bytes)/1024)
case bytes < 1024*1024*1024:
return fmt.Sprintf("%.2fM", float64(bytes)/1024/1024)
default:
return fmt.Sprintf("%.2fG", float64(bytes)/1024/1024/1024)
}
}
// short string format
func toS(d time.Duration) string {
u := uint64(d)
if u < uint64(time.Second) {
switch {
case u == 0:
return "0"
case u < uint64(time.Microsecond):
return fmt.Sprintf("%.2fns", float64(u))
case u < uint64(time.Millisecond):
return fmt.Sprintf("%.2fus", float64(u)/1000)
default:
return fmt.Sprintf("%.2fms", float64(u)/1000/1000)
}
} else {
switch {
case u < uint64(time.Minute):
return fmt.Sprintf("%.2fs", float64(u)/1000/1000/1000)
case u < uint64(time.Hour):
return fmt.Sprintf("%.2fm", float64(u)/1000/1000/1000/60)
default:
return fmt.Sprintf("%.2fh", float64(u)/1000/1000/1000/60/60)
}
}
}
</pre>

### 分布式系统中的二阶段提交事务
准备阶段 提交阶段

1.准备阶段：事务协调者(事务管理器)给每个参与者(资源管理器)发送Prepare消息，每个参与者要么直接返回失败(如权限验证失败)，要么在本地执行事务，写本地的redo和undo日志，但不提交，到达一种“万事俱备，只欠东风”的状态。(关于每一个参与者在准备阶段具体做了什么目前我还没有参考到确切的资料，但是有一点非常确定：参与者在准备阶段完成了几乎所有正式提交的动作，有的材料上说是进行了“试探性的提交”，只保留了最后一步耗时非常短暂的正式提交操作给第二阶段执行。)

2.提交阶段：如果协调者收到了参与者的失败消息或者超时，直接给每个参与者发送回滚(Rollback)消息；否则，发送提交(Commit)消息；参与者根据协调者的指令执行提交或者回滚操作，释放所有事务处理过程中使用的锁资源。(注意:必须在最后阶段释放锁资源)
###ML
####简单介绍
1. 模式识别
模式识别=机器学习
2. 数据挖掘
数据挖掘=机器学习+数据库
3. 统计学习
统计学习近似等于机器学习
4. 计算机视觉
计算机视觉=图像处理+机器学习
5. 语音识别
语音识别=语音处理+机器学习
6. 自然语言处理
自然语言处理=文本处理+机器学习

####常用方法
1、回归算法
2、神经网络
3、SVM（支持向量机）
4、聚类算法
5、降维算法
6、推荐算法
7、其他
除了以上算法之外，机器学习界还有其他的如高斯判别，朴素贝叶斯，决策树等等算法。但是上面列的六个算法是使用最多，
影响最广，种类最全的典型。机器学习界的一个特色就是算法众多，发展百花齐放。


监督学习算法：
线性回归，逻辑回归，神经网络，SVM
无监督学习算法：
聚类算法，降维算法
特殊算法：
推荐算法

###Nginx
nginx -t 检测配置文件错误

###Linux根据进程号PID找到对应程序文件所在的目录
ps -ef|grep '' //程序名称
//得到进程号PID,比如说PID:2333

cd /proc/PID 比如这个 cd /proc/2333

进到这里之后，执行

ls -ail  // 找到  exe->****

//根据端口号获取程序
netstat -anp|grep 8080




fdisk –l    //查看硬盘分区情况

df –h //查看当前硬盘使用情况

cat /proc/cpuinfo  //查看服务器硬件信息

dmidecode -t 4   //查看CPU信息
### 字符编码
在计算机内存中，统一使用Unicode编码，当需要保存到硬盘或者需要传输的时候，就转换为UTF-8编码。

用记事本编辑的时候，从文件读取的UTF-8字符被转换为Unicode字符到内存里，编辑完成后，保存的时候再把Unicode转换为UTF-8保存到文件。

### 记一次https配置问题
如果你的Apple设备无法访问https网址，那么很有可能你的nginx配置出了问题：
<pre>
	#添加这个
	ssl_ciphers HIGH:!aNULL:!MD5;
	ssl_prefer_server_ciphers on;
</pre>

### 可能是下一代的互联网:)
凑个热闹：

http://127.0.0.1:43110/1F7L7DZNeGNMWBCux9zjQPu3YRdtzQToKG

### 打开或者关闭端口
关闭端口
iptables -A INPUT -p tcp --dport 111 -j DROP
打开端口
iptables -A INPUT -p tcp --dport 111 -j ACCEPT

### vscode打开指定文件
control+p 

### RPC还是REST 
RPC更偏向内部调用，REST更偏向外部调用。所以中国的技术圈子更倡导RPC，比如阿里开源的dubbo。美国的技术圈子更倡导REST，比如spring cloud，是个纯REST的项目，不支持RPC。大概是美国的技术圈，保留的初心多那么一点点吧 ; 如果你的系统很复杂，用RPC就要小心地去控制复杂度了，用REST反而会简单些 ;通过RPC能解耦服务，这才是使用RPC的真正目的。通过RPC能解耦服务，这才是使用RPC的真正目的;一个高性能RPC框架最重要的四个点就是：传输协议，框架线程模型，IO模型，零拷贝

### 云风-翻墙工具
xtunnel.go 这个程序运行在本地
<pre>
package main

import "net"
import "log"
import "container/list"
import "io"

import "sync"

const bindAddr = "127.0.0.1:1080"
const serverAddr = "www.yourvps.com:2011"
const bufferSize = 4096
const maxConn = 0x10000
const xor = 0x64

type tunnel struct {
	id int
	*list.Element
	send  chan []byte
	reply io.Writer
}

type bundle struct {
	t [maxConn]tunnel
	*list.List
	*xsocket
	sync.Mutex
}

type xsocket struct {
	net.Conn
	*sync.Mutex
}

func (s xsocket) Read(data []byte) (n int, err error) {
	n, err = io.ReadFull(s.Conn, data)
	if n > 0 {
		for i := 0; i < n; i++ {
			data[i] = data[i] ^ xor
		}
	}

	return
}

func (s xsocket) Write(data []byte) (n int, err error) {
	s.Lock()
	defer s.Unlock()
	log.Println("Send", len(data))
	for i := 0; i < len(data); i++ {
		data[i] = data[i] ^ xor
	}
	x := 0
	all := len(data)

	for all > 0 {
		n, err = s.Conn.Write(data)
		if err != nil {
			n += x
			return
		}
		if all != n {
			log.Println("Write only", n, all)
		}
		all -= n
		x += n
		data = data[n:]
	}

	return all, err
}

func (t *tunnel) processBack(c net.Conn) {
	//c.SetReadTimeout(1e7) 原版中有的
	var buf [bufferSize]byte
	for {
		n, err := c.Read(buf[4:])
		if n > 0 {
			t.sendBack(buf[:4+n])
		}
		e, ok := err.(net.Error)
		if !(ok && e.Timeout()) && err != nil {
			log.Println(n, err)
			return
		}
	}
}

func (t *tunnel) sendClose() {
	buf := [4]byte{
		byte(t.id >> 8),
		byte(t.id & 0xff),
		0,
		0,
	}
	t.reply.Write(buf[:])
}

func (t *tunnel) sendBack(buf []byte) {
	buf[0] = byte(t.id >> 8)
	buf[1] = byte(t.id & 0xff)
	length := len(buf) - 4
	buf[2] = byte(length >> 8)
	buf[3] = byte(length & 0xff)
	t.reply.Write(buf)
}

func (t *tunnel) process(c net.Conn, b *bundle) {
	go t.processBack(c)
	send := t.send

	for {
		buf, ok := <-send
		if !ok {
			c.Close()
			return
		}
		n, err := c.Write(buf)
		if err != nil {
			b.free(t.id)
		} else if n != len(buf) {
			log.Println("Write", n, len(buf))
		}
	}
}

func (t *tunnel) open(b *bundle, c net.Conn) {
	t.send = make(chan []byte)
	t.reply = b.xsocket
	go t.process(c, b)
}

func (t *tunnel) close() {
	close(t.send)
}

func newBundle(c net.Conn) *bundle {
	b := new(bundle)
	b.List = list.New()
	for i := 0; i < maxConn; i++ {
		t := &b.t[i]
		t.id = i
		t.Element = b.PushBack(t)
	}
	b.xsocket = &xsocket{c, new(sync.Mutex)}
	return b
}

func (b *bundle) alloc(c net.Conn) *tunnel {
	b.Lock()
	defer b.Unlock()
	f := b.Front()
	if f == nil {
		return nil
	}
	t := b.Remove(f).(*tunnel)
	t.Element = nil
	t.open(b, c)
	return t
}

func (b *bundle) free(id int) {
	b.Lock()
	defer b.Unlock()
	t := &b.t[id]
	if t.Element == nil {
		t.sendClose()
		t.Element = b.PushBack(t)
		t.close()
	}
}

func (b *bundle) get(id int) *tunnel {
	b.Lock()
	defer b.Unlock()
	t := &b.t[id]
	if t.Element != nil {
		return nil
	}
	return t
}

func servSocks(b *bundle) {
	a, err := net.ResolveTCPAddr("tcp", bindAddr)
	if err != nil {
		log.Fatal(err)
	}
	l, err2 := net.ListenTCP("tcp", a)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("xtunnelc bind %s", a)
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(c.RemoteAddr())
		b.alloc(c)
	}
}

func mainServer(c net.Conn) {
	b := newBundle(c)
	go servSocks(b)
	var header [4]byte
	for {
		_, err := b.Read(header[:])
		if err != nil {
			log.Fatal(err)
		}
		id := int(header[0])<<8 | int(header[1])
		length := int(header[2])<<8 | int(header[3])
		log.Println("Recv", id, length)
		if length == 0 {
			b.free(id)
		} else {
			t := b.get(id)
			buf := make([]byte, length)
			n, err := b.Read(buf)
			if err != nil {
				log.Fatal(err)
			} else if n != len(buf) {
				log.Println("Read", n, len(buf))
			}
			if t != nil {
				t.send <- buf
			}
		}
	}
}

func start(addr string) {
	a, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	c, err2 := net.DialTCP("tcp", nil, a)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("xtunnelc connect %s", a)
	mainServer(c)
}

func main() {
	start(serverAddr)
}

</pre>
xtunneld.go 这个程序运行在墙外
<pre>
package main

import "net"
import "log"
import "container/list"
import "io"

import "sync"

const socksServer = "127.0.0.1:1080"
const bindAddr = ":2011"
const bufferSize = 4096
const maxConn = 0x10000
const xor = 0x64

var socksAddr *net.TCPAddr

func init() {
	_, err := net.ResolveTCPAddr("tcp", socksServer)
	if err != nil {
		log.Fatal(err)
	}
}

type tunnel struct {
	id int
	*list.Element
	send  chan []byte
	reply io.Writer
}

type bundle struct {
	t [maxConn]tunnel
	*list.List
	*xsocket
}

type xsocket struct {
	net.Conn
	*sync.Mutex
}

func (s xsocket) Read(data []byte) (n int, err error) {
	n, err = io.ReadFull(s.Conn, data)
	if n > 0 {
		for i := 0; i < n; i++ {
			data[i] = data[i] ^ xor
		}
	}

	return
}

func (s xsocket) Write(data []byte) (n int, err error) {
	s.Lock()
	defer s.Unlock()
	log.Println("Send", len(data))
	for i := 0; i < len(data); i++ {
		data[i] = data[i] ^ xor
	}
	x := 0
	all := len(data)

	for all > 0 {
		n, err = s.Conn.Write(data)
		if err != nil {
			n += x
			return
		}
		all -= n
		x += n
		data = data[n:]
	}

	return all, err
}

func (t *tunnel) processBack(c net.Conn) {
	//c.SetReadTimeout(1e7)  原版中有的
	var buf [bufferSize]byte
	for {
		n, err := c.Read(buf[4:])
		if n > 0 {
			t.sendBack(buf[:4+n])
		}
		e, ok := err.(net.Error)
		if !(ok && e.Timeout()) && err != nil {
			log.Println(n, err)
			return
		}
	}
}

func (t *tunnel) sendClose() {
	buf := [4]byte{
		byte(t.id >> 8),
		byte(t.id & 0xff),
		0,
		0,
	}
	t.reply.Write(buf[:])
}

func (t *tunnel) sendBack(buf []byte) {
	buf[0] = byte(t.id >> 8)
	buf[1] = byte(t.id & 0xff)
	length := len(buf) - 4
	buf[2] = byte(length >> 8)
	buf[3] = byte(length & 0xff)
	t.reply.Write(buf)
}

func connectSocks() net.Conn {
	c, err := net.DialTCP("tcp", nil, socksAddr)
	if err != nil {
		return nil
	}
	log.Println(c.RemoteAddr())
	return c
}

func (t *tunnel) process() {
	c := connectSocks()
	if c == nil {
		t.sendClose()
	} else {
		go t.processBack(c)
	}
	send := t.send

	for {
		buf, ok := <-send
		if !ok {
			if c != nil {
				c.Close()
			}
			return
		}
		if c != nil {
			n, err := c.Write(buf)
			if err != nil {
				log.Println("tunnel", n, err)
				t.sendClose()
			}
		}
	}
}

func (t *tunnel) open(reply io.Writer) {
	t.send = make(chan []byte)
	t.reply = reply
	go t.process()
}

func (t *tunnel) close() {
	close(t.send)
}

func newBundle(c net.Conn) *bundle {
	b := new(bundle)
	b.List = list.New()
	for i := 0; i < maxConn; i++ {
		t := &b.t[i]
		t.id = i
		t.Element = b.PushBack(t)
	}
	b.xsocket = &xsocket{c, new(sync.Mutex)}
	return b
}

func (b *bundle) free(id int) {
	t := &b.t[id]
	if t.Element == nil {
		t.Element = b.PushBack(t)
		t.close()
	}
}

func (b *bundle) get(id int) *tunnel {
	t := &b.t[id]
	if t.Element != nil {
		b.Remove(t.Element)
		t.Element = nil
		t.open(b.xsocket)
	}
	return t
}

func servTunnel(c net.Conn) {
	b := newBundle(c)
	var header [4]byte
	for {
		_, err := b.Read(header[:])
		if err != nil {
			log.Fatal(err)
		}
		id := int(header[0])<<8 | int(header[1])
		length := int(header[2])<<8 | int(header[3])
		log.Println("Recv", id, length)
		if length == 0 {
			b.free(id)
		} else {
			t := b.get(id)
			buf := make([]byte, length)
			_, err := b.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			t.send <- buf
		}
	}
}

func start(addr string) {
	a, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	l, err2 := net.ListenTCP("tcp", a)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("xtunneld bind %s", a)
	c, err3 := l.Accept()
	if err3 != nil {
		log.Fatal(err3)
	}
	l.Close()
	servTunnel(c)
}

func main() {
	start(bindAddr)
}

</pre>

### ssh 端口转发
<pre>
 实用例子

有A,B,C 3台服务器, A,C有公网IP, B是某IDC的服务器无公网IP. A通过B连接C的80端口(A<=>B<=>C), 那么在B上执行如下命令即可: 

$ ssh -CfNg -L 6300:127.0.0.1:80 userc@C
$ ssh -CfNg -R 80:127.0.0.1:6300 usera@A

服务器A和服务器C之间, 利用跳板服务器B建立了加密隧道. 在A上连接127.0.0.1:80, 就等同C上的80端口. 需要注意的是, 服务器B上的6300端口的数据没有加密, 可被监听, 例: 

# tcpdump -s 0-i lo port 6300


既然SSH可以传送数据，那么我们可以让那些不加密的网络连接，全部改走SSH连接，从而提高安全性。假定我们要让8080端口的数据，都通过SSH传向远程主机，命令就这样写：
ssh -D 8080 user@host

嗯，这里自己搭建了一个自用的梯子：
//1.在本机执行
ssh -f -N -D 127.0.0.1:7079 root@xxxxxx.com
//2.在火狐浏览器配置中配置 设置-选项-高级-网络-连接-手动配置代理-SOCKS主机  记得将 【使用 SOCKS v5 代理 DNS】勾上
//3.在中转服务器【一般是云主机】 修改ssh config
vi  /etc/ssh/ssh_config //在最后加上

AllowTcpForwarding yes
GatewayPorts       yes

//4.ok


基本参数解释：
-f   输入密码后进入后台模式
-N   不执行远程命令,用于端口转发 
-D   socket5代理
-L   tcp转发  

如果想实现sock5全局代理，可以参考下面：
#!/bin/bash
# Create new chain
iptables -t nat -N REDSOCKS
 
# Ignore LANs and some other reserved addresses.
iptables -t nat -A REDSOCKS -d 0.0.0.0/8 -j RETURN
iptables -t nat -A REDSOCKS -d 10.0.0.0/8 -j RETURN
iptables -t nat -A REDSOCKS -d 127.0.0.0/8 -j RETURN
iptables -t nat -A REDSOCKS -d 169.254.0.0/16 -j RETURN
iptables -t nat -A REDSOCKS -d 172.16.0.0/12 -j RETURN
iptables -t nat -A REDSOCKS -d 192.168.0.0/16 -j RETURN
iptables -t nat -A REDSOCKS -d 224.0.0.0/4 -j RETURN
iptables -t nat -A REDSOCKS -d 240.0.0.0/4 -j RETURN
 
# Anything else should be redirected to port 31338
iptables -t nat -A REDSOCKS -p tcp -j REDIRECT --to-ports 31338
 
# Any tcp connection made by `linuxaria' should be redirected, put your username here.
iptables -t nat -A OUTPUT -p tcp -m owner --uid-owner linuxaria -j REDSOCKS
此配置实现了把所有除本地局域网连接以外的TCP连接全部转发到 31338 端口, 显然你应该用代理软件提前监听这个端口, 当然也可以是其他任意指定的代理监听端口.

//列出所有正在监听的端口以及程序名称
netstat -lntp

linux ~ 与 /

~代表你的/home/用户明目录
假设你的用户名是x，那么~/就是/home/x/

linux 中的 buffer 与 cache 

buffer[缓冲]是用于存放要输出到disk（块设备）的数据的，而cache[缓存]是存放从disk上读出的数据。这二者是为了提高IO性能的，并由OS管理。

-/+ buffers/cache:   10321516   14355944

即-buffers/cache，表示一个应用程序认为系统被用掉多少内存，这里是 10321516 KB；
即+buffers/cache，表示一个应用程序认为系统还有多少内存，这里是 14355944 KB；

Linux文件系统之inode

理解inode，要从文件储存说起。

文件储存在硬盘上，硬盘的最小存储单位叫做"扇区"（Sector）。每个扇区储存512字节（相当于0.5KB）。

操作系统读取硬盘的时候，不会一个个扇区地读取，这样效率太低，而是一次性连续读取多个扇区，即一次性读取一个"块"（block）。这种由多个扇区组成的"块"，是文件存取的最小单位。"块"的大小，最常见的是4KB，即连续八个 sector组成一个 block。

文件数据都储存在"块"中，那么很显然，我们还必须找到一个地方储存文件的元信息，比如文件的创建者、文件的创建日期、文件的大小等等。这种储存文件元信息的区域就叫做inode，中文译名为"索引节点"。

每一个文件都有对应的inode，里面包含了与该文件有关的一些信息。

inode也会消耗硬盘空间，所以硬盘格式化的时候，操作系统自动将硬盘分成两个区域。一个是数据区，存放文件数据；另一个是inode区（inode table），存放inode所包含的信息。

1、一个Inode对应一个文件，而一个文件根据其大小，会占用多块blocks。
2、更为准确的来说，一个文件只对应一个Inode。因为硬链接其实不是创建新文件，只是在Directory中写入了新的对应关系而已。
3、当我们删除文件的时候，只是把Inode标记为可用，文件在block中的内容是没有被清除的，只有在有新的文件需要占用block的时候，才会被覆盖。
</pre>

### 关于期权基本常识
期权中的陷阱

我们来看看青蛙和凤凰这过程中期权问题：

凤凰御姐许诺给青蛙的未来股权，叫做期权。期权都有附加条件，凤凰御姐给的是4年，每年兑现四分之一。
触发条件到达，掏钱将期许变为现实股权，并签订股权协定，叫行权。只是这部分股权通常有附加条件，比如御姐代持。
行权后的股权如何变现？卖给谁、多少钱、谁定价？一开始青蛙就不清楚，最后只能懵逼。

CEO在期权上坑员工，通常是许诺期权的时候：

告诉你多少股，但不说占公司股份百分比，这样以后他可以随意变更总股数。不是股份制公司，没有在工商登记股份数的，讲股份数是没有法律依据的。
说行权价格到时候董事会订，到时候再说。这就变成了他想订多少就多少；很多初创公司这点倒好，一般都是白送。

不说行权后的股权关键附加条款，比如：你离开公司是否可以带走？如果不能，到时候按什么价格收回？反正到时候他随便定价，而且能一口咬定这是董事会定的，把董事会说的跟临时工一样，董事会就变成了背锅侠。

如果大家遇到这些问题，可以翻翻自己的期权协议，仔细看看，也可以网上找律师帮你把把关，看看是否有明显不合理问题。

陷阱背后的算盘

不合理的期权协议背后无非两种情况：

* 第一种，CEO不清楚期权发放知识，也没钱找专业律师，尤其初创公司。这个情有可原，而且我可以提供一个比较合理的方案建议给他。
* 第二种，连续创业的，法律知识丰富的。这就是诚心坑人的。

如果是第二种，那么CEO在招聘CTO之初就没有觉得CTO多么重要，也没有当做长期合作伙伴，已经在打主意怎么坑这个CTO了。试问这种情况下CTO能长期帮助公司发展么？这种早晚会出问题。

一个好的期权设计

好的许诺期权应该在一开始就说好：

1. 占公司股权比例数
2. 行权条件，比如技术人员通常按年，销售人员通常按业绩
3. 行权价格
4. 兑换为股权后股权的退出机制，是否不允许带走？是否需要强制收回
5. 如果定了强制收回，就需要说明退出价格计算方法

我总结这些年的各种经验，给老板们建议：

1. 创业公司，第一期将10%股权做为期权，用于人才招揽，并且将这部分作为期权池的股权，放到一个持股公司，持股公司到上海崇明岛、新疆等有优惠政策地方注册。
2. 公司CEO同时是这个持股公司的法人、总经理，代表该股份在母公司的投票权。
3. 许诺给员工的期权，每年到期行权。行权后将CEO自己在持股公司的股份转给员工。比如母公司员工应该兑现1%期权，就相当于是在持股公司给员工转股10%，并变更持股公司的工商登记，保障员工利益。
4. 员工离职，没有兑现的期权要清空，已经兑现的股权要强行收回。
5. 收回价格按：员工的股权比例* 公司价值；公司价值按：max（公司最近一轮融资的0.3到0.5，公司上一年总现金回款，工商注册资本）。
6. 收回的的期权放回期权池（持股公司，CEO持有），激励未来的核心员工。
7. 期权交换过程中的税金，谁受益，谁帮CEO出。

简称创业公司期权七铁律。

### Golang 使用通道实现常规锁的功能
<pre>
package main

import "fmt"

type myMap interface {
    push(key string, e interface {}) interface{} 
    pop(key string) interface{}
}

type myMapPair struct {
    key string
    value interface {}
}

type mapChan struct {
    push_req chan * myMapPair
    push_rep chan interface{}
    pop_req chan string
    pop_rep chan interface{}
}

func (c *mapChan) push(key string, e interface{}) interface{} {
    c.push_req <- & myMapPair {key,e}
    return <- c.push_rep
}

func (c *mapChan) pop(key string) interface {} {
    c.pop_req <- key
    return <- c.pop_rep
}

func newMap() myMap {
    c := mapChan { 
        push_req : make (chan * myMapPair),
        push_rep : make (chan interface{}),
        pop_req : make (chan string),
        pop_rep : make (chan interface{}),
    }
    m := make(map[string] interface {})
    go func() {
        for {
            select {
            case r := <- c.push_req :
                if v , exist := m[r.key] ; exist {
                    c.push_rep <- v
                } else {
                    m[r.key] = r.value
                    c.push_rep <- nil
                }
            case r := <- c.pop_req:
                if v,exist := m[r] ; exist {
                    m[r] = nil
                    c.pop_rep <- v
                } else {
                    c.pop_rep <- nil
                }
            }
        }
    } ()
    return &c   
}

func main() {
    m := newMap()
    fmt.Println(m.push("hello","world"))
    fmt.Println(m.push("hello","world"))
    fmt.Println(m.pop("hello"))
    fmt.Println(m.pop("hello"))
}
</pre>
### SQL语句监测优化工具
https://github.com/Meituan-Dianping/SQLAdvisor
###  zabbix gitlab
启动zabbix命令
 
/usr/sbin/zabbix_agentd -c /etc/zabbix/zabbix_agentd.conf

启动 gitlab 命令

/usr/bin/gitlab-ctl start

github提交不显示绿点问题

git config user.email //对比两个邮箱是否相同，如果不相同就使用命令 ：git config --global user.email  DONE!!!
### Abandon
  https://github.com/exacity/deeplearningbook-chinese/
### TCP UDP 
TCP和UDP是OSI模型中的运输层中的协议。TCP提供可靠的通信传输，而UDP则常被用于让广播和细节控制交给应用的通信传输。TCP与UDP基本区别：

1. 基于连接与无连接。
2. TCP要求系统资源较多，UDP较少。
3. UDP程序结构较简单。
4. 流模式（TCP）与数据报模式(UDP)。
5. TCP保证数据正确性，UDP可能丢包。
6. TCP保证数据顺序，UDP不保证。
7. TCP面向连接（如打电话要先拨号建立连接）;UDP是无连接的，即发送数据之前不需要建立连接。
8. TCP提供可靠的服务。也就是说，通过TCP连接传送的数据，无差错，不丢失，不重复，且按序到达;UDP尽最大努力交付，即不保证可靠交付。
9. TCP面向字节流，实际上是TCP把数据看成一连串无结构的字节流;UDP是面向报文的，UDP没有拥塞控制，因此网络出现拥塞不会使源主机的发送速率降低（对实时应用很有用，如IP电话，实时视频会议等）。
10. 每一条TCP连接只能是点到点的;UDP支持一对一，一对多，多对一和多对多的交互通信。
11. TCP首部开销20字节;UDP的首部开销小，只有8个字节。
12. TCP的逻辑通信信道是全双工的可靠信道，UDP则是不可靠信道。

TCP（Transmission Control Protocol 传输控制协议）是一种面向连接的、可靠的、基于字节流的传输层通信协议，由IETF的RFC 793定义。在简化的计算机网络OSI模型中，完成第四层传输层所指定的功能。

UDP 是User Datagram Protocol的简称， 中文名是用户数据报协议，是OSI（Open System Interconnection，开放式系统互联） 参考模型中一种无连接的传输层协议，提供面向事务的简单不可靠信息传送服务，IETF RFC 768是UDP的正式规范。UDP在IP报文的协议号是17。
### 查看程序依赖包与依赖关系
<pre>
//以nginx为例
ldd $(which /usr/local/nginx/sbin/nginx)
</pre>
//X-Forwarded-For:简称XFF头，它代表客户端，也就是HTTP的请求端真实的IP，只有在通过了HTTP 代理或者负载均衡服务器时才会添加该项。

nginx.conf中的 proxy_set_header Remoteip $proxy_add_x_forwarded_for;

ECS中的nginx访问不了不一定是服务器问题，很有可能是安全组设置没有设定。
### nginx 配置示例
<pre>
#user  example;   ＃使用的用户和组
worker_processes  8; 
#启动进程数，可根据机器核数来修改 一般等于cpu总核数或总核数的两倍，两个四核cpu，总核数为8
#指定错误日志存放的路径，错误日志记录级别可选项为debug info notice warn error crit
#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;
error_log   /usr/local/nginx/logs/error.log debug;
#指定pid存放的路径
#pid        logs/nginx.pid;
#指定文件描述符数量
worker_rlimit_nofile 65535;
events {
#使用的i/o模型，linux系统推荐采用epoll模型
    use epoll;
#工作模式及连接数上限
    worker_connections  1024;  
}

http {
     #设定http服务器，利用它的反向代理功能提供负载均衡支持
    include       mime.types;   #文件扩展名与文件类型映射表
    default_type  application/octet-stream;  #默认的文件类型
     #设置使用的字符集，如果一个网站有多种字符集，不要轻易设置
     #charset gb2312;
    #server_names_hash_bush_bucket_sizp 128;  #服务器名字的hash名大小?
    #nginx默认会用client_header_buffer_size这个buffer来读取header值，如果header过大，它会使用large_client_header_buffers来读取
    #client_header_buffer_size 32k;
    #large_client_header_buffers 4 32k;
    #log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
    #                  '$status $body_bytes_sent "$http_referer" '
    #                  '"$http_user_agent" "$http_x_forwarded_for"';
    #access_log  logs/access.log  main;  #定义的日志格式和存放路径
    #设置客户端能够上传的文件大小
    client_max_body_size  8m;
    sendfile        on;  #开启高效文件传输模式
    #tcp_nopush     on;  #防止网络阻塞
    #keepalive_timeout  0;
    keepalive_timeout  65; #超时时间
    #开启gzip压缩
    #gzip  on;
    server {
        listen       80;
        server_name  localhost;
        #charset koi8-r;
        location / {
            #root   html;   wj
            root /usr/local/nginx/html;
            #index  index.html index.htm;  wj
            index index.php index.html index.htm;
        }
 
        ##对常见格式文件，在浏览器本地缓存的天数
        location ~*^.+.(jpg|jpeg|gif|css|png|js|ico|xml)$
        {
            access_log off;
            expires 30d;
            root /usr/local/nginx/html;
        }
        #error_page  404              /404.html;
        # redirect server error pages to the static page /50x.html
        #
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
           # root   html; wj
             root /usr/local/nginx/html;
        }
        # proxy the PHP scripts to Apache listening on 127.0.0.1:80
        #
        #location ~ \.php$ {
        #    proxy_pass   http://127.0.0.1;
        #}
        # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
        #
        #location ~ \.php$ {  
        #    root           html; 
        #    fastcgi_pass   127.0.0.1:9000;  
        #    fastcgi_index  index.php; 
        #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name; 
        #    include        fastcgi_params; 
        #}  
        ##Parse all.php file in the /var/www/nginx-default directory
        location ~\.php$
        {
             fastcgi_split_path_info ^(.+\.php)(.*)$;
             fastcgi_pass 127.0.0.1:9000;
             fastcgi_index index.php;
             fastcgi_param SCRIPT_FILENAME /usr/local/nginx/html$fastcgi_script_name;
             include fastcgi_params;
             fastcgi_param QUERY_STRING $query_string;
             fastcgi_param REQUEST_METHOD $request_method;
             fastcgi_param CONTENT_TYPE $content_type;
             fastcgi_param CONTENT_LENGTH $content_length;
             fastcgi_intercept_errors on;
             fastcgi_ignore_client_abort off;
             fastcgi_connect_timeout 60; #连接超时时间
             fastcgi_send_timeout 180; #发送超时时间
             fastcgi_read_timeout 180; #读取超时时间
             fastcgi_buffer_size 128k; #设置FastCGI服务器相应头部的缓存区大小
             fastcgi_buffers 4 256k; #设置FastCGI进程返回信息的缓存区数量的大小

             fastcgi_busy_buffers_size 256k;
             fastcgi_temp_file_write_size 256k;
        }
        # deny access to .htaccess files, if Apache's document root
        # concurs with nginx's one
        #
        #location ~ /\.ht {  wj
        location ~ /\.ht {
        #    deny  all;  wj
            deny  all;
        #} wj
        }
    }

    # another virtual host using mix of IP-, name-, and port-based configuration
    #
    #server {
    #    listen       8000;
    #    listen       somename:8080;
    #    server_name  somename  alias  another.alias;
    #    location / {
    #        root   html;
    #        index  index.html index.htm;
    #    }
    #}
}
</pre>
### 服务器快照原理
快照激活后，应用服务器可以对快照卷进行读写操作。应用服务器下发写请求后，数据将直接写入快照卷，并在独享映射表中记录数据在快照卷中的存放位置。<br>
快照卷用的是存储池的空间，一般存储池预留的20%就是给快照卷用的！

### 使用企业微信号发送消息
<pre>
package services

import (
    "bufio"
    "bytes"
    "encoding/json"
    "errors"
    "strings"
    "io/ioutil"
    "net/http"
    "os"
)

const (
    requestError = errors.New("request error,check url or network")
    agentidX = 0
    corpidX = ""
    corpsecretX = ""
    sendurl   = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`
    get_token = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=`
)

type access_token struct {
    Access_token string `json:"access_token"`
    Expires_in   int    `json:"expires_in"`
}

type send_msg struct {
    Touser  string            `json:"touser"`
    Toparty string            `json:"toparty"`
    Totag   string            `json:"totag"`
    Msgtype string            `json:"msgtype"`
    Agentid int               `json:"agentid"`
    Text    map[string]string `json:"text"`
    Safe    int               `json:"safe"`
}

type send_msg_error struct {
    Errcode int    `json:"errcode`
    Errmsg  string `json:"errmsg"`
}

//通过corpid 和 corpsecret 获取token 
func Get_token(corpid, corpsecret string) (at access_token, err error) {
    resp, err := http.Get(get_token + corpid + "&corpsecret=" + corpsecret)
    if err != nil {
        return
    }
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        err = requestError
        return
    }
    buf, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(buf, &at)
    if at.Access_token == "" {
        err = errors.New("corpid or corpsecret error.")
    }
    return
}

func Parse(jsonpath string) ([]byte, error) {
    var zs = []byte("//")
    File, err := os.Open(jsonpath)
    if err != nil {
        return nil, err
    }
    defer File.Close()
    var buf []byte
    b := bufio.NewReader(File)
    for {
        line, _, err := b.ReadLine()
        if err != nil {
            if err.Error() == "EOF" {
                break
            }
            return nil, err
        }
        line = bytes.TrimSpace(line)
        if len(line) <= 0 {
            continue
        }
        index := bytes.Index(line, zs)
        if index == 0 {
            continue
        }
        if index > 0 {
            line = line[:index]
        }
        buf = append(buf, line...)
    }
    return buf, nil
}

//发送消息
func Send_msg(Access_token string, msgbody []byte) error {
	body := bytes.NewBuffer(msgbody)
    resp, err := http.Post(sendurl+Access_token, "application/json", body)
    if resp.StatusCode != 200 {
        return requestError
    }
    buf, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    var e send_msg_error
    err = json.Unmarshal(buf, &e)
    if err != nil {
        return err
	}
    if e.Errcode != 0 && e.Errmsg != "ok" {
        return errors.New(string(buf))
    }
    return nil
}

func QiYeWeiXinSendMsg(touser,message_body string) error {
	if  strings.TrimSpace(touser) ==""{
		touser = "@all"
	}
    var m send_msg = send_msg{Touser:touser, Msgtype: "text", Agentid:agentidX, Text: map[string]string{"content":message_body}}
    token, err := Get_token(corpidX,corpsecretX)
    if err != nil {
        println(err.Error())
        return err
    }
    buf, err := json.Marshal(m)
    if err != nil {
        return err
    }
    err = Send_msg(token.Access_token, buf)
    if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
</pre>
### 数据库备份数据shell脚本
<pre>
#!/bin/bash
#你要修改的地方从这里开始
MYSQL_USER=root     #mysql用户名
MYSQL_PASS=         #mysql密码
FTP_IP=             #远程ftp地址
FTP_USER=           #远程ftp用户名
FTP_PASS=           #远程ftp密码
FTP_backup=         #远程ftp上存放备份文件的目录，需要先在FTP上面建好
WEB_DATA=/home/wwwroot     #本地要备份的网站数据
#你要修改的地方从这里结束
 
if [ ! -f /usr/bin/ftp ]; then
    yum install ftp -y
fi
if [ ! -d /home/backup ]; then
    mkdir /home/backup
fi
  
#定义备份文件的名字
DataBakName=Data_$(date +"%Y%m%d").tar.gz
OldData=Data_$(date -d -5day +"%Y%m%d").tar.gz
 
#删除本地3天前的数据
rm -rf /home/backup/Data_$(date -d -3day +"%Y%m%d").tar.gz
cd /home/backup
  
#导出数据库,一个数据库一个压缩文件
for db in `/usr/local/mysql/bin/mysql -u$MYSQL_USER -p$MYSQL_PASS -B -N -e 'SHOW DATABASES' | xargs`; do
    (/usr/local/mysql/bin/mysqldump -u$MYSQL_USER -p$MYSQL_PASS ${db} -q --skip-lock-tables | gzip -9 - > ${db}.sql.gz;
    echo dumped /home/backup/${db}.sql.gz)    
done
 
#将导出的数据库和网站目录压缩为一个文件
tar zcf /home/backup/$DataBakName $WEB_DATA /home/backup/*.sql.gz

#删除本地已导出的数据库
rm -rf /home/backup/*.sql.gz

#上传到FTP空间,删除FTP空间5天前的数据
ftp -v -n $FTP_IP << END
user $FTP_USER $FTP_PASS
type binary
cd $FTP_backup
delete $OldData
put $DataBakName
bye
END
</pre>
### MySQL distinct
select (distinct id) from table 

delete table_a from table_a,table_b where table_a.id=table_b.id
### 100亿加减法思路
负数在计算机中以补码的形式存储。
负数的补码表示方法是：将负数表示成二进制原码（负数最高位是1，正数最高位是0）然后将原码取反（1变0，0变1），即反码，将反码加1（最后一位上加1），即转化为补码。如用八位
二进制表示-5，第一步，原码10000101，反码01111010，加1变为补码：01111011

首先要明白这道题目的考查点是什么，一是大家首先要对计算机原理的底层细节要清楚、要知道加减法的位运算原理和知道计算机中的算术运算会发生越界的情况，二是要具备一定的面向对象的设计思想。

首先，计算机中用固定数量的几个字节来存储的数值，所以计算机中能够表示的数值是有一定的范围的，为了便于讲解和理解，我们先以byte类型的整数为例，它用1个字节进行存储，表示的最大数值范围为-128到+127。-1在内存中对应的二进制数据为11111111，如果两个-1相加，不考虑Java运算时的类型提升，运算后会产生进位，二进制结果为1,11111110，由于进位后超过了byte类型的存储空间，所以进位部分被舍弃，即最终的结果为11111110，也就是-2，这正好利用溢位的方式实现了负数的运算。-128在内存中对应的二进制数据为10000000，如果两个-128相加，不考虑Java运算时的类型提升，运算后会产生进位，二进制结果为1,00000000，由于进位后超过了byte类型的存储空间，所以进位部分被舍弃，即最终的结果为00000000，也就是0，这样的结果显然不是我们期望的，这说明计算机中的算术运算是会发生越界情况的，两个数值的运算结果不能超过计算机中的该类型的数值范围。由于Java中涉及表达式运算时的类型自动提升，我们无法用byte类型来做演示这种问题和现象的实验，大家可以用下面一个使用整数做实验的例子程序体验一下：
<pre>
	int a = Integer.MAX_VALUE;

	int b = Integer.MAX_VALUE;

	int sum = a + b;

	System.out.println(“a=”+a+”,b=”+b+”,sum=”+sum);
</pre>
先不考虑long类型，由于int的正数范围为2的31次方，表示的最大数值约等于2*1000*1000*1000，也就是20亿的大小，所以，要实现一个一百亿的计算器，我们得自己设计一个类可以用于表示很大的整数，并且提供了与另外一个整数进行加减乘除的功能，大概功能如下：

* 这个类内部有两个成员变量，一个表示符号，另一个用字节数组表示数值的二进制数
* 有一个构造方法，把一个包含有多位数值的字符串转换到内部的符号和字节数组中
* 提供加减乘除的功能

### Linux 上的base64加解密

echo -n "snailwarrior" | base64

c25haWx3YXJyaW9y

echo -n 选项没有输出字符串结尾的' '换行字符，因此字符串"snailwarrior"精确的base64编码是"c25haWx3YXJyaW9y"，可以用php函数来检验哦。文件方式等进行的"snailwarrior"字符串编码都带入了对' '的编码，因此，不小心就会发生莫名的编码错误哦。

Base64解码

echo -n "c25haWx3YXJyaW9y" | base64 -d
### 关于OAuth(Open Authorization)协议
OAuth 本身不存在一个标准的实现，后端开发者自己根据实际的需求和标准的规定实现。其步骤一般如下：

* 客户端要求用户给予授权
* 用户同意给予授权
* 根据上一步获得的授权，向认证服务器请求令牌（token）
* 认证服务器对授权进行认证，确认无误后发放令牌
* 客户端使用令牌向资源服务器请求资源
* 资源服务器使用令牌向认证服务器确认令牌的正确性，确认无误后提供资源

在oauth2.0的流程中，用户登录了第三方的系统后，会先跳去服务方获取一次性用户授权凭据，再跳回来把它交给第三方，第三方的服务器会把授权凭据以及服务方给它的的身份凭据一起交给服务方，这样，服务方一可以确定第三方得到了用户对此次服务的授权（根据用户授权凭据），二可以确定第三方的身份是可以信任的（根据身份凭据），所以，最终的结果就是，第三方顺利地从服务方获取到了此次所请求的服务。

OAuth2.0成员

* Resource Owner（资源拥有者：用户）
* Client （第三方接入平台：请求者）
* Resource Server （服务器资源：数据中心）
* Authorization Server （认证服务器）


第三方客户端的授权模式

客户端必须得到用户的授权（authorization grant），才能获得令牌（access token）。OAuth 2.0定义了四种授权方式。

* 授权码模式（authorization code）
* 简化模式（implicit）
* 密码模式（resource owner password credentials）
* 客户端模式（client credentials）
### Traceroute netcat
* Traceroute

使用 traceroute 查找一个请求都经过了哪些网关
<pre>
traceroute www.baidu.com
</pre>
* netcat 


扫描端口：
        nc -nvvz -w2 127.0.0.1 1-1000
        最基本的功能之一了，扫描目标ip的端口段，然后用-w加上一个超时时间。参数里加上r使得端口扫描变得随机一些，对方log里看起来不那么像是被扫描的（其实还是很容易看出来）。

连接到目标：
        nc -nvv 127.0.0.1 8089
        作为测试客户端非常常用的功能，即连接到目标ip的某个端口上。连接上之后会把stdin的数据发给server。 
        在确定服务端正常开启的情况下连接被拒多半要去检查下防火墙设置。

服务器：
        nc -lvv -p 8089
        让nc作为一个server监听8089端口，把stdin的数据发给client。

传文件：
        其实和nc没啥关系，主要利用了系统IO重定向或pipe的功能，标准化带来的好处多多。
        发送端通过pipe把文件数据传递给stdin，或者通过重定向把stdin重定向到某个文件上，然后接受端只要对stdout进行重定向把它定向到目标文件上去即可。
        比如：
        server: nc -lvv -p 8089 < my.txt
        client: nc -nvv serverip 8089 > my.txt
        当你在纠结两台机子传个文件到底是用网络共享好还是利用中间ftp好又或者是scp甚至于给自己发个邮件另一台机子上去下载的时候，这简直是你的诺亚方舟。
        这个功能说成传文件说的有些狭隘了，总之就是传递数据，理论上可以把接收端的数据（可能是某种格式编码的）直接传给对应的处理软件去执行，举个例子比如说看个视频什么的。

得到对端机子的shell：
        -e的邪恶能力终于派上用场了。
        一台机子上 nc -lvv -p 8089 -t -e cmd.exe
        然后你去连它的时候神奇的事情发生了，自己的终端会变成对方机子的终端，可以随便做邪恶的事情（在用户权限范围下），原理的话估计和我上面写的类似。
        如果对方机子的nc支持-d，那就更理想一点，加上这个参数使得nc后台运行更加难以发现，更进一步的，修改nc.exe的文件名使查看进程的时候别人也不容易起疑。这个NB的功能除了拿来动真刀真枪之外，还能怎么玩其实主要看-e后面的这个用户程序怎么去写吧，这个还是自己去找灵感。

蜜罐：
        还是作为server的基本功能来说，思路不一样会发现用途立刻就大相径庭了。比如说你没事儿listen一个端口，然后-o 来打log，闲来喝喝茶翻翻log，你就能知道都有哪些人对你的这个端口感兴趣（嘿嘿嘿...）
        更进一步的，因为许多小坏蛋们总是用一些安全渗透工具来找空子，毕竟是工具嘛，识别一些漏洞总是有一组傻乎乎的规则的，如果想要捉弄他们让他们误以为真的有漏洞而采取进一步措施的话，可以结合nc的数据传输功能和-e再加上自写程序来伪造一个。不过小心玩火自焚喽。

反向链接：
        前面说一个拿到server shell的情况，要求server端绑定一个shell程序到端口上。那么实际应用情境中，一般就是小坏蛋给小绵羊电脑上种一个server啦，然后每次小坏蛋去连接都能得到shell。但由于种种问题，小坏蛋要直接连到小绵羊电脑可能比较麻烦，尤其是小绵羊电脑在一个局域网中，通过nat上网的情况下。
        为了解决这个问题，可以换一种思路，小绵羊作为主动发起端并自己绑定一个shell程序到nc，要实现这个步骤可能是小坏蛋哪天千辛万苦入了小绵羊电脑设了一个计划任务，也可能是小坏蛋到小绵羊家里玩的时候在小绵羊电脑里使了个坏...然后小坏蛋电脑只需要安心开着server等小绵羊上钩，每次小绵羊连接到server就自动把自己卖了出去...
        原理的话，猜测是在阻塞connect成功之后fork然后exec，再wait，和主动连接差不多的。不过反向链接省去了正面突围的大麻烦，功德一件，善哉善哉。


        看到比较典型的大概有上面这几种，熟悉的人肯定已经熟悉的不能再熟了。通过和各个系统工具的组合想必还有更多好玩的用法的，所以nc可是一个很益智的玩具呢，当然它不仅仅是玩具，必要时也可能帮你解决实际问题。
        综合来看，nc在许多方面的功能都有比较好用的独立软件，这也侧面反应了nc在这些功能上都有许多做的不够好的地方（没有足够多针对性的优化和扩展功能）。但它胜在麻雀虽小但五脏俱全，这一点十分讨人喜欢，在许多没有特点工作环境的场合会是非常有吸引力的选择。
        虽然还没有深度使用这个工具，不过有预感结合自编脚本来用的话，应该会表现的更加出彩一些。

        就像语言影响人的思维面一样，工具的掌握让人更加清晰的认识到何事可为何事不可为，从而进一步影响决策和行为。

        为了进化而努力。

### mysql regexp
<pre>
SELECT * FROM users_bank where  card_no  regexp '^[0-9]+$';  -- 找出所有只是数字      not REGEXP 不只是数字
</pre>
mysql 中 emun 与 tinyint 

1.  enum的存储原理我仔细查看了下手册。是这样的:

在建立这个字段时，我们会给他规定一个范围比如enum('a','b','c'),这时mysql内部会建立一张hash结构的map表，类似:0000 -> a，0001 -> b，0002 -> c。

当我插入一条数据，此字段的值位a或b或c时，他存储在里面的不是这个字符，而是对应他的索引，也就是那个0000或0001或0002。
同样，enum在mysql手册上的说明:
<pre>
ENUM('value1','value2',...)
1或2个字节，取决于枚举值的个数(最多65,535个值)
除非enum的个数超过了一定数量，否则他所占的存储空间也总是1字节。
</pre>
2. tinyint:
<pre>
类型  字节  最小值  最大值
      (带符号的/无符号的)  (带符号的/无符号的)
TINYINT  1  -128  127
他的最小存储所占空间也是1字节。
</pre>
最后，Enum，既然要用它，就不必要使用什么0，1，2来代替实际的字符串了。甚至中文字符串。他并不会对数据库性能进行多余开销。因为对于它来说，你使用'0','1','2'和'张三','李四','王五'数据表所占的存储空间一样。但是考虑到我们实际应用时数据需要从db服务器回传到web app，所以在网络传输时，当然还是尽可能的传输小数据比较好。所以如果很在意这些，还是不用它好了。
### HTTrack - 克隆任意网站
HTTrack可以克隆指定网站－把整个网站下载到本地。

httrack http://163.com -O /tmp/163

### python中的self字段
<pre>

class Test(object):
#类的方法中需要添加self
    def add(self,a,b):
        print(a+b)

    def display(self):
        print("hello")
    

test = Test()
test.add(1,2)
test.display()

#普通方法不需要加self
def addTwo(a,b):
    print(a+b)

addTwo(2,3)

</pre>

### 工作线程

<pre>

//使用工作线程控制数量

package main

import (
	"fmt"
	"time"

	"github.com/xgfone/go-tools/worker"
)

func ExampleDispatcher() {
	JobQueue := make(chan interface{}, 10)
	dispatcher := worker.NewDispatcher(100, worker.FuncTask(func(job interface{}) {
		fmt.Println("Receive job",job)
	}))
	dispatcher.Dispatch(JobQueue)

	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27} {
		JobQueue <- i
	}

	time.Sleep(time.Second)
	dispatcher.Stop()
}
func main(){
	 ExampleDispatcher() 
}
</pre>

### Golang csv转换成 xlsx

<pre>
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/tealeg/xlsx"
)

const (
	bom0 = 0xef
	bom1 = 0xbb
	bom2 = 0xbf
)

// BOMReader returns a Reader that discards the BOM header.
func BOMReader(r io.Reader) io.Reader {
	buf := bufio.NewReader(r)
	b, err := buf.Peek(3)
	if err != nil {
		return buf
	}
	if b[0] == bom0 && b[1] == bom1 && b[2] == bom2 {
		buf.Discard(3)
	}
	return buf
}

func csv2xlsx(csvPath string) {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(BOMReader(csvFile))

	xlsxFile := xlsx.NewFile()
	xlsx.SetDefaultFont(10, "Verdana")

	sheetNo := 0
	sheetName := fmt.Sprintf("Sheet%d", sheetNo)
	sheet, _ := xlsxFile.AddSheet(sheetName)

	lineNo := 0
	fields, err := reader.Read()
	for err == nil {
		lineNo++
		//  a sheet can contain 1048576 rows, 16384 columns.
		if lineNo%1000000 == 0 {
			sheetNo++
			sheetName = fmt.Sprintf("Sheet%d", sheetNo)
			sheet, _ = xlsxFile.AddSheet(sheetName)
		}

		row := sheet.AddRow()
		for _, field := range fields {
			cell := row.AddCell()
			cell.Value = field
		}

		fields, err = reader.Read()
	}

	if err != nil && err != io.EOF {
		fmt.Printf(err.Error())
		return
	}

	fileName := strings.TrimSuffix(filepath.Base(csvPath), ".csv")
	outFile := fileName + ".xlsx"
	xlsxFile.Save(outFile)
}

func showHelp() {
	fmt.Println()
	fmt.Println("-")
	fmt.Println("Usage: csv2xlsx")
	fmt.Println("       csv2xlsx [FILEPATH]")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("     csv2xlsx")
	fmt.Println("       -Convert all csv files in the folder to xlsx.")
	fmt.Println("     csv2xlsx ./test.csv")
	fmt.Println("       -Convert test.csv to xlsx.")
	fmt.Println()
}

func main() {
	showHelp()

	if len(os.Args) > 1 {
		csv2xlsx(os.Args[1])
		return
	}

	csvFiles, _ := filepath.Glob("./*.csv")
	for _, csvFile := range csvFiles {
		fmt.Printf("Processing csv file: %s\n", csvFile)
		csv2xlsx(csvFile)
	}
}

</pre>

### git merge 

* git merge –no-ff 可以保存你之前的分支历史。能够更好的查看 merge历史，以及branch 状态。
* git merge 则不会显示 feature，只保留单条分支记录。

基本使用分支方法：

<pre>
a. 创建develop分支
git branch dev       # 本地新建分支
git push -u origin dev     # 将本地分支推送到远端 

b. 在Dev分支基础上再建立功能分支【次步可忽略】
git checkout -b feature1 dev   # 在本地Dev分支基础上再新建功能分支1，并且切换到该分支下
git push -u origin feature1  # 将本地分支1推送到远端

-- 做一些改动    
git status
git add some-file
git commit   

c. 完成功能开发
git push origin dev # 拉取Dev分支最新代码
git checkout dev  # 从feature1分支切回都Dev分支
git merge --no-ff feature1 # 将本地分支1内容合并到Dev分支
git push origin dev  # 将合并后的Dev分支推到远端

-- 删除功能分支feature1 【次步可忽略】
git branch feature1 
git branch -d feature1 # 删除本地此分支
git push origin --delete feature1  # 删除远端此分支

d. 开始Relase  【感觉这一步好繁琐】
git checkout -b release-0.1.0 dev

-- Optional: Bump version number, commit
-- Prepare release,commit

e. 完成Release

git checkout master
git merge --no-ff release-0.1.0
git push

git checkout dev    # 切回Dev分支，将release-0.1.0合并到Dev
git merge --no-ff release-0.1.0
git push

git branch -d release-0.1.0  # 然后删除临时正式分支 release-0.1.0
git push origin --delete release-0.1.0  

git tag -a v0.1.0 master  # 打上标签
git push --tags

f. 开始Hotfix
git checkout -b hotfix-0.1.1 master 

g. 完成 Hotfix 
git checkout master
git merge --no-ff hotfix-0.1.1
git push

git checkout dev
git merge --no-ff hotfix-0.1.1
git push
git branch -d hotfix-0.1.1
git tag -a v0.1.1 master
git push --tags

</pre>
### golang 交叉编译 

<pre>
// 目标程序是 linux 64位
GOOS=linux GOARCH=amd64 go build -o djason
// golang在Linux下编译生成.a与.so静态链接库、动态链接库文件
go build -buildmode=c-archive -o lib.a
go build -buildmode=c-shared -o test.so test-so.go
</pre>
### 支付宝、微信支付
https://github.com/shengzhi/payment

### MySQL 索引失效

在索引列上使用函数使得索引失效的是常见的索引失效原因之一，因此尽可能的避免在索引列上使用函数。尽管可以使用基于函数的索引来
解决索引失效的问题，但如此一来带来的比如磁盘空间的占用以及列上过多的索引导致DML性能的下降。

查询条件使用函数在索引列上，或者对索引列进行运算，运算包括(+，-，*，/，! 等) 错误的例子：select * from test where id-1=9; 正确的例子：select * from test where id=10 。

### Golang 二叉树

<pre>
/*
功能：二叉树
*/

package main

import (
    "fmt"
    "sync"
)

//二叉树节点结构
type Node struct {
    data  int
    left  *Node
    right *Node
}

//二叉查找树结构
type BST struct {
    root *Node
    lock sync.RWMutex
}

//插入方法(判断位置 中序遍历)
func insertNode(node *Node, addNode *Node) {
    if addNode.data < node.data {
        if node.left == nil {
            node.left = addNode
        } else {
            insertNode(node.left, addNode)
        }
    } else {
        if node.right == nil {
            node.right = addNode
        } else {
            insertNode(node.right, addNode)
        }
    }
}

//插入操作
func (t *BST) Insert(data int) {
    t.lock.Lock()
    defer t.lock.Unlock()
    node := &Node{
        data:  data,
        left:  nil,
        right: nil,
    }
    if t.root == nil {
        t.root = node
    } else {
        insertNode(t.root, node)
    }
}

//先序遍历
func PreOrderTraverse(bst *Node) {
    if bst != nil {
        fmt.Printf("%d  ", bst.data)
        PreOrderTraverse(bst.left)
        PreOrderTraverse(bst.right)
    }
}

//中序遍历
func InOrderTraverse(bst *Node) {
    if bst != nil {
        InOrderTraverse(bst.left)
        fmt.Printf("%d  ", bst.data)
        InOrderTraverse(bst.right)
    }
}

//后序遍历
func PostOrderTraverse(bst *Node) {
    if bst != nil {
        PostOrderTraverse(bst.left)
        PostOrderTraverse(bst.right)
        fmt.Printf("%d  ", bst.data)
    }
}

//查找
func SearchBST(node *Node, key int) bool {
    if node == nil {
        return false
    }
    if key == node.data {
        return true
    }
    if key < node.data {
        return SearchBST(node.left, key)
    } else {
        return SearchBST(node.right, key)
    }
}

//删除执行函数
func remove(node *Node, key int) *Node {
    if node == nil {
        return nil
    }
    if key == node.data {
        if node.left == nil && node.right == nil {
            node = nil
            return nil
        }
        if node.left == nil {
            node = node.right
            return node
        }
        if node.right == nil {
            node = node.left
            return node
        }
        rightside := node.right
        for {
            if rightside != nil && rightside.left != nil {
                rightside = rightside.left
            } else {
                break
            }
        }
        node.data = rightside.data
        node.right = remove(node.right, node.data)
        return node
    }
    if key < node.data {
        node.left = remove(node.left, key)
        return node
    } else {
        node.right = remove(node.right, key)
        return node
    }
}

//删除操作
func (t *BST) Remove(key int) {
    t.lock.Lock()
    defer t.lock.Unlock()
    remove(t.root, key)
}

func main() {
    var t BST
    t.Insert(2)
    t.Insert(6)
    t.Insert(5)
    t.Insert(1)
    t.Insert(10)
    t.Insert(8)

    fmt.Println("PREORDER...")
    PreOrderTraverse(t.root)
    fmt.Println("\nINORDER...")
    InOrderTraverse(t.root)
    fmt.Println("\nPOSTORDER...")
    PostOrderTraverse(t.root)
    fmt.Printf("\n")
    fmt.Println("Search...")
    fmt.Println(SearchBST(t.root, 6))
    fmt.Println("Remove...")
    t.Remove(6)
    fmt.Println("PREORDER...")
    PreOrderTraverse(t.root)
    fmt.Println("\nINORDER...")
    InOrderTraverse(t.root)
    fmt.Println("\nPOSTORDER...")
    PostOrderTraverse(t.root)
    fmt.Printf("\n")
    fmt.Println("Search...")
    fmt.Println(SearchBST(t.root, 6))
}
</pre>

### Golang 高并发
例子1

<pre>
package main 

import (
	"sync"
	"runtime"
	"time"
)

var (
    MaxWorker = 1000   //最大工作线程数
	MaxQueue  = 20000  //最大队列数
	i         = 0
	wg        sync.WaitGroup
)

func doTask() {
	//操作
	println(1)
   	time.Sleep(200 * time.Millisecond)
    wg.Done()
}

//这里模拟的http接口,每次请求抽象为一个job  实际工作操作
func handle() {
	job := Job{}
	i++
	println(i) 
    JobQueue <- job
}

type Worker struct {
    quit chan bool
}

func NewWorker() Worker {
    return Worker{
        quit: make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
    go func() {
        for {
            select {
            case <-JobQueue:
                // we have received a work request.
                doTask()
            case <-w.quit:
                // we have received a signal to stop
                return
            }
        }
    }()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
    go func() {
        w.quit <- true
    }()
}

type Job struct {
}

var JobQueue chan Job = make(chan Job, MaxQueue)

type Dispatcher struct {
}

func NewDispatcher() *Dispatcher {
    return &Dispatcher{}
}

func (d *Dispatcher) Run() {
    // starting n number of workers
    for i := 0; i < MaxWorker; i++ {
        worker := NewWorker()
        worker.Start()
    }
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
    d := NewDispatcher()
	d.Run()
    for i := 0; i < MaxQueue; i++ {
        wg.Add(1)
        handle()
    }
    wg.Wait()
}

</pre>

例子2

<pre>
package main

import (
     "fmt"
     "time"
     "runtime"
     "os"
)

// MaxWorker 与 DispatchNumControl 遵循木桶理论
var (
    //待处理的任务数
    JOBSNUM = 10
    //创建的工作任务线程数
    MaxWorker = 5
    //可以正常运行的协程数
    DISPATCHNUM = 10
)

type Payload struct {
    Num int
}

//待执行的工作
type Job struct {
     Payload Payload
}

//任务channel
var JobQueue = make(chan Job, JOBSNUM)

//用于控制并发处理的协程数
var DispatchNumControl = make(chan bool, DISPATCHNUM)

//执行任务的工作者单元
type Worker struct {
	WorkerPool chan chan Job //工作者池--每个元素是一个工作者的私有任务channel
	JobChannel chan Job      //每个工作者单元包含一个任务管道 用于获取任务
	quit       chan bool     //退出信号
	no         int           //编号
}

//调度中心
type Dispatcher struct {
	//工作者池
	WorkerPool chan chan Job
	//工作者数量
	MaxWorkers int
}

//创建一个新工作者单元
func NewWorker(workerPool chan chan Job, no int) Worker {
	println("创建一个新工作者单元")
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		no:         no,
	}
}

//循环  监听任务和结束信号
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel
            fmt.Println("w.WorkerPool <- w.JobChannel", w)
			select {
			case job := <-w.JobChannel:
                // 收到任务
                print("++++++++++++++++++++++++++正在执行任务++++++++++++++++++++++++++++++")
				fmt.Println(job)
				time.Sleep(100 * time.Second)
			case <-w.quit:
				// 收到退出信号
				return
			}
		}
	}()
}

// 停止信号
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

//创建调度中心
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

//工作者池的初始化
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 1; i < d.MaxWorkers+1; i++ {
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
	}
	go d.dispatch()
}

//使用有缓冲channel控制并发量
func Limit(work Job) bool {
   select {
   case <-time.After(time.Millisecond * 50):
      //超过限额，限制协程
      return false
   case DispatchNumControl <- true:
      //正常，放入任务队列
      return true
   }
}

//调度
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			println("job := <-JobQueue:")
			go func(job Job) {
                if Limit(job){
                    // fmt.Println("等待空闲worker(任务多的时候会阻塞这里)")
                    jobChannel := <-d.WorkerPool
                    // 将任务放到上述woker的私有任务channel中
                    jobChannel <- job
                }else{
                    println("我很忙，暂时不处理任务zZZ")
                }
			}(job)
        }
	}   
}

func AddQueue() {
	for i := 1; i <= 10; i++ {
		// 新建一个任务
		payLoad := Payload{Num: i}
		work := Job{Payload: payLoad}
		// 任务放入任务队列channel
		JobQueue <- work
		println("JobQueue <- work", i)
	    println("当前协程数:", runtime.NumGoroutine()-2)  // -2 
		time.Sleep(100 * time.Millisecond)
	}
	os.Exit(0)
}

func main() {  
	dispatcher := NewDispatcher(MaxWorker)
    dispatcher.Run()
    AddQueue()
}
</pre>

### string 性能优化

<pre>
func str2bytes(s string) []byte {
   x := (*[2]uintptr)(unsafe.Pointer(&s))
   h := [3]uintptr{x[0], x[1], x[1]}
   return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
   return *(*string)(unsafe.Pointer(&b))
}

func main(){
	s :="I want to know more about this world"
	b := str2bytes(s)
	s2 := bytes2str(b)
	fmt.Println(b, s2)
}
</pre>

### Hook目的

过滤一些关键函数调用，在函数执行前，先执行自己的挂钩函数。达到监控函数调用，改变函数功能的目的。

### shell 

//筛选出占用超过1G的文件夹
du -sh * |grep G

### 捕获杀死进程命令【Linux】

<pre>
package main

import (
 "fmt"
 "os"
 "os/signal"
 "syscall"
)

func main() {
	go SignalProc()

	done := make(chan bool, 1)
	for {
		select {
		case <-done:
			break
		}
	}
	fmt.Println("exit")

}

func SignalProc() {
	sigs:= make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP,os.Interrupt)

	for {
		msg := <-sigs
		fmt.Println("Recevied signal:", msg)

		switch msg {
		default:
			fmt.Println("get sig=%v\n", msg)
		case syscall.SIGHUP:
			fmt.Println("get sighup\n")
		case syscall.SIGUSR1:
			fmt.Println("SIGUSR1 test")
		case syscall.SIGUSR2:
			fmt.Println("SIGUSR2 test")
		}
	}
}
</pre>

### flysnow overwall
this had been running at remote server
<pre>
package main
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)
func main() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}
}
func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}
	if hostPortURL.Opaque == "443" { //https访问
		address = hostPortURL.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}
	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(b[:n])
	}
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}
</pre>

### 乐观锁与悲观锁 使用场景

* [1]如果对读的响应度要求非常高，比如证券交易系统，那么适合用乐观锁，因为悲观锁会阻塞读
* [2]如果读远多于写，那么也适合用乐观锁，因为用悲观锁会导致大量读被少量的写阻塞
* [3]如果写操作频繁并且冲突比例很高，那么适合用悲观写独占锁

悲观锁示例：

那么是否包上事务就万事大吉了呢？
    显然不是。因为如果同时有两个事务都分别SELECT到相同的vip_member记录，那么一样的会发生数据覆盖问题。那有什么办法可以解决呢？难道要设置事务隔离级别为SERIALIZABLE，考虑到性能不现实。
    我们知道InnoDB支持行锁。查看MySQL官方文档（innodb locking reads）了解到InnoDB在读取行数据时可以加两种锁：读共享锁和写独占锁。

如果事务A先获得了某行的写共享锁，那么事务B就必须等待事务A commit或者roll back之后才可以访问行数据。
显然要解决员状态更新问题，不能加读共享锁，只能加写共享锁，把SQL改写成如下:

<pre>
   vipMember = SELECT * FROM vip_member WHERE uid=1001 LIMIT 1 FOR UPDATE # 查uid为1001的会员
   if vipMember.end_at < NOW():
	   UPDATE vip_member SET start_at=NOW(), end_at=DATE_ADD(NOW(), INTERVAL 1 MONTH), active_status=1, updated_at=NOW() 	 WHERE uid=1001
   else:
	   UPDATE vip_member SET end_at=DATE_ADD(end_at, INTERVAL 1 MONTH), active_status=1, updated_at=NOW() WHERE uid=1001
</pre>

乐观锁示例：

上面一种加锁方案是一种悲观锁机制。而且SELECT...FOR UPDATE方式也不太常用，联想到CAS实现的乐观锁机制，于是我想到了第三种解决方案：乐观锁。
具体来说也挺简单，首先SELECT SQL不作任何修改，然后在UPDATE SQL的WHERE条件中加上SELECT出来的vip_memer的end_at条件。
这样可以根据UPDATE返回值来判断是否更新成功，如果返回值是0则表明存在并发更新，那么只需要重试一下就好了。
<pre>
vipMember = SELECT * FROM vip_member WHERE uid=1001 LIMIT 1 # 查uid为1001的会员
cur_end_at = vipMember.end_at
if vipMember.end_at < NOW():
   UPDATE vip_member SET start_at=NOW(), end_at=DATE_ADD(NOW(), INTERVAL 1 MONTH), active_status=1, updated_at=NOW() WHERE uid=1001 AND end_at=cur_end_at
else:
   UPDATE vip_member SET end_at=DATE_ADD(end_at, INTERVAL 1 MONTH), active_status=1, updated_at=NOW() WHERE uid=1001 AND end_at=cur_end_at
</pre>

### Windows 开机自启动设置

将需要执行的可执行文件放到 <span color=red>C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup</span> 目录下,开始->执行 msconfig命令，看是否存在。

### 限制Goruntine数量

<pre>
//限制goroutine数量
package main

import (
	"net"
	"strconv"
	"sync"
)
func demo9MincIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func demo9Miplist(cidr string) []string {
	var list []string
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		ip = net.ParseIP(cidr)
		list = append(list, ip.String())
		return list
	}
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); demo9MincIP(ip) {
		list = append(list, ip.String())
	}
	return list
}

func demo9chantest(wg *sync.WaitGroup, routineCtl chan string){
	defer wg.Done()
	for ip := range routineCtl{
		println("do something..."+ip)
	}
}
	
func main() {	
	routineCtl := make(chan string,100) 
	var i int = 0
	var portlist = []int{21,22,23,25,587,53,79,80,88,110,111,113,135,139,161,264,389,443,445,512,513,514,548,554,593,873,1099,1433,1521,2049,3260,5432,5900,6000,9100,9160,10000,11211,27017,27018,44818,47808,8080,8443,8554,3306,9999,500,3389}
	ips := demo9Miplist("192.168.1.0/24")
	var processNum = 100
	var wg = &sync.WaitGroup{}
	for i:=0; i<processNum; i++{
		wg.Add(1)
		go demo9chantest(wg, routineCtl)
	}
	for _, ports := range portlist {
		for _,ip :=range ips[1:len(ips)-1]{
			ip = ip+":"+strconv.Itoa(ports)
			routineCtl <- ip
			i++
		}			
	}
	close(routineCtl)
	wg.Wait()
	println("i的数量:")
	println(i)
}

</pre>
