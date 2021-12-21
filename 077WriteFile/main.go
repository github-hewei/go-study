package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Create("d:/temp/qqq.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var buf []string

	rdx := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 4000000000; i++ {
		n := fmt.Sprintf("%d", randomInt(1000000000, 9000000000, rdx))
		buf = append(buf, n)
		if i % 1000000 == 0 {
			_, err := file.WriteString(strings.Join(buf, "\n") + "\n")
			if err != nil {
				panic(err)
			}
			fmt.Println(i)
			buf = buf[:0]
		}
	}

	fmt.Println("time:", time.Since(start))
}

// randomInt 获取指定区间内的随机数
func randomInt(min int, max int, r *rand.Rand) int {
	if min > max {
		panic("the min is greater than max!")
	}
	n := max - min + 1
	return r.Intn(n) + min
}
