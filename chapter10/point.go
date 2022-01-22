package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Polar struct {
	X, Y, Z float64
}

func Abs(p Point) (dis float64) {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}

func Scale(p *Point, s float64) {
	p.X = p.X * s;
	p.Y = p.Y * s;
}

func main() {
	p := Point{2, 3}
	fmt.Println(Abs(p))
	Scale(&p, 2)
	fmt.Println(p)
}
