package main

import "fmt"

type employee struct {
	salary float32
}

func (this *employee) giveRaise(r float32) {
	this.salary = this.salary * (1 + r)
}

func main() {
	e := employee{100}
	e.giveRaise(0.5)
	fmt.Println(e)
}
