package heap

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/chou-shu-lcof/
// input: n = 10
// output: 12, "1, 2, 3, 4, 5, 6, 8, 9, 10, 12"
var factors = []int{2, 3, 5}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func nthUglyNumber(n int) int {
	h := &hp{
		sort.IntSlice{1},
	}
	seen := map[int]struct{}{1: {}}
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, ok := seen[next]; !ok {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}

func nthUglyNumberByDP(n int) int {
	var min func(i, j int) int
	min = func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}
