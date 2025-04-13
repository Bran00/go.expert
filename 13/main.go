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

func (c Client) Desactive() {
	c.Active = false
	fmt.Printf("Client %s is now inactive\n", c.Name)
}

func main() {
	brando := Client{
		Name: "Brando",
		Age: 30,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", brando.Name, brando.Age, brando.Active)
	
	brando.City = "Brasilia"
	brando.State = "DF"
	brando.Desactive()
	fmt.Print(brando)
}