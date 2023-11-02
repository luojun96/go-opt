package string

// https://leetcode.cn/problems/rings-and-rods/?envType=daily-question&envId=2023-11-02
func countPoints(rings string) int {
	cnt := 0
	rods := make(map[byte][]byte, 10)
	a := []byte(rings)
	for i := 0; i < len(a)/2; i++ {
		ring, rod := a[2*i], a[2*i+1]
		rods[rod] = append(rods[rod], ring)
	}
	for _, rings := range rods {
		if len(rings) < 3 {
			continue
		}
		if checkRings(rings) {
			cnt++
		}
	}
	return cnt
}

func checkRings(r []byte) bool {
	var check func(color byte) bool
	check = func(c byte) bool {
		for i := 0; i < len(r); i++ {
			if r[i] == c {
				return true
			}
		}
		return false
	}
	return check('R') && check('G') && check('B')
}

const (
	rodCount    = 10
	colorNumber = 3
)

func countPoints2(rings string) int {
	state := make([][]int, rodCount)
	for i := 0; i < rodCount; i++ {
		state[i] = make([]int, colorNumber)
	}

	var getColorId func(color byte) int
	getColorId = func(color byte) int {
		res := 0
		switch color {
		case 'R':
			res = 0
		case 'G':
			res = 1
		case 'B':
			res = 2
		}
		return res
	}

	n := len(rings)
	for i := 0; i < n/2; i++ {
		color := rings[2*i]
		rod := rings[2*i+1] - '0'
		state[rod][getColorId(color)] = 1
	}

	cnt := 0
	for i := 0; i < rodCount; i++ {
		flag := true
		for j := 0; j < colorNumber; j++ {
			if state[i][j] == 0 {
				flag = false
				break
			}
		}
		if flag {
			cnt++
		}
	}

	return cnt
}
