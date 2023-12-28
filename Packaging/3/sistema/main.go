package main

import "github.com/danubiobwm/goexpert/Packaging/3/math"

func main() {
	m := math.NewMath(1, 2)
	println("Soma de 1+2=", m.Add())
}
