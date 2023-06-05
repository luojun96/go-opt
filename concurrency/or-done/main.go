package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// https://chat.forefront.ai/share/vyxp5ih7vc47408b
func task(ctx context.Context, id int) {
	duration := time.Duration(rand.Intn(5)+1) * time.Second
	select {
	case <-time.After(duration):
		fmt.Printf("Task %d completed in %v\n", id, duration)
	case <-ctx.Done():
		fmt.Printf("Task %d canceled\n", id)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithCancel(context.Background())

	// 启动 5 个并发任务
	for i := 1; i <= 5; i++ {
		go task(ctx, i)
	}

	// 等待一个任务完成
	time.Sleep(3 * time.Second)

	// 取消其他任务
	cancel()

	// 等待一段时间，以便观察任务取消的效果
	time.Sleep(5 * time.Second)
}
