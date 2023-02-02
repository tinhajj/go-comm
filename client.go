package main

import "fmt"

func client() {
	go send()
	go read()
}

func send() {
	for i := 0; i < 3; i++ {
		in <- 70
	}
	close(in)
}

func read() {
	for message := range out {
		fmt.Println(message)
	}
}
