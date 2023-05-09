package main

func soma(a, b *int) int { // Para mudar o endereço na memoria é necessario passar o ponteiro usando * caso contrario será somente uma cópia
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	println(soma(&minhaVar1, &minhaVar2)) // Para interpretar a soma é necessario passar o &
	println(minhaVar1)
	println(minhaVar2)

}
