package main

import (
	"fmt"
	"math"
)

// nextPrime returns the next prime number
// 获取下一个质数
func nextPrime(n int) int {
	if n < 2 {
		return 2
	}

	for {
		n++
		if isPrime(n) {
			return n
		}
	}
}

// isPrime checks if a number is prime
// 判断一个数是不是质数
// 质数（prime number）的定义是，在大于1的自然数中，除了1和它本身以外不再有其他因数的数，质数只能被1和自身整除。
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// primeGenerator generates prime numbers
// 生成质数的迭代器
func primeGenerator(done <-chan bool, operation func(int) int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		num := 1
		for {
			num = operation(num)
			select {
			case <-done:
				return
			case ch <- num:
			}
		}
	}()
	return ch
}

func main() {
	done := make(chan bool)

	i := 0
	for p := range primeGenerator(done, nextPrime) {
		i++
		fmt.Println("Prime: ", p)

		if i >= 10 {
			break
		}
	}

	done <- true
}
