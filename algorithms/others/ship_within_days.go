package others

import "sort"

func shipWithinDays(weights []int, days int) int {
	l, r := 0, 0
	for _, w := range weights {
		if w > l {
			l = w
		}
		r += w
	}

	return l + sort.Search(r-l, func(x int) bool {
		x += l
		day := 1
		sum := 0
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= days
	})
}
