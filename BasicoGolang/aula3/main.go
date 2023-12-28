package main

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
	a := "Hello world!"
	b = false
	c = 30
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
