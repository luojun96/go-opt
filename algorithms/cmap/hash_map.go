package cmap

type MyHashMap struct {
	hashMap map[int]int
}

func NewMyHashMap() MyHashMap {
	return MyHashMap{make(map[int]int)}
}

func (m *MyHashMap) Put(key int, value int) {
	m.hashMap[key] = value
}

func (m *MyHashMap) Get(key int) int {
	if value, ok := m.hashMap[key]; ok {
		return value
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	delete(m.hashMap, key)
}
