package main

// func Soma[T int | float64](m map[string]T) T {
// 	var soma T
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

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

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool { // comparable só funciona caso seja possivel a comparação
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Fabricio": 1000, "Maria": 2000, "Sergio": 3000}
	m2 := map[string]float64{"Fabricio": 1000.35, "Maria": 2000.10, "Sergio": 3000.20}

	m3 := map[string]MyNumber{"Fabricio": 1000, "Maria": 2000, "Sergio": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))

}
