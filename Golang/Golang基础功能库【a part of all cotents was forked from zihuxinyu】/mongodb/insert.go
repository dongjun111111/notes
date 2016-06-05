// 演示mongodb插入数据操作
// 主要使用Insert函数
// 函数原型
// func (*Collection) Insert
// 	func (c *Collection) Insert(docs ...interface{}) error
package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// get mongodb db
func getDB() *mgo.Database {
	session, err := mgo.Dial("172.16.27.134:10001")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("test")
	return db
}

func main() {
	insert()
	insertMuti()
	insertArray()
	insertNesting()
	insertMap()
	insertObjectId()
}

// 插入单条数据
func insert() {
	db := getDB()

	c := db.C("user")
	type User struct {
		Name string "bson:`name`"
		Age  int    "bson:`age`"
	}

	err := c.Insert(&User{Name: "Tom", Age: 20})
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
}

// 插入多条记录
func insertMuti() {
	db := getDB()

	c := db.C("user")
	type User struct {
		Name string "bson:`name`"
		Age  int    "bson:`age`"
	}

	err := c.Insert(&User{Name: "Tom", Age: 20}, &User{Name: "Anny", Age: 28})
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
}

// 插入数组格式
func insertArray() {
	db := getDB()
	c := db.C("user")

	type User struct {
		Name   string   "bson:`name`"
		Age    int      "bson:`age`"
		Groups []string "bson:`groups`"
	}

	err := c.Insert(&User{
		Name:   "Tom",
		Age:    20,
		Groups: []string{"news", "sports"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
}

// 插入嵌套数据
func insertNesting() {
	db := getDB()

	c := db.C("user")

	type Toy struct {
		Name string "bson:`name`"
	}
	type User struct {
		Name string "bson:`name`"
		Age  int    "bson:`age`"
		Toys []Toy
	}

	err := c.Insert(&User{
		Name: "Tom",
		Age:  20,
		Toys: []Toy{Toy{Name: "dog"}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
}

// 插入map格式的数据
func insertMap() {
	db := getDB()
	c := db.C("user")

	user := map[string]interface{}{
		"name":   "Tom",
		"age":    20,
		"groups": []string{"news", "sports"},
		"toys": []map[string]interface{}{
			map[string]interface{}{
				"name": "dog",
			},
		},
	}

	err := c.Insert(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
}

// 插入关联其它集合ObjectId的数据
// 要使用bson.ObjectIdHex函数对字符串进行转化
// 函数原型
// func ObjectIdHex
//	func ObjectIdHex(s string) ObjectId
func insertObjectId() {
	db := getDB()
	c := db.C("user")

	user := map[string]interface{}{
		"name":     "Tom",
		"age":      20,
		"group_id": bson.ObjectIdHex("540046baae59489413bd7759"),
	}

	err := c.Insert(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
}
