package main

import "fmt"

func main() {
	x := 10

	fmt.Printf("X: %d\t%b\t%#x\n", x, x, x)

	y := x << 1
	fmt.Printf("Y: %d\t%b\t%#x\n", y, y, y)
}