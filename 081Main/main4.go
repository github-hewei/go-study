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

	stopCh := make(chan struct{})
	finishedCh := make(chan struct{})

	go func(stopCh, finishedCh chan struct{}) {
		for {
			select {
			case <- stopCh:
				fmt.Println("Stopped")
				finishedCh <- struct{}{}
				return
			default:
				fmt.Println("wait...")
				time.Sleep(time.Second * 10)
			}
		}
	}(stopCh, finishedCh)

	<-sig
	stopCh <- struct{}{}
	<-finishedCh
	fmt.Println("Finished")
}

