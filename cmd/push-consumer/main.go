package main

import (
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	js.Subscribe("demo.>", func(msg *nats.Msg) {
		msg.Ack()
		log.Printf("message: %s\n", msg.Data)
	},
		nats.Durable("qqryq"),
		nats.ManualAck(),
	)

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
