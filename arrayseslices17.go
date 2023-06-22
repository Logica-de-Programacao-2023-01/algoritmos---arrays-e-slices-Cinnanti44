package main

import "fmt"

func main() {
	// Criando o Array de inteiros com 10 elementos
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Calculando a soma dos elementos nas posições pares do Array
	var sum int
	for i, value := range arr {
		if i%2 == 0 {
			sum += value
		}
	}

	// Imprimindo a soma dos elementos nas posições pares do Array
	fmt.Println("A soma dos elementos nas posições pares do Array é:", sum)
}
