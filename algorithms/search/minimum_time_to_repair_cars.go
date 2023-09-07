package search

import "math"

// https://leetcode.cn/problems/minimum-time-to-repair-cars/?envType=daily-question&envId=2023-09-07
func repairCars(ranks []int, cars int) int64 {
	l, r := 1, ranks[0]*cars*cars
	var check func(int) bool
	check = func(m int) bool {
		cnt := 0
		for _, rank := range ranks {
			cnt += int(math.Sqrt(float64(m / rank)))
		}
		return cnt >= cars
	}

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return int64(l)
}
