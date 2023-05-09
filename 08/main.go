package main

import (
	"errors"
	"fmt"
)

//	func sum(a int, b int) int {
//		return a + b
//	}
//
// se a variavel for do mesmo tipo pode passar somente 1x o tipo
func sum(a, b int) int {
	return a + b
}

// Adiciona mais de um tipo
func sum1(c, d int) (int, bool) {
	if c+d >= 50 {
		return c + d, true
	}
	return c + d, false
}

func sum2(e, f int) (int, error) {
	if e+f >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return e + f, nil
}

func main() {
	fmt.Println(sum(2, 4))
	fmt.Println(sum1(50, 4))
	valor, err := sum2(50, 20)
	// condicional de erro
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}
