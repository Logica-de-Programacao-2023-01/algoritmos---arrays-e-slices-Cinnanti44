package main

import "fmt"

func main() {
	var n int
	fmt.Print("Informe o tamanho dos arrays: ")
	fmt.Scan(&n)

	// Lendo os elementos do primeiro array
	arr1 := make([]int, n)
	fmt.Println("Informe os elementos do primeiro array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr1[i])
	}

	// Lendo os elementos do segundo array
	arr2 := make([]int, n)
	fmt.Println("Informe os elementos do segundo array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr2[i])
	}

	// Somando os elementos dos dois arrays
	arr3 := make([]int, n)
	for i := 0; i < n; i++ {
		arr3[i] = arr1[i] + arr2[i]
	}

	// Imprimindo o resultado
	fmt.Println("Resultado da soma:")
	fmt.Println(arr3)
}
