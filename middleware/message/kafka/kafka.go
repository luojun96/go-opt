package main

import "context"

type Kafka struct {
	network string
	address string
	context context.Context
}

func New(address string) *Kafka {
	return &Kafka{
		network: "tcp",
		address: address,
		context: context.Background(),
	}
}
