package heap

import (
	"container/heap"
)

// https://leetcode.cn/problems/top-k-frequent-elements/description/
// since this is a topk issue, we can use priority queue or quick select sort to handle it. But the key to use isn't value, but the frequent of an element
// then we need to get the frequent count of eah element in the array
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequentByHeap(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}

	h := &IHeap{}
	heap.Init(h)

	for i, v := range m {
		if h.Len() < k {
			heap.Push(h, [2]int{i, v})
		} else {
			top := heap.Pop(h)
			if v > top.([2]int)[1] {
				heap.Push(h, [2]int{i, v})
			} else {
				heap.Push(h, top)
			}
		}
	}
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).([2]int)[0])
	}
	return res
}
