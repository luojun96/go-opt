package array

// https://leetcode.cn/problems/design-circular-deque/
type node struct {
	prev, next *node
	val        int
}

type MyCircularDeque struct {
	head, tail     *node
	capacity, size int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{capacity: k}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}

	node := &node{val: value}
	if q.IsEmpty() {
		q.head = node
		q.tail = node
	} else {
		node.next = q.head
		q.head.prev = node
		q.head = node
	}

	q.size++
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}

	node := &node{val: value}
	if q.IsEmpty() {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		node.prev = q.tail
		q.tail = node
	}
	q.size++
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	q.head = q.head.next
	if q.head != nil {
		q.head.prev = nil
	}
	q.size--
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}

	q.tail = q.tail.prev
	if q.tail != nil {
		q.tail.next = nil
	}
	q.size--
	return true
}

func (q *MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.head.val
}

func (q *MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.tail.val
}

func (q *MyCircularDeque) IsEmpty() bool {
	return q.size == 0
}

func (q *MyCircularDeque) IsFull() bool {
	return q.size == q.capacity
}
