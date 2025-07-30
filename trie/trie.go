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

type XORNode struct {
	End      bool
	Children [2]*XORNode
}

type XORTrie struct {
	root *XORNode
}

func NewXORTrie() *XORTrie {
	return &XORTrie{root: &XORNode{}}
}

func (this *XORTrie) Insert(num int) {
	p := this.root
	for i := 30; i >= 0; i-- {
		t := num >> i & 1
		if p.Children[t] == nil {
			p.Children[t] = &XORNode{}
		}
		p = p.Children[t]
	}
	p.End = true
}

func (this *XORTrie) Search(num int) int {
	res := 0
	p := this.root
	for i := 30; i >= 0; i-- {
		t := num >> i & 1
		if p.Children[t^1] != nil {
			res += 1 << i
			p = p.Children[t^1]
		} else if p.Children[t] != nil {
			p = p.Children[t]
		} else {
			break
		}
	}
	return res
}

// 421
func findMaximumXOR(nums []int) int {
	trie := NewXORTrie()
	maxXOR := 0
	maxInt := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _, num := range nums {
		maxXOR = maxInt(maxXOR, trie.Search(num))
		trie.Insert(num)
	}
	return maxXOR
}
