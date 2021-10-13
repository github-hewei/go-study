package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.Now() // 当前时间

	fmt.Println(t1.Year())   // 获取年份
	fmt.Println(t1.Month())  // 获取月份 英文
	fmt.Println(t1.Day())    // 获取日期
	fmt.Println(t1.Hour())   // 获取小时
	fmt.Println(t1.Minute()) // 获取分钟
	fmt.Println(t1.Second()) // 获取秒数 没有导0

	// 格式化时间为 Y-m-d H:i:s
	fmt.Println(t1.Format("2006-01-02 15:04:05"))

	// 解析时间
	t2, err := time.Parse("2006-01-02 15:04:05", "1994-11-22 08:01:02")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2.Format("2006-01-02 15:04:05"))

	// 解析日期
	t3, err := time.Parse("2006-01-02", "1994-11-22")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3.Format("2006-01-02 15:04:05"))

	// 获取30分钟之后的时间
	t4 := t1.Add(time.Minute * 30) // 30分钟之后的时间
	fmt.Println(t4.Format("2006-01-02 15:04:05"))

	// 获取 5个月 零 1天 后的时间
	t5 := t1.AddDate(0, 5, 1)
	fmt.Println(t5.Format("2006-01-02 15:04:05"))
	fmt.Println(t5.Format("2006-01-02"))

	// 比较时间
	t6 := time.Now().Add(time.Second)

	fmt.Println(t6.Before(t1)) // false
	fmt.Println(t6.After(t1)) // true

	// 通过时间获取时间戳
	t7 := t1.Unix()
	fmt.Println(t7)

	// 时间戳转时间
	t7 += 86400
	t8 := time.Unix(t7, 0)
	fmt.Println(t8.Format("2006-01-02 15:04:05"))


	// 判断一个日期是否大于当前时间

	str := "2021-10-14"

	v1, err := time.Parse("2006-01-02", str)
	fmt.Println(v1.Format("2006-01-02 15:04:05"))

	if v1.After(time.Now()) {
		fmt.Println("大于当前时间")
	}

}
