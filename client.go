package main

type Client struct {
	Connection *Connection
}

func NewClient() *Client {
	return &Client{
		Connection: nil,
	}
}

func (c *Client) Send(m Message) {
	c.Connection.Send(m)
}
