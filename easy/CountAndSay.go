package easy

// https://leetcode-cn.com/problems/count-and-say/

func countAndSay(n int) (res string) {
	var counter func([]byte) []byte
	counter = func(s []byte) (res []byte) {
		if len(s) == 0 {
			return []byte{'1'}
		}

		pre := s[0]
		count := 1
		for i := 1; i < len(s); i++ {
			if pre != s[i] {
				res = append(res, byte('0'+count))
				res = append(res, pre)
				pre = s[i]
				count = 0
			}
			count++
		}
		if count != 0 {
			res = append(res, byte('0'+count))
			res = append(res, pre)
		}

		return
	}

	var cal []byte
	for i := 1; i <= n; i++ {
		cal = counter(cal)
	}

	return string(cal)
}
