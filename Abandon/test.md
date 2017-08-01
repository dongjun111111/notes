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
