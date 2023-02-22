package others

func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			r = mid - 1
			continue
		}
		if target > nums[mid] {
			l = mid + 1
			continue
		}
	}
	return -1
}

// BinarySearchLarge 二分查找比给定值大的元素
func BinarySearchLarge(nums []int, e int) (res int) {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] <= e {
			low = mid + 1
		} else {
			high = mid
			res = mid
		}
	}
	if nums[low] <= e {
		res = -1
	}
	return
}
