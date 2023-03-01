package main

type Broker struct {
	Counter *Counter

	Router      *Router
	Connections []*Connection
}

func NewBroker(r *Router) *Broker {
	return &Broker{
		Counter:     new(Counter),
		Router:      r,
		Connections: []*Connection{},
	}
}

func (b *Broker) ConnectTo(c *Client) {
	connection := NewConnection(c)
	b.Connections = append(b.Connections, connection)
	c.Connection = connection

	go func() {
		for msg := range connection.Tube.In {
			connection.Tube.Out <- b.Router.Handle(msg)
		}
	}()
}
