package main

import (
	"fmt"
	"sync"
)

func main() {

	var once = sync.Once{}

	for i := 0; i < 10; i++ {
		//此函数只会执行 一次
		once.Do(func() {
			fmt.Println("only once")
		})
	}

	// WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。
	// 每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。
	var wg = sync.WaitGroup{}

	wg.Add(2)

	wg.Done() // wg.Done 等于 wg.Add(-1)
	wg.Done()

	wg.Wait() // 如果线程数没有归零会一直阻塞下去，会报错

}
