package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	defer request.Body.Close()
	defer fmt.Println("Request completed")

	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
}
