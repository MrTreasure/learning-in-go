package main

import (
	"context"
	"ele"
	"fmt"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	NetworkAddr = "127.0.0.1:9090"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	handler := &eleThrift{}
	processor := ele.NewEleThriftProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)

	fmt.Println("thrift server in", NetworkAddr)

	server.Serve()
}

type eleThrift struct {
	list []*ele.Student
}

func (p *eleThrift) FindStuById(ctx context.Context, id int32) (stu *ele.Student, err error) {
	for _, v := range p.list {
		if v.ID == id {
			stu = v
			break
		}
	}
	return stu, nil
}

func (p *eleThrift) FindStuByName(ctx context.Context, name string) (stu *ele.Student, err error) {
	for _, v := range p.list {
		if v.Name == name {
			stu = v
			break
		}
	}
	return stu, nil
}

func (p *eleThrift) FindStuByDuty(ctx context.Context, duty string) ([]*ele.Student, error) {
	list := make([]*ele.Student, 0, 10)
	for _, v := range p.list {
		if v.Duty == duty {
			list = append(list, v)
		}
	}
	return list, nil
}

func (p *eleThrift) Put(ctx context.Context, stu *ele.Student) error {
	p.list = append(p.list, stu)
	return nil
}
