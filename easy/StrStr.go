package easy

// https://leetcode-cn.com/problems/implement-strstr/

func strStr(haystack string, needle string) (res int) {
	lh := len(haystack)
	ln := len(needle)
	if ln == 0 {
		return
	}
	if lh < ln {
		res = -1
		return
	}
	var i int
	for i < lh {
		if haystack[i] != needle[0] {
			i++
			continue
		}
		if i+ln > lh {
			break
		}
		if haystack[i:i+ln] == needle {
			res = i
			return
		}
		i++
	}
	return -1
}
