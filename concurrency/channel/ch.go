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
				fmt.Println("closing...")
				return
			default:
				fmt.Println("running...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	fmt.Println()
	close(closing)
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(15 * time.Second):
		fmt.Println("cleanup time is over...")
	}
	fmt.Println("gracefully exit")
}

func doCleanup(closed chan struct{}) {
	fmt.Println("do cleanup...")
	time.Sleep(10 * time.Second)
	close(closed)
}
