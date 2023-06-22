package main

import "fmt"

func main() {
	array := [3]int{2, 4, 6}

	var soma int = 0
	for i := 0; i < len(array); i++ {
		soma += array[i]
	}

	fmt.Println("A soma dos elementos é:", soma)

}
