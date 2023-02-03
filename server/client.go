package server

type Client struct {
	In     chan int
	server *Server
}

func NewClient() *Client {
	return &Client{
		In: make(chan int),
	}
}

func (c *Client) Connect(s *Server) {
	c.server = s
	s.clients = append(s.clients, c)
}

func (c *Client) Send(message int) {
	if c.server == nil {
		return
	}

	c.server.In <- message
}
