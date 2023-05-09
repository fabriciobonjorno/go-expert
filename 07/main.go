package main

import "fmt"

func main() {
	salarios := map[string]int{"Fabricio": 1000, "João": 2000, "Mario": 3000}
	fmt.Println(salarios["Fabricio"])
	delete(salarios, "João")
	salarios["Maria"] = 2000
	fmt.Println(salarios["Maria"])

	sal := make(map[string]int) // inicia vazio
	sal["Marcos"] = 12000
	sal1 := map[string]int{} // inicia vazio
	sal1["Mar"] = 20000

	for name, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", name, salario)
	}

	// ignora valor utilizando _ blank identify
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
