package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Printf("failed to create trace file: %v ", err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		fmt.Printf("failed to start tracing: %v", err)
	}
	defer trace.Stop()

	fmt.Println("jun")
}
