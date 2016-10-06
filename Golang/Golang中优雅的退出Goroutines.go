package main

/*
在go语言中优雅退出goroutines,通常需要做以下3点：

1. 向各个goroutines发通知，令其退出，如shutdown.

2. 等待各个goroutines都退出，如: sync.WaitGroup.

3. 在退出goroutine之前，确保数据不丢失（1.停止生产数据。2.关闭数据channel messages. 3. 消费者goroutine检查判断数据channel messages是否有效，若无效，则退出。）
*/
import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func consumer(messages <-chan int, shutdown <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for {

		select {

		case message, ok := <-messages:
			//do something.
			if ok {

				fmt.Println(message)

			} else {

				//no data , exit.
				fmt.Println("no data, exit.")
				return
			}

		case _ = <-shutdown:
			//we `re done!
			//shutdown now , messages buffered channel data may be lost.
			fmt.Println("all done!")
			return

		}
	}
}

func main() {

	//signals handle.
	shutdown := make(chan int)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {

		sig := <-sigs
		fmt.Println(sig)
		//shutdown now , messages buffered channel data may be lost.
		//close(shutdown)
		/*or*/ shutdown <- 0
	}()
	messages := make(chan int, 16)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go consumer(messages, shutdown, wg)
	for i := 0; i < 10; i++ {

		messages <- i
	}
	close(messages) //flush messages channel , no data loss.
	fmt.Println("wait!")
	wg.Wait()

}
