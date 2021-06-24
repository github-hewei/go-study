package main

import (
	"fmt"
)

func main() {

	//定义字符串
	var str1 string
	fmt.Println(str1)

	var str2 string = "Hello"
	fmt.Println(str2)

	str3 := "World"
	fmt.Println("len=", len(str3), "str3=", str3)

	//字符串是多个字符的切片
	str4 := "Hello World"

	//遍历
	for i, s := range str4 {
		fmt.Println(i, s) //i输出索引，s输出的是字符ascii编码的十进制数字
	}

	str5 := "你好，世界"
	for i, s := range str5 {
		fmt.Println(i, s) //每个汉字三个字节，索引为0-3-6...
	}

	//切片操作
	str6 := "Hello World"
	slice1 := str6[:5]
	slice2 := str6[6:]
	fmt.Println(slice1) //Hello 字符串类型
	fmt.Println(slice2) //World 字符串类型

	//通过索引遍历
	for i := len(str6); i > 0; i-- {
		fmt.Println(str6[i-1])
	}

}
