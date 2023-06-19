package main

import (
	"fmt"
	"time"
)

func doRun() {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		fmt.Println("loop", i)
		select {
		case ch <- i:
			fmt.Println("send", i)
		case v := <-ch:
			fmt.Println("receive", v)
		}
	}
	time.Sleep(60 * time.Second)
}
