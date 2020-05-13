package easy

// https://leetcode-cn.com/problems/reverse-integer/

import "math"

const INT_MAX = 1<<31 - 1
const INT_MIN = ^INT_MAX

func reverse(x int) int {
	var result int
	for x != 0 {
		t := x % 10
		result = result*10 + t
		x = x / 10
	}

	if result > math.MaxInt32 || result < INT_MIN {
		return 0
	}

	return result
}
