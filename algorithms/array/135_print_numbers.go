package main

import (
	"math"
	"strconv"
)

// https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
// input: n = 1
// output: [1,2,3,4,5,6,7,8,9]
func printNumbers(n int) []int {
	res := make([]int, 0, int(math.Pow10(n)-1))
	count := 0
	var dfs func(num []byte, index int, digit int)
	dfs = func(num []byte, index int, digit int) {
		if index == digit {
			tmp, _ := strconv.Atoi(string(num))
			res[count] = tmp
			count++
			return
		}
		for i := '0'; i <= '9'; i++ {
			num[index] = byte(i)
			dfs(num, index+1, digit)
		}

	}

	for i := 0; i <= n; i++ {
		num := make([]byte, i)
		for j := '1'; j <= '9'; j++ {
			num[0] = byte(j)
			dfs(1, num, i)
		}
	}
	return res
}
