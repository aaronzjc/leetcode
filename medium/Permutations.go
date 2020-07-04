package medium

// https://leetcode-cn.com/problems/permutations/

// TIPS:
// 没什么特别的，回溯法，也叫DFS。

func permute(nums []int) [][]int {
	var res [][]int
	var gen func([]int, []int)
	gen = func(n []int, r []int) {
		if len(n) == 0 {
			res = append(res, r)
			return
		}

		for k, v := range n {
			tmp := []int{}
			if k > 0 {
				tmp = append(tmp, n[:k]...)
			}
			if k < len(n)-1 {
				tmp = append(tmp, n[k+1:]...)
			}
			rr := make([]int, len(r))
			rr = append(r, v)
			gen(tmp, rr)
		}
	}

	gen(nums, []int{})
	return res
}
