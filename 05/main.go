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
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 5

	fmt.Println(meuArray)
	fmt.Println(meuArray[1])
	fmt.Println(len(meuArray) - 1)
	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor de %d é %d \n", i, v)
	}
}
