package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	capacity      = 10
	producerCount = 10
	consumerCount = 20
)

type Product struct {
	Name string
	Id   int
}

func Produce(ctx context.Context, id int, q chan<- Product) {
	log.Printf("Producer %d starts to work...", id)
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Printf("Producer %d is interrupted", id)
			return
		case <-ticker.C:
			number := rand.Intn(1000)
			product := Product{
				Name: fmt.Sprintf("Product-%d", number),
				Id:   number,
			}
			q <- product
			log.Printf("Producer %d produces %s\n", id, product.Name)
		}
	}
}

func Consume(ctx context.Context, id int, q <-chan Product) {
	log.Printf("Consumer %d comes in...", id)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Consumer %d is interrupted...", id)
			return
		case product, ok := <-q:
			if ok {
				time.Sleep(5 * time.Second)
				log.Printf("Consumer %d buys %s", id, product.Name)
			}
		}
	}
}

func main() {
	// queue, capacity = 10
	q := make(chan Product, capacity)
	defer close(q)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start produce
	for i := 0; i < producerCount; i++ {
		go Produce(ctx, i, q)
	}

	// start consume
	for i := 0; i < consumerCount; i++ {
		go Consume(ctx, i, q)
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	log.Println("Ask consumers to leave and producers to stop working...")
}
