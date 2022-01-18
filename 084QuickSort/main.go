package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// 构造一个随机数切片
	list := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = rand.Intn(1000000)
	}

	QuickSort(list)
	fmt.Println(list)

}

// QuickSort 快速排序算法
func QuickSort(list []int) {
	max := len(list) - 1

	if max < 1 {
		return
	}

	head, tail, index := 0, max, max
	pivot := list[0]

	for head < tail {
		if list[index] > pivot {
			list[tail] = list[index]
			tail--
			index = tail
		} else {
			list[head] = list[index]
			head++
			index = head
		}
	}

	list[index] = pivot
	QuickSort(list[:index])
	QuickSort(list[index+1:])
}
