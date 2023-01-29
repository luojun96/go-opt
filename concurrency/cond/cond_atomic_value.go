package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

func loadNewConfig() Config {
	return Config{
		NodeName: "beijing",
		Addr:     "192.168.31.18",
		Count:    rand.Int31(),
	}
}

func Run() {
	var config atomic.Value
	config.Store(Config{})
	var cond = sync.NewCond(&sync.Mutex{})

	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast()
		}
	}()

	go func() {
		for {
			cond.L.Lock()
			cond.Wait()
			c := config.Load().(Config)
			fmt.Printf("new config: %+v\n", c)
			cond.L.Unlock()
		}
	}()

	select {}
}
