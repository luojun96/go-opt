package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.cn/problems/top-k-frequent-elements/
// input: nums = [1,1,1,2,2,3], k = 2
// output: [1,2]
// since this is a topk issue, we can use priority queue or quick select sort to handle it. But the key to use isn't value, but the frequent count of an element
// Then we need get the frequent count of each element in the array
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
H
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	fmt.Printf("k = %d\n", k)
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}
	fmt.Printf("m: %v\n", m)
	h := &IHeap{}
	heap.Init(h)
	for i, v := range m {
		fmt.Printf("heap length: %d\n", h.Len())
		if h.Len() < k {
			heap.Push(h, [2]int{i, v})
			fmt.Printf("the heap isn't full, then push [%d: %d]\n", i, v)
		} else {
			top := heap.Pop(h)
			fmt.Printf("Pop the top [%d: %d]\n", top.([2]int)[0], top.([2]int)[1])
			if v > top.([2]int)[1] {
				heap.Push(h, [2]int{i, v})
				fmt.Printf("push [%d: %d] into heap\n", i, v)
			} else {
				heap.Push(h, top)
				fmt.Printf("push top [%d: %d] back to heap\n", top.([2]int)[0], top.([2]int)[1])
			}
		}
	}
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).([2]int)[0])
	}
	return res
}
