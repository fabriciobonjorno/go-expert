package main

import (
	"fmt"
)

func main() {
	salarios := map[string]int{"Fabricio": 1000, "Fabio": 2000, "João": 3000}
	fmt.Println(salarios["Fabricio"])
	delete(salarios, "Fabricio")
	salarios["Fab"] = 5000
	fmt.Println(salarios["Fab"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Fab"] = 5000
	// sal["Fab"] = 5000

	for nome, salario := range salarios {
		fmt.Printf("O sálario do %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O sálario é %d\n", salario)
	}

}
