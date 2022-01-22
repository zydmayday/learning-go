package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

type Car2 struct {
	Engine
}

type SportsCar struct {
	Car
	Car2
}

func (this Car2) Start() {
	fmt.Println("Car2 Start")
}

func (this Car) Start() {
	fmt.Println("Car Start")
}

func (this Car) Stop() {
	fmt.Println("Stop")
}

func main() {
	c := SportsCar{}
	c.Car.Start()
	c.Car2.Start()
	c.Stop()
}
