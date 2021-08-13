package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, err := nats.Connect("nats://localhost:4222",
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
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	fmt.Printf("ConnectedUrl: %s\n", nc.ConnectedUrl())

	i := 0
	for {
		i++
		nc.Publish("demo.cluster", []byte("Msg "+strconv.Itoa(i)))
		fmt.Printf("Sending %d\n", i)
		time.Sleep(2 * time.Second)
	}
}
