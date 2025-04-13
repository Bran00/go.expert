package main

import "fmt"

type Client struct {
	Name string
	Age int
	Active bool
}

func main() {
	brando := Client{
		Name: "Brando",
		Age: 30,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", brando.Name, brando.Age, brando.Active)

	brando.Active = false
	fmt.Printf("Name: %s, Age: %d, Active: %t\n", brando.Name, brando.Age, brando.Active)
}