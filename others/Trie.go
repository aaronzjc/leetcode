package others

import (
	"fmt"
	"strconv"
)

// Trie - 字典树
// 字典树常用于高效敏感词查找，路由匹配等

type TrieNode struct {
	Val      byte
	Children []*TrieNode
	Ref      int
	Flag     bool
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Val: byte('#'),
		},
	}
}

func (t *Trie) Insert(word string) bool {
	// 这里牺牲一点点性能先搜索是为了防止重复的字符串插入，导致计数异常
	if t.Search(word) {
		return true
	}

	w := []byte(word)
	i, lw := 0, len(w)
	node := t.Root
	for ; i < lw; i++ {
		found := false
		for _, v := range node.Children {
			if v.Val == w[i] {
				node.Ref++
				node = v
				found = true
				break
			}
		}

		if !found {
			node.Children = append(node.Children, &TrieNode{
				Val: w[i],
			})
			node.Ref++
			node = node.Children[len(node.Children)-1]
		}
	}

	node.Flag = true

	return true
}

func (t *Trie) Search(word string) bool {
	w := []byte(word)
	i, lw := 0, len(w)

	node := t.Root
	for ; i < lw; i++ {
		found := false
		for _, v := range node.Children {
			if v.Val == w[i] {
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

func (t *Trie) Del(word string) bool {
	if t.Search(word) == false {
		return false
	}
	w := []byte(word)
	i, lw := 0, len(w)
	node := t.Root
	t.Root.Ref--
	for ; i < lw; i++ {
		for k, v := range node.Children {
			if v.Val != w[i] {
				continue
			}
			// 只有分支才需要减少引用，如果本身是单词节点，不需要减
			if v.Flag {
				v.Flag = false
			} else {
				v.Ref--
			}

			// 如果分支的引用为0了，直接削了
			if v.Ref == 0 {
				node.Children = append(node.Children[:k], node.Children[k+1:]...)
				return true
			}

			node = v
		}
	}

	return true
}

func (t *Trie) Dump() []string {
	var l []*TrieNode
	var res []string
	l = append(l, t.Root)
	for len(l) > 0 {
		for _, node := range l {
			res = append(res, fmt.Sprintf("%s%s", string(node.Val), strconv.Itoa(node.Ref)))
			l = l[1:]
			for _, child := range node.Children {
				l = append(l, child)
			}
		}
	}
	return res
}
