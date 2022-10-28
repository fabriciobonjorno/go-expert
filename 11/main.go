package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	fabricio := Cliente{
		Nome:  "fabricio",
		Idade: 18,
		Ativo: true,
	}
	fabricio.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", fabricio.Nome, fabricio.Idade, fabricio.Ativo)

}
