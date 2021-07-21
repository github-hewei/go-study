package main

import "fmt"

func main() {

	slice1 := []int{}

	if slice1 == nil {
		fmt.Println("[]int{} = nil")
	}

	slice2 := make([]int, 0)
	if slice2 == nil {
		fmt.Println("[]int{} = nil")
	}

	slice3 := [10000][0]int{}

	for i := range slice3 {
		fmt.Println(i)
	}

}
