package main

// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/description/
// input: s = "abaccdeff"
// output: 'b'
func firstUniqChar(s string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

type pair struct {
	ch  byte
	pos int
}

func firstUniqCharByQueue(s string) byte {
	n := len(s)
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}

	q := []pair{}
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		ch := bytes[i] - 'a'
		if pos[ch] == n {
			pos[ch] = i
			q = append(q, pair{ch, i})
		} else {
			pos[ch] = n + 1
			for len(q) > 0 && pos[q[0].ch] == n+1 {
				q = q[1:]
			}
		}
	}

	if len(q) > 0 {
		return bytes[q[0].pos]
	}
	return ' '
}
