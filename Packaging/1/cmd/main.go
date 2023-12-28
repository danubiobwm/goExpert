package main

import (
	"fmt"

	"github.com/danubiobwm/goexpert/Packaging/1/math"
)

func main() {

	m := math.NewMath(2, 3)
	fmt.Println(m.Add())
	fmt.Println(math.X)
}
