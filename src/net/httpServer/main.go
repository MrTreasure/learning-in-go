package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "URL: "+req.URL.Path)
	fmt.Fprintf(w, "%s<br/>", req.UserAgent())
	buffer := strings.Builder{}
	for _, v := range req.Cookies() {
		buffer.WriteString(fmt.Sprintf("%s<br/>", v.String()))
	}
	fmt.Fprintf(w, "%s", buffer.String())
}
