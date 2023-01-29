package main

func maxNumberOfBalloons(text string) int {
	count := [26]int{}
	for i := 0; i < len(text); i++ {
		count[text[i] - 'a']++
	}

	// 'b' -> 1
	min := count[1]

	// 'a' -> 0
	min = Min(min, count[0])

	// 'l' -> 11
	min = Min(min, count[11] / 2)

	// 'o' -> 14
	min = Min(min, count[14] / 2)

	// 'n' -> 13
	min = Min(min, count[13])

	return min
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
