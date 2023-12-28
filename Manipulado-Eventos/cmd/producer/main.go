package main

import "github.com/danubiobwm/goExpert/fcutils/packages/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}
