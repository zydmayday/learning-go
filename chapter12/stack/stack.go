package stack

import (
	"errors"
)

type Stack interface {
	Len() int
	IsEmpty() bool
	Push(x float32)
	Pop() (float32, error)
}

type GStack struct {
	xs []float32
}

func (s *GStack) Len() int {
	return len(s.xs)
}

func (s *GStack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *GStack) Push(x float32) {
	s.xs = append(s.xs, x)
}

func (s *GStack) Pop() (res float32, e error) {
	if s.IsEmpty() {
		e = errors.New("No elements")
	}
	res = s.xs[s.Len()-1]
	s.xs = s.xs[:s.Len()-1]
	return
}
