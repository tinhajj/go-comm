package server

type Server struct {
	In      chan int
	clients []*Client
}

func NewServer() *Server {
	s := &Server{
		In: make(chan int),
	}
	go func() {
		for input := range s.In {
			s.handle(input)
		}
	}()
	return s
}

// Decide what logic should handle the input here
func (s *Server) handle(input int) {
	// Echo the input to the clients right away for now.  Later a real
	// handler/router will do something before it goes to the clients.
	for _, c := range s.clients {
		c.In <- input
	}
}
