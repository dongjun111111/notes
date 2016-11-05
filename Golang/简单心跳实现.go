package 心跳实现

//针对tcp长连接做的应用层心跳检测
import (
	"sync"
	"time"
)

//连接过期处理器接口类
type DierInterface interface {
	Die()
}

type PingTimer struct {
	mu    sync.Mutex
	ptmr  *time.Timer   //ping定时器
	pTime time.Time     //记录最近ping时间戳
	dier  DierInterface //心跳死亡后的执行函数
}

//注册连接过期的处理器
func (t *PingTimer) SetDier(pi DierInterface) {
	t.dier = pi
	t.pTime = time.Now()
}

//定时器唤醒处理函数
func (t *PingTimer) Process() {
	bDie := false
	t.mu.Lock()
	t.ptmr = nil

	//如果连接未激活时间超过最大时长，则认为该连接过期，销毁处理
	if t.pTime.Nanosecond() != 0 && (time.Since(t.pTime) > (time.Duration)(opts.PingTimeOut)*time.Second) {
		bDie = true
	}

	t.mu.Unlock()

	if bDie { //销毁，并且关闭定时器
		t.dier.Die()
		t.Clear()
	} else {
		//重置心跳检测定时器
		t.Set()
	}
}

//设置定时器
func (t *PingTimer) Set() {
	t.mu.Lock()
	defer t.mu.Unlock()
	xylogs.Debug("Set PingTimer.")
	t.ptmr = time.AfterFunc(time.Duration(GOpts.PingInterval)*time.Second, func() { t.Process() })
}

//刷新激活时间戳
func (t *PingTimer) RefleshTime() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.pTime = time.Now()
}

//销毁函数
func (t *PingTimer) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.ptmr == nil {
		return
	}
	t.ptmr.Stop()
	t.ptmr = nil
}
