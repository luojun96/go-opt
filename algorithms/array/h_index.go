package array

import "sort"

// https://leetcode.cn/problems/h-index/?envType=daily-question&envId=2023-10-29
func hIndex(citations []int) int {
	h := 0
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] > h {
			h++
		}
	}
	return h
}

// https://leetcode.cn/problems/h-index-ii/description/?envType=daily-question&envId=2023-10-30
func hIndex2(citations []int) int {
	l, r := 0, len(citations)-1
	for l <= r {
		m := l + (r-l)>>1
		if citations[m] >= len(citations)-m {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return len(citations) - l
}

func hIndex3(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(x int) bool {
		return citations[x] >= n-x
	})
}
