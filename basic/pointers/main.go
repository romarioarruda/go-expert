package main

func main() {
	// Memory => Address => Value
	a := 10

	var pointer *int = &a
	println(*pointer)

	*pointer = 20
	println(*pointer)

	b := a
	println(b)

	valueA := 10
	valueB := 20
	println("Value from A before:", valueA)
	sumResult := sum(&valueA, valueB)
	println("sumResult:", sumResult)

	println("Value from A now:", valueA)
}

func sum(a *int, b int) int {
	*a = 30

	return *a + b
}