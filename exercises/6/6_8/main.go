package main

import (
	"fmt"
)

func main() {
	hello := pegaHello()
	hello()
}

func pegaHello() func() {
	return func() {
		fmt.Println("Hello world!")
	}
}
