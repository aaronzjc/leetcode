package medium

import "testing"

func TestLongestPalindromic(t *testing.T) {
	seeds := []struct {
		input  string
		expect string
	}{
		{"cbbd", "bb"},
		{"abbadeffedgg", "deffed"},
		{"abccea", "cc"},
		{"a", "a"},
		{"aa", "aa"},
	}

	for _, v := range seeds {
		result := longestPalindrome(v.input)
		if result != v.expect {
			t.Error(v.input, result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("LongestPalindromic Passed !")
}
