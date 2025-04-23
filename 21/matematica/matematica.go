package matematica

func Soma[T int | float64] (a, b T) {
	return a + b
}