package main

import "fmt"

func main() {
	array := [10]int{3, 8, 1, 6, 9, 2, 7, 5, 4, 0}

	var valor int
	fmt.Print("Informe um valor: ")
	fmt.Scan(&valor)

	encontrado := false
	for _, v := range array {
		if v == valor {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println("O valor", valor, "está presente no Array")
	} else {
		fmt.Println("O valor", valor, "não está presente no Array")
	}
}
