package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Users struct {
	Name  string `validate:"required"`
	Age   int    `validate:"required,min=1,max=150"`
	Email string `validate:"required,email"`
}

var (
	ValidateErrorsMsg = map[string]string{
		"Name.required":  "姓名不能为空",
		"Age.required":   "年龄不能为空",
		"Age.min":        "年龄不能小于1",
		"Age.max":        "年龄不能大于150",
		"Email.required": "邮箱不能为空",
		"Email.email":    "邮箱格式有误",
	}
)

func main() {

	// 验证类型参考链接
	// https://pkg.go.dev/github.com/go-playground/validator/v10

	validate := validator.New()

	email := "hemailads@163.com"
	err := validate.Var(email, "required,email")
	if err != nil {
		panic(err) // Key: '' Error:Field validation for '' failed on the 'email' tag
	}
	// fmt.Println(err)

	age := 1
	err = validate.Var(age, "required,min=0,max=150")
	if err != nil {
		panic(err)
	}

	user := &Users{
		Name:  "",
		Age:   160,
		Email: "",
	}

	err = validate.Struct(user)
	// fmt.Printf("%#v\n", err)

	// !!!
	// 062Assert/main.go 中的例子解释了此处 err.(validator.ValidationErrors) 断言的作用

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			val, ok := ValidateErrorsMsg[e.Field()+"."+e.Tag()]
			if ok {
				fmt.Println(val)
			} else {
				fmt.Println("参数有误")
			}
		}
	}

	fmt.Println(*user)
}
