// json解码
// 尝试用字典进行接收
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// json格式字符串
	var jsonUsers = []byte(`[
		{"id": "1", "name": "Anny"},
		{"id": "2", "name": "Tom"}
	]`)

	// 尝试用字典进行接收
	// 因为json格式的键通常为字符串类型
	// 值有可能为各种类型，如果整形、浮点、数组、对象等
	// 所以map存储的值定义为interface{}类型
	var users []map[string]interface{}
	err := json.Unmarshal(jsonUsers, &users)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", users)
	// output: [map[id:1 name:Anny] map[id:2 name:Tom]]

}
