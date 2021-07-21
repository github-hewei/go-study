package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	slice2 := slice1[2:5:6]

	fmt.Println("len=", len(slice2), "cap=", cap(slice2), "slice2=", slice2)

	//切片中元素与源切片之间时引用传值
	slice2[0] = 999
	fmt.Println("len=", len(slice2), "cap=", cap(slice2), "slice2=", slice2)
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), "slice1=", slice1)

	//没有触发切片扩容的时候，切片中元素与源切片之间时引用关系
	slice2 = append(slice2, 6)
	slice2[0] = 888
	fmt.Println("len=", len(slice2), "cap=", cap(slice2), "slice2=", slice2)
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), "slice1=", slice1)

	//触发切片扩容之后，新切片会被重新分配内存地址，与源切片不存在引用关系
	slice2 = append(slice2, 7)
	slice2[0] = 777
	fmt.Println("len=", len(slice2), "cap=", cap(slice2), "slice2=", slice2)
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), "slice1=", slice1)

	//根据一个切片创建一个新切片，并不需要保持跟源切片的引用关系
	slice3 := append([]int{}, slice1[2:5:6]...)
	slice3[0] = 666
	fmt.Println("len=", len(slice3), "cap=", cap(slice3), "slice2=", slice3)
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), "slice1=", slice1)

	slice1 = nil

}
