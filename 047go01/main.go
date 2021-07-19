package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var ws = sync.WaitGroup{}

func dump(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(time.Second)
	}
	ws.Done()
}

func main() {

	var n int = 2

	ws.Add(n)

	for i := 0; i < n; i++ {
		go dump("Hello World: " + strconv.Itoa(i))
	}

	ws.Wait()

}
