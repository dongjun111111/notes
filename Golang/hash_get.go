package utils

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

//初始化
func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

//添加元素 成功返回ture
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

//删除元素
func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

//清除元素
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

//判断元素是否存在
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

//获取元素数量
func (set *HashSet) Len() int {
	return len(set.m)
}

//判断一个HashSet是否相同
func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}


//生成快照
func (set *HashSet) Elements() []interface{} {
	initialLen := set.Len()
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
		if actualLen < initialLen {
			snapshot = snapshot[:actualLen]
		}
	}
	return snapshot
}

//返回字符串
func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	frist := true
	for key := range set.m {
		if frist {
			frist = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

//真包含
func (set *HashSet) IsSuperset(other *HashSet) bool {
	if other == nil {
		return false
	}
	oneLen := set.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if set.Contains(v) {
			return false
		}
	}
	return true
}
