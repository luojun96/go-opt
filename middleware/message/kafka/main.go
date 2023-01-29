package main

import (
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	address := "localhost:9092"
	topic := "test"
	partition := 0

	k := New(address)
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
