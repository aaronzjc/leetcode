package easy

// https://leetcode-cn.com/problems/reverse-integer/

import "math"

// IntMax 最大整数
const IntMax = 1<<31 - 1

// IntMin 最小整数
const IntMin = ^IntMax

func reverse(x int) int {
	var result int
	for x != 0 {
		t := x % 10
		result = result*10 + t
		x = x / 10
	}

	if result > math.MaxInt32 || result < IntMin {
		return 0
	}

	return result
}
