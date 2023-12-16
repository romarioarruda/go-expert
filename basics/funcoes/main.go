package main

import "fmt"

type Aluno struct {
	Name string
	Nota float64
}

func main() {
	var alunos = []Aluno{
		{Name: "Romario", Nota: 10.0},
		{Name: "Tati", Nota: 9.0},
		{Name: "Maria", Nota: 8.0},
		{Name: "Joao", Nota: 6.0},
		{Name: "Jose", Nota: 5.0},
	}
	fmt.Println("Todos os alunos", alunos)

	var melhoresAlunos = filter(alunos, func(aluno Aluno) bool {
		return aluno.Nota > 6
	})

	fmt.Println("Melhores alunos", melhoresAlunos)

	var pioresAlunos = filter(alunos, func(aluno Aluno) bool {
		return aluno.Nota < 7
	})

	fmt.Println("Piores alunos", pioresAlunos)
}

func filter(slice []Aluno, function func(input Aluno) bool) []Aluno {
	var resultado []Aluno

	for _, valor := range slice {
		if function(valor) {
			resultado = append(resultado, valor)
		}
	}

	return resultado
}
