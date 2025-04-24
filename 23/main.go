package main

func main() {
	a := 1
	b := 2

	switch a {
		case 1:
			println("a")
		case 2:
			println("b")
	}

	if (a > b) {
		println("a is greater than b")
	} else {
		println("a is not greater than b")
	}
}