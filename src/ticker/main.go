package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go tickerFn(ctx)
	time.Sleep(time.Second * 11)
	cancel()
}

func tickerFn(ctx context.Context) {
	ticker := time.Tick(time.Second * 1)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker:
			fmt.Println("send from tick...")
		}
	}
}
