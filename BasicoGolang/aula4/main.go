package main

import "fmt"

const a = "Hello world!"

// Declaração de Tipagem
type ID int

// Tipagem primitivos
var (
	b bool    = true
	c int     = 10
	d string  = "Danubio"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de F é %T", f)
}
