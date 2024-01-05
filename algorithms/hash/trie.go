package hash

type Trie struct {
	children     []*Trie
	isEndingChar bool
}

func Constructor() Trie {
	return Trie{
		children:     make([]*Trie, 26),
		isEndingChar: false,
	}
}

func (t *Trie) Insert(word string) {
	p := t
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if p.children[index] == nil {
			node := Constructor()
			p.children[index] = &node
		}
		p = p.children[index]
	}
	p.isEndingChar = true
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	p := t
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if p.children[index] == nil {
			return nil
		}
		p = p.children[index]
	}
	return p
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	if node != nil && node.isEndingChar {
		return true
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	if t.searchPrefix(prefix) != nil {
		return true
	}
	return false
}
