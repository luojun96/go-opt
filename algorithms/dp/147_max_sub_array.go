package main

// https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
// input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// output: 6
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
