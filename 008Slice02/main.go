package main

import "fmt"

func main() {

	//创建一个空切片
	var slice1 []int
	fmt.Println("len=", len(slice1), "cap=", cap(slice1), slice1)

	//往切片中追加元素，切片会自动扩容
	for i := 0; i < 3; i++ {
		slice1 = append(slice1, i)
		fmt.Println("len=", len(slice1), "cap=", cap(slice1), slice1)
	}

	//创建切片并赋值
	slice2 := []int{1, 2, 3}
	fmt.Println("len=", len(slice2), "cap=", cap(slice2), slice2)

	//创建切片make
	slice3 := make([]int, 5, 10)
	fmt.Println("len=", len(slice3), "cap=", cap(slice3), slice3)
	slice3[4] = 999
	fmt.Println("len=", len(slice3), "cap=", cap(slice3), slice3)

	//去除切片中的空字符串
	slice4 := []string{"red", "", "yellow", "", "blue", "green", ""}

	//1. 值传递方式
	slice5 := func1(slice4)
	fmt.Println("len=", len(slice5), "cap=", cap(slice5), slice5)

	//2. 引用传递方式，传slice4的内存地址过去，所以也不需要接收返回值
	func2(&slice4)
	fmt.Println("len=", len(slice4), "cap=", cap(slice4), slice4)

	//3. 引用传递方式，将返回值赋值到第二个函数实参
	//往slice4中追加几个空字符串，用来测试
	slice4 = append(slice4, "", "", "")
	//使用new创建一个切片的内存地址，用来接收新的切片结果
	slice6 := new([]string)
	fmt.Println(slice6, &slice6)
	//slice6已经是一个引用变量了，所以这里不用&符号了
	result := func3(slice4, slice6)
	fmt.Println("len=", len(slice4), "cap=", cap(slice4), slice4)
	fmt.Println("len=", len(*slice6), "cap=", cap(*slice6), *slice6)
	fmt.Println("len=", len(*slice6), "cap=", cap(*slice6), slice6)
	fmt.Println(result)
}

/**
 * 去除参数1中的空字符串，将新值写入到引用传值的参数2，并返回一个布尔类型的返回值
 */
func func3(slice1 []string, slice2 *[]string) bool {
	for i, l := 0, len(slice1); i < l; i++ {
		if slice1[i] != "" {
			//(*slice2)[i] = slice1[i] // index out of range
			*slice2 = append(*slice2, slice1[i])
		}
	}
	return true
}

/**
 * 在原切片基础上操作并且函数调用方式为引用传递
 */
func func2(slice1 *[]string) {
	i := 0
	for _, s := range *slice1 {
		if s != "" {
			(*slice1)[i] = s
			i++
		}
	}
	*slice1 = (*slice1)[:i]
}

/**
 * 去除切片中的空字符串，将新切片数组返回
 */
func func1(slice1 []string) []string {
	slice2 := make([]string, 0)
	for i, length := 0, len(slice1); i < length; i++ {
		if slice1[i] != "" {
			slice2 = append(slice2, slice1[i])
		}
	}
	return slice2
}
