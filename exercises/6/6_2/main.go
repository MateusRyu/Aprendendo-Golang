package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println(add(x...))
	fmt.Println(soma(x))
}

func add(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func soma(x []int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
