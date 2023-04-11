package main

// https://leetcode.cn/problems/qIsx9U/

type MovingAverage struct {
	size int
	sum  int
	q    []int
}

func NewMovingAverage(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (m *MovingAverage) Next(val int) float64 {
	if len(m.q) == m.size {
		m.sum -= m.q[0]
		m.q = m.q[1:]
	}

	m.sum += val
	m.q = append(m.q, val)
	return float64(m.sum) / float64(len(m.q))
}
