package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	m1 := map[int]string{}
	fmt.Printf("地址:%p\n", m1)

	m1[1] = "red"
	m1[10] = "blue"

	fmt.Printf("地址:%p\n", m1)

	p1 := Person{}
	fmt.Printf("地址:%p\n", &p1)

	p1.age = 20
	p1.name = "tom"

	//结构体的内存地址就是结构体中第一个元素的内存地址
	fmt.Printf("地址:%p\t地址:%p\t地址:%p\n", &p1, &p1.name, &p1.age)

}
