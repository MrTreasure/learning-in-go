package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.google.com",
	"http://golang.org",
	"http://bolg.golang.org",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err.Error())
			continue
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
