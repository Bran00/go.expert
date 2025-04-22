package main

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable] (a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Brando": 1000, "Luca": 2000, "Gustavo": 3000}
	m2 := map[string]float64{"Brando": 1000.65, "Luca": 2000.32, "Gustavo": 3000.98}
	println(Soma(m))
	println(Soma(m2))

	println(Compara(10, 10.4))
}
