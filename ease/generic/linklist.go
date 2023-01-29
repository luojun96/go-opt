package main

import "fmt"

type node[T comparable] struct {
	data T
	prev *node[T]
	next *node[T]
}

type list[T comparable] struct {
	head, tail *node[T]
	len        int
}

func (l *list[T]) isEmpty() bool {
	return l.head == nil && l.tail == nil
}

func (l *list[T]) add(data T) {
	n := &node[T]{
		data: data,
		prev: nil,
		next: l.head,
	}
	if l.isEmpty() {
		l.head = n
		l.tail = n
	}
	l.head.prev = n
	l.head = n
}

func (l *list[T]) push(data T) {
	n := &node[T]{
		data: data,
		prev: l.tail,
		next: nil,
	}
	if l.isEmpty() {
		l.head = n
		l.tail = n
	}

	l.tail.next = n
	l.tail = n
}

func (l *list[T]) del(data T) {
	for p := l.head; p != nil; p = p.next {
		if data == p.data {
			if p == l.head {
				l.head = p.next
			}
			if p == l.tail {
				l.tail = p.prev
			}
			if p.prev != nil {
				p.prev.next = p.next
			}
			if p.next != nil {
				p.next.prev = p.prev
			}
			return
		}
	}
}

func (l *list[T]) print() {
	if l.isEmpty() {
		fmt.Println("the link list is empty")
		return
	}
	for p := l.head; p != nil; p = p.next {
		fmt.Printf("[%v] -> ", p.data)
	}
	fmt.Println("nil")
}
