package queue

// https://leetcode.cn/problems/design-circular-queue/

type MyCircularQueue struct {
	Items []int
	Size  int
	Head  int
	Tail  int
}

func New(k int) MyCircularQueue {
	items := make([]int, k+1)
	return MyCircularQueue{
		Items: items,
		Size:  k + 1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Items[this.Tail] = value
	this.Tail = (this.Tail + 1) % this.Size
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = (this.Head + 1) % this.Size
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Items[this.Head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Items[(this.Tail - 1 + this.Size)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == this.Tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.Tail+1)%this.Size == this.Head
}
