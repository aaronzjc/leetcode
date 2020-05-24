package easy

// https://leetcode-cn.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	var i, j int
	for j < l {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	nums = nums[:i]
	return i
}
