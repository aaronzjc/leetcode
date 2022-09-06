package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	seeds := []struct {
		input  int
		expect []string
	}{
		{1, []string{"()"}},
		{2, []string{"()()", "(())"}},
		{3, []string{"()()()", "((()))", "(())()", "()(())", "(()())"}},
	}

	for _, v := range seeds {
		result := generateParenthesis(v.input)
		if !tools.IsArrEquals(result, v.expect, false) {
			t.Error(v.input, v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("GenParenthese passed !")
}
