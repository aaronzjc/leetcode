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
		{"a", []string{"=1", "a=0"}},
		{"abc", []string{"=2", "a=1", "b=1", "c=0"}},
		{"abc", []string{"=2", "a=1", "b=1", "c=0"}},
		{"abd", []string{"=3", "a=2", "b=2", "c=0", "d=0"}},
		{"acd", []string{"=4", "a=3", "b=2", "c=1", "c=0", "d=0", "d=0"}},
		{"b", []string{"=5", "a=3", "b=1", "b=2", "c=1", "c=0", "d=0", "d=0"}},
		{"bc", []string{"=6", "a=3", "b=2", "b=2", "c=1", "c=0", "c=0", "d=0", "d=0"}},
		{"bcd", []string{"=7", "a=3", "b=2", "b=2", "c=1", "c=0", "c=0", "d=0", "d=0", "d=0"}},
	}

	for _, w := range inserts {
		if !trie.Insert(w.word) {
			t.Error("insert error", w.word)
		}
		dp := trie.Dump()
		if tools.IsStringArrEquals(dp, w.expect, true) == false {
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
