package main

import (
	"fmt"

	"github.com/Bran00/go.expert/matematica"
	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.SomaTeste{Tipo: "Soma"}
	fmt.Println("Resultado: ", s)
	fmt.Println((matematica.A))
	fmt.Println(carro)
	fmt.Println(uuid.New())
}
