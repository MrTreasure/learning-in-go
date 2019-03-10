package main

import (
	"exam/Car"
	"fmt"
	"reflect"
)

type BMW struct {
}

type Hello func() string

const zero = 0.0

func main() {
	car := &BMW{}
	addCar(car)

	fmt.Println(reflect.TypeOf(zero))
	fmt.Println(reflect.ValueOf(zero))
}

func addCar(car Car.Car) {
	car.Start()
	car.Di()
	car.Stop()
}

func (car *BMW) Stop() {
	fmt.Println("BMW was stpped")
}

func (car *BMW) Di() {
	fmt.Println("Di Di Di")
}

func (car *BMW) Start() {
	fmt.Println("BMW is running")
}

func (car *BMW) Error() string {
	return "something crash"
}
