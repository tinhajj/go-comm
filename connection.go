package main

var counter = new(Counter)

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

func (c *Connection) Send(m Message) uint64 {
	m.ID = counter.Next()
	c.Tube.In <- m
	return m.ID
}

func (c *Connection) WaitFor(ID int) {
}
