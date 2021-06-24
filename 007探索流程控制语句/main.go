package main

import "fmt"

func main() {

	fmt.Println("----------")
	_if()

}

func _if() {
	a := 1
	b := 2

	//表达式判断
	if b > a {
		fmt.Println("b > a")
	}

	//OR条件
	if b < a || true {
		fmt.Println("||")
	}

	//AND条件
	if true && b > a {
		fmt.Println("&&")
	}

	//通过()控制运算优先级
	if (true) && (true || false) {
		fmt.Println("()")
	}
}
