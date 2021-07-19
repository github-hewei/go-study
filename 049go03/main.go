package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan string)

	go pump1(c1)
	go pump2(c2)
	go suck(c1, c2)

	time.Sleep(time.Duration(time.Second * 30))
}

func suck(c1 chan int, c2 chan string) {
	chRate := time.Tick(time.Duration(time.Second * 5))
	for {
		select {
		case v1 := <-c1:
			fmt.Println("Channel 1 : ", v1)
		case v2 := <-c2:
			fmt.Println("Channel 2 : ", v2)
		case <-chRate:
			fmt.Println("Log...")
		}
	}
}

func pump1(c1 chan int) {
	for i := 0; ; i++ {
		c1 <- i
		time.Sleep(time.Duration(time.Second))
	}
}

func pump2(c2 chan string) {
	for i := 0; ; i++ {
		c2 <- strconv.Itoa(i)
		time.Sleep(time.Duration(time.Second))
	}
}
