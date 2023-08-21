package array

// https://leetcode.cn/problems/maximize-distance-to-closest-person/
func maxDistToClosest(seats []int) int {
	res := 0
	l := 0

	for l < len(seats) && seats[l] == 0 {
		l++
	}
	res = max(res, l)

	for l < len(seats) {
		r := l + 1
		for r < len(seats) && seats[r] == 0 {
			r++
		}
		if r == len(seats) {
			res = max(res, r-l-1)
		} else {
			res = max(res, (r-l)/2)
		}
		l = r
	}

	return res
}
