package hard

// https://leetcode-cn.com/problems/jump-game-ii/

// TIPS:
// 贪心算法。每一次往前跳的时候，都选择最优的。对于这个题目，每次跳选择跳完后索引最远的。

func jump(nums []int) (res int) {
	l := len(nums) - 1
	var i, j, max, maxIdx int
	for i <= l {
		if i == l {
			break
		}
		if i+nums[i] >= l {
			res++
			break
		}
		max = -1
		maxIdx = -1
		for j = 0; j <= nums[i]; j++ {
			if nums[i+j]+j >= max {
				max = nums[i+j] + j
				maxIdx = j
			}
		}

		if maxIdx > -1 {
			i = i + maxIdx
			res++
		}
	}
	return
}
