package main

import "fmt"

func main() {
	a := 0 == 10
	b := 0 != 10
	c := 0 <= 10
	d := 0 < 10
	e := 0 >= 10
	f := 0 > 10

	fmt.Println(a, b, c, d, e, f)
}