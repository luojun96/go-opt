package hash

import (
	"testing"
)

func TestNewTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	res := trie.Search("apple")
	if res == false {
		t.Fail()
	}
	res = trie.Search("app")
	if res == true {
		t.Fail()
	}
	res = trie.StartsWith("app")
	if res == false {
		t.Fail()
	}
	trie.Insert("app")
	res = trie.Search("app")
	if res == false {
		t.Fail()
	}
}
