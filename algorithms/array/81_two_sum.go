package algorithms

// https://leetcode.cn/problems/two-sum/
// input: nums = [2,7,11,15], target = 9
// output: [0,1]

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, k}
		}
		m[v] = k
	}
	return nil
}
