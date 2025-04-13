package main

import "fmt"

type Adress struct {
	Logra 	string
	Number 	int
	City 		string
	State		string
}

type Client struct {
	Name 		string
	Age 		int
	Active 	bool
	Adress
}

func main() {
	brando := Client{
		Name: "Brando",
		Age: 30,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", brando.Name, brando.Age, brando.Active)

	brando.Active = false
	brando.City = "Brasilia"
	brando.State = "DF"
	fmt.Print(brando)
}