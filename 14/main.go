package main

import "fmt"

type Adress struct {
	Logra 	string
	Number 	int
	City 		string
	State		string
}

type Person interface {
	Desactive()
}

type Company struct {
	Nome string
}

func (c Company) Desactive()  {
	
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

func Desactivation(person Person) {
	person.Desactive()
}

func main() {
	brando := Client{
		Name: "Brando",
		Age: 30,
		Active: true,
	}

	myCompany := Company{}

	Desactivation(myCompany)
	fmt.Printf("Name: %s, Age: %d, Active: %t\n", brando.Name, brando.Age, brando.Active)
}