package main

import (
	"fmt"
)

func main() {
	a := closure(3)
	b := closure(5)

	a()
	a()
	a()
	a()
	fmt.Println("----------------------")
	b()
	b()
}

func closure(n int) func() {
	x := n

	return func() {
		x++
		fmt.Println(x)
	}
}
