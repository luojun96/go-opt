package algorithms

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
// input: [1,3,2,6,5]
//
//	    5
//	   / \
//	  2   6
//	 / \
//	1   3
//
// output: false
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}

	root := postorder[len(postorder)-1]

	// get the started index of right sub nodes
	i := len(postorder) - 2
	for ; i >= 0; i-- {
		if postorder[i] < root {
			break
		}
	}

	// check the values of left sub nodes
	j := i
	for ; j >= 0; j-- {
		if postorder[j] > root {
			return false
		}
	}

	left, right := true, true
	if i >= 0 {
		left = verifyPostorder(postorder[:i+1])
	}

	if i < len(postorder)-1 {
		right = verifyPostorder(postorder[i+1 : len(postorder)-1])
	}

	return left && right

}
