package main

type Broker struct {
	Router      *Router
	Connections []*Connection
}

func NewBroker(r *Router) *Broker {
	return &Broker{
		Router:      r,
		Connections: []*Connection{},
	}
}

func (b *Broker) NewConnection() *Connection {
	conn := NewConnection()
	b.Connections = append(b.Connections, conn)

	go func() {
		for msg := range conn.Tube.In {
			conn.Tube.Out <- b.Router.Handle(msg)
		}
	()

	return conn
}
