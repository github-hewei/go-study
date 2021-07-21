package main

import "fmt"

func main() {
	//Copy函数的使用
	//Copy函数的参数：参数1：目标函数，参数2：源函数

	slice1 := []int{9, 9}
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("slice1=", slice1, "slice2=", slice2) // slice1= [9 9] slice2= [1 2 3 4 5 6 7 8 9]
	//copy切片时会替换到指定的下标
	copy(slice2, slice1)
	fmt.Println("slice1=", slice1, "slice2=", slice2) // slice1= [9 9] slice2= [9 9 3 4 5 6 7 8 9]

	//Copy切片中的一部分
	slice3 := slice2[:]
	fmt.Println("slice2=", slice2, "slice3=", slice3) // slice2= [9 9 3 4 5 6 7 8 9] slice3= [9 9 3 4 5 6 7 8 9]

	//把789拷贝到前面一份
	copy(slice3[:3], slice3[6:])
	fmt.Println("slice2=", slice2, "slice3=", slice3) // slice2= [9 9 3 4 5 6 7 8 9] slice3= [9 9 3 4 5 6 7 8 9]

	//问题：为什么改变slice3变量，slice2变量也受到了影响发生了改变？
	//slice2 和 slice3 是同一个内存地址吗？
	fmt.Println(&slice1, &slice2, &slice3)

	fmt.Println("-------------------------------------")

	//测试一下
	slice4 := []int{1, 2, 3, 4, 5, 6}
	slice5 := slice4[:]
	slice6 := slice4[:3]
	fmt.Println(slice4, slice5, slice6)    // [1 2 3 4 5 6] [1 2 3 4 5 6] [1 2 3]
	fmt.Println(&slice4, &slice5, &slice6) // &[1 2 3 4 5 6] &[1 2 3 4 5 6] &[1 2 3]

	slice5[0] = 999
	slice6[2] = 999

	fmt.Println(slice4, slice5, slice6)    // [999 2 999 4 5 6] [999 2 999 4 5 6] [999 2 999]
	fmt.Println(&slice4, &slice5, &slice6) // &[999 2 999 4 5 6] &[999 2 999 4 5 6] &[999 2 999]

	//结果
	//可见切片之间是引用关系
	//哪怕切片的长度都不一致，索引相同的元素之前依然是引用关系
	//以下三个元素的内存地址是相同的：
	fmt.Println(&slice4[0], &slice5[0], &slice6[0]) // 0xc00000a390 0xc00000a390 0xc00000a390

	fmt.Println("-------------------------------------")

	//问题：
	//如何值传递复制一个切片到新的变量呢？
	slice7 := []int{1, 2, 3, 4, 5, 6}
	slice8 := []int{}

	// for _, item := range slice7 {
	// 	slice8 = append(slice8, item)
	// }

	//以上代码可简写为
	slice8 = append(slice8, slice7...)

	slice7[0] = 999
	slice8[1] = 999

	//此时两个切片中的元素没有引用关系了
	fmt.Println("slice7=", slice7, "slice8=", slice8) // slice7= [999 2 3 4 5 6] slice8= [1 999 3 4 5 6]
	fmt.Println(&slice7[0], &slice8[0])               // 0xc00013e090 0xc00013e0c0

	fmt.Println("-------------------------------------")
	//结论
	//在根据一个现有的切片创建一个新的切片的时候，有两种方式

	slice10 := []int{1, 2, 3, 4, 5, 6}

	//第一种：创建切片语法，切片中元素会引用传递
	slice11 := slice10[1:5]                               //取{2 3 4 5}
	fmt.Println("slice10=", slice10, "slice11=", slice11) // slice10= [1 2 3 4 5 6] slice11= [2 3 4 5]
	fmt.Println(&slice10[1], &slice11[0])                 // 0xc00013e0f8 0xc00013e0f8

	//第二种：通过append方式，切片中元素会值传递
	slice12 := make([]int, 0)
	slice12 = append(slice12, slice10[1:5]...)
	fmt.Println("slice10=", slice10, "slice12=", slice12) // slice10= [1 2 3 4 5 6] slice12= [2 3 4 5]
	fmt.Println(&slice10[1], &slice12[0])                 // 0xc00013e0f8 0xc000126080

}
