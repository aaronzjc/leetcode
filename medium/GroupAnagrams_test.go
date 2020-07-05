package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestGroupAnagrams(t *testing.T) {
	seeds := []struct {
		input  []string
		expect [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
		{
			[]string{"tea", "and", "ace", "ad", "eat", "dans"},
			[][]string{{"tea", "eat"}, {"and"}, {"ace"}, {"dans"}, {"ad"}},
		},
		{
			[]string{"", "b"},
			[][]string{{""}, {"b"}},
		},
		{
			[]string{"", ""},
			[][]string{{"", ""}},
		},
		{
			[]string{"hos", "boo", "nay", "deb", "wow", "bop", "bob", "brr", "hey", "rye", "eve", "elf", "pup", "bum", "iva", "lyx", "yap", "ugh", "hem", "rod", "aha", "nam", "gap", "yea", "doc", "pen", "job", "dis", "max", "oho", "jed", "lye", "ram", "pup", "qua", "ugh", "mir", "nap", "deb", "hog", "let", "gym", "bye", "lon", "aft", "eel", "sol", "jab"},
			[][]string{{"sol"}, {"wow"}, {"gap"}, {"hem"}, {"yap"}, {"bum"}, {"ugh", "ugh"}, {"aha"}, {"jab"}, {"eve"}, {"bop"}, {"lyx"}, {"jed"}, {"iva"}, {"rod"}, {"boo"}, {"brr"}, {"hog"}, {"nay"}, {"mir"}, {"deb", "deb"}, {"aft"}, {"dis"}, {"yea"}, {"hos"}, {"rye"}, {"hey"}, {"doc"}, {"bob"}, {"eel"}, {"pen"}, {"job"}, {"max"}, {"oho"}, {"lye"}, {"ram"}, {"nap"}, {"elf"}, {"qua"}, {"pup", "pup"}, {"let"}, {"gym"}, {"nam"}, {"bye"}, {"lon"}},
		},
	}

	for _, v := range seeds {
		result := groupAnagrams(v.input)
		if !tools.IsL2StrArrayEquals(v.expect, result, false) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("GroupAnagrams passed !")
}
