package main

import (
	"fmt"
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

	// confront with te subject set by the publisher
	nc.Subscribe("demo.>", func(msg *nats.Msg) {
		fmt.Printf("%s\n", msg.Data)
	})

	// quick cheat against premature ending of the program
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
