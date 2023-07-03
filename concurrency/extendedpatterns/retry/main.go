package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vardius/gollback"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := gollback.Retry(ctx, 5, func(ctx context.Context) (interface{}, error) {
		fmt.Println("try")
		return nil, errors.New("failed")
	})

	fmt.Println(res)
	fmt.Println(err)
}
