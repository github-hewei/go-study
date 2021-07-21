package main

import (
	"fmt"
	"unsafe"
)

type s1 struct {
	a [1000000]int
}

type s2 struct {
	a *[1000000]int
}

func (s s1) f1() {
	fmt.Println("f1", unsafe.Sizeof(s))
}

//在定义结构体的函数的时候，结构体前面*的作用
//避免重复复制结构体的数据值，浪费资源
func (s *s1) f2() {
	fmt.Println("f2", unsafe.Sizeof(s))
}

func main() {

	a := [1000000]int{}

	fmt.Println("a size", unsafe.Sizeof(a)) //8000000

	ss1 := s1{a}
	fmt.Println("ss1 size", unsafe.Sizeof(ss1)) //80000

	ss2 := s2{&a}
	fmt.Println("ss2 size", unsafe.Sizeof(ss2)) //8

	ss1.f1()
	ss1.f2()

}
