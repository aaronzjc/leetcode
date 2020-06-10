package medium

import (
	"sort"
)

// https://leetcode-cn.com/problems/combination-sum-ii/

// TIPS:
// 和前一个题目一模一样，但是，这个情况需要考虑去重。怎么去重？排序后，下一次遍历和之前一样的话，就跳过。因为你走过的路，之前的人肯定走过。

func combinationSum2(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	var resolve func([]int, int, []int)
	resolve = func(choices []int, target int, solutions []int) {
		if target == 0 {
			res = append(res, solutions)
			return
		}
		for i := 0; i < len(choices); i++ {
			if choices[i] > target {
				return
			}
			if i > 0 && choices[i] == choices[i-1] {
				continue
			}
			tmp := make([]int, len(solutions))
			copy(tmp, solutions)
			tmp = append(tmp, choices[i])
			resolve(choices[i+1:], target-choices[i], tmp)
		}
	}

	resolve(candidates, target, []int{})

	return
}
