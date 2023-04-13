package algorithms

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
// input: nums = [4,1,4,6]
// output: [1,6] or [6,1]
func singleNumbers(nums []int) []int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}
	div := 1
	for (div & ret) == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for _, num := range nums {
		if (div & num) != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof
func singleNumbersByTraverse(nums []int) int {
	res := -1
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if count, ok := m[nums[i]]; !ok {
			m[nums[i]] = 1
		} else {
			m[nums[i]] = count + 1
		}
	}

	for k, v := range m {
		if v == 1 {
			res = k
			break
		}
	}
	return res
}
