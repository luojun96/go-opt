package algorithms

// https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/
type MaxQueue struct {
	queue   []int
	maxList []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue:   []int{},
		maxList: []int{},
	}
}

func (m *MaxQueue) Max_value() int {
	if len(m.queue) == 0 {
		return -1
	}
	return m.maxList[0]
}

func (m *MaxQueue) Push_back(value int) {
	m.queue = append(m.queue, value)
	for len(m.maxList) > 0 && m.maxList[len(m.maxList)-1] < value {
		m.maxList = m.maxList[:len(m.maxList)-1]
	}
	m.maxList = append(m.maxList, value)
}

func (m *MaxQueue) Pop_front() int {
	if len(m.queue) == 0 {
		return -1
	}
	res := m.queue[0]
	m.queue = m.queue[1:]
	if res == m.maxList[0] {
		m.maxList = m.maxList[1:]
	}

	return res
}
