package main

import "fmt"

func main() {
	// Criando o array de inteiros com 5 elementos
	array := [5]int{3, 6, 7, 9, 12}

	// Criando um novo slice que contém apenas os elementos múltiplos de 3
	var slice []int
	for _, valor := range array {
		if valor%3 == 0 {
			slice = append(slice, valor)
		}
	}

	// Imprimindo o novo slice
	fmt.Println(slice)
}
