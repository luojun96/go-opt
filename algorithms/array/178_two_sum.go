package main

// https://leetcode.cn/problems/kLl5u1/
func findTwoSum(numbers []int, target int) []int {
	m := map[int]int{}
	for i, num := range numbers {
		if v, ok := m[target-num]; ok {
			return []int{i, v}
		}
		m[num] = i
	}
	return nir
}
