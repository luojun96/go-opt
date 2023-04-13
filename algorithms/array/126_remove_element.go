package algorithms

// https://leetcode.cn/problems/remove-element/
// input: nums = [0,1,2,2,3,0,4,2], val = 2
// output: 5, nums = [0,1,4,0,3]
func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
