package main

import "fmt"

type innerS struct {
	a, b int
}

type innerS2 struct {
	a, b string
}

type outerS struct {
	innerS
	innerS2
	c int
	a int
}

func main() {
	o := outerS{}
	o.a = 1
	o.innerS.b = 2
	o.c = 3
	o.innerS.a = 4
	fmt.Println(o)
}
