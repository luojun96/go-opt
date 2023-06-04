package main

import (
	"fmt"
	"time"
)

func doRun() {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		fmt.Println("loop", i)
		select {
		case ch <- i:
			fmt.Println("send", i)
		case v := <-ch:
			fmt.Println("receive", v)
		}
	}
	time.Sleep(60 * time.Second)
	// Execute()
	// chArr := [4]chan struct{}{
	// 	make(chan struct{}),
	// 	make(chan struct{}),
	// 	make(chan struct{}),
	// 	make(chan struct{}),
	// }

	// for i := 0; i < 4; i++ {
	// 	go func(i int) {
	// 		for {
	// 			<-chArr[i%4]
	// 			fmt.Printf("I'm %d\n", i+1)
	// 			time.Sleep(1 * time.Second)
	// 			chArr[(i+1)%4] <- struct{}{}
	// 		}
	// 	}(i)
	// }

	// chArr[0] <- struct{}{}
	// select {}
}
