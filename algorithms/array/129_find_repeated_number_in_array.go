package array

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/description/
// input: [2, 3, 1, 0, 2, 5, 3]
// output: 2 or 3
func findRepeatNumber(nums []int) int {
	res := -1
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		_, ok := m[num]
		if !ok {
			m[num] = true
		} else {
			res = num
			break
		}
	}
	return res
}
