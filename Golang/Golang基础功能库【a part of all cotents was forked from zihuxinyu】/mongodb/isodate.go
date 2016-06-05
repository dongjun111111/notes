// 处理一下mongodb中的isodate时间问题
// 在golang中一般都是通过time.Time进行接收
package main

import (
	"fmt"
	"time"

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
	db := getDB()
	c := db.C("isodate")

	// mongodb中存储的数据如下所示
	//	{
	//	    "_id" : ObjectId("572f3c68e43001d2c1703aa7"),
	//	    "time" : ISODate("2015-07-08T09:29:14.002Z")
	//	}
	type Model struct {
		Id   bson.ObjectId `bson:"_id,omitempty"`
		Time time.Time     `bson:"time"`
	}
	m := Model{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex("572f3c68e43001d2c1703aa7")}).One(&m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", m)
	// output: {Id:ObjectIdHex("572f3c68e43001d2c1703aa7") Time:2015-07-08 17:29:14.002 +0800 CST}
	// 从输出可以看出时间被转换为了CST时间格式，从数字上来看比mongodb中的快8个小时

	// 插入当前时间
	now := time.Now()
	fmt.Printf("%+v\n", now)
	// output: 2016-05-12 14:34:00.998011694 +0800 CST
	err = c.Insert(Model{Time: now})
	if err != nil {
		panic(err)
	}

	// 时间字符串转到到time.Time格式
	// 使用time.Parse方法进行转换
	timeString := "2016-05-12 14:34:00.998011694 +0800 CST"
	t, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)

}
