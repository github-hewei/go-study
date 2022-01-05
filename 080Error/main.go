package main

import (
	"fmt"
	"github.com/pkg/errors"
)

// MyErrorCode 全局的错误号类型，用于Api调用之间传递
type MyErrorCode int

// 全局错误号具体定义
const (
	ErrorBookNotFoundCode MyErrorCode = iota + 1
	ErrorBookHasBeenBorrowedCode
)

// 内部的错误map，用来对应错误号和错误信息
var errCodeMap = map[MyErrorCode]string{
	ErrorBookNotFoundCode:        "Book was not found",
	ErrorBookHasBeenBorrowedCode: "Book has been borrowed",
}

// Sentinel Error: 即全局定义的 Static 错误变量
// 注意，这里的全局error 是没有保存堆栈信息的，所以需要在初始调用处使用 errors.Wrap
var (
	ErrorBookNotFound        = NewMyError(ErrorBookNotFoundCode)
	ErrorBookHasBeenBorrowed = NewMyError(ErrorBookHasBeenBorrowedCode)
)

func NewMyError(code MyErrorCode) *MyError {
	return &MyError{
		Code:    code,
		Message: errCodeMap[code],
	}
}

type MyError struct {
	Code    MyErrorCode
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func main() {

	books := []string{
		"Hamlet",
		"Jane Eyre",
		"War and Peace",
	}

	for _, bookName := range books {
		fmt.Println("start", bookName)

		err := borrowOne(bookName)
		if err != nil {
			fmt.Printf("%+v\n", err)
			// panic(err)
		}

		fmt.Println("end", bookName)
	}

}

func borrowOne(bookName string) error {
	// Step1: 找书
	err := searchBook(bookName)

	// Step2: 处理
	// 特殊业务场景：如果发现书被借走了，下次再来就行了，不需要作为错误处理
	if err != nil {
		// 提取 error 这个 interface 底层的错误码，一般在api的返回前才提取
		// As 获取错误的具体实现

		var myError = new(MyError)
		if errors.As(err, &myError) {
			fmt.Printf("error code is %d, message is %s\n", myError.Code, myError.Message)
		}

		// 特殊逻辑，对应场景2，指定错误 ErrorBookHasBeenBorrowed 时，打印即可，不返回错误
		// Is - 判断错误是否为指定类型
		if errors.Is(err, ErrorBookHasBeenBorrowed) {
			fmt.Printf("book %s has been borrowed, I will come back later!\n", bookName)
			err = nil
		}

	}

	return err
}

func searchBook(bookName string) error {
	// 下面两个error都是不带堆栈信息的，所以初次调用得用 Warp方法
	// 如果已有 堆栈信息，应调用 WithMessage 方法
	if len(bookName) > 10 {
		return errors.Wrapf(ErrorBookNotFound, "bookName is `%s`", bookName)
	} else if len(bookName) > 8 {
		return errors.Wrapf(ErrorBookHasBeenBorrowed, "bookName is `%s`", bookName)
	} else {
		return nil
	}
}
