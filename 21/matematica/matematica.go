package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type SomaTeste struct {
	Tipo string
}
