package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"wesley": 2000, "Joao": 3000, "Danubio": 4000}
	m2 := map[string]float64{"wesley": 20.0, "Joao": 30.0, "Danubio": 40.0}

	m3 := map[string]MyNumber{"wesley": 20, "Joao": 30, "Danubio": 40}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 9))
}
