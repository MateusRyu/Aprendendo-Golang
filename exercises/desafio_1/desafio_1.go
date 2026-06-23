package exercises

import (
	"fmt"
)

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("(%d) %v\n", x, rune(x))
	}
}
