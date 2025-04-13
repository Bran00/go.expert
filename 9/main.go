package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(12, 15, 23, 54, 56, 32, 87))
}

func sum(numeros ...int) (int) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
