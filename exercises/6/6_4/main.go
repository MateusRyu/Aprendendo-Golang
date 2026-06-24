package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	p1 := pessoa{
		nome:      "João",
		sobrenome: "Silva",
		idade:     30,
	}

	p1.apresentar()
}

func (p pessoa) apresentar() {
	fmt.Printf("Nome: %s %s\nIdade: %d\n", p.nome, p.sobrenome, p.idade)
}
