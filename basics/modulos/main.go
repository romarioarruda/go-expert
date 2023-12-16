package main

import (
	"fmt"

	"go-expert/basics/modulos/calculos"
)

func main() {
	var soma = calculos.Soma(10, 20)
	fmt.Printf("O resultado da soma é: %v\n", soma)

	var multiplicacao = calculos.Multiplicar(10, 20)
	fmt.Printf("O resultado da multiplicacao é: %v\n", multiplicacao)
}
