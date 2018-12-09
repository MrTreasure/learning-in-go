package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)

	for i := 0; i < 100000; i++ {
		intChan <- i
	}

	for i := 0; i < 12; i++ {
		// chan 是线程安全的管道，不会出现重复取数
		go calc(intChan, resultChan)
	}

	time.Sleep(time.Second * 10)
}

func calc(taskChan chan int, resChan chan int) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resChan <- v
		}
	}
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
