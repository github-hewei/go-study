package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var test = make(map[int]int)
var ws = sync.WaitGroup{}
var lock = sync.RWMutex{}

func main() {
	ws.Add(2)

	go func() {
		for i := 0; i < 1000000000; i++ {
			n := rand.Intn(100)
			lock.Lock()
			test[n] = n
			lock.Unlock()
		}
		ws.Done()
	}()

	go func() {
		for {
			lock.Lock()
			for k, v := range test {
				fmt.Println(k, v)
			}
			lock.Unlock()
		}
	}()

	ws.Wait()
}


