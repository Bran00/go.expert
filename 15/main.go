package main

func main() {

	first := 20
	println(first)
  pointer := &first
	println(pointer)
	var b = pointer
	println(*b)
	*b = 30
	println(*b)
	println(first)
}

//#F023