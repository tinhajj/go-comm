package main

import (
	"fmt"
	"go-comm/server"
)

func main() {
	fmt.Println("start")
	s := server.NewServer()
	c1 := server.NewClient()
	c2 := server.NewClient()
	c1.Connect(s)
	c2.Connect(s)

	go func() {
		for message := range c1.In {
			fmt.Println("client 1 received:", message)
		}
	}()

	go func() {
		for message := range c2.In {
			fmt.Println("client 2 received:", message)
		}
	}()

	c1.Send(3)
	c1.Send(3)
	c1.Send(3)

	c2.Send(3)
	c2.Send(3)
	c2.Send(3)

	fmt.Println("end")
}
