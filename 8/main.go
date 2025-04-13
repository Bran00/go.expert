package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(1, 4)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Valor:", valor)
}

func sum(a, b int) (int, error) {
	if a + b >= 50 {
		return 0, errors.New("sum is greater than or equal to 50")
	}
	return a + b, nil
}
