package main

import "fmt"

type S struct {
	a, b int
}

func main() {
	s := new(S)
	s.a = 1
	s.b = 2
	fmt.Println(s)

	s2 := new(S)
	(*s2).a = 3
	(*s2).b = 4
	fmt.Println(s2)
}