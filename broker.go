package main

type Broker struct {
	In          chan Message
	Router      *Router
	Connections []*Connection
}

func NewBroker(r *Router) *Broker {
	return &Broker{
		Connections: []*Connection{},
		Router:      r,
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

func (b *Broker) Add(m Message) int {
	return 3
}
