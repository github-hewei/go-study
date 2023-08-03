package main

import (
	"fmt"
	"os"
)

func main() {

	var path string

	fmt.Println("请输入要查询的目录")
	_, err := fmt.Scan(&path)
	if err != nil {
		panic(err)
	}
	fmt.Println("读取目录：", path)

	handle, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		panic(err)
	}

	fileInfo, err := handle.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, file := range fileInfo {
		fmt.Println("Name: ", file.Name())
		fmt.Println("Mode: ", file.Mode())
		fmt.Println("IsDir: ", file.IsDir())
		fmt.Println("Size: ", file.Size())
		fmt.Println()
	}

}
