package main

import "fmt"

// const a = "Hello world!"
type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Hello world!"
	e float64 = 1.2
	f ID      = 1 // Cria seu próprio tipo de dado
)

func main() {
	// b = true
	a := "My name is"
	a = "Your name is"
	fmt.Printf("O tipo de a é %T e valor é %v\n", a, a)
	fmt.Printf("O tipo de b é %T e valor é %v\n", b, b)
	fmt.Printf("O tipo de c é %T e valor é %v\n", c, c)
	fmt.Printf("O tipo de d é %T e valor é %v\n", d, d)
	fmt.Printf("O tipo de e é %T e valor é %v\n", e, e)
	fmt.Printf("O tipo de f é %T e valor é %v\n", f, f)

}
