package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	danubio := Cliente{Nome: "Danubio", Idade: 38, Ativo: true}

	fmt.Println(danubio)
	fmt.Println(danubio.Idade)
	danubio.Ativo = false
	fmt.Println(danubio.Ativo)
}
