package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		default:
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}
