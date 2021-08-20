package main

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	interval = 1 * time.Second // we will send a message every 1 second
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	sub, _ := js.PullSubscribe("demo.>", "lalabbbhhy7tgwwwwwla")

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		msgs, _ := sub.Fetch(10, nats.Context(ctx))
		for _, msg := range msgs {
			msg.Ack()
			log.Printf("message: %s\n", msg.Data)
		}
	}
}
