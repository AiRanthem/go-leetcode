package leetcode

type Trie struct {
	children []*Trie
	isWord   bool
}

// func Constructor() Trie {
// 	return Trie{
// 		children: make([]*Trie, 26),
// 	}
// }

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		t.isWord = true
		return
	}
	c := word[0] - 'a'
	if t.children[c] == nil {
		t.children[c] = &Trie{
			children: make([]*Trie, 26),
		}
	}
	t.children[c].Insert(word[1:])
}

func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return t.isWord
	}
	c := word[0] - 'a'
	if t.children[c] == nil {
		return false
	} else {
		return t.children[c].Search(word[1:])
	}
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	c := prefix[0] - 'a'
	if t.children[c] == nil {
		return false
	} else {
		return t.children[c].StartsWith(prefix[1:])
	}
}
