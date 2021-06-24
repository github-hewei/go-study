package main

import "fmt"

func f1(args ...interface{}) {
	//fmt.Println(args)
	fmt.Println(args...)
	for k, item := range args {
		fmt.Println(k, item)
	}
}

func f2(args ...[]int) {
	for k, item := range args {
		fmt.Println(k, item)
	}
}

func main() {
	f1(1, 2.0, "string", 'a')
	f2([]int{1, 2}, []int{2, 3}, []int{})
}
