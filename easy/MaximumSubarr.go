package easy

// https://leetcode-cn.com/problems/maximum-subarray/

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
