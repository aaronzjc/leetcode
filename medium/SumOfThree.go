package medium

import (
	"sort"
)

// https://leetcode-cn.com/problems/3sum/

func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	var i, j, k int
	l := len(nums) - 1
	for i <= l {
		j = i + 1
		for j <= l {
			k = l
			for j != k {
				s := nums[i] + nums[j] + nums[k]
				if s > 0 {
					k--
					continue
				} else if s == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
				break
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
