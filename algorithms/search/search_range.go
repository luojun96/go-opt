package search

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
// use binary search to handle this issue.
func searchRange(nums []int, target int) []int {
	// first := bSearchFirstEqual(nums, target)
	// last := bSearchLastEqual(nums, target)

	first := bSearchEqual(nums, target, true)
	last := bSearchEqual(nums, target, false)
	return []int{first, last}
}

func bSearchEqual(a []int, v int, isFirst bool) int {
	l, h := 0, len(a)-1
	for l <= h {
		m := l + (h-l)>>1
		if a[m] > v {
			h = m - 1
		} else if a[m] < v {
			l = m + 1
		} else {
			if isFirst {
				if m == 0 || a[m-1] != v {
					return m
				} else {
					h = m - 1
				}
			} else {
				if m == len(a)-1 || a[m+1] != v {
					return m
				} else {
					l = m + 1
				}
			}
		}
	}
	return -1
}
func bSearchFirstEqual(a []int, v int) int {
	l, h := 0, len(a)-1
	for l <= h {
		m := l + (h-l)>>1
		if a[m] > v {
			h = m - 1
		} else if a[m] < v {
			l = m + 1
		} else {
			if m == 0 || a[m-1] != v {
				return m
			} else {
				h = m - 1
			}
		}
	}
	return -1
}

func bSearchLastEqual(a []int, v int) int {
	l, h := 0, len(a)-1
	for l <= h {
		m := l + (h-l)>>1
		if a[m] > v {
			h = m - 1
		} else if a[m] < v {
			l = m + 1
		} else {
			if m == len(a)-1 || a[m+1] != v {
				return m
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
