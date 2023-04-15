package array

// https://leetcode.cn/problems/WGki4K/description/
func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	for num, count := range m {
		if count == 1 {
			return num
		}
	}
	return 0
}
