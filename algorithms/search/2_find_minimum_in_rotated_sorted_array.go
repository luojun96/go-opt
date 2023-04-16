package search

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
// input: nums = [3,4,5,1,2]
// output: 1

func findMinInRotatedArray(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
