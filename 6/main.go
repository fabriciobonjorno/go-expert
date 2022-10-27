package main

import "fmt"

func main() {
	meuSlice := []int{2, 4, 6, 8, 10}

	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)

	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice[:0]), cap(meuSlice[:0]), meuSlice[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice[:4]), cap(meuSlice[:4]), meuSlice[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice[2:]), cap(meuSlice[2:]), meuSlice[2:])

	meuSlice = append(meuSlice, 12)

	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)
	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice[2:]), cap(meuSlice[2:]), meuSlice[2:])

	for i, v := range meuSlice {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}

}
