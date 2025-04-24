package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"uno", "dos", "tres", "cuatro", "cinco"}
	for _, v := range numeros {
		println(v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}
}