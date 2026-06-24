package main

import (
	"fmt"
)

func main() {
	executaCallback(callback, "Olá, mundo!")
}

func callback(s string) {
	fmt.Println(s)
}

func executaCallback(f func(s string), texto string) {
	f(texto)
}
