package main

import "fmt"

func main() {
	pessoa := [][]string{
		{"João", "Carvalho", "Futebol"},
		{"Maria", "Silva", "Balé"},
		{"Pedro", "Santos", "Valorant"},
	}

	for _, p := range pessoa {
		fmt.Println("Nome:", p[0])
		fmt.Println("Sobrenome:", p[1])
		fmt.Println("Hobby:", p[2])
		fmt.Println()
	}
}
