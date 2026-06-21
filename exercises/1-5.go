package main

import "fmt"

type age int
var x age
var y int

func main() {
	fmt.Printf("X\ns value: %v\nType: %T\n", x, x)
	x = 42
	fmt.Printf("Value: %v\nType: %T\n", x, x)

	y = int(x)
	fmt.Printf("\nY\nValue: %v\nType: %T\n", y, y)
}