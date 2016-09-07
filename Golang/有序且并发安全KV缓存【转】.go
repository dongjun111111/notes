//有序且并发安全KV缓存，来自于henrylee2cn
//高频操作中慎用判断
package main

import (
	"fmt"
	"sort"
	"sync"
)

type KV struct {
	count int
	keys  []string
	hash  map[string]interface{}
	lock  sync.RWMutex
}

// 新建KV缓存(preCapacity为预申请内存容量)
func NewKV(preCapacity uint) *KV {
	return &KV{
		keys: make([]string, 0, int(preCapacity)),
		hash: make(map[string]interface{}, int(preCapacity)),
	}
}

// 添加kv键值对
func (this *KV) Set(k string, v interface{}) {
	this.lock.Lock()
	if _, ok := this.hash[k]; !ok {
		this.keys = append(this.keys, k)
		sort.Strings(this.keys)
		this.count++
	}
	this.hash[k] = v
	this.lock.Unlock()
}

// 获取数据长度
func (this *KV) Count() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.count
}

// 由key检索value
func (this *KV) Get(k string) (interface{}, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	v, ok := this.hash[k]
	return v, ok
}

// 根据key排序，返回有序的vaule切片
func (this *KV) Values() []interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	vals := make([]interface{}, this.count)
	for i := 0; i < this.count; i++ {
		vals[i] = this.hash[this.keys[i]]
	}
	return vals
}

func main() {
	test := NewKV(10)
	test.Set("Y", "6767")
	test.Set("g", "ffff")
	test.Set("d", "ggg")
	res, _ := test.Get("Y")
	res2, _ := test.Get("g")
	res3, _ := test.Get("d")
	fmt.Println(res, res2, res3)
	fmt.Println(test.Count())
}
