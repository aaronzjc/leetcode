package hard

import "testing"

func TestWildcardMatching(t *testing.T) {
	seeds := []struct {
		input  string
		match  string
		expect bool
	}{
		{"", "a", false},
		{"a", "a*", true},
		{"aa", "a", false},
		{"aa", "a?", true},
		{"acdcb", "a*c?b", false},
		{"mississippi", "m??*ss*?i*pi", false},
		{"aaaa", "***a", true},
		{"cb", "?a", false},
		{"aa", "*", true},
		{"ab", "?*", true},
		{"adceb", "*a*b", true},
		{"abcde", "*a*bcd*cdf*", false},
		{"abbbabaaabbabbabbabaabbbaabaaaabbbabaaabbbbbaaababbb", "**a*b*aa***b***bbb*ba*a", false},
		{"babbbbaabababaabbababaababaabbaabababbaaababbababaaaaaabbabaaaabababbabbababbbaaaababbbabbbbbbbbbbaabbb", "b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a", false},
	}

	for _, v := range seeds {
		result := isMatch(v.input, v.match)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("WildcardMatching passed !")
}
