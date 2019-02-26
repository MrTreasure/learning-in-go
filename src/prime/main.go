package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		primes := sieve()
		for {
			fmt.Println(<-primes)
		}
	}()
	time.Sleep(10 * time.Second)
}

func generate() chan int {
	out := make(chan int)
	go func() {
		for i := 2; ; i++ {
			out <- i
		}
	}()
	return out
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}
