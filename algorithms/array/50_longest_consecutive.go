package array

// https://leetcode.cn/problems/WhsWhI/
// input: nums = [100,4,200,1,3,2]
// output: 4
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0
	for num, _ := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
