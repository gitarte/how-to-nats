package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.QueueSubscribe("service.potęgator", "dupa", func(msg *nats.Msg) {
		fmt.Printf("Got: %s\n", msg.Data)
		i, _ := strconv.ParseInt(string(msg.Data), 10, 64)
		p := math.Pow(float64(i), 2)
		msg.Respond([]byte(fmt.Sprintf("%f", p)))
	})

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
