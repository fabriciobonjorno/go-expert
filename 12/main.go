package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	fabricio := Cliente{
		Nome:  "fabricio",
		Idade: 18,
		Ativo: true,
	}
	fabricio.Ativo = false
	fabricio.Cidade = "Cuiabá"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s", fabricio.Nome, fabricio.Idade, fabricio.Ativo, fabricio.Cidade)

}
