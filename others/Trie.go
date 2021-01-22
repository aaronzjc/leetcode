package others

import (
	"fmt"
	"strconv"
)

// Trie - 字典树
// 字典树常用于高效敏感词查找，路由匹配等

type TrieNode struct {
	Val      rune
	Children []*TrieNode
	Ref      int
	Flag     bool
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
	if t.Search(word) {
		return true
	}

	w := []rune(word)
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
	w := []rune(word)
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
		return true
	}
	w := []rune(word)
	i, lw := 0, len(w)
	node := t.Root
	for ; i < lw; i++ {
		for k, v := range node.Children {
			if v.Val != w[i] {
				continue
			}
			v.Ref--
			if v.Ref == 0 {
				node.Children = append(node.Children[:k], node.Children[k+1:]...)
			}
			break
		}
	}

	if node.Ref > 0 {
		node.Flag = false
	}

	return true
}

func (t *Trie) Dump() []string {
	var l []*TrieNode
	var res []string
	l = append(l, t.Root)
	for len(l) > 0 {
		for _, node := range l {
			res = append(res, fmt.Sprintf("%s=%s", string(node.Val), strconv.Itoa(node.Ref)))
			l = l[1:]
			for _, child := range node.Children {
				l = append(l, child)
			}
		}
	}
	return res
}
