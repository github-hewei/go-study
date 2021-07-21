package main

func main() {

	println("main")
	var (
		a int = 10
		b int = 20
	)
	println("a = ", a)
	println("b = ", b)
	test(&a, &b)
	println("a = ", a)
	println("b = ", b)
}

func test(a *int, b *int) {
	println("test")
	//交换两个变量的值
	*a, *b = *b, *a
	//*a在等号左边为左值，表示将要把数据写入到*a指向的内存区
	//*a在等号右边为右值，表示将*a内存区的数据读取出来赋值到等号左边对应的变量
}
