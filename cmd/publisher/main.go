package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	interval = 1 * time.Second // we will send a message every 1 second
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
			subject := fmt.Sprintf("demo.ququq.lalala.%d", i) // subject (equivalent of the topic) is dynamic
			msg := fmt.Sprintf("msg number %d", i)            // message is always a text of some kind. May be JSON or better CludEvents spec
			fmt.Println(subject, msg)                         // irrelevant, just to know what is going on
			nc.Publish(subject, []byte(msg))                  // here we send the message as slice of bytes to nats server
			tmr.Reset(interval)                               // we must reset the timer if we want to reuse it
		case <-sig: // Ctrl+C captured
			return
		}
	}
}
