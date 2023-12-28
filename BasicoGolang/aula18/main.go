package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello world!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da Variavel %T e o valor e %v \n", t, t)
}
