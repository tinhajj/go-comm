package main

type Handler func(m Message) Message

type Message struct {
	ID   uint64
	Name string
	Data []byte
}

type Router struct {
	Routes map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		Routes: make(map[string]Handler),
	}
}

func (r *Router) AddRoute(name string, handler Handler) {
	r.Routes[name] = handler
}

func (r *Router) Handle(msg Message) Message {
	h, ok := r.Routes[msg.Name]
	if !ok {
		return Message{
			ID:   msg.ID,
			Name: "nohandle",
			Data: nil,
		}
	}

	return h(msg)
}
