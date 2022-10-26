package main

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
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)

}
