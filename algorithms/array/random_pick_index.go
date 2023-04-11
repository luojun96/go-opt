package main

import (
	"math/rand"
)

// https://leetcode.cn/problems/random-pick-index/

type Solution struct {
	nums  []int
	cache map[int][]int
}

func NewSolution(nums []int) Solution {
	m := make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		m[v] = append(m[v], i)
	}
	return Solution{nums: nums, cache: m}
}

func (s *Solution) Pick(target int) int {
	a := s.cache[target]
	return a[rand.Intn(len(a))]
	// res := 0
	// l := 0

	// for i, num := range s.nums {
	// 	if num == target {
	// 		l++
	// 		if rand.Intn(l) == 0 {
	// 			res = i
	// 		}
	// 	}
	// }
	// for i := 0; i < len(s.nums); i++ {
	// 	if s.nums[i] == target {
	// 		l++
	// 		if rand.Intn(l) == 0 {
	// 			res = i
	// 		}
	// 	}
	// }

	// return res
}
