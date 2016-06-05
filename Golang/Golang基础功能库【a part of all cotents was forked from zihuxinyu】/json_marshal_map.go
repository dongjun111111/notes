// json编码
// 从map类型编程成json格式字符串
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]interface{}{
		"id":      1,
		"name":    "Socrates",
		"friends": []string{"Plato", "Aristotle"},
	}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
