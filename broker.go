package main

type Broker struct {
	Connections []*Connection
}

func NewBroker() *Broker {
	return &Broker{
		Connections: []*Connection{},
	}
}

func (b *Broker) ConnectTo(c *Client) {
	t := NewTube()

	b.Connections = append(b.Connections, &Connection{
		Client: c,
		Tube:   t,
	})

	go func() {
		for msg := range c.Tube.In {
			c.Tube.Out <- msg * 2
		}
	}()

	c.Tube = t
}
