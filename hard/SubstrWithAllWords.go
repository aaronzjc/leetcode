package hard

// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/

// Tips:
// 截取定长的字符串匹配字符数组。如果匹配了，则添加索引，否则，往后移。
// 注意数组拷贝和删除一个元素。

func findSubstring(s string, words []string) (res []int) {
	ls := len(s)
	var i, j int

	for _, v := range words {
		j += len(v)
	}

	checkStrMatch := func(str string, words []string) bool {
		if str == "" || len(words) == 0 {
			return false
		}
		var i int
		tmp := make([]string, len(words))
		copy(tmp, words) // 不要改变初始数组
		match := 0
	MAT:
		lt := len(tmp)
		for i = 0; i < lt; i++ {
			sl := len(tmp[i])
			if str[:sl] == tmp[i] {
				tmp[i] = tmp[lt-1] // 因为单词数组的顺序不重要，所以这里可以这么做
				tmp = tmp[:lt-1]
				str = str[sl:]
				match++
				goto MAT
			}
		}

		return match == len(words)
	}

	for j <= ls {
		if checkStrMatch(s[i:j], words) {
			res = append(res, i)
		}
		i++
		j++
	}
	return
}
