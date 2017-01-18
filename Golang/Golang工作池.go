package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type GoroutinePool struct {
	Queue          chan func() error
	Number         int
	Total          int
	result         chan error
	finishCallback func()
}

// 初始化
func (self *GoroutinePool) Init(number int, total int) {
	self.Queue = make(chan func() error, total)
	self.Number = number
	self.Total = total
	self.result = make(chan error, total)
}

// 开始工作
func (self *GoroutinePool) Start() {
	for i := 0; i < self.Number; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}

				err := task()
				self.result <- err
			}
		}()
	}

	// 获得每个work的执行结果
	for j := 0; j < self.Total; j++ {
		res, ok := <-self.result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}

	// 所有任务都执行完成，回调函数
	if self.finishCallback != nil {
		self.finishCallback()
	}
}

// 停止工作
func (self *GoroutinePool) Stop() {
	close(self.Queue)
	close(self.result)
}

// 添加任务
func (self *GoroutinePool) AddTask(task func() error) {
	self.Queue <- task
}

// 设置结束回调
func (self *GoroutinePool) SetFinishCallback(callback func()) {
	self.finishCallback = callback
}

func Download_test() {
	urls := []string{
		"http://dlsw.baidu.com/sw-search-sp/soft/44/17448/Baidusd_Setup_4.2.0.7666.1436769697.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/3a/12350/QQ_V7.4.15197.0_setup.1436951158.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/9d/14744/ChromeStandalone_V43.0.2357.134_Setup.1436927123.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/6f/15752/iTunes_V12.2.1.16_Setup.1436855012.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/70/17456/BaiduAn_Setup_5.0.0.6747.1435912002.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/40/12856/QIYImedia_1_06_v4.0.0.32.1437470004.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/42/37473/BaiduSoftMgr_Setup_7.0.0.1274.1436770136.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/49/16988/YoudaoNote_V4.1.0.300_setup.1429669613.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/55/11339/bdbrowserSetup-7.6.100.2089-1212_11000003.1437029629.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/53/21734/91zhushoupc_Windows_V5.7.0.1633.1436844901.exe",
	}
	pool := new(GoroutinePool)
	pool.Init(3, len(urls))
	for i := range urls {
		url := urls[i]
		pool.AddTask(func() error {
			return download(url)
		})
	}
	isFinish := false
	pool.SetFinishCallback(func() {
		func(isFinish *bool) {
			*isFinish = true
		}(&isFinish)
	})
	pool.Start()
	for !isFinish {
		time.Sleep(time.Millisecond * 100)
	}
	pool.Stop()
	fmt.Println("所有操作已完成！")
}

func download(url string) error {
	fmt.Println("开始下载... ", url)
	sp := strings.Split(url, "/")
	filename := sp[len(sp)-1]
	file, err := os.Create("D:/gopath/src/test/do/" + filename)
	if err != nil {
		return err
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	length, err := io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	fmt.Println("## 下载完成！ ", url, " 文件长度：", length)
	return nil
}

func main() {
	Download_test()
}
