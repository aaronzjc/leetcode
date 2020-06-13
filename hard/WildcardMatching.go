package hard

// https://leetcode-cn.com/problems/wildcard-matching/

func isMatch(s string, p string) bool {
	var pre byte
	var ps []byte
	for i := 0; i < len(p); i++ {
		if p[i] == pre && p[i] == '*' {
			continue
		}
		ps = append(ps, p[i])
		pre = p[i]
	}

	var checkStr func(string, []byte) bool
	checkStr = func(str string, pattern []byte) bool {
		ls, lp := len(str), len(pattern)
		if str == string(pattern) {
			return true
		}
		if lp > 1 {
			if pattern[lp-1] != '*' {
				if pattern[lp-1] != '?' && pattern[lp-1] != str[ls-1] {
					return false
				}
			}
		}
		var i int
		for i < ls && i < lp {
			if pattern[i] == '*' {
				if i == lp-1 {
					return true
				}
				var pl int
				for i := 0; i < lp; i++ {
					if pattern[i] != '*' {
						pl++
					}
				}
				for j := 0; j <= ls-pl; j++ {
					match := checkStr(string(str[i+j:]), pattern[i+1:])
					if match {
						return true
					}
				}
				return false
			} else {
				if pattern[i] == '?' || pattern[i] == str[i] {
					i++
					continue
				}
				return false
			}
		}
		if i > 0 && i == lp && i == ls {
			return true
		}
		if len(pattern[i:]) == 1 && pattern[i] == '*' {
			return true
		}
		return false
	}

	return checkStr(s, ps)
}
