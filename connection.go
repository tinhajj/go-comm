package main

type Connection struct {
	Tube *Tube
}

func NewConnection() *Connection {
	return &Connection{
		Tube: NewTube(100),
	}
}

func (c *Connection) Send(m Message) {
	c.Tube.In <- m
}
