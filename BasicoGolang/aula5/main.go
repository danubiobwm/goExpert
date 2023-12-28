package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 30
	fmt.Println(len(meuArray) - 1)
	fmt.Println(meuArray[len(meuArray)-1])
	fmt.Println(meuArray[2])

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é o valor é %d \n", i, v)
	}

}
