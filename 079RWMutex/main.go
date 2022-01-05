package main

import (
	"fmt"
	"sync"
)

func main() {
	w := sync.WaitGroup{}

	// 演示切片不安全的并发写入，发现有数据丢失，实际写入数据不足2000条
	list := make([]int, 0)

	for i := 0; i < 2; i++ {
		w.Add(1)
		go func() {
			defer w.Done()

			for j := 0; j < 1000; j++ {
				list = append(list, 1)
			}

		}()
	}

	w.Wait()
	fmt.Println(list, len(list), cap(list))
}
