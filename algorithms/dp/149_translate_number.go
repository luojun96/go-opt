package algorithms

import (
	"strconv"
)

// https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
// input: 12258
// output: 5
func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}
