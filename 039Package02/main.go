package main

//引入包的名称不需要必须和文件名一致
//可以自定义包的名称

import (
	"fmt"

	//采用包的默认名称 p1
	"github.com/github-hewei/go-study/038Package01/p1"

	//采用包的命名 p2
	p2 "github.com/github-hewei/go-study/038Package01/p1"
)

func main() {

	p1.A = "AAAAA"
	p1.Run1()

	fmt.Println(p1.Run2().F2())

	fmt.Println("-------------------------------------------------------")

	p2.A = "BBBBB"
	p2.Run1()

	fmt.Println(p2.Run2().F2())

}
