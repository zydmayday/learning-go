package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b *B) change() { b.thing = 10 }

func (b B) write() string { return fmt.Sprint(b) }

func main() {
	// go会自动的侦测是值还是指针，然后帮助进行转换
	var b1 B // b1 是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2 是指针
	b2.change()
	fmt.Println(b2.write())
}
