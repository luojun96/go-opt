package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter AnonymousCounter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Println("count: ", counter.Count())
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		fmt.Println("incr")
		counter.Incr()
		time.Sleep(time.Second)
	}
}

type AnonymousCounter struct {
	sync.RWMutex
	count uint64
}

func (ac *AnonymousCounter) Incr() {
	ac.Lock()
	ac.count++
	ac.Unlock()
}

func (ac *AnonymousCounter) Count() uint64 {
	ac.RLock()
	defer ac.RLocker().Unlock()
	return ac.count
}
