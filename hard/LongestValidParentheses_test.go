package hard

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	seeds := []struct {
		input  string
		expect int
	}{
		{"(()", 2},
		{"", 0},
		{"(", 0},
		{"()", 2},
		{")()())", 4},
		{"()(()", 2},
		{"(()(((()", 2},
		{"(()(((())))()", 12},
		{")(()((((())", 4},
		{")(((((()())()()))()(()))(", 22},
		{"()(((()(()))))", 14},
	}

	for _, v := range seeds {
		result := longestValidParentheses(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("LongestValidParentheses passed !")
}
