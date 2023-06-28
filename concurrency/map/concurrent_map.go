package main

import (
	"strconv"
	"sync"
)

const ShardCount = 32

type ConcurrentMap []*ConcurrentMapShared

type ConcurrentMapShared struct {
	rw    sync.RWMutex
	items map[string]interface{}
}

func New() ConcurrentMap {
	m := make(ConcurrentMap, ShardCount)

	for i := 0; i < ShardCount; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]interface{})}
	}
	return m
}

func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShared {
	i, _ := strconv.Atoi(key)
	return m[uint(i)%uint(ShardCount)]
}

func (m ConcurrentMap) Set(key string, value interface{}) {
	shard := m.GetShard(key)
	shard.rw.Lock()
	shard.items[key] = value
	shard.rw.Unlock()
}

func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	shard := m.GetShard(key)
	shard.rw.RLock()
	val, ok := shard.items[key]
	shard.rw.RUnlock()
	return val, ok
}
