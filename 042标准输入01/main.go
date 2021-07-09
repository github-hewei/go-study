package main

import (
	"fmt"
	"os"
)

func main() {

	for key, item := range os.Args {
		fmt.Println(key, item)
	}

}
