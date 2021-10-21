package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background()) // 根节点和一级节点
	defer cancel()

	go test(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("11111")
	case <-time.After(time.Second * 10): // 10秒后退出
		fmt.Println("main timeout")
	}

}

func test(parent context.Context) {
	ctx, cancel := context.WithCancel(parent) // 二级节点
	defer cancel()

	go test2(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("test1 cancel")
	case <-time.After(time.Second * 5): // 5秒后退出，推出后此节点的所有子节点全部退出，执行 Done 分支
		fmt.Println("test1 timeout")
	}

}

func test2(parent context.Context) {
	ctx, cancel := context.WithCancel(parent) // 三级节点
	defer cancel()

	go test3(ctx, time.Second * 1) // 1秒后退出，输出 test3 timeout
	go test3(ctx, time.Second * 2) // 2秒后退出，输出 test3 timeout
	go test3(ctx, time.Second * 9) // 9秒后退出，父级节点5秒后退出，所以5秒后输出 test3 cancel

	select {
	case <-ctx.Done():
		fmt.Println("test2 cancel")
	case <-time.After(time.Second * 8): // 8秒后退出，但是父级节点5秒后退出，所以5秒后输出 test cancel
		fmt.Println("test2 timeout")
	}

}

func test3(parent context.Context, t time.Duration) {
	ctx, cancel := context.WithCancel(parent) // 四级节点
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("test3 cancel")
	case <-time.After(t):
		fmt.Println("test3 timeout")
	}
}
