package main

import (
	"fmt"
	"time"
)

type Message struct {
	id  int
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	//RabitMQ
	go func() {
		i := 0
		for {
			i++
			msg := Message{i, "RabitMQ"}

			c1 <- msg
			time.Sleep(time.Second * 2)
		}
	}()

	//Kafka
	go func() {
		i := 0
		for {
			i++
			msg := Message{i, "Kafka"}

			c2 <- msg
			time.Sleep(time.Second * 1)
		}
	}()

	for {

		select {
		case msg1 := <-c1:
			fmt.Printf("Received from RabbitMQ: %d - %s\n", msg1.id, msg1.Msg)

		case msg2 := <-c2:
			fmt.Printf("Received from Kafka: %d - %s\n", msg2.id, msg2.Msg)

		case <-time.After(time.Second * 3):
			println("timeout")

			// default:
			// 	println("nothing ready")
		}

	}
}
