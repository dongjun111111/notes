// 演示mongodb update操作
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
	updateAll()
}

// 更新单条记录
//	func (c *Collection) Update(selector interface{}, update interface{}) error
func update() {
	// 更新前数据：
	//	{
	//	    "_id" : ObjectId("56fdce98189df8759fd61e5b"),
	//	    "name" : "Tom",
	//	    "age" : 20
	//	}

	// 更新的mongodb语句
	//	db.getCollection('user').update(
	//	    { "_id": ObjectId("56fdce98189df8759fd61e5b") },
	//	    { "age": 21}
	//	)
	selector := bson.M{"_id": bson.ObjectIdHex("56fdce98189df8759fd61e5b")}
	data := bson.M{"age": 21}
	err := getDB().C("user").Update(selector, data)
	if err != nil {
		panic(err)
	}
	// 更新后数据：
	//	{
	//	    "_id" : ObjectId("56fdce98189df8759fd61e5b"),
	//	    "age" : 21
	//	}

}

// 更新不存在的数据
func updateNoExistData() {
	selector := bson.M{"_id": bson.ObjectIdHex("16fdce98189df8759fd61e5b")}
	data := bson.M{"age": 21}
	err := getDB().C("user").Update(selector, data)
	if err != nil {
		fmt.Println(err == mgo.ErrNotFound)
		// output: true
	}
}

// 更新单个字段值，需要在更新的数据data中使用$set来标识
func updateBySet() {
	// 更新前数据：
	//	{
	//	    "_id" : ObjectId("571de968a99cff2c68264807"),
	//	    "name" : "Tom",
	//	    "age" : 20
	//	}

	// 更新的mongodb语句
	//	db.getCollection('user').update(
	//	    { "_id": ObjectId("571de968a99cff2c68264807") },
	//	    { "$set": { "age": 20 } }
	//	)
	selector := bson.M{"_id": bson.ObjectIdHex("571de968a99cff2c68264807")}
	data := bson.M{"$set": bson.M{"age": 21}}
	err := getDB().C("user").Update(selector, data)
	if err != nil {
		panic(err)
	}
	// 更新后数据：
	//	{
	//	    "_id" : ObjectId("571de968a99cff2c68264807"),
	//	    "age" : 21,
	//	    "name" : "Tom"
	//	}

}

// 批量更新数据
//	func (c *Collection) UpdateAll(selector interface{}, update interface{}) (info *ChangeInfo, err error)
func updateAll() {
	selector := bson.M{"name": "Tom"}
	data := bson.M{"$set": bson.M{"age": 22}}
	changeInfo, err := getDB().C("user").UpdateAll(selector, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", changeInfo)
	// output: &{Updated:2 Removed:0 UpsertedId:<nil>}
}

// 直接根据mongodb中的id进行更新
// 	func (c *Collection) UpdateId(id interface{}, update interface{}) error
//	类似
//	err := collection.Update(bson.M{"_id": id}, update)
func updateId() {
	id := bson.ObjectIdHex("571de968a99cff2c68264807")
	data := bson.M{"$set": bson.M{"age": 30}}
	err := getDB().C("user").UpdateId(id, data)
	if err != nil {
		panic(err)
	}
}

// 更新数据，如果数据不存在则创建该数据
//  func (c *Collection) Upsert(selector interface{}, update interface{}) (info *ChangeInfo, err error)
func upsert() {
	selector := bson.M{"key": "max"}
	data := bson.M{"$set": bson.M{"value": 30}}
	changeInfo, err := getDB().C("config").Upsert(selector, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", changeInfo)
	// 首次执行output: &{Updated:0 Removed:0 UpsertedId:ObjectIdHex("571df02ea99cff2c6826480a")}
	// 再次执行output: &{Updated:1 Removed:0 UpsertedId:<nil>}
}

// 同upsert一致，只是selector条件换成id
// func (c *Collection) UpsertId(id interface{}, update interface{}) (info *ChangeInfo, err error)
// 类似
// info, err := collection.Upsert(bson.M{"_id": id}, update)
func upsetId() {
	id := bson.ObjectIdHex("571df02ea99cff2c6826480b")
	data := bson.M{"$set": bson.M{"key": "max", "value": 30}}
	changeInfo, err := getDB().C("config").UpsertId(id, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", changeInfo)
	// 首次执行output: &{Updated:0 Removed:0 UpsertedId:ObjectIdHex("571df02ea99cff2c6826480b")}
	// 再次执行output: &{Updated:1 Removed:0 UpsertedId:<nil>}
}
