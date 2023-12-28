package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thead 1: running
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	//Thead 2: Task A is running
	go task("A", &waitGroup)
	//Thead 3: Task B is running
	go task("B", &waitGroup)

	//Thead 4: Task B is running
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
