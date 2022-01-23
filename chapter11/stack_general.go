package main

import (
	"errors"
	"fmt"
)

type Any interface{}

type Stack interface {
	Len() int
	IsEmpty() bool
	Push(x Any)
	Pop() (Any, error)
}

type GStack struct {
	xs []Any
}

func (s *GStack) Len() int {
  return len(s.xs)
}

func (s *GStack) IsEmpty() bool {
  return s.Len() == 0
}

func (s *GStack) Push(x Any) {
  s.xs = append(s.xs, x)
}

func (s *GStack) Pop() (res Any, e error) {
	if s.IsEmpty() {
		e = errors.New("No elements")
	}
	res = s.xs[s.Len() - 1]
	s.xs = s.xs[:s.Len() - 1]
	return
}


func main() {
	stack := GStack{}

  fmt.Println(stack.IsEmpty())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack)

	v, _ := stack.Pop()
  fmt.Println(v)
  fmt.Println(stack)
}