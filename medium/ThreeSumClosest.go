package medium

import (
	"sort"
)

// https://leetcode-cn.com/problems/3sum-closest/

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -1 * n
}

func threeSumClosest(nums []int, target int) (res int) {
	sort.Ints(nums)
	var i, j, k, minus, s int
	l := len(nums) - 1
	pre := -1
	for i <= l {
		j = i + 1
		for j <= l {
			k = j + 1
			for k <= l {
				s = nums[i] + nums[j] + nums[k]
				minus = abs(s - target)
				if minus == 0 {
					res = s
					return
				}
				if pre < 0 {
					pre = minus
					res = s
					k++
					continue
				}
				if minus < pre {
					pre = minus
					res = s
				}
				k++
			}
			j++
			for j < l && nums[j] == nums[j-1] {
				j++
			}
		}
		i++
		for i < l && nums[i] == nums[i-1] {
			i++
		}
	}
	return
}
