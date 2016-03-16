##Raft
###算法描述
通常来说，在分布式环境下，可以通过两种手段达成一致：

-  Symmetric, leader-less<br>
所有Server都是对等的，Client可以和任意Server进行交互

- Asymmetric, leader-based <br>
任意时刻，有且仅有1台Server拥有决策权，Client仅和该Leader交互,Raft算法采用该思路。
###基本概念
Raft算法将Server划分为3种角色：

-   Leader:负责Client交互和log复制，同一时刻系统中最多存在1个
-   Follower:被动响应请求RPC，从不主动发起请求RPC
-   Candidate:由Follower向Leader转换的中间状态

####Terms
众所周知，在分布式环境中，“时间同步”本身是一个很大的难题，但是为了识别“过期信息”，时间信息又是必不可少的。Raft为了解决这个问题，将时间切分为一个个的Term，可以认为是一种“逻辑时间”。

1. 每个Term至多存在1个Leader
2. 某些Term由于选举失败，不存在Leader
3. 每个Server本地维护currentTerm
####Heartbeats and Timeouts

1.       所有的Server均以Follower角色启动，并启动选举定时器
2.       Follower期望从Leader或者Candidate接收RPC
3.       Leader必须广播Heartbeat重置Follower的选举定时器
4.       如果Follower选举定时器超时，则假定Leader已经crash，发起选举
####Leader election
自增currentTerm，由Follower转换为Candidate，设置votedFor为自身，并行发起RequestVote RPC，不断重试，直至满足以下任一条件：

1.       获得超过半数Server的投票，转换为Leader，广播Heartbeat
2.       接收到合法Leader的AppendEntries RPC，转换为Follower
3.       选举超时，没有Server选举成功，自增currentTerm，重新选举
#####细节补充：
1.       Candidate在等待投票结果的过程中，可能会接收到来自其它Leader的AppendEntries RPC。如果该Leader的Term不小于本地的currentTerm，则认可该Leader身份的合法性，主动降级为Follower；反之，则维持Candidate身份，继续等待投票结果
2.       Candidate既没有选举成功，也没有收到其它Leader的RPC，这种情况一般出现在多个节点同时发起选举（如图Split Vote），最终每个Candidate都将超时。为了减少冲突，这里采取“随机退让”策略，每个Candidate重启选举定时器（随机值），大大降低了冲突概率
####正常操作流程：
1.       Client发送command给Leader
2.       Leader追加command至本地log
3.       Leader广播AppendEntriesRPC至Follower
4.       一旦日志项committed成功：
     
1)     Leader应用对应的command至本地StateMachine，并返回结果至Client

2)     Leader通过后续AppendEntriesRPC将committed日志项通知到Follower

3)     Follower收到committed日志项后，将其应用至本地StateMachine
###Safety
为了保证整个过程的正确性，Raft算法保证以下属性时刻为真：

1.       Election Safety
在任意指定Term内，最多选举出一个Leader
2.       Leader Append-Only
Leader从不“重写”或者“删除”本地Log，仅仅“追加”本地Log
3.       Log Matching
如果两个节点上的日志项拥有相同的Index和Term，那么这两个节点[0, Index]范围内的Log完全一致
4.       Leader Completeness
如果某个日志项在某个Term被commit，那么后续任意Term的Leader均拥有该日志项
5.       State Machine Safety
一旦某个server将某个日志项应用于本地状态机，以后所有server对于该偏移都将应用相同日志项
###Log compaction
随着系统的持续运行，操作日志不断膨胀，导致日志重放时间增长，最终将导致系统可用性的下降。快照（Snapshot）应该是用于“日志压缩”最常见的手段，Raft也不例外.
###Client interaction
典型的用户交互流程：

1.       Client发送command给Leader
            若Leader未知，挑选任意节点，若该节点非Leader，则重定向至Leader
2.       Leader追加日志项，等待commit，更新本地状态机，最终响应Client
3.       若Client超时，则不断重试，直至收到响应为止
细心的读者可能已经发现这里存在漏洞：Leader在响应Client之前crash，如果Client简单重试，可能会导致command被执行多次。
####Raft给出的方案：
Client赋予每个command唯一标识，Leader在接收command之前首先检查本地log，若标识已存在，则直接响应。如此，只要Client没有crash，可以做到“Exactly Once”的语义保证。
####个人建议：
尽量保证操作的“幂等性”，简化系统设计！
