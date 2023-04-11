package main

// https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
// input: [1, 2, 3, 2, 2, 2, 5, 4, 2]
// output: 2
func findMajorityElement(nums []int) int {
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

func majorityElementByHash(nums []int) int {
	m := make(map[int]int)
	var countNums func()
	countNums = func() {
		for _, num := range nums {
			if count, ok := m[num]; !ok {
				m[num] = 1
			} else {
				count++
				m[num] = count
			}
		}
	}
	countNums()
	res, count := nums[0], 1
	for n, c := range m {
		if c > count {
			res = n
			count = c
		}
	}
	return res
}
