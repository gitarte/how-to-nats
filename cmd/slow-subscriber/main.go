package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	draggingTime = 2 * time.Second
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL, nats.ErrorHandler(func(c *nats.Conn, s *nats.Subscription, err error) {
		fmt.Println(err)
	}))
	defer nc.Close()

	sub, _ := nc.Subscribe("demo.>", func(msg *nats.Msg) {
		time.Sleep(draggingTime)
		// Pending returns the number of queued messages and queued bytes in the client for this subscription.
		queued, _, _ := msg.Sub.Pending()

		// Dropped returns the number of known dropped messages for this subscription.
		// This will correspond to messages dropped by violations of PendingLimits. If
		// the server declares the connection a SlowConsumer, this number may not be
		// valid.
		dropped, _ := msg.Sub.Dropped()
		fmt.Printf("%s     Queued: %d,  Dropped: %d\n", msg.Data, queued, dropped)
	})

	fmt.Printf("Subscribed %s\n", sub.Queue)

	sub.SetPendingLimits(10, 1024)

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
