package main

type Connection struct {
	Client *Client
	Tube   *Tube
}

func NewConnection(c *Client) *Connection {
	return &Connection{
		Client: c,
		Tube:   NewTube(100),
	}
}

func (c *Connection) Send(m Message) {
	c.Tube.In <- m
}
