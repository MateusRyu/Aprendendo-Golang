package main

import "fmt"

const (
	_ = iota + 2026
	a
	b
	c
	d
)

func main()  {
	fmt.Println(a, b, c, d)
}