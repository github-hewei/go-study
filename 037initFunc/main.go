package main

import "fmt"

func init() {
	fmt.Println("Init 1 Runing")
}
func init() {
	fmt.Println("Init 2 Runing")
}
func init() {
	fmt.Println("Init 3 Runing")
	test()
}

func main() {
	fmt.Println("Main Runing")
	test()
}

func test() {
	fmt.Println("Test Runing")
}

/*
Init 1 Runing
Init 2 Runing
Init 3 Runing
Test Runing
Main Runing
Test Runing
*/
