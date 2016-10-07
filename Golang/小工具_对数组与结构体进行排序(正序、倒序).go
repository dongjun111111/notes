package main

import (
	"fmt"
	"sort"
)

//---------结构体排序START---------
type Person struct {
	Name string // 姓名
	Age  int    // 年纪
}

type PersonWrapper struct {
	people []Person
	by     func(p, q *Person) bool
}

func (pw PersonWrapper) Len() int { // 重写 Len() 方法
	return len(pw.people)
}
func (pw PersonWrapper) Swap(i, j int) { // 重写 Swap() 方法
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}
func (pw PersonWrapper) Less(i, j int) bool { // 重写 Less() 方法
	return pw.by(&pw.people[i], &pw.people[j])
}

//---------结构体排序END---------

func main() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
	fmt.Println("------------正序---------------")
	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)
	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
	fmt.Println("------------倒序---------------")
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
	fmt.Println("-------------结构体(特定字段)排序-----------")
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}
	fmt.Println(people)
	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return q.Age < p.Age //Age 递减排序
	}})
	fmt.Println(people)
	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return p.Name < q.Name //Name 递增排序
	}})
	fmt.Println(people)
}
