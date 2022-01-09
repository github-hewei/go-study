package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	// 优雅的退出（收到退出信号时处理完当前的工作再退出）

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	// 模拟并发进行的处理业务逻辑
	for i := 1; i <= 10; i++ {

		go func(i int) {
			for {
				// 希望把当前周期休眠完，再优雅的退出
				fmt.Println(i, time.Now().Unix())
				time.Sleep(time.Duration(i) * time.Second)
			}
		}(i)

	}

	fmt.Println(<-sig)
}

