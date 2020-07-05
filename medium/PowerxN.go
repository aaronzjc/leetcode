package medium

// https://leetcode-cn.com/problems/powx-n/

// TIPS:
//递归。二分法降低计算次数。

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n > 0 {
		if n%2 == 0 {
			return myPow(x*x, n/2)
		}
		return myPow(x, n-1) * x
	}
	return 1.0 / myPow(x, -1*n)
}
