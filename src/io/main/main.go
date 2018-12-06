package main

import (
	"bufio"
	"fmt"
	"os"
)

type student struct {
	Name   string
	Age    int
	Scrore float32
}

func main() {
	str := "stu01 18 89.92"
	var stu student
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Scrore)
	fmt.Println(stu)

	reader := bufio.NewReader(os.Stdin)
	str2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed, err:", err)
	}

	fmt.Printf("read str succ, res:%s\n", str2)
}
