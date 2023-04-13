package algorithms

import (
	"math/rand"
	"time"
)

// https://leetcode.cn/problems/find-the-kth-largest-integer-in-the-array/description/?languageTags=golang
// input: nums = ["3","6","7","10"], k = 4
// output: "3"
func findkthLargestNumber(nums []string, k int) string {
	rand.Seed(time.Now().UnixNano())
	return doQuickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func doQuickSelect(a []string, l, r, index int) string {
	q := DoRandomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return doQuickSelect(a, q+1, r, index)
	}
	return doQuickSelect(a, l, q-1, index)
}

func DoRandomPartition(a []string, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return DoPartition(a, l, r)
}

func DoPartition(a []string, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if cmp(x, a[j]) >= 0 {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func cmp(a, b string) int {
	lens := len(a)
	if lens > len(b) {
		return 1
	} else if lens < len(b) {
		return -1
	}
	for i := 0; i < lens; i++ {
		if a[i] != b[i] {
			return int(a[i]) - int(b[i])
		}
	}
	return 0
}
