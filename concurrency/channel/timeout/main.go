package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	closed := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go doCleanup(closed)

	select {
	case <-closed:
		fmt.Println("cleanup is done...")
	case <-ctx.Done():
		fmt.Println("cleanup time is over...")
	}
	fmt.Println("graceful shutdown...")
}

func doCleanup(closed chan struct{}) {
	// do cleanup
	fmt.Println("do cleanup...")
	time.Sleep(3 * time.Second)
	close(closed)
}
