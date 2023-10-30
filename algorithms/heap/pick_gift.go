package heap

import (
	"container/heap"
	"math"
)

// https://leetcode.cn/problems/take-gifts-from-the-richest-pile/?envType=daily-question&envId=2023-10-28
type highHeap []int

// Len is the number of elements in the collection.
func (h highHeap) Len() int {
	return len(h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (h highHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

// Swap swaps the elements with indexes i and j.
func (h highHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *highHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *highHeap) Pop() any {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

func pickGifts(gifts []int, k int) int64 {
	h := make(highHeap, 0, len(gifts))
	for _, gift := range gifts {
		heap.Push(&h, gift)
	}

	for k > 0 {
		g := heap.Pop(&h).(int)
		if g == 1 {
			heap.Push(&h, g)
			break
		}
		ng := int(math.Sqrt(float64(g)))
		heap.Push(&h, ng)
		k--
	}

	res := 0
	for h.Len() > 0 {
		res += heap.Pop(&h).(int)
	}
	return int64(res)
}
