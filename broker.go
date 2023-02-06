package main

type Broker struct {
	Router      *Router
	Connections []*Connection
}

func NewBroker(r *Router) *Broker {
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
			c.Tube.Out <- b.Router.Handle(msg)
		}
	}()

	c.Tube = t
}
