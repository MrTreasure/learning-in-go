package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int)
	go write(intChan)
	go read(intChan)

	time.Sleep(10 * time.Second)
}

func read(ch chan int) {
	for {
		var b, c int
		b = <-ch
		c = <-ch
		fmt.Println(b, c)
		time.Sleep(time.Second)
	}
}

func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}
