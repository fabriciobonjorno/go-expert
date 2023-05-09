package main

// type Cliente struct {
// 	nome string
// }

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}

}

// func (c Cliente) andou() {
// 	c.nome = "Fabricio Comeli"
// 	fmt.Printf("O client %v andou!\n", c.nome)
// }

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
	// fabricio := Cliente{
	// 	nome: "fabricio",
	// }
	// fabricio.andou()
	// fmt.Printf("O valor da struct nome é %v\n", fabricio.nome)
}
