package main

import "fmt"

func main() {

	//定义一个数组
	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr1)

	//定义一个切片
	arr2 := arr1[:]
	fmt.Println(arr2)

	//切片的语法[low, heigh, max] [低位下标, 高位下标（不含）, 最大下标]
	//切片的len = heigh - low
	//切片的cap = max - low
	arr3 := arr1[1:4:5]
	fmt.Println(arr3, "len = ", len(arr3), "cap = ", cap(arr3))

	fmt.Println(arr3[0])

	arr3[2] = 44
	//arr3[3] = 66 //溢出报错，因为超出了切片的空间长度 cap

	fmt.Println(arr3) //[2 3 44]

	//不指定数组长度，解释器根据赋值推算数组长度
	arr4 := [...]int{1, 2, 3}
	fmt.Println(arr4, "len =", len(arr4), "cap = ", cap(arr4))

}
