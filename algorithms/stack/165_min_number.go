package algorithms

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof
// input: [3,30,34,5,9]
// output: "3033459"
func minNumber(nums []int) string {
	var res string
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	quickSort(strs, 0, len(strs)-1)
	return strings.Join(strs, "")
}

func quickSort(strs []string, left int, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	tmp := strs[i]

	for i < j {
		for strings.Compare((strs[j]+strs[left]), (strs[left]+strs[j])) >= 0 && i < j {
			j--
		}

		for strings.Compare((strs[i]+strs[left]), (strs[left]+strs[i])) <= 0 && i < j {
			i++
		}
		tmp = strs[i]
		strs[i] = strs[j]
		strs[j] = tmp
	}
	strs[i] = strs[left]
	strs[left] = tmp
	quickSort(strs, left, i-1)
	quickSort(strs, i+1, right)
}
