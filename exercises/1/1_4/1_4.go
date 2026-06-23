package exercises

import "fmt"

type age int

var x age

func main() {
	fmt.Printf("Value: %v\nType: %T\n", x, x)
	x = 42
	fmt.Printf("Value: %v\nType: %T\n", x, x)
}
