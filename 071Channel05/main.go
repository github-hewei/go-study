package main

import (
	"fmt"
	"time"
)

func main() {

	// 生产消费者模型

	ch := make(chan int, 10)

	go producer(ch)

	n := 5
	done := make(chan bool, n)

	for i := 0; i < n; i++ {
		go consumer(i, ch, done)
	}

	for i := 0; i < n; i++ {
		<- done
	}
}

func consumer(n int, ch <-chan int, done chan<- bool) {
	for {
		if v, ok := <-ch; ok {
			fmt.Println(n, "-消费: ", v)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	done <- true
}

func producer(ch chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Println("生产: ", i)
		ch <- i
	}
	close(ch)
}
