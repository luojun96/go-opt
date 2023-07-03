package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	cancelCtx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()
	go task(cancelCtx)
	time.Sleep(6 * time.Second)
}

func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
			i++
		}
	}
}
