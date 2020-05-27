package medium

import (
	"math"
)

// https://leetcode-cn.com/problems/divide-two-integers/

// Tips:
// 不能用除法，就用减法不断减。但是每次减去除数太慢了，所以每次，倍增除数。然后剩余的数再走这个流程。
// 例如：
// 29/2 = 29 -2 = 27 - 4 = 23 - 8 = 15 - 16。这一步得出的商是(2^0 + 2^1 + 2^2 + 2^3 + 2^4 = 2^5 - 1)。发现少于除数，然后再将15进行这一步。
// 符合递归的思路。

func divide(dividend int, divisor int) (res int) {
	if divisor == 0 {
		panic("divisor can't be 0")
	}

	var negtive bool
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		negtive = true
	}

	if dividend < 0 {
		dividend -= dividend + dividend
	}

	if divisor < 0 {
		divisor -= divisor + divisor
	}

	if (divisor == 1 || divisor == -1) && (dividend > math.MaxInt32 || dividend < ^math.MaxInt32) {
		if negtive {
			return ^math.MaxInt32
		}

		return math.MaxInt32
	}
	var div func(int, int, int) int
	div = func(num int, d int, res int) int {
		if num < d {
			return res + 0
		}
		if num == d {
			return res + 1
		}
		c := 0
		t := d
		for num >= t {
			num -= t
			t += t
			c++
		}

		res = res + ((1 << c) - 1)

		if num >= d {
			return div(num, d, res)
		}

		return res
	}

	res = div(dividend, divisor, 0)

	if negtive {
		res -= res + res
	}
	return
}
