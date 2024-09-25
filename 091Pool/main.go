package main

import (
	"fmt"
	"math/rand"
	"time"
)

func TaskPool(size int, taskCh <-chan func() any, retCh chan<- any) {
	for i := 0; i < size; i++ {
		go func(taskCh <-chan func() any, retCh chan<- any) {
			for {
				task := <-taskCh
				result := task()
				retCh <- result
			}
		}(taskCh, retCh)
	}
}

func main() {
	retCh := make(chan any, 256)
	taskCh := make(chan func() any, 256)
	doneCh := make(chan struct{})

	TaskPool(100, taskCh, retCh)

	total := 10
	done := 0

	go func() {
		for {
			ret := <-retCh
			fmt.Println(ret)
			done++

			if done == total {
				doneCh <- struct{}{}
				return
			}
		}
	}()

	for i := 0; i < total; i++ {
		taskCh <- func() any {
			n := rand.Intn(10)
			time.Sleep(time.Second * time.Duration(n))
			return n
		}
	}

	<-doneCh
}
