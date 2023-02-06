package main

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Handle(message int) int {
	return message * 10
}
