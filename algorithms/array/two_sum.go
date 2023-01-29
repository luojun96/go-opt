package main

// nums = [2,7,11,15], target = 9 => output: [0,1]
// for i:=0; i < len(nums); i++ {
// 		for j := i + 1; i < len(nums); i++ {
// 			if nums[i] + nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// }
// 复杂度分析：时间复杂度O(n^2), 空间复杂度O(1)
// use hashmap
// m := map[int][int]{}
// for i:=0; i < len(nums); i++ {
//		m[nums[i]] = i
// }
// then loop nums, try to find the element in which the index is: m[target - nums[i]]
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		if v, ok := m[other]; ok && v != i {
			return []int{i, v}
		}
	}
	return []int{-1, -1}
}
