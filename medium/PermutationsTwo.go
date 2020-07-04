package medium

import (
	"sort"
)

// https://leetcode-cn.com/problems/permutations-ii/

// TIPS:
//和上一个题目一样，回溯法。不一样的是去重，去重就是同一次个选择时，如果选过了，则不继续选了。要先排序。

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var gen func([]int, []int)
	gen = func(n []int, r []int) {
		ln := len(n)
		if ln == 0 {
			res = append(res, r)
			return
		}

		for k, v := range n {
			if k > 0 && v == n[k-1] {
				continue
			}
			tmp := []int{}
			if k > 0 {
				tmp = append(tmp, n[:k]...)
			}
			if k < len(n)-1 {
				tmp = append(tmp, n[k+1:]...)
			}
			rr := make([]int, len(r))
			copy(rr, r)
			rr = append(rr, v)
			gen(tmp, rr)
		}
	}

	gen(nums, []int{})
	return res
}
