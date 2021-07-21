package main

import (
	"fmt"
	"os"
)

func main() {

	//通过defer和recover捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常", err)
		}
	}()

	//通过 os.Create 创建文件
	f1, err := os.Create("d:/temp/temp.go")
	if err != nil {
		//通过panic抛出异常
		panic(err)
	}

	fmt.Println(f1)

	//往文件中写入10行字符串
	for i := 0; i < 10; i++ {
		n, err := f1.WriteString("Hello World\n")
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
	}

	//关闭文件资源句柄
	f1.Close()

}
