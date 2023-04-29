package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type Once struct {
	sync.Once
}

func (o *Once) Done() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.Once))) == 1
}

func main() {
	var once Once
	fmt.Println(once.Done()) // false
	once.Do(func() {
		time.Sleep(time.Second)
		println("Hello, World!")
	})
	fmt.Println(once.Done()) // true
}
