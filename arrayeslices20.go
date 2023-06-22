package main

import "fmt"

func main() { // Lendo o array de inteiros
	var n int
	fmt.Print("Informe o tamanho do array: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Informe o elemento %d: ", i)
		fmt.Scan(&arr[i])
	}

	// Verificando se o array está ordenado em ordem crescente
	ordenado := true
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			ordenado = false
			break
		}
	}

	// Imprimindo o resultado
	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}

}
