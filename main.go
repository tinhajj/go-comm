package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	r := NewRouter()
	r.AddRoute("restart", func(msg Message) Message {
		return Message{
			ID: msg.ID,
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
		for message := range c1.Out {
			fmt.Println("client 1 received:", message)
			counter++
		}
	}()

	go func() {
		for message := range c2.Out {
			fmt.Println("client 2 received:", message)
		}
	}()

	now := time.Now()
	future := now.Add(time.Millisecond * 1000)


	for time.Now().Before(future) {
		c1.Send()
	}


	fmt.Println("sleep main routine for 500ms")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("counted", counter)

	fmt.Println("end")
}
