//counter 实现简单的计数器
package counter

import (
	"fmt"
	"sync"
)

type Counter struct {
	mutex sync.Mutex
	datas map[string]*Data
}

type Data struct {
	current int64
	max     int64
}

func NewCounter() *Counter {
	var counter Counter
	counter.datas = make(map[string]*Data)
	return &counter
}

func (this *Counter) Add(tp string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	d, ok := this.datas[tp]
	if ok {
		d.current = d.current + 1
		if d.current > d.max {
			d.max = d.current
		}
	} else {
		d = new(Data)
		d.current = 1
		d.max = d.current

		this.datas[tp] = d
	}
}

func (this *Counter) Sub(tp string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	d, ok := this.datas[tp]
	if ok {
		d.current = d.current - 1
	}
}

func (this *Counter) String() string {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	res := "[cnt]"
	for k, v := range this.datas {
		tmp := fmt.Sprintf("\t[%s current is %d max is %d]", k, v.current, v.max)
		res = res + tmp
	}
	return res
}
