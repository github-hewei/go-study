package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 大文件拷贝

	// 读取文件
	readHandle, err := os.Open("d:/temp/WeChatSetup.exe")
	if err != nil {
		fmt.Println(err)
	}
	defer readHandle.Close()

	// 创建新文件
	writeHandle, err := os.Create("d:/temp/WechatCopy.exe")
	if err != nil {
		fmt.Println(err)
	}
	defer writeHandle.Close()

	buf := make([]byte, 4096)

	for {
		n, err := readHandle.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读取完成")
			break
		}
		fmt.Println(n)
		_, err = writeHandle.Write(buf[:])
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("完成")

}
