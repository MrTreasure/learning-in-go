package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand := func() interface{} {
		num := rand.Intn(50000000)
		fmt.Println(num)
		return num
	}
	done := make(chan interface{})

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))

	fmt.Println("Primes:")

	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took : %v", time.Since(start))

}

func take(done chan interface{}, inStream chan int, num int) chan int {
	takeStream := make(chan int)

	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case v := <-inStream:
				takeStream <- v
			}
		}
	}()

	return takeStream
}

func repeatFn(done chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func toInt(done chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

func primeFinder(done chan interface{}, inStream <-chan int) chan int {
	primeStream := make(chan int)
	go func() {
		defer close(primeStream)
		for {
			select {
			case <-done:
				return
			case v := <-inStream:
				for i := 2; i < v; i++ {
					if v%i == 0 {
						primeStream <- v
					}
				}
			}
		}
	}()
	return primeStream
}
