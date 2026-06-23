package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%d dividido por 4 sobra %d\n", i, i%4)
	}
}