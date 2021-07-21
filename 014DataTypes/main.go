package main

import "fmt"

var (

	//整形
	v4 uint8  // 无符号 8 位整型 (0 到 255)
	v5 uint16 // 无符号 16 位整型 (0 到 65535)
	v6 uint32 // 无符号 32 位整型 (0 到 4294967295)
	v7 uint64 // 无符号 64 位整型 (0 到 18446744073709551615)

	v8  int8  // 有符号 8 位整型 (-128 到 127)
	v9  int16 // 有符号 16 位整型 (-32768 到 32767)
	v10 int32 // 有符号 32 位整型 (-2147483648 到 2147483647)
	v11 int64 // 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

	//浮点类型
	v12 float32    // IEEE-754 32位浮点型数
	v13 float64    // IEEE-754 64位浮点型数
	v14 complex64  // 32 位实数和虚数
	v15 complex128 // 64 位实数和虚数

	//其他数字类型
	v16 byte    // 类似 uint8
	v17 rune    // 类似 int32
	v18 uint    // 32 或 64 位
	v19 int     // 与 uint 一样大小
	v20 uintptr // 无符号整型，用于存放一个指针

)

func main() {
	//基本数据类型

	fmt.Println("--- 布尔型数据 ---")
	var v1 bool
	fmt.Println(v1)
	var v2 bool = true
	fmt.Println(v2)
	v3 := false
	fmt.Println(v3)

	//字符串类型
	var v21 string = "abc"
	fmt.Println(v21)

	v22 := "def"
	fmt.Println(v22)

	v23 := string("gjt")
	fmt.Println(v23)

}
