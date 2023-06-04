package dp

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	// test tables
	var tests = []struct {
		given    []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		actual := maxProfit(tt.given)
		if actual != tt.expected {
			t.Errorf("maxProfit(%v) = %v; expected %v", tt.given, actual, tt.expected)
		}
	}
}
