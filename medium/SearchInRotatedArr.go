package medium

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

// TIP:
// 比较笨的方式，在二分查找基础上，判断不同的边界情况。

func search(nums []int, target int) int {
	var mid int
	low, high := 0, len(nums)-1
	if high < 0 {
		return -1
	}
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if low == high {
			break
		}
		if nums[mid] > target {
			if nums[low] < target {
				high = mid
			} else if nums[low] > target {
				if nums[low] > nums[mid] {
					high = mid
				} else {
					low = mid + 1
				}
			} else {
				return low
			}
		} else {
			if nums[high] > target {
				low = mid + 1
			} else if nums[high] < target {
				if nums[high] < nums[mid] {
					low = mid + 1
				} else {
					high = mid
				}
			} else {
				return high
			}
		}
	}
	return -1
}
