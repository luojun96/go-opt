package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		for {
			select {
			case <-closing:
				return
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("cleanup time is over...")
	}
	fmt.Println("gracefully exit")
}

func doCleanup(closed chan struct{}) {
	time.Sleep(time.Minute)
	close(closed)
}
