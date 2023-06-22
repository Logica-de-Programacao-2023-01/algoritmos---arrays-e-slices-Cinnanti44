package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Criando um novo Slice que contenha apenas os elementos pares do Array original
	var slice []int
	for _, value := range arr {
		if value%2 == 0 {
			slice = append(slice, value)
		}
	}

	// Imprimindo o novo Slice
	fmt.Println(slice)
}
