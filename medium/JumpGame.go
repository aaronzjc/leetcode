package medium

// https://leetcode-cn.com/problems/jump-game/

// TIPS:
// 贪心的思路，寻找下一跳最大。看能不能跳到最后。

func canJump(nums []int) bool {
	var i, j, pre, next int
	ln := len(nums)
	for i < ln {
		if nums[i] == 0 && i != ln-1 {
			return false
		}
		pre = 0
		for j = 0; j <= nums[i]; j++ {
			if i+j >= ln-1 {
				// 可以直接返回true，为了测试覆盖率
				next = ln
				break
			}
			if j+nums[i+j] >= pre {
				pre = j + nums[i+j]
				next = i + j
			}
		}
		i = next
	}
	return true
}
