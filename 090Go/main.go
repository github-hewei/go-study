package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Test struct {
	num int
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan struct{}, 50)

	go func() {
		for {
			fmt.Printf("------------------%d\n", len(ch))
			time.Sleep(time.Second)
		}
	}()

	test := &Test{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go request(fmt.Sprintf("%d", i), wg, ch, test)
	}

	wg.Wait()
	fmt.Println("num: ", test.num)
}

func request(url string, wg *sync.WaitGroup, ch chan struct{}, test *Test) error {
	fmt.Printf("[%s] start\n", url)
	n := rand.Intn(10)
	time.Sleep(time.Second * time.Duration(n))
	fmt.Printf("[%s] ok (%d) \n", url, n)
	<-ch
	wg.Done()

	for i := 0; i < 1000; i++ {
		test.num++
	}

	return nil
}
