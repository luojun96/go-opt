package main

import "fmt"

const queueSize = 10

func start(id string, f func(string) int, next chan<- struct{}) (chan<- struct{}, chan<- struct{}, <-chan int) {
	queue := make(chan struct{}, queueSize)
	quit := make(chan struct{})
	result := make(chan int)

	go func() {
		total := 0
		for {
			select {
			case <-quit:
				result <- total
			case v := <-queue:
				total += f(id)
				if next != nil {
					next <- v
				}
			}
		}
	}()

	return queue, quit, result
}

func newAirportSecurityCheckChannel(id string, queue <-chan struct{}) {
	go func(id string) {
		fmt.Printf("goroutine-%s: airportSecurityCheckChannel is ready...\n", id)
		// start xRayCheck routine
		queue3, quit3, result3 := start(id, xRayCheck, nil)
		// start bodyCheck routine
		queue2, quit2, result2 := start(id, bodyCheck, queue3)
		// start idCheck routine
		queue1, quit1, result1 := start(id, idCheck, queue2)

		for {
			select {
			case v, ok := <-queue:
				if !ok {
					close(quit1)
					close(quit2)
					close(quit3)
					total := max(<-result1, <-result2, <-result3)
					fmt.Printf("goroutine-%s: airportSecurityCheckChannel time cost:%d\n", id, total)
					fmt.Printf("goroutine-%s: airportSecurityCheckChannel closed\n", id)
					return
				}
				fmt.Printf("goroutine-%s: a passenger enter into queue.\n", id)
				queue1 <- v
			}
		}
	}(id)
}

func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}
