package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Cidade    string
}

func (p *Pessoa) ToString() string {
	return fmt.Sprintf("Nome: %s, Sobrenome: %s, Idade: %d, Cidade: %s", p.Nome, p.Sobrenome, p.Idade, p.Cidade)
}

func main() {

	var pessoa Pessoa
	pessoa.Nome = "Senhor"
	pessoa.Sobrenome = "Teste"
	pessoa.Idade = 30
	pessoa.Cidade = "Sao Paulo"

	fmt.Println(pessoa)
	fmt.Println(pessoa.ToString())

}
