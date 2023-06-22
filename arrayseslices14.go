package main

import "fmt"

func main() {
	slice := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	var indice1, indice2 int
	fmt.Print("digite o indice do primeiro elemento:")
	fmt.Scanln(&indice1)
	fmt.Print("digite o indice do segundo elemento:")
	fmt.Scanln(&indice2)

	// Realizando a troca de posição dos elementos
	slice[indice1], slice[indice2] = slice[indice2], slice[indice1]

	// Imprimindo o Slice resultante
	fmt.Println(slice)
}
