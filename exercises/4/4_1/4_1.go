package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	for i, v := range array {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	fmt.Printf("%T", array)
}
