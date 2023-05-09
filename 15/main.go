package main

// Memória -> Endereço -> Valor
// A variavél -> ponteiro que tem um endereço na memória  -> Valor
func main() {
	a := 10
	var ponteiro *int = &a
	println(ponteiro)
	*ponteiro = 20
	b := &a
	*b = 30
	println(*b) // o * é usado para mostrar o valor na memória
	println(a)
}
