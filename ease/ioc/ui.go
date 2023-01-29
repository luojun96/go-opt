package main

type Widget struct {
	X, Y int
}

type Label struct {
	Widget
	Text string
}

type Button struct {
	Label
}

type ListBox struct {
	Widget
	Texts []string
	Index int
}
