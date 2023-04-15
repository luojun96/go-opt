package string

import "strings"

// https://leetcode.cn/problems/lexicographically-smallest-string-after-applying-operations/
func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	res := s
	s = s + s
	g := gcd(b, n)

	for i := 0; i < n; i += g {
		t := []byte(s[i : i+n])
		add(t, n, a, 1)
		if b%2 != 0 {
			add(t, n, a, 0)
		}
		str := string(t)
		if strings.Compare(str, res) < 0 {
			res = str
		}
	}
	return res
}

func add(arr []byte, n int, a int, start int) {
	minVal, times := 10, 0
	for i := 0; i < 10; i++ {
		added := (int(arr[start]-'0') + i*a) % 10
		if added < minVal {
			minVal = added
			times = i
		}
	}

	for i := start; i < n; i += 2 {
		arr[i] = byte((int(arr[i]-'0')+times*a)%10) + '0'
	}
}
func gcd(num1 int, num2 int) int {
	for num2 != 0 {
		temp := num1
		num1 = num2
		num2 = temp % num2
	}
	return num1
}
