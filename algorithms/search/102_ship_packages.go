package algorithms

import "sort"

// https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/
// input: weights= [1,2,3,4,5,6,7,8,9,10], days = 5
// output: 15
func shipWithinDays(weights []int, days int) int {
	l, r := 0, 0
	for _, w := range weights {
		if  w > l {
			l = w
		}
		r += w
	}

	return l + sort.Search(r-l, func(x int) bool {
		x += l
		day := 1
		sum := 0
		for _, w := range weights {
			if sum + w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= days
	})
}
