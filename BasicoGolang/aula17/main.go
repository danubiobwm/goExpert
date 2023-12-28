package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c Cliente) andou() {
	c.nome = "Danubio Araujo"
	fmt.Printf("O cliente %v andou", c.nome)
	fmt.Printf("\n")
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
	//fmt.Printf("O cliente Possui %v ", c.saldo)
	//fmt.Printf("\n")
}

func main() {
	danubio := Cliente{nome: "Danubio"}
	conta := Conta{saldo: 5558}
	conta.simular(558)

	danubio.andou()

	fmt.Printf("O valor da struct com nome %v", danubio.nome)
	fmt.Printf("\n")
	fmt.Printf("O valor da conta %v", conta.saldo)
}
