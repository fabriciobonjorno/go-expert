package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
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

func main() {
	fabricio := Cliente{
		Nome:  "fabricio",
		Idade: 18,
		Ativo: true,
	}
	fabricio.Cidade = "Cuiabá"
	fabricio.Desativar()

	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s", fabricio.Nome, fabricio.Idade, fabricio.Ativo, fabricio.Cidade)

}
