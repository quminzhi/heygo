package trie

type TrieNode struct {
	Val      int
	Children [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	p := this.root
	for _, ch := range word {
		c := int(ch - 'a')
		if p.Children[c] == nil {
			p.Children[c] = &TrieNode{}
		}
		p = p.Children[c]
	}
	p.Val++
}

func (this *Trie) Search(word string) bool {
	p := this.root
	for _, ch := range word {
		c := int(ch - 'a')
		if p.Children[c] == nil {
			return false
		}
		p = p.Children[c]
	}
	return p.Val != 0
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for _, ch := range prefix {
		c := int(ch - 'a')
		if p.Children[c] == nil {
			return false
		}
		p = p.Children[c]
	}
	return true
}
