package main

import "sync"

type Counter struct {
	lock  sync.RWMutex
	count uint64
}

func (c *Counter) Incr() {
	c.lock.Lock()
	c.count++
	c.lock.Unlock()
}

func (c *Counter) Count() uint64 {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.count
}
