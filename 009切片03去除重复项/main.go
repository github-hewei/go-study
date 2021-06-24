package main

import "fmt"

func main() {

	slice1 := []string{"red", "yellow", "red", "blue", "green", "blue"}

	//去除重复项，值传递版本
	slice2 := func1(slice1)
	fmt.Println("len=", len(slice2), "cap=", cap(slice2), slice2)

	//去除重复项，引用传递版本
	func2(&slice1)
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), slice1)

}

/**
 * 去除切片数组中的重复项，引用传值版本
 */
func func2(slice1 *[]string) {
	slice2 := make([]string, 0)
	for _, a1 := range *slice1 {
		flag := true
		for _, a2 := range slice2 {
			if a1 == a2 {
				flag = false
				break
			}
		}
		if flag {
			slice2 = append(slice2, a1)
		}
	}
	*slice1 = slice2
}

/**
 * 去除切片数组中的重复项元素
 */
func func1(slice1 []string) []string {
	slice2 := make([]string, 0)
	for _, a1 := range slice1 {
		flag := true
		for _, a2 := range slice2 {
			if a1 == a2 {
				flag = false
				break
			}
		}
		if flag {
			slice2 = append(slice2, a1)
		}
	}
	return slice2
}
