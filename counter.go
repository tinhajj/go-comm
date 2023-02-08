package main

import "sync"

type Counter struct {
	mu  sync.Mutex
	num uint64
}

func (c *Counter) Next() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	current := c.num
	c.num++

	return current
}
