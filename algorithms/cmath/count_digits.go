package cmath

func countDigits(num int) int {
	t, res := num, 0
	for t != 0 {
		if num%(t%10) == 0 {
			res += 1
		}
		t /= 10
	}
	return res
}
