package main

import "fmt"

func main() {
	esporteFavoito := "xadrez"

	switch esporteFavoito {
	case "futebol":
		fmt.Println("Quadra de futebol")
	case "natação":
		fmt.Println("Piscina")
	case "xadrez":
		fmt.Println("Clube de xadrez")
	}
}