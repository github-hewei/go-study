package main

import (
	"errors"
	"fmt"
)

//定义 int 类型除法运算的函数
func Divide(varDividee, varDivider int) (result int, err error) {
	if varDivider == 0 {
		err = errors.New("除数不能为0")
	} else {
		result = varDividee / varDivider
	}
	return
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常：", err)
		}
	}()

	if result, err := Divide(100, 10); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

	if result, err := Divide(100, 0); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

}
