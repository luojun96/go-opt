package search

import "testing"

func TestReversePairs(t *testing.T) {
	nums := []int{7, 5, 6, 4}
	if reversePairs(nums) != 5 {
		t.Error("reversePairs error")
	}

	nums = []int{1, 3, 2, 3, 1}
	if reversePairs(nums) != 4 {
		t.Error("reversePairs error")
	}
}
