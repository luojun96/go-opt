package heap

import (
	"container/heap"
)

// https://leetcode.cn/problems/kth-largest-element-in-an-array/
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
func findKthLargest(nums []int, k int) int {
	h := make(intHeap, 0, k)
	heap.Init(&h)

	for _, v := range nums {
		if len(h) < k {
			heap.Push(&h, v)
		} else {
			top := (heap.Pop(&h)).(int)
			if top < v {
				heap.Push(&h, v)
			} else {
				heap.Push(&h, top)
			}
		}
	}
	return (heap.Pop(&h)).(int)
}

// func heap() {
// 	nums := []int{1, 5, 2, 6, 3}
// 	k := 3
// 	res := findKthLargest(nums, k)
// 	fmt.Println(res)
// }
