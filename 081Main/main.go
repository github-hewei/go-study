package main

import (
	"fmt"
	"net/http"
)

func main()  {
	// 下面代码实现了一个简单的http服务端，但有一个问题
	// 如果程序因为代码问题而意外退出(例如panic), 无法和kill这种人为强制杀死的情况进行区分
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", mux)

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	panic("test")
}
