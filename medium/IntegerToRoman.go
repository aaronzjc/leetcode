package medium

// https://leetcode-cn.com/problems/integer-to-roman/

import (
	"sort"
	"strings"
)

var (
	flags  = []string{"I", "V", "X", "L", "C", "D", "M"}
	rels   = []int{1, 5, 10, 50, 100, 500, 1000}
	flags2 = []string{"IV", "IX", "XL", "XC", "CD", "CM"}
	rels2  = []int{4, 9, 40, 90, 400, 900}
)

func intToRoman(num int) (res string) {
	var chosen int
	var chosenStr string
	l := len(rels)
	l2 := len(rels2)
	for num > 0 {
		idx := sort.Search(l, func(i int) bool {
			return rels[i] > num
		})
		if idx > 0 && rels[idx-1] == num {
			res += flags[idx-1]
			break
		}
		idx2 := sort.Search(l2, func(i int) bool {
			return rels2[i] > num
		})
		if idx2 > 0 && rels2[idx2-1] == num {
			res += flags2[idx2-1]
			break
		}
		chosen = rels[idx-1]
		chosenStr = flags[idx-1]
		if idx2 > 0 && rels[idx-1] < rels2[idx2-1] {
			chosen = rels2[idx2-1]
			chosenStr = flags2[idx2-1]
		}
		counter := num / chosen
		res += strings.Repeat(chosenStr, counter)
		num = num - counter*chosen
	}
	return
}
