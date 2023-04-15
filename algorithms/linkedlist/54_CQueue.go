package linkedlist

// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
type CQueue struct {
	instack  []int
	outstack []int
}

func NewCQueue() CQueue {
	return CQueue{
		instack:  []int{},
		outstack: []int{},
	}
}

func (cq *CQueue) AppendTail(value int) {
	cq.instack = append(cq.instack, value)
}

func (cq *CQueue) DeleteHead() int {
	if len(cq.outstack) == 0 && len(cq.instack) == 0 {
		return -1
	}

	var res int
	if len(cq.outstack) > 0 {
		res = cq.outstack[len(cq.outstack)-1]
		cq.outstack = cq.outstack[:len(cq.outstack)-1]
	} else if len(cq.outstack) == 0 && len(cq.instack) > 0 {
		for len(cq.instack) > 0 {
			if len(cq.instack) == 1 {
				res = cq.instack[0]
				cq.instack = []int{}
				break
			}

			v := cq.instack[len(cq.instack)-1]
			cq.instack = cq.instack[:len(cq.instack)-1]
			cq.outstack = append(cq.outstack, v)
		}
	}

	return res
}
