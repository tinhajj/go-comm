package main

type Handler func(data []byte) Message

type Message struct {
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

func (r *Router) Handle(message Message) Message {
	h, ok := r.Routes[message.Name]
	if !ok {
		return Message{
			Name: "nohandle",
			Data: nil,
		}
	}

	return h(message.Data)
}
