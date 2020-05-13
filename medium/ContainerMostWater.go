package medium

// https://leetcode-cn.com/problems/container-with-most-water/

func maxArea(height []int) int {
	var result int
	var left, right int
	right = len(height) - 1
	for left < right {
		h := height[left]
		if h > height[right] {
			h = height[right]
		}
		w := right - left

		area := w * h
		if result < area {
			result = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
