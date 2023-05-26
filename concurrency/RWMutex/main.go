package main

import (
	"fmt"
	"time"
)

func main() {
	var counter Counter
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
