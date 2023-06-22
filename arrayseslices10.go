package main

import (
	"fmt"
	"sort"
)

func main() {
	// Criando o slice com valores aleatórios
	slice := []int{5, 3, 8, 1, 9, 2, 7, 4, 6, 0}

	// Ordenando o slice
	sort.Ints(slice)

	// Imprimindo o valor mínimo e o valor máximo
	fmt.Println("Valor mínimo:", slice[0])
	fmt.Println("Valor máximo:", slice[len(slice)-1])
}
