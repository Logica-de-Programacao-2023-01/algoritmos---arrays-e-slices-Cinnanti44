package main

import "fmt"

func main() { // Criando o array bidimensional com valores aleatórios
	array := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// Solicitando ao usuário o índice de linha e coluna
	var linha, coluna int
	fmt.Print("Informe o índice da linha: ")
	fmt.Scanln(&linha)
	fmt.Print("Informe o índice da coluna: ")
	fmt.Scanln(&coluna)

	// Imprimindo o valor armazenado na posição informada
	fmt.Println("Valor armazenado:", array[linha][coluna])

}
