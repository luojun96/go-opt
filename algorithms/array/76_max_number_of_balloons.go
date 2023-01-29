package main

// https://leetcode.cn/problems/maximum-number-of-balloons/
// input: text = "nlaebolko"
// output: 1
func maxNumberOfBalloons(text string) int {
	cnt := [5]int{}
	for _, ch := range text {
		if ch == 'b' {
			cnt[0]++
		} else if ch == 'a' {
			cnt[1]++
		} else if ch == 'l' {
			cnt[2]++
		} else if ch == 'o' {
			cnt[3]++
		} else if ch == 'n' {
			cnt[4]++
		}
	}

	cnt[2] /= 2
	cnt[3] /= 2
	ans := cnt[0]
	for _, v := range cnt[1:] {
		if v < ans {
			ans = v
		}
	}

	return ans
}
