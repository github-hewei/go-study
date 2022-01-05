package main

import (
	"fmt"
	"sync"
)

func main() {
	w := sync.WaitGroup{}
	m := sync.Mutex{}

	// 演示安全的并发写入，加入锁机制，保证并发安全
	list := make([]int, 0)

	for i := 0; i < 2; i++ {
		w.Add(1)
		go func() {
			defer w.Done()

			for j := 0; j < 1000; j++ {
				m.Lock()
				list = append(list, 1)
				m.Unlock()
			}

		}()
	}

	w.Wait()
	fmt.Println(list, len(list), cap(list))
}
