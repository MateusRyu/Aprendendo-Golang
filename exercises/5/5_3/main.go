package main

import "fmt"

type veículo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veículo
	traçãoNasQuatro bool
}

type sedan struct {
	veículo
	modeloLuxo bool
}

func main() {
	caminhonete1 := caminhonete{
		veículo: veículo{
			portas: 4,
			cor:    "preta",
		},
		traçãoNasQuatro: true,
	}

	sedan1 := sedan{veículo{portas: 4, cor: "branca"}, true}

	fmt.Println("Caminhonete:")
	fmt.Println("Portas:", caminhonete1.portas)
	fmt.Println("Cor:", caminhonete1.cor)
	fmt.Println("Tração nas quatro rodas:", caminhonete1.traçãoNasQuatro)

	fmt.Println("\nSedan:")
	fmt.Println("Portas:", sedan1.portas)
	fmt.Println("Cor:", sedan1.cor)
	fmt.Println("Modelo de luxo:", sedan1.modeloLuxo)
}
