package main

import (
	"fmt"
)

func main() {
	x := func() {
		fmt.Println("Função anônima")
	}

	x()
}
