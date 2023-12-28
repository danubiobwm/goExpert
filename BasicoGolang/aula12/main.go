package main

import "fmt"

type Endereco struct {
	Logradoro string
	Numero    int
	Cidade    string
	Estado    string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Addes Endereco
}

func main() {
	danubio := Cliente{Nome: "Danubio", Idade: 38, Ativo: true}
	danubio.Addes.Cidade = "Jua"

	fmt.Println(danubio)
	fmt.Println(danubio.Idade)
	danubio.Ativo = false
	fmt.Println(danubio.Ativo)
}
