package main

import (
	"fmt"
)

func main() {

	map1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(map1)

	//根据键名删除集合中的元素
	delete(map1, 3)

	fmt.Println(map1)

	//slice1 := []int{1, 2, 3}
	//delete(slice1, 0) //报错。delete函数是针对集合的，不能删除切片中的元素

	//那么删除切片中的某个元素怎么删除
	slice1 := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), slice1)
	delIndex := 3

	//删除值不重置索引
	slice1[delIndex] = ""
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), slice1)

	//删除并重置索引
	slice1 = append(slice1[:delIndex], slice1[delIndex+1:]...)
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), slice1)

	//根据指定切片的值来删除元素
	delValue := "e"
	slice2 := make([]string, 0)
	for _, val := range slice1 {
		if val != delValue {
			slice2 = append(slice2, val)
		}
	}
	fmt.Println("len=", len(slice2), "cap=", cap(slice2), slice2)
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), slice1)
	//释放 slice1
	slice1 = nil
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), slice1)

}
