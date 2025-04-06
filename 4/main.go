package main

import "fmt"

const a = "Hello"

type ID int

var (
    b bool    = true
    c int     = 42
    d string  = "World"
    e float64 = 3.14
    f ID      = 1
)

func main() {
   fmt.Printf("O tipo de E é %T", f)
}
