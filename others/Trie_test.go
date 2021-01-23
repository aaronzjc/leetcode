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
		{"a", []string{"#1", "a0"}},
		{"abc", []string{"#2", "a1", "b1", "c0"}},
		{"abc", []string{"#2", "a1", "b1", "c0"}},
		{"abd", []string{"#3", "a2", "b2", "c0", "d0"}},
		{"acd", []string{"#4", "a3", "b2", "c1", "c0", "d0", "d0"}},
		{"b", []string{"#5", "a3", "b0", "b2", "c1", "c0", "d0", "d0"}},
		{"bc", []string{"#6", "a3", "b1", "b2", "c1", "c0", "c0", "d0", "d0"}},
		{"bcd", []string{"#7", "a3", "b2", "b2", "c1", "c1", "c0", "d0", "d0", "d0"}},
	}

	for _, w := range inserts {
		if !trie.Insert(w.word) {
			t.Error("insert error", w.word)
		}
		dp := trie.Dump()
		if tools.IsStringArrEquals(w.expect, trie.Dump(), false) == false {
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

	deletes := []struct {
		word   string
		result bool
		expect []string
	}{
		{"c", false, []string{"#7", "a3", "b2", "b2", "c1", "c1", "c0", "d0", "d0", "d0"}},
		{"b", true, []string{"#6", "a3", "b2", "b2", "c1", "c1", "c0", "d0", "d0", "d0"}},
		{"a", true, []string{"#5", "a3", "b2", "b2", "c1", "c1", "c0", "d0", "d0", "d0"}},
		{"acd", true, []string{"#4", "a2", "b2", "b2", "c1", "c0", "d0", "d0"}},
		{"abc", true, []string{"#3", "a1", "b2", "b1", "c1", "d0", "d0"}},
		{"abd", true, []string{"#2", "b2", "c1", "d0"}},
	}

	for _, v := range deletes {
		res := trie.Del(v.word)
		dp := trie.Dump()
		if res != v.result || !tools.IsStringArrEquals(dp, v.expect, false) {
			t.Error(v, res, dp)
		}
	}

	t.Log("Trie success")
}
