package easy

// https://leetcode-cn.com/problems/length-of-last-word/

func lengthOfLastWord(s string) (res int) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res == 0 {
				continue
			}
			return
		}
		res++
	}
	return
}
