package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thead 1: running
func main() {
	//Thead 2: Task A is running
	go task("A")
	//Thead 3: Task B is running
	go task("B")

	//Thead 4: Task B is running
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	//Nada Aqui //Sair do programa
	time.Sleep(10 * time.Second)
}
