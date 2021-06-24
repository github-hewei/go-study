package main

import "fmt"

func main() {

	//定义一个数组，长度为10，存储的数据类型为int
	var a1 [10]int
	fmt.Println(a1)
	//赋值
	a1[0] = 1
	a1[9] = 2
	fmt.Println(a1)

	//var a2 [3]int{1, 2, 3} //语法错误
	a2 := [5]int{1, 2, 3}
	fmt.Println(a2)

	//通过len函数获取数组长度
	fmt.Println("a2 len = ", len(a2)) //数组长度
	fmt.Println("a2 cap = ", cap(a2)) //数组空间大小

}
