package main

import (
	"dispatcher/balance"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	insts := make([]*balance.Instance, 0, 16)
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	var banlancer balance.Balancer
	var conf = os.Args[1]

	if conf == "random" {
		banlancer = &balance.RandomBalance{}
	} else if conf == "roundrobin" {
		banlancer = &balance.RoundRobinBalance{}
	} else {
		banlancer = &balance.RandomBalance{}
	}

	for {
		inst, err := balancer.DoBalance(insts)
		if err != nil {
			panic(err)
		}
		fmt.Println(inst.GetHost(), inst.GetPort())
		time.Sleep(time.Second * 2)
	}
}
