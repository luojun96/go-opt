package hash

import "strings"

// https://leetcode.cn/problems/replace-words/
// input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// output: "the cat was rat by the bat"

func replaceWords(dictionary []string, sentence string) string {
	dictionarySet := map[string]bool{}
	for _, v := range dictionary {
		dictionarySet[v] = true
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			if dictionarySet[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}
	return strings.Join(words, " ")
}
