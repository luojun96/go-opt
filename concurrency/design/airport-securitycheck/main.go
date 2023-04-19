package main

import "time"

func main() {
	passengers := 100
	queue := make(chan struct{}, 30)
	newAirportSecurityCheckChannel("channel1", queue)
	newAirportSecurityCheckChannel("channel2", queue)
	newAirportSecurityCheckChannel("channel3", queue)

	time.Sleep(5 * time.Second)
	for i := 0; i < passengers; i++ {
		queue <- struct{}{}
	}
	time.Sleep(60 * time.Second)
	close(queue)
	time.Sleep(120 * time.Second)
}
