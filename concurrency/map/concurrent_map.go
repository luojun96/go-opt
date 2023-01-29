package main

import (
	"fmt"
	"sync"
)

const SHARD_COUNT = 32

type ConcurrentMap []*ConcurrentMapShared

type ConcurrentMapShared struct {
	sync.RWMutex
	items map[string]interface{}
}

func New() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)

	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]interface{})}
	}
	return m
}

func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

func (m ConcurrentMap) Set(key string, value interface{}) {
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	shard := m.GetShard(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}
func main() {
	fmt.Println("vim-go")
}
