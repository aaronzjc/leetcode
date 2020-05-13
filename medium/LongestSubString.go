package medium

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func searchArr(e byte, arr []byte) int {
	for k, v := range arr {
		if v == e {
			return k
		}
	}
	return -1
}

func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	var w []byte
	str := []byte(s)
	max := 0
	for i := 0; i < strLen; i++ {
		if idx := searchArr(str[i], w); idx >= 0 {
			w = w[(idx + 1):]
		}
		w = append(w, str[i])
		if len(w) > max {
			max = len(w)
		}
	}
	return max
}
