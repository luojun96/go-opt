package algorithms

import (
	"strings"
)

type TrieNode struct {
	data         byte
	children     [26]*TrieNode
	isEndingChar bool
}

func (t *TrieNode) insert(s string) {
	p := t
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if p.children[index] == nil {
			newNode := &TrieNode{data: s[i]}
			p.children[index] = newNode
		}
		p = p.children[index]
	}
	p.isEndingChar = true

}

func (t *TrieNode) find(s string) string {
	p := t
	i := 0
	for ; i < len(s); i++ {
		index := s[i] - 'a'
		if p.isEndingChar {
			return s[:i]
		}
		if p.children[index] == nil {
			return s
		}
		p = p.children[index]
	}
	return s
}

func replaceWords_01(dictionary []string, sentence string) string {
	root := &TrieNode{data: '/'}
	for _, s := range dictionary {
		root.insert(s)
	}

	words := strings.Split(sentence, " ")
	for i, s := range words {
		words[i] = root.find(s)
	}
	return strings.Join(words, " ")
}
