package main

import "fmt"

func main() {

	//通过make函数创建一个集合(映射)
	map1 := make(map[int]string)
	fmt.Println("len=", len(map1), map1)

	//写入值
	map1[1] = "green"
	fmt.Println("len=", len(map1), map1)

	//通过赋值运算符创建一个集合
	map2 := map[string]int{}
	fmt.Println("len=", len(map2), map2)

	//写入值
	map2["money"] = 100
	fmt.Println("len=", len(map2), map2)

	//创建集合并赋值
	map3 := map[int]int{1: 2, 3: 4}
	fmt.Println("len=", len(map3), map3)

	//复杂集合
	map4 := map[int][]int{1: {1, 2, 3}, 2: {1, 1, 1}}
	fmt.Println("len=", len(map4), map4)

}
