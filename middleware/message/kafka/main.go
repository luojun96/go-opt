package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// kafka client configuration
type Kafka struct {
	network string
	address string
	context context.Context
}

func main() {
	topic := "test"
	partition := 0

	k := &Kafka{
		network: "tcp",
		address: "localhost:9092",
		context: context.Background(),
	}

	conn, err := kafka.DialLeader(k.context, k.network, k.address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one")},
		kafka.Message{Value: []byte("two")},
		kafka.Message{Value: []byte("three")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
