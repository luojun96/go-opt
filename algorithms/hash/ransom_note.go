package hash

// https://leetcode.cn/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int, len(magazine))
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		count, ok := m[ransomNote[i]]
		if !ok {
			return false
		}
		count--
		if count < 0 {
			return false
		}
		m[ransomNote[i]] = count
	}

	return true
}
