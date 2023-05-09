package main

import (
	"fmt"

	"go-expert/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Andar())
	fmt.Println("O resultado: %v", s)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())

}
