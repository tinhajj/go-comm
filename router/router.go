package server

type Router struct {
	In  chan int
	Out chan int
}

func NewRouter() *Router {
	r := &Router{
		In:  make(chan int),
		Out: make(chan int),
	}
	go func() {
		for input := range r.In {
			r.handle(input)
		}
		close(r.Out)
	}()
	return r
}

// Decide what logic should handle the input here
func (r *Router) handle(input int) {
	// Echo the input to the output right away for now.  Later a real
	// handler will do something before it goes to the output
	r.Out <- input
}
