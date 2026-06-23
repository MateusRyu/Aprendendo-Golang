package exercises

import "fmt"

func main() {
	esporteFavorito := "xadrez"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Quadra de futebol")
	case "natação":
		fmt.Println("Piscina")
	case "xadrez":
		fmt.Println("Clube de xadrez")
	}
}
