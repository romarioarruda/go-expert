package main

import "fmt"

type Nome string

var (
	a string
	b Nome
)

func main() {
	a = "hello world"
	b = "Romario"
	fmt.Printf("O tipo de a é %T\n", a)
	fmt.Printf("O tipo de b é %T\n", b)
}
