package main

import "fmt"

const a = "Hello, world!"

var (
	b bool    = true
	c int     = 10
	d string  = "Hello, world!"
	e float64 = 1.2
)

func main() {
	// var a string = "X"
	// a := "B" // usar : para atribuir somente a primeira vez
	// a = "D"
	fmt.Printf("O tipo de E é %T", e)
}
