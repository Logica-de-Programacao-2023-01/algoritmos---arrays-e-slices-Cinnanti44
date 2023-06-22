package main

import "fmt"

func main() {
	// Criar um Array de inteiros com 7 elementos e preencher com alguns valores iniciais
	array := [7]int{2, 4, 6, 8, 10, 12, 14}

	// Solicitar ao usuário que informe um número para ser adicionado ao primeiro e ao último elemento do Array
	var numero int
	fmt.Print("Informe um número para ser adicionado ao primeiro e ao último elemento do Array: ")
	fmt.Scanln(&numero)

	// Adicionar o número informado pelo usuário ao primeiro e ao último elemento do Array
	array[0] += numero
	array[len(array)-1] += numero

	// Imprimir o Array resultante
	fmt.Println("Array resultante:", array)

}
