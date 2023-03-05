package main

import (
	"fmt"
	"time"
)

func example1() {
	r := NewRouter()
	r.AddRoute("restart", func(msg Message) Message {
		return Message{
			ID:   msg.ID,
			Name: "restartapproved",
			Data: []byte{},
		}
	})

	b := NewBroker(r)

	c1 := b.NewConnection()
	c2 := b.NewConnection()

	counter := 0
	go func() {
		for range c1.Tube.Out {
			//fmt.Println("client 1 received:", message)
			counter++
		}
	}()

	go func() {
		for range c2.Tube.Out {
		}
	}()

	now := time.Now()
	future := now.Add(time.Millisecond * 1000)

	for time.Now().Before(future) {
		c1.Send(Message{
			ID:   0,
			Name: "restart",
			Data: nil,
		})
	}

	fmt.Println("sleep main routine for 500ms")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("received messages:", counter, "in 1000ms")
	fmt.Println("end")
}
