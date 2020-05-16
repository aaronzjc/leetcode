package medium

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

var (
	m = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) (res []string) {
	var i, l int
	for _, v := range []byte(digits) {
		vs := string(v)
		vsv := m[vs]
		if len(res) == 0 {
			res = vsv
			continue
		}
		i = 0
		l = len(res)
		for i < l {
			for _, vv := range vsv {
				res = append(res, res[i]+vv)
			}
			i++
		}
		res = res[i:]
	}
	return
}
