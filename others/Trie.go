package others

// Trie - 字典树
// 字典树常用于高效敏感词查找，路由匹配等

type TrieNode struct {
	Val rune
	Children []*TrieNode
	Ref int
}

func NewNode(r rune) *TrieNode {
	return &TrieNode{
		Val: r,
	}
}

func (tn *TrieNode) Add(r rune) bool {
	// 如果元素存在，不重复添加
	for _, v := range tn.Children {
		if v.Val == r {
			return false
		}
	}

	node := &TrieNode{
		Val: r,
	}
	tn.Children = append(tn.Children, node)
	tn.Ref++
	return true
}

func (tn *TrieNode) Del(r rune) bool {
	for _, v := range tn.Children {
		if v.Val == r {
			
		}
	}
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: new(TrieNode),
	}
}

func (t *Trie) Insert(word string) bool {
	if len(word) == 0 {
		return false
	}
	w := []rune(word)
	for _, v := range t.Root.Children {
		if v.Val == word[]
	}
}

func (t *Trie) Del(word string) {

}

func (t *Trie) Has(word string) bool {

}