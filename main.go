package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	r := NewRouter()
	b := NewBroker(r)
	c1 := NewClient()
	c2 := NewClient()

	b.ConnectTo(c1)
	b.ConnectTo(c2)

	go func() {
		for message := range c1.Tube.Out {
			fmt.Println("client 1 received:", message)
		}
	}()

	go func() {
		for message := range c2.Tube.Out {
			fmt.Println("client 2 received:", message)
		}
	}()

	c1.Send()
	c1.Send()
	c1.Send()

	c2.Send()
	c2.Send()
	c2.Send()

	fmt.Println("end")
}
