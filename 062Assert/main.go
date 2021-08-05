package main

import (
	"fmt"
	"reflect"
)

type ErrorField struct {
	msg string
	field string
	tag string
}
func (ef *ErrorField) Field() string {
	return ef.field
}
func (ef *ErrorField) Tag() string {
	return ef.tag
}
func (ef *ErrorField) Msg() string {
	return ef.msg
}

// MyErrors 是 ErrorField 结构体的切片
type MyErrors []ErrorField

// 实现 error 接口的 Error 方法
func (err MyErrors) Error() string {
	var str string
	for _, e := range err {
		str += e.Msg() + "; "
	}
	return str
}

// 测试方法，用来返回一个 error 错误
func test() error {
	var err = MyErrors{}
	err = append(err, ErrorField{"MSG1", "FIELD1", "TAG1"}, ErrorField{"MSG2", "FIELD2", "TAG2"})
	return err
}

func main() {

	// 断言的基本用法
	var i interface{} = "abc"
	fmt.Printf("%T\n", i)
	fmt.Println(reflect.ValueOf(i).Kind())

	a, b := i.(int) // a 是变量的值,类型是断言后的变量，b 是断言的结果
	fmt.Println(a, b)
	fmt.Printf("%T\n", a)

	// 扩展
	err := test()
	if err != nil {
		// panic(err)
		fmt.Println(err) // err 的值是调用 Error 方法返回的 string
	}

	// 理解一下 err.(MyErrors) 的作用
	// err 未断言的情况下就是 error，断言后就是 MyErrors
	err = test()
	if err != nil {
		for _, val := range err.(MyErrors) {
			fmt.Printf("msg: %s, field: %s, tag: %s\n", val.Msg(), val.Field(), val.Tag())
		}
	}

}
