package main

import "fmt"

//tread 1
func main() {
	canal := make(chan string)
	//tread 2
	go func() {
		canal <- "Olá Mundo!"
	}()

	//tread 1
	msg := <-canal
	fmt.Println(msg)

}
