package easy

// https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) (res string) {
	if len(strs) == 0 {
		return ""
	}
	for _, v := range strs {
		if v == "" {
			return ""
		}
	}
	var c byte
	var i int
	for {
		c = 0
		for _, v := range strs {
			if i >= len(v) {
				return
			}
			if c == 0 {
				c = v[i]
				continue
			}
			if v[i] != c {
				return
			}
		}
		res += string(c)
		i++
	}
}
