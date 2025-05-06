package main

import (
	"github.com/Bran00/go.expert/37/4/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}

//go work init ./math ./sistema
//go mod tidy -e