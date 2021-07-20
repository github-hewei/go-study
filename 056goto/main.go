package main

import (
	"fmt"
	"math/rand"
)

func main() {
BEGIN:
	n := rand.Intn(10)
	fmt.Println(n)
	if n != 5 {
		goto BEGIN
	}
}
