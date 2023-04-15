package heap

import (
	"fmt"
)

func findKthLargest01(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := partition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func partition(a []int, l, r int) int {
	pivot := a[r]
	i := l
	for j := l; j < r; j++ {
		if a[j] <= pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}

func heap() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	res := findKthLargest01(nums, k)
	fmt.Println(res)
}
