package main

import "fmt"

func main() {
	valores := [6]float64{10.0, 5.0, 8.0, 6.5, 7.5, 9.0}

	var soma float64 = 0.0

	for i := 0; i < len(valores); i++ {
		soma += valores[i]
	}

	media := soma / float64(len(valores))

	fmt.Printf("A média dos valores é: %.2f", media)
}
