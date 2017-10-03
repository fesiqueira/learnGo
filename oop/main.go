package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Pessoa struct {
	Nome     string
	Ocupacao string
	Idade    int
	Altura   float64
}

func (p *Pessoa) Falar() {
	random := rand.Intn(5)

	time.Sleep(time.Duration(random) * time.Second)
	fmt.Println("Olá, meu nome é", p.Nome, "e sao", time.Now().Second())
}

type Aluno struct {
	Pessoa
	Semestre int
}

func NewAluno(nome, ocupacao string, idade, semestre int, altura float64) *Aluno {
	return &Aluno{
		Pessoa: Pessoa{
			Nome:     nome,
			Ocupacao: ocupacao,
			Idade:    idade,
			Altura:   altura,
		},
		Semestre: semestre,
	}
}

type Professor struct {
	Pessoa
	Salario float64
}

func main() {
	aluno1 := NewAluno("Felipe", "Estagiario", 21, 6, 1.79)
	aluno2 := NewAluno("Jose", "Estagiario", 21, 6, 1.79)
	aluno3 := NewAluno("Roberto", "Estagiario", 21, 6, 1.79)
	aluno4 := NewAluno("Maria", "Estagiario", 21, 6, 1.79)

	var alunos = []*Aluno{aluno1, aluno2, aluno3, aluno4}

	for _, aluno := range alunos {
		aluno.Falar()
	}

	var input string
	fmt.Scanln(&input)
}
