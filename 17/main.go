package main

// type Cliente struct {
// 	nome string
// }

// func (c Cliente) andou() {
// 	c.nome = "Fabricio Bonjorno"
// 	fmt.Printf("O cliente %s andou!\n", c.nome)
// }
// func main() {
// 	fabricio := Cliente{
// 		nome: "Fabricio",
// 	}
// 	fabricio.andou()
// 	fmt.Printf("O valor da struct com o nome %v\n", fabricio.nome)
// }

// Sem ponteiros

// type Conta struct {
// 	saldo int
// }

// func (c Conta) simular(valor int) int {
// 	c.saldo += valor
// 	println(c.saldo)
// 	return c.saldo
// }
// func main() {
// 	conta := Conta{saldo: 100}
// 	conta.simular(200)
// 	println(conta.saldo)
// }

// Com ponteiros

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}
func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}
