package main

type Client struct {
	Broker *Broker
	Out    chan Message
}

func NewClient() *Client {
	return &Client{
		Broker: nil,
		Out:    make(chan Message),
	}
}

func (c *Client) Send() {
	c.Broker.Add(Message{
		ID:   0,
		Name: "restart",
		Data: nil,
	})
}
