package main

import "fmt"

func receive(nome string, hello chan<- string) {
	hello <- nome
}

func read(data <-chan string) {
	fmt.Println(<-data)
}

// Thread 1
func main() {
	hello := make(chan string)
	go receive("Hello World", hello)
	read(hello)
}