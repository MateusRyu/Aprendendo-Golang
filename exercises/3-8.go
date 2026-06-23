package main

import "fmt"

func main() {
	x := 2
	switch {
	case x == 0:
		fmt.Println("zero")
	case x == 1:
		fmt.Println("um")
	case x == 2:
		fmt.Println("dois")
	}
}