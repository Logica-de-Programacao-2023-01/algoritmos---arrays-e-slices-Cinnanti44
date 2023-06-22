package main

import "fmt"

func main() {
	matriz := [3][4]int{}

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			matriz[i][j] = i + j
		}
	}

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Printf("\n")
	}

}
