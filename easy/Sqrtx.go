package easy

// https://leetcode-cn.com/problems/sqrtx/

func mySqrt(x int) (res int) {
	left := 1
	right := x
	var mid int
	for left <= right {
		mid = (left + right) / 2
		cal := mid * mid
		if cal <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return
}
