package stack

// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
type MinStack struct {
	elements []int
	mins     []int
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	m.elements = append(m.elements, x)
	if len(m.mins) == 0 {
		m.mins = append(m.mins, x)
	} else {
		min := m.Min()
		if x < min {
			m.mins = append(m.mins, x)
		} else {
			m.mins = append(m.mins, min)
		}
	}
}

func (m *MinStack) Pop() {
	if len(m.elements) == 0 {
		panic("empty stack")
	}
	m.elements = m.elements[:len(m.elements)-1]
	m.mins = m.mins[:len(m.mins)-1]
}

func (m *MinStack) Top() int {
	n := len(m.elements)
	if n == 0 {
		panic("empty stack")
	}
	return m.elements[n-1]
}

func (m *MinStack) Min() int {
	n := len(m.mins)
	return m.mins[n-1]
}
