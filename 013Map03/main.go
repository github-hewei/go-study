package main

import "fmt"

func main() {
	//创建一个函数，传入一个map集合，返回两个切片数组分别包含map的key和value

	map1 := map[int]string{100: "red", 200: "blue", 300: "green", 400: "yellow"}

	keys, values := map2slice(map1)
	fmt.Println(keys)
	fmt.Println(values)

}

/**
 * 传如一个map，返回两个切片数组分别包含所有键和所有值
 */
func map2slice(arg map[int]string) ([]int, []string) {
	slice1 := make([]int, 0)
	slice2 := make([]string, 0)
	for key, value := range arg {
		slice1 = append(slice1, key)
		slice2 = append(slice2, value)
	}
	return slice1, slice2
}
