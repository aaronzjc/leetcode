package medium

// https://leetcode-cn.com/problems/integer-to-roman/

import (
	"sort"
)

var (
	flags = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	rels  = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
)

func intToRoman(num int) (res string) {
	for num > 0 {
		idx := sort.Search(len(rels), func(i int) bool {
			return rels[i] > num
		})
		if idx > 0 && rels[idx-1] == num {
			res += flags[idx-1]
			break
		}
		res += flags[idx-1]
		num = num - rels[idx-1]
	}
	return
}
