package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sema       = semaphore.NewWeighted(int64(maxWorkers))
	tasks      = make([]int, maxWorkers*10)
)

func main() {
	ctx := context.Background()

	log.Printf("Max workers: %d, tasks: %d", maxWorkers, len(tasks))
	for i := range tasks {
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}
		go func(i int) {
			defer sema.Release(1)
			log.Printf("Task %d started", i)
			time.Sleep(100 * time.Millisecond)
			tasks[i] = i + 1
		}(i)
	}

	// request all workers to ensure all tasks are done
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Fatalf("Failed to acquire semaphore: %v", err)
	}

	log.Printf("Tasks: %v", tasks)
}
