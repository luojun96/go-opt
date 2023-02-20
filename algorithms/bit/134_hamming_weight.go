package main

// https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof
// input: 11
// output: 3
func hammingWeight(num uint32) int {
	var res int
	for ; num > 0; num &= num - 1 {
		res++
	}
	return res
}
