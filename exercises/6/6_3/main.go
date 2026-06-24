package main

import (
	"fmt"
)

func main() {
	fmt.Println("Abre arquivo!")
	defer fmt.Println("Fecha arquivo!")
	fmt.Println("Lê arquivo!")
	fmt.Println("Escreve no arquivo!")
}
