package main

import "github.com/Bran00/go.expert/37/3/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}

//go mod edit -replace github.com/Bran00/go.expert/37/3/math=../math

/* go mod tidy
go: found github.com/Bran00/go.expert/37/3/math in github.com/Bran00/go.expert/37/3/math v0.0.0-00010101000000-000000000000 */