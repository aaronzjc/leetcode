package medium

import "sort"

// https://leetcode-cn.com/problems/group-anagrams/

// TIPS:
// 暴力匹配。将出现的字母分析成一个map。然后后面的字符串匹配map，如果匹配到了，表明相异，追加。
// 性能不好，主要是在构建字符串模式，和查找那一块，有优化的空间。看了一个其他的答案，可以构建一个26字母的出现次数的数组key。这样查找和匹配都很快了。

func groupAnagrams(strs []string) [][]string {
	sort.Strings(strs)
	var res [][]string
	var m []map[byte]int
	checkGroup := func(str string) (res int) {
		ms := make(map[byte]int)
		for i := 0; i < len(str); i++ {
			ms[str[i]]++
		}
		var i int
		var find bool
		for ; i < len(m); i++ {
			find = true
			if len(m[i]) != len(ms) {
				find = false
			} else {
				for kk, vv := range m[i] {
					c, ok := ms[kk]
					if !ok || c != vv {
						find = false
						break
					}
				}
			}
			if find {
				res = i
				break
			}
		}
		if !find {
			m = append(m, ms)
			res = len(m) - 1
		}
		return
	}

	var empty []string
	var g int
	for _, v := range strs {
		if v == "" {
			empty = append(empty, "")
			continue
		}
		g = checkGroup(v)
		if g <= len(res)-1 {
			res[g] = append(res[g], v)
			continue
		}
		res = append(res, []string{v})
	}
	if len(empty) > 0 {
		res = append(res, empty)
	}

	return res
}
