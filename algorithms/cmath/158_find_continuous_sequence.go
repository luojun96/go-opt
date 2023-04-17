package cmath

// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/solutions/
// input: target = 9
// output: [[2,3,4],[4,5]]
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	for l, r := 1, 2; l < r; {
		sum := ((l + r) * (r - l + 1)) >> 1
		if sum == target {
			arr := make([]int, r-l+1)
			for i := l; i <= r; i++ {
				arr[i-l] = i
			}
			res = append(res, arr)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return res
}
