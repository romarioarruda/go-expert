package main

import (
	"errors"
	"fmt"
)

func main() {
	var notaRomario, _ = pegaNotaPorAluno("romario")
	fmt.Println("Nota romario:", notaRomario)

	var notaMaria, _ = pegaNotaPorAluno("maria")
	fmt.Println("Nota maria:", notaMaria)

	notaJoao, err := pegaNotaPorAluno("joao")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Nota joao:", notaJoao)
}

func pegaNotaPorAluno(aluno string) (int, error) {
	alunos := map[string]int{
		"romario": 10,
		"tati":    9,
		"maria":   8,
	}

	valor, existe := alunos[aluno]
	if !existe {
		return 0, errors.New("aluno não identificado")
	}

	return valor, nil
}
