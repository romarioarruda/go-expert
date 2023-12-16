package main

import "fmt"

type Number interface{ int | float64 }

func main() {
	var mapInteiros = map[string]int{"Romario": 10000, "Tati": 9000}

	var inteirosSomados = Soma(mapInteiros)
	fmt.Printf("Resultado da soma: %v\n", inteirosSomados)

	var mapFloats = map[string]float64{"Romario": 10000.5, "Tati": 9000.9}

	var floatSomados = Soma(mapFloats)
	fmt.Printf("Resultado da soma: %v\n", floatSomados)
}

func Soma[T Number](mapa map[string]T) T {
	var total T

	for _, valor := range mapa {
		total += valor
	}

	return total
}
