package search

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/description/
// input: nums = [2,2,3,0,1]
// output: 0
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return nums[low]
}
