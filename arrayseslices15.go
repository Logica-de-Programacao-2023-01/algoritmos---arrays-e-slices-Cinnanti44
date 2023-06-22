package main

import "fmt"

func main() {
	array := [10]float64{1.2, 6.7, 8.9, 4.5, 2.1, 10.0, 3.4, 5.6, 7.8, 9.0}
	slice := make([]float64, 0)

	for _, element := range array {
		if element > 5 {
			slice = append(slice, element)
		}
	}

	fmt.Println("Array original:", array)
	fmt.Println("Slice resultante:", slice)

}
