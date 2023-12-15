package others

type DLinkedNode struct {
	prev  *DLinkedNode
	next  *DLinkedNode
	key   int
	value int
}

type LRUCache struct {
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
	size     int
	capacity int
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func New(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}

	node := l.cache[key]
	// move node to the head of linked list
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) Put(key, value int) {
	if _, ok := l.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		l.cache[key] = node
		// add node to the head of linked list
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			// delete the tail node from the linked list and the cache
			removedNode := l.removeTail()
			delete(l.cache, removedNode.key)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = value
		// move node to the head of linked list
		l.moveToHead(node)
	}
}

func (l *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
