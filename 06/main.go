package main

import "fmt"

func main() {
	sli := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d %v\n", len(sli), cap(sli), sli)

	fmt.Printf("len=%d cap=%d %v\n", len(sli[:0]), cap(sli[:0]), sli[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(sli[:4]), cap(sli[:4]), sli[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(sli[2:]), cap(sli[2:]), sli[2:])

	sli = append(sli, 12)

	fmt.Printf("len=%d cap=%d %v\n", len(sli[2:]), cap(sli[2:]), sli[2:])
}
