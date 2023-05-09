package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	// Address // adiciona composição
	Endereço Address // adiciona tipo

}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Address struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// adiciona metedo com function
func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	fab := Client{
		Nome:  "Fabricio",
		Idade: 33,
		Ativo: true,
	}

	minhaEmpresa := Empresa{Nome: "Rocket"}
	Desativacao(minhaEmpresa)
	Desativacao(fab)

}
