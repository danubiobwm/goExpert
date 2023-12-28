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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {
	danubio := Cliente{Nome: "Danubio", Idade: 38, Ativo: true}
	danubio.Addes.Cidade = "Jua"

	danubio.Desativar()

}
