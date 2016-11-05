package main

//伪消息队列 channel
import (
	"fmt"
	"runtime"
	"time"
)

var url_read_mysql_lis = make(chan string, 2000)

func write_channel(queue chan string, data string) {
	ok := true
	for ok {
		select {
		case <-time.After(time.Second * 2):
			println("write channel timeout")
			ok = true
		case queue <- data:
			ok = false
		}
	}
}

func read_channel(queue chan string) string {
	ok := true
	for ok {
		select {
		case <-time.After(time.Second * 2):
			println("read channel timeout")
			ok = true
		case i := <-queue:
			ok = false
			return i
		}
	}
	return ""
}

func main() {
	fmt.Printf("11111\n")
	for {

		go func() {
			write_channel(url_read_mysql_lis, "33333333333333")
		}()
		ss := read_channel(url_read_mysql_lis)
		fmt.Println(ss)
		time.Sleep(1 * time.Second)
		runtime.Gosched()
	}

}
