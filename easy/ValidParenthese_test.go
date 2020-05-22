package easy

import "testing"

func TestValidParenthese(t *testing.T) {
	seeds := []struct {
		input  string
		expect bool
	}{
		{"()()()({[]})[", false},
		{"([({})])", true},
		{"[[[", false},
		{")))", false},
		{"((]", false},
		{"[[))", false},
		{"(}}", false},
	}

	for _, v := range seeds {
		result := isValid(v.input)
		if result != v.expect {
			t.Error(v.input, v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("IsValidParenthese passed !")
}
