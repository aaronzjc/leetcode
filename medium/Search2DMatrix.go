package medium

// https://leetcode-cn.com/problems/search-a-2d-matrix/

// Tips:
// 有点像二维空间下的二分查找，要做好地址映射

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	left, right := 0, rows*cols-1
	mid := (left + right) / 2

	for left <= right {
		y := mid / cols
		x := mid - y*cols

		if matrix[y][x] == target {
			return true
		}
		if matrix[y][x] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return false
}
