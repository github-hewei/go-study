package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 演示“指数退避策略”
	//if err := WaitForServer("localhost"); err != nil {
	//	log.Fatal(err)
	//}

	// 演示“指数退避策略”
	url := "localhost"
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	go func() {
		time.Sleep(30 * time.Second)
		cancel()
	}()

	err := WaitForCallbackWithContext(ctx, func() error {
		_, err := http.Head(url)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}

// WaitForCallbackWithContext 尝试执行cb函数，直到cb函数返回nil或者超时
func WaitForCallbackWithContext(ctx context.Context, cb func() error) error {
	tries := 0

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out after %s", ctx.Err())
		case <-time.After(time.Second << uint(tries)):
			if err := cb(); err != nil {
				log.Printf("server not responding (%s); retrying....\n", err)
				tries++
			} else {
				return nil
			}
		}
	}
}

// WaitForServer 尝试链接url对应的服务器
// 在一分钟之内，使用指数退避策略重试连接
// 所有的尝试失败后，返回错误
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)

		if err == nil {
			return nil
		}

		log.Printf("server not responding (%s); retrying....\n", err)
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
