package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.TODO()
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go task(cancelCtx)
	time.Sleep(3 * time.Second)
	fmt.Println("Cancel the task")
	cancel()
	time.Sleep(1 * time.Second)
}

func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
			i++
		}
	}
}
