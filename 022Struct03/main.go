package main

import (
	"fmt"
)

type Car struct {
	brand  string  // 品牌
	color  string  // 颜色
	price  float64 // 价格
	length float64 // 长度
}

func main() {
	var c1 Car = Car{}
	fmt.Println("c1 = ", c1)

	c2 := Car{"BMW", "red", 300000, 5.4}

	//结构体的内存地址和结构体中第一个元素的内存地址相同
	fmt.Println(&c2)
	fmt.Println(&(c2.brand))

	desc(&c2)

}

func desc(c *Car) {
	fmt.Println("汽车介绍")
	fmt.Printf("企业品牌：%s\n", (*c).brand)
	fmt.Printf("企业颜色：%s\n", (*c).color)
	fmt.Printf("企业价格：%v元\n", (*c).price)
	fmt.Printf("企业长度：%v米\n", (*c).length)
}
