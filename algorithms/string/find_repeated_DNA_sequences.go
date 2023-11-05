package string

// https://leetcode.cn/problems/repeated-dna-sequences/description/?envType=daily-question&envId=2023-11-05

const length = 10

func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	cnt := map[string]int{}
	for i := 0; i <= len(s)-length; i++ {
		sub := s[i : i+length]
		cnt[sub]++
		if cnt[sub] == 2 {
			res = append(res, sub)
		}
	}
	return res
}
