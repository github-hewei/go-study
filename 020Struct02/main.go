package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
		sex  byte
	}

	man1 := Person{name: "zhangsan", sex: 'f'}
	fmt.Println(man1)

	type Test struct {
		man   Person
		count int
		slice [][]int
	}

	test1 := Test{Person{"tom", 20, 1}, 10, make([][]int, 5)}
	fmt.Println(test1)

}
