package medium

import "sort"

// https://leetcode-cn.com/problems/combination-sum/

// TIPS:
// 又是回溯。选择当前一个节点，然后判断接下来的节点。

func combinationSum(candidates []int, target int) (res [][]int) {
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
			jj := target / choices[i]
			for j := 0; j < jj; j++ {
				tmp := make([]int, len(solutions))
				copy(tmp, solutions)
				sum := 0
				for k := 0; k <= j; k++ {
					tmp = append(tmp, choices[i])
					sum += choices[i]
				}

				resolve(choices[i+1:], target-sum, tmp)
			}
		}
	}

	resolve(candidates, target, []int{})

	return
}
