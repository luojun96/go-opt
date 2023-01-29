package main

import "math/rand"

// https://leetcode.cn/problems/random-pick-index/

type Solution struct {
	data []int
}

func Constructor(nums []int) Solution {
	return Solution{data: nums}
}

func (this *Solution) Pick(target int) int {
	var res int
	cnt := 0
	for i, num := range this.data {
		if num == target {
			cnt++
			if rand.Intn(cnt) == 0 {
				res = i	
			}
		}
	}
	return res
}
