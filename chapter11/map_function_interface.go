package main

import "fmt"

type obj interface{}

func mf(o obj) obj {
	switch o.(type) {
	case int:
		fmt.Println("It is int")
		return o.(int) * 2
	case string:
		fmt.Println("It is string")
		return o.(string) + o.(string)
	}
	return o
}

func mapFunc(mf func(o obj) obj, os []obj) (res []obj) {
	res = make([]obj, len(os))
	for i, v := range os {
		res[i] = mf(v)
	}
	return
}

func main() {
	data := []obj{10, 20, "abc", "hello world"}
	res := mapFunc(mf, data)
	fmt.Println(res)
}
