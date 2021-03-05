package easy

// https://leetcode-cn.com/problems/maximum-subarray/
// Tips:
// 遍历数组，逐个相加。只有加上当前数字和大于0才继续推进，否则重置。

func maxSubArray(nums []int) int {
	ln := len(nums)

	res := nums[0]
	var tmp int
	for i := 0; i < ln; i++ {
		tmp += nums[i]
		if tmp > res {
			res = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	return res
}
