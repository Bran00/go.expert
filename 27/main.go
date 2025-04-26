package main

import (
	"encoding/json"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println("Error:", err)
	}
	println(string(res))

	jsonPuro := []byte(`{"N":2,"S":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println("Error:", err)
	}
	println(contaX.Saldo)
}
