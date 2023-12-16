package main

import "fmt"

type Aluno struct {
	Name  string
	Nota  float64
	Ativo bool
}

func (aluno *Aluno) Desativar() {
	aluno.Ativo = false
	fmt.Printf("Aluno %v foi desativado\n", aluno)
}

func main() {
	var aluno = Aluno{
		Name:  "Romario",
		Nota:  10.0,
		Ativo: true,
	}

	fmt.Printf("Aluno: %v\n", aluno)

	aluno.Desativar()

	fmt.Printf("Aluno: %v\n", aluno)
}
