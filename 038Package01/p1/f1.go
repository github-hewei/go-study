package p1

import "fmt"

//包内部的变量首字母大写表示此变量是Public公开的
//可以在包外部访问到
//首字母小写表示此变量是Private私有的
//只可以在包的内部访问

var a string // 对外不可见，只能在包内部使用
var A string // public 公开的，对外可见

func f1() {
	fmt.Println("private func f1() runing")
}

func Run1() {
	fmt.Println("Run 1")
	f1()
	fmt.Println("Private a: ", a)
	fmt.Println("Public A: ", A)
}
