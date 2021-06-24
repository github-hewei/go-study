package main

import (
	"fmt"
)

func main() {

	//冒泡排序
	slice1 := []int{9, 1, 4, 8, 6, 6, 4, 4, 2, 7}
	bubbleSort(&slice1)
	fmt.Println(slice1)

	//选择排序
	slice2 := []int{9, 1, 4, 8, 6, 6, 4, 4, 2, 7}
	selectSort(&slice2)
	fmt.Println(slice2)

	//二分查找
	slice3 := []int{1, 3, 4, 5, 9, 13, 19, 45, 50, 77, 100, 120, 132, 144, 156, 177, 188}
	index := binarySearch(slice3, 7)
	fmt.Println(index)

}

/**
 * 二分查找
 */
func binarySearch(slice []int, index int) int {

	return 0
}

/**
 * 选择排序
 */
func selectSort(slice *[]int) {
	len := len(*slice)
	for i, ii := 0, len-1; i < ii; i++ {
		n := 0
		for j, jj := 0, len-i; j < jj; j++ {
			if (*slice)[j] > (*slice)[n] {
				n = j
			}
		}
		(*slice)[len-1-i], (*slice)[n] = (*slice)[n], (*slice)[len-1-i]
	}
}

/**
 * 冒泡排序
 */
func bubbleSort(slice *[]int) {
	len := len(*slice)
	for i, ii := 0, len-1; i < ii-1; i++ {
		for j, jj := 0, len-i-1; j < jj; j++ {
			if (*slice)[j] > (*slice)[j+1] {
				(*slice)[j], (*slice)[j+1] = (*slice)[j+1], (*slice)[j]
			}
		}
	}
}
