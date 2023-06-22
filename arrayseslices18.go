package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("O número precisa ser positivo.")
		return
	}

	// Inicializando o array com os n primeiros números primos
	array := make([]int, n)
	count := 0
	for i := 2; count < n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			array[count] = i
			count++
		}
	}

	fmt.Println(array)

}
