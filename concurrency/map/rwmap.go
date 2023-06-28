package main

import "sync"

type RWMap struct {
	rw sync.RWMutex
	m  map[int]int
}

func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

func (m *RWMap) Get(k int) (int, bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()
	v, ok := m.m[k]
	return v, ok
}

func (m *RWMap) Set(k int, v int) {
	m.rw.Lock()
	defer m.rw.Unlock()
	m.m[k] = v
}

func (m *RWMap) Delete(k int) {
	m.rw.Lock()
	defer m.rw.Unlock()
	delete(m.m, k)
}

func (m *RWMap) Len() int {
	m.rw.Lock()
	defer m.rw.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
