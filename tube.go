package main

type Tube struct {
	In  chan Message
	Out chan Message
}

func NewTube(size int) *Tube {
	return &Tube{
		In:  make(chan Message, size),
		Out: make(chan Message, size),
	}
}

func (t *Tube) Send(Message) int {
	return 0
}
