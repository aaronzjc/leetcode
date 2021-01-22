package others

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	inserts := []struct {
		word   string
		expect []string
	}{
		{"a", []string{"1", "a0"}},
		{"abc", []string{"2", "a1", "b1", "c0"}},
		{"abc", []string{"2", "a1", "b1", "c0"}},
		{"abd", []string{"3", "a2", "b2", "c0", "d0"}},
		{"acd", []string{"4", "a3", "b2", "c1", "c0", "d0", "d0"}},
		{"b", []string{"5", "a3", "b1", "b2", "c1", "c0", "d0", "d0"}},
		{"bc", []string{"6", "a3", "b2", "b2", "c1", "c0", "c0", "d0", "d0"}},
		{"bcd", []string{"7", "a3", "b2", "b2", "c1", "c0", "c0", "d0", "d0", "d0"}},
	}

	for _, w := range inserts {
		if !trie.Insert(w.word) {
			t.Error("insert error", w.word)
		}
		dp := trie.Dump()
		if tools.IsStringArrEquals(w.expect, dp, false) == false {
			t.Error("dump error", "dp =", dp, "expect =", w.expect)
		}
	}

	searches := []struct {
		word   string
		expect bool
	}{
		{"a", true},
		{"ab", false},
		{"abd", true},
		{"ad", false},
		{"abd", true},
		{"b", true},
		{"bd", false},
		{"bc", true},
	}

	for _, v := range searches {
		search := trie.Search(v.word)
		if search != v.expect {
			t.Error(v, search)
		}
	}

	t.Log(trie.Dump())

	t.Fatal("Trie success")
}
