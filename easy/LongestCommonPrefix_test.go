package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	seeds := []struct {
		input  []string
		exepct string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"abcd", "def"}, ""},
		{[]string{}, ""},
		{[]string{"abc", ""}, ""},
		{[]string{"c", "c"}, "c"},
	}

	for _, v := range seeds {
		result := longestCommonPrefix(v.input)
		if result != v.exepct {
			t.Error(v.input, v.exepct, result)
			t.Fatalf("failed !")
		}
	}

	t.Log(" LongestCommonPrefix passed !")
}
