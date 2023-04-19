package main

import (
	"fmt"
	"time"
)

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	fmt.Printf("\tgoroutine-%s-idCheck: idCheck ok\n", id)
	return idCheckTmCost
}

func bodyCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	fmt.Printf("\tgoroutine-%s-bodyCheck: bodyCheck ok\n", id)
	return bodyCheckTmCost
}

func xRayCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	fmt.Printf("\tgoroutine-%s-xRayCheck: xRayCheck ok\n", id)
	return xRayCheckTmCost
}

func start(id string, f func(string) int, next chan<- struct{}) (chan<- struct{}, chan<- struct{}, <-chan int) {
	queue := make(chan struct{}, 10)
	quit := make(chan struct{})
	result := make(chan int)

	go func() {
		total := 0
		for {
			select {
			case <-quit:
				result <- total
				return
			case v := <-queue:
				total += f(id)
				if next != nil {
					next <- v
				}
			default:

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
				queue1 <- v
			default:
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

func main() {
	passengers := 200
	queue := make(chan struct{}, 30)
	newAirportSecurityCheckChannel("channel1", queue)
	newAirportSecurityCheckChannel("channel2", queue)
	newAirportSecurityCheckChannel("channel3", queue)

	time.Sleep(5 * time.Second)
	for i := 0; i < passengers; i++ {
		queue <- struct{}{}
	}
	time.Sleep(time.Duration(100 * time.Second))
	close(queue)
	time.Sleep(100 * time.Second)
}
