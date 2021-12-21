package main

import (
	"fmt"
	"sync"
)

func main() {

	// 并发不安全
	//count := 0
	//wait := sync.WaitGroup{}
	//
	//for i := 0; i < 100; i++ {
	//	wait.Add(1)
	//	go func() {
	//		for j := 0; j < 10000; j++ {
	//			count ++
	//		}
	//		wait.Done()
	//	}()
	//}
	//
	//wait.Wait()
	//fmt.Println(count)

	// 并发安全
	count := 0
	wait := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 100; i++ {
		wait.Add(1)
		go func() {
			for j := 0; j < 10000; j++ {
				mu.Lock()
				count ++
				mu.Unlock()
			}
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println(count)
}


