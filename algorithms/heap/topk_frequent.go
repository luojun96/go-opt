package heap

import "container/heap"

type Item struct {
	Value     int
	Frequency int
}

type FrequencyHeap []*Item

// Len is the number of elements in the collection.
func (fh *FrequencyHeap) Len() int {
	return len(*fh)
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
func (fh *FrequencyHeap) Less(i int, j int) bool {
	return (*fh)[i].Frequency < (*fh)[j].Frequency
}

// Swap swaps the elements with indexes i and j.
func (fh *FrequencyHeap) Swap(i int, j int) {
	(*fh)[i], (*fh)[j] = (*fh)[j], (*fh)[i]
}

func (fh *FrequencyHeap) Push(x any) {
	item := x.(*Item)
	*fh = append(*fh, item)
}

func (fh *FrequencyHeap) Pop() any {
	old := *fh
	n := fh.Len()
	item := (*fh)[n-1]
	old[n-1] = nil
	*fh = old[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	items := make(map[int]int, len(nums))
	for _, v := range nums {
		items[v]++
	}

	fh := FrequencyHeap{}
	heap.Init(&fh)

	for value, frequency := range items {
		if fh.Len() < k {
			heap.Push(&fh, &Item{
				Value:     value,
				Frequency: frequency,
			})
		} else {
			top := heap.Pop(&fh).(*Item)
			if frequency > top.Frequency {
				heap.Push(&fh, &Item{
					Value:     value,
					Frequency: frequency,
				})
			} else {
				heap.Push(&fh, top)
			}
		}
	}

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(&fh).(*Item)
		res = append(res, item.Value)
	}
	return res
}
