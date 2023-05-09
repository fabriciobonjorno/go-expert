package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	fab := Client{
		Nome:  "Fabricio",
		Idade: 33,
		Ativo: true,
	}

	fab.Nome = "Fabri" // muda valor
	fmt.Println(fab.Nome)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", fab.Nome, fab.Idade, fab.Ativo)

}
