package main

import "fmt"

func main() {
	array := [6]float64{1.5, 2.3, 4.0, 3.2, 2.1, 5.7}

	var valor float64
	fmt.Println("Informe o valor a ser adicionado em todas as posições:")
	fmt.Scan(&valor)

	for i := 0; i < len(array); i++ {
		array[i] += valor
	}

	fmt.Println("Array resultante:", array)

}
