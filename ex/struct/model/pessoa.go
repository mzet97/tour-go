package model

import (
	"fmt"
	"time"
)

type Pessoa struct {
	Nome             string
	DataDeNascimento time.Time
	Endereco         Endereco
	Idade            int
}

// method pessoa.CalculaIdade
func (p *Pessoa) CalculaIdade() {
	anoDeNacimiento := p.DataDeNascimento.Year()
	year := time.Now().Year()
	p.Idade = year - anoDeNacimiento
	fmt.Println("Idade calculada:", p.Idade)
}

// function model.IdadeAtual
func IdadeAtual(p Pessoa) int {
	anoDeNacimiento := p.DataDeNascimento.Year()
	year := time.Now().Year()
	return year - anoDeNacimiento
}
