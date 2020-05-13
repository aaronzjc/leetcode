package medium

// https://leetcode-cn.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	sArr := []byte(s)
	if len(sArr) <= 1 {
		return s
	}
	var res []byte
	for i := 0; i < len(sArr); i++ {
		match := false
		for j := len(sArr) - 1; j > i; j-- {
			if j-i < len(res)-1 {
				break
			}
			end := -1
			m, n := i, j
			for m <= n {
				if sArr[m] == sArr[n] {
					if end == -1 {
						end = n
					}
					m++
					match = true
				} else {
					match = false
					if m > i {
						break
					}
				}
				n--
			}
			if match && (end-i+1) > len(res) {
				res = sArr[i : end+1]
			}
		}
	}

	return string(res)
}
