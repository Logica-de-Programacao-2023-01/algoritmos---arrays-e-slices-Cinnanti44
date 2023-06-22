package main

import "fmt"

func main() {
	slice := []int{3, 8, 1, 6, 9}

	var valor int
	fmt.Print("Informe um valor inteiro: ")
	fmt.Scan(&valor)

	encontrado := false
	for _, v := range slice {
		if v == valor {
			encontrado = true
			break
		}
	}

	if !encontrado {
		slice = append(slice, valor)
		fmt.Println("O valor", valor, "foi adicionado ao Slice")
	} else {
		fmt.Println("O valor", valor, "já está presente no Slice")
	}

	fmt.Println("Slice resultante:", slice)
}
