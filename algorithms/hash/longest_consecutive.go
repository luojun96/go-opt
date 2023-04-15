package hash

// https://leetcode.cn/problems/WhsWhI/

import "fmt"

func longestConsecutive(nums []int) int {
	num := map[int]bool{}
	for _, n := range nums {
		num[n] = true
	}
	longest := 0
	for n := range num {
		if !num[n-1] {
			curNum := n
			curStreak := 1

			for num[curNum+1] {
				curNum++
				curStreak++
			}

			if longest < curStreak {
				longest = curStreak
			}
		}
	}
	return longest
}
func hash() {
	fmt.Println("vim-go")
}
