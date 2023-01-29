package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// wait group
type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func WaitGroup() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go worker(&counter, &wg)
	}

	wg.Wait()

	fmt.Println(counter.Count())
}

// or-done
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	default:
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()

	return orDone
}

func orReflect(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	default:
	}

	orDone := make(chan interface{})
	go func() {
		defer close(OrDone)

		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		reflect.Select(cases)
	}()

	return orDone
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func OrDone() {
	start := time.Now()

	<-or(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(01*time.Second),
		sig(50*time.Second),
	)

	// fmt.Printf("done after %v", time.Since(start))
}

// fan-in
func faninReflect(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir: reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})	
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)	
			if !ok {
				cases = append(cases[:i], cases[i+1]...)
				continue
			}
			out <- v.Interface()
	}()

	return out
}
// fan-out
func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])	
			}
		}
		for v := range ch {
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async {
					go func() {
						out[i] <- v	
					}()
				} else {
					out[i] <- v
				}
			}
		}
	}()

}
// stream
func asStream(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	s := make(chan interface{})
	go func() {
		defer close(s)
		for _, v := range values {
			select {
			case <-done:
				return	
			case s <- v:
			}
		}
	}()
	return s
}
