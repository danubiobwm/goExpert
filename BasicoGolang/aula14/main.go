package main

import "fmt"

type Endereco struct {
	Logradoro string
	Numero    int
	Cidade    string
	Estado    string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

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

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	danubio := Cliente{Nome: "Danubio", Idade: 38, Ativo: true}
	danubio.Addes.Cidade = "Jua"

	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)

}
