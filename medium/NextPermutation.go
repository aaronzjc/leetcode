package medium

import "sort"

// https://leetcode-cn.com/problems/next-permutation/

// Tips:
// 双指针，从后往前，如果在后面找到比前面大的，则交换过去。然后从小到大排序后半部分。

func nextPermutation(nums []int) {
	var i, j int
	i = len(nums) - 1
	for i >= 0 {
		j = len(nums) - 1
		for j > i {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				sort.Ints(nums[i+1:])
				return
			}
			j--
		}
		i--
	}
	for i, j = 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}
