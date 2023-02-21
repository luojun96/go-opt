package main

// https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/
// input: [0,1,3]
// output: 2
func findMissingNumber(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	for i := 0; ; i++ {
		if !m[i] {
			return i
		}
	}
}
