package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(5, 10))
	fmt.Println(route("/"))
}

func add(x, y int) int {
	return x + y
}

func route(path string) (int, string) {
	switch path {
	case "/":
		return 200, "Home"
	case "/about":
		return 200, "About"
	default:
		return 404, "Not Found"
	}
}
