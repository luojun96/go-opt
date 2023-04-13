package algorithms

import (
	"fmt"
)

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		cache:    map[int]*DLinkedNode{},
	}

	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (t *LRUCache) Get(key int) int {
	if _, ok := t.cache[key]; !ok {
		return -1
	}
	node := t.cache[key]
	// move to list head
	t.moveToHead(node)
	return node.value
}

func (t *LRUCache) Put(key int, value int) {
	if _, ok := t.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		t.cache[key] = node
		// move to list head
		t.addToHead(node)
		t.size++
		if t.size > t.capacity {
			// remove the list tail
			r := t.removeTail()
			delete(t.cache, r.key)
			t.size--
		}
	} else {
		node := t.cache[key]
		node.value = value
		// move to list head
		t.moveToHead(node)
	}
}

func (t *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = t.head
	node.next = t.head.next
	t.head.next.prev = node
	t.head.next = node
}

func (t *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	t.tail.prev = node.prev
}

func (t *LRUCache) moveToHead(node *DLinkedNode) {
	t.removeNode(node)
	t.addToHead(node)
}

func (t *LRUCache) removeTail() *DLinkedNode {
	node := t.tail.prev
	t.removeNode(node)
	return node
}

func algorithms() {
	obj := Constructor(3)
	obj.Put(1, 2)
	obj.Put(2, 3)
	obj.Put(3, 4)

	fmt.Println(obj.Get(1))
}
