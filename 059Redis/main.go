package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
	"time"
)

var (
	redisPool *redis.Pool
)

func main() {

	// Redis 连接池
	redisPool = &redis.Pool{
		MaxIdle     : 1,
		MaxActive   : 3,
		IdleTimeout : 6 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:16379")
		},
	}

	// 通过通道限制协程的并行数量
	ch := make(chan int, 3)

	for i := 0; i < 10; i++ {
		ch <- 1
		go func(ch chan int) {
			conn := redisPool.Get()
			// 必须断开连接才会释放回连接池
			defer conn.Close()
			res, err := conn.Do("PING")
			// 当连接池耗尽时执行命令会返回错误：panic: redigo: connection pool exhausted
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
			<- ch
		}(ch)
	}


	// 连接Redis
	conn, err := redis.Dial("tcp", "127.0.0.1:16379")
	if err != nil {
		panic(err)
	}

	// 关闭Redis连接
	defer conn.Close()

	// 执行Redis命令
	ret, err := conn.Do("SELECT", 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	// 执行SET
	ret, err = conn.Do("SET", "kkk", "vvv")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	// 执行GET
	ret, err = conn.Do("GET", "kkk")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	// redis.Args
	a := redis.Args{1, 2, 3}
	fmt.Println(a)
	a = a.Add([]int{1, 2, 3}, "abc")
	fmt.Println(a)
	a = a.AddFlat(map[string]int{"a": 1, "b": 2})
	fmt.Println(a, reflect.ValueOf(a).Kind())
	for _, val := range a {
		fmt.Println(val, reflect.ValueOf(val).Kind())
	}

}
