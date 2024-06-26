package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c", "d"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}
	for range values {
		<-done
	}
}
