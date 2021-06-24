package main

import "fmt"

//------------------------------------------------------------------------
// ../033Error01/main.go 文件中的代码是从菜鸟教程copy来的。
// 感觉对自定义的异常实现的有点别扭，参考了 errors.New() 方法中的代码之后
// 修改成本下面的样子
//------------------------------------------------------------------------

//定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

//实现 `error` 接口
//error接口的定义:
// type error interface {
//  Error() string
// }
//如何实现接口呢???
//Go 中没有 implements 关键字，判断一个类型是否实现了一个接口是完全是自动地。
//所有定义了 `Error() string` 方法的类型我们称它实现了 error 接口。

func (de *DivideError) Error() string {
	strFormat := `
Cannot proceed, the divider is zero.
dividee: %d
divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

//定义 int 类型除法运算的函数
func Divide(varDividee, varDivider int) (result int, err error) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		err = &dData
	} else {
		result = varDividee / varDivider
	}
	return
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常:", err)
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
