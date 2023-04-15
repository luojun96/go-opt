package array

// https://leetcode.cn/problems/majority-element/
// input: nums = [3,2,3]
// output: 3
func majorityElement(nums []int) int {
	num, count := nums[0], 1
	for _, n := range nums {
		if num == n {
			count++
		} else {
			count--
			if count == 0 {
				num = n
				count = 1
			}
		}
	}
	return num
}
