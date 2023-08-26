package main

import "fmt"

type Number interface {
	int | float64
}

func sum[T Number](a T, b T) T {
	return a + b
}

//comparable values must be the same type
func compare[T comparable](a, b T) bool {
	return a == b
}

func main() {
	resultInt := sum(10, 10)
	fmt.Printf("resultInt: %v, type: %T\n", resultInt, resultInt)

	resultFloat := sum(10.444, 8.7)
	fmt.Printf("resultFloat: %v, type: %T\n", resultFloat, resultFloat)

	compareResult := compare(10, 10)
	fmt.Printf("Compare result: %v\n", compareResult)
}
