package main

import "fmt"

func main() {
	salarios := map[string]int{"Danubio": 1000, "Wesley": 2000, "maria": 4000}
	//fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 10000
	//fmt.Println(salarios["Wes"])

	//funçao make
	/*sal := make(map[string]int)
	sal1 := map[string]int{}
	sal1["Danubio"] = 3000*/

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	//ignorar uma variavel com  _
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
