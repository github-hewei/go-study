package main

import "fmt"

func main() {

	//创建集合
	map1 := map[int]string{1: "blue", 2: "red", 3: "green"}

	//遍历集合
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	//判断某个键是否存在于集合中
	value1, has1 := map1[0]
	fmt.Println("value=", value1, "has=", has1)
	value2, has2 := map1[2]
	fmt.Println("value=", value2, "has=", has2)

	//封装一个函数，将切片中的数据去重并统计出现次数
	str := []string{"red", "blue", "green", "red", "green", "yellow", "red"}
	map2 := arrayQueue(str)
	fmt.Println(map2)

}

/**
 * 统计切片中的单词重复次数
 */
func arrayQueue(slice []string) map[string]int {
	map1 := make(map[string]int)
	for _, item := range slice {
		// if _, has := map1[item]; has {
		// 	map1[item]++
		// } else {
		// 	map1[item] = 1
		// }

		//以上代码可以简写为如下代码，因为集合map1，默认值为int类型的0
		map1[item]++
	}
	return map1
}
