package main

import (
	"fmt"
)

func main() {
	x := struct {
		pessoas map[int]string
		ids     []int
	}{
		pessoas: map[int]string{
			10:  "João",
			22:  "Maria",
			300: "José",
		},
		ids: []int{300, 22, 10},
	}

	for _, id := range x.ids {
		fmt.Println(x.pessoas[id])
	}
}
