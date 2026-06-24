package main

import "fmt"

type pessoa struct {
	nome              string
	sobrenome         string
	sabores_favoritos []string
}

func main() {
	fulano := pessoa{
		nome:              "Fulano",
		sobrenome:         "da Silva",
		sabores_favoritos: []string{"chocolate", "morango", "baunilha"},
	}

	ciclano := pessoa{"Ciclano", "de Souza", []string{"flocos", "doce de leite", "coco"}}

	fmt.Println(fulano.nome, fulano.sobrenome)
	fmt.Println("Sabores favoritos:")
	for _, sabor := range fulano.sabores_favoritos {
		fmt.Println("-", sabor)
	}

	fmt.Println("\n", ciclano.nome, ciclano.sobrenome)
	fmt.Println("Sabores favoritos:")
	for _, sabor := range ciclano.sabores_favoritos {
		fmt.Println("-", sabor)
	}

	// Fim do exercício anterior

	fmt.Println("\n-------- MAP --------")
	pessoas := map[string]pessoa{
		"da Silva": fulano,
		"de Souza": ciclano,
	}

	for _, p := range pessoas {
		fmt.Println(p.nome, p.sobrenome)
		fmt.Println("Sabores favoritos:")
		for _, sabor := range p.sabores_favoritos {
			fmt.Println("-", sabor)
		}
		fmt.Println()
	}
}
