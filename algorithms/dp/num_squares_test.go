package dp

import (
	"testing"
)

func TestNumSquares(t *testing.T) {
	// use test tables
	var tests = []struct {
		given    int
		expected int
	}{
		{12, 3},
		{13, 2},
		{14, 3},
	}

	for _, tt := range tests {
		actual := numSquares(tt.given)
		if actual != tt.expected {
			t.Errorf("numSquares(%v): expected %v, actual %v", tt.given, tt.expected, actual)
		}
	}
}
