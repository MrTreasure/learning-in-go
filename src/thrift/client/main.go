package main

import (
	"context"
	"ele"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	HOST = "127.0.0.1"
	PORT = "9090"
)

var (
	startTime int64
	endTime   int64
)

func main() {
	ctx := context.Background()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))

	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error geting transport:", err)
		os.Exit(1)
	}

	client := ele.NewEleThriftClientFactory(useTransport, protocolFactory)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to"+HOST+":"+PORT, "", err)
		os.Exit(1)
	}

	defer transport.Close()
	students := getData()

	for _, v := range students {
		err := client.Put(ctx, &v)
		if err != nil {
			fmt.Printf("Error in put, err: %s\n", err)
		}
		os.Exit(1)
		fmt.Printf("Add student %v\n", v)
		time.Sleep(time.Second * 2)
	}

	time.Sleep(time.Second * 5)
	startTime = currentTimeMillis()
	stu1, err := client.FindStuById(ctx, 1)
	if err != nil {
		fmt.Printf("Error in findStuByID, err: %s\n", err)
	}
	endTime = currentTimeMillis()
	fmt.Printf("本次调用耗时: %d - %d = %d ms, %v\n", endTime, startTime, (endTime - startTime), stu1)

	time.Sleep(time.Second * 5)
	startTime = currentTimeMillis()
	stu2, err := client.FindStuByName(ctx, "一哥")
	if err != nil {
		fmt.Printf("Error in findStuByName, err: %s\n", err)
	}
	endTime = currentTimeMillis()
	fmt.Printf("本次调用耗时: %d - %d = %d ms, %v\n", endTime, startTime, (endTime - startTime), stu2)
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

func getData() []ele.Student {
	stuByte, err := ioutil.ReadFile("stu.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading json data: ", err)
		os.Exit(1)
	}
	stus := make([]ele.Student, 10)
	err = json.Unmarshal(stuByte, &stus)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error unmarshal data: ", err)
		os.Exit(1)
	}

	return stus
}
