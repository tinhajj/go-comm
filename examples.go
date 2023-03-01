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
	c1 := NewClient()
	c2 := NewClient()

	b.ConnectTo(c1)
	b.ConnectTo(c2)

	counter := 0
	go func() {
		for message := range c1.Connection.Tube.Out {
			fmt.Println("client 1 received:", message)
			counter++
		}
	}()

	go func() {
		for message := range c2.Connection.Tube.Out {
			fmt.Println("client 2 received:", message)
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

	fmt.Println("counted", counter, "in 1000ms")

	fmt.Println("end")
}

func example2() {
	/*
		websocketserver

		it receives a message
		its from the frontend asking to do something and it wants a response
		i send it to the client
		i need to wait for the response
		i send it to the frontend

		client sends me a message
		i send it to the frontend
	*/
}
