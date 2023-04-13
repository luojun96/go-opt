package algorithms

// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
type MinStack struct {
	Elements []int
	MinS     []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	m.Elements = append(m.Elements, x)
	if len(m.MinS) == 0 {
		m.MinS = append(m.MinStack, x)
	} else {
		min := m.Min()
		if x < min {
			m.MinS = append(m.MinS, x)
		} else {
			m.MinS = append(m.MinS, min)
		}
	}
}

func (m *MinStack) Pop() {
	if len(m.Elements) == 0 {
		panic("empty stack")
	}
	m.Elements = m.Elements[:len(m.Elements)-1]
	m.MinS = m.MinS[:len(m.MinS)-1]
}

func (m *MinStack) Top() int {
	n := len(m.Elements)
	if n == 0 {
		panic("empty stack")
	}
	return m.Elements[n-1]
}

func (m *MinStack) Min() int {
	n := len(m.MinS)
	return m.MinS[n-1]
}
