package main

import "fmt"

/**
 * 定义结构体
 */
type Person struct {
	name   string
	sex    int8
	age    uint8
	height float32
	weight float32
}

func main() {

	//先初始化再通过 . 赋值
	man1 := Person{}
	fmt.Println(man1)

	man1.name = "张三"
	man1.sex = -1
	man1.age = 29
	man1.height = 1.75
	man1.weight = 70.0

	fmt.Println(man1)

	//初始化按顺序赋值，必须按顺序并指定全部值
	man2 := Person{"李四", 1, 30, 0, 77.00}
	fmt.Println(man2)

	//初始化，并根据键名赋值，无需赋值全部
	man3 := Person{height: 180, name: "小明"}
	fmt.Println(man3)

}
