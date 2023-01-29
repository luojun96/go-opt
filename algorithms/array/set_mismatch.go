package main

// https://leetcode.cn/problems/set-mismatch/
// input: nums = [1,2,2,4]
// output: [2,3]
func findErrorNums(nums []int) []int {
	n := len(nums)
	dup := -1
	for i := 0; i < n; i++ {
		index := absInt(nums[i]) - 1
		if nums[index] < 0 {
			dup = absInt(nums[i])
		} else {
			nums[index] *= -1
		}
	}

	missing := -1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			missing = i + 1
		}
	}
	return []int{dup, missing}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
