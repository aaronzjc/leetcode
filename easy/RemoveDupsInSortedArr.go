package easy

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

// Tips: 很简单。注意怎么删除重复元素。

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	i := 0
	j := 1
	for j < l {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	nums = nums[:i+1]
	return i + 1
}
