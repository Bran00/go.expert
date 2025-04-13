package main

import "fmt"

func main() {

	salarios := map[string]int{
		"Pedro": 1000,
		"Maria": 1500,
		"João": 2000,
		"José": 2500,
	}
	fmt.Println(salarios["José"])
	delete(salarios, "José")
	salarios["Brando"] = 6000
	
	sal := make(map[string]int)
	sal["Brando"] = 6000
	sal1 := map[string]int{}
	sal1["Brando"] = 8000

	for nome, salario := range salarios {
		fmt.Printf("O sálario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O sálario é %d\n",  salario)
	}
}