package main

import "fmt"

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 冒泡排序
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Person struct {
	firstName, lastName string
}

type Persons []Person

func (p Persons) Len() int { return len(p) }
func (p Persons) Less(i, j int) bool {
	if p[i].firstName == p[j].firstName {
		return p[i].lastName < p[j].lastName
	} else {
		return p[i].firstName < p[j].firstName
	}
}
func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data) //conversion to type IntArray
	Sort(a)
	if !IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

func persons() {
	data := []Person{{"zhang", "yidong"}, {"zhang", "erdong"}, {"li", "hui"}}
	a := Persons(data) //conversion to type IntArray
	Sort(a)
	if !IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

func main() {
	ints()
	persons()
}
