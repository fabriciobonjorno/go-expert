package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Pessoa interface {
	Desativar()
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(p Pessoa) {
	p.Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func main() {
	fabricio := Cliente{
		Nome:  "fabricio",
		Idade: 18,
		Ativo: true,
	}
	fabricio.Cidade = "Cuiabá"

	minhaEmpresa := Empresa{
		Nome: "Bonjorno",
	}
	// fabricio.Desativar()
	Desativacao(fabricio)
	Desativacao(minhaEmpresa)

	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s", fabricio.Nome, fabricio.Idade, fabricio.Ativo, fabricio.Cidade)

}
