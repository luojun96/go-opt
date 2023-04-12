package main

import "math/rand"

// https://leetcode.cn/problems/linked-list-random-node/
// Google: 给你一个未知长度的链表，请你设计一个算法，只能遍历一遍，随机地返回链表中的一个节点。
// 扩展：给你一个未知长度的序列，如何在其中随机地选择k个元素？
// 算法：水塘抽样算法
// 实现：先解决只抽取一个元素的问题，当遇到第i个元素时，应该有1/i的概率选择该元素，1 - i/i 的概率保持原有的选择
type RandomNode struct {
	head *ListNode
}

func NewRandomNode(head *ListNode) RandomNode {
	return RandomNode{head: head}
}

func (s *RandomNode) GetRandom() int {
	res := 0
	i := 0
	p := s.head
	for p != nil {
		i++
		if rand.Intn(i) == 0 {
			res = p.Val
		}
		p = p.Next
	}
	return res
}
