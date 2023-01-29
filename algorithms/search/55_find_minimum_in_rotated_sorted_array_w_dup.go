package main

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/description/
// input: nums = [2,2,3,0,1]
// output: 0
func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	for l < r {
		mid := l + (r - l)>>1
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r--
		}
	}
	return nums[l]
}
