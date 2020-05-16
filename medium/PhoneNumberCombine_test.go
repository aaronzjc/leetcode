package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestLettersCombination(t *testing.T) {
	seeds := []struct {
		input  string
		expect []string
	}{
		{"3", []string{"d", "e", "f"}},
		{"32", []string{"da", "db", "dc", "ea", "eb", "ec", "fa", "fb", "fc"}},
	}

	for _, v := range seeds {
		result := letterCombinations(v.input)
		if !tools.IsStringArrEquals(result, v.expect, false) {
			t.Error(v.input, v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("LettersCombination passed !")
}
