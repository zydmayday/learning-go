package main

import (
	"bufio"
	"chapter12/stack"
	"fmt"
	"os"
	"strconv"
)

func wrap(s stack.Stack, op func(float32, float32) float32) {
	first, _ := s.Pop()
	second, _ := s.Pop()
	res := op(second, first)
	s.Push(res)
}

func minus(a float32, b float32) float32 {
	return a - b
}

func plus(a float32, b float32) float32 {
	return a + b
}

func multi(a float32, b float32) float32 {
	return a * b
}

func devide(a float32, b float32) float32 {
	return a / b
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	s := stack.GStack{}
	for {
		fmt.Println("input number or operator:")
		input, err := inputReader.ReadString('\n')
		input = input[0 : len(input)-2]
		if input == "q" {
			break
		}
		if err != nil {
			fmt.Printf("Error occurs reading, %v\n", err)
			return
		}
		switch input {
		case "+":
			wrap(&s, plus)
		case "-":
			wrap(&s, minus)
		case "*":
			wrap(&s, multi)
		case "/":
			wrap(&s, devide)
		default:
			{
				res, _ := strconv.ParseFloat(input, 32)
				s.Push(float32(res))
			}
		}
		fmt.Println(s)
	}
}
