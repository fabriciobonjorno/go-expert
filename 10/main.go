package main

import (
	"fmt"
)

// função variadicas passa um range de valores
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// closures função dentro de função
func main() {
	total := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9) * 2
	}()

	fmt.Println(total)
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
