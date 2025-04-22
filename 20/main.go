package main

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Brando": 1000, "Luca": 2000, "Gustavo": 3000}
	m2 := map[string]float64{"Brando": 1000.65, "Luca": 2000.32, "Gustavo": 3000.98}
	println(Soma(m))
	println(Soma(m2))
}
