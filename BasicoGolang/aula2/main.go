package main

const a = "Hello world!"

var (
	b bool    = true
	c int     = 10
	d string  = "Danubio"
	e float64 = 1.2
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
}

func x() {}
