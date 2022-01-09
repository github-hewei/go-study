package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	// 引入 Signal
	// kill工具是Linux系统中，往进程发送一个信号。所以我们去实现，捕捉信号的功能

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello2)

	// http服务改造成异步
	go http.ListenAndServe(":8080", mux)

	fmt.Println(<-sig)
}

func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello2")
	panic("test")
}
