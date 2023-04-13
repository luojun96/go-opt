package algorithms

// https://leetcode.cn/problems/longest-common-prefix/
// input: strs = ["flower","flow","flight"]
// output: "fl"
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	count := len(strs)

	for i := 0; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}

	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
