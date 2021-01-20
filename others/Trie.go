package others

// Trie - 字典树
// 字典树常用于高效敏感词查找，路由匹配等

type TrieNode struct {
	Val rune
	Children map[rune]*TrieNode
	Flag bool
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
	var node *TrieNode
	w := []rune(word)
	i, lw := 0, len(w)
	for ; i < lw; i++ {
		found := false
		if node == nil {
			node = t.Root
		}
		for k, v := range node.Children {
			if k == w[i] {
				node = v
				found = true
				break
			}
		}

		if !found {
			node.Children[w[i]] = &TrieNode{
				Val: w[i],
			}
			node = node.Children[w[i]]
		}

		// 如果找到最后一个元素了，那么设置符号位，表明是一个单词
		if i == lw - 1 {
			node.Flag = true
		}
	}

	return true
}

func (t *Trie) Search(word string) bool {
	var node *TrieNode
	w := []rune(word)
	i, lw := 0, len(w)

	var found bool
	for ; i < lw; i++ {
		if node == nil {
			node = t.Root
		}

		for k, v := range node.Children {
			if k == w[i] {
				found = true
				node = v
				break
			}
		}

		if !found {
			return false
		}
	}

	return node.Flag
}