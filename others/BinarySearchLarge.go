package others

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
