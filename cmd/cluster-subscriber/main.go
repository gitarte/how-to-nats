package main

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, _ := nats.Connect("nats://localhost:6222",
		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			fmt.Printf("Disconnect\n")
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			fmt.Printf("Reconnect to %s\n", conn.ConnectedUrl())
		}),
		nats.ErrorHandler(func(conn *nats.Conn, subscription *nats.Subscription, err error) {
			fmt.Printf("Error %s!", err)
		}),
		nats.ClosedHandler(func(conn *nats.Conn) {
			fmt.Println("Closed!")
		}),
		nats.DiscoveredServersHandler(func(conn *nats.Conn) {
			fmt.Printf("Server found! Servers: %s\n", conn.DiscoveredServers())
		}),
	)
	defer nc.Close()

	fmt.Printf("DiscoveredServers: %s\n", nc.DiscoveredServers())
	fmt.Printf("ConnectedUrl: %s\n", nc.ConnectedUrl())

	nc.Subscribe("demo.cluster", func(msg *nats.Msg) {
		fmt.Printf("%s\n", msg.Data)
	})

	// quick cheat against premature ending of the program
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
