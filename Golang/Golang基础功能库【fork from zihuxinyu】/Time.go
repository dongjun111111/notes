package Library

import (
	"time"
	"github.com/astaxie/beego"
)

const TIME_LAYOUT_OFTEN = "2006-01-02 15:04:05"

// 解析常用的日期时间格式：2014-01-11 16:18:00，东八区
func TimeParseOften(value string) (time.Time, error) {
	local, _ := time.LoadLocation("Local")
	return time.ParseInLocation(TIME_LAYOUT_OFTEN, value, local)
}

//返回当前时区的当前时间
func TimeLocal() ( time.Time) {
	stime := "2006-01-02 15:04:05 -07:00 "
	datastring := beego.DateFormat(time.Now(), stime)
	rtime, _ := beego.DateParse(datastring, stime)
	return rtime
}

//返回当前时区的当前时间
func DateLocal() ( time.Time) {
	sdate := "2006-01-02"
	datastring := beego.DateFormat(time.Now(), sdate)
	rtime, _ := beego.DateParse(datastring, sdate)
	return rtime
}

//返回当前时区的当前时间
func TimeLocalString() ( string) {
	datastring := beego.DateFormat(TimeLocal(), TIME_LAYOUT_OFTEN)
	return datastring
}

//得到多少分钟前的时间
func TheTimeString(counts time.Duration) (string) {
	baseTime := time.Now()
	date := baseTime.Add(counts)

	//beego.Debug("TheTimeString",date)
	datastring := beego.DateFormat(date, TIME_LAYOUT_OFTEN)
	return datastring


}
