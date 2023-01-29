package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	cancelCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	go task1(cancelCtx)
	time.Sleep(4 * time.Second)
}

func task1(ctx context.Context) {
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
