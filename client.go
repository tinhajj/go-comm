package main

type Client struct {
	Tube *Tube
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Send() {
	c.Tube.In <- Message{
		Name: "restart",
		Data: nil,
	}
}
