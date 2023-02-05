package main

type Tube struct {
	In  chan int
	Out chan int
}

func NewTube() *Tube {
	return &Tube{
		In:  make(chan int),
		Out: make(chan int),
	}
}
