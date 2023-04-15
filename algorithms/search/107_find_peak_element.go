package search

import "math"

// https://leetcode.cn/problems/find-peak-element/
// input: nums = [1,2,3,1]
// output: 2
func findPeakElement(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1
	for {
		mid := left + (right-left)>>1
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func findPeakElementByIteration(nums []int) (res int) {
	for i, v := range nums {
		if v > nums[res] {
			res = i
		}
	}
	return
}
