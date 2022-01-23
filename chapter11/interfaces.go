package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Shape struct {
	
}

// 接口不需要显示的引用
type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

type Rectangle struct {
	l, w float32
}

func (r *Rectangle) Area() float32 {
	return r.l * r.w
}

type Circle struct {
	r float32
}

func (c *Circle) Area() float32 {
	return 2 * c.r * math.Pi
}

func printArea(s Shaper) {
	fmt.Println(s.Area())
}

func main() {
	s := new(Square)
	s.side = 10

	// Square依然可以赋值给Shaper
	var si Shaper = s

	fmt.Println(si.Area())

	// 多态，同一种类型在不同的实例上表现出不同的行为
	printArea(s)
	printArea(si)

	r := &Rectangle{10, 20}
	shapes := []Shaper{s, r, &Circle{10}}
	for _, s := range shapes {
		printArea(s)
	}
}
