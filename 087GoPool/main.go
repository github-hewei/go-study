package main

import (
	"fmt"
	"github.com/github-hewei/go-study/087GoPool/gopool"
	"time"
)

func main() {
	p := gopool.NewPool(10, 10, 10)

	for {
		p.Schedule(func() {
			fmt.Println("Hello World")
		})

		time.Sleep(time.Second)
	}
}
