package main

// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

type CQueue struct {
	instack, outstack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (c *CQueue) AppendTail(value int) {
	c.instack = append(c.instack, value)
}

func (c *CQueue) DeleteHead() int {
	if len(c.outstack) == 0 {
		if len(c.instack) == 0{
			return -1
		}
		c.in2out()
	}	

	value := c.outstack[len(c.outstack)-1]
	c.outstack = c.outstack[:len(c.outstack)-1]
	return value
}

func (c *CQueue) in2out() {
	for len(c.inStack) > 0 {
		c.outstack = append(c.outstack, c.instack[len(c.instack)-1])
		c.instack = c.instack[:len(c.instack)-1]
	}
}
