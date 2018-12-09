package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)

	for v := range t.C {
		fmt.Println("incoming:", v)
	}
}
