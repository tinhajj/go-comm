package main

type Broker struct {
	Counter *Counter

	In          chan Message
	Router      *Router
	Connections []*Connection
}

func NewBroker(r *Router) *Broker {
	return &Broker{
		Counter:     new(Counter),
		In:          make(chan Message, 100),
		Router:      r,
		Connections: []*Connection{},
	}
}

func (b *Broker) ConnectTo(c *Client) {
	b.Connections = append(b.Connections, &Connection{
		Client: c,
	})

	go func() {
		for msg := range b.In {
			c.Out <- b.Router.Handle(msg)
		}
	}()

	c.Broker = b
}

func (b *Broker) Add(m Message) uint64 {
	m.ID = b.Counter.Next()
	b.In <- m
	return m.ID
}
