package hard

// https://leetcode-cn.com/problems/first-missing-positive/

// TIPS:
// 参考的官方题解。hard级别的题目果然不是一般的脑回路。
// 首先，要想到利用hash来求。其次，用符号位作为hash值也比较灵性。
// 还是要多学习，开阔自己的思路。

func firstMissingPositive(nums []int) (res int) {
	contain := 0
	n := len(nums)
	for k, v := range nums {
		if v == 1 {
			contain++
		}
		if v <= 0 || v > n {
			nums[k] = 1
		}
	}
	if contain == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	for i := 0; i < n; i++ {
		v := nums[i]
		if v < 0 {
			v = -1 * v
		}
		if v == n {
			nums[0] = -1 * v
		} else {
			if nums[v] > 0 {
				nums[v] = -1 * nums[v]
			}
		}
	}

	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return n
	}

	return n + 1
}
