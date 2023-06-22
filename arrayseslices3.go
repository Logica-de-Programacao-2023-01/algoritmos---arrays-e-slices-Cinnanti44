package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("Informe o tamanho do slice: ")
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Println("Informe o valor do elemento: ", i)
		fmt.Scan(&slice[i])
	}

	fmt.Println("Slice:", slice)

	var soma int
	for _, valor := range slice {
		soma += valor
	}

	fmt.Println("A soma dos elementos do slice é :", soma)

}
