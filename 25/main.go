package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Hello, world!") Escreve somente strings
	// if err != nil {
	// 	panic(err)
	// }

	tamanho, err := f.Write([]byte("Hello, world! Welcome")) // adicona os bytes, slash de dados
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	f.Close()

	//// Leitura

	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Leitura de pouco em pouco abrindo o arquivo

	file1, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file1)
	buffer := make([]byte, 15)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// remove
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
