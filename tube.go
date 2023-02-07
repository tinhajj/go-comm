package main

type Tube struct {
	In  chan Message
	Out chan Message
}

func NewTube() *Tube {
	return &Tube{
		In:  make(chan Message),
		Out: make(chan Message),
	}
}

func (t *Tube) Send(Message) int {
	return 0
}
