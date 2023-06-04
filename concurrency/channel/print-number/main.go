package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	products := make(chan int, 100)
	defer close(products)

	done := make(chan bool)

	// start to produce products
	fmt.Println("start 1 producer...")
	go producer(products, done)

	fmt.Println("start 10 consumers...")
	for i := 0; i < 10; i++ {
		go consumer(products, done)
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	fmt.Println("stop all producers and consumers...")
	close(done)
}

func producer(products chan<- int, done <-chan bool) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			fmt.Println("producer is interrupted...")
			return
		case <-ticker.C:
			number := rand.Intn(100)
			products <- number
			fmt.Println("\t", "produce product: ", number)
		}
	}
}

func consumer(products <-chan int, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("consumer is interrupted...")
			return
		case product := <-products:
			fmt.Println("consume product: ", product)
		}
	}
}
