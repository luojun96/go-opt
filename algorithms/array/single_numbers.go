package main

func singleNumbers(nums []int) []int {
	res := []int{}
	hash := map[int]bool{}
	for i := 0; i < len(nums) - 1; i++ {
		_, ok := hash[i]
		if !ok {
			for j := i + 1; j < len(nums); j++ {
				if _, okj := hash[j]; okj && hash[j] {
					continue
				}
				if nums[i] == nums[j] {

				}
			}
		}
	}
	return res
}
