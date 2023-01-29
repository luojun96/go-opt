package main

import "testing"

func TestLinkList(t *testing.T) {
	var l = list[int]{}
	l.add(1)
	l.add(2)
	l.push(3)
	l.push(4)
	l.add(5)
	l.print()

	l.del(5)
	l.del(1)
	l.del(4)
	l.print()
}
