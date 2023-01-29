package main

import (
	"sync"
	"time"
)
func main() {
	var counter Counter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				counter.Count()
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		counter.Incr()
		time.Sleep(time.Second)
	}
}

type Counter struct {
	mu sync.RWMutex
	count uint64
}

type anonymousCounter struct {
	sync.RWMutex
	count uint64
}

func (ac *anonymousCounter) Incr() {
	ac.Lock()
	ac.count++
	ac.Unlock()
}

func (ac *anonymousCounter) Count() uint64 {
	ac.RLock()
	defer ac.RLocker().Unlock()
	return ac.count
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}


