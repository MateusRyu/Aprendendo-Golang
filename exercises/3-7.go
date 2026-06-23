package main

import "fmt"

func main() {
	if x:=100; x<100 {
		fmt.Printf("%d é menor que 100\n", x)
	} else if x>100 {
		fmt.Printf("%d é maior que 100\n", x)
	} else {
	 	fmt.Printf("%d é igual que 100\n", x)
	}
}