package main

type Triangle struct {
	Side1, Side2, Side3 int
}

func (t Triangle) accept(v Visitor) {
	v(t)
}
