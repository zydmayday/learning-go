package main

import "fmt"

type S1 struct{}
type S2 struct{}

type I interface{
  M()
}

func (s S1) M() {}
func (s S2) M() {}

func main() {
	s1 := S1{}
	s2 := S2{}

	for _, s := range []I{s1, s2} {
		if _, ok := s.(S1); ok {
			fmt.Println("S1")
		} else {
			fmt.Println("Other type")
		}
	}
}