package main

import (
	"fmt"
	"time"
)

func dump(s string, c chan string) {
	// 向 channel 中写入数据
	c <- s
	time.Sleep(time.Second * 2)
}

func main() {
	c := make(chan string)

	go dump("Hello", c)
	go dump("World", c)

	// 从 channel 中取出数据时会阻塞程序
	v1 := <-c
	v2 := <-c

	fmt.Println(v1)
	fmt.Println(v2)

}
