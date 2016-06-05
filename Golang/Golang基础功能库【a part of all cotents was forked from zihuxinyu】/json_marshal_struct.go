// json编码
// 从struct类型编码成json格式字符串
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 声明一个结构体
	type ColorGroup struct {
		ID     int      `json:"id,string"`
		Name   string   `json:"name,omitempty"`
		Colors []string `json:"colors"`
	}

	// 对ColorGroup类型的变量进行编码
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// output: {"id":"1","name":"Reds","colors":["Crimson","Red","Ruby","Maroon"]}

	// 如果没有设置Name属性值，因为标记为了omitempty属性，则在编码成json的时候会忽略Name属性
	group = ColorGroup{
		ID:     1,
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err = json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// output: {"id":"1","colors":["Crimson","Red","Ruby","Maroon"]}

	// 如果没有设置Colors值，因为没有omitempty属性，会输出nil
	group = ColorGroup{
		ID:   1,
		Name: "Reds",
	}
	b, err = json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// output: {"id":"1","name":"Reds","colors":null}

}
