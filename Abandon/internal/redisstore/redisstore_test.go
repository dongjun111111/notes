package redisstore

import (
	"testing"
)

/*
[root@AY1312201614528802d5Z redisstore]# go test -test.bench=".*"
PASS
Benchmark_SavePushMsg	    3000	    474275 ns/op
Benchmark_GetPushMsg	    5000	    336761 ns/op
ok  	galopush/redisstore	3.211s
*/
func Benchmark_SavePushMsg(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	db := NewStorager("115.28.128.9:6379", "12345678", 0)
	if db == nil {
		panic("connect to mstore error")
	}
	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		msg := []byte("hello world,I am a PushMsg")
		db.SavePushMsg("testid", msg)
	}
}

func Benchmark_GetPushMsg(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	db := NewStorager("115.28.128.9:6379", "12345678", 0)
	if db == nil {
		panic("connect to mstore error")
	}
	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		db.GetPushMsg("testid")
	}
}
