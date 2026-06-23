package main

import "fmt"

func main() {
	ano := 1999
	for {
		fmt.Println(ano)
		ano++
		if ano > 2026 {
			break
		}
	}
}