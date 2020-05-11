package medium

import "testing"

func TestLongestSubString(t *testing.T) {
	input := "pwwkew"
	expect := len("wke")

	result := lengthOfLongestSubstring(input)
	if result != expect {
		t.Error("failed !")
	}

	t.Log("LongestSubString passed !")
}
