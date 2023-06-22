package main

import "fmt"

func main() {
	// Criando o slice com valores iniciais
	slice := []string{"foo", "bar", "foo", "baz", "foo", "qux", "foo", "quux"}

	// Solicitando o valor a ser removido
	var valor string
	fmt.Println("Informe o valor a ser removido:")
	fmt.Scan(&valor)

	// Removendo as ocorrências do valor no slice
	novoSlice := []string{}
	for _, s := range slice {
		if s != valor {
			novoSlice = append(novoSlice, s)
		}
	}

	// Imprimindo o slice resultante
	fmt.Println("Slice resultante:", novoSlice)

}
