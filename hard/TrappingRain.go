package hard

import "fmt"

// https://leetcode-cn.com/problems/trapping-rain-water/

// TIPS:
// 我的思路比较笨拙，就是把整个区间划分成一个个完整的能盛水的区间容器。然后，依次计算容器的容量，相加。

func trap(height []int) (res int) {
	// 找到有效的容器范围
	var findContainer func(int, []int) (int, int)
	findContainer = func(start int, height []int) (left int, right int) {
		// 找到左边界
		for start < len(height)-1 {
			if height[start+1] < height[start] {
				break
			}
			start++
		}
		if start == len(height)-1 {
			return start, start
		}
		left = start
		right = start
		var min int
		for i := right + 1; i < len(height); i++ {
			if height[i] > height[start] {
				right = i
				return
			}
			if height[i] >= min {
				right = i
				min = height[i]
			}
		}
		return
	}

	// 计算区间的容量
	var calWater func(int, int, []int) int
	calWater = func(left int, right int, height []int) (total int) {
		min := height[left]
		if height[right] < min {
			min = height[right]
		}
		total = 0
		for i := left; i < right; i++ {
			if height[i] >= min {
				continue
			}
			total += min - height[i]
		}
		return
	}

	i := 0
	for i < len(height) {
		left, right := findContainer(i, height)
		fmt.Println(left, right)
		if left == right {
			break
		}
		i = right
		if left == right-1 {
			continue
		}
		area := calWater(left, right, height)
		res += area
	}

	return
}
