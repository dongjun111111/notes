// 演示mongodb find操作
// func (c *Collection) Find(query interface{}) *Query
// bson.M{}
// 	type M map[string]interface{}
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
	findPage()
}

// 查找单条记录
func findOne() {
	db := getDB()

	c := db.C("user")

	// 用struct接收,一般情况下都会这样处理
	type User struct {
		Name string "bson:`name`"
		Age  int    "bson:`age`"
	}
	user := User{}
	err := c.Find(bson.M{"name": "Tom"}).One(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	// output: {Tom 20}

	// 用bson.M结构接收，当你不了解返回的数据结构格式时，可以用这个先查看，然后再定义struct格式
	// 在处理mongodb组合查询时，经常这么干
	result := bson.M{}
	err = c.Find(nil).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// output: map[_id:ObjectIdHex("56fdce98189df8759fd61e5b") name:Tom age:20]

}

// 查找多条记录
func findMuit() {
	db := getDB()

	c := db.C("user")

	// 使用All方法，一次性消耗较多内存，如果数据较多，可以考虑使用迭代器
	type User struct {
		Id   bson.ObjectId `bson:"_id,omitempty"`
		Name string        "bson:`name`"
		Age  int           "bson:`age`"
	}
	var users []User
	err := c.Find(nil).All(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	// output: [{ObjectIdHex("56fdce98189df8759fd61e5b") Tom 20}...]

	// 使用迭代器获取数据可以避免一次占用较大内存
	var user User
	iter := c.Find(nil).Iter()
	for iter.Next(&user) {
		fmt.Println(user)
	}
	// output:
	// {ObjectIdHex("56fdce98189df8759fd61e5b") Tom 20}
	// {ObjectIdHex("56fdce98189df8759fd61e5c") Tom 20}
	// ...
}

// 查找指定字段
func findField() {
	db := getDB()

	c := db.C("user")

	// 只读取name字段
	type User struct {
		Name string "bson:`name`"
	}
	var users []User
	err := c.Find(bson.M{}).Select(bson.M{"name": 1}).All(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	// output: [{Tom} {Tom} {Anny}...]

	// 只排除_id字段
	type User2 struct {
		Name string "bson:`name`"
		Age  int    "bson:`age`"
	}
	var users2 []User2
	err = c.Find(bson.M{}).Select(bson.M{"_id": 0}).All(&users2)
	if err != nil {
		panic(err)
	}
	fmt.Println(users2)
	// output: [{Tom 20} {Tom 20} {Anny 28}...]

}

// 查询嵌套格式数据
func findNesting() {
	db := getDB()

	c := db.C("user")

	// 使用嵌套的struct接收数据
	type User struct {
		Name string "bson:`name`"
		Age  int    "bson:`age`"
		Toys []struct {
			Name string "bson:`name`"
		}
	}
	var users User
	// 只查询toys字段存在的
	err := c.Find(bson.M{"toys": bson.M{"$exists": true}}).One(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	// output: {Tom 20 [{dog}]}
}

// 排序
// 使用Sort函数
// func (q *Query) Sort(fields ...string) *Query
func findSort() {
	db := getDB()

	c := db.C("user")

	type User struct {
		Id   bson.ObjectId `bson:"_id,omitempty"`
		Name string        "bson:`name`"
		Age  int           "bson:`age`"
	}
	var users []User
	// 按照age字段降序排列，如果升序去掉横线"-"就可以了
	err := c.Find(nil).Sort("-age").All(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	// output:
	// [{ObjectIdHex("56fdce98189df8759fd61e5d") Anny 28} ...]
	// ...
}

// 分页查询
// 使用Skip函数和Limit函数
// func (q *Query) Skip(n int) *Query
// func (q *Query) Limit(n int) *Query
func findPage() {
	db := getDB()

	c := db.C("user")

	type User struct {
		Id   bson.ObjectId `bson:"_id,omitempty"`
		Name string        "bson:`name`"
		Age  int           "bson:`age`"
	}
	var users []User
	// 表示从偏移位置为2的地方开始取两条记录
	err := c.Find(nil).Sort("-age").Skip(2).Limit(2).All(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	// output:
	// [{ObjectIdHex("56fdce98189df8759fd61e5d") Anny 20} ...]
	// ...
}

// 查找数据总数
func count() {
	db := getDB()

	c := db.C("user")

	// 查找表总数
	count, err := c.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
	// output: 8

	// 结合find条件查找
	count, err = c.Find(bson.M{"name": "Tom"}).Count()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
	// output: 6

}
