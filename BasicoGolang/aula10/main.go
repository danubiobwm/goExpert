package main

import (
	"fmt"
)

func main() {

	total := func() int {
		return soma(1, 5, 45, 100, 545, 600) * 2
	}()

	fmt.Println(total)

	fmt.Println(soma(1, 5, 45, 100, 545, 600))
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
