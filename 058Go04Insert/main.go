package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

func main() {

	// 纳秒 time.Nanosecond => 1
	// 微秒 time.Microsecond -> 1000 * time.Nanosecond -> 1000
	// 毫秒 time.Millisecond -> 1000 * time.Microsecond -> 1000000
	// 秒 time.Second -> 1000 * time.Millisecond -> 1000000000
	// 分钟 time.Minute -> 60 * time.Second -> 60000000000
	// 小时 time.Hour -> 60 * time.Minute -> 3600000000000

	t := time.Now().UnixNano()

	db, err := sql.Open("mysql", "root:root@tcp(localhost:13301)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	// 高并发下按需开启。否则会在短时间内产生大量 tcp 短链接，导致cpu爆满，甚至耗尽系统最大tcp连接数
	db.SetMaxOpenConns(50) // 长连接数 测试结果：50个连接，10W条写入SQL，耗时60s
	db.SetMaxIdleConns(10) // 闲置连接数

	// 同步方式写入10000条数据，耗时 61700.0556ms
	//for i := 0; i < 10000; i++ {
	//	ret, err := db.Exec("insert into `test` (`text`) values (?)", time.Now().UnixNano())
	//	if err != nil {
	//		panic(err)
	//	}
	//	id, err := ret.LastInsertId()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(id)
	//}

	// 17208.927
	//for j := 0; j < 100; j ++ {
	//	num := 100
	//	c := make(chan int64, num)
	//	for i := 0; i < num; i++ {
	//		i := i
	//		go func(db *sql.DB, c chan int64) {
	//			fmt.Println(i)
	//			ret, err := db.Exec("insert into `test` (`text`) values (?)", time.Now().UnixNano())
	//			if err != nil {
	//				panic(err)
	//			}
	//			id, err := ret.LastInsertId()
	//			if err != nil {
	//				panic(err)
	//			}
	//			c <- id
	//		}(db, c)
	//	}
	//
	//	for i := 0; i < num; i++ {
	//		fmt.Println( <- c )
	//	}
	//
	//}

	var ws = sync.WaitGroup{}
	num := 100000

	ws.Add(num)

	c1 := make(chan int64, 100)

	// 并发执行10000条写入SQL耗时6s左右
	for i := 0; i < num; i++ {
		c1 <- 1
		go func(db *sql.DB, c1 chan int64) {
			ret, err := db.Exec("insert into `test` (`text`) values (?)", time.Now().UnixNano())
			if err != nil {
				panic(err)
			}
			id, err := ret.LastInsertId()
			if err != nil {
				panic(err)
			}
			<-c1
			fmt.Println(id)
			ws.Done()
		}(db, c1)
	}

	ws.Wait()

	fmt.Println(float64(time.Now().UnixNano()-t) / 1e6)
}
