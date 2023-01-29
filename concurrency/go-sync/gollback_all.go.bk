package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vardius/gollback"
)

func main() {
	rs, errs := gollback.All(
		context.Background(),
		func(ctx context.Context) (interface{}, error) {
			time.Sleep(3 * time.Second)
			return 1, nil
		},
		func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("failed")
		},
		func(ctx context.Context) (interface{}, error) {
			return 3, nil
		},
	)

	fmt.Println(rs)
	fmt.Println(errs)
}
