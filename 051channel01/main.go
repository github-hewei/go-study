package main

import "fmt"

func main() {

	// <-chan 只可以读的通道
	// chan-> 只可以写的通道

	// 定义一个可以读写的通道，大小为4
	c1 := make(chan int, 4)

	c1 <- 1
	c1 <- 2
	c1 <- 3
	c1 <- 4
	//c1 <- 5 通道写满无法写入会导致程序无限阻塞，所以会报错

	<-c1
	<-c1
	<-c1
	<-c1
	//<-c1 通道一直读取不到数据会导致程序无限阻塞，所以会报错

	c1 <- 1
	c1 <- 2

	v1, ok := <-c1
	fmt.Println(v1, ok)

	v2, ok := <-c1
	fmt.Println(v2, ok)

	c1 <- 2
	c1 <- 3
	c1 <- 4

	// c2 是一个只读的通道，可以通过传递其他已经写入数据的通道来给它赋值
	c2 := test(c1)

	close(c1)

	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println(<-c2)

	// c2 <- 1 只读通道不能写入数据
	// close(c2) 只读的通道不能close

	c3 := make(chan<- int, 1)
	c3 <- 1

	fmt.Println(c3) // 变量内存地址
	close(c3)
}

func test(c chan int) <-chan int {
	return c
}
