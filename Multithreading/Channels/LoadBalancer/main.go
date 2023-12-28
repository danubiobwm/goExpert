package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d got %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	Qtd := 10

	for i := 0; i < Qtd; i++ {
		go worker(i, data)
	}

	go worker(1, data)
	go worker(2, data)

	for i := 0; i < 100; i++ {
		data <- i
	}
}
