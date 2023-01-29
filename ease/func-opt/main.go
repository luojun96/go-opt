package main

import (
	"time"
)

func main() {
	NewServer("localhost", 9000, nil)

	conf := Config{
		Protocol: "tcp",
		Timeout:  time.Duration(60),
	}
	NewServer("localhost", 9000, &conf)

	NewServerWithOptions("localhost", 1024)
	NewServerWithOptions("localhost", 2048, Protocol("udp"))
	NewServerWithOptions("localhost", 8080, Timeout(300*time.Second), MaxConns(10000))
}
