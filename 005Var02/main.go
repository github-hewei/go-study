package main

import "fmt"

func main() {

	//声明变量
	var (
		a int
		b string
	)

	//赋值
	a = 10
	b = "abc"

	fmt.Println(a, b)

	// := 声明变量并赋值
	// := 操作符只能在函数内部使用
	// 系统会自动判断数据类型
	c := 123
	d := "efg"
	e := [3]string{"a", "b", "c"}

	fmt.Println(c, d, e)

}
