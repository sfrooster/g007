package main

import (
	"fmt"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	fmt.Println("NATS Code")

	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Simple Sync Subscriber
	var timeout = 5 * time.Second

	sub, err := nc.SubscribeSync("foo")

	if err != nil {
		panic(err)
	}

	_, err = sub.NextMsg(timeout)

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err = nc.ChanSubscribe("foo", ch)
	_ = <-ch

	// Unsubscribe
	sub.Unsubscribe()

	// Requests
	_, err = nc.Request("help", []byte("help me"), 10*time.Millisecond)

	// Replies
	// nc.Subscribe("help", func(m *Msg) {
	// 	nc.Publish(m.Reply, []byte("I can help!"))
	// })

	// Close connection
	nc, _ = nats.Connect("nats://localhost:4222")
	nc.Close()
}
