package main

import (
	"context"
	"log"
	"time"

	"github.com/mdlayher/schedgroup"
)

func main() {
	sg := schedgroup.New(context.Background())

	for i := 0; i < 3; i++ {
		n := i + 1
		sg.Delay(time.Duration(n)*100*time.Millisecond, func() {
			log.Println(n)
		})
	}

	if err := sg.Wait(); err != nil {
		log.Fatalf("failed to wait: %v", err)
	}
}
