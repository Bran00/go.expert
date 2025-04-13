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
    var myArray [3]int
    myArray[0] = 1
    myArray[1] = 2
    myArray[2] = 10

    for i, v := range myArray {
        fmt.Printf("this value is %d and the index is %d\n", v, i)
    }
}
