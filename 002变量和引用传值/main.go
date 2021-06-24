package main

import "fmt"

//变量声明，声明变量时需要指定数据类型
var (
	name string
	age  int
	isOk bool
)

//全局变量定义后可以不在代码中使用，可以编译成功
//函数内部的局部变量定义后必须使用，不使用不能编译成功

func main() {
	name = "HEWEI"
	age = 27
	isOk = true
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)

	fmt.Println("------------------------------")

	//指针操作，引用传值
	var a int

	var b *int = &a

	fmt.Println(b)  //当前为引用的a的内存地址
	fmt.Println(*b) //解引用，取值运算符

	a = 100

	fmt.Println(b)
	fmt.Println(*b) //当变量a的值改变时，通过b的内存地址能够读取到变量a改变之后的值

	*b = 1000 //当通过b存储的内存地址改变值时，变量a的值也能够更新

	fmt.Println(a)  //1000
	fmt.Println(b)  //0xc000...
	fmt.Println(*b) //1000

	fmt.Println("------------------------")

	var c int = 1

	var d = new(int) //申请一块内存地址准备存储int类型数据
	fmt.Println(d)
	d = &c //把d的地址指向c变量的内存地址

	fmt.Println(d)  //此时内存地址改变
	fmt.Println(*d) //1

	c = 2000
	fmt.Println(d)  //0xc000...
	fmt.Println(*d) //2000

}
