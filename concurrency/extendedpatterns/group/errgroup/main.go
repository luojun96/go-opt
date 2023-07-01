package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	fmt.Println("execute three tasks concurrently, and exit when one of them fails")
	exitWhenError()

	fmt.Println("execute three tasks concurrently, and exit when all tasks are done, and print all errors")
	exitWithAllErrors()

	fmt.Println("execute tasks concurrently, and exit when timeout")
	withContext()
}

func exitWhenError() {
	var g errgroup.Group

	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(6 * time.Second)
		fmt.Println("exec #2")
		return errors.New("failed to exec #2")
	})

	g.Go(func() error {
		time.Sleep(8 * time.Second)
		fmt.Println("exec #3")
		return nil
	})

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Printf("failed to exec: %v\n", err)
	}
}

func exitWithAllErrors() {
	var g errgroup.Group
	results := make([]error, 3)

	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		results[0] = nil
		return nil
	})

	g.Go(func() error {
		time.Sleep(4 * time.Second)
		fmt.Println("exec #2")
		results[1] = errors.New("failed to exec #2")
		return nil
	})

	g.Go(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("exec #3")
		results[2] = errors.New("failed to exec #3")
		return nil
	})

	if err := g.Wait(); err == nil {
		fmt.Println("all tasks are done")
	} else {
		fmt.Printf("failed to exec: %v\n", err)
	}

	fmt.Println("all errors:")
	for _, result := range results {
		if result != nil {
			fmt.Printf("error: %v\n", result)
		}
	}
}

type Result string
type Search func(ctx context.Context, query string) (Result, error)

func fakeSearch(kind string) Search {
	return func(ctx context.Context, query string) (Result, error) {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Printf("timeout: %s\n", kind)
			return "", ctx.Err()
		default:
		}

		fmt.Printf("search: %s\n", kind)
		return Result(fmt.Sprintf("%s result for %q.", kind, query)), nil
	}
}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func withContext() {
	Google := func(ctx context.Context, query string) ([]Result, error) {
		g, ctx := errgroup.WithContext(ctx)
		searches := []Search{Web, Image, Video}
		results := make([]Result, len(searches))
		for i, search := range searches {
			i, search := i, search // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				result, err := search(ctx, query)
				if err == nil {
					results[i] = result
				}
				results[i] = result
				return err
			})
		}

		if err := g.Wait(); err != nil {
			return nil, err
		}
		return results, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	results, err := Google(ctx, "golang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
