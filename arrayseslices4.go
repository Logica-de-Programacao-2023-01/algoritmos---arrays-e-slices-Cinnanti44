package main

import "fmt"

func main() {
	matriz := [3][2]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println("Informe o valor do elemento: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	fmt.Println("Matriz:", matriz)

}
