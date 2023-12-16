package main

import "fmt"

type Aluno struct {
	Name  string
	Ativo bool
}

func (aluno *Aluno) Desativar() {
	aluno.Ativo = false
	fmt.Printf("Aluno %v foi desativado\n", aluno)
}

func main() {
	var aluno = Aluno{
		Name:  "Romario",
		Ativo: true,
	}

	fmt.Printf("Aluno: %v\n", aluno)

	aluno.Desativar()

	fmt.Printf("Aluno: %v\n", aluno)

	valor1 := 10

	var valor2 *int = &valor1

	*valor2 = 20

	fmt.Println("valor2", valor2)
	fmt.Println("valor1", &valor1)
}
