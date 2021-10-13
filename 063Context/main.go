package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx1 := context.Background()
	fmt.Println(ctx1)

	ctx2 := context.TODO()
	fmt.Println(ctx2)

	ctx3, cancel1 := context.WithCancel(ctx1)
	cancel1()
	fmt.Println(ctx3)

	ctx4, cancel2 := context.WithTimeout(ctx1, time.Second * 10)
	cancel2()
	fmt.Println(ctx4)


}
