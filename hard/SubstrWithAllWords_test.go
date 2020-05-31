package hard

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	seeds := []struct {
		str    string
		words  []string
		expect []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"a", []string{"a", "a"}, []int{}},
		{"", []string{}, []int{}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
	}

	for _, v := range seeds {
		result := findSubstring(v.str, v.words)
		if !tools.IsIntArrEquals(result, v.expect, false) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("FindSubString passed !")
}
