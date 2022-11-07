package main

// func SomaInteiro(m map[string]int) int {
// 	var soma int
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func SomaFloat(m map[string]float64) float64 {
// 	var soma float64
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

type MyNumber int

// Constraint
type Number interface {
	~int | ~float64 // o ~ ele considera o outro type ou seja ele executa o MyNumber
}

// Generics
// func Soma[T int | float64](m map[string]T) T { modo direto
func Soma[T Number](m map[string]T) T { //usando constraint
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// func Compara[T Number](a T, b T) bool { sem constraint
func Compara[T comparable](a T, b T) bool { // usando constraint
	if a == b {
		return true
	}
	return false

}

func main() {
	m := map[string]int{"Fabricio": 1000, "João": 2000, "Mario": 3000}
	m2 := map[string]float64{"Fabricio": 1000.00, "João": 2000.55, "Mario": 3000.33}

	m3 := map[string]MyNumber{"Fabricio": 1000, "João": 2000, "Mario": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
