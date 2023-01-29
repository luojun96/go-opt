package main

type IntSetWithUndo struct {
	data map[int]bool
	undo Undo
}

func NewIntSetWithUndo() IntSetWithUndo {
	return IntSetWithUndo{data: make(map[int]bool)}
}

func (set *IntSetWithUndo) Undo() error {
	return set.undo.Undo()
}

func (set *IntSetWithUndo) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSetWithUndo) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSetWithUndo) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}
