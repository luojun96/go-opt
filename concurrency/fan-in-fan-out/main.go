package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Duration(job) * time.Millisecond)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 扇出：启动 3 个 worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送 9 个任务
	for j := 1; j <= 9; j++ {
		jobs <- j * 100
	}
	close(jobs)

	// 扇入：收集任务结果
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 1; i <= 9; i++ {
			fmt.Printf("Result: %d\n", <-results)
		}
		wg.Done()
	}()

	wg.Wait()
}
