package medium

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) (res []int) {
	res = []int{-1, -1}
	low, high := 0, len(nums)-1
	if high < 0 {
		return
	}
	for low <= high {
		if res[0] == -1 {
			if nums[low] == target {
				res[0] = low
			} else {
				low++
			}
		}
		if res[1] == -1 {
			if nums[high] == target {
				res[1] = high
			} else {
				high--
			}
		}
		if res[0] != -1 && res[1] != -1 {
			return
		}
	}

	return
}
