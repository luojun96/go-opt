package main

// https://leetcode.cn/problems/subarray-sum-equals-k/
// exp: nums = [1,1,1], k = 2 => 2 ([1,1], [1,1])
// 		nums = [1,2,3], k = 3 => 2 ([1,2], [3])
// 利用前缀和散列表， pre[i] = pre[i-1] + num[i], sub array [i,j], pre[i] - pre[j-1] = k
// use hash map to store the count of each pre sum: map[presum]count
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{
		0: 1,
	}
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}
