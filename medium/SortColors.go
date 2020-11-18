package medium

// https://leetcode-cn.com/problems/sort-colors/

func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		if nums[i] == 0 {
			nums[i] = nums[p0]
			nums[p0] = 0
			p0++
		}
		if nums[i] == 2 {
			nums[i] = nums[p2]
			nums[p2] = 2
			p2--
			// 换过来的值可能时0或者2，这时候需要重新比较交换一次
			if nums[i] != 1 {
				i--
			}
		}
	}
}
