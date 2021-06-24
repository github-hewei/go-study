package main

import "fmt"

func main() {

	//普通函数
	fmt.Println(add(3, 4))

	//匿名函数
	f1 := func() {}
	fmt.Println(&f1)

	//匿名函数直接调用
	f2 := func(a, b int) int { return a + b }(1, 2)
	fmt.Println(f2)

	fmt.Println(f3(7, 8))
}

func add(a, b int) int {
	return a + b
}

func f3(a, b int) int {
	defer func() { fmt.Println(111) }()
	defer func() { fmt.Println(222) }()
	return a + b
}
