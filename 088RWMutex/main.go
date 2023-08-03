package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	Num int
	mu  sync.RWMutex
}

func (c *Counter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Num = c.Num + 1
}

func (c *Counter) Read() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	time.Sleep(time.Second)
	return c.Num
}

func NewCounter() *Counter {
	counter := new(Counter)
	counter.mu = sync.RWMutex{}
	return counter
}

func main() {
	counter := NewCounter()
	goroutineNum := 1000

	ch := make(chan struct{}, goroutineNum)

	for i := 0; i < goroutineNum; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Add()
			}

			ch <- struct{}{}
		}()
	}

	for i := 0; i < goroutineNum; i++ {
		<-ch
		fmt.Println("num:", counter.Read())
	}
}
