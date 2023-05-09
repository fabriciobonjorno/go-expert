package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	// Address // adiciona composição
	Endereço Address // adiciona tipo

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

func main() {

	fab := Client{
		Nome:  "Fabricio",
		Idade: 33,
		Ativo: true,
	}
	fab.Desativar()
	// fab.Nome = "Fabri" // muda valor
	// fab.Cidade = "Cuiabá"
	// fab.Address.Cidade = "Cuiabá" // mesmo valor mas menos cleaning
	fab.Endereço.Cidade = "Cuiabá" // adicionado através de tipo
	fmt.Println(fab.Nome)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", fab.Nome, fab.Idade, fab.Ativo)

}
