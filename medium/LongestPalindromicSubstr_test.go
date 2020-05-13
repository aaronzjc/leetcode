package medium

import "testing"

func TestLongestPalindromic(t *testing.T) {
	input := "cbbd"
	expect := "bb"

	result := longestPalindrome(input)

	if result != expect {
		t.Fatalf("failed !")
	}

	t.Log("LongestPalindromic Passed !")
}
