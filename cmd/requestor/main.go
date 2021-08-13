package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	interval = 1 * time.Second
	timeout  = 10 * time.Second
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// create a timer with duration of 1 second
	tmr := time.NewTimer(interval)

	// we need to be able to handle Ctrl+C gracefully
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	i := 0 // this is just iteration counter
	for {
		select {
		case <-tmr.C:
			i++
			res, err := nc.Request("service.potęgator", []byte(strconv.Itoa(i)), timeout)
			if err != nil {
				fmt.Printf("%s\n", err)
			} else {
				fmt.Printf("%d: %s\n", i, res.Data)
			}
			tmr.Reset(interval)
		case <-sig: // Ctrl+C captured
			return
		}
	}
}
