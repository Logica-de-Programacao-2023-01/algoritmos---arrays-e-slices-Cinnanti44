package main

import "fmt"

func main() {
	array := [4]float64{2.0, 3.5, 1.0, 4.5}

	var produto float64 = 1.0
	for i := 0; i < len(array); i++ {
		produto *= array[i]
	}

	fmt.Println("O produto dos elementos do array é: ", produto)
	
}
