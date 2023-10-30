package array

import "sort"

// https://leetcode.cn/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/description/?envType=daily-question&envId=2023-10-27
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	return calMax(horizontalCuts, h) * calMax(verticalCuts, w)
}

func calMax(a []int, board int) int {
	res, pre := 0, 0
	for _, v := range a {
		res = max(v-pre, res)
		pre = v
	}
	return max(res, board-pre)
}
