package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"` // adiciona tags no json
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta) // usando o marshal ele guarda o valor
	if err != nil {
		println(err)
	}
	println(string(res))
	json.NewEncoder(os.Stdout).Encode(conta) // faz o processo de serialização entregando o resultado

	jsonPuro := []byte(`{"n": 2, "s": 200}`)

	// processo inverso Json para Struct
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX) // vai na memoria e muda o valor usando o endereço
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)

}
