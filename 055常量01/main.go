package main

import "fmt"

const PI = 3.1415926

const NAME string = "abcdef"

// 常量只支持三种数据类型，数字字符串和布尔值
const A, B, C = 1, "2", true

const (
	a = iota
	b
	c
	d
)

const (
	e = 1 << iota
	f = 2 << iota
	g
	h
)

func main() {

	fmt.Println(PI)
	fmt.Println(NAME)
	fmt.Println(A, B, C)
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h)

}
