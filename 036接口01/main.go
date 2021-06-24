package main

import "fmt"

//定义接口
type test interface {
	m1() string
	m2() string
}

type s1 struct {
	a1 int
	a2 int
}

func (s *s1) m1() string {
	return "Hello M1"
}

func (s *s1) m2() string {
	return "Hello M2"
}

func demo() test {
	s := s1{1, 2}
	return &s
}

func dump(arg ...interface{}) {
	fmt.Println(arg...)
	for k, v := range arg {
		fmt.Println(k, v)
	}
}

func main() {

	//实现了某个接口中的方法，程序就判断实现了某个接口
	fmt.Println(demo())

	dump(1, "string", 'a', 1.22)

}
