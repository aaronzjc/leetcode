package easy

// https://leetcode-cn.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)
	var mid int
	for low < high {
		mid = (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
