package main

import (
	"fmt"
	"time"
)

func main() {

	// 10秒后执行
	tick := time.Tick(time.Duration(time.Second * 5))
	<-tick
	<-tick
	tick = nil

	go tick1()
	go tick2()

	time.Sleep(time.Duration(time.Second * 30))

}

// 第一种定时程序
func tick1() {
	for {
		fmt.Println(1, formatTime(time.Now()))
		time.Sleep(time.Duration(time.Second * 2))
	}
}

func tick2() {
	tick := time.Tick(time.Duration(time.Second * 2))
	for {
		fmt.Println(2, formatTime(time.Now()))
		<-tick
	}
}

func formatTime(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
