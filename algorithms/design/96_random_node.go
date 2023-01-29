package main

import "math/rand"

// https://leetcode.cn/problems/linked-list-random-node/

type ListNode struct {
	Val int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	var res int
	for node, i := this.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 {
			res = node.Val	
		}
	}
	return res
}
