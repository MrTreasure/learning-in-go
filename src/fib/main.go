package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	result := <-fib(50)

	fmt.Printf("fib computed %d, spend: %s, get %d", 20, time.Since(start).String(), result)
}

func fib(n int) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)
		if n <= 2 {
			result <- 1
			return
		}
		result <- <-fib(n-1) + <-fib(n-2)
	}()
	return result
}
