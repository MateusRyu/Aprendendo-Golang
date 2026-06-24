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
	fulano := pessoa{"Fulano", "da Silva", 30}
	fmt.Println(fulano)
	mudeMe(&fulano)
	fmt.Println(fulano)
}

func mudeMe(p *pessoa) {
	p.idade += 1
}
