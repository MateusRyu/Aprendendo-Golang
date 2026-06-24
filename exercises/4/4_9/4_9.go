package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"carvalho_joão": {"Futebol", "Leitura"},
		"silva_maria":   {"Balé", "Pintura"},
		"santos_pedro":  {"Valorant", "Ciclismo"},
	}

	for nome, hobbies := range pessoas {
		fmt.Println("Nome:", nome)
		for i, hobby := range hobbies {
			fmt.Printf("Hobby %d: %s\n", i, hobby)
		}
		fmt.Println()
	}

	// fim do exercício anterior

	fmt.Println("4-9:\n")
	pessoas["braga_ana"] = []string{"Cozinhar", "Fofocar"}

	for nome, hobbies := range pessoas {
		fmt.Println("Nome:", nome)
		for i, hobby := range hobbies {
			fmt.Printf("Hobby %d: %s\n", i, hobby)
		}
		fmt.Println()
	}
}
