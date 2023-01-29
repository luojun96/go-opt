package main

import (
	"container/list"
	"fmt"
)

type ExamRoom struct {
	seat *list.List
	n    int
}

func Constructor(N int) ExamRoom {
	return ExamRoom{
		seat: list.New(),
		n:    N - 1,
	}
}

func (e *ExamRoom) Seat() int {
	if e.seat.Len() == 0 {
		e.seat.PushFront(0)
		return 0
	}

	element := e.seat.Front()
	prev := element.Value.(int)
	max := prev
	addVal := 0
	addFront := element
	for ; element != nil; element = element.Next() {
		value := element.Value.(int)
		distant := (value - prev) / 2
		if distant > max {
			max = distant
			addFront = element
			addVal = prev + distant
		}
		prev = value
	}
	distant := e.n - prev
	if distant > max {
		e.seat.PushBack(e.n)
		return e.n
	}
	e.seat.InsertBefore(addVal, addFront)
	return addVal

}

func (e *ExamRoom) Leave(p int) {
	for element := e.seat.Front(); element != nil; element = e.Next() {
		if element.Value.(int) == p {
			e.seat.Remove(element)
			return
		}
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
