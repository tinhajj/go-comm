package main

var in chan int = make(chan int)
var out chan int = make(chan int)

func server() {
	for message := range in {
		handle(message)
	}
	close(out)
}

func handle(message int) {
	out <- message
}
